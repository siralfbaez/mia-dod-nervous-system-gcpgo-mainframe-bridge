resource "google_project_iam_member" "gke_vertex_ai" {
  project = var.project_id
  role    = "roles/aiplatform.user"
  member  = "serviceAccount:${var.gke_sa_email}"
}

resource "google_project_iam_member" "gke_alloydb" {
  project = var.project_id
  role    = "roles/alloydb.client"
  member  = "serviceAccount:${var.gke_sa_email}"
}
