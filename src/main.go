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
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID, // 复制核心 | PID
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}

func child() {
	cmd := exec.Command(os.Args[2])

	// 实际在这里执行 docker 命令的参数
	syscall.Sethostname([]byte("Container"))

	// 在本文件系统中
	// 不允许运行其他程序 | 不允许 set-user-id/set-group-id | 从 linux 2.4 mount 默认参数
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	// 挂载根目录
	syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	// 设置完成后运行bash
	if err := cmd.Run(); err != nil {
		panic(err)
	}
	syscall.Unmount("/proc", 0)
}
