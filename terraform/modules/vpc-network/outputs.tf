output "vpc_id" {
  value = google_compute_network.main_vpc.id
}

output "subnet_id" {
  value = google_compute_subnetwork.gke_subnet.id
}