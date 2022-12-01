package qapischema

type alternateField func(a *Alternate)
type alternativeField func(a *Alternative)
type featureField func(f *Feature)
type branchField func(b *Branch)
type commandField func(c *Command)
type enumField func(e *Enum)
type enumValueField func(e *EnumValue)
type eventField func(e *Event)
type memberField func(m *Member)
type pragmaField func(p *Pragma)
type structField func(s *Struct)
type unionField func(u *Union)

type eventDataBoxedTrue struct{}
