package container

import (
	"docker/src/utils_"
	"os"
	"os/exec"
	"syscall"
)

func CreateParentProcess(is_interactive, is_tty bool, args []string) *exec.Cmd {
	// 拼接原命令再次运行 docker 进入 Init 分支 (创建子进程)
	// "/proc/self/exe" 等同于 docker 命令 本身 即: docker child args...
	cmd := exec.Command("/proc/self/exe", "child", args[0])

	// 设置进程程隔离 thread-isolation, 用户id隔离
	cmd.SysProcAttr = &syscall.SysProcAttr{
		/*
			Linux Namespace:
				CLONE_NEWUTS: UTS Namespace 隔离 nodename 和 dominname
				CLONE_NEWPID: PID Namespace 隔离进程ID
				CLONE_NEWS: Mount Namespace 隔离进程看到挂载点视图 (文件系统)
				CLONE_NEWIPC: IPC Namespace 隔离 System V IPC 和 POSIX Message Queues
				CLONE_NEWUSER: User Namespace 隔离用户组ID 防止拥有(root)权限
		*/
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWUSER,
		// map 外界的id(非root) 到 容器内的 uer id 和 group id
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getuid(),
				Size:        1, // default
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 0,
				HostID:      os.Getgid(),
				Size:        1,
			},
		},
	}

	// 根据 flag 判断是否附加到 stdout/in/err
	utils_.AttachCmdToStd(cmd, is_interactive, is_tty)

	return cmd
}
