package solaredgedb

import (
	"context"
	log "solaredge-sync/logging"
	"time"
)

func (db *SolarEdgeDB) DeleteMeasurements(startTime, stopTime time.Time) {
	deleteAPI := db.client.DeleteAPI()
	err := deleteAPI.DeleteWithName(context.Background(), db.org, db.bucket, startTime, stopTime, "")
	if err != nil {
		log.Errorf("Error when deleting data: %v", err)
	}
}
