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

package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nullhash/kcm/types"
	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(filePath string) {
	home := os.Getenv("HOME")
	if home == "" {
		log.Println("error while reading environment variable")
		return
	}
	var kubeConfig types.KubeConfig
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println("unable to read file : ", err)
		return
	}
	err = yaml.Unmarshal(yamlFile, &kubeConfig)
	if err != nil {
		log.Println("unable to unmarshal config yaml")
		return
	}
	for _, context := range kubeConfig.Contexts {
		dirName := home + "/.kcm/" + context.Name
		tmpConfig := kubeConfig
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			os.Mkdir(dirName, os.ModePerm)
		}
		tmpConfig.CurrentContext = context.Name
		buffer, err := yaml.Marshal(tmpConfig)
		if err != nil {
			log.Println("unable to marshal config : ", err)
			continue
		}
		err = ioutil.WriteFile(dirName+"/config", buffer, 0600)
		if err != nil {
			log.Println("unable to update kcm config : ", err)
			continue
		}
		log.Println("kcm context created : ", dirName+"/config")
	}
}
