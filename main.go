package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-prettier/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "prettier",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewPrettierTrailingCommaRule(),
				rules.NewPrettierMaxLenRule(),
			},
		},
	})
}
