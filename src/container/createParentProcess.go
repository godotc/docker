package container

import (
	"docker/src/alert"
	"docker/src/utils_"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func CreateParentProcess(containerName string, is_interactive, is_tty bool, args []string) *exec.Cmd {
	// 拼接原命令再次运行 docker 进入 Init 分支 (创建子进程)
	// "/proc/self/exe" 等同于 docker 命令 本身 即: docker child args...
	args = append([]string{containerName}, args[0:]...)
	logFilePath := filepath.Join(ROOT_FOLDER_PATH_PREFIX, containerName, LOG_FILENAME)
	cmd := exec.Command("/proc/self/exe", "child", strings.Join(args, " "))

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

	imageFolderPath := IMAGE_FOLDER_PATH
	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFIX, containerName, ROOTFS_NAME)

	// Copy image
	if _, err := os.Stat(rootFolderPath); os.IsNotExist(err) {
		utils_.Err(utils_.CopyFileOrDirectory(imageFolderPath, rootFolderPath), "013")
	}

	// 根据 flag 判断是否附加到 stdout/in/err, 或则输出到日志文件
	if is_tty {
		utils_.AttachCmdToStd(cmd, is_interactive, is_tty)
	} else {
		// detach mode
		logFile, err := os.Create(logFilePath)
		utils_.Err(err, "014")
		cmd.Stdout = logFile
		alert.Debug(containerName)
	}

	return cmd
}
