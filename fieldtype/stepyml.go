package fieldtype

import (
	"fmt"

	envmanModels "github.com/bitrise-io/envman/models"
	"github.com/bitrise-io/step-yml-linter/lint"
	"gopkg.in/yaml.v3"
)

// StepYML ...
type StepYML struct {
	Title         Title   `yaml:"title,omitempty"`
	Summary       *string `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description   *string `json:"description,omitempty" yaml:"description,omitempty"`
	Website       *string `json:"website,omitempty" yaml:"website,omitempty"`
	SourceCodeURL *string `json:"source_code_url,omitempty" yaml:"source_code_url,omitempty"`
	SupportURL    *string `json:"support_url,omitempty" yaml:"support_url,omitempty"`
	// PublishedAt         *time.Time                          `json:"published_at,omitempty" yaml:"published_at,omitempty"`
	// Source              *StepSourceModel                    `json:"source,omitempty" yaml:"source,omitempty"`
	AssetURLs       map[string]string `json:"asset_urls,omitempty" yaml:"asset_urls,omitempty"`
	HostOsTags      []string          `json:"host_os_tags,omitempty" yaml:"host_os_tags,omitempty"`
	ProjectTypeTags []string          `json:"project_type_tags,omitempty" yaml:"project_type_tags,omitempty"`
	TypeTags        []string          `json:"type_tags,omitempty" yaml:"type_tags,omitempty"`
	// Dependencies        []DependencyModel                   `json:"dependencies,omitempty" yaml:"dependencies,omitempty"`
	// Toolkit             *StepToolkitModel                   `json:"toolkit,omitempty" yaml:"toolkit,omitempty"`
	// Deps                *DepsModel                          `json:"deps,omitempty" yaml:"deps,omitempty"`
	IsRequiresAdminUser *bool                               `json:"is_requires_admin_user,omitempty" yaml:"is_requires_admin_user,omitempty"`
	IsAlwaysRun         *bool                               `json:"is_always_run,omitempty" yaml:"is_always_run,omitempty"`
	IsSkippable         *bool                               `json:"is_skippable,omitempty" yaml:"is_skippable,omitempty"`
	RunIf               *string                             `json:"run_if,omitempty" yaml:"run_if,omitempty"`
	Timeout             *int                                `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Meta                map[string]interface{}              `json:"meta,omitempty" yaml:"meta,omitempty"`
	Inputs              []envmanModels.EnvironmentItemModel `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs             []envmanModels.EnvironmentItemModel `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	parsed              bool
}

// UnmarshalYAML ...
func (*StepYML) UnmarshalYAML(node *yaml.Node) error {
	var val StepYML

	if err := node.Decode(&val); err != nil {
		return err
	}

	var msgs []string
	for _, check := range stepYMLChecks {
		if err := check(val); err != nil {
			msgs = append(msgs, err.Error())
		}
	}

	if len(msgs) > 0 {
		lint.Warnings = append(lint.Warnings, lint.Warning{Messages: msgs, Line: node.Line, Column: node.Column})
	}

	return nil
}

var stepYMLChecks = []func(StepYML) error{
	DescriptionIsLongerThanSummary,
}

// DescriptionIsLongerThanSummary ...
func DescriptionIsLongerThanSummary(stepYML StepYML) error {
	if len(*stepYML.Summary) > len(*stepYML.Description) {
		return fmt.Errorf("summary should be shorter than description")
	}
	return nil
}
