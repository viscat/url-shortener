variable "region" {
  default = "eu-west-2"
  type    = string
}

variable "tfe_token" {
  description = "Terraform Cloud token"
  type        = string
}
