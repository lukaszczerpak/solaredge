package deletedata

import (
	"solaredge-sync/common"
	log "solaredge-sync/logging"
	"solaredge-sync/solaredgedb"
	"solaredge-sync/util"
	"time"
)

func Delete(cfg *common.AppConfig, startTime, stopTime time.Time) {
	// <START_DATE, STOP_DATE> range is inclusive thus STOP_DATE's time must be 23:59:59
	stopTime = time.Date(stopTime.Year(), stopTime.Month(), stopTime.Day(), 23, 59, 59, 0, stopTime.Location())

	db := solaredgedb.New(cfg.Influxdb.Url, cfg.Influxdb.Token, cfg.Influxdb.Org, cfg.Influxdb.Bucket)
	defer db.Close()

	log.Infof("Period from %v to %v => deleting data",
		startTime.Format(util.DATE_FORMAT), stopTime.Format(util.DATE_FORMAT))

	db.DeleteMeasurements(cfg, startTime, stopTime)
}
