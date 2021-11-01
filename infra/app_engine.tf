resource "google_app_engine_application" "this" {
  project     = data.google_project.this.project_id
  location_id = var.gcp_app_engine_region
}

resource "google_app_engine_domain_mapping" "this" {
  domain_name = var.custom_domain_name

  ssl_settings {
    ssl_management_type = "AUTOMATIC"
  }
}
