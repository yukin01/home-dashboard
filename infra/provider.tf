terraform {
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "yukin01"

    workspaces {
      name = "home-dashboard"
    }
  }
}

provider "google" {
  credentials = file("service-account.json")
  project     = var.gcp_project
  region      = var.gcp_region
  version     = "~> 3.19"
}

data "google_project" "this" {}
