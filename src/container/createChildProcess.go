package container

import (
	"docker/src/utils_"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

const (
	MAX_CONTAINER_ID     = 32
	IMAGE_DIR_PATH       = "/var/lib/docker/images/base"
	ROOT_DIR_PATH_PREFIX = "/var/lib/docker/containers/"
)

func CreateChildProcess(args []string) error {
	containerId := utils_.GenerateContainerId(MAX_CONTAINER_ID)
	imageDirPath := IMAGE_DIR_PATH
	rootDirPath := ROOT_DIR_PATH_PREFIX + containerId

	// Copy image
	if _, err := os.Stat(rootDirPath); os.IsNotExist(err) {
		if err := utils_.CopyFileOrDirectory(imageDirPath, rootDirPath); err != nil {
			return err
		}
	}
	// 切换主机名
	if err := syscall.Sethostname([]byte(containerId)); err != nil {
		return err
	}
	// 先改变文件系统的根, 无法访问其他位置
	if err := syscall.Chroot(rootDirPath); err != nil {
		return err
	}
	if err := syscall.Chdir("/"); err != nil {
		return err
	}
	// 在第二层中 挂载 内存中的 虚拟文件系统 proc
	if err := syscall.Mount("proc", "/proc", "proc", 0, ""); err != nil {
		return err
	}

	// 在$PATH中查找命令
	path, err := exec.LookPath(args[0])
	utils_.Err(err)
	fmt.Println(path)

	// 开辟第三层,运行 bash 命令, 附带所有参数 与 第二层的所有环境变量(mount)
	if err := syscall.Exec(path, args[0:], os.Environ()); err != nil { // /bin/bash args... env..
		return err
	}

	return nil
}
