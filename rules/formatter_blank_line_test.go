package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_FormatterBlankLine(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "blank at start of file",
			Content: `
resource "null_resource" "null" {
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterBlankLineRule(),
					Message: "too many blank lines at start of file",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 1},
					},
				},
			},
		},
		{
			Name: "blank at end of file",
			Content: `resource "null_resource" "null" {
}

`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterBlankLineRule(),
					Message: "too many blank lines at end of file",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 1},
						End:      hcl.Pos{Line: 4, Column: 1},
					},
				},
			},
		},
		{
			Name: "blank at middle of file",
			Content: `resource "null_resource" "first" {
}



resource "null_resource" "second" {
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterBlankLineRule(),
					Message: "too many blank lines at middle of file",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 1},
						End:      hcl.Pos{Line: 5, Column: 1},
					},
				},
			},
		},
		{
			Name: "blank at start of block",
			Content: `resource "null_resource" "null" {

  triggers = {
	foo = "bar"
  }
}
`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterBlankLineRule(),
					Message: "too many blank lines at start of block",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 2, Column: 1},
						End:      hcl.Pos{Line: 3, Column: 1},
					},
				},
			},
		},
		{
			Name: "blank at end of block",
			Content: `resource "null_resource" "null" {
  triggers = {
	foo = "bar"
  }

}
`,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterBlankLineRule(),
					Message: "too many blank lines at end of block",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 5, Column: 1},
						End:      hcl.Pos{Line: 6, Column: 1},
					},
				},
			},
		},
		{
			Name: "only blank line",
			Content: `
`,
			Expected: helper.Issues{
				{
					Rule: NewFormatterBlankLineRule(),
					Message: "too many blank lines at start of file",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 1, Column: 1},
						End:      hcl.Pos{Line: 2, Column: 1},
					},
				},
			},
		},
	}

	rule := NewFormatterBlankLineRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
