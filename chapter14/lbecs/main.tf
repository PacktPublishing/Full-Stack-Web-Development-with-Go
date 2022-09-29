variable "aws_region" {
  default = "us-east-1"
}

variable "aws_access_key" {
  type = string
}
variable "aws_secret_key" {
  type = string
}

terraform {
  required_version = ">= 1.2.0"
}

provider "aws" {
  region     = var.aws_region
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}


resource "aws_vpc" "lbecs-vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  tags                 = {
    env = "dev"
  }
}

resource "aws_subnet" "lbecs-subnet" {
  availability_zone = "us-east-1a"
  cidr_block        = "10.0.0.0/24"
  vpc_id            = aws_vpc.lbecs-vpc.id
}

resource "aws_subnet" "lbecs-subnet-1" {
  availability_zone = "us-east-1b"
  cidr_block        = "10.0.1.0/24"
  vpc_id            = aws_vpc.lbecs-vpc.id
}

resource "aws_internet_gateway" "lbecs-igw" {
  vpc_id = aws_vpc.lbecs-vpc.id

  tags = {
    Name = "Internet Gateway"
  }
}

resource "aws_default_route_table" "lbecs-subnet-default-route-table" {
  default_route_table_id = aws_vpc.lbecs-vpc.default_route_table_id
  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = "${aws_internet_gateway.lbecs-igw.id}"
  }
}

resource "aws_security_group" "lbecs-security-group" {
  name        = "allow_http"
  description = "Allow HTTP inbound traffic"
  vpc_id      = aws_vpc.lbecs-vpc.id

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "Allow HTTP for all"
    from_port   = 80
    to_port     = 3333
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_lb" "lbecs-load-balancer" {
  name               = "load-balancer"
  internal           = false
  load_balancer_type = "application"
  security_groups    = [aws_security_group.lbecs-security-group.id]
  subnets            = [aws_subnet.lbecs-subnet.id, aws_subnet.lbecs-subnet-1.id]
  tags               = {
    env = "dev"
  }
}


resource "aws_lb_target_group" "lbecs-load-balancer-target-group" {
  name        = "load-balancer-tg"
  port        = 80
  protocol    = "HTTP"
  target_type = "ip"
  vpc_id      = aws_vpc.lbecs-vpc.id
}

resource "aws_lb_listener" "lbecs-load-balancer-listener" {
  load_balancer_arn = aws_lb.lbecs-load-balancer.arn
  port              = "80"
  protocol          = "HTTP"
  default_action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.lbecs-load-balancer-target-group.arn
  }
}


resource "aws_ecs_cluster" "lbecs-ecs-cluster" {
  name = "lbecs-ecs-cluster"
}

resource "aws_ecs_task_definition" "lbecs-ecs-task-definition" {
  family                   = "service"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  container_definitions    = jsonencode([
    {
      name         = "lbecs-ecs-cluster-chapter14"
      image        = "ghcr.io/nanikjava/golangci/chapter12:latest"
      cpu          = 512
      memory       = 2048
      essential    = true
      portMappings = [
        {
          containerPort = 3333
        }
      ]
    }
  ])
}

resource "aws_ecs_service" "lbecs-ecs-service" {
  name            = "lbecs-ecs-service"
  cluster         = aws_ecs_cluster.lbecs-ecs-cluster.id
  task_definition = aws_ecs_task_definition.lbecs-ecs-task-definition.arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = [aws_subnet.lbecs-subnet.id, aws_subnet.lbecs-subnet-1.id]
    security_groups  = [aws_security_group.lbecs-security-group.id]
    assign_public_ip = true
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.lbecs-load-balancer-target-group.arn
    container_name   = "lbecs-ecs-cluster-chapter14"
    container_port   = 3333
  }

  tags = {
    env = "dev"
  }
}

