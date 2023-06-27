# formatter_trailing_comma

This rule ensures that tuple element always end with comma.

## Example

```hcl
data "aws_iam_policy_document" "example" {
  statement {
    sid = "1"
    actions = [
      "ec2:Describe*",
      "ec2:Get*",
      "ec2:List*"
    ]
    resources = ["arn:aws:s3:::*"]
  }
}
```
