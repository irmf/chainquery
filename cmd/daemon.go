package cmd

import (
	"log"

	"github.com/irmf/chainquery/apiactions"
	"github.com/irmf/chainquery/config"
	"github.com/irmf/chainquery/daemon"
	"github.com/irmf/chainquery/db"
	"github.com/irmf/chainquery/lbrycrd"
	"github.com/irmf/chainquery/twilio"
	"github.com/spf13/cobra"
)

func init() {
	serveCmd.AddCommand(daemonCmd)
}

var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run only Daemon routines",
	Long:  `Run only Daemon routines, without the API Server. Check github.com/irmf/chainquery#what-does-chainquery-consist-of`,
	Run: func(cmd *cobra.Command, args []string) {
		config.InitSlack()
		twilio.InitTwilio()
		apiactions.AutoUpdateCommand = config.GetAutoUpdateCommand()
		//Main Chainquery DB connection
		dbInstance, err := db.Init(config.GetMySQLDSN(), config.GetDebugQueryMode())
		if err != nil {
			log.Panic(err)
		}
		defer db.CloseDB(dbInstance)

		lbrycrdClient := lbrycrd.Init()
		defer lbrycrdClient.Shutdown()
		daemon.DoYourThing()
	},
}
