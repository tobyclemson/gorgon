package command

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/tobyclemson/gorgon/test/support/fs"
	"os/exec"
	"path/filepath"
	"testing"
)

func buildStreamCapturingCommand(
	binary string,
	args ...string,
) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	cmd := exec.Command(binary, args...)
	var outputBuffer, errorBuffer bytes.Buffer
	cmd.Stdout = &outputBuffer
	cmd.Stderr = &errorBuffer

	return cmd, &outputBuffer, &errorBuffer
}

func Run(
	t *testing.T,
	name string,
	args ...string,
) (*exec.Cmd, *bytes.Buffer, *bytes.Buffer) {
	workingDirectory := fs.GetWorkingDirectory(t)
	binary := filepath.Join(workingDirectory, "..", "..", name)

	cmd, outputBuffer, errorBuffer :=
		buildStreamCapturingCommand(binary, args...)
	logCommand(t, "Executing command", cmd)

	err := cmd.Run()
	logLines(t, "Standard Output", outputBuffer)
	logLines(t, "Standard Error", errorBuffer)
	assert.Nil(t, err)

	return cmd, outputBuffer, errorBuffer
}
