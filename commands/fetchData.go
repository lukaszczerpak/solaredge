package commands

import (
	"solaredge-sync/fetchdata"
	log "solaredge-sync/logging"
	"solaredge-sync/util"
	"time"

	"github.com/spf13/cobra"
)

// fetchDataCmd represents the fetchData command
var fetchDataCmd = &cobra.Command{
	Use:   "fetch-data START_DATE STOP_DATE",
	Short: "Fetches measurements from SolarEdge API and stores in database.",
	Long:  `Dates must be provided in a format YYYY-MM-DD, ie.: 2021-08-13`,
	Run: func(cmd *cobra.Command, args []string) {
		startTime, err := time.ParseInLocation(util.DATE_FORMAT, args[0], config.General.Location)
		if err != nil {
			log.Fatalf("Invalid start time: %v", err)
		}

		stopTime, err := time.ParseInLocation(util.DATE_FORMAT, args[1], config.General.Location)
		if err != nil {
			log.Fatalf("Invalid stop time: %v", err)
		}

		fetchdata.Fetch(&config, startTime, stopTime)
	},
}

func init() {
	fetchDataCmd.Flags().BoolVar(&config.General.DeleteBeforeWrite, "delete", false, "Delete old data before writing new data")
	rootCmd.AddCommand(fetchDataCmd)
}
