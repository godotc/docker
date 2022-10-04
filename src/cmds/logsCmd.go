package cmds

import (
	"docker/src/container"

	"github.com/spf13/cobra"
)

func InitLogsCmd() *cobra.Command {
	logsCmd := &cobra.Command{
		Use:   "logs",
		Short: "Fetch the logs of a container",
		Run: func(self *cobra.Command, args []string) {
			container.DisplayContainerLog(args[0])
		},
	}
	return logsCmd
}
