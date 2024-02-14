package cmd

import "github.com/spf13/cobra"

func RegisterCmd(rootCmd *cobra.Command) {

	// shared core commands:
	// server:
	rootCmd.AddCommand(ServerCmd())
	// migrations:
	rootCmd.AddCommand(MigrateCmd())
	//swagger:
	rootCmd.AddCommand(SwaggerCmd())

	// module specific commands will go here:
}
