package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_FormatterEmptyFile(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name:    "file has no content",
			Content: ``,
			Expected: helper.Issues{
				{
					Rule:    NewFormatterEmptyFileRule(),
					Message: "file has no content",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 0, Column: 0},
						End:      hcl.Pos{Line: 0, Column: 0},
					},
				},
			},
		},
	}

	rule := NewFormatterEmptyFileRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
