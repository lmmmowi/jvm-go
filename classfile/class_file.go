package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []FiledInfo
	methods      []MethodInfo
	attributes   []AttributeInfo
}

type ConstantPool struct {
	count uint16
}

type FiledInfo struct {
}

type MethodInfo struct {
}

type AttributeInfo struct {
}

func (cf *ClassFile) readMagic(cr *ClassReader) {
	cf.magic = cr.ReadUint32()
}

func (cf *ClassFile) readMinorVersion(cr *ClassReader) {
	cf.minorVersion = cr.ReadUint16()
}

func (cf *ClassFile) readMajorVersion(cr *ClassReader) {
	cf.majorVersion = cr.ReadUint16()
}

func (cf *ClassFile) readConstantPool(cr *ClassReader) {
	constantPool := ConstantPool{}
	constantPool.count = cr.ReadUint16()
	cf.constantPool = constantPool
}

func (cf *ClassFile) Print() {
	fmt.Println(cf)
	fmt.Printf("magic number: %X\n", cf.magic)
	fmt.Printf("minor version %d\n", cf.minorVersion)
	fmt.Printf("major version: %d\n", cf.majorVersion)
	fmt.Printf("constant pool: %#v\n", cf.constantPool)
	//hex_string_data := hex.EncodeToString(cf.magic)
}
