package solaredgedb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"solaredge/common"
	log "solaredge/logging"
	"solaredge/solaredgeapi/model"
)

func (db *SolarEdgeDB) WriteMeasurements(cfg *common.AppConfig, telemetries []model.InverterTelemetry) {
	writeAPI := db.client.WriteAPI(db.org, db.bucket)
	errorsCh := writeAPI.Errors()
	go func() {
		for err := range errorsCh {
			log.Errorf("Error writing to database: %v", err)
		}
	}()
	for _, t := range telemetries {
		ts, fields, err := InverterTelemetryToInflux(t, cfg.General.Location)
		if err != nil {
			log.Errorf("Error when preparing data to write: %v", err)
			log.Debugf("Skipped telemetry: %+v", t)
			continue
		}

		tags := map[string]string{
			"site":     cfg.SolarEdge.SiteId,
			"slave_id": "1",
			"sn":       cfg.SolarEdge.InverterSn,
		}

		p := influxdb2.NewPoint(cfg.Influxdb.Measurement, tags, fields, ts)
		writeAPI.WritePoint(p)
	}

	writeAPI.Flush()
}
