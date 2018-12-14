package cmd

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// not thread safe
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func TestVersion(t *testing.T) {
	result := captureStdout(func() { versionCommand.Run(nil, nil) })
	if !strings.HasPrefix(result, "kcm version v") {
		t.Error("Expected version string, got ", result)
	}
}
