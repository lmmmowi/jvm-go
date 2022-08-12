package instruction

type BytecodeReader struct {
	code []byte
	pc   int
}

func (reader *BytecodeReader) Reset(code []byte, pc int) {
	reader.code = code
	reader.pc = pc
}

func (reader *BytecodeReader) ReadUint8() uint8 {
	val := reader.code[reader.pc]
	reader.pc++
	return val
}

func (reader *BytecodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *BytecodeReader) ReadUint16() uint16 {
	high := uint16(reader.ReadInt8())
	low := uint16(reader.ReadInt8())
	return (high << 8) | low
}

func (reader *BytecodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}
