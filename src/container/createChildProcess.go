package container

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

func CreateChildProcess(args []string) (error, string) {

	cmdArgs := strings.Split(args[0], " ")
	containerName := cmdArgs[0]
	rootFolderPath := filepath.Join(ROOT_FOLDER_PATH_PREFIX, containerName, ROOTFS_NAME)

	// 切换主机名
	if err := syscall.Sethostname([]byte(containerName)); err != nil {
		return err, "007"
	}
	// 先改变文件系统的根, 无法访问其他位置
	if err := syscall.Chroot(rootFolderPath); err != nil {
		return err, "008"
	}
	if err := syscall.Chdir("/"); err != nil {
		return err, "009"
	}
	// 在第二层中 挂载 内存中的 虚拟文件系统 proc
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return err, "010"
	}

	// 在$PATH中查找命令
	path, err := exec.LookPath(cmdArgs[1])
	if err != nil {
		return err, "011"
	}

	// 开辟第三层,运行 bash 命令, 附带所有参数 与 第二层的所有环境变量(mount)
	if err := syscall.Exec(path, cmdArgs[1:], os.Environ()); err != nil { // /bin/bash args... env..
		return err, "012"
	}

	return nil, ""
}
