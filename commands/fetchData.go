package commands

import (
	"solaredge/common"
	log "solaredge/logging"
	"solaredge/solaredgeapi"
	"solaredge/solaredgedb"
	"solaredge/util"
	"time"

	"github.com/spf13/cobra"
)

// fetchDataCmd represents the fetchData command
var fetchDataCmd = &cobra.Command{
	Use:   "fetch START_DATE STOP_DATE",
	Short: "Fetches measurements from SolarEdge API and stores in database.",
	Long:  `Dates must be provided in a format YYYY-MM-DD, ie.: 2021-08-13`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		startTime, err := time.ParseInLocation(util.DATE_FORMAT, args[0], config.General.Location)
		if err != nil {
			log.Fatalf("Invalid start time: %v", err)
		}

		stopTime, err := time.ParseInLocation(util.DATE_FORMAT, args[1], config.General.Location)
		if err != nil {
			log.Fatalf("Invalid stop time: %v", err)
		}

		fetchData(&config, startTime, stopTime)
	},
}

func init() {
	fetchDataCmd.Flags().BoolVar(&config.General.DeleteBeforeWrite, "delete", false, "Delete old data before writing new data")
	rootCmd.AddCommand(fetchDataCmd)
}

func fetchData(cfg *common.AppConfig, startTime, stopTime time.Time) {
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
			db.DeleteMeasurements(rangeStartTime, rangeStopTime)
		}

		log.Infof("Period from %v to %v => Writing %d telemetries to database",
			rangeStartTime.Format(util.DATE_FORMAT), rangeStopTime.Format(util.DATE_FORMAT), len(data.Data.Telemetries))
		db.WriteMeasurements(cfg, data.Data.Telemetries)
	}
}
