resource "aws_security_group" "main" {
  description = "sg for ssh and application ports"
  name        = "${local.service_name}-sg"
}

resource "aws_security_group_rule" "ssh" {
  type              = "ingress"
  description       = "ssh"
  from_port         = 22
  to_port           = 22
  protocol          = "tcp"
  cidr_blocks       = ["${var.my_ip}/32"]
  security_group_id = aws_security_group.main.id
}

resource "aws_security_group_rule" "tcp_access" {
  type              = "ingress"
  description       = "for tcp server"
  from_port         = 10000
  to_port           = 10020
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.main.id
}

resource "aws_security_group_rule" "ssl_outbound" {
  type              = "egress"
  description       = "ssl outbound"
  from_port         = 443
  to_port           = 443
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.main.id
}

resource "aws_security_group_rule" "http_outbound" {
  type              = "egress"
  description       = "http outbound"
  from_port         = 80
  to_port           = 80
  protocol          = "tcp"
  cidr_blocks       = ["0.0.0.0/0"]
  security_group_id = aws_security_group.main.id
}

resource "aws_instance" "main" {
  instance_type = "t2.micro"
  ami           = "ami-0cd7ad8676931d727" # Ubuntu Server 22.04 LTS (64 bit, x86)
  vpc_security_group_ids = [
    aws_security_group.main.id
  ]
  key_name = aws_key_pair.main.key_name
  user_data_base64 = base64encode(join("\n", [
    "#cloud-config",
    yamlencode({
      packages : [
        "git",
        "golang"
      ]
      package_update : true
      package_updated : true
      # なぜか clone できない
      runcmd : [
        ["git clone https://github.com/shimoch-214/protohackers_problems.git $HOME/protohackers_problems"]
      ]
    })
  ]))
}

resource "aws_key_pair" "main" {
  key_name   = var.key_name
  public_key = file(var.public_key_path)
}
