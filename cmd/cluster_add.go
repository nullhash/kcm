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

	"github.com/nullhash/kcm/util"
	"github.com/spf13/cobra"
)

var (
	clusterAddCommandHelpText = `This command is to add cluster config`
)

// addCmd represents the cluster add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is to add cluster config",
	Long:  clusterAddCommandHelpText,
	Run: func(addCmd *cobra.Command, args []string) {

		// To check argument is passed or missing
		if len(args) == 0 {
			fmt.Println("cluster name is missing..")
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

		// To check config file exists on the given path by the user
		if !util.CheckFileOrDirectoryExists(clusterOptions.config) {
			fmt.Println("Error: config file does not exist on the given path!")
			os.Exit(1)
		}

		// // To check cluster directory exists or not
		if util.CheckFileOrDirectoryExists(clusterPath) {
			fmt.Println("Error: Cluster name " + args[0] + " already exists!")
			os.Exit(1)
		} else {
			err = os.MkdirAll(clusterPath, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		// copy confile file from users specified location to .kcm cluster directory
		_, err = util.CopyConfigFile(clusterOptions.config, clusterConfigPath)
		if err != nil {
			// To delete the clustername directory if something fails, so that again same cluster name can be retried.
			errDeleteDirectory := util.DeleteDirectory(clusterPath)
			if errDeleteDirectory != nil {
				fmt.Println("error while deleting directory on fail. Try again with changing the cluster name or manually delete the directory with cluster name provided by you inside $HOME/.kcm/<cluster_name>. Error - ", errDeleteDirectory.Error())
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Cluster config " + args[0] + ", created successfully!")
	},
}

func init() {
	clusterCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&clusterOptions.config, "config", "", clusterOptions.config, "Absolute path of cluster kubeconfig file")
	addCmd.MarkFlagRequired("config")
}
