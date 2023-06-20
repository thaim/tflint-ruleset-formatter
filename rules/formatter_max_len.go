package rules

import (
	"fmt"

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

type formatterMaxLenConfig struct {
	Length int `hclext:"length,optional"`
}

func newFormatterMaxLenConfig() formatterMaxLenConfig {
	return formatterMaxLenConfig{
		Length: 80,
	}
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
	config := newFormatterMaxLenConfig()
	if err := runner.DecodeRuleConfig(r.Name(), &config); err != nil {
		return err
	}

	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	for name, file := range files {
		if err := r.checkLineLength(runner, name, file, config); err != nil {
			return err
		}

	}
	return nil
}

func (r *FormatterMaxLenRule) checkLineLength(runner tflint.Runner, filename string, file *hcl.File, config formatterMaxLenConfig) error {
	tokens, diags := hclsyntax.LexConfig(file.Bytes, filename, hcl.InitialPos)
	if diags.HasErrors() {
		return diags
	}

	for _, token := range tokens {
		if token.Range.End.Column >= config.Length {
			err := runner.EmitIssue(
				r,
				fmt.Sprintf("Line length is too long (current: %d, max: %d)", token.Range.End.Column, config.Length),
				token.Range,
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
