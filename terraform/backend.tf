terraform {
  backend "s3" {
    bucket = var.state_bucket
    region = var.region
    use_lockfile = true
  }
}

resource "aws_s3_bucket" "tf-state" {
    bucket = var.state_bucket
}
