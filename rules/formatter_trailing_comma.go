package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/project"
)

// FormatterTrailingCommaRule checks whether list/tuple values are terminated with a comma
type FormatterTrailingCommaRule struct {
	tflint.DefaultRule
}

// FormatterTrailingCommaRule returns a new rule
func NewFormatterTrailingCommaRule() *FormatterTrailingCommaRule {
	return &FormatterTrailingCommaRule{}
}

// Name returns the rule name
func (r *FormatterTrailingCommaRule) Name() string {
	return "formatter_trailing_comma"
}

// Enabled returns whether the rule is enabled by default
func (r *FormatterTrailingCommaRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FormatterTrailingCommaRule) Severity() tflint.Severity {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *FormatterTrailingCommaRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// Check checks whether list/tuple values are terminated with a comma
func (r *FormatterTrailingCommaRule) Check(runner tflint.Runner) error {
	logger.Debug("start FormatterTrailingCommaRule")
	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
		tuple, ok := expr.(*hclsyntax.TupleConsExpr)
		if !ok {
			return nil
		}
		if len(tuple.Exprs) <= 1 {
			return nil
		}

		// convert TupleConsExpr to tokens and count number of commas
		file, _ := runner.GetFile(tuple.SrcRange.Filename)
		tokens, diags := hclsyntax.LexConfig(
			file.Bytes[tuple.SrcRange.Start.Byte:tuple.SrcRange.End.Byte],
			tuple.SrcRange.Filename,
			tuple.SrcRange.Start,
		)
		if diags.HasErrors() {
			return diags
		}
		count := 0
		for _, token := range tokens {
			if token.Type == hclsyntax.TokenComma {
				count++
			}
		}

		if count != len(tuple.Exprs) {
			msg := fmt.Sprintf("List value should end with a comma (actual: %d, expected: %d)", count, len(tuple.Exprs))
			if err := runner.EmitIssue(r, msg, expr.Range()); err != nil {
				return hcl.Diagnostics{
					{
						Severity: hcl.DiagError,
						Summary:  "Failed to emit issue",
						Detail:   err.Error(),
					},
				}
			}
		}
		return nil
	}))

	if diags.HasErrors() {
		return diags
	}
	return nil
}
