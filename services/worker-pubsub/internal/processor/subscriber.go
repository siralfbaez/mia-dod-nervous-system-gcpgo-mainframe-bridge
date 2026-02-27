package processor

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

type PassengerRecord struct {
	PNR      string `json:"pnr"`
	Name     string `json:"name"`
	FlightID string `json:"flight_id"`
}

type RecordProcessor struct {
	DB *sql.DB
}

// ProcessMessage simulates a Pub/Sub message handling loop
func (p *RecordProcessor) ProcessMessage(ctx context.Context, msg []byte) error {
	var record PassengerRecord
	if err := json.Unmarshal(msg, &record); err != nil {
		return fmt.Errorf("failed to unmarshal cloud-native record: %v", err)
	}

	// Upsert logic for AlloyDB/PostgreSQL
	query := `
		INSERT INTO reservations (pnr, passenger_name, flight_id) 
		VALUES ($1, $2, $3)
		ON CONFLICT (pnr) DO UPDATE SET passenger_name = EXCLUDED.passenger_name;`

	_, err := p.DB.ExecContext(ctx, query, record.PNR, record.Name, record.FlightID)
	if err != nil {
		return fmt.Errorf("database metabolic failure: %v", err)
	}

	log.Printf("📥 Successfully processed PNR: %s into AlloyDB", record.PNR)
	return nil
}
