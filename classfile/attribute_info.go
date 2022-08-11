package classfile

const (
	ConstantValue                        = "ConstantValue"
	Code                                 = "Code"
	StackMapTable                        = "StackMapTable"
	Exceptions                           = "Exceptions"
	InnerClasses                         = "InnerClasses"
	EnclosingMethod                      = "EnclosingMethod"
	Synthetic                            = "Synthetic"
	Signature                            = "Signature"
	SourceFile                           = "SourceFile"
	SourceDebugExtension                 = "SourceDebugExtension"
	LineNumberTable                      = "LineNumberTable"
	LocalVariableTable                   = "LocalVariableTable"
	LocalVariableTypeTable               = "LocalVariableTypeTable"
	Deprecated                           = "Deprecated"
	RuntimeVisibleAnnotations            = "RuntimeVisibleAnnotations"
	RuntimeInvisibleAnnotations          = "RuntimeInvisibleAnnotations"
	RuntimeVisibleParameterAnnotations   = "RuntimeVisibleParameterAnnotations"
	RuntimeInvisibleParameterAnnotations = "RuntimeInvisibleParameterAnnotations"
	AnnotationDefault                    = "AnnotationDefault"
	BoostrapMethods                      = "BoostrapMethods"
)

type AttributeInfo interface {
	print(placeHolder int)
}

func readAttributes(reader *ClassReader) []AttributeInfo {
	return reader.readTable(readAttribute).([]AttributeInfo)
}

func readAttribute(reader *ClassReader) AttributeInfo {
	nameIndex := reader.ReadUint16()
	name := reader.classFile.getConstant(nameIndex)
	attrLength := reader.ReadUint32()

	switch name {
	case Code:
		return readCodeAttribute(reader)
	case Signature:
		return readSignatureAttribute(reader)
	case LineNumberTable:
		return readLineNumberTableAttribute(reader)
	case LocalVariableTable:
		return readLocalVariableTableAttribute(reader)
	case LocalVariableTypeTable:
		return readLocalVariableTypeTableAttribute(reader)
	case RuntimeVisibleAnnotations:
		return readRuntimeVisibleAnnotationsAttribute(reader)
	default:
		return UnparsedAttribute{
			Name: name,
			Info: reader.ReadBytes(int(attrLength)),
		}
	}
}
