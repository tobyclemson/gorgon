package command

import (
	"bytes"
	"strings"
)

type commandOutput struct {
	Header string
	Body   []string
	Raw    []string
}

func OutputFrom(buffer *bytes.Buffer) *commandOutput {
	lines := strings.Split(buffer.String(), "\n")
	var header string
	if len(lines) > 0 {
		header = lines[0]
	}
	var body []string
	if len(lines) > 2 {
		body = lines[2 : len(lines)-2]
	}
	return &commandOutput{
		Header: header,
		Body:   body,
		Raw:    lines,
	}
}
