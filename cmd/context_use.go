package cmd

import (
	"log"

	"github.com/nullhash/kcm/kcmmanager/context"

	"github.com/spf13/cobra"
)

var useContext = &cobra.Command{
	Use:   "use",
	Short: "This command is for kubeconfig from kcm",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("no context provided")
			return
		}
		if len(args) > 1 {
			log.Println("more than one context not supported")
			return
		}
		context.UseContext(args[0])
	},
}
