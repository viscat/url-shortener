terraform {
  required_version = "~> 0.12"
  required_providers {
    tfe = "~> 0.15.0"
  }

  backend "remote" {
    organization = "viscat"

    workspaces {

      name = "sandbox-platform"
    }

  }
}