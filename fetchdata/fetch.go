package fetchdata

import (
	"solaredge-sync/common"
	log "solaredge-sync/logging"
	"solaredge-sync/solaredgeapi"
	"solaredge-sync/solaredgedb"
	"solaredge-sync/util"
	"time"
)

func Fetch(cfg *common.AppConfig, startTime, stopTime time.Time) {
	// <START_DATE, STOP_DATE> range is inclusive thus STOP_DATE's time must be 23:59:59
	stopTime = time.Date(stopTime.Year(), stopTime.Month(), stopTime.Day(), 23, 59, 59, 0, stopTime.Location())

	db := solaredgedb.New(cfg.Influxdb.Url, cfg.Influxdb.Token, cfg.Influxdb.Org, cfg.Influxdb.Bucket)
	defer db.Close()
	api := solaredgeapi.New(cfg.SolarEdge.ApiKey)

	for rangeStartTime := startTime; rangeStartTime.Before(stopTime); rangeStartTime = rangeStartTime.AddDate(0, 0, 7) {
		rangeStopTime := rangeStartTime.AddDate(0, 0, 7).Add(time.Second * -1)
		if rangeStopTime.After(stopTime) {
			rangeStopTime = stopTime
		}

		if !cfg.General.DeleteBeforeWrite {
			exists, err := db.CheckIfDataExists(rangeStartTime, rangeStopTime)
			if err != nil {
				log.Errorf("Checking DB failed: %v", err)
				continue
			}
			if exists {
				log.Infof("Period from %v to %v => data exists in db, skipping",
					rangeStartTime.Format(util.DATE_FORMAT), rangeStopTime.Format(util.DATE_FORMAT))
				continue
			}
		}

		log.Infof("Period from %v to %v => fetching data from SolarEdge API",
			rangeStartTime.Format(util.DATE_FORMAT), rangeStopTime.Format(util.DATE_FORMAT))

		data, err := api.FetchData(cfg.SolarEdge.SiteId, cfg.SolarEdge.InverterSn, rangeStartTime, rangeStopTime)
		if err != nil {
			log.Errorf("Error on fetching data from SolarEdge API: %v", err)
			continue
		}

		if cfg.General.DeleteBeforeWrite {
			log.Infof("Period from %v to %v => deleting old data, if exists",
				rangeStartTime.Format(util.DATE_FORMAT), rangeStopTime.Format(util.DATE_FORMAT))
			db.DeleteMeasurements(cfg, rangeStartTime, rangeStopTime)
		}

		log.Infof("Period from %v to %v => Writing %d telemetries to database",
			rangeStartTime.Format(util.DATE_FORMAT), rangeStopTime.Format(util.DATE_FORMAT), len(data.Data.Telemetries))
		db.WriteMeasurements(cfg, data.Data.Telemetries)
	}
}
