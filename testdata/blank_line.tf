
// blank line at the beginning of the file
resource "aws_instance" "blank_line_at_the_beginning_of_the_block" {

  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"
}
// blank line between the blocks



resource "aws_instance" "blank_line_at_the_end_of_the_block" {
  ami           = "ami-0c55b159cbfafe1f0"
  instance_type = "t2.micro"

}

resource "aws_instance" "blank_at_block_middle" {
  ami = "ami-0c55b159cbfafe1f0"


  instance_type = "t2.micro"
}
// blank line at the end of the file

