package alert

import (
	"fmt"
)

func Show(err error, num string) {
	switch num {
	case "001":
		ErrorShow("Can not get is_tty value.")
	case "002":
		ErrorShow("Can not get is_interactive value.")
	case "003":
		ErrorShow("Can not get is_detach value.")
	case "004":
		ErrorShow("can not use '-it' and '-d' at the same time.")
	case "005":
		ErrorShow("Can not get container name.")
	case "006":
		ErrorShow("can not start the cmd .")
	case "007":
		ErrorShow("Failed on syscall.Sethostname.")
	case "008":
		ErrorShow("Failed on syscall.chroot.")
	case "009":
		ErrorShow("Failed on syscall.Chdir.")
	case "010":
		ErrorShow("Failed on syscall.Mount 'proc' folder.")
	case "011":
		ErrorShow("Failed on exec.LookPath. ")
	case "012":
		ErrorShow("Failed on syscall.Exec cmd. ")
	case "013":
		ErrorShow("Failed on copy base image.")
	case "014":
		ErrorShow("Can not create container.log file.")
	case "015":
		ErrorShow("Can not open container.log file.")
	case "016":
		ErrorShow("can not write output to container.log file.")
	default:
	}
	panic(err)
}

func ErrorShow(msg string) {
	fmt.Printf(" [ERROR] %s\n", msg)
}

func Debug(msg string) {
	fmt.Println(msg)
}
