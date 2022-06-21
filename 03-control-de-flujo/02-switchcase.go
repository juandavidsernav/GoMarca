package main

import (
	"fmt"
	"runtime"
)

func main() {
	arch := runtime.GOARCH

	fmt.Println("La arquitectura de su computador es:")

	switch arch {
	case "386":
		fmt.Println("x86 de 32 bits")
	case "amd64":
		fmt.Println("x86 de 64 bits")
	default:
		fmt.Println(arch)
	}

	fmt.Println("El sistema operativo de su computador es:")

	os := runtime.GOOS

	switch os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("GNU/Linux")
	case "hurd":
		fmt.Println("GNU/Hurd")
	default:
		fmt.Println(os)
	}

}
