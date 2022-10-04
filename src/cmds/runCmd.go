package cmds

import (
	"docker/src/container"
	"docker/src/utils_"
	"os"

	"github.com/spf13/cobra"
)

func InitRunCmd() *cobra.Command {
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a command in a new container",
		Run: func(self *cobra.Command, args []string) { // call back
			is_tty, err := self.Flags().GetBool("tty")
			utils_.Err(err)
			is_interactive, err := self.Flags().GetBool("interactive")
			utils_.Err(err)

			//fmt.Println("is_tty:", is_tty, "  is_interactive:", is_interactive)

			cmd := container.CreateParentProcess(is_interactive, is_tty, args)
			utils_.Err(cmd.Start())

			cmd.Wait()
			os.Exit(-1)
		},
	}
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a presudo-TTY")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")

	return runCmd
}
