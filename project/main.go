package project

import (
	"fmt"
	"runtime/debug"
)

var version string

// ReferenceLink returns the rule reference link
func ReferenceLink(name string) string {
	return fmt.Sprintf("https://github.com/thaim/tflint-ruleset-formatter/blob/%s/docs/rules/%s.md", GetVersion(), name)
}

func SetVersion(v string) {
	version = v
}

func GetVersion() string {
	if version != "" {
		return version
	}

	i, ok := debug.ReadBuildInfo()
	if !ok || i.Main.Version == "(devel)" {
		return "main"
	}

	return i.Main.Version
}
