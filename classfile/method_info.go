package classfile

type MethodInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}

func readMethods(reader *ClassReader) []MethodInfo {
	count := int(reader.ReadUint16())
	methods := make([]MethodInfo, count)
	for i := 0; i < count; i++ {
		methods[i] = readMethod(reader)
	}
	return methods
}

func readMethod(reader *ClassReader) MethodInfo {
	return MethodInfo{
		AccessFlags:     reader.ReadUint16(),
		NameIndex:       reader.ReadUint16(),
		DescriptorIndex: reader.ReadUint16(),
		Attributes:      readAttributes(reader),
	}
}
