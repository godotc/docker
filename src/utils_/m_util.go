package utils_

import (
	"os"
	"os/exec"
)

func ErrCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func AttachCmdToStd(cmd *exec.Cmd) {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
}
