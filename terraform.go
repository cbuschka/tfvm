package tfvm

import (
	"os"
	"os/exec"
	"syscall"
)

type Terraform struct {
	version string
	path string
}

func (this *Terraform) Run(args ...string) (int, error) {
	cmd := exec.Command(this.path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Env = os.Environ()
	if err := cmd.Start(); err != nil {
		return -1, err
	}

	if err := cmd.Wait(); err != nil {
		// https://stackoverflow.com/questions/10385551/get-exit-code-go
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
		} else {
			return -1, err
		}
	}

	return 0, nil
}
