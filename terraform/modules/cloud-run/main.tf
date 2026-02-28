resource "google_cloud_run_v2_service" "gateway" {
  name     = "signal-gateway"
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = "gcr.io/${var.project_id}/signal-gateway:latest"
      env {
        name  = "GCP_PROJECT_ID"
        value = var.project_id
      }
    }
    vpc_access {
      connector = var.vpc_connector_id
      egress    = "ALL_TRAFFIC"
    }
  }
}
resource "google_cloud_run_v2_service" "signal_gateway" {
  name     = "signal-gateway"
  location = var.region

  template {
    containers {
      image = "gcr.io/${var.project_id}/signal-gateway:latest"
    }
  }
}