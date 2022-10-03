package main

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func LoadRootCobraCmds() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "docker [Command]",
	}

	// version
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show the Docker version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("---------")
		},
	}
	rootCmd.AddCommand(versionCmd)

	// ps
	psCmd := &cobra.Command{
		Use:   "ps",
		Short: "List containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(strings.Join(args, " "))
		},
	}
	psCmd.Flags().BoolP("all", "a", false, "Show all containers (default shows just running)")
	psCmd.Flags().BoolP("quiet", "q", false, "Only display container IDs")
	rootCmd.AddCommand(psCmd)

	// run
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run a command in a new container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("---------")
		},
	}
	rootCmd.AddCommand(runCmd)

	// exec
	execCmd := &cobra.Command{
		Use:   "exec",
		Short: "Run a command in a running container",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("---------")
		},
	}
	rootCmd.AddCommand(execCmd)

	// start
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start one or more stopped containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("---------")
		},
	}
	rootCmd.AddCommand(startCmd)

	// stop
	stopCmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop one or more running containers",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("---------")
		},
	}
	rootCmd.AddCommand(stopCmd)

	return rootCmd
}
