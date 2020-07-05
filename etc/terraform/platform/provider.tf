provider "aws" {
  region  = var.region
  version = "~> 2"
}

provider "tfe" {
  token    = var.tfe_token
  version  = "~> 0.15.0"
}