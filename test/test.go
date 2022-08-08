package main

import (
	"flag"
	"fmt"
	"github.com/lmmmowi/jvm-go/classfile"
	"log"
	"os"
)

func main() {
	flag.Parse()
	fileName := flag.Args()[0]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("read fail", err)
		return
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("read fail", err)
		return
	}

	fmt.Println("file size:", fi.Size())
	data := make([]byte, fi.Size())
	n, err := file.Read(data)
	fmt.Println("read", n, "bytes")

	classFile := classfile.ParseClassFile(data)
	classFile.Print()
}
