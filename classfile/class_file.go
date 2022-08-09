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
	fields       []FiledInfo
	methods      []MethodInfo
	attributes   []AttributeInfo
}

type FiledInfo struct {
}

type MethodInfo struct {
}

type AttributeInfo struct {
}

func ParseClassFile(data []byte) *ClassFile {
	reader := newClassReader(data)
	cf := &ClassFile{}
	cf.readMagic(&reader)
	cf.readMinorVersion(&reader)
	cf.readMajorVersion(&reader)
	cf.readConstantPool(&reader)
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

func (cf *ClassFile) Print() {
	fmt.Printf("magic number: %X\n", cf.magic)
	fmt.Printf("minor version %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)

	for i := 1; i < len(cf.constantPool); i++ {
		fmt.Printf("constant pool[%d]: %T => %v\n", i, cf.constantPool[i], constantInfoStringify(cf.constantPool, i))
	}

	//hex_string_data := hex.EncodeToString(cf.magic)
}
