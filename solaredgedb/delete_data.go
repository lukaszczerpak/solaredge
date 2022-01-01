package solaredgedb

import (
	"context"
	"solaredge-sync/common"
	log "solaredge-sync/logging"
	"time"
)

func (db *SolarEdgeDB) DeleteMeasurements(cfg *common.AppConfig, startTime, stopTime time.Time) {
	deleteAPI := db.client.DeleteAPI()
	err := deleteAPI.DeleteWithName(context.Background(), cfg.Influxdb.Org, cfg.Influxdb.Bucket, startTime, stopTime, "")
	if err != nil {
		log.Errorf("Error when deleting data: %v", err)
	}
}
