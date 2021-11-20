variable "instance_type" {}

resource "aws_instance" "default" {
  ami = "ami-0c3fd0f5d33134a76"

  instance_type = var.instance_type

  vpc_security_group_ids = [aws_security_group.default.id]

  user_data = file("./user_data.sh")
}

resource "aws_security_group" "default" {
  name = "ec2"

  ingress {
    from_port = 80
    to_port = 80
    protocol = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

output "public_dns" {
  value = aws_instance.default.public_dns
}
