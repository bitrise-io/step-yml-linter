package main

import (
	"github.com/bitrise-io/step-yml-linter/lint"
	"github.com/bitrise-io/step-yml-linter/step"
)

var checks = []func(step.YML){
	// title
	func(s step.YML) {
		if s.Title.Value == "" {
			lint.Log(s.Title, "title is required")
		}
	},

	//summary
	func(s step.YML) {
		if s.Summary.Value == "" {
			lint.Log(s.Summary, "summary is required")
		}
		if len(s.Summary.Value) > len(s.Description.Value) {
			lint.Log(s.Summary, "summary must be shorter than description")
		}
	},

	// description
	func(s step.YML) {
		if s.Description.Value == "" {
			lint.Log(s.Description, "description is required")
		}
	},

	// inputs
	func(s step.YML) {
		for _, inp := range s.Inputs {
			if len(inp.Options.Description.Value) == 0 {
				lint.Log(inp.Options.Description, "description is required")
			}
		}
	},
}
