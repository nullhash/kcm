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
