package fieldtype

import (
	"fmt"

	"github.com/bitrise-io/step-yml-linter/lint"
	"gopkg.in/yaml.v3"
)

// Title ...
type Title struct {
	Raw          interface{}
	Value        *string
	Line, Column int
}

// UnmarshalYAML ...
func (t *Title) UnmarshalYAML(node *yaml.Node) error {
	var val *string

	if err := node.Decode(&val); err != nil {
		return err
	}

	t.Value, t.Line, t.Column = val, node.Line, node.Column

	var msgs []string
	for _, check := range titleChecks {
		if err := check(val); err != nil {
			msgs = append(msgs, err.Error())
		}
	}

	if len(msgs) > 0 {
		lint.Warnings = append(lint.Warnings, lint.Warning{Messages: msgs, Line: node.Line, Column: node.Column})
	}

	return nil
}

var titleChecks = []func(*string) error{
	Length,
}

// Length returns error if title is not set or empty
func Length(value *string) error {
	if value == nil || len(*value) == 0 {
		return fmt.Errorf("title cannot be empty")
	}
	return nil
}
