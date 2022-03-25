package metadata

import (
	"fmt"
	"go/ast"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/emicklei/proto"
)

type Message struct {
	Name                string
	Fields              []*Field
	IsArray             bool
	ElementType         string
	CustomProtoComments []string
	CustomProtoOptions  []string
}

func (m *Message) ProtoAttributes() string {
	var s strings.Builder
	for i, f := range m.Fields {
		s.WriteString(f.Proto(i + 1))
	}
	return s.String()
}

func (m *Message) adjustType(messages map[string]*Message) {
	for _, f := range m.Fields {
		f.Type = adjustType(f.Type, messages)
	}
}

func (m *Message) loadOptions(protoMessage *proto.Message) {
	if protoMessage.Comment != nil {
		m.CustomProtoComments = clearLines(protoMessage.Comment.Lines)
	}
	for _, e := range protoMessage.Elements {
		if opt, ok := e.(*proto.Option); ok {
			m.CustomProtoOptions = append(m.CustomProtoOptions, fmt.Sprintf("    option %s = {", opt.Name))
			m.CustomProtoOptions = append(m.CustomProtoOptions, printProtoLiteral(opt.Constant.OrderedMap, 2)...)
			m.CustomProtoOptions = append(m.CustomProtoOptions, "    };")
		}

		if f, ok := e.(*proto.NormalField); ok {
			for _, field := range m.Fields {
				if ToSnakeCase(field.Name) == f.Name {
					if f.Comment != nil {
						field.CustomProtoComments = clearLines(f.Comment.Lines)
					}
					var hasComplexOption bool
					for i, opt := range f.Options {
						var prefix string
						if i > 0 {
							prefix = "        "
						}
						var suffix string
						if i+1 < len(f.Options) {
							suffix = ", "
						}
						if opt.Constant.Source != "" {
							if hasComplexOption {
								field.CustomProtoOptions = append(field.CustomProtoOptions, fmt.Sprintf("%s%s = %s%s", prefix, opt.Name, opt.Constant.Source, suffix))
							} else {
								field.CustomProtoOptions = append(field.CustomProtoOptions, fmt.Sprintf("%s = %s%s", opt.Name, opt.Constant.Source, suffix))
							}
							continue
						}
						hasComplexOption = true
						field.CustomProtoOptions = append(field.CustomProtoOptions, fmt.Sprintf("%s%s = {\n", prefix, opt.Name))
						field.CustomProtoOptions = append(field.CustomProtoOptions, printProtoLiteral(opt.Constant.OrderedMap, 3)...)
						field.CustomProtoOptions = append(field.CustomProtoOptions, fmt.Sprintf("        }%s", suffix))
					}
					break
				}
			}
		}
	}
}

func createStructMessage(name string, s *ast.StructType) (*Message, error) {
	fields := make([]*Field, 0)
	for _, f := range s.Fields.List {
		if len(f.Names) == 0 || !firstIsUpper(f.Names[0].Name) {
			continue
		}
		typ, err := exprToStr(f.Type)
		if err != nil {
			return nil, err
		}
		fields = append(fields, &Field{Name: f.Names[0].Name, Type: typ})
	}
	return &Message{
		Name:   name,
		Fields: fields,
	}, nil
}

func createArrayMessage(name string, s *ast.ArrayType) (*Message, error) {
	elt, err := exprToStr(s.Elt)
	if err != nil {
		return nil, err
	}
	return &Message{
		Name:        name,
		IsArray:     true,
		ElementType: elt,
	}, nil
}

func createAliasMessage(name string, s *ast.Ident) (*Message, error) {
	str, err := exprToStr(s)
	if err != nil {
		return nil, err
	}
	return &Message{
		Name:        name,
		ElementType: str,
	}, nil
}

func customType(typ string) bool {
	typ = strings.TrimPrefix(typ, "*")
	return firstIsUpper(typ)
}

func firstIsUpper(s string) bool {
	ru, _ := utf8.DecodeRuneInString(s[0:1])
	return unicode.IsUpper(ru)
}

func adjustType(typ string, messages map[string]*Message) string {
	if m, ok := messages[typ]; ok {
		var prefix string
		if m.IsArray {
			prefix = "[]"
		}
		if m.ElementType != "" {
			return prefix + typ + "." + m.ElementType
		}
	}

	return typ
}

func originalAndElementType(typ string) (original, element string) {
	typ = strings.TrimPrefix(typ, "[]")
	t := strings.Split(typ, ".")
	return t[0], strings.Join(t[1:], ".")
}

func (m *Message) HasComplexAttribute() bool {
	for _, f := range m.Fields {
		if customType(f.Type) || strings.HasPrefix(f.Type, "[]") {
			return true
		}
	}

	return false
}

func (m *Message) AdapterToGo(src, dst string) []string {
	res := make([]string, 0)
	for _, f := range m.Fields {
		attrName := UpperFirstCharacter(f.Name)
		res = append(res, bindToGo(src, fmt.Sprintf("%s.%s", dst, attrName), attrName, f.Type, false)...)
	}
	return res
}

func (m *Message) AdapterToProto(src, dst string) []string {
	res := make([]string, 0)
	for _, f := range m.Fields {
		res = append(res, bindToProto(src, dst, UpperFirstCharacter(f.Name), f.Type)...)
	}
	return res
}

func (m *Message) ProtoName() string {
	return regexp.MustCompile("Params$").ReplaceAllString(m.Name, "Request")
}
