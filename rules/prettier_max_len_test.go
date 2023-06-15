package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_PrettierMaxLenRune(t *testing.T) {
	tests := []struct {
		Name     string
		Content    string
		Expected helper.Issues
	}{
		{
			Name: "line length is too long",
			Content: `
resource "aws_instance" "too_long_resource_name_for_aws_instance_since_prettier_ruleset_requires_max_len_80" {
  ami           = "ami-0ff8a91507f77f867"
  instance_type = "t3.micro"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewPrettierMaxLenRule(),
					Message: "Line length is too long",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 26},
						End:      hcl.Pos{Line: 2, Column: 108},
					},
				},
				{
					Rule:    NewPrettierMaxLenRule(),
					Message: "Line length is too long",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 108},
						End:      hcl.Pos{Line: 2, Column: 109},
					},
				},
				{
					Rule:    NewPrettierMaxLenRule(),
					Message: "Line length is too long",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 110},
						End:      hcl.Pos{Line: 2, Column: 111},
					},
				},
			},
		},
	}

	rule := NewPrettierMaxLenRule()

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
