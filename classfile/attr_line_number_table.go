package classfile

import "encoding/json"

type LineNumberTableAttribute struct {
	Table []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	StartPc    uint16
	LineNumber uint16
}

func readLineNumberTableAttribute(reader *ClassReader) LineNumberTableAttribute {
	return LineNumberTableAttribute{
		Table: reader.readTable(readLineNumberTableEntry).([]LineNumberTableEntry),
	}
}

func readLineNumberTableEntry(reader *ClassReader) LineNumberTableEntry {
	return LineNumberTableEntry{
		StartPc:    reader.ReadUint16(),
		LineNumber: reader.ReadUint16(),
	}
}

func (attr LineNumberTableAttribute) print(placeHolder int) {
	for i, entry := range attr.Table {
		jsonStr, _ := json.Marshal(entry)
		_println(placeHolder, "-[%d]: %v", i, string(jsonStr))
	}
}
