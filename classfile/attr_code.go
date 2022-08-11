package classfile

import "reflect"

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
		ExceptionTable: reader.readTable(readExceptionTableEntry).([]ExceptionTableEntry),
		AttributeTable: readAttributes(reader),
	}
}

func readExceptionTableEntry(reader *ClassReader) ExceptionTableEntry {
	return ExceptionTableEntry{
		StartPc:   reader.ReadUint16(),
		EndPc:     reader.ReadUint16(),
		HandlerPc: reader.ReadUint16(),
		CatchType: reader.ReadUint16(),
	}
}

func (attr CodeAttribute) print(placeHolder int) {
	_println(placeHolder, "-MaxStack: %d", attr.MaxStack)
	_println(placeHolder, "-MaxLocals: %d", attr.MaxLocals)
	_println(placeHolder, "-Codes: %d", attr.Code)

	_println(placeHolder, "-ExceptionTable: %v", attr.ExceptionTable)

	_println(placeHolder, "-AttributeTable:")
	for _, subAttr := range attr.AttributeTable {
		_println(placeHolder+1, "-%v:", reflect.TypeOf(subAttr).Name())
		subAttr.print(placeHolder + 2)
	}
}
