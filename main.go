package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("running %v as %d \n", os.Args[2:], os.Getpid())
	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
	}
	cmd.Run()

}
func child() {
	fmt.Printf("running %v as %d \n", os.Args[2:], os.Getpid())
	syscall.Sethostname([]byte("container"))
	syscall.Mount("proc", "proc", "proc", 0, "")
	syscall.Mount()
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	//	cmd.SysProcAttr = &syscall.SysProcAttr{
	//		Cloneflags: syscall.CLONE_NEWUTS,
	//	}
	cmd.Run()

}
