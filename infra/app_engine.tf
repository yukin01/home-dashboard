resource "google_app_engine_application" "default" {
  project     = data.google_project.this.project_id
  location_id = var.gcp_app_engine_region
}
