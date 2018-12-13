package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var listConfig = &cobra.Command{
	Use:   "list",
	Short: "This command is for kubeconfig from kcm",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("This is list config")
	},
}
