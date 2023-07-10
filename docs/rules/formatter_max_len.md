# formatter_max_len

This rule ensures the limitation of code line length.

## Configuration

* `length`: Maximum line length allowed (integer, default 80).


## Example

```hcl
locals {
  application_name = var.environment == "production" ? "production-service" : "development-service"
}
```


## Why
Shorter line length is easier to read.


## How To Fix
There are several ways to fix:

1. Use parentheses to split into multiple lines

```hcl
locals {
  applicatioin_name = (
    var.environment == "production" ? "production-service" : "development-service"
  )
}
```

1. Use local values to store temporary local variables

```
locals {
  application_name = "${var.environment}-service"
}

```
