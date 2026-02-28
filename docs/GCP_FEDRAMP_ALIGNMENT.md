# GCP FedRAMP High Alignment Matrix

This bridge is architected to meet the **FedRAMP High** authorization boundary requirements for Aviation and Government data.

## Data Sovereignty & Locality
* **Regionality:** All processing is restricted to `us-central1` and `us-east1` to comply with US-only data residency requirements.
* **Encryption:** FIPS 140-2 validated modules are used for all TLS 1.3 encryption in transit.

## Boundary Protection
* **Cloud Armor:** Protects the `signal-gateway` (Cloud Run) from Layer 7 DDoS attacks.
* **VPC Service Controls:** Prevents data exfiltration by restricting service access to a defined perimeter.