package qapischema

import (
	"errors"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/lex"
	"github.com/digitalocean/go-qemu/qapi-schema/internal/parse"
)

// Parse parses a QAPI input document and returns its parse tree.
func Parse(input string) (*Tree, error) {
	tokens := lex.Lex(input)

	schema, err := parse.Parse(schemaParser, tokens)
	if errors.Is(err, parse.ErrNonExhaustiveParse) {
		err = nil
	}
	return schema, err
}

// Node is a marker interface that indentifies the implementing
// concrete type as a QAPI type.
type Node interface {
	QAPINode()
}

// Tree is a QAPI parse tree.
type Tree struct {
	Node     Node
	Children []*Tree
}

// Root is the root of the QAPI document. On its own, it is
// unremarkable. Client code will want to visit its subtrees.
type Root struct{}

// QAPINode marks Root as a QAPI node.
func (r Root) QAPINode() {}

// Include is a QAPI node that refers to another QAPI document
// that contains dependent types.
type Include string

// QAPINode marks Include as a QAPI node.
func (i Include) QAPINode() {}

// Pragma is a QAPI node that parameterizes the upstream QEMU
// QAPI code generator.
type Pragma struct {
	DocRequired             bool
	CommandNameExceptions   []string
	CommandReturnExceptions []string
	MemberNameExceptions    []string
}

// QAPINode marks Pragma as a QAPI node.
func (p Pragma) QAPINode() {}

// Cond expresses the conditions under which the upstream QEMU
// code generator will include a type or field.
type Cond struct {
	If  *CondIf
	All *CondAll
	Any *CondAny
	Not *CondNot
}

// CondIf will result in code generation if the specified config
// option is set.
type CondIf string

// CondNot will result in code generation if the specified config
// option is not set.
type CondNot string

// CondAll will result in code generation if all of the specified
// config options are set.
type CondAll []string

// CondAny will result in code generation if any of the specified
// config options are set.
type CondAny []string

// Feature specifies a QAPI document feature.
type Feature struct {
	Name string
	Cond Cond
}

// TypeRef refers to a QAPI type or an array of a QAPI type.
// If it is an array type, then TypeArray will be set and Type
// will be the empty string and vice versa.
type TypeRef struct {
	Type      string
	TypeArray TypeRefArray
}

// TypeRefArray is a stronger Go type for expressing a string
// that is meant to represent an array of a QAPI type.
type TypeRefArray string

// Enum is a QAPI type whose value is only one of many defined
// variants.
type Enum struct {
	Name     string
	Values   []EnumValue
	Prefix   string
	If       *Cond
	Features []Feature
}

// QAPINode marks Enum as a QAPI node.
func (e *Enum) QAPINode() {}

// EnumValue is one possible variant for a QAPI enum.
type EnumValue struct {
	Value    string
	Cond     *Cond
	Features []Feature
}

// Struct is a QAPI type that is a combined record of many
// other types.
type Struct struct {
	Name     string
	Members  []Member
	Base     string
	If       *Cond
	Features []Feature
}

// QAPINode marks Struct as a QAPI node.
func (s *Struct) QAPINode() {}

// Member is a field type for a QAPI object.
type Member struct {
	Name     string
	Type     TypeRef
	If       *Cond
	Features []Feature
	Optional bool
}

// Union is a QAPI type that is like a sum-type enum.
type Union struct {
	Name          string
	Base          UnionBase
	Discriminator string
	Branches      []Branch
	If            *Cond
	Features      []Feature
}

// QAPINode marks Union as a QAPI node.
func (u *Union) QAPINode() {}

// UnionBase is the base object of the union.
type UnionBase struct {
	Name    string
	Members []Member
}

// Branch describes one of the active variants for the Union type.
type Branch struct {
	Name string
	Type TypeRef
	If   *Cond
}

// Event is a QAPI type that describes something that has happened.
type Event struct {
	Name     string
	Data     EventData
	If       *Cond
	Features []Feature
}

// QAPINode marks Event as a QAPI node.
func (e *Event) QAPINode() {}

// EventData contains an Event's data.
type EventData struct {
	Boxed    bool
	Selector EventDataSelector
}

// EventDataSelector contains the name of the event's boxed type
// or its own fields that comprise the event.
type EventDataSelector struct {
	Ref   string
	Embed *EventDataUnboxed
}

// EventDataUnboxed contains the event's own fields, rather than
// embedding another type's fields as its fields.
type EventDataUnboxed struct {
	Type    string
	Members []Member
}

// Command is a QAPI type that describes a QMP command that can be
// sent to the QEMU process.
type Command struct {
	Name            string
	Data            CommandData
	Returns         *TypeRef
	SuccessResponse *bool
	Gen             *bool
	AllowOOB        *bool
	AllowPreconfig  *bool
	Coroutine       *bool
	Boxed           *bool
	If              *Cond
	Features        []Feature
}

// QAPINode marks Command as a QAPI node.
func (c *Command) QAPINode() {}

// CommandData contains the Command's fields.
type CommandData struct {
	Embed *CommandDataEmbed
	Ref   *CommandDataRef
}

// CommandDataRef refers to another QAPI type that should be
// used as the argument for this Command.
type CommandDataRef string

// CommandDataEmbed refers to the fields to embed in the
// Command.
type CommandDataEmbed struct {
	Members []Member
}

// Alternate is a QAPI type that describes what types can
// be suitable alternates.
type Alternate struct {
	Name     string
	Data     []Alternative
	If       *Cond
	Features []Feature
}

// QAPINode marks Alternate as a QAPI node.
func (a *Alternate) QAPINode() {}

// Alternative is one such alternative that could be contained
// in an Alternate.
type Alternative struct {
	Name string
	Type TypeRef
	If   *Cond
}
