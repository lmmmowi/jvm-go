package classfile

type AttributeInfo struct {
	AttributeNameIndex uint16
	Info               []byte
}

func readAttributes(reader *ClassReader) []AttributeInfo {
	count := int(reader.ReadUint16())
	attributes := make([]AttributeInfo, count)
	for i := 0; i < count; i++ {
		attributes[i] = readAttribute(reader)
	}
	return attributes
}

func readAttribute(reader *ClassReader) AttributeInfo {
	nameIndex := reader.ReadUint16()
	length := reader.ReadUint32()
	info := reader.ReadBytes(int(length))

	return AttributeInfo{
		AttributeNameIndex: nameIndex,
		Info:               info,
	}
}
