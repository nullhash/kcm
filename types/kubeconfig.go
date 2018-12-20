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

package types

import (
	"time"
)

type KubeConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Cluster string `yaml:"cluster"`
			User    string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			AuthProvider struct {
				Config struct {
					AccessToken string    `yaml:"access-token"`
					CmdArgs     string    `yaml:"cmd-args"`
					CmdPath     string    `yaml:"cmd-path"`
					Expiry      time.Time `yaml:"expiry"`
					ExpiryKey   string    `yaml:"expiry-key"`
					TokenKey    string    `yaml:"token-key"`
				} `yaml:"config"`
				Name string `yaml:"name"`
			} `yaml:"auth-provider"`
		} `yaml:"user"`
	} `yaml:"users"`
}

type KcmConfig struct {
	Contexts    []KcmContext `yaml:"contexts"`     // Contexts ...
	ConfigFiles []string     `yaml:"config-files"` // ConfigFiles ...
}

type KcmContext struct {
	Name           string `yaml:"name"`
	ClusterName    string `yaml:"cluster-name"`
	ContextName    string `yaml:"context-name"`
	ConfigFilePath string `yaml:"config-file-path"`
	KcmContextPath string `yaml:"kcm-context-path"`
}
