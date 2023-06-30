package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "formatter",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewFormatterTrailingCommaRule(),
				rules.NewFormatterMaxLenRule(),
				rules.NewFormatterEOFRule(),
				rules.NewFormatterBlankLineRule(),
			},
		},
	})
}
