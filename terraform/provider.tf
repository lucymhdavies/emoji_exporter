provider "aws" {
  region  = "eu-west-1"
  profile = "lmhd_root"
}

terraform {
  backend "s3" {
    bucket  = "lmhd-root-terraform"
    key     = "emoji_exporter.tfstate"
    region  = "eu-west-1"
    profile = "lmhd_root"
  }
}
