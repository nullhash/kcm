// Copyright © 2018 kcm
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
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/harshvkarn/kcm/util"
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

		if len(args) == 0 {
			fmt.Println("cluster name is missing..")
			os.Exit(1)
		}

		home, err := util.GetHomeDir()
		if err != nil {
			log.Fatal("Unable to get home dir, Please try again. error - " + err.Error())
		}

		if util.CheckFileOrDirectoryExists(home + "/.kcm/" + args[0]) {
			fmt.Println("Error: Cluster name " + args[0] + " already exists!")
			os.Exit(1)
		} else {
			err = os.MkdirAll(home+"/.kcm/"+args[0], 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		_, err = copyConfigFile(clusterOptions.config, home+"/.kcm/"+args[0]+"/config")
		if err != nil {
			// To delete the clustername directory if something fails, so that again same cluster name can be retried.
			errDeleteDirectory := deleteDirectory(home + "/.kcm/" + args[0])
			if errDeleteDirectory != nil {
				fmt.Println("error while deleting directory on fail. Try again with changing the cluster name or manually delete the directory with cluster name provided by you inside $HOME/.kcm/<cluster_name>")
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

// copyConfigFile copies file content from sourc (src) to destination (dst) path.
func copyConfigFile(src, dst string) (int64, error) {
	if !util.CheckFileOrDirectoryExists(src) {
		return 0, errors.New("Error: config file does not exist on the given path")
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// deleteDirectory is to delete directory in a given path.
func deleteDirectory(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
