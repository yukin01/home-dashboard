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
      }
    }
  }
}

# gcloud run services add-iam-policy-binding pubsub-tutorial \
#   --member=serviceAccount:cloud-run-pubsub-invoker@PROJECT_ID.iam.gserviceaccount.com \
#   --role=roles/run.invoker
resource "google_cloud_run_service_iam_member" "member" {
  location = google_cloud_run_service.worker.location
  project  = google_cloud_run_service.worker.project
  service  = google_cloud_run_service.worker.name
  role     = "roles/run.invoker"
  member   = "serviceAccount:${google_service_account.worker_pubsub.email}"
}
