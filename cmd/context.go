package cmd

import (
	"github.com/spf13/cobra"
)

type ContextOptions struct{}

var (
	contextOptions = &ContextOptions{}
)

var context = &cobra.Command{
	Use:   "context",
	Short: "A brief description of your command",
	Long:  "",
}

func init() {
	context.AddCommand(
		listContext,
		deleteContext,
		resetContext,
	)
}
