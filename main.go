package main

import (
	"fmt"
	"os"
	"os/exec"
)

/* docker
docker run image <cmd> <param>

*/

/*
go equivalent
go run main.go

*/

func main() {
	switch os.Args[1] {
	case "run":
		run()
		//	case "child":
		// child()
	default:
		panic("bad command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	// cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd := exec.Command(os.Args[2], os.Args[3:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// create namespace
	// cmd.SysProcAttr = &syscall.SysProcAttr{
	// 	Cloneflags: syscall.CLONE_NEWUTS,
	// }
	cmd.Run()

}

// func child() {
// 	fmt.Printf("Running %v\n", os.Args[2:])
// 	syscall.Sethostname([]byte("container"))
// 	cmd := exec.Command(os.Args[2], os.Args[3:]...)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr

// 	cmd.Run()
// }

func must(err error) {
	if err != nil {
		panic(err)
	}
}
