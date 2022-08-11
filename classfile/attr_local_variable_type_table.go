package classfile

import "encoding/json"

type LocalVariableTypeTableAttribute struct {
	Table []LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	StartPc   uint16
	Length    uint16
	Name      string
	Signature string
	Index     uint16
}

func readLocalVariableTypeTableAttribute(reader *ClassReader) LocalVariableTypeTableAttribute {
	return LocalVariableTypeTableAttribute{
		Table: reader.readTable(readLocalVariableTypeTableEntry).([]LocalVariableTypeTableEntry),
	}
}

func readLocalVariableTypeTableEntry(reader *ClassReader) LocalVariableTypeTableEntry {
	return LocalVariableTypeTableEntry{
		StartPc:   reader.ReadUint16(),
		Length:    reader.ReadUint16(),
		Name:      reader.classFile.getConstant(reader.ReadUint16()),
		Signature: reader.classFile.getConstant(reader.ReadUint16()),
		Index:     reader.ReadUint16(),
	}
}

func (attr LocalVariableTypeTableAttribute) print(placeHolder int) {
	for i, entry := range attr.Table {
		jsonStr, _ := json.Marshal(entry)
		_println(placeHolder, "-[%d]: %v", i, string(jsonStr))
	}
}
