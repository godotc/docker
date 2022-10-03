package main

import (
	"docker/src/utils_"
	"fmt"
	"math/rand"
	"os/exec"
	"time"
)

func PrintHelpManual() {
	context := "argument error"
	fmt.Printf("%v\n", context)
}

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
	fmt.Printf("Copying %s => %s\n", src, dst)
	cmd := exec.Command("cp", "-r", src, dst)
	utils_.AttachCmdToStd(cmd)
	return cmd.Run()
}
