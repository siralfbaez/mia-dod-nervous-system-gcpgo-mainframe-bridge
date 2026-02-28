resource "google_kms_key_ring" "main" {
  name     = "mainframe-bridge-keyring"
  location = var.region
}

resource "google_kms_crypto_key" "alloydb_key" {
  name            = "alloydb-key"
  key_ring        = google_kms_key_ring.main.id
  rotation_period = "7776000s" # 90 days
}