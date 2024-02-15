package main

import (
	"os"

	"github.com/amrebada/go-modules/modules"
	"github.com/amrebada/neon-core/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "go-modules",
		Short: "go-modules module CLI",
		Long:  `go-modules module CLI`,
	}
	mainModule := modules.NewAppModule()
	cmd.RegisterCmd(rootCmd, mainModule)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
