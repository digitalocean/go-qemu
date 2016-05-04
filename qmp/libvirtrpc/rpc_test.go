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
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/davecgh/go-xdr/xdr2"
	"github.com/digitalocean/go-qemu/qmp"
)

var (
	// dc229f87d4de47198cfd2e21c6105b01
	testUUID = [uuidSize]byte{
		0xdc, 0x22, 0x9f, 0x87, 0xd4, 0xde, 0x47, 0x19,
		0x8c, 0xfd, 0x2e, 0x21, 0xc6, 0x10, 0x5b, 0x01,
	}
)

var testHeader = []byte{
	0x20, 0x00, 0x80, 0x86, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x01, // procedure
	0x00, 0x00, 0x00, 0x00, // type
	0x00, 0x00, 0x00, 0x01, // serial
	0x00, 0x00, 0x00, 0x00, // status
}

var testEventHeader = []byte{
	0x00, 0x00, 0x00, 0xb0, // length
	0x20, 0x00, 0x80, 0x87, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x06, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x02, // serial
	0x00, 0x00, 0x00, 0x00, // status
}

var testEvent = []byte{
	0x00, 0x00, 0x00, 0x01, // callback id

	// domain name ("test")
	0x00, 0x00, 0x00, 0x04, 0x74, 0x65, 0x73, 0x74,

	// uuid (dc229f87d4de47198cfd2e21c6105b01)
	0xdc, 0x22, 0x9f, 0x87, 0xd4, 0xde, 0x47, 0x19,
	0x8c, 0xfd, 0x2e, 0x21, 0xc6, 0x10, 0x5b, 0x01,

	// domain id (14)
	0x00, 0x00, 0x00, 0x0e,

	// event name (BLOCK_JOB_COMPLETED)
	0x00, 0x00, 0x00, 0x13, 0x42, 0x4c, 0x4f, 0x43,
	0x4b, 0x5f, 0x4a, 0x4f, 0x42, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x00,

	// seconds (1462211891)
	0x00, 0x00, 0x00, 0x00, 0x57, 0x27, 0x95, 0x33,

	// microseconds (931791)
	0x00, 0x0e, 0x37, 0xcf,

	// event json data
	// ({"device":"drive-ide0-0-0","len":0,"offset":0,"speed":0,"type":"commit"})
	0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x48,
	0x7b, 0x22, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x3a, 0x22, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2d, 0x69, 0x64, 0x65, 0x30, 0x2d, 0x30, 0x2d,
	0x30, 0x22, 0x2c, 0x22, 0x6c, 0x65, 0x6e, 0x22,
	0x3a, 0x30, 0x2c, 0x22, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x22, 0x3a, 0x30, 0x2c, 0x22, 0x73,
	0x70, 0x65, 0x65, 0x64, 0x22, 0x3a, 0x30, 0x2c,
	0x22, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3a, 0x22,
	0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x22, 0x7d,
}

var testDomainResponse = []byte{
	0x00, 0x00, 0x00, 0x38, // length
	0x20, 0x00, 0x80, 0x86, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x17, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x01, // serial
	0x00, 0x00, 0x00, 0x00, // status

	// domain name ("test")
	0x00, 0x00, 0x00, 0x04, 0x74, 0x65, 0x73, 0x74,

	// uuid (dc229f87d4de47198cfd2e21c6105b01)
	0xdc, 0x22, 0x9f, 0x87, 0xd4, 0xde, 0x47, 0x19,
	0x8c, 0xfd, 0x2e, 0x21, 0xc6, 0x10, 0x5b, 0x01,

	// domain id (14)
	0x00, 0x00, 0x00, 0x0e,
}

var testRegisterEvent = []byte{
	0x00, 0x00, 0x00, 0x20, // length
	0x20, 0x00, 0x80, 0x87, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x04, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x02, // serial
	0x00, 0x00, 0x00, 0x00, // status
	0x00, 0x00, 0x00, 0x01, // callback id
}

var testDeregisterEvent = []byte{
	0x00, 0x00, 0x00, 0x1c, // length
	0x20, 0x00, 0x80, 0x87, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x05, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x01, // serial
	0x00, 0x00, 0x00, 0x00, // status
}

var testAuthReply = []byte{
	0x00, 0x00, 0x00, 0x1c, // length
	0x20, 0x00, 0x80, 0x86, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x42, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x01, // serial
	0x00, 0x00, 0x00, 0x00, // status
}

var testConnectReply = []byte{
	0x00, 0x00, 0x00, 0x1c, // length
	0x20, 0x00, 0x80, 0x86, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x01, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x02, // serial
	0x00, 0x00, 0x00, 0x00, // status
}

var testDisconnectReply = []byte{
	0x00, 0x00, 0x00, 0x1c, // length
	0x20, 0x00, 0x80, 0x86, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x02, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x01, // serial
	0x00, 0x00, 0x00, 0x00, // status
}

var testRunReply = []byte{
	0x00, 0x00, 0x00, 0x74, // length
	0x20, 0x00, 0x80, 0x87, // program
	0x00, 0x00, 0x00, 0x01, // version
	0x00, 0x00, 0x00, 0x01, // procedure
	0x00, 0x00, 0x00, 0x01, // type
	0x00, 0x00, 0x00, 0x02, // serial
	0x00, 0x00, 0x00, 0x00, // status

	// {"return":{"qemu":{"micro":1,"minor":5,"major":2},"package":""},"id":"libvirt-53"}
	0x00, 0x00, 0x00, 0x52, 0x7b, 0x22, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x22, 0x3a, 0x7b, 0x22,
	0x71, 0x65, 0x6d, 0x75, 0x22, 0x3a, 0x7b, 0x22,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x22, 0x3a, 0x31,
	0x2c, 0x22, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x22,
	0x3a, 0x35, 0x2c, 0x22, 0x6d, 0x61, 0x6a, 0x6f,
	0x72, 0x22, 0x3a, 0x32, 0x7d, 0x2c, 0x22, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x22, 0x3a,
	0x22, 0x22, 0x7d, 0x2c, 0x22, 0x69, 0x64, 0x22,
	0x3a, 0x22, 0x6c, 0x69, 0x62, 0x76, 0x69, 0x72,
	0x74, 0x2d, 0x35, 0x33, 0x22, 0x7d, 0x00, 0x00,
}

var testErrorMessage = []byte{
	0x00, 0x00, 0x00, 0x37, // code
	0x00, 0x00, 0x00, 0x0a, // domain id

	// message ("Requested operation is not valid: domain is not running")
	0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x37,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x20, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x6e,
	0x6f, 0x74, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x3a, 0x20, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20,
	0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x00,

	// error level
	0x00, 0x00, 0x00, 0x02,
}

var testDomain = domain{
	Name: "test-domain",
	UUID: testUUID,
	ID:   1,
}

type mockLibvirt struct {
	net.Conn
	test net.Conn
}

func setupTest(t *testing.T) *mockLibvirt {
	serv, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}

	conn, err := net.Dial("tcp", serv.Addr().String())
	if err != nil {
		log.Fatal(err)
	}

	m := &mockLibvirt{Conn: conn}

	go func() {
		for {
			conn, _ := serv.Accept()
			m.test = conn
			go m.handle(m.test)
		}
	}()

	return m
}

func (m *mockLibvirt) handle(conn net.Conn) {
	for {
		buf := make([]byte, packetLengthSize+headerSize)
		conn.Read(buf)

		// extract program
		prog := binary.BigEndian.Uint32(buf[4:8])

		// extract procedure
		proc := binary.BigEndian.Uint32(buf[12:16])

		switch prog {
		case programRemote:
			switch proc {
			case procAuthList:
				conn.Write(testAuthReply)
			case procConnectOpen:
				conn.Write(testConnectReply)
			case procConnectClose:
				conn.Write(testDisconnectReply)
			case procDomainLookupByName:
				conn.Write(testDomainResponse)
			}
		case programQEMU:
			switch proc {
			case qemuConnectDomainMonitorEventRegister:
				conn.Write(testRegisterEvent)
			case qemuConnectDomainMonitorEventDeregister:
				conn.Write(testDeregisterEvent)
			case qemuDomainMonitor:
				conn.Write(testRunReply)
			}
		}
	}
}

func TestNew(t *testing.T) {
	domain := "test-1"

	conn := setupTest(t)
	rpc := New(domain, conn)

	if rpc.Domain != domain {
		t.Errorf("expected domain %q, got %q", domain, rpc.Domain)
	}
}

func TestQMPEvent(t *testing.T) {
	e := &event{
		CallbackID:   1,
		Domain:       testDomain,
		Event:        "test",
		Seconds:      1,
		Microseconds: 1,
		Details:      []byte(`{"device":"drive-virtio-disk0","len":0,"offset":0,"speed":0,"type":"commit"}`),
	}

	qe, err := qmpEvent(e)
	if err != nil {
		t.Error(err)
	}

	if qe.Event != e.Event {
		t.Errorf("expected event %q, got %q", e.Event, qe.Event)
	}

	if qe.Timestamp.Seconds != int64(e.Seconds) {
		t.Errorf("expected seconds to be %q, got %q", e.Seconds, qe.Timestamp.Seconds)
	}

	if qe.Timestamp.Microseconds != int64(e.Microseconds) {
		t.Errorf("expected microseconds to be %q, got %q", e.Microseconds, qe.Timestamp.Microseconds)
	}

	expected := "drive-virtio-disk0"
	actual, ok := qe.Data["device"]
	if !ok {
		t.Error("expected event data 'device'")
	}

	if actual != expected {
		t.Errorf("expected device %q, got %q", expected, actual)
	}
}

func TestExtractHeader(t *testing.T) {
	r := bytes.NewBuffer(testHeader)
	h, err := extractHeader(r)
	if err != nil {
		t.Error(err)
	}

	if h.Program != programRemote {
		t.Errorf("expected Program %q, got %q", programRemote, h.Program)
	}

	if h.Version != programVersion {
		t.Errorf("expected version %q, got %q", programVersion, h.Version)
	}

	if h.Procedure != procConnectOpen {
		t.Errorf("expected procedure %q, got %q", procConnectOpen, h.Procedure)
	}

	if h.Type != Call {
		t.Errorf("expected type %q, got %q", Call, h.Type)
	}

	if h.Status != StatusOK {
		t.Errorf("expected status %q, got %q", StatusOK, h.Status)
	}
}

func TestPktLen(t *testing.T) {
	data := []byte{0x00, 0x00, 0x00, 0xa} // uint32:10
	r := bytes.NewBuffer(data)

	expected := uint32(10)
	actual, err := pktlen(r)
	if err != nil {
		t.Error(err)
	}

	if expected != actual {
		t.Errorf("expected packet length %q, got %q", expected, actual)
	}
}

func TestDecodeEvent(t *testing.T) {
	e, err := decodeEvent(testEvent)
	if err != nil {
		t.Error(err)
	}

	expCbID := uint32(1)
	if e.CallbackID != expCbID {
		t.Errorf("expected callback id %d, got %d", expCbID, e.CallbackID)
	}

	expName := "test"
	if e.Domain.Name != expName {
		t.Errorf("expected domain %s, got %s", expName, e.Domain.Name)
	}

	expUUID := testUUID
	if !bytes.Equal(e.Domain.UUID[:], expUUID[:]) {
		t.Errorf("expected uuid:\t%x, got\n\t\t\t%x", expUUID, e.Domain.UUID)
	}

	expID := 14
	if e.Domain.ID != expID {
		t.Errorf("expected id %d, got %d", expID, e.Domain.ID)
	}

	expEvent := "BLOCK_JOB_COMPLETED"
	if e.Event != expEvent {
		t.Errorf("expected %s, got %s", expEvent, e.Event)
	}

	expSec := uint64(1462211891)
	if e.Seconds != expSec {
		t.Errorf("expected seconds to be %d, got %d", expSec, e.Seconds)
	}

	expMs := uint32(931791)
	if e.Microseconds != expMs {
		t.Errorf("expected microseconds to be %d, got %d", expMs, e.Microseconds)
	}

	expDetails := []byte(`{"device":"drive-ide0-0-0","len":0,"offset":0,"speed":0,"type":"commit"}`)
	if e.Domain.ID != expID {
		t.Errorf("expected data %s, got %s", expDetails, e.Details)
	}
}

func TestDecodeError(t *testing.T) {
	expected := "Requested operation is not valid: domain is not running"

	err := decodeError(testErrorMessage)
	if err.Error() != expected {
		t.Errorf("expected error %s, got %s", expected, err.Error())
	}
}

func TestEncode(t *testing.T) {
	data := "test"
	buf, err := encode(data)
	if err != nil {
		t.Error(err)
	}

	dec := xdr.NewDecoder(bytes.NewReader(buf.Bytes()))
	res, _, err := dec.DecodeString()
	if err != nil {
		t.Error(err)
	}

	if res != data {
		t.Errorf("expected %s, got %s", data, res)
	}
}

func TestRegister(t *testing.T) {
	rpc := &Monitor{}
	rpc.callbacks = make(map[uint32]chan response)
	id := uint32(1)
	c := make(chan response)

	rpc.register(id, c)
	if _, ok := rpc.callbacks[id]; !ok {
		t.Error("expected callback to register")
	}
}

func TestDeregister(t *testing.T) {
	id := uint32(1)

	rpc := &Monitor{}
	rpc.callbacks = map[uint32]chan response{
		id: make(chan response),
	}

	rpc.deregister(id)
	if _, ok := rpc.callbacks[id]; ok {
		t.Error("expected callback to deregister")
	}
}

func TestAddStream(t *testing.T) {
	id := uint32(1)
	c := make(chan *event)

	rpc := &Monitor{}
	rpc.events = make(map[uint32]chan *event)

	rpc.addStream(id, c)
	if _, ok := rpc.events[id]; !ok {
		t.Error("expected event stream to exist")
	}
}

func TestRemoveStream(t *testing.T) {
	id := uint32(1)

	conn := setupTest(t)
	rpc := New("test", conn)
	rpc.events[id] = make(chan *event)

	err := rpc.removeStream(id)
	if err != nil {
		t.Error(err)
	}

	if _, ok := rpc.events[id]; ok {
		t.Error("expected event stream to be removed")
	}
}

func TestStream(t *testing.T) {
	id := uint32(1)
	c := make(chan *event, 1)

	rpc := &Monitor{}
	rpc.events = map[uint32]chan *event{
		id: c,
	}

	rpc.stream(testEvent)
	e := <-c

	if e.Event != "BLOCK_JOB_COMPLETED" {
		t.Error("expected event")
	}
}

func TestSerial(t *testing.T) {
	count := uint32(10)
	rpc := &Monitor{}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			rpc.serial()
			wg.Done()
		}()
	}

	wg.Wait()

	expected := count + uint32(1)
	actual := rpc.serial()
	if expected != actual {
		t.Errorf("expected serial to be %d, got %d", expected, actual)
	}
}

func TestLookup(t *testing.T) {
	id := uint32(1)
	c := make(chan response)
	name := "test"

	conn := setupTest(t)
	rpc := New(name, conn)

	rpc.register(id, c)

	d, err := rpc.lookup(name)
	if err != nil {
		t.Error(err)
	}

	if d == nil {
		t.Error("nil domain returned")
	}

	if d.Name != name {
		t.Errorf("expected domain %s, got %s", name, d.Name)
	}
}

func TestConnect(t *testing.T) {
	conn := setupTest(t)
	rpc := New("test", conn)

	err := rpc.Connect()
	if err != nil {
		t.Error(err)
	}
}

func TestDisconnect(t *testing.T) {
	conn := setupTest(t)
	rpc := New("test", conn)

	err := rpc.Disconnect()
	if err != nil {
		t.Error(err)
	}
}

func TestRun(t *testing.T) {
	conn := setupTest(t)
	rpc := New("test", conn)

	res, err := rpc.Run([]byte(`{"query-version"}`))
	if err != nil {
		t.Error(err)
	}

	type version struct {
		Return struct {
			Package string `json:"package"`
			QEMU    struct {
				Major int `json:"major"`
				Micro int `json:"micro"`
				Minor int `json:"minor"`
			} `json:"qemu"`
		} `json:"return"`
	}

	var v version
	err = json.Unmarshal(res, &v)
	if err != nil {
		t.Error(err)
	}

	expected := 2
	if v.Return.QEMU.Major != expected {
		t.Errorf("expected qemu major version %d, got %d", expected, v.Return.QEMU.Major)
	}
}

func TestEvents(t *testing.T) {
	conn := setupTest(t)
	rpc := New("test", conn)

	stream, err := rpc.Events()
	if err != nil {
		t.Error(err)
	}

	// send an event
	go func() {
		conn.test.Write(append(testEventHeader, testEvent...))
	}()

	var e qmp.Event
	select {
	case e = <-stream:
	case <-time.After(time.Second * 1):
		t.Error("expected event, received timeout")
	}

	expected := "drive-ide0-0-0"
	if e.Data["device"] != expected {
		t.Errorf("expected device %q, got %q", expected, e.Data["device"])
	}
}
