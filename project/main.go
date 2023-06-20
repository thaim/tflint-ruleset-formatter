package project

import (
	"fmt"
	"runtime/debug"
)

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/thaim/tflint-ruleset-formatter/blob/%s/docs/rules/%s.md", getVersion(), name)
}

func getVersion() string {
	i, ok := debug.ReadBuildInfo()
	if !ok || i.Main.Version == "(devel)" {
		return "main"
	}

	return i.Main.Version
}
