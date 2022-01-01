package commands

import (
	"solaredge-sync/deletedata"
	log "solaredge-sync/logging"
	"solaredge-sync/util"
	"time"

	"github.com/spf13/cobra"
)

var deleteDataCmd = &cobra.Command{
	Use:   "delete-data START_DATE STOP_DATE",
	Short: "Deletes measurements from database for the specified time window.",
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

		deletedata.Delete(&config, startTime, stopTime)
	},
}

func init() {
	rootCmd.AddCommand(deleteDataCmd)
}
