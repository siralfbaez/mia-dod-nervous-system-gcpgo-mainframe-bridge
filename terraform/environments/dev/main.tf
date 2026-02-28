module "vpc" {
  source = "../../modules/vpc-network"
  region = "us-central1"
}

module "kms" {
  source = "../../modules/kms-encryption"
  region = "us-central1"
}

module "database" {
  source     = "../../modules/alloydb"
  vpc_id     = module.vpc.vpc_id
  kms_key_id = module.kms.key_id
}

module "compute_gke" {
  source    = "../../modules/gke-cluster"
  vpc_id    = module.vpc.vpc_id
  subnet_id = module.vpc.subnet_id
}