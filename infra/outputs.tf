output "app_engine_custom_domain_record" {
  value = google_app_engine_domain_mapping.this.resource_records[0]
}
