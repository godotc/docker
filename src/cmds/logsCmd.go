package cmds

import (
	"fmt"

	"github.com/spf13/cobra"
)

func InitLogsCmd() *cobra.Command {
	logsCmd := &cobra.Command{
		Use:   "logs",
		Short: "Fetch the logs of a container",
		Run: func(self *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
	return logsCmd
}
