terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

variable "aws_access_key" {
  type = string
}
variable "aws_secret_key" {
  type = string
}

provider "aws" {
  region     = "us-east-1"
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
}


resource "aws_vpc" "default-vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true
  tags                 = {
    env = "dev"
  }
}

resource "aws_subnet" "default-subnet" {
  cidr_block = "10.0.0.0/24"
  vpc_id     = aws_vpc.default-vpc.id
}

resource "aws_instance" "app_server" {
  ami             = "ami-0ff8a91507f77f867"
  instance_type   = "t2.nano"
  subnet_id       = aws_subnet.default-subnet.id

  tags = {
    Name = "Chapter14"
  }
}