package fieldtype

import (
	"fmt"

	"github.com/bitrise-io/step-yml-linter/lint"
	"gopkg.in/yaml.v3"
)

// IO ...
type IO struct {
	Key     String
	Value   String
	Options Opts `yaml:"-"`
	l, c    int
	Parent  *String
}

type Opts struct {
	IsExpand          Bool                   `yaml:"is_expand,omitempty"`
	SkipIfEmpty       Bool                   `yaml:"skip_if_empty,omitempty"`
	Title             String                 `yaml:"title,omitempty"`
	Description       String                 `yaml:"description,omitempty"`
	Summary           String                 `yaml:"summary,omitempty"`
	Category          String                 `yaml:"category,omitempty"`
	ValueOptions      []String               `yaml:"value_options,omitempty"`
	IsRequired        Bool                   `yaml:"is_required,omitempty"`
	IsDontChangeValue Bool                   `yaml:"is_dont_change_value,omitempty"`
	IsTemplate        Bool                   `yaml:"is_template,omitempty"`
	IsSensitive       Bool                   `yaml:"is_sensitive,omitempty"`
	Unset             Bool                   `yaml:"unset,omitempty"`
	Meta              map[String]interface{} `yaml:"meta,omitempty"`
	Parent            *String
}

type input struct {
	Options Opts `yaml:"opts,omitempty"`
}

// UnmarshalYAML ...
func (t *IO) UnmarshalYAML(node *yaml.Node) error {
	t.l, t.c = node.Line, node.Column

	var i input

	if err := node.Decode(&i); err != nil {
		return err
	}

	t.Options = i.Options

	var v map[String]yaml.Node

	if err := node.Decode(&v); err != nil {
		return err
	}

	if len(v) != 2 {
		return fmt.Errorf("invalid input map")
	}

	for key, value := range v {
		if key.Value != "opts" {
			t.Key = key
			if err := value.Decode(&t.Value); err != nil {
				return err
			}
		} else {
			t.Options.Parent = &String{}
			*t.Options.Parent = key
		}
	}

	return nil
}

// Line ...
func (t IO) Line() int {
	return t.l
}

// Column ...
func (t IO) Column() int {
	return t.c
}

// Line ...
func (t Opts) Line() int {
	return t.Parent.l
}

// Column ...
func (t Opts) Column() int {
	return t.Parent.c
}

// IsEmpty ...
func (t Opts) IsEmpty() bool {
	if t.Parent == nil {
		return false
	}
	return t.Parent.c == 0 || t.Parent.l == 0
}

// ParentField ...
func (t Opts) ParentField() lint.Field {
	if t.Parent == nil {
		return nil
	}
	return *t.Parent
}

// IsEmpty ...
func (t IO) IsEmpty() bool {
	return t.c == 0 || t.l == 0
}

// ParentField ...
func (t IO) ParentField() lint.Field {
	if t.Parent == nil {
		return nil
	}
	return *t.Parent
}

// FieldName ...
func (t IO) FieldName() string {
	if t := t.Parent; t != nil {
		return t.Value
	}
	return ""
}

// FieldName ...
func (t Opts) FieldName() string {
	if t := t.Parent; t != nil {
		return t.Value
	}
	return ""
}
