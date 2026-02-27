resource "google_alloydb_cluster" "main_cluster" {
  cluster_id = var.cluster_id
  location   = var.region
  network    = var.vpc_id

  initial_user {
    password = var.db_password
    user     = "admin_architect"
  }
}

resource "google_alloydb_instance" "primary_instance" {
  cluster       = google_alloydb_cluster.main_cluster.name
  instance_id   = "${var.cluster_id}-primary"
  instance_type = "PRIMARY"

  machine_config {
    cpu_count = 4
  }

  # High Availability configuration for FedRAMP/Mission-Critical workloads
  availability_type = "REGIONAL"
}
