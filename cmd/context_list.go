package cmd

import (
	"github.com/nullhash/kcm/kcmmanager/kcmcontext"
	"github.com/spf13/cobra"
)

var listContext = &cobra.Command{
	Use:   "list",
	Short: "This command is for kubeconfig from kcm",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		kcmcontext.ListContext()
	},
}
