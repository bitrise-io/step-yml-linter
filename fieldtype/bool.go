package fieldtype

import (
	"github.com/bitrise-io/step-yml-linter/lint"
	"gopkg.in/yaml.v3"
)

// Bool ...
type Bool struct {
	Raw    interface{}
	Value  bool
	l, c   int
	Parent *String
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

// IsEmpty ...
func (t Bool) IsEmpty() bool {
	return t.c == 0 || t.l == 0
}

// ParentField ...
func (t Bool) ParentField() lint.Field {
	if t.Parent == nil {
		return nil
	}
	return *t.Parent
}

// FieldValue ...
func (t Bool) FieldValue() string {
	if t := t.Parent; t != nil {
		return t.Value
	}
	return ""
}
