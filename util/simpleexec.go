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
