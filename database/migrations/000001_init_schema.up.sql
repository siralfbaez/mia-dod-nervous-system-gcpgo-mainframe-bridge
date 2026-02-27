CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE reservations (
    pnr VARCHAR(6) PRIMARY KEY,
    passenger_name TEXT NOT NULL,
    flight_id VARCHAR(10),
    seat_assignment VARCHAR(4),
    last_updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    sync_status VARCHAR(20) DEFAULT 'PENDING'
);

CREATE INDEX idx_reservations_flight ON reservations(flight_id);