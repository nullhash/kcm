/*
Copyright 2018 The nullhash Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		// checkConfigExists(usr.HomeDir + "/.kcm/maya-staging/config")
		files, err := ioutil.ReadDir(usr.HomeDir + "/.kcm/")
		if err != nil {
			fmt.Println("Nothing Found here! :(")
		}
		for _, f := range files {
			// fmt.Println(f.Name())
			configPath := usr.HomeDir + "/.kcm/" + f.Name() + "/config"
			if checkConfigExists(configPath) {
				fmt.Println(f.Name())
			}
		}
	},
}

func init() {
	clusterCmd.AddCommand(listCmd)

}

func checkConfigExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}
