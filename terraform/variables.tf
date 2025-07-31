variable "aws_access_key" {
  sensitive = true
}

variable "aws_secret_key" {
  sensitive = true
}

variable "aws_region" {
  default = "us-east-1"
}

variable "state_bucket" {
  default = "koronet-tf-state"
}