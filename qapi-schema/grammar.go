package qapischema

import (
	"strings"

	"github.com/digitalocean/go-qemu/qapi-schema/internal/parse"
	"github.com/digitalocean/go-qemu/qapi-schema/internal/token"
)

var schemaParser = parse.Map(
	parse.ManyOf(topLevelExprParser),
	func(tlds []*Tree) *Tree {
		return &Tree{Node: Root{}, Children: tlds}
	},
)

var topLevelExprParser = parse.OneOf(directiveParser, definitionParser)

var definitionParser = parse.OneOf(
	enumObjParser,
	structObjParser,
	unionObjParser,
	alternateObjParser,
	commandObjParser,
	eventObjParser,
)

var unionObjParser = parse.Map(
	parse.All3(objStartParser, unionParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []unionField, parse.Empty]) *Tree {
		var u Union
		for _, field := range all.Second {
			field(&u)
		}
		return &Tree{Node: &u}
	},
)

var unionParser = parse.Map(
	parse.All3(unionNameParser, commaParser, atLeastOneSeparatedBy(unionFieldParser, commaParser)),
	func(all parse.Seq3[string, parse.Empty, []unionField]) []unionField {
		namer := func(u *Union) { u.Name = all.First }
		return append([]unionField{namer}, all.Third...)
	},
)

var unionNameParser = parse.Map(
	parse.All3(exactStrParser("union"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var unionFieldParser = parse.OneOf(
	unionBaseParser,
	unionDataParser,
	discriminatorParser,
	parse.Map(condKVParser, func(c Cond) unionField {
		return func(u *Union) {
			u.If = &c
		}
	}),
	parse.Map(featuresKVParser, func(fs []Feature) unionField {
		return func(u *Union) {
			u.Features = fs
		}
	}),
)

var unionBaseParser = parse.Map(
	parse.All3(
		exactStrParser("base"),
		colonParser,
		parse.OneOf(
			parse.Map(membersObjParser, func(ms []Member) UnionBase { return UnionBase{Members: ms} }),
			parse.Map(strParser, func(s string) UnionBase { return UnionBase{Name: s} }),
		),
	),
	func(all parse.Seq3[parse.Empty, parse.Empty, UnionBase]) unionField {
		return func(u *Union) {
			u.Base = all.Third
		}
	},
)

var unionDataParser = parse.Map(
	parse.All3(exactStrParser("data"), colonParser, branchesParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, []Branch]) unionField {
		return func(u *Union) {
			u.Branches = all.Third
		}
	},
)

var discriminatorParser = parse.Map(
	parse.All3(exactStrParser("discriminator"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) unionField {
		return func(u *Union) {
			u.Discriminator = all.Third
		}
	},
)

var branchesParser = parse.Map(
	parse.All3(objStartParser, atLeastOneSeparatedBy(branchParser, commaParser), objStopParser),
	func(all parse.Seq3[parse.Empty, []Branch, parse.Empty]) []Branch {
		return all.Second
	},
)

var branchParser = parse.OneOf(extBranchParser, basicBranchParser)

var extBranchParser = parse.Map(
	parse.All3(strParser, colonParser, extBranchObjParser),
	func(all parse.Seq3[string, parse.Empty, []branchField]) Branch {
		br := Branch{Name: all.First}
		for _, field := range all.Third {
			field(&br)
		}
		return br
	},
)

var extBranchObjParser = parse.Map(
	parse.All3(objStartParser, extBranchInteriorParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []branchField, parse.Empty]) []branchField {
		return all.Second
	},
)

var extBranchInteriorParser = atLeastOneSeparatedBy(extBranchInteriorFieldParser, commaParser)

var extBranchInteriorFieldParser = parse.OneOf(
	parse.Map(typeTypeRefParser, func(t TypeRef) branchField {
		return func(b *Branch) { b.Type = t }
	}),
	parse.Map(condKVParser, func(c Cond) branchField {
		return func(b *Branch) { b.If = &c }
	}),
)

var basicBranchParser = parse.Map(
	parse.All3(strParser, colonParser, typeRefParser),
	func(all parse.Seq3[string, parse.Empty, TypeRef]) Branch {
		return Branch{Name: all.First, Type: all.Third}
	},
)

var alternateObjParser = parse.Map(
	parse.All3(objStartParser, alternateParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []alternateField, parse.Empty]) *Tree {
		var alt Alternate
		for _, field := range all.Second {
			field(&alt)
		}
		return &Tree{Node: &alt}
	},
)

var alternateParser = parse.Map(
	parse.All3(alternateNameParser, commaParser, atLeastOneSeparatedBy(alternateFieldParser, commaParser)),
	func(all parse.Seq3[string, parse.Empty, []alternateField]) []alternateField {
		namer := func(a *Alternate) { a.Name = all.First }
		return append([]alternateField{namer}, all.Third...)
	},
)

var alternateNameParser = parse.Map(
	parse.All3(exactStrParser("alternate"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var alternateFieldParser = parse.OneOf(
	parse.Map(alternateDataParser, func(all parse.Seq3[parse.Empty, parse.Empty, []Alternative]) alternateField {
		return func(a *Alternate) { a.Data = all.Third }
	}),
	parse.Map(condKVParser, func(c Cond) alternateField {
		return func(a *Alternate) { a.If = &c }
	}),
	parse.Map(featuresKVParser, func(fs []Feature) alternateField {
		return func(a *Alternate) { a.Features = fs }
	}),
)

var alternateDataParser = parse.All3(exactStrParser("data"), colonParser, alternativesParser)

var alternativesParser = parse.Map(
	parse.All3(objStartParser, atLeastOneSeparatedBy(alternativeParser, commaParser), objStopParser),
	func(all parse.Seq3[parse.Empty, []Alternative, parse.Empty]) []Alternative {
		return all.Second
	},
)

var alternativeParser = parse.OneOf(
	alternativeBasicParser,
	alternativeExtParser,
)

var alternativeExtParser = parse.Map(
	parse.All3(strParser, colonParser, alternativeExtObjParser),
	func(all parse.Seq3[string, parse.Empty, []alternativeField]) Alternative {
		alt := Alternative{Name: all.First}
		for _, field := range all.Third {
			field(&alt)
		}
		return alt
	},
)

var alternativeExtObjParser = parse.Map(
	parse.All3(objStartParser, alternativeExtInteriorParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []alternativeField, parse.Empty]) []alternativeField {
		return all.Second
	},
)

var alternativeExtInteriorParser = atLeastOneSeparatedBy(alternativeExtFieldParser, commaParser)

var alternativeExtFieldParser = parse.OneOf(
	parse.Map(
		parse.All3(exactStrParser("type"), colonParser, strParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, string]) alternativeField {
			return func(a *Alternative) { a.Type = TypeRef{Type: all.Third} }
		},
	),
	parse.Map(condKVParser, func(c Cond) alternativeField {
		return func(a *Alternative) { a.If = &c }
	}),
)

var alternativeBasicParser = parse.Map(
	parse.All3(strParser, colonParser, typeRefParser),
	func(all parse.Seq3[string, parse.Empty, TypeRef]) Alternative {
		return Alternative{Name: all.First, Type: all.Third}
	},
)

var commandObjParser = parse.Map(
	parse.All3(objStartParser, commandParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []commandField, parse.Empty]) *Tree {
		var cmd Command
		for _, field := range all.Second {
			field(&cmd)
		}
		return &Tree{Node: &cmd}
	},
)

var commandParser = parse.Map(
	parse.All2(
		commandNameParser,
		parse.OneOf(
			parse.Map(
				parse.All2(commaParser, atLeastOneSeparatedBy(commandFieldParser, commaParser)),
				func(all parse.Seq2[parse.Empty, []commandField]) []commandField {
					return all.Second
				},
			),
			parse.Map(parse.Nothing(), func(parse.Empty) []commandField {
				return []commandField{func(c *Command) {}}
			}),
		),
	),
	func(all parse.Seq2[string, []commandField]) []commandField {
		namer := func(c *Command) { c.Name = all.First }
		return append([]commandField{namer}, all.Second...)
	},
)

var commandFieldParser = parse.OneOf(
	parse.Map(commandDataUnboxedParser, func(dat CommandDataEmbed) commandField {
		return func(c *Command) { c.Data.Embed = &dat }
	}),
	parse.Map(commandDataKVParser, func(s string) commandField {
		d := CommandDataRef(s)
		return func(c *Command) { c.Data.Ref = &d }
	}),
	parse.Map(commandBoxedBoolKVParser, func(b bool) commandField {
		return func(c *Command) {
			t := true
			c.Boxed = &t
		}
	}),
	parse.Map(coroutineKVParser, func(b bool) commandField {
		return func(c *Command) { c.Coroutine = &b }
	}),
	parse.Map(returnsKVParser, func(t TypeRef) commandField {
		return func(c *Command) { c.Returns = &t }
	}),
	parse.Map(successResponseKVParser, func(b bool) commandField {
		return func(c *Command) { c.SuccessResponse = &b }
	}),
	parse.Map(genKVParser, func(b bool) commandField {
		return func(c *Command) { c.Gen = &b }
	}),
	parse.Map(allowOOBKVParser, func(b bool) commandField {
		return func(c *Command) { c.AllowOOB = &b }
	}),
	parse.Map(allowPreconfigKVParser, func(b bool) commandField {
		return func(c *Command) { c.AllowPreconfig = &b }
	}),
	parse.Map(featuresKVParser, func(fs []Feature) commandField {
		return func(c *Command) { c.Features = fs }
	}),
	parse.Map(condKVParser, func(c Cond) commandField {
		return func(cmd *Command) { cmd.If = &c }
	}),
)

var commandNameParser = parse.Map(
	parse.All3(exactStrParser("command"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var commandDataUnboxedParser = parse.OneOf(
	parse.Map(
		parse.All3(exactStrParser("data"), colonParser, membersObjParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, []Member]) CommandDataEmbed {
			return CommandDataEmbed{Members: all.Third}
		},
	),
)

var commandBoxedBoolKVParser = parse.Map(
	parse.All3(exactStrParser("boxed"), colonParser, boolParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) bool {
		return all.Third
	},
)

var commandDataKVParser = parse.Map(
	parse.All3(exactStrParser("data"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var coroutineKVParser = parse.Map(
	parse.All3(exactStrParser("coroutine"), colonParser, boolParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) bool {
		return all.Third
	},
)

var allowPreconfigKVParser = parse.Map(
	parse.All3(exactStrParser("allow-preconfig"), colonParser, boolParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) bool {
		return all.Third
	},
)

var allowOOBKVParser = parse.Map(
	parse.All3(exactStrParser("allow-oob"), colonParser, boolParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) bool {
		return all.Third
	},
)

var genKVParser = parse.Map(
	parse.All3(exactStrParser("gen"), colonParser, boolParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) bool {
		return all.Third
	},
)

var successResponseKVParser = parse.Map(
	parse.All3(exactStrParser("success-response"), colonParser, boolParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) bool {
		return all.Third
	},
)

var returnsKVParser = parse.Map(
	parse.All3(exactStrParser("returns"), colonParser, typeRefParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, TypeRef]) TypeRef {
		return all.Third
	},
)

var eventObjParser = parse.Map(
	parse.All3(objStartParser, eventParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []eventField, parse.Empty]) *Tree {
		var ev Event
		for _, field := range all.Second {
			field(&ev)
		}
		return &Tree{Node: &ev}
	},
)

var eventParser = parse.OneOf(
	parse.Map(
		parse.All3(eventNameParser, commaParser, atLeastOneSeparatedBy(eventFieldParser, commaParser)),
		func(all parse.Seq3[string, parse.Empty, []eventField]) []eventField {
			namer := func(e *Event) { e.Name = all.First }
			return append([]eventField{namer}, all.Third...)
		},
	),
	parse.Map(
		parse.All2(eventNameParser, parse.Nothing()),
		func(all parse.Seq2[string, parse.Empty]) []eventField {
			namer := func(e *Event) { e.Name = all.First }
			return []eventField{namer}
		},
	),
)

var eventFieldParser = parse.OneOf(
	parse.Map(condKVParser, func(c Cond) eventField {
		return func(e *Event) { e.If = &c }
	}),
	parse.Map(featuresKVParser, func(fs []Feature) eventField {
		return func(e *Event) { e.Features = fs }
	}),
	parse.Map(eventDataUnboxedParser, func(evd EventDataUnboxed) eventField {
		return func(e *Event) { e.Data.Selector.Embed = &evd }
	}),
	parse.Map(eventBoxedKVParser, func(evd eventDataBoxedTrue) eventField {
		return func(e *Event) { e.Data.Boxed = true }
	}),
	parse.Map(eventDataFlatParser, func(s string) eventField {
		return func(e *Event) { e.Data.Selector.Ref = s }
	}),
)

var eventDataUnboxedParser = parse.Map(
	parse.All3(
		exactStrParser("data"), colonParser, parse.OneOf(
			parse.Map(membersObjParser, func(ms []Member) EventDataUnboxed {
				return EventDataUnboxed{Members: ms}
			}),
			parse.Map(strParser, func(s string) EventDataUnboxed {
				return EventDataUnboxed{Type: s}
			}),
		),
	),
	func(all parse.Seq3[parse.Empty, parse.Empty, EventDataUnboxed]) EventDataUnboxed {
		return all.Third
	},
)

var eventBoxedKVParser = parse.Map(
	parse.All3(exactStrParser("boxed"), colonParser, trueParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, bool]) eventDataBoxedTrue {
		return eventDataBoxedTrue{}
	},
)

var eventDataFlatParser = parse.Map(
	parse.All3(exactStrParser("data"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var eventNameParser = parse.Map(
	parse.All3(exactStrParser("event"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var typeTypeRefParser = parse.Map(
	parse.All3(exactStrParser("type"), colonParser, typeRefParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, TypeRef]) TypeRef {
		return all.Third
	},
)

var structObjParser = parse.Map(
	parse.All3(objStartParser, structParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []structField, parse.Empty]) *Tree {
		var st Struct
		for _, field := range all.Second {
			field(&st)
		}
		return &Tree{Node: &st}
	},
)

var structParser = parse.Map(
	parse.All3(structNameParser, commaParser, atLeastOneSeparatedBy(structFieldParser, commaParser)),
	func(all parse.Seq3[string, parse.Empty, []structField]) []structField {
		namer := func(s *Struct) { s.Name = all.First }
		return append([]structField{namer}, all.Third...)
	},
)

var structFieldParser = parse.OneOf(
	parse.Map(structDataParser, func(ms []Member) structField {
		return func(s *Struct) { s.Members = ms }
	}),
	parse.Map(baseParser, func(b string) structField {
		return func(s *Struct) { s.Base = b }
	}),
	parse.Map(featuresKVParser, func(fs []Feature) structField {
		return func(s *Struct) { s.Features = fs }
	}),
	parse.Map(condKVParser, func(c Cond) structField {
		return func(s *Struct) { s.If = &c }
	}),
)

var structNameParser = parse.Map(
	parse.All3(exactStrParser("struct"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var structDataParser = parse.Map(
	parse.All3(exactStrParser("data"), colonParser, membersObjParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, []Member]) []Member {
		return all.Third
	},
)

var membersObjParser = parse.Map(
	parse.All3(objStartParser, zeroOrMoreSeparatedBy(memberParser, commaParser), objStopParser),
	func(all parse.Seq3[parse.Empty, []Member, parse.Empty]) []Member {
		return all.Second
	},
)

var baseParser = parse.Map(
	parse.All3(exactStrParser("base"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string { return all.Third },
)

var memberParser = parse.Map(
	parse.OneOf(basicMemberParser, extMemberParser),
	func(m Member) Member {
		trimmed := strings.TrimPrefix(m.Name, "*")
		if trimmed != m.Name {
			m.Name = trimmed
			m.Optional = true
		}
		return m
	},
)

var extMemberParser = parse.Map(
	parse.All3(strParser, colonParser, extMemberObjParser),
	func(all parse.Seq3[string, parse.Empty, parse.Seq2[TypeRef, []memberField]]) Member {
		m := Member{
			Name: all.First,
			Type: all.Third.First,
		}

		for _, field := range all.Third.Second {
			field(&m)
		}

		return m
	},
)

var extMemberObjParser = parse.Map(
	parse.All3(objStartParser, extMemberInteriorParser, objStopParser),
	func(all parse.Seq3[parse.Empty, parse.Seq2[TypeRef, []memberField], parse.Empty]) parse.Seq2[TypeRef, []memberField] {
		return all.Second
	},
)

var extMemberInteriorParser = parse.OneOf(
	parse.Map(
		parse.All3(typeTypeRefParser, commaParser, atLeastOneSeparatedBy(optionalMemberFieldParser, commaParser)),
		func(all parse.Seq3[TypeRef, parse.Empty, []memberField]) parse.Seq2[TypeRef, []memberField] {
			var r parse.Seq2[TypeRef, []memberField]
			r.First = all.First
			r.Second = all.Third
			return r
		},
	),
	parse.Map(typeTypeRefParser, func(t TypeRef) parse.Seq2[TypeRef, []memberField] {
		var r parse.Seq2[TypeRef, []memberField]
		r.First = t
		r.Second = nil
		return r
	}),
)

var optionalMemberFieldParser = parse.OneOf(
	parse.Map(featuresKVParser, func(fs []Feature) memberField {
		return func(m *Member) { m.Features = fs }
	}),
	parse.Map(condKVParser, func(c Cond) memberField {
		return func(m *Member) { m.If = &c }
	}),
)

var basicMemberParser = parse.Map(
	parse.All3(strParser, colonParser, typeRefParser),
	func(all parse.Seq3[string, parse.Empty, TypeRef]) Member {
		return Member{Name: all.First, Type: all.Third}
	},
)

var enumObjParser = parse.Map(
	parse.All3(objStartParser, enumParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []enumField, parse.Empty]) *Tree {
		var en Enum
		for _, field := range all.Second {
			field(&en)
		}
		return &Tree{Node: &en}
	},
)

var enumParser = parse.Map(
	parse.All3(enumNameParser, commaParser, atLeastOneSeparatedBy(enumFieldParser, commaParser)),
	func(all parse.Seq3[string, parse.Empty, []enumField]) []enumField {
		namer := func(e *Enum) { e.Name = all.First }
		return append([]enumField{namer}, all.Third...)
	},
)

var enumFieldParser = parse.OneOf(
	parse.Map(enumDataParser, func(vals []EnumValue) enumField {
		return func(e *Enum) { e.Values = vals }
	}),
	parse.Map(prefixParser, func(s string) enumField {
		return func(e *Enum) { e.Prefix = s }
	}),
	parse.Map(condKVParser, func(c Cond) enumField {
		return func(e *Enum) { e.If = &c }
	}),
	parse.Map(featuresKVParser, func(fs []Feature) enumField {
		return func(e *Enum) { e.Features = fs }
	}),
)

var enumNameParser = parse.Map(
	parse.All3(exactStrParser("enum"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var enumDataParser = parse.Map(
	parse.All3(exactStrParser("data"), colonParser, enumValuesParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, []EnumValue]) []EnumValue {
		return all.Third
	},
)

var featuresKVParser = parse.Map(
	parse.All3(exactStrParser("features"), colonParser, featuresParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, []Feature]) []Feature {
		return all.Third
	},
)

var condKVParser = parse.Map(
	parse.All3(exactStrParser("if"), colonParser, condParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, Cond]) Cond {
		return all.Third
	},
)

var prefixParser = parse.Map(
	parse.All3(exactStrParser("prefix"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var enumValuesParser = parse.Map(
	parse.All3(
		arrayStartParser,
		atLeastOneSeparatedBy(enumValueParser, commaParser),
		arrayStopParser,
	),
	func(all parse.Seq3[parse.Empty, []EnumValue, parse.Empty]) []EnumValue {
		return all.Second
	},
)

var enumValueParser = parse.OneOf(
	parse.Map(strParser, func(s string) EnumValue { return EnumValue{Value: s} }),
	enumValueObjParser,
)

var enumValueObjParser = parse.Map(
	parse.All3(
		objStartParser,
		atLeastOneSeparatedBy(enumValueInteriorFieldParser, commaParser),
		objStopParser,
	),
	func(all parse.Seq3[parse.Empty, []enumValueField, parse.Empty]) EnumValue {
		var ev EnumValue
		for _, field := range all.Second {
			field(&ev)
		}
		return ev
	},
)

var enumValueInteriorFieldParser = parse.OneOf(
	parse.Map(
		parse.All3(exactStrParser("name"), colonParser, strParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, string]) enumValueField {
			return func(e *EnumValue) { e.Value = all.Third }
		},
	),
	parse.Map(condKVParser, func(c Cond) enumValueField {
		return func(e *EnumValue) { e.Cond = &c }
	}),
	parse.Map(featuresKVParser, func(fs []Feature) enumValueField {
		return func(e *EnumValue) { e.Features = fs }
	}),
)

var typeRefParser = parse.OneOf(
	parse.Map(strParser, func(s string) TypeRef { return TypeRef{Type: s} }),
	arrayTypeParser,
)

var arrayTypeParser = parse.Map(
	parse.All3(arrayStartParser, strParser, arrayStopParser),
	func(all parse.Seq3[parse.Empty, string, parse.Empty]) TypeRef {
		return TypeRef{TypeArray: TypeRefArray(all.Second)}
	},
)

var featuresParser = parse.Map(
	parse.All3(arrayStartParser, atLeastOneSeparatedBy(featureParser, commaParser), arrayStopParser),
	func(all parse.Seq3[parse.Empty, []Feature, parse.Empty]) []Feature {
		return all.Second
	},
)

var featureParser = parse.OneOf(
	parse.Map(strParser, func(s string) Feature { return Feature{Name: s} }),
	extFeatureObjParser,
)

var extFeatureObjParser = parse.Map(
	parse.All3(objStartParser, atLeastOneSeparatedBy(extFeatureInteriorFieldParser, commaParser), objStopParser),
	func(all parse.Seq3[parse.Empty, []featureField, parse.Empty]) Feature {
		var f Feature
		for _, field := range all.Second {
			field(&f)
		}
		return f
	},
)

var extFeatureInteriorFieldParser = parse.OneOf(
	parse.Map(
		parse.All3(exactStrParser("name"), colonParser, strParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, string]) featureField {
			return func(f *Feature) {
				f.Name = all.Third
			}
		},
	),
	parse.Map(condKVParser, func(c Cond) featureField {
		return func(f *Feature) {
			f.Cond = c
		}
	}),
)

var condParser = parse.OneOf(
	parse.Map(strParser, func(s string) Cond {
		cif := CondIf(s)
		return Cond{If: &cif}
	}),
	allCondParser,
	anyCondParser,
	notCondParser,
)

var allCondParser = parse.Map(
	parse.All3(objStartParser, allCondInteriorParser, objStopParser),
	func(all parse.Seq3[parse.Empty, *CondAll, parse.Empty]) Cond {
		return Cond{All: all.Second}
	},
)

var allCondInteriorParser = parse.Map(
	parse.All3(exactStrParser("all"), colonParser, strArrayParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, []string]) *CondAll {
		v := CondAll(all.Third)
		return &v
	},
)

var anyCondParser = parse.Map(
	parse.All3(objStartParser, anyCondInteriorParser, objStopParser),
	func(all parse.Seq3[parse.Empty, *CondAny, parse.Empty]) Cond {
		return Cond{Any: all.Second}
	},
)

var anyCondInteriorParser = parse.Map(
	parse.All3(exactStrParser("any"), colonParser, strArrayParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, []string]) *CondAny {
		v := CondAny(all.Third)
		return &v
	},
)

var notCondParser = parse.Map(
	parse.All3(objStartParser, notCondInteriorParser, objStopParser),
	func(all parse.Seq3[parse.Empty, *CondNot, parse.Empty]) Cond {
		return Cond{Not: all.Second}
	},
)

var notCondInteriorParser = parse.Map(
	parse.All3(exactStrParser("not"), colonParser, strParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) *CondNot {
		v := CondNot(all.Third)
		return &v
	},
)

var directiveParser = parse.OneOf(includeObjParser, pragmaObjParser)

var pragmaObjParser = parse.Map(
	parse.All3(objStartParser, pragmaKeyAndDataParser, objStopParser),
	func(all parse.Seq3[parse.Empty, Pragma, parse.Empty]) *Tree {
		return &Tree{Node: all.Second}
	},
)

var pragmaKeyAndDataParser = parse.Map(
	parse.All3(exactStrParser("pragma"), colonParser, pragmaInteriorParser),
	func(all parse.Seq3[parse.Empty, parse.Empty, Pragma]) Pragma {
		return all.Third
	},
)

var pragmaInteriorParser = parse.Map(
	parse.All3(objStartParser, pragmaFieldsParser, objStopParser),
	func(all parse.Seq3[parse.Empty, []pragmaField, parse.Empty]) Pragma {
		var p Pragma
		for _, field := range all.Second {
			field(&p)
		}
		return p
	},
)

var pragmaFieldsParser = parse.Map(
	parse.ManyOf(atLeastOneSeparatedBy(pragmaFieldParser, commaParser)),
	func(slices [][]pragmaField) []pragmaField {
		var flattened []pragmaField
		for _, sl := range slices {
			flattened = append(flattened, sl...)
		}
		return flattened
	},
)

var pragmaFieldParser = parse.OneOf(
	parse.Map(
		parse.All3(exactStrParser("doc-required"), colonParser, boolParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, bool]) pragmaField {
			return func(p *Pragma) { p.DocRequired = all.Third }
		},
	),
	parse.Map(
		parse.All3(exactStrParser("command-name-exceptions"), colonParser, strArrayParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, []string]) pragmaField {
			return func(p *Pragma) { p.CommandNameExceptions = all.Third }
		},
	),
	parse.Map(
		parse.All3(exactStrParser("command-returns-exceptions"), colonParser, strArrayParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, []string]) pragmaField {
			return func(p *Pragma) { p.CommandReturnExceptions = all.Third }
		},
	),
	parse.Map(
		parse.All3(exactStrParser("member-name-exceptions"), colonParser, strArrayParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, []string]) pragmaField {
			return func(p *Pragma) { p.MemberNameExceptions = all.Third }
		},
	),

	// For backwards compatibility with very old QEMU releases
	parse.Map(
		parse.All3(exactStrParser("returns-whitelist"), colonParser, strArrayParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, []string]) pragmaField {
			return func(p *Pragma) { p.CommandReturnExceptions = all.Third }
		},
	),
	parse.Map(
		parse.All3(exactStrParser("name-case-whitelist"), colonParser, strArrayParser),
		func(all parse.Seq3[parse.Empty, parse.Empty, []string]) pragmaField {
			return func(p *Pragma) { p.CommandNameExceptions = all.Third }
		},
	),
)

var includeObjParser = parse.Map(
	parse.All3(objStartParser, includeParser, objStopParser),
	func(all parse.Seq3[parse.Empty, string, parse.Empty]) *Tree {
		return &Tree{Node: Include(all.Second)}
	},
)

var includeParser = parse.Map(
	parse.All3(
		exactStrParser("include"),
		colonParser,
		strParser,
	),
	func(all parse.Seq3[parse.Empty, parse.Empty, string]) string {
		return all.Third
	},
)

var objStartParser = parse.Exactly(token.Token{Type: token.LeftCurly, Literal: "{"})

var objStopParser = parse.Exactly(token.Token{Type: token.RightCurly, Literal: "}"})

var strArrayParser = parse.Map(
	parse.All3(arrayStartParser, atLeastOneStrParser, arrayStopParser),
	func(all parse.Seq3[parse.Empty, []string, parse.Empty]) []string {
		return all.Second
	},
)

var arrayStartParser = parse.Exactly(token.Token{Type: token.LeftSquare, Literal: "["})

var arrayStopParser = parse.Exactly(token.Token{Type: token.RightSquare, Literal: "]"})

func atLeastOneSeparatedBy[T any, U any](elemParser parse.Parser[T], sepParser parse.Parser[U]) parse.Parser[[]T] {
	return parse.OneOf(
		parse.Map(
			parse.All2(
				parse.ConsumeAtLeastOne(
					parse.Map(
						parse.All2(elemParser, sepParser),
						func(all parse.Seq2[T, U]) T { return all.First },
					),
				),
				parse.Map(elemParser, func(elem T) []T { return []T{elem} }),
			),
			func(all parse.Seq2[[]T, []T]) []T { return append(all.First, all.Second...) },
		),
		parse.Map(elemParser, func(elem T) []T { return []T{elem} }),
	)
}

func zeroOrMoreSeparatedBy[T any, U any](elemParser parse.Parser[T], sepParser parse.Parser[U]) parse.Parser[[]T] {
	return parse.OneOf(
		atLeastOneSeparatedBy(elemParser, sepParser),
		parse.Map(parse.Nothing(), func(parse.Empty) []T { return nil }),
	)
}

var atLeastOneStrParser = atLeastOneSeparatedBy(strParser, commaParser)

var colonParser = parse.Exactly(token.Token{Type: token.Colon, Literal: ":"})

var commaParser = parse.Exactly(token.Token{Type: token.Comma, Literal: ","})

func exactStrParser(exactly string) parse.Parser[parse.Empty] {
	return parse.Map(
		parse.All3(
			quoteParser,
			parse.Exactly(token.Token{Type: token.AlphaNumeric, Literal: exactly}),
			quoteParser,
		),
		func(_ parse.Seq3[parse.Empty, parse.Empty, parse.Empty]) parse.Empty {
			return parse.Empty{}
		},
	)
}

var strParser = parse.Map(
	parse.All3(quoteParser, alphaNumParser, quoteParser),
	func(all parse.Seq3[parse.Empty, string, parse.Empty]) string {
		return all.Second
	},
)

var alphaNumParser = parse.Map(
	parse.MatchesToken(func(t token.Token) bool { return t.Type == token.AlphaNumeric }),
	func(t token.Token) string { return t.Literal },
)

var quoteParser = parse.Exactly(token.Token{Type: token.SingleQuote, Literal: "'"})

var boolParser = parse.OneOf(trueParser, falseParser)

var trueParser = parse.Map(
	parse.Exactly(token.Token{Type: token.AlphaNumeric, Literal: "true"}),
	func(parse.Empty) bool {
		return true
	})

var falseParser = parse.Map(
	parse.Exactly(token.Token{Type: token.AlphaNumeric, Literal: "false"}),
	func(parse.Empty) bool {
		return false
	})
