# TFLint Ruleset with stricter
[![Build Status](https://github.com/terraform-linters/tflint-ruleset-template/workflows/build/badge.svg?branch=main)](https://github.com/terraform-linters/tflint-ruleset-template/actions)

TFLint ruleset plugin for Terraform Language with stricter syntax check.
This ruleset achieves stricter syntax checking that cannot be detected in 'terraform fmt'.

## Requirements

- TFLint v0.40+
- Go v1.20

## Installation

TODO: This template repository does not contain release binaries, so this installation will not work. Please rewrite for your repository. See the "Building the plugin" section to get this template ruleset working.

You can install the plugin with `tflint --init`. Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "stricter" {
  enabled = true

  version = "0.1.0"
  source  = "github.com/thaim/tflint-ruleset-stricter"
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
plugin "stricter" {
  enabled = true
}
EOS
$ tflint
```
