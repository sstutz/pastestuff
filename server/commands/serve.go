package commands

import (
	"github.com/spf13/cobra"
	"github.com/sstutz/pastestuff/server/routers"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the application's web server",
	Long:  "Starts the application's web server",
	Run:   serve,
}

func init() {
	RootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	routers.ListenAndServe()
}
