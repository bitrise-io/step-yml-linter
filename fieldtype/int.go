package fieldtype

import (
	"gopkg.in/yaml.v3"
)

// Int ...
type Int struct {
	Raw   interface{}
	Value int
	l, c  int
}

// UnmarshalYAML ...
func (t *Int) UnmarshalYAML(node *yaml.Node) error {
	t.l, t.c = node.Line, node.Column
	if err := node.Decode(&t.Raw); err != nil {
		return err
	}
	return node.Decode(&t.Value)
}

// Line ...
func (t Int) Line() int {
	return t.l
}

// Column ...
func (t Int) Column() int {
	return t.c
}
