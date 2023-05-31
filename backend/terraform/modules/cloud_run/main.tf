resource "google_cloud_run_service" "default" {
  name     = var.name
  location = var.location

  template {
    spec {
      service_account_name = var.service_account_name
      containers {
        image = var.image
        env {
          name = "DB_NAME"
          value = var.db_name
        }
        env {
          name = "DB_HOST"
          value = var.db_host
        }
        env {
          name = "DB_USER"
          value = var.db_user
        }
        env {
          name = "DB_PASSWORD"
          value = var.db_password
        }
        env {
          name = "EMAIL_FROM"
          value = var.email_from
        }
        env {
          name = "EMAIL_USERNAME"
          value = var.email_username
        }
        env {
          name = "EMAIL_PASSWORD"
          value = var.email_password
        }
        # メモリ1Gib
        resources {
          limits = { "memory" : "1Gi" }
        }
      }
    }
    metadata {
      annotations = {
        "run.googleapis.com/cloudsql-instances" = var.cloud_sql_instance
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

# ドメインを紐づける
resource "google_cloud_run_domain_mapping" "default" {
  location = google_cloud_run_service.default.location
  name     = var.domain

  metadata {
    namespace = var.project
  }

  spec {
    route_name = google_cloud_run_service.default.name
  }
}

