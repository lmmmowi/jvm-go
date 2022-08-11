package classfile

import "encoding/json"

type LocalVariableTableAttribute struct {
	Table []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	StartPc    uint16
	Length     uint16
	Name       string
	Descriptor string
	Index      uint16
}

func readLocalVariableTableAttribute(reader *ClassReader) LocalVariableTableAttribute {
	return LocalVariableTableAttribute{
		Table: reader.readTable(readLocalVariableTableEntry).([]LocalVariableTableEntry),
	}
}

func readLocalVariableTableEntry(reader *ClassReader) LocalVariableTableEntry {
	return LocalVariableTableEntry{
		StartPc:    reader.ReadUint16(),
		Length:     reader.ReadUint16(),
		Name:       reader.classFile.getConstant(reader.ReadUint16()),
		Descriptor: reader.classFile.getConstant(reader.ReadUint16()),
		Index:      reader.ReadUint16(),
	}
}

func (attr LocalVariableTableAttribute) print(placeHolder int) {
	for i, entry := range attr.Table {
		jsonStr, _ := json.Marshal(entry)
		_println(placeHolder, "-[%d]: %v", i, string(jsonStr))
	}
}
