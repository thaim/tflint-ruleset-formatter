# formatter_empty_file

This rule ensures there are no empty Terraform files.

## Example

```
$ tflint

1 issue(s) found:

empty_file.tf:0:0: Warning - file has no content (formatter_empty_file)
```

## Why

Having empty files in a project is confusing and gives unused maintenance and confusing, e.g. when browsing through folders.

## How To Fix

Remove the empty files from the project.
