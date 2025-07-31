
module "vpc" {
  source = "terraform-aws-modules/vpc/aws"

  name = "koronet-vpc"
  cidr = "10.0.0.0/16"

  azs             = ["eu-west-1a", "eu-west-1b", "eu-west-1c"]
  private_subnets = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets  = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]

  # This is only for cost reasons, if this would be applied to prod, I would enable one NAT Gateway per AZ for High Availability
  enable_nat_gateway = true
  single_nat_gateway = true
  one_nat_gateway_per_az = false
  
  # enable_vpn_gateway = true

  tags = {
    Terraform = "true"
    Environment = "dev"
  }
}
