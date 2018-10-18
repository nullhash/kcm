package util

import (
	"os"

	homedir "github.com/mitchellh/go-homedir"
)

// GetHomeDir gives the current user home directory path
func GetHomeDir() (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return home, err
}

// CheckFileOrDirectoryExists to check directory or file is present on the given path.
func CheckFileOrDirectoryExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
