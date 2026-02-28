# Integration Standards: Mainframe-to-Cloud Bridge

## 1. Encoding Sovereignty (CCSID 37)
- **Standard:** All raw data originating from the IBM iSeries must be treated as `EBCDIC (CCSID 37)` byte-streams.
- **Rule:** No UTF-8 string assumptions shall be made until the data has passed through the `pkg/encoding-utils` translator.
- **Validation:** Every signal must include a checksum to ensure no byte-loss during the 0x41 (UTF-8) to 0xC1 (EBCDIC) conversion.

## 2. Resilience & Circuit Breaking
- **Thresholds:** All calls to legacy Mainframe modules must be wrapped in a `pkg/resilience` Circuit Breaker.
- **Fail-Fast:** Max timeout is **2000ms**. If 5 consecutive failures occur, the circuit opens for a **30-second cooling period**.
- **Observability:** Every "Open" state must be logged to Cloud Trace with the `service_impact=degraded` tag.

## 3. Contractual Integrity (Protobuf)
- **Schema:** All inter-service communication (Gateway -> Worker) must adhere to the `api/proto/signal.proto` definition.
- **Versioning:** No breaking changes to the proto schema. Fields must be marked `deprecated` before removal to maintain backward compatibility with in-flight Pub/Sub messages.