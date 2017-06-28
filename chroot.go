package main

import (
	"os"
	"os/exec"
	"syscall"
)

func must(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	must(syscall.Chroot(os.Args[1]))
	must(syscall.Chdir("/"))
	cmd := exec.Command("/bin/sh")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWPID}
	must(cmd.Run())
}
