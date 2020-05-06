# gcloud run deploy pubsub-tutorial --image gcr.io/PROJECT_ID/pubsub
resource "google_cloud_run_service" "worker" {
  name     = "worker-service"
  location = var.gcp_region

  template {
    spec {
      containers {
        image = "us.gcr.io/${var.gcp_project}/worker:latest"
        env {
          name  = "REMO_ACCESS_TOKEN"
          value = var.remo_access_token
        }
        env {
          name  = "PROJECT_ID"
          value = var.gcp_project
        }
      }
      service_account_name = google_service_account.worker_run.email
    }
  }

  # https://github.com/terraform-providers/terraform-provider-google/issues/5898
  autogenerate_revision_name = true
}

resource "google_service_account" "worker_run" {
  account_id   = "worker-run"
  display_name = "Runtime service account for Cloud Run"
}

# gcloud run services add-iam-policy-binding pubsub-tutorial \
#   --member=serviceAccount:cloud-run-pubsub-invoker@PROJECT_ID.iam.gserviceaccount.com \
#   --role=roles/run.invoker
resource "google_cloud_run_service_iam_member" "worker_pubsub_run_invoker" {
  location = google_cloud_run_service.worker.location
  project  = google_cloud_run_service.worker.project
  service  = google_cloud_run_service.worker.name
  role     = "roles/run.invoker"
  member   = "serviceAccount:${google_service_account.worker_pubsub.email}"
}
