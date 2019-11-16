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
	default:
		panic("whaat??")
	}
}

func run() {
	fmt.Printf("got %v\n", os.Args[2:])

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}

	err := cmd.Run()
	if err != nil {
		must(err)
	}
}

func must(err error) {
	panic(err)
}
