package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

type PrettierMaxLenRule struct {
	tflint.DefaultRule
}

func NewPrettierMaxLenRule() *PrettierMaxLenRule {
	return &PrettierMaxLenRule{}
}

func (r *PrettierMaxLenRule) Name() string {
	return "prettier_max_len"
}

func (r *PrettierMaxLenRule) Enabled() bool {
	return true
}

func (r *PrettierMaxLenRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *PrettierMaxLenRule) Link() string {
	return ""
}

func (r *PrettierMaxLenRule) Check(runner tflint.Runner) error {
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

func (r *PrettierMaxLenRule) checkLineLength(runner tflint.Runner, filename string, file *hcl.File) error {
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
