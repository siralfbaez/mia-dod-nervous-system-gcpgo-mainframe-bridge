-- High-level setup for iSeries Journaling to feed CDC
-- To capture deltas without high overhead
CHGJRNE JRN(RESERVATIONS/RESEV_JRN) JRNSTATE(*ENABLED)
STRJRNPF FILE(RESERVATIONS/PASSENGER_MANIFEST) JRN(RESERVATIONS/RESEV_JRN) IMAGES(*BOTH)
-- This ensures CDC tools like Debezium or DataFusion can read the EBCDIC stream