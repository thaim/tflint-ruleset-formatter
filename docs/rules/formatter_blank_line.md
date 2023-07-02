# formatter_blank_line

This rule ensures that there are no extra blank lines.

## Example

```hcl

// file start with blank lines.
resource "null_resource" "foo" {
}
```

```hcl
resource "null_resource" "foo" {
}
// file start with more than one blank lines.


```


```hcl
resource "null_resource" "foo" {
}
// more than two blank lines between blocks.



resource "null_resource" "bar" {
}
```
