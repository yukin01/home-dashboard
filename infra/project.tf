# gcloud projects add-iam-policy-binding PROJECT_ID \
#    --member=serviceAccount:service-PROJECT-NUMBER@gcp-sa-pubsub.iam.gserviceaccount.com \
#    --role=roles/iam.serviceAccountTokenCreator
resource "google_project_iam_member" "pubsub_token_creator" {
  role   = "roles/iam.serviceAccountTokenCreator"
  member = "serviceAccount:service-${data.google_project.this.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_project_iam_member" "cloud_run_firestore_user" {
  role   = "roles/datastore.user"
  member = "serviceAccount:${google_service_account.worker_run.email}"
}
