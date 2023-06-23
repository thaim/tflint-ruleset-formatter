package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_FormatterEOF(t *testing.T) {
	tests := []struct {
		Name     string
		Content    string
		Expected helper.Issues
	}{
		{
			Name: "no newline at end of file",
			Content: `
resource "aws_instance" "test" {
  ami           = "ami-0ff8a91507f77f867"
  instance_type = "t3.micro"
}`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterEOFRule(),
					Message: "no new line at end of file",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 1, Column: 2},
					},
				},
			},
		},
	}

	rule := NewFormatterEOFRule()

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
