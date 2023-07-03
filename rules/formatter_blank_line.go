package rules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
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

	hclsyntaxBody, ok := file.Body.(*hclsyntax.Body)
	if !ok {
		logger.Debug(fmt.Sprintf("cannot cast file.Body to hclsyntax.Body in %s", name))
		return nil
	}
	bodyLines := splitNewline(string(file.Bytes))
	for _, block := range hclsyntaxBody.Blocks {
		r.checkBlockStart(runner, name, block, bodyLines)
		r.checkBlockEnd(runner, name, block, bodyLines)
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

	for line = firstNonBlankLine+1; line < lastNonBlankLine - 1; line++ {
		if bodyLines[line] == "" && bodyLines[line+1] == "" && bodyLines[line+2] == "" {
			block := file.OutermostBlockAtPos(hcl.Pos{Line: line+1, Column: 1})
			if block != nil && (block.DefRange.Start.Line < line+1 || block.DefRange.End.Line >= line+1) {
				logger.Debug(fmt.Sprintf("ignore block at %d in %s (expect line: %d)", block.DefRange.Start.Line, name, line+1))
				block = nil
			}

			if block != nil {
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

func (r *FormatterBlankLineRule) checkBlockStart(runner tflint.Runner, name string, block *hclsyntax.Block, bodyLines []string) error {
	logger.Debug(fmt.Sprintf("check block start at %d and end at %d in %s", block.Range().Start.Line, block.Range().End.Line, name))

	line := block.Range().Start.Line + 1
	for ; line < block.Range().End.Line; line++ {
		if bodyLines[line-1] != "" {
			break
		}
	}
	if line != block.Range().Start.Line + 1 {
		err := runner.EmitIssue(
			r,
			"too many blank lines at start of block",
			hcl.Range{
				Filename: name,
				Start: hcl.Pos{Line: block.Range().Start.Line+1, Column: 1},
				End: hcl.Pos{Line: line, Column: 1},
			},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *FormatterBlankLineRule) checkBlockEnd(runner tflint.Runner, name string, block *hclsyntax.Block, bodyLines []string) error {
	logger.Debug(fmt.Sprintf("check block start at %d and end at %d in %s", block.Range().Start.Line, block.Range().End.Line, name))

	line := 1
	for ; block.Range().End.Line - line > block.Range().Start.Line; line++ {
		if bodyLines[block.Range().End.Line - line - 1] != "" {
			break
		}
	}
	if line != 1 {
		err := runner.EmitIssue(
			r,
			"too many blank lines at end of block",
			hcl.Range{
				Filename: name,
				Start: hcl.Pos{Line: block.Range().End.Line - line + 1, Column: 1},
				End: hcl.Pos{Line: block.Range().End.Line, Column: 1},
			},
		)
		if err != nil {
			return err
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
