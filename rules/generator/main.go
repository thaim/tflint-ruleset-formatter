package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

type metadata struct {
	RuleName   string
	RuleNameCC string
}

func main() {
	buf := bufio.NewReader(os.Stdin)
	fmt.Print("Rule name? (e.g. formatter_trailing_comma): ")
	ruleName, err := buf.ReadString('\n')
	if err != nil {
		panic(err)
	}
	ruleName = strings.Trim(ruleName, "\n")

	meta := &metadata{RuleNameCC: toCamel(ruleName), RuleName: ruleName}

	generateFile(fmt.Sprintf("rules/%s.go", ruleName), "rules/rule.go.tmpl", meta)
	generateFile(fmt.Sprintf("rules/%s_test.go", ruleName), "rules/rule_test.go.tmpl", meta)
	generateFile(fmt.Sprintf("docs/rules/%s.md", ruleName), "docs/rules/rule.md.tmpl", meta)

	fmt.Println(`
TODO:
1. Remove all "TODO" comments from generated files.
2. Write implementation of the rule.
3. Add a link to the generated documentation into docs/rules/README.md`)
}

func toCamel(ruleName string) string {
	var name string
	for _, word := range strings.Split(ruleName, "_") {
		name += strings.Title(word)
	}
	return name
}

func generateFile(fileName string, tmplName string, meta interface{}) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tmpl, err := template.ParseFiles(tmplName)
	if err != nil {
		panic(err)
	}

	if err := tmpl.Execute(file, meta); err != nil {
		panic(err)
	}

	fmt.Printf("Generated: %s\n", fileName)
}
