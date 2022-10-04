package utils_

import (
	"docker/src/alert"
	"os"
	"os/exec"
)

func Err(err error, code string) {
	if err != nil {
		alert.Show(err, code)
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
