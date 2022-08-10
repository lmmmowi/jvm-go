package classfile

const (
	Code                      = "Code"
	RuntimeVisibleAnnotations = "RuntimeVisibleAnnotations"
	Signature                 = "Signature"
)

type AttributeInfo interface {
	getName() string
	print()
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
	name := reader.classFile.constantPool.stringify(nameIndex)
	attrLength := reader.ReadUint32()

	switch name {
	case Code:
		return readCodeAttribute(reader)
	case Signature:
		return readSignatureAttribute(reader)
	case RuntimeVisibleAnnotations:
		return readRuntimeVisibleAnnotationsAttribute(reader)
	default:
		return UnparsedAttribute{
			Name: name,
			Info: reader.ReadBytes(int(attrLength)),
		}
	}
}
