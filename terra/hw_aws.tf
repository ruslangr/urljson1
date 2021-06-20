provider "aws" {
  profile	= "default"
  region	= var.region
}

data "aws_availability_zones" "available" {}

locals {
  cluster_name = "ruslangr-eks"
}

module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "2.66.0"
  name                 = "ruslangr-vpc"
  cidr                 = "10.0.0.0/16"
  azs                  = data.aws_availability_zones.available.names
  public_subnets     = var.pub_subnet

  enable_nat_gateway   = true
  enable_dns_hostnames = true

}

data "aws_subnet_ids" "ruslangr-net-ids" {
  depends_on = [
    module.vpc
  ]
  vpc_id =  module.vpc.vpc_id
}

resource "aws_db_subnet_group" "ruslangr-net-gr" {
  name       = "main"
  subnet_ids = data.aws_subnet_ids.ruslangr-net-ids.ids
}