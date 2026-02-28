# Terraform Infrastructure Documentation

## Overview
The infrastructure is managed as a set of modular components to ensure environment parity between `dev` and `prod`.

## Module Dependency Graph
1. **Network Layer:** Provisions a VPC with custom subnets and Private Google Access.
2. **Security Layer:** Deploys KMS KeyRings for CMEK (Customer Managed Encryption Keys).
3. **IAM Layer:** Configures Service Accounts with Least Privilege (e.g., `roles/pubsub.publisher`).
4. **Data Layer:** Deploys an AlloyDB Cluster with automated backups and encryption.
5. **Compute Layer:** Deploys a GKE Autopilot cluster and a Cloud Run Gateway.



## How to Deploy
1. Navigate to `terraform/environments/dev`.
2. Run `terraform init`.
3. Run `terraform plan -var-file="terraform.tfvars"`.
4. Run `terraform apply`.
