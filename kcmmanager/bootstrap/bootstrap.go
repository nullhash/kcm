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

package bootstrap

import (
	"log"
	"os"

	"github.com/nullhash/kcm/kcmmanager/kcmconfig"
)

func setupKCMHome() {
	if _, err := os.Stat("/home/shovan/.kcm"); os.IsNotExist(err) {
		log.Println("kcm home not available creating kcm home")
		os.Mkdir("/home/shovan/.kcm", os.ModePerm)
		return
	}
	log.Println("kcm home available skip creating kcm home")
}

func setupKCMConfig() {
	if _, err := os.Stat("/home/shovan/.kcm/config"); os.IsNotExist(err) {
		if file, err := os.Create("/home/shovan/.kcm/config"); err == nil {
			log.Println("kcm config not available creating kcm config")
			defer file.Close()
			return
		}
		log.Fatal("unable to create kcm config")
		return
	}
	log.Println("kcm config available skip creating kcm config")

}

func Bootstrap() {
	setupKCMHome()
	setupKCMConfig()
	defaultKubeConfigPath := "/home/shovan/.kube/config"
	kcmconfig.LoadConfig(defaultKubeConfigPath)
}
