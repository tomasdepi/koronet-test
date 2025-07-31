provider "aws" {
  //access_key = var.aws_access_key
  //secret_key = var.aws_secret_key
  region = var.aws_region

  skip_credentials_validation = true
  skip_requesting_account_id  = true
  skip_metadata_api_check     = true
  access_key                  = "mock_access_key"
  secret_key                  = "mock_secret_key"

}
