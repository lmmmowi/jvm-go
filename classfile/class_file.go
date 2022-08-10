package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool []ConstantInfo
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []FieldInfo
	methods      []MethodInfo
	attributes   []AttributeInfo
}

func ParseClassFile(data []byte) *ClassFile {
	reader := newClassReader(data)
	cf := &ClassFile{}
	cf.readMagic(&reader)
	cf.readMinorVersion(&reader)
	cf.readMajorVersion(&reader)
	cf.readConstantPool(&reader)
	cf.readAccessFlags(&reader)
	cf.readThisClass(&reader)
	cf.readSuperClass(&reader)
	cf.readInterfaces(&reader)
	cf.readFields(&reader)
	cf.readMethods(&reader)
	cf.readAttributes(&reader)
	return cf
}

func (cf *ClassFile) readMagic(reader *ClassReader) {
	cf.magic = reader.ReadUint32()
}

func (cf *ClassFile) readMinorVersion(reader *ClassReader) {
	cf.minorVersion = reader.ReadUint16()
}

func (cf *ClassFile) readMajorVersion(reader *ClassReader) {
	cf.majorVersion = reader.ReadUint16()
}

func (cf *ClassFile) readConstantPool(reader *ClassReader) {
	cf.constantPool = readConstantPool(reader)
}

func (cf *ClassFile) readAccessFlags(reader *ClassReader) {
	cf.accessFlags = reader.ReadUint16()
}

func (cf *ClassFile) readThisClass(reader *ClassReader) {
	cf.thisClass = reader.ReadUint16()
}

func (cf *ClassFile) readSuperClass(reader *ClassReader) {
	cf.superClass = reader.ReadUint16()
}

func (cf *ClassFile) readInterfaces(reader *ClassReader) {
	count := int(reader.ReadUint16())
	interfaces := make([]uint16, count)
	for i := 0; i < count; i++ {
		interfaces[i] = reader.ReadUint16()
	}
	cf.interfaces = interfaces
}

func (cf *ClassFile) readFields(reader *ClassReader) {
	cf.fields = readFields(reader)
}

func (cf *ClassFile) readMethods(reader *ClassReader) {
	cf.methods = readMethods(reader)
}

func (cf *ClassFile) readAttributes(reader *ClassReader) {
	cf.attributes = readAttributes(reader)
}

func (cf *ClassFile) Print() {
	// 魔数
	fmt.Printf("magic number: %X\n", cf.magic)

	// 编译版本
	fmt.Printf("minor version %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)

	// 类名&接口
	fmt.Printf("this class: %v\n", constantInfoStringify(cf.constantPool, int(cf.thisClass)))
	fmt.Printf("super class: %v\n", constantInfoStringify(cf.constantPool, int(cf.superClass)))
	for i := 0; i < len(cf.interfaces); i++ {
		fmt.Printf("interface[%d]: %v\n", i, constantInfoStringify(cf.constantPool, int(cf.interfaces[i])))
	}

	// 字段
	for i := 0; i < len(cf.fields); i++ {
		field := cf.fields[i]
		name := constantInfoStringify(cf.constantPool, int(field.NameIndex))
		descriptor := constantInfoStringify(cf.constantPool, int(field.DescriptorIndex))
		fmt.Printf("field[%d]: %b (%v) %v\n", i, field.AccessFlags, name, descriptor)
	}

	// 方法
	for i := 0; i < len(cf.methods); i++ {
		method := cf.methods[i]
		name := constantInfoStringify(cf.constantPool, int(method.NameIndex))
		descriptor := constantInfoStringify(cf.constantPool, int(method.DescriptorIndex))
		fmt.Printf("method[%d]: %b %v %v\n", i, method.AccessFlags, name, descriptor)
	}

	// 常量池
	//for i := 1; i < len(cf.constantPool); i++ {
	//	fmt.Printf("constant pool[%d]: %T => %v\n", i, cf.constantPool[i], constantInfoStringify(cf.constantPool, i))
	//}
}
