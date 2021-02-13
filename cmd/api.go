package cmd

import (
	"log"

	"github.com/irmf/chainquery/apiactions"
	"github.com/irmf/chainquery/config"
	"github.com/irmf/chainquery/db"
	"github.com/irmf/chainquery/lbrycrd"
	swagger "github.com/irmf/chainquery/swagger/apiserver"
	"github.com/irmf/chainquery/twilio"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start only the api server",
	Long:  `This runs the API Server for chainquery only. The daemon does not run, however, the db is still required and all APIs are available.`,
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
		swagger.InitApiServer(config.GetAPIHostAndPort())
	},
}
