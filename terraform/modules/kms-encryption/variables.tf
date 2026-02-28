variable "region" {
  type        = string
  description = "The GCP region for KMS"
}

variable "project_id" {
  type        = string
  description = "The GCP Project ID"
  default     = ""
}