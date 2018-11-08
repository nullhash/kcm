package util

import (
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
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
