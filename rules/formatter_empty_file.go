package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/project"
)

// FormatterEmptyFileRule checks if a Terraform file actually has any content
type FormatterEmptyFileRule struct {
	tflint.DefaultRule
}

func NewFormatterEmptyFileRule() *FormatterEmptyFileRule {
	return &FormatterEmptyFileRule{}
}

func (r *FormatterEmptyFileRule) Name() string {
	return "formatter_empty_file"
}

func (r *FormatterEmptyFileRule) Enabled() bool {
	return true
}

func (r *FormatterEmptyFileRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FormatterEmptyFileRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FormatterEmptyFileRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	for name, file := range files {
		if err := r.checkEmpty(runner, name, file); err != nil {
			return err
		}

	}
	return nil
}

func (r *FormatterEmptyFileRule) checkEmpty(runner tflint.Runner, filename string, file *hcl.File) error {
	if len(file.Bytes) == 0 {
		fileRange := hcl.Range{
			Filename: filename,
			Start:    hcl.Pos{Line: 0, Column: 0},
			End:      hcl.Pos{Line: 0, Column: 0},
		}

		return runner.EmitIssue(
			r,
			"file has no content",
			fileRange,
		)
	}

	return nil
}
