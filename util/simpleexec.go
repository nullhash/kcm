package util

import (
	"errors"
	"io/ioutil"
	"os"
	"os/exec"
)

func ExecuteSingleCommand(name string, arg ...string) (string, error) {
	rescueStdout := os.Stdout
	rescueStderr := os.Stderr
	or, ow, err := os.Pipe()
	if err != nil {
		return "", err
	}
	er, ew, err := os.Pipe()
	if err != nil {
		return "", err
	}
	os.Stdout = ow
	os.Stderr = ew
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err = cmd.Start(); err != nil {
		return "", err
	}
	cmd.Wait()
	ow.Close()
	ew.Close()
	stdout, _ := ioutil.ReadAll(or)
	stderr, _ := ioutil.ReadAll(er)
	os.Stdout = rescueStdout
	os.Stderr = rescueStderr
	if len(string(stderr)) != 0 {
		return "", errors.New(string(stderr))
	}
	return string(stdout), nil
}
