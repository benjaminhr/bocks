// +build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// go run main.go run cmd/args
func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("whaat??")
	}
}

func run() {
	fmt.Printf("%v\n", os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}

	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func child() {
	fmt.Printf("%v\n", os.Getpid())
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := syscall.Chroot("/home/rootfs")
	checkError(err)
	err = os.Chdir("/")
	checkError(err)
	err = syscall.Mount("proc", "proc", "proc", 0, "")
	checkError(err)
	err = cmd.Run()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
