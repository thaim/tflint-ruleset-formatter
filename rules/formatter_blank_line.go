package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/logger"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/thaim/tflint-ruleset-formatter/project"
)

// TODO: Write the rule's description here
// FormatterBlankLineRule checks ...
type FormatterBlankLineRule struct {
	tflint.DefaultRule
}

// NewFormatterBlankLineRule returns new rule with default attributes
func NewFormatterBlankLineRule() *FormatterBlankLineRule {
	return &FormatterBlankLineRule{}
}

// Name returns the rule name
func (r *FormatterBlankLineRule) Name() string {
	return "formatter_blank_line"
}

// Enabled returns whether the rule is enabled by default
func (r *FormatterBlankLineRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *FormatterBlankLineRule) Severity() tflint.Severity {
	return tflint.WARNING
}

// Link returns the rule reference link
func (r *FormatterBlankLineRule) Link() string {
	return project.ReferenceLink(r.Name())
}

// TODO: Write the details of the inspection
// Check checks ...
func (r *FormatterBlankLineRule) Check(runner tflint.Runner) error {
	files, err := runner.GetFiles()
	if err != nil {
		return err
	}

	for name, file := range files {
		if err := r.checkTooManyBlankLines(runner, name, file); err != nil {
			return err
		}
	}

	return nil
}

func (r *FormatterBlankLineRule) checkTooManyBlankLines(runner tflint.Runner, name string, file *hcl.File) error {
	logger.Debug("start FormtterBlankLineRule")

	if len(file.Bytes) == 0 {
		return nil
	}

	runes := []rune(string(file.Bytes))
	line := 0
	for ; line < len(runes); line++ {
		if runes[line] != '\n' {
			break
		}
	}
	logger.Debug(fmt.Sprintf("found %d newline at start of %s", line, name))
	if line != 0 {
		issueRange := hcl.Range{
			Filename: name,
			Start: hcl.Pos{Line: 1, Column: 1},
			End: hcl.Pos{Line: 1+line, Column: 1},
		}
		runner.EmitIssue(
			r,
			"too many blank lines",
			issueRange,
		)
		logger.Debug(fmt.Sprintf("found too many blank line at start of %s", name))
	}

	totalLine := countLines(runes)
	logger.Debug(fmt.Sprintf("totalLine: %d in %s", totalLine, name))
	for line = 0; line < totalLine; line++ {
		if runes[len(runes)-line-1] != '\n' {
			break
		}
	}
	logger.Debug(fmt.Sprintf("found %d newline at end of %s", line, name))
	if line > 1 {
		issueRange := hcl.Range{
			Filename: name,
			Start: hcl.Pos{Line: totalLine-(line-1), Column: 1},
			End: hcl.Pos{Line: totalLine, Column: 1},
		}
		runner.EmitIssue(
			r,
			"too many blank lines",
			issueRange,
		)
		logger.Debug(fmt.Sprintf("found too many blank line at end of %s", name))
	}

	return nil
}

func countLines(runes []rune) int {
	line := 1
	for i := 0; i < len(runes); i++ {
		if runes[i] == '\n' {
			line++
		}
	}
	return line
}
