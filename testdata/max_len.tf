resource "aws_instance" "foo" {
  ami           = "ami-0ff8a91507f77f867"
  instance_type = "t3.micro"

  tags = {
    Name = "too-long-name-for-aws-instance-since-prettier-ruleset-requires-max-len-80"
  }
}
