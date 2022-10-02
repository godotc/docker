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
		Run()
	case "init":
		Init()
	default:
		panic("Argument  error")
	}
}

func Run() {
	// 拼接原命令再次运行 docker 进入 Init 分支 (创建子进程)
	cmd := exec.Command(os.Args[0], "init", os.Args[2])

	// 设置进程程隔离 thread-isolation
	cmd.SysProcAttr = &syscall.SysProcAttr{
		/*
			Linux Namespace:
				CLONE_NEWUTS: UTS Namespace 隔离 nodename 和 dominname
				CLONE_NEWPID: PID Namespace 隔离进程ID
				CLONE_NEWS: Mount Namespace 隔离进程看到挂载点视图 (文件系统)
		*/
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil { // 进入 子进程
		panic(err)
	}
}

func Init() {

	// 切换主机名
	syscall.Sethostname([]byte("Container"))
	// 在第二/三层中 挂载 内存中的 虚拟文件系统 proc
	syscall.Mount("proc", "/proc", "proc", 0, "")

	// 改变文件系统的根, 无法访问其他位置
	syscall.Chroot("rootfs")
	syscall.Chdir("/")

	// 开辟第三层,运行 bash 命令, 附带所有参数 与 第二层的所有环境变量(mount)
	if err := syscall.Exec(os.Args[2], os.Args[3:], os.Environ()); // /bin/bash args... env...
	err != nil {
		panic(err)
	}

	// 取消 mount，退出取消 'mdocker init bash[1]' 这条线程的显示
	syscall.Unmount("/proc", 0)
}
