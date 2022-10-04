package main

import (
	"docker/src/cmds"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use: "docer [Command]",
	}

	rootCmd.AddCommand(cmds.InitRunCmd())
	rootCmd.AddCommand(cmds.InitChildCmd())
	rootCmd.AddCommand(cmds.InitLogsCmd())

	rootCmd.Execute()
}
