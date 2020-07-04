terraform {
    required_version = "~> 0.12"

    backend "remote" {
        organization = "viscat"
        workspaces {
            name = "url-shortener"
        }

    }
}