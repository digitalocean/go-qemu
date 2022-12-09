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

package gen

import (
	"container/list"
	"flag"
	"go/format"
	"io/ioutil"

	qapischema "github.com/digitalocean/go-qemu/qapi-schema"
)

const specURL = `https://gitlab.com/qemu-project/qemu/-/raw/v7.1.0/qapi/qapi-schema.json`

var (
	inputSpec = flag.String("input", specURL, "Input spec")
	outputGo  = flag.String("output", "", "Generated code output")
	templates = flag.String("templates", "", "Template directory")
)

// Generate parses a QMP API spec and generates a Go implementation.
func Generate() error {
	flag.Parse()

	src, err := getQAPI(*inputSpec)
	if err != nil {
		return err
	}

	tree, err := qapischema.Parse(string(src))
	if err != nil {
		return err
	}

	tree, err = completeParseTree(tree, *inputSpec)
	if err != nil {
		return err
	}

	symbols := lowerParseTree(tree)

	need := neededTypes(symbols)
	bs, err := renderAPI(*templates, symbols, need)
	if err != nil {
		return err
	}
	formatted, err := format.Source(bs)
	if err != nil {
		// Ignore the error, we're just trying to write out the
		// unformatted code to help debugging.
		ioutil.WriteFile(*outputGo, bs, 0640)
		return err
	}

	return ioutil.WriteFile(*outputGo, formatted, 0640)
}

// completeParseTree walks the parse tree and expands the Include statements.
// It returns the completely resolved parse tree.
func completeParseTree(tree *qapischema.Tree, base string) (*qapischema.Tree, error) {
	if inc, ok := tree.Node.(qapischema.Include); ok {
		src, err := resolvePath(base, string(inc))
		if err != nil {
			return nil, err
		}

		// Any include directives in a subtree of this node will use
		// *this* include directive as its main path.
		base = src

		srcBytes, err := getQAPI(src)
		if err != nil {
			return nil, err
		}

		root, err := qapischema.Parse(string(srcBytes))
		if err != nil {
			return nil, err
		}

		tree.Children = root.Children
	}

	for i, subtree := range tree.Children {
		st, err := completeParseTree(subtree, base)
		if err != nil {
			return tree, err
		}

		tree.Children[i] = st
	}

	return tree, nil
}

// lowerParseTree converts the parsed types from package qapischema into
// types that are wired up for Go code generation.
func lowerParseTree(tree *qapischema.Tree) map[name]interface{} {
	types := make(map[name]interface{})

	remaining := list.New()
	remaining.PushFront(tree)

	for remaining.Len() > 0 {
		tree := remaining.Remove(remaining.Front()).(*qapischema.Tree)

		for _, t := range tree.Children {
			remaining.PushBack(t)
		}

		switch node := tree.Node.(type) {
		case qapischema.Root:
			// Nothing to do, subtrees were already added to the queue above.
		case qapischema.Include:
			// Nothing to do, subtrees were already added to the queue above.
		case qapischema.Pragma:
			// Ignored by go-qemu.
		case *qapischema.Struct:
			v := structType{
				Name: name(node.Name),
				Base: name(node.Base),
			}

			for _, member := range node.Members {
				f := field{
					Name:     name(member.Name),
					Optional: member.Optional,
				}

				if member.Type.TypeArray != "" {
					f.Type = name(member.Type.TypeArray)
					f.List = true
				} else {
					f.Type = name(member.Type.Type)
				}
				v.Fields = append(v.Fields, f)
			}

			types[v.Name] = v

		case *qapischema.Command:
			v := command{
				Name: name(node.Name),
			}

			if node.Boxed != nil {
				v.BoxedInput = *node.Boxed
			}

			if node.Data.Ref != nil {
				v.Inputs.Ref = name(*node.Data.Ref)
			} else if node.Data.Embed != nil {
				var fs fields
				for _, member := range node.Data.Embed.Members {
					f := field{
						Name:     name(member.Name),
						Optional: member.Optional,
					}

					if member.Type.TypeArray != "" {
						f.List = true
						f.Type = name(member.Type.TypeArray)
					} else {
						f.Type = name(member.Type.Type)
					}

					fs = append(fs, f)
					v.Inputs.AnonFields = fs
				}
			}

			if node.Returns != nil {
				// TODO: Output.Name is confusing
				if node.Returns.TypeArray != "" {
					v.Output = field{
						Name: name(node.Returns.TypeArray),
						List: true,
						Type: name(node.Returns.TypeArray),
						// TODO: Optional:
					}
				} else {
					v.Output = field{
						Name: name(node.Returns.Type),
						Type: name(node.Returns.Type),
						// TODO: Optional:
					}
				}
			}

			types[v.Name] = v
		case *qapischema.Event:
			v := event{
				Name: name(node.Name),
			}

			if node.Data.Selector.Embed != nil {
				v.Data = fieldsOrRef{
					Ref: name(node.Data.Selector.Ref),
				}
				if v.Data.Ref != "" {
					v.BoxedData = true
				}

				for _, member := range node.Data.Selector.Embed.Members {
					f := field{
						Name:     name(member.Name),
						Optional: member.Optional,
					}

					if member.Type.TypeArray != "" {
						f.List = true
						f.Type = name(member.Type.TypeArray)
					} else {
						f.Type = name(member.Type.Type)
					}

					v.Data.AnonFields = append(v.Data.AnonFields, f)
				}
			}

			types[v.Name] = v
		case *qapischema.Alternate:
			v := alternate{
				Name:    name(node.Name),
				Options: make(map[name]name),
			}

			for _, opt := range node.Data {
				n := opt.Type.Type
				if opt.Type.TypeArray != "" {
					n = string(opt.Type.TypeArray)
				}
				v.Options[name(opt.Name)] = name(n)
			}

			types[v.Name] = v
		case *qapischema.Union:
			if node.Discriminator == "" {
				v := simpleUnion{
					Name:    name(node.Name),
					Options: make(map[name]name),
				}

				for _, branch := range node.Branches {
					b := string(branch.Type.Type)
					if branch.Type.TypeArray != "" {
						b = string(branch.Type.TypeArray)
					}
					v.Options[name(branch.Name)] = name(b)
				}

				types[v.Name] = v
			} else {
				v := flatUnion{
					Name:          name(node.Name),
					Discriminator: name(node.Discriminator),
					Base: fieldsOrRef{
						Ref: name(node.Base.Name),
					},
					Options: make(map[name]name),
				}

				for _, member := range node.Base.Members {
					f := field{
						Name:     name(member.Name),
						Optional: member.Optional,
					}

					if member.Type.TypeArray != "" {
						f.List = true
						f.Type = name(member.Type.TypeArray)
					} else {
						f.Type = name(member.Type.Type)
					}

					v.Base.AnonFields = append(v.Base.AnonFields, f)
				}

				for _, branch := range node.Branches {
					b := string(branch.Type.Type)
					if branch.Type.TypeArray != "" {
						b = string(branch.Type.TypeArray)
					}
					v.Options[name(branch.Name)] = name(b)
				}

				types[v.Name] = v
			}

		case *qapischema.Enum:
			v := enum{
				Name: name(node.Name),
			}

			for _, val := range node.Values {
				v.Values = append(v.Values, name(val.Value))
			}

			types[v.Name] = v
		}
	}

	return types
}
