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
			utils_.Err(err, "001")

			is_interactive, err := self.Flags().GetBool("interactive")
			utils_.Err(err, "002")

			is_detach, err := self.Flags().GetBool("detach")
			utils_.Err(err, "003")

			containerName, err := self.Flags().GetString("name")
			utils_.Err(err, "004")
			if containerName == "" {
				containerName = utils_.GenerateContainerId(container.MAX_CONTAINER_ID)
			}

			cmd := container.CreateParentProcess(containerName, is_interactive, is_tty, args)
			utils_.Err(cmd.Start(), "006")

			// 如果 没有 -d 则 wait
			if !is_detach {
				cmd.Wait()
			}

			os.Exit(-1)
		},
	}
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a presudo-TTY")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().StringP("name", "n", "", "Assign a name to the container")

	return runCmd
}
