package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_FormatterTrailingComma(t *testing.T) {
	tests := []struct {
		Name     string
		Content    string
		Expected helper.Issues
	}{
		{
			Name: "require trailing comma",
			Content: `
data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"
    actions = [
      "s3:ListBucket",
      "s3:GetObject"
    ]
    resources = ["arn:aws:s3:::*"]
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterTrailingCommaRule(),
					Message: "no new line at end of file",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 6, Column: 6},
						End:      hcl.Pos{Line: 6, Column: 20},
					},
				},
			},
		},
	}

	rule := NewFormatterTrailingCommaRule()

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			runner := helper.TestRunner(t, map[string]string{"resource.tf": test.Content})

			if err := rule.Check(runner); err != nil {
				t.Fatalf("unexpected error occurred: %s", err)
			}

			helper.AssertIssues(t, test.Expected, runner.Issues)
		})
	}
}
