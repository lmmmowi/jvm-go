package classfile

import (
	"fmt"
	"reflect"
)

func (cf *ClassFile) Print() {
	printMagicNumber(cf)
	printVersion(cf)
	printClass(cf)
	printInterfaces(cf)
	printFields(cf)
	printMethods(cf)

	// 常量池
	//for i := 1; i < len(cf.constantPool); i++ {
	//	fmt.Printf("constant pool[%d]: %T => %v\n", i, cf.constantPool[i], constantStringify(cf.constantPool, i))
	//}
}

func _println(placeHolder int, format string, a ...interface{}) {
	for i := 0; i < placeHolder; i++ {
		fmt.Print("\t")
	}
	fmt.Printf(format+"\n", a...)
}

func printMagicNumber(cf *ClassFile) {
	_println(0, "magic number: %X", cf.magic)
}

func printVersion(cf *ClassFile) {
	_println(0, "minor version: %d", cf.minorVersion)
	_println(0, "major version: %d", cf.majorVersion)
}

func printClass(cf *ClassFile) {
	_println(0, "this class: %v", cf.getConstant(cf.thisClass))
	_println(0, "super class: %v", cf.getConstant(cf.superClass))
}

func printInterfaces(cf *ClassFile) {
	for i := 0; i < len(cf.interfaces); i++ {
		_println(0, "interface[%d]: %v", i, cf.getConstant(cf.interfaces[i]))
	}
}

func printFields(cf *ClassFile) {
	for i, field := range cf.fields {
		name := cf.getConstant(field.NameIndex)
		descriptor := cf.getConstant(field.DescriptorIndex)
		_println(0, "field[%d]: %b (%v) %v", i, field.AccessFlags, name, descriptor)
		printAttributes(field.Attributes)
	}
}

func printMethods(cf *ClassFile) {
	for i, method := range cf.methods {
		name := cf.getConstant(method.NameIndex)
		descriptor := cf.getConstant(method.DescriptorIndex)
		_println(0, "method[%d]: %b %v %v", i, method.AccessFlags, name, descriptor)
		printAttributes(method.Attributes)
	}
}

func printAttributes(attributes []AttributeInfo) {
	for _, attr := range attributes {
		_println(1, "-%v:", reflect.TypeOf(attr).Name())
		attr.print(2)
	}
}
