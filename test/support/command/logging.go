package command

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

func logCommand(t *testing.T, title string, cmd *exec.Cmd) {
	t.Logf("%v: %v", title, cmd.String())
	t.Log()
}

func logLines(t *testing.T, title string, linesBuffer *bytes.Buffer) {
	t.Logf("%v:", title)
	t.Log()
	lines := strings.Split(linesBuffer.String(), "\n")
	for _, line := range lines {
		t.Log(line)
	}
}
