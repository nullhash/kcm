// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/harshvkarn/kcm/util"
	"github.com/spf13/cobra"
)

var (
	clusterUseCommandHelpText = `This command helps to change or use the cluster or config`
)

// useCmd represents the use command
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "This command helps to change or use the cluster or config",
	Long:  clusterUseCommandHelpText,
	Run: func(cmd *cobra.Command, args []string) {

		// To check argument is passed or missing
		if len(args) == 0 {
			fmt.Println("cluster name is missing. Use - kcm cluster use <cluster-name>")
			os.Exit(1)
		}

		// clusterPath contains cluster directory path ( $home/.kcm/<cluster-name>)
		clusterPath, err := util.GetClusterPath(args[0])
		if err != nil {
			log.Fatal("Unable to get cluster directory path. Error - ", err.Error())
		}

		// clusterConfigPath contains specific cluster config file path ( $home/.kcm/<cluster-name>/config)
		clusterConfigPath, err := util.GetClusterConfigPath(args[0])
		if err != nil {
			log.Fatal("Unable to get cluster config file path. Error - ", err.Error())
		}

		// kubeconfigiPath contains path to the config file which is used by kubectl
		kubeconfigPath, err := util.GetKubeconfigEnvValue()
		if err != nil {
			fmt.Println("Error - ", err.Error())
			os.Exit(1)
		}

		// To check cluster directory exists or not
		if !util.CheckFileOrDirectoryExists(clusterPath) {
			fmt.Println("cluster name does not exist. Use - kcm cluster add <cluster-name>--config=<config-file-path>")
			os.Exit(1)
		}

		// To check cluster directory contains config or not
		if !util.CheckFileOrDirectoryExists(clusterConfigPath) {
			log.Fatal("Unable to find config in cluster directory.")
		}

		// _, err = createLink(home+"/.kcm/"+args[0]+"/config", home+"/.kcm/config")
		// if err != nil {
		// 	log.Fatal("Unable to create link, Please try again. error - " + err.Error())
		// 	os.Exit(1)
		// }

		// Copy config file from cluster directory to KUBECONFIG path so to be used
		_, err = util.CopyConfigFile(clusterConfigPath, kubeconfigPath)
		if err != nil {
			log.Fatal("Unable to copy config, Please try again. error - " + err.Error())
		}
		fmt.Println("Using cluster " + args[0])

	},
}

func init() {
	clusterCmd.AddCommand(useCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// useCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// useCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// func createLink(source string, destination string) (string, error) {
// 	var cmd = "ln -sf " + source + " " + destination
// 	out, err := util.ExeCmd(cmd)
// 	if err != nil {
// 		return "", err
// 	}
// 	return out, nil
// }
