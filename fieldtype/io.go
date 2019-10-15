package fieldtype

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// IO ...
type IO struct {
	Key     String
	Value   String
	Options opts `yaml:"-"`
	l, c    int
}

type opts struct {
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
}

type input struct {
	Options opts `yaml:"opts,omitempty"`
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
