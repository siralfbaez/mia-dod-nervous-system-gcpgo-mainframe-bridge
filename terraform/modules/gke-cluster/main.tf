resource "google_container_cluster" "primary" {
  name     = "mia-dod-gke-cluster"
  location = var.region

  # Use the VPC we just created
  network    = var.vpc_id
  subnetwork = var.subnet_id

  # Enabling a Private Cluster for high-security workloads
  private_cluster_config {
    enable_private_nodes    = true
    enable_private_endpoint = false # Allow master access via authorized networks
    master_ipv4_cidr_block  = "172.16.0.0/28"
  }

  ip_allocation_policy {
    cluster_secondary_range_name  = "gke-pods"
    services_secondary_range_name = "gke-services"
  }

  # Removing default node pool to replace with specialized pools
  remove_default_node_pool = true
  initial_node_count       = 1
}

resource "google_container_node_pool" "primary_nodes" {
  name       = "mainframe-bridge-pool"
  location   = var.region
  cluster    = google_container_cluster.primary.name
  node_count = 3

  node_config {
    machine_type = "e2-standard-4"
    
    # IAM Service Account with least privilege
    service_account = var.service_account_email
    oauth_scopes    = ["https://www.googleapis.com/auth/cloud-platform"]

    labels = {
      workload = "nervous-system-core"
    }
  }
}
