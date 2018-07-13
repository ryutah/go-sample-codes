package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type output struct {
	stdout *bytes.Buffer
	stderr *bytes.Buffer
}

func getCommandOutput(cmd string, args ...string) (*output, error) {
	output := output{
		stdout: new(bytes.Buffer),
		stderr: new(bytes.Buffer),
	}

	c := exec.Command(cmd, args...)
	c.Stdout = output.stdout
	c.Stderr = output.stderr
	if err := c.Run(); err != nil {
		return nil, fmt.Errorf(
			"failed to exec command [%v %v]: details: %v",
			cmd, strings.Join(args, " "),
			strings.TrimSpace(output.stderr.String()),
		)
	}

	return &output, nil
}
