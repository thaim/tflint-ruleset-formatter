# TFLint Ruleset with formatter
[![Build Status](https://github.com/thaim/tflint-ruleset-formatter/workflows/build/badge.svg?branch=main)](https://github.com/thaim/tflint-ruleset-formatter/actions)

TFLint ruleset plugin for Terraform Language with more format check.
This ruleset achieves more syntax checking that cannot be detected in 'terraform fmt'.

## Requirements

- TFLint v0.40+
- Go v1.20

## Installation
Declare a config in `.tflint.hcl` as follows:

```hcl
plugin "formatter" {
  enabled = true
  version = "0.2.4"
  source  = "github.com/thaim/tflint-ruleset-formatter"

  signing_key = <<-KEY
-----BEGIN PGP PUBLIC KEY BLOCK-----

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
//0y4Ag=
=opj9
-----END PGP PUBLIC KEY BLOCK-----
KEY
}
```

You can customize rule specific configurations. See [rule docs](docs/rules/) for more details.

## Rules

|Rule|Description|Severity|Enabled by default|
| --- | --- | --- | --- |
|[formatter_blank_line](docs/rules/formatter_blank_line.md)|ensures that there are no extra blank lines|WARNING|✔
|[formatter_empty_file](docs/rules/formatter_empty_file.md)|ensures that files are not empty|WARNING|✔
|[formatter_eof](docs/rules/formatter_eof.md)|ensures that file end with new line|WARNING|✔
|[formatter_max_len](docs/rules/formatter_max_len.md)|ensures the limitation of code line length|WARNING|✔
|[formatter_trailing_comma](docs/rules/formatter_trailing_comma.md)|ensures that tuple element always end with comma|WARNING|✔

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
plugin "formatter" {
  enabled = true
}
EOS
$ tflint
```

## Writing a new rule
Run command below and answer the options.

```
$ go run rules/generator/main.go
```
