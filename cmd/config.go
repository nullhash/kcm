package cmd

import (
	"github.com/spf13/cobra"
)

var configCommand = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long:  "",
}

func init() {
	configCommand.AddCommand(
		listConfig,
		deleteConfig,
		resetConfig,
	)
}
