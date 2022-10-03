package main

import (
	"docker/src/utils_"
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		PrintHelpManual()
		return
	}

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
	cmd.Wait()
}

func Init() {
	imageDirPath := "/var/lib/docker/images/base"
	rootDirPath := "/var/lib/docker/containers/rootfs"
	if _, err := os.Stat(rootDirPath); os.IsNotExist(err) {
		utils_.ErrCheck(CopyFileOrDirectory(imageDirPath, rootDirPath))
	}
	time.Sleep(time.Second)

	// 切换主机名
	utils_.ErrCheck(syscall.Sethostname([]byte("Container")))

	// 先改变文件系统的根, 无法访问其他位置
	utils_.ErrCheck(syscall.Chroot(rootDirPath))
	utils_.ErrCheck(syscall.Chdir("/"))

	// 在第二/三层中 挂载 内存中的 虚拟文件系统 proc
	utils_.ErrCheck(syscall.Mount("proc", "/proc", "proc", 0, ""))

	path, err := exec.LookPath(os.Args[2])
	utils_.ErrCheck(err)
	fmt.Println(path)

	// 开辟第三层,运行 bash 命令, 附带所有参数 与 第二层的所有环境变量(mount)
	err = syscall.Exec(path, os.Args[3:], os.Environ()) // /bin/bash args... env...
	utils_.ErrCheck(err)

	// 取消 mount，退出取消 'mdocker init bash[1]' 这条线程的显示
	syscall.Unmount("/proc", 0)
}

func CopyFileOrDirectory(src string, dst string) error {
	fmt.Printf("Copying %s => %s\n", src, dst)
	cmd := exec.Command("cp", "-r", src, dst)
	utils_.AttachCmdToStd(cmd)
	return cmd.Run()
}
