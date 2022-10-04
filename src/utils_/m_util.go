package utils_

import (
	"os"
	"os/exec"
)

func Err(err error) {
	if err != nil {
		panic(err)
	}
}

// If tty, attached to stderr and stdin. If interactive, attached to stdout
func AttachCmdToStd(cmd *exec.Cmd, is_interactive, is_tty bool) {
	if is_tty {
		if is_interactive {
			cmd.Stdout = os.Stdout
		}
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
	}
}
