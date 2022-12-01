# qapi-schema

Package qapi-schema is a fully compliant[^1] QAPI[^2] schema language parser.
The QAPI schema language looks approximately like JSON, but it differs
slightly in many ways, which can confuse a regular JSON parser.

## Usage

If you want to parse QAPI schema from your Go code, all you have to do
is called `qapischema.Parse`:

```go
input := `{ 'struct': 'DiskThing',
            'data': {
                'diskname': {
                    'type':'str',
                    'if':'DISKS_HAVE_NAMES' } } }`
schema, _ := qapischema.Parse(input)
```

The above code snippet would produce a `*qapischema.Tree` that looks like this:

```txt
&qapischema.Tree{
  Node: qapischema.Root{  },
  Children: []*qapischema.Tree{
    {
      Node: &qapischema.Struct{
        Name: "DiskThing",
        Members: []qapischema.Member{
          {
            Name: "diskname",
            Type: qapischema.TypeRef{
              Type: "str",
            },
            If: &qapischema.Cond{
              If: &qapischema.CondIf("DISKS_HAVE_NAMES"),
            },
          },
        },
      },
    },
  },
}
```

Once the QAPI has been parsed, you can walk the `*qapischema.Tree` and do
whatever it is that you need to do.

The `Node` field in `*qapischema.Tree` is an interface type, and so a type
assertion is required to identify and access the QAPI-type-specific data
in the node you're visiting within the tree.

```go
func visit(tree *qapischema.Tree) {
	switch data := tree.Node.(type) {
	// Root node, no data, traverse the subtrees in the .Children field.
	case qapischema.Root:
	case qapischema.Include:
	case qapischema.Pragma:
	case *qapischema.Struct:
	case *qapischema.Union:
	case *qapischema.Event:
	case *qapischema.Command:
	case *qapischema.Alternate:
	}

	// Process the rest of the document
	for _, t := range tree.Children {
		visit(t)
	}
}

func main() {
	tree, _ := qapischema.Parse(input)

	visit(tree)
}
```

## Reporting defects

There is a lot of room for improvement with how this parser emits diagnostic
information. That is, it doesn't emit any at all. The parser will simply stop
parsing when it's not able to parse something. It won't complain, it will just
stop.

So, when it comes to identifying which part of the document the parser did not
understand, just compare the input schema to the output until you find the
first element in the input schema that is missing from the parse tree.

There are two utilities included in this module: `qapilex` and `qapiparse`.

`qapiparse` parses QAPI from its stdin and prints a pretty string representation
of the parse tree to stdout. This can be very helpful for figuring out where the
parser stopped.

In your bug report, please include the QAPI input that surfaced the failure to
parse. If possible, try to reduce the QAPI input down to a minimal viable
reproducer.

## Acknowledgements

Many thanks to:

* Thorsten Ball, the author of _Writing an Interpreter in Go_[^3]. The lessons
  in that book's chapter on lexing were directly applied to create package
  `internal/lex`.
* Jeremy Brown, whose GopherCon 2022 talk[^4] demonstrated simple and elegant
  ways to write parser combinators in Go, which directly inspired much of
  package `internal/parse`.

[^1]: At least, it's intended to be fully compliant. If it is not, please
file a bug.

[^2]: https://qemu.readthedocs.io/en/latest/devel/qapi-code-gen.html#introduction

[^3]: https://interpreterbook.com/

[^4]: [GopherCon 2022: Jeremy Brown - Parsing w/o Code Generators: Parser Combinators in Go with Generics](
https://www.youtube.com/watch?v=x5p_SJNRB4U)
