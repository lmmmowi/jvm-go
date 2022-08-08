package main

import (
	"fmt"
	"github.com/lmmmowi/jvm-go/vm"
)

func main() {
	options := vm.ParseOptions()
	startJVM(options)
}

func startJVM(options *vm.Options) {
	if options.VersionFlag {
		fmt.Println("jvm-go v0.0.0")
	}

	fmt.Println("java home:", options.JavaHomePath)

	if options.MainClass != "" {
		fmt.Println("main class:", options.MainClass)

		for index, arg := range options.Args {
			fmt.Println("arg-", index, ":", arg)
		}
	}
}
