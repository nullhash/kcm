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
	"github.com/spf13/cobra"
)

// ClusterOptions stores information of cluster
type ClusterOptions struct {
	// clusterName string
	config string
}

var (
	clusterCommandHelpText = `
The following commands helps managing cluster, i.e. adding, removing or editing
Usage: kcm cluster <subcommand> [options] [args]
Examples:
 # List available Clusters:
   $ kcm cluster list
 # Use the available cluster:
   $ kcm cluster use --name <cluster-name>
`
	clusterOptions = &ClusterOptions{}
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "A brief description of your command",
	Long:  clusterCommandHelpText,
}

func init() {
	//	rootCmd.AddCommand(clusterCmd)

}

// Validate verifies whether a cluster name or file path is provided or not, followed by
// sub command. It returns nil and proceeds to execute the command if there is
// no error and returns an error if it is missing.
// func (c *ClusterOptions) Validate(cmd *cobra.Command, clusterNameVerify bool, filePathVerify bool) error {
// 	if clusterNameVerify {
// 		if len(c.clusterName) == 0 {
// 			return errors.New("error: --clustername not specified")
// 		}
// 	}
// 	if filePathVerify {
// 		if len(c.filePath) == 0 {
// 			return errors.New("error: --filepath not specified")
// 		}
// 	}
// 	return nil
// }
