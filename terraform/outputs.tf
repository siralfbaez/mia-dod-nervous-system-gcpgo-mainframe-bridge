output "gateway_url" {
  description = "The public URL of the Signal Gateway (Cloud Run)"
  value       = module.cloud_run_gateway.service_url
}

output "gke_cluster_endpoint" {
  description = "The IP address of the GKE cluster master"
  value       = module.gke.cluster_endpoint
}

output "alloydb_connection_name" {
  description = "The connection string for the AlloyDB cluster"
  value       = module.alloydb.cluster_id
}

output "kms_key_id" {
  description = "The resource ID of the encryption key"
  value       = module.kms.key_id
}