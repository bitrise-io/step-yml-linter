package step

import (
	"time"

	"github.com/bitrise-io/step-yml-linter/fieldtype"
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
