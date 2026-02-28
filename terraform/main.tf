# 1. Network Layer (The Skeleton)
module "vpc" {
  source     = "./modules/vpc-network"
  project_id = var.project_id
  region     = var.region
}

# 2. Security & Identity (The Vault)
module "kms" {
  source     = "./modules/kms-encryption"
  project_id = var.project_id
  region     = var.region
}

# 3. Permissions (The Shield)
module "iam" {
  source     = "./modules/iam-roles"
  project_id = var.project_id
}

# 4. Messaging (The Synapses)
module "pubsub" {
  source     = "./modules/pubsub"
  project_id = var.project_id
}

# 5. Entry Point (The Front Door - Signal Gateway)
module "cloud_run_gateway" {
  source          = "./modules/cloud-run"
  project_id      = var.project_id
  region          = var.region
  vpc_id          = module.vpc.vpc_id
  pubsub_topic    = module.pubsub.topic_name
  service_account = module.iam.bridge_sa_email
}

# 6. Data & Metabolism (The Heart - AlloyDB)
module "alloydb" {
  source     = "./modules/alloydb"
  project_id = var.project_id
  region     = var.region
  vpc_id     = module.vpc.vpc_id
  kms_key_id = module.kms.key_id
}

# 7. Compute (The Brain - GKE Processing Cluster)
module "gke" {
  source     = "./modules/gke-cluster"
  project_id = var.project_id
  region     = var.region
  vpc_id     = module.vpc.vpc_id
  subnet_id  = module.vpc.subnet_id
}