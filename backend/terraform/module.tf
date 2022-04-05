provider "google" {
  credentials = "${file("../credentials/puzzles-345814-24d9b5fcb5ca.json")}"
  project = "puzzles-345814"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}
terraform {
  # version指定
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 4.1.0"
    }
  }
}

module "cloud_run" {
  source    = "./modules/cloud_run"
  name      = "puzzles"
  image     = "gcr.io/puzzles-345814/puzzles"
  location  = "asia-northeast1"
  project   = "puzzles-345814"
}
