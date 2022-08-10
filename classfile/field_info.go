package classfile

const (
	FIELD_ACC_PUBLIC    = 0x0001
	FIELD_ACC_PRIVATE   = 0x0002
	FIELD_ACC_PROTECTED = 0x0004
	FIELD_ACC_STATIC    = 0x0008
	FIELD_ACC_FINAL     = 0x0010
	FIELD_ACC_VOLATILE  = 0x0040
	FIELD_ACC_TRANSIENT = 0x0080
	FIELD_ACC_SYNTHETIC = 0x1000
	FIELD_ACC_ENUM      = 0x4000
)

type FieldInfo struct {
	AccessFlags     uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Attributes      []AttributeInfo
}

func readFields(reader *ClassReader) []FieldInfo {
	count := int(reader.ReadUint16())
	fields := make([]FieldInfo, count)
	for i := 0; i < count; i++ {
		fields[i] = readField(reader)
	}
	return fields
}

func readField(reader *ClassReader) FieldInfo {
	return FieldInfo{
		AccessFlags:     reader.ReadUint16(),
		NameIndex:       reader.ReadUint16(),
		DescriptorIndex: reader.ReadUint16(),
		Attributes:      readAttributes(reader),
	}
}
