package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/project"
	"github.com/thaim/tflint-ruleset-formatter/rules"
)

var (
	version = "main"
	commit  = "none"
	date    = "unknown"
)

func init() {
	project.SetVersion(version)
}

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "formatter",
			Version: project.GetVersion(),
			Rules: []tflint.Rule{
				rules.NewFormatterTrailingCommaRule(),
				rules.NewFormatterMaxLenRule(),
				rules.NewFormatterEOFRule(),
				rules.NewFormatterBlankLineRule(),
				rules.NewFormatterEmptyFileRule(),
			},
		},
	})
}
