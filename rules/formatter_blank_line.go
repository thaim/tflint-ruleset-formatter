package rules

import (
	"fmt"
	"strings"

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
	err := r.checkFileStart(runner, name, runes)
	if err != nil {
		return err
	}

	err = r.checkFileEnd(runner, name, runes)
	if err != nil {
		return err
	}

	err = r.checkFileMiddle(runner, name, file)
	if err != nil {
		return err
	}

	return nil
}

func (r *FormatterBlankLineRule) checkFileStart(runner tflint.Runner, name string, runes []rune) error {
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
		err := runner.EmitIssue(
			r,
			"too many blank lines at start of file",
			issueRange,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *FormatterBlankLineRule) checkFileEnd(runner tflint.Runner, name string, runes []rune) error {
	totalLine := countLines(runes)
	line := 0

	for ; line < totalLine; line++ {
		if runes[len(runes)-line-1] != '\n' {
			break
		}
	}

	if line > 1 {
		issueRange := hcl.Range{
			Filename: name,
			Start: hcl.Pos{Line: totalLine-(line-1), Column: 1},
			End: hcl.Pos{Line: totalLine, Column: 1},
		}
		err := runner.EmitIssue(
			r,
			"too many blank lines at end of file",
			issueRange,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *FormatterBlankLineRule) checkFileMiddle(runner tflint.Runner, name string, file *hcl.File) error {
	bodyLines := splitNewline(string(file.Bytes))
	logger.Debug(fmt.Sprintf("start check File Middle for %s with total lines = %d", name, len(bodyLines)))

	firstNonBlankLine := 1
	for ; firstNonBlankLine <= len(bodyLines); firstNonBlankLine++ {
		if bodyLines[firstNonBlankLine-1] != "" {
			break
		}
	}

	line := 0
	for ; line < len(bodyLines); line++ {
		if bodyLines[len(bodyLines) - line - 1] != "" {
			break
		}
	}
	lastNonBlankLine := len(bodyLines) - line

	if (lastNonBlankLine <= firstNonBlankLine) {
		// empty file
		logger.Debug(fmt.Sprintf("empty file %s", name))
		return nil
	}
	logger.Debug(fmt.Sprintf("firstNonBlankLine = %d, lastNonBlankLine = %d", firstNonBlankLine, lastNonBlankLine))

	for line = firstNonBlankLine+1; line < lastNonBlankLine - 1; line++ {
		if bodyLines[line] == "" && bodyLines[line+1] == "" && bodyLines[line+2] == "" {
			block := file.OutermostBlockAtPos(hcl.Pos{Line: line+1, Column: 1})
			if block != nil && (block.DefRange.Start.Line < line+1 || block.DefRange.End.Line >= line+1) {
				logger.Debug(fmt.Sprintf("ignore block at %d in %s (expect line: %d)", block.DefRange.Start.Line, name, line+1))
				block = nil
			}

			if block != nil {
				logger.Debug(fmt.Sprintf("found too many blank lines in block at %s in line %d (block %s %s)", name, line, block.Type, block.Labels[1]))

				err := runner.EmitIssue(
					r,
					"too many blank lines in the middle of file",
					block.DefRange,
				)
				if err != nil {
					return err
				}

				line = block.DefRange.End.Line - 1
			} else {
				logger.Debug(fmt.Sprintf("found too many blank lines in non block at %s in line %d", name, line))
				lineEnd := line+2
				for ; lineEnd < lastNonBlankLine; lineEnd++ {
					if bodyLines[lineEnd] != "" {
						break
					}
				}

				err := runner.EmitIssue(
					r,
					"too many blank lines at middle of file",
					hcl.Range{
						Filename: name,
						Start: hcl.Pos{Line: line+1, Column: 1},
						End: hcl.Pos{Line: lineEnd, Column: 1},
					},
				)
				if err != nil {
					return err
				}

				line = lineEnd
			}
		}
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

func splitNewline(body string) []string {
	return strings.Split(body, "\n")
}
