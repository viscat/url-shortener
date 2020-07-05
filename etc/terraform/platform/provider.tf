provider "aws" {
  region  = var.region
  version = "~> 2"
  profile = "personal"
}

provider "tfe" {
  token    = var.tfe_token
  version  = "~> 0.15.0"
}