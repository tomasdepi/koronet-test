# terraform {
#   backend "s3" {
#     bucket = "koronet-tf-state"
#     region = "us-east-1"
#     use_lockfile = true
#   }
# }

resource "aws_s3_bucket" "tf-state" {
  bucket = var.state_bucket
}
