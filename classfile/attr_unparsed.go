package classfile

import "fmt"

type UnparsedAttribute struct {
	Name string
	Info []byte
}

func (attr UnparsedAttribute) getName() string {
	return attr.Name
}

func (attr UnparsedAttribute) print() {
	fmt.Printf("\t\t-Info: %v\n", attr.Info)
}
