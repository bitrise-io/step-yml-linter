package fieldtype

import (
	"gopkg.in/yaml.v3"
)

// String ...
type String struct {
	Raw   interface{}
	Value string
	l, c  int
}

// UnmarshalYAML ...
func (t *String) UnmarshalYAML(node *yaml.Node) error {
	t.l, t.c = node.Line, node.Column
	if err := node.Decode(&t.Raw); err != nil {
		return err
	}
	return node.Decode(&t.Value)
}

// Line ...
func (t String) Line() int {
	return t.l
}

// Column ...
func (t String) Column() int {
	return t.c
}
