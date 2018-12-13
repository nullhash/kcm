package cmd

import (
	"github.com/spf13/cobra"
)

type ConfigOptions struct {
	configPath string
}

var (
	configOptions = &ConfigOptions{}
)

var config = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long:  "",
}

func init() {
	config.AddCommand(
		listConfig,
		deleteConfig,
		resetConfig,
	)
}
