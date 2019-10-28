package fieldtype

import (
	"github.com/bitrise-io/step-yml-linter/lint"
	"gopkg.in/yaml.v3"
)

// Int ...
type Int struct {
	Raw    interface{}
	Value  int
	l, c   int
	Parent *String
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

// IsEmpty ...
func (t Int) IsEmpty() bool {
	return t.c == 0 || t.l == 0
}

// ParentField ...
func (t Int) ParentField() lint.Field {
	if t.Parent == nil {
		return nil
	}
	return *t.Parent
}

// FieldValue ...
func (t Int) FieldValue() string {
	if t := t.Parent; t != nil {
		return t.Value
	}
	return ""
}
