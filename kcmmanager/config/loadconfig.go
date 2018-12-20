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
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/nullhash/kcm/types"
	yaml "gopkg.in/yaml.v2"
)

type loadConfigUtil struct {
	home       string
	filePath   string
	kcmConfig  types.KcmConfig
	kubeConfig types.KubeConfig
}

func newLoadConfigUtil(filePath string) (*loadConfigUtil, error) {
	lcu := loadConfigUtil{}
	home := os.Getenv("HOME")
	if home == "" {
		return nil, errors.New("error while reading environment variable")
	}
	lcu.home = home
	lcu.filePath = filePath
	return &lcu, nil
}

func (lcu *loadConfigUtil) readKubeConfig() error {
	var kubeConfig types.KubeConfig
	yamlFile, err := ioutil.ReadFile(lcu.filePath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &kubeConfig)
	if err != nil {
		return err
	}
	lcu.kubeConfig = kubeConfig
	return nil
}

func (lcu *loadConfigUtil) readKcmConfig() error {
	var kcmConfig types.KcmConfig
	kcmYaml, err := ioutil.ReadFile(lcu.home + "/.kcm/config")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(kcmYaml, &kcmConfig)
	if err != nil {
		return err
	}
	lcu.kcmConfig = kcmConfig
	return nil
}

/*
func (lcu *loadConfigUtil) deleteOldResources() {
	kcmContexts := []types.KcmContext{}
	configFiles := []string{}
	for _, context := range lcu.kcmConfig.Contexts {
		if context.ConfigFilePath != lcu.filePath {
			kcmContexts = append(kcmContexts, context)
			continue
		}
		if _, err := os.Stat(context.KcmContextPath); !os.IsNotExist(err) {
			os.RemoveAll(context.KcmContextPath)
		}
	}
	for _, configfile := range lcu.kcmConfig.ConfigFiles {
		if configfile != lcu.filePath {
			configFiles = append(configFiles, configfile)
		}
	}
	lcu.kcmConfig.Contexts = kcmContexts
	lcu.kcmConfig.ConfigFiles = configFiles
}
*/

func (lcu *loadConfigUtil) saveKcmContext() {
	for _, context := range lcu.kubeConfig.Contexts {
		fmt.Println("Provide dislplay name of context " + context.Name + " of cluster " + context.Context.Cluster)
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		rand.Seed(time.Now().UnixNano())
		dirName := lcu.home + "/.kcm/random-" + strconv.Itoa(rand.Intn(99999))
		kcmContext := types.KcmContext{
			Name:           input.Text(),
			ClusterName:    context.Context.Cluster,
			ContextName:    context.Name,
			ConfigFilePath: lcu.filePath,
			KcmContextPath: dirName,
		}
		tmpConfig := lcu.kubeConfig
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			os.Mkdir(dirName, os.ModePerm)
		}
		tmpConfig.CurrentContext = context.Name
		buffer, err := yaml.Marshal(tmpConfig)
		if err != nil {
			fmt.Println("unable to marshal config : ", err)
			continue
		}
		err = ioutil.WriteFile(dirName+"/config", buffer, 0600)
		if err != nil {
			fmt.Println("unable to update kcm config : ", err)
			continue
		}
		fmt.Println("kcm context created : ", dirName+"/config")
		lcu.kcmConfig.Contexts = append(lcu.kcmConfig.Contexts, kcmContext)
	}
	lcu.kcmConfig.ConfigFiles = append(lcu.kcmConfig.ConfigFiles, lcu.filePath)
}

func (lcu *loadConfigUtil) saveKcmConfig() error {
	buffer, err := yaml.Marshal(lcu.kcmConfig)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(lcu.home+"/.kcm/config", buffer, 0600)
	if err != nil {
		return err
	}
	return nil
}

func LoadConfig(filePath string) {
	lcm, err := newLoadConfigUtil(filePath)
	if err != nil {
		return
	}
	if err := lcm.readKubeConfig(); err != nil {
		return
	}

	if err := lcm.readKcmConfig(); err != nil {
		return
	}
	lcm.saveKcmContext()
	if err := lcm.saveKcmConfig(); err != nil {
		return
	}
}
