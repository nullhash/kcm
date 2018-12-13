package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var resetConfig = &cobra.Command{
	Use:   "reset",
	Short: "This command is for kubeconfig from kcm",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("This is reset config")
	},
}
