package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	fmt.Printf("Process => %v [%d]\n", os.Args, os.Getpid())
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("Argument  error")
	}
}

func run() {

	// 在主进程中设置线程隔离后，替换参数进入子进程分支
	cmd := exec.Command(os.Args[0], append([]string{"child"}, os.Args[2])...)

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID, // 进程隔离,隔离线程ID
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	// 先进入子进程，隔离后，再进行更改hostname等操作
	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	// 实际在这里执行 docker 命令的参数
	cmd := exec.Command(os.Args[2])
	syscall.Sethostname([]byte("Container"))

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	cmd.Run()
}
