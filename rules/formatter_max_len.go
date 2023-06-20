package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/project"
)

type FormatterMaxLenRule struct {
	tflint.DefaultRule
}

func NewFormatterMaxLenRule() *FormatterMaxLenRule {
	return &FormatterMaxLenRule{}
}

func (r *FormatterMaxLenRule) Name() string {
	return "formatter_max_len"
}

func (r *FormatterMaxLenRule) Enabled() bool {
	return true
}

func (r *FormatterMaxLenRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FormatterMaxLenRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FormatterMaxLenRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	for name, file := range files {
		if err := r.checkLineLength(runner, name, file); err != nil {
			return err
		}

	}
	return nil
}

func (r *FormatterMaxLenRule) checkLineLength(runner tflint.Runner, filename string, file *hcl.File) error {
	tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return diags
	}

	for _, token := range tokens {
		if token.Range.End.Column >= 80 {
			err := runner.EmitIssue(
				r,
				"Line length is too long",
				token.Range,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
