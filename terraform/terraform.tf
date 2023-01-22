terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.51.0"
    }
  }

  backend "s3" {
    region  = "ap-northeast-1"
    encrypt = true
  }

  required_version = "1.3.7"
}

provider "aws" {
  region = "ap-northeast-1"
  default_tags {
    tags = {
      "service" = local.service_name
    }
  }
}
