# SOP: Incident Response (Nervous System Bridge)

## Severity 1: Data Corruption / Encoding Mismatch
- **Indicator:** `TranslatorError: Invalid EBCDIC sequence` spikes in Cloud Logging.
- **Action:** Immediately halt the `worker-pubsub` subscription to prevent corrupt data from reaching AlloyDB.
- **Mitigation:** Revert to the last known stable `mapping.go` and trigger a manual "Replay" of the dead-letter queue signals.

## Severity 2: Circuit Breaker Cascade
- **Indicator:** `pkg/resilience` state transitions to `OPEN`.
- **Action:** Notify the Mainframe Ops team of potential MIPS saturation.
- **Mitigation:** The AI Diagnostic Agent (`ai-agent`) will automatically analyze the last 100 signals to determine if the issue is a network blip or a mainframe deadlock.

## Severity 3: AI Diagnostic Drift
- **Indicator:** Gemini/Vertex AI returns low-confidence scores for system health.
- **Action:** Bypass AI-driven auto-scaling and return to hard-coded GKE horizontal pod autoscaler (HPA) limits.
- **Follow-up:** Review the Vertex AI `inference_test.go` logs to identify if the model requires a new system prompt or fine-tuning.