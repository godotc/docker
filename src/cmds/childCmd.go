package cmds

import (
	"docker/src/container"
	"docker/src/utils_"

	"github.com/spf13/cobra"
)

func InitChildCmd() *cobra.Command {
	childCmd := &cobra.Command{
		Use: "child",
		Run: func(self *cobra.Command, args []string) {
			utils_.Err(container.CreateChildProcess(args))
		},
	}
	return childCmd
}
