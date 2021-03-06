package utils

import (
	"os/exec"
)

// GitInit ...
func GitInit(dir string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", "git init")
	cmd.Dir = dir
	return cmd.CombinedOutput()
}

// GetVersion ...
func GetVersion() (string, error) {
	cmd := exec.Command("bash", "-c", "git tag --sort=-version:refname | head -n 1")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}
