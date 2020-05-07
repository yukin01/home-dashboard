# gcloud pubsub topics create myRunTopic
resource "google_pubsub_topic" "worker" {
  name = "worker-topic"
}

# gcloud iam service-accounts create cloud-run-pubsub-invoker \
#    --display-name "Cloud Run Pub/Sub Invoker"
resource "google_service_account" "worker_pubsub" {
  account_id   = "worker-pubsub"
  display_name = "Worker service invoker for pub/sub"
}

# gcloud pubsub subscriptions create myRunSubscription --topic myRunTopic \
#   --push-endpoint=SERVICE-URL/ \
#   --push-auth-service-account=cloud-run-pubsub-invoker@PROJECT_ID.iam.gserviceaccount.com
resource "google_pubsub_subscription" "worker" {
  name  = "worker-subscription"
  topic = google_pubsub_topic.worker.name

  push_config {
    push_endpoint = google_cloud_run_service.worker.status[0].url
    oidc_token {
      service_account_email = google_service_account.worker_pubsub.email
    }
  }
}

resource "google_cloud_scheduler_job" "worker" {
  name     = "worker-job"
  schedule = "0 */1 * * *"             # hourly
  region   = var.gcp_app_engine_region # temporarily

  pubsub_target {
    topic_name = google_pubsub_topic.worker.id
    data       = base64encode("Published by Cloud Scheduler")
  }
}
