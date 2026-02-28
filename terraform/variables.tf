variable "project_id" {
  type        = string
  description = "The GCP Project ID where the bridge will be deployed"
}

variable "region" {
  type        = string
  description = "The primary GCP region for all resources"
  default     = "us-central1"
}

variable "environment" {
  type        = string
  description = "The environment name (e.g., dev, prod)"
  default     = "dev"
}

variable "mainframe_ip_range" {
  type        = string
  description = "The CIDR range for the legacy mainframe interconnect"
  default     = "10.0.0.0/24"
}
