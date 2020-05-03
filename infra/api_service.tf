resource "google_project_service" "cloud_run" {
  service                    = "run.googleapis.com"
  disable_dependent_services = true
}
