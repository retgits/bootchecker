package cmd

import (
	"fmt"
	"os/exec"
)

func runner(command string) string {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = "."
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Sprintf("Error while executing command [%s]: %s", command, err.Error())
	}
	return string(output)
}
