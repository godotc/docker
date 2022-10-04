package utils_

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GenerateContainerId(n uint) string {
	rand.Seed(time.Now().UnixNano())
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"
	length := len(letters)

	str := make([]byte, n)
	for i := range str {
		str[i] = letters[rand.Intn(length)]
	}
	return string(str)
}

func CopyFileOrDirectory(src string, dst string) error {
	info, err := os.Stat(src)
	if err != nil {
		return err
	}

	if info.IsDir() {
		if err := os.MkdirAll(dst, 0777); err != nil {
			return err
		}
		// 获取目录中所有文件
		if list, err := ioutil.ReadDir(src); err == nil {
			for _, item := range list {
				// 递归 复制每一个文件/文件夹
				if err = CopyFileOrDirectory(filepath.Join(src, item.Name()), filepath.Join(dst, item.Name())); err != nil {
					return err
				}
			}
		} else {
			return err
		}
	} else {
		content, err := ioutil.ReadFile(src)
		if err != nil {
			return err
		}
		// copy
		if err := ioutil.WriteFile(dst, content, 0777); err != nil {
			return err
		}
	}

	return nil
}
