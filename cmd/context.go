package cmd

import (
	"github.com/spf13/cobra"
)

type ContextOptions struct{}

var (
	contextOptions = &ContextOptions{}
)

var contextCommand = &cobra.Command{
	Use:   "context",
	Short: "A brief description of your command",
	Long:  "",
}

func init() {
	contextCommand.AddCommand(
		listContext,
		useContext,
		deleteContext,
		resetContext,
	)
}
