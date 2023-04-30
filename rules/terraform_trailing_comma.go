package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
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
	logger.Debug("start TerraformTrailingCommaRule")
	diags := runner.WalkExpressions(tflint.ExprWalkFunc(func(expr hcl.Expression) hcl.Diagnostics {
		// Check if the expression is a literal
		if lit, ok := expr.(*hclsyntax.LiteralValueExpr); ok {
			// Check if the literal is a list or tuple
			logger.Debug("literal value expression: %s", lit.Val.GoString())
			if lit.Val.Type().IsListType() || lit.Val.Type().IsTupleType() {
				// Check if the last element is a comma
				// treat all as error for debug
				if true {
					if err := runner.EmitIssue(r, "List value should end with a comma.", lit.Range()); err != nil {
						return hcl.Diagnostics{
							{
								Severity: hcl.DiagError,
								Summary:  "Failed to emit issue",
								Detail:   err.Error(),
							},
						}
					}
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
