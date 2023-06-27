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
  statement {
    sid = "2"
    actions = [
      "ec2:GetObject",
      "ec2:ListBucket",
    ]
    resources = ["arn:aws:s3:::*"]
  }
}
