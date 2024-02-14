package cmd

import (
	"log"
	"os"

	"github.com/amrebada/go-modules/core"
	"github.com/amrebada/go-modules/modules"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// Server is a command to run server
func ServerCmd() (serverCmd *cobra.Command) {
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run module server",
		Long:  `Run module server`,
	}
	serverCmd.Run = func(cmd *cobra.Command, args []string) {

		env := "dev"

		if os.Getenv("ENV") != "" && lo.Contains([]string{"dev", "prod", "test"}, os.Getenv("ENV")) {
			env = os.Getenv("ENV")
		}

		isMigrate := false
		isSwagger := false
		core.NewConfig(core.Stage(env)).
			SetMigrate(isMigrate).
			SetSwagger(isSwagger)

		app := core.NewServer()
		app.MainModule = modules.NewAppModule()
		app.RegisterMainModule()
		err := app.Start()
		if err != nil {
			log.Println(err)
		}

	}
	return
}
