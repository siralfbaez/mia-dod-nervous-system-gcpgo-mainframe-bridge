# NIST 800-53 Control Mapping (High Baseline)

| Control ID | Name | Implementation Detail |
| :--- | :--- | :--- |
| **AC-2** | Account Management | GKE Role-Based Access Control (RBAC) integrated with Google Cloud IAM. |
| **AC-17** | Remote Access | Mainframe-to-Cloud traffic restricted to Dedicated Interconnect or IPsec VPN via Cloud Router. |
| **AU-3** | Content of Audit Records | All EBCDIC-to-UTF8 translation events logged via Cloud Logging with non-repudiation. |
| **IA-9** | Service Identification | Services use Workload Identity to authenticate to AlloyDB without hardcoded secrets. |
| **SC-12** | Cryptographic Key Establishment | KMS handles automated 90-day rotation for all data-at-rest encryption keys. |