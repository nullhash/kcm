package util

import (
	"errors"
	"io"
	"os"
	"os/exec"
	"strings"

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

// ExeCmd executes the given command
func ExeCmd(cmd string) (string, error) {
	// fmt.Println("command is ", cmd)
	// splitting head => g++ parts => rest of the command
	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:len(parts)]

	out, err := exec.Command(head, parts...).Output()
	if err != nil {
		return "", err
		// fmt.Printf("%s", err)
	}
	// fmt.Printf("%s", out)

	return string(out), nil
}

// DeleteDirectory is to delete directory in a given path.
func DeleteDirectory(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return err
	}
	return nil
}

// CopyConfigFile copies file content from sourc (src) to destination (dst) path.
func CopyConfigFile(src, dst string) (int64, error) {
	if !CheckFileOrDirectoryExists(src) {
		return 0, errors.New("Error: config file does not exist on the given path")
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

// GetClusterPath returns specific cluster directory path
func GetClusterPath(cluster string) (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/.kcm/" + cluster, nil
}

// GetClusterConfigPath returns specific cluster config file path
func GetClusterConfigPath(cluster string) (string, error) {
	home, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return home + "/.kcm/" + cluster + "/config", nil
}

// GetKubeconfigEnvValue returns the kubeconfig environment value
func GetKubeconfigEnvValue() (string, error) {
	kubeconfigValue := os.Getenv("KUBECONFIG")
	if kubeconfigValue == "" {
		return "", errors.New("KUBECONFIG environment variable not set")
	}
	return kubeconfigValue, nil
}
