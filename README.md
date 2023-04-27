# TFLint Ruleset with prettier
[![Build Status](https://github.com/thaim/tflint-ruleset-prettier/workflows/build/badge.svg?branch=main)](https://github.com/thaim/tflint-ruleset-prettier/actions)

TFLint ruleset plugin for Terraform Language with prettier syntax check.
This ruleset achieves prettier syntax checking that cannot be detected in 'terraform fmt'.

## Requirements

- TFLint v0.40+
- Go v1.20

## Installation

TODO: This template repository does not contain release binaries, so this installation will not work. Please rewrite for your repository. See the "Building the plugin" section to get this template ruleset working.

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "prettier" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/thaim/tflint-ruleset-prettier"
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```

You can run the built plugin like the following:

```
$ cat << EOS > .tflint.hcl
plugin "prettier" {
  enabled = true
}
EOS
$ tflint
```
