package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Println(os.Args)
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("Argument  error")
	}
}

func run() {
	cmd := exec.Command(os.Args[2])
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS, // 进程隔离
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	cmd.Run()
}
