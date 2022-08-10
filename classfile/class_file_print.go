package classfile

import "fmt"

func (cf *ClassFile) Print() {
	printMagicNumber(cf)
	printVersion(cf)
	printClass(cf)
	printInterfaces(cf)
	printFields(cf)
	printMethods(cf)

	// 常量池
	//for i := 1; i < len(cf.constantPool); i++ {
	//	fmt.Printf("constant pool[%d]: %T => %v\n", i, cf.constantPool[i], stringify(cf.constantPool, i))
	//}
}

func printMagicNumber(cf *ClassFile) {
	fmt.Printf("magic number: %X\n", cf.magic)
}

func printVersion(cf *ClassFile) {
	fmt.Printf("minor version: %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)
}

func printClass(cf *ClassFile) {
	fmt.Printf("this class: %v\n", cf.constantPool.stringify(cf.thisClass))
	fmt.Printf("super class: %v\n", cf.constantPool.stringify(cf.superClass))
}

func printInterfaces(cf *ClassFile) {
	for i := 0; i < len(cf.interfaces); i++ {
		fmt.Printf("interface[%d]: %v\n", i, cf.constantPool.stringify(cf.interfaces[i]))
	}
}

func printFields(cf *ClassFile) {
	for i, field := range cf.fields {
		name := cf.constantPool.stringify(field.NameIndex)
		descriptor := cf.constantPool.stringify(field.DescriptorIndex)
		fmt.Printf("field[%d]: %b (%v) %v\n", i, field.AccessFlags, name, descriptor)
		printAttributes(field.Attributes)
	}
}

func printMethods(cf *ClassFile) {
	for i, method := range cf.methods {
		name := cf.constantPool.stringify(method.NameIndex)
		descriptor := cf.constantPool.stringify(method.DescriptorIndex)
		fmt.Printf("method[%d]: %b %v %v\n", i, method.AccessFlags, name, descriptor)
		printAttributes(method.Attributes)
	}
}

func printAttributes(attributes []AttributeInfo) {
	for _, attr := range attributes {
		fmt.Printf("\t-%v:\n", attr.getName())
		attr.print()
	}
}
