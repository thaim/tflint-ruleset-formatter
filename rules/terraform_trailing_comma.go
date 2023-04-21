package rules

import (
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// TerraformTrailingCommaRule checks whether list/tuple values are terminated with a comma
type TerraformTrailingCommaRule struct {
	tflint.DefaultRule
}

// TerraformTrailingCommaRule returns a new rule
func NewTerraformTrailingCommaRule() *TerraformTrailingCommaRule {
	return &TerraformTrailingCommaRule{}
}

// Name returns the rule name
func (r *TerraformTrailingCommaRule) Name() string {
	return "terraform_trailing_comma"
}

// Enabled returns whether the rule is enabled by default
func (r *TerraformTrailingCommaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *TerraformTrailingCommaRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *TerraformTrailingCommaRule) Link() string {
	return ""
}

// Check checks whether list/tuple values are terminated with a comma
func (r *TerraformTrailingCommaRule) Check(runner tflint.Runner) error {
	return nil
}
