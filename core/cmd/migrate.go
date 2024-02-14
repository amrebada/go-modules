package cmd

import (
	"os"

	"github.com/amrebada/go-modules/core"
	"github.com/amrebada/go-modules/modules"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

// MigrateCmd is a command to run migrations
func MigrateCmd() (migrateCmd *cobra.Command) {
	migrateCmd = &cobra.Command{
		Use:   "migrate",
		Short: "Run auto migration",
		Long:  `Run auto migration`,
	}
	migrateCmd.Run = func(cmd *cobra.Command, args []string) {

		env := "dev"

		if os.Getenv("ENV") != "" && lo.Contains([]string{"dev", "prod", "test"}, os.Getenv("ENV")) {
			env = os.Getenv("ENV")
		}

		isMigrate := true
		isSwagger := false
		core.NewConfig(core.Stage(env)).
			SetMigrate(isMigrate).
			SetSwagger(isSwagger)

		app := core.NewServer()
		app.MainModule = modules.NewAppModule()
		app.RegisterMainModule()
		os.Exit(0)

	}
	return
}
