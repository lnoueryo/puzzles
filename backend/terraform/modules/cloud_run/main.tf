provider "google" {
  project = "puzzles-345814"
  region  = "asia-northeast1"
  zone    = "asia-northeast1-a"
}
resource "google_cloud_run_service" "default" {
  name     = var.name
  location = var.location

  # metadata {
  #   annotations = {
  #     "run.googleapis.com/cloudsql-instances" = "trim-tide-313616:asia-northeast1:gmap"
  #   }
  # }

  template {
    spec {
      containers {
        image = "gcr.io/puzzles-345814/puzzles"
        resources {
          limits = { "memory" : "1Gi" }
        }
      }
    }
  }
}
data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.default.location
  project     = google_cloud_run_service.default.project
  service     = google_cloud_run_service.default.name

  policy_data = data.google_iam_policy.noauth.policy_data
}

resource "google_cloud_run_domain_mapping" "default" {
  location = google_cloud_run_service.default.location
  name     = "puzzles-api.jounetsism.biz"

  metadata {
    namespace = var.project
  }

  spec {
    route_name = google_cloud_run_service.default.name
  }
}

