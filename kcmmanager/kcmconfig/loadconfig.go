package kcmconfig

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nullhash/kcm/types"
	yaml "gopkg.in/yaml.v2"
)

func LoadConfig(filePath string) {
	if filePath == "" {
		log.Println("no file path provided")
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
		dirName := "/home/shovan/.kcm/" + context.Name
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
