package fieldtype

import (
	"github.com/bitrise-io/step-yml-linter/lint"
	"gopkg.in/yaml.v3"
)

// String ...
type String struct {
	Raw    interface{}
	Value  string
	l, c   int
	Parent *String
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

// IsEmpty ...
func (t String) IsEmpty() bool {
	return t.c == 0 || t.l == 0
}

// ParentField ...
func (t String) ParentField() lint.Field {
	if t.Parent == nil {
		return nil
	}
	return *t.Parent
}

// FieldName ...
func (t String) FieldName() string {
	if t := t.Parent; t != nil {
		return t.Value
	}
	return ""
}
