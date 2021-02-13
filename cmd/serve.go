package cmd

import (
	"log"

	"github.com/irmf/chainquery/apiactions"
	"github.com/irmf/chainquery/config"
	"github.com/irmf/chainquery/daemon"
	"github.com/irmf/chainquery/db"
	"github.com/irmf/chainquery/lbrycrd"
	swagger "github.com/irmf/chainquery/swagger/apiserver"
	"github.com/irmf/chainquery/twilio"

	"github.com/pkg/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run Daemon routines and the API Server",
	Long:  `Run Daemon routines and the API Server, check github.com/irmf/chainquery#what-does-chainquery-consist-of`,
	Args:  cobra.OnlyValidArgs,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("codeprofile") {
			defer profile.Start(profile.NoShutdownHook).Stop()
		}
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

		go swagger.InitApiServer(config.GetAPIHostAndPort())
		daemon.DoYourThing()
	},
}
