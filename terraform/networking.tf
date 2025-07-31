
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "koronet-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["us-east-1a", "us-east-1b", "us-east-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  # Since this should be as closed as prod environments, creating one NAT per AZ
  enable_nat_gateway = true
  single_nat_gateway = false
  one_nat_gateway_per_az = true
}

resource "aws_security_group" "koronet_lb_sg" {
  vpc_id = module.vpc.vpc_arn
  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_security_group" "service_security_group" {
  vpc_id = module.vpc.vpc_arn
  ingress {
    from_port = 0
    to_port   = 0
    protocol  = "-1"
    security_groups = [aws_security_group.koronet_lb_sg.id]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_lb" "koronet_alb" {
  name               = "koronet"
  internal = false
  load_balancer_type = "application"
  subnets  = module.vpc.public_subnets
  
  security_groups = [aws_security_group.koronet_lb_sg.id]
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.koronet_alb.arn
  port              = "80"
  protocol          = "HTTP"

  default_action {
    type = "forward"
    target_group_arn = aws_lb_target_group.webapp.arn
  }
}

resource "aws_lb_target_group" "webapp" {
  name        = "koronet-webapp"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = module.vpc.vpc_arn
  health_check {
    matcher = "200"
    path    = "/"
    interval            = 10
    timeout             = 2
    healthy_threshold   = 3
    unhealthy_threshold = 3
  }
  depends_on = [aws_lb.koronet_alb]
}
