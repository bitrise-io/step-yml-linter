package fieldtype

import (
	"gopkg.in/yaml.v3"
)

// Bool ...
type Bool struct {
	Raw   interface{}
	Value bool
	l, c  int
}

// UnmarshalYAML ...
func (t *Bool) UnmarshalYAML(node *yaml.Node) error {
	t.l, t.c = node.Line, node.Column
	if err := node.Decode(&t.Raw); err != nil {
		return err
	}
	return node.Decode(&t.Value)
}

// Line ...
func (t Bool) Line() int {
	return t.l
}

// Column ...
func (t Bool) Column() int {
	return t.c
}
