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
	"errors"
	"fmt"
	"io"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	clusterAddCommandHelpText = `This command is to add cluster config`
)

// func CToGoString(c []byte) string {
// 	n := -1
// 	for i, b := range c {
// 		if b == 0 {
// 			break
// 		}
// 		n = i
// 	}
// 	return string(c[:n+1])
// }

// addCmd represents the cluster add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is to add cluster config",
	Long:  clusterAddCommandHelpText,
	Run: func(addCmd *cobra.Command, args []string) {

		// stores the home dir of the user
		// used to know the location of .kcm directory as .kcm directory
		// is in home
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if checkFileOrDirectoryExists(home + "/.kcm/" + clusterOptions.clusterName) {
			fmt.Println("error: clustername should be unique..")
			os.Exit(1)
		} else {
			err = os.MkdirAll(home+"/.kcm/"+clusterOptions.clusterName, 0755)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}

		_, err = copyConfigFile(clusterOptions.filePath, home+"/.kcm/"+clusterOptions.clusterName+"/config")
		if err != nil {
			// To delete the clustername directory if something fails, so that again same cluster name can be retried.
			errDeleteDirectory := deleteDirectory(home + "/.kcm/" + clusterOptions.clusterName)
			if errDeleteDirectory != nil {
				fmt.Println("error while deleting directory on fail. Try again with changing the cluster name or manually delete the directory with cluster name provided by you inside $HOME/.kcm/<cluster_name>")
			}
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Cluster config created successfully..")
	},
}

func init() {
	clusterCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&clusterOptions.clusterName, "clustername", "", clusterOptions.clusterName, "unique cluster name")
	addCmd.MarkFlagRequired("clustername")

	addCmd.Flags().StringVarP(&clusterOptions.filePath, "filepath", "", clusterOptions.filePath, "absolute path of cluster kubeconfig file")
	addCmd.MarkFlagRequired("filepath")
}

// checkFileOrDirectoryExists to check directory or file is present on the given path.
func checkFileOrDirectoryExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// copyConfigFile copies file content from sourc (src) to destination (dst) path.
func copyConfigFile(src, dst string) (int64, error) {
	if !checkFileOrDirectoryExists(src) {
		return 0, errors.New("error: config file does not exist on given path")
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
