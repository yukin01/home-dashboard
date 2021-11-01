terraform {
  required_version = "1.0.10"
  backend "remote" {
    hostname     = "app.terraform.io"
    organization = "yukin01"

    workspaces {
      name = "home-dashboard"
    }
  }
  required_providers {
    google = {
      version = "3.90.0"
      source  = "hashicorp/google"
    }
  }
}

provider "google" {
  project = var.gcp_project
  region  = var.gcp_region
}

data "google_project" "this" {}
