package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/project"
)

type FormatterEOFRule struct {
	tflint.DefaultRule
}

func NewFormatterEOFRule() *FormatterEOFRule {
	return &FormatterEOFRule{}
}

func (r *FormatterEOFRule) Name() string {
	return "formatter_eof"
}

func (r *FormatterEOFRule) Enabled() bool {
	return true
}

func (r *FormatterEOFRule) Severity() tflint.Severity {
	return tflint.WARNING
}

func (r *FormatterEOFRule) Link() string {
	return project.ReferenceLink(r.Name())
}

func (r *FormatterEOFRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}
	for name, file := range files {
		if err := r.checkEOF(runner, name, file); err != nil {
			return err
		}

	}
	return nil
}

func (r *FormatterEOFRule) checkEOF(runner tflint.Runner, filename string, file *hcl.File) error {
	if len(file.Bytes) == 0 {
		return nil
	}

	last := file.Bytes[len(file.Bytes)-1]
	if last != '\n' {
		eofRange := hcl.Range{
			Filename: filename,
			Start:    hcl.Pos{Line: 1, Column: 1},
			End:    hcl.Pos{Line: 1, Column: 2},
		}
		err := runner.EmitIssue(
			r,
			"no new line at end of file",
			eofRange,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
