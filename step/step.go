package step

import (
	"reflect"
	"strings"
	"time"

	"github.com/bitrise-io/step-yml-linter/fieldtype"
	"gopkg.in/yaml.v3"
)

// YML ...
type YML struct {
	Title         fieldtype.String `yaml:"title,omitempty"`
	Summary       fieldtype.String `yaml:"summary,omitempty"`
	Description   fieldtype.String `yaml:"description,omitempty"`
	Website       fieldtype.String `yaml:"website,omitempty"`
	SourceCodeURL fieldtype.String `yaml:"source_code_url,omitempty"`
	SupportURL    fieldtype.String `yaml:"support_url,omitempty"`
	PublishedAt   time.Time        `yaml:"published_at,omitempty"`
	Source        struct {
		Git    fieldtype.String `yaml:"git,omitempty"`
		Commit fieldtype.String `yaml:"commit,omitempty"`
	} `yaml:"source,omitempty"`
	AssetURLs       map[fieldtype.String]fieldtype.String `yaml:"asset_urls,omitempty"`
	HostOsTags      []fieldtype.String                    `yaml:"host_os_tags,omitempty"`
	ProjectTypeTags []fieldtype.String                    `yaml:"project_type_tags,omitempty"`
	TypeTags        []fieldtype.String                    `yaml:"type_tags,omitempty"`
	Dependencies    []struct {
		Manager fieldtype.String `yaml:"manager,omitempty"`
		Name    fieldtype.String `yaml:"name,omitempty"`
	} `yaml:"dependencies,omitempty"`
	Toolkit struct {
		Bash struct {
			EntryFile fieldtype.String `yaml:"entry_file,omitempty"`
		} `yaml:"bash,omitempty"`
		Go struct {
			PackageName fieldtype.String `yaml:"package_name"`
		} `yaml:"go,omitempty"`
	} `yaml:"toolkit,omitempty"`
	Deps struct {
		Brew []struct {
			Name    fieldtype.String `yaml:"name,omitempty"`
			BinName fieldtype.String `yaml:"bin_name,omitempty"`
		} `yaml:"brew,omitempty"`
		AptGet []struct {
			Name    fieldtype.String `yaml:"name,omitempty"`
			BinName fieldtype.String `yaml:"bin_name,omitempty"`
		} `yaml:"apt_get,omitempty"`
		CheckOnly []struct {
			Name fieldtype.String `yaml:"name,omitempty"`
		} `yaml:"check_only,omitempty"`
	} `yaml:"deps,omitempty"`
	IsRequiresAdminUser fieldtype.Bool                   `yaml:"is_requires_admin_user,omitempty"`
	IsAlwaysRun         fieldtype.Bool                   `yaml:"is_always_run,omitempty"`
	IsSkippable         fieldtype.Bool                   `yaml:"is_skippable,omitempty"`
	RunIf               fieldtype.String                 `yaml:"run_if,omitempty"`
	Timeout             fieldtype.Int                    `yaml:"timeout,omitempty"`
	Meta                map[fieldtype.String]interface{} `yaml:"meta,omitempty"`
	Inputs              []fieldtype.IO                   `yaml:"inputs,omitempty"`
	Outputs             []fieldtype.IO                   `yaml:"outputs,omitempty"`
}

// UnmarshalYAML ...
func (y *YML) UnmarshalYAML(node *yaml.Node) error {
	var fields map[fieldtype.String]yaml.Node
	if err := node.Decode(&fields); err != nil {
		return err
	}

	var yml = &YML{}

	ymlValue := reflect.ValueOf(yml).Elem()

	for i := 0; i < ymlValue.NumField(); i++ {
		tag, ok := ymlValue.Type().Field(i).Tag.Lookup("yaml")
		if !ok {
			continue
		}

		fieldName := strings.Split(tag, ",")[0]

		for key, val := range fields {
			if key.Value == fieldName {
				switch ymlValue.Field(i).Interface().(type) {
				case fieldtype.String:
					var s fieldtype.String
					if err := val.Decode(&s); err != nil {
						return err
					}
					s.Parent = &fieldtype.String{}
					*s.Parent = key
					ymlValue.Field(i).Set(reflect.ValueOf(s))
				case fieldtype.Int:
					var s fieldtype.Int
					if err := val.Decode(&s); err != nil {
						return err
					}
					s.Parent = &fieldtype.String{}
					*s.Parent = key
					ymlValue.Field(i).Set(reflect.ValueOf(s))
				case fieldtype.Bool:
					var s fieldtype.Bool
					if err := val.Decode(&s); err != nil {
						return err
					}
					s.Parent = &fieldtype.String{}
					*s.Parent = key
					ymlValue.Field(i).Set(reflect.ValueOf(s))
				case []fieldtype.IO:
					var s []fieldtype.IO
					if err := val.Decode(&s); err != nil {
						return err
					}
					for ioi := range s {
						s[ioi].Parent = &fieldtype.String{}
						*s[ioi].Parent = key

						s[ioi].Key.Parent = &fieldtype.String{}
						*s[ioi].Key.Parent = key

						s[ioi].Value.Parent = &fieldtype.String{}
						*s[ioi].Value.Parent = s[ioi].Key

						//
						// s[ioi].Options.Parent = &fieldtype.String{}
						// *s[ioi].Options.Parent = key

						s[ioi].Options.Parent.Parent = &fieldtype.String{}
						*s[ioi].Options.Parent.Parent = s[ioi].Key

						s[ioi].Options.Title.Parent = &fieldtype.String{}
						*s[ioi].Options.Title.Parent = *s[ioi].Options.Parent

						s[ioi].Options.Summary.Parent = &fieldtype.String{}
						*s[ioi].Options.Summary.Parent = *s[ioi].Options.Parent

						s[ioi].Options.Description.Parent = &fieldtype.String{}
						*s[ioi].Options.Description.Parent = *s[ioi].Options.Parent

						s[ioi].Options.IsDontChangeValue.Parent = &fieldtype.String{}
						*s[ioi].Options.IsDontChangeValue.Parent = *s[ioi].Options.Parent

						s[ioi].Options.IsExpand.Parent = &fieldtype.String{}
						*s[ioi].Options.IsExpand.Parent = *s[ioi].Options.Parent

						s[ioi].Options.IsRequired.Parent = &fieldtype.String{}
						*s[ioi].Options.IsRequired.Parent = *s[ioi].Options.Parent

						s[ioi].Options.IsSensitive.Parent = &fieldtype.String{}
						*s[ioi].Options.IsSensitive.Parent = *s[ioi].Options.Parent

						s[ioi].Options.IsTemplate.Parent = &fieldtype.String{}
						*s[ioi].Options.IsTemplate.Parent = *s[ioi].Options.Parent

						// s[ioi].Options.Meta.Parent = &fieldtype.String{}
						// *s[ioi].Options.Meta.Parent = *s[ioi].Options.Parent

						s[ioi].Options.SkipIfEmpty.Parent = &fieldtype.String{}
						*s[ioi].Options.SkipIfEmpty.Parent = *s[ioi].Options.Parent

						s[ioi].Options.Unset.Parent = &fieldtype.String{}
						*s[ioi].Options.Unset.Parent = *s[ioi].Options.Parent

						// s[ioi].Options.ValueOptions.Parent = &fieldtype.String{}
						// *s[ioi].Options.ValueOptions.Parent = *s[ioi].Options.Parent
					}
					ymlValue.Field(i).Set(reflect.ValueOf(s))
				}
			}
		}
	}

	*y = *yml

	return nil
}
