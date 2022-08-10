package classfile

import "fmt"

type CodeAttribute struct {
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableEntry
	AttributeTable []AttributeInfo
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func readCodeAttribute(reader *ClassReader) CodeAttribute {
	return CodeAttribute{
		MaxStack:       reader.ReadUint16(),
		MaxLocals:      reader.ReadUint16(),
		Code:           reader.ReadBytes(int(reader.ReadUint32())),
		ExceptionTable: readExceptionTable(reader),
		AttributeTable: readAttributes(reader),
	}
}

func readExceptionTable(reader *ClassReader) []ExceptionTableEntry {
	count := int(reader.ReadUint16())
	exceptionTableEntries := make([]ExceptionTableEntry, count)
	for i := 0; i < count; i++ {
		exceptionTableEntries[i] = ExceptionTableEntry{
			StartPc:   reader.ReadUint16(),
			EndPc:     reader.ReadUint16(),
			HandlerPc: reader.ReadUint16(),
			CatchType: reader.ReadUint16(),
		}
	}
	return exceptionTableEntries
}

func (attr CodeAttribute) getName() string {
	return Code
}

func (attr CodeAttribute) print() {
	fmt.Printf("\t\t-MaxStack: %d\n", attr.MaxStack)
	fmt.Printf("\t\t-MaxLocals: %d\n", attr.MaxLocals)
	fmt.Printf("\t\t-Codes: %d\n", attr.Code)
	fmt.Printf("\t\t-ExceptionTable: %v\n", attr.ExceptionTable)
	fmt.Printf("\t\t-AttributeTable: %v\n", attr.AttributeTable)
}
