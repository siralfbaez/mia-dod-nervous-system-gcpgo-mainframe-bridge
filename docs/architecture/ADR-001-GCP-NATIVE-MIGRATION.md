# ADR-001: Legacy Mainframe to GCP Native Migration Strategy

## Status
Proposed (Draft for Review)

## Context
The existing Airline Reservation and Inventory system is a 15+ year-old C/C++ monolith running on legacy Mainframe architecture (z/OS / iSeries). It utilizes EBCDIC encoding, VSAM/DB2 data structures, and handles high-throughput passenger transactions. 

The business goal is to modernize this into a scalable, cloud-native environment on **Google Cloud Platform (GCP)** to improve agility, reduce MIPS costs, and enable AI-driven insights.

### Challenges
* **Data Integrity:** Character encoding mismatch (EBCDIC to UTF-8/ASCII).
* **High Stakes:** Zero-downtime requirement for reservation availability.
* **Tightly Coupled Logic:** Monolithic C/C++ modules with hard-to-trace dependencies.

## Decision
We will adopt the **Strangler Fig Pattern** facilitated by an intermediary **"Nervous System" Architecture** built in Go.

### 1. Phased Decoupling (Strangler Fig)
Instead of a "Big Bang" migration, we will incrementally replace specific functional domains (e.g., Seat Selection, Loyalty) as independent microservices on **Google Kubernetes Engine (GKE)**.

### 2. Integration Layer (The Nervous System)
* **API Gateway (Apigee):** Expose legacy mainframe functions as REST/gRPC endpoints.
* **Translation Engine:** A dedicated Go-based service to handle real-time **CCSID mapping** and **EBCDIC/UTF-8 conversion** to ensure data survival across environments.

### 3. Data Metabolism
* **CDC (Change Data Capture):** Utilize iSeries Journaling to stream delta changes to **AlloyDB (PostgreSQL)** in real-time.
* **Dual-Write Strategy:** During the coexistence phase, the system will maintain state across both the Mainframe and GCP to allow for immediate rollback.



## Consequences

### Positive
* **Risk Mitigation:** Each microservice can be tested and rolled back independently.
* **Cost Efficiency:** Offloading "read-heavy" traffic to GCP immediately reduces Mainframe MIPS consumption.
* **Scalability:** Enables the use of **Vertex AI** for predictive pricing and inventory management.

### Negative
* **Increased Complexity:** Maintaining data consistency between the Mainframe and GCP requires strict **Contract Validation**.
* **Latency:** The "hop" between GCP and on-prem Mainframe adds milliseconds; mitigated by caching strategies in `pkg/resilience`.

## Compliance & Security
* **FedRAMP High Alignment:** All data in transit is encrypted via TLS 1.3, and data at rest utilizes GCP KMS (Key Management Service).
* **Auditability:** Full logging via Cloud Logging to meet SOX and ITAR requirements.
