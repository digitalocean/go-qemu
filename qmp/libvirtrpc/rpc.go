// Copyright 2016 The go-qemu Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package libvirtrpc is a pure Go implementation of the libvirt RPC protocol.
// For more information on the protocol, see https://libvirt.org/internals/rpc.html
package libvirtrpc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
	"sync"
	"sync/atomic"

	"github.com/davecgh/go-xdr/xdr2"
	"github.com/digitalocean/go-qemu/qmp"
)

var _ qmp.Monitor = &Monitor{}

// request and response types
const (
	// Call is used when making calls to the remote server.
	Call = iota

	// Reply indicates a server reply.
	Reply

	// Message is an asynchronous notification.
	Message

	// Stream represents a stream data packet.
	Stream

	// CallWithFDs is used by a client to indicate the request has
	// arguments with file descriptors.
	CallWithFDs

	// ReplyWithFDs is used by a server to indicate the request has
	// arguments with file descriptors.
	ReplyWithFDs
)

// request and response statuses
const (
	// StatusOK is always set for method calls or events.
	// For replies it indicates successful completion of the method.
	// For streams it indicates confirmation of the end of file on the stream.
	StatusOK = iota

	// StatusError for replies indicates that the method call failed
	// and error information is being returned. For streams this indicates
	// that not all data was sent and the stream has aborted.
	StatusError

	// StatusContinue is only used for streams.
	// This indicates that further data packets will be following.
	StatusContinue
)

// magic program numbers
// see: https://libvirt.org/git/?p=libvirt.git;a=blob_plain;f=src/remote/remote_protocol.x;hb=HEAD
const (
	programVersion   = 1
	programRemote    = 0x20008086
	programQEMU      = 0x20008087
	programKeepAlive = 0x6b656570
)

// libvirt procedure identifiers
const (
	procConnectOpen        = 1
	procConnectClose       = 2
	procDomainLookupByName = 23
	procAuthList           = 66
)

// qemu procedure identifiers
const (
	qemuDomainMonitor                       = 1
	qemuConnectDomainMonitorEventRegister   = 4
	qemuConnectDomainMonitorEventDeregister = 5
	qemuDomainMonitorEvent                  = 6
)

const (
	// packet length, in bytes.
	packetLengthSize = 4

	// packet header, in bytes.
	headerSize = 24

	// UUID size, in bytes.
	uuidSize = 16
)

// header is a libvirt rpc packet header
type header struct {
	// Program identifier
	Program uint32

	// Program version
	Version uint32

	// Remote procedure identifier
	Procedure uint32

	// Call type, e.g., Reply
	Type uint32

	// Call serial number
	Serial uint32

	// Request status, e.g., StatusOK
	Status uint32
}

// packet represents a RPC request or response.
type packet struct {
	// Size of packet, in bytes, including length.
	// Len + Header + Payload
	Len    uint32
	Header header
}

// libvirt's domain response type
type domain struct {
	Name string
	UUID [uuidSize]byte
	ID   int
}

// libvirt domain event
type event struct {
	CallbackID   uint32
	Domain       domain
	Event        string
	Seconds      uint64
	Microseconds uint32
	Padding      uint8
	Details      []byte
}

// internal rpc response
type response struct {
	Payload []byte
	Status  uint32
}

// libvirt error response
type libvirtError struct {
	Code     uint32
	DomainID uint32
	Padding  uint8
	Message  string
	Level    uint32
}

// Monitor implements LibVirt's remote procedure call protocol.
type Monitor struct {
	// Domain name as seen by libvirt, e.g., stage-lb-1
	Domain string

	conn net.Conn
	r    *bufio.Reader
	w    *bufio.Writer

	// method callbacks
	cm        sync.Mutex
	callbacks map[uint32]chan response

	// event listeners
	em     sync.Mutex
	events map[uint32]chan *event

	// next request serial number
	s uint32
}

// New configures a new RPC Monitor connection.
// The provided domain should be the name of the domain as seen
// by libvirt, e.g., stage-lb-1.
func New(domain string, conn net.Conn) *Monitor {
	l := &Monitor{
		Domain:    domain,
		conn:      conn,
		s:         0,
		r:         bufio.NewReader(conn),
		w:         bufio.NewWriter(conn),
		callbacks: make(map[uint32]chan response),
		events:    make(map[uint32]chan *event),
	}

	go l.listen()

	return l
}

// listen processes incoming data and routes the
// decoded responses to their respective handler.
func (rpc *Monitor) listen() {
	for {
		// response packet length
		length, err := pktlen(rpc.r)
		if err != nil {
			if err == io.EOF {
				return
			}

			// invalid packet
			continue
		}

		// response header
		h, err := extractHeader(rpc.r)
		if err != nil {
			// invalid packet
			continue
		}

		// payload: packet length minus what was previously read
		size := int(length) - (packetLengthSize + headerSize)
		buf := make([]byte, size)
		for n := 0; n < size; {
			nn, err := rpc.r.Read(buf)
			if err != nil {
				// invalid packet
				continue
			}

			n += nn
		}

		// route response to caller
		rpc.route(h, buf)
	}
}

// Connect establishes communication with the libvirt server.
// The underlying libvirt socket connection must be previously established.
func (rpc *Monitor) Connect() error {
	payload := struct {
		Padding [3]byte
		Name    string
		Flags   uint32
	}{
		// This is what came of a stupid amount of tcpdump.
		// Why this is needed remains a mystery.
		Padding: [3]byte{0x1, 0x0, 0x0},
		Name:    "qemu:///system",
		Flags:   0,
	}

	buf, err := encode(&payload)
	if err != nil {
		return err
	}

	// libvirt requires that we call auth-list prior to connecting,
	// event when no authentication is used.
	resp, err := rpc.request(procAuthList, programRemote, &buf)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Status != StatusOK {
		return decodeError(r.Payload)
	}

	resp, err = rpc.request(procConnectOpen, programRemote, &buf)
	if err != nil {
		return err
	}

	r = <-resp
	if r.Status != StatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Disconnect shuts down communication with the libvirt server.
// The underlying net.Conn is not closed.
func (rpc *Monitor) Disconnect() error {
	// close event streams
	for id := range rpc.events {
		if err := rpc.removeStream(id); err != nil {
			return err
		}
	}

	// inform libvirt we're done
	resp, err := rpc.request(procConnectClose, programRemote, nil)
	if err != nil {
		return err
	}

	r := <-resp
	if r.Status != StatusOK {
		return decodeError(r.Payload)
	}

	return nil
}

// Events streams QEMU QMP Events.
// If a problem is encountered setting up the event monitor connection
// an error will be returned. Errors encountered during streaming will
// cause the returned event channel to be closed.
func (rpc *Monitor) Events() (<-chan qmp.Event, error) {
	d, err := rpc.lookup(rpc.Domain)
	if err != nil {
		return nil, err
	}

	payload := struct {
		Padding [4]byte
		Domain  domain
		Event   [2]byte
		Flags   [2]byte
	}{
		Padding: [4]byte{0x0, 0x0, 0x1, 0x0},
		Domain:  *d,
		Event:   [2]byte{0x0, 0x0},
		Flags:   [2]byte{0x0, 0x0},
	}

	buf, err := encode(&payload)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.request(qemuConnectDomainMonitorEventRegister, programQEMU, &buf)
	if err != nil {
		return nil, err
	}

	res := <-resp
	dec := xdr.NewDecoder(bytes.NewReader(res.Payload))

	cbID, _, err := dec.DecodeUint()
	if err != nil {
		return nil, err
	}

	c := make(chan qmp.Event)
	go func() {
		stream := make(chan *event)
		rpc.addStream(cbID, stream)

		// process events
		for e := range stream {
			qe, err := qmpEvent(e)
			if err != nil {
				close(c)
				break
			}

			c <- *qe
		}
	}()

	return c, nil
}

// Run executes the given QAPI command against a domain's QEMU instance.
// For a list of available QAPI commands, see:
//	http://git.qemu.org/?p=qemu.git;a=blob;f=qapi-schema.json;hb=HEAD
func (rpc *Monitor) Run(cmd []byte) ([]byte, error) {
	d, err := rpc.lookup(rpc.Domain)
	if err != nil {
		return nil, err
	}

	payload := struct {
		Domain  domain
		Command []byte
		Flags   uint32
	}{
		Domain:  *d,
		Command: cmd,
		Flags:   0,
	}

	buf, err := encode(&payload)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.request(qemuDomainMonitor, programQEMU, &buf)
	if err != nil {
		return nil, err
	}

	res := <-resp
	if res.Status != StatusOK {
		return nil, decodeError(res.Payload)
	}

	r := bytes.NewReader(res.Payload)
	dec := xdr.NewDecoder(r)
	data, _, err := dec.DecodeFixedOpaque(int32(r.Len()))
	if err != nil {
		return nil, err
	}

	// drop QMP control characters from start and end of line
	return data[4 : len(data)-2], err
}

// callback sends rpc responses to their respective caller.
func (rpc *Monitor) callback(id uint32, res response) {
	c, ok := rpc.callbacks[id]
	if ok {
		c <- res
	}

	rpc.deregister(id)
}

// route sends incoming packets to their listeners.
func (rpc *Monitor) route(h *header, buf []byte) {
	// route events to their respective listener
	if h.Program == programQEMU && h.Procedure == qemuDomainMonitorEvent {
		rpc.stream(buf)
		return
	}

	// send responses to caller
	res := response{
		Payload: buf,
		Status:  h.Status,
	}
	rpc.callback(h.Serial, res)
}

// serial provides atomic access to the next sequential request serial number.
func (rpc *Monitor) serial() uint32 {
	return atomic.AddUint32(&rpc.s, 1)
}

// stream decodes domain events and sends them
// to the respective event listener.
func (rpc *Monitor) stream(buf []byte) {
	e, err := decodeEvent(buf)
	if err != nil {
		// event was malformed, drop.
		return
	}

	// send to event listener
	if c, ok := rpc.events[e.CallbackID]; ok {
		c <- e
	}
}

// addStream configures the routing for an event stream.
func (rpc *Monitor) addStream(id uint32, stream chan *event) {
	rpc.em.Lock()
	rpc.events[id] = stream
	rpc.em.Unlock()
}

// removeStream notifies the libvirt server to stop sending events
// for the provided callback id. Upon successful de-registration the
// callback handler is destroyed.
func (rpc *Monitor) removeStream(id uint32) error {
	close(rpc.events[id])

	payload := struct {
		CallbackID uint32
	}{
		CallbackID: id,
	}

	buf, err := encode(&payload)
	if err != nil {
		return err
	}

	resp, err := rpc.request(qemuConnectDomainMonitorEventDeregister, programQEMU, &buf)
	if err != nil {
		return err
	}

	res := <-resp
	if res.Status != StatusOK {
		return decodeError(res.Payload)
	}

	rpc.em.Lock()
	delete(rpc.events, id)
	rpc.em.Unlock()

	return nil
}

// register configures a method response callback
func (rpc *Monitor) register(id uint32, c chan response) {
	rpc.cm.Lock()
	rpc.callbacks[id] = c
	rpc.cm.Unlock()
}

// deregister destroys a method response callback
func (rpc *Monitor) deregister(id uint32) {
	close(rpc.callbacks[id])

	rpc.cm.Lock()
	delete(rpc.callbacks, id)
	rpc.cm.Unlock()
}

// request performs a libvirt RPC request.
// The returned channel is used by the caller to receive the asynchronous
// call response. The channel is closed once a response has been sent.
func (rpc *Monitor) request(proc uint32, program uint32, payload *bytes.Buffer) (<-chan response, error) {
	serial := rpc.serial()
	c := make(chan response)

	rpc.register(serial, c)

	size := packetLengthSize + headerSize
	if payload != nil {
		size += payload.Len()
	}

	p := packet{
		Len: uint32(size),
		Header: header{
			Program:   program,
			Version:   programVersion,
			Procedure: proc,
			Type:      Call,
			Serial:    serial,
			Status:    StatusOK,
		},
	}

	// write header
	err := binary.Write(rpc.w, binary.BigEndian, p)
	if err != nil {
		return nil, err
	}

	// write payload
	if payload != nil {
		err = binary.Write(rpc.w, binary.BigEndian, payload.Bytes())
		if err != nil {
			return nil, err
		}
	}

	if err := rpc.w.Flush(); err != nil {
		return nil, err
	}

	return c, nil
}

// lookup returns a domain as seen by libvirt.
func (rpc *Monitor) lookup(name string) (*domain, error) {
	payload := struct {
		Name string
	}{name}

	buf, err := encode(&payload)
	if err != nil {
		return nil, err
	}

	resp, err := rpc.request(procDomainLookupByName, programRemote, &buf)
	if err != nil {
		return nil, err
	}

	r := <-resp
	if r.Status != StatusOK {
		return nil, decodeError(r.Payload)
	}

	dec := xdr.NewDecoder(bytes.NewReader(r.Payload))

	var d domain
	_, err = dec.Decode(&d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// encode XDR encodes the provided data.
func encode(data interface{}) (bytes.Buffer, error) {
	var buf bytes.Buffer
	_, err := xdr.Marshal(&buf, data)

	return buf, err
}

// decodeError extracts an error message from the provider buffer.
func decodeError(buf []byte) error {
	var e libvirtError

	dec := xdr.NewDecoder(bytes.NewReader(buf))
	_, err := dec.Decode(&e)
	if err != nil {
		return err
	}

	return errors.New(e.Message)
}

// decodeEvent extracts an event from the given byte slice.
// Errors encountered will be returned along with a nil event.
func decodeEvent(buf []byte) (*event, error) {
	var e event

	dec := xdr.NewDecoder(bytes.NewReader(buf))
	_, err := dec.Decode(&e)
	if err != nil {
		return nil, err
	}

	return &e, nil
}

// pktlen determines the length of an incoming rpc response.
// If an error is encountered reading the provided Reader, the
// error is returned and response length will be 0.
func pktlen(r io.Reader) (uint32, error) {
	buf := make([]byte, packetLengthSize)

	for n := 0; n < cap(buf); {
		nn, err := r.Read(buf)
		if err != nil {
			return 0, err
		}

		n += nn
	}

	return binary.BigEndian.Uint32(buf), nil
}

// extractHeader returns the decoded header from an incoming response.
func extractHeader(r io.Reader) (*header, error) {
	buf := make([]byte, headerSize)

	for n := 0; n < cap(buf); {
		nn, err := r.Read(buf)
		if err != nil {
			return nil, err
		}

		n += nn
	}

	h := &header{
		Program:   binary.BigEndian.Uint32(buf[0:4]),
		Version:   binary.BigEndian.Uint32(buf[4:8]),
		Procedure: binary.BigEndian.Uint32(buf[8:12]),
		Type:      binary.BigEndian.Uint32(buf[12:16]),
		Serial:    binary.BigEndian.Uint32(buf[16:20]),
		Status:    binary.BigEndian.Uint32(buf[20:24]),
	}

	return h, nil
}

// qmpEvent takes a libvirt event structure and returns the qmp equivalent.
func qmpEvent(e *event) (*qmp.Event, error) {
	var qe qmp.Event

	if e.Details != nil {
		if err := json.Unmarshal(e.Details, &qe.Data); err != nil {
			return nil, err
		}
	}

	qe.Event = e.Event
	qe.Timestamp.Seconds = int64(e.Seconds)
	qe.Timestamp.Microseconds = int64(e.Microseconds)

	return &qe, nil
}
