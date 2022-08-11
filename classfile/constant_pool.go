package classfile

import (
	"fmt"
	"math"
)

const (
	CONSTANT_Utf8               = 1
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_Class              = 7
	CONSTANT_String             = 8
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_NameAndType        = 12
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface{}

type ConstantClassInfo struct {
	NameIndex uint16
}

type ConstantStringInfo struct {
	StringIndex uint16
}

type ConstantFieldRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantMethodRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantInterfaceRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func readConstantPool(reader *ClassReader) []ConstantInfo {
	count := int(reader.ReadUint16())
	constantPool := make([]ConstantInfo, count)
	for i := 1; i < count; i++ {
		constantPool[i] = readConstantInfo(reader)
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch constantPool[i].(type) {
		case int64, float64:
			i++
		}
	}
	return constantPool
}

func readConstantInfo(reader *ClassReader) ConstantInfo {
	tag := reader.ReadUint8()

	switch tag {
	case CONSTANT_Utf8:
		return readConstantUtf8Info(reader)
	case CONSTANT_Integer:
		return readConstantIntegerInfo(reader)
	case CONSTANT_Float:
		return readConstantFloatInfo(reader)
	case CONSTANT_Long:
		return readConstantLongInfo(reader)
	case CONSTANT_Double:
		return readConstantDoubleInfo(reader)
	case CONSTANT_Class:
		return readConstantClassInfo(reader)
	case CONSTANT_String:
		return readConstantStringInfo(reader)
	case CONSTANT_Fieldref:
		return readConstantFieldRefInfo(reader)
	case CONSTANT_Methodref:
		return readConstantMethodRefInfo(reader)
	case CONSTANT_InterfaceMethodref:
		return readConstantInterfaceRefInfo(reader)
	case CONSTANT_NameAndType:
		return readConstantNameAndTypeInfo(reader)
	case CONSTANT_MethodHandle:
		return readConstantMethodHandleInfo(reader)
	case CONSTANT_MethodType:
		return readConstantMethodTypeInfo(reader)
	case CONSTANT_InvokeDynamic:
		return readConstantInvokeDynamicInfo(reader)
	default:
		return nil
	}
}

/*
CONSTANT_Utf8_info {
    u1 tag;
	u2 length;
    u1 bytes[length];
}
*/
func readConstantUtf8Info(reader *ClassReader) []byte {
	length := int(reader.ReadUint16())
	return reader.ReadBytes(length)
}

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantIntegerInfo(reader *ClassReader) int32 {
	return int32(reader.ReadUint32())
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantFloatInfo(reader *ClassReader) float32 {
	return math.Float32frombits(reader.ReadUint32())
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantLongInfo(reader *ClassReader) int64 {
	return int64(reader.ReadUint64())
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantDoubleInfo(reader *ClassReader) float64 {
	return math.Float64frombits(reader.ReadUint64())
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
func readConstantClassInfo(reader *ClassReader) ConstantClassInfo {
	return ConstantClassInfo{NameIndex: reader.ReadUint16()}
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
func readConstantStringInfo(reader *ClassReader) ConstantStringInfo {
	return ConstantStringInfo{StringIndex: reader.ReadUint16()}
}

/*
CONSTANT_Fieldref_info {
    u1 tag;
	u2 class_index;
    u2 name_and_type_index;
}
*/
func readConstantFieldRefInfo(reader *ClassReader) ConstantFieldRefInfo {
	return ConstantFieldRefInfo{
		ClassIndex:       reader.ReadUint16(),
		NameAndTypeIndex: reader.ReadUint16(),
	}
}

/*
CONSTANT_Methodref_info {
    u1 tag;
	u2 class_index;
    u2 name_and_type_index;
}
*/
func readConstantMethodRefInfo(reader *ClassReader) ConstantMethodRefInfo {
	return ConstantMethodRefInfo{
		ClassIndex:       reader.ReadUint16(),
		NameAndTypeIndex: reader.ReadUint16(),
	}
}

/*
CONSTANT_Interfaceref_info {
    u1 tag;
	u2 class_index;
    u2 name_and_type_index;
}
*/
func readConstantInterfaceRefInfo(reader *ClassReader) ConstantInterfaceRefInfo {
	return ConstantInterfaceRefInfo{
		ClassIndex:       reader.ReadUint16(),
		NameAndTypeIndex: reader.ReadUint16(),
	}
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
	u2 name_index;
    u2 descriptor_index;
}
*/
func readConstantNameAndTypeInfo(reader *ClassReader) ConstantNameAndTypeInfo {
	return ConstantNameAndTypeInfo{
		NameIndex:       reader.ReadUint16(),
		DescriptorIndex: reader.ReadUint16(),
	}
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
	u1 reference_kind;
    u2 reference_index;
}
*/
func readConstantMethodHandleInfo(reader *ClassReader) ConstantMethodHandleInfo {
	return ConstantMethodHandleInfo{
		ReferenceKind:  reader.ReadUint8(),
		ReferenceIndex: reader.ReadUint16(),
	}
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
func readConstantMethodTypeInfo(reader *ClassReader) ConstantMethodTypeInfo {
	return ConstantMethodTypeInfo{
		DescriptorIndex: reader.ReadUint16(),
	}
}

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
	u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
func readConstantInvokeDynamicInfo(reader *ClassReader) ConstantInvokeDynamicInfo {
	return ConstantInvokeDynamicInfo{
		BootstrapMethodAttrIndex: reader.ReadUint16(),
		NameAndTypeIndex:         reader.ReadUint16(),
	}
}

func constantStringify(constantInfos []ConstantInfo, i uint16) string {
	info := constantInfos[i]
	switch info.(type) {
	case []byte:
		return string(info.([]byte))
	case ConstantClassInfo:
		nameIndex := info.(ConstantClassInfo).NameIndex
		return constantStringify(constantInfos, nameIndex)
	case ConstantStringInfo:
		stringIndex := info.(ConstantStringInfo).StringIndex
		return constantStringify(constantInfos, stringIndex)
	case ConstantFieldRefInfo:
		classIndex := info.(ConstantFieldRefInfo).ClassIndex
		nameAndTypeIndex := info.(ConstantFieldRefInfo).NameAndTypeIndex
		className := constantStringify(constantInfos, classIndex)
		nameAndDescriptor := constantStringify(constantInfos, nameAndTypeIndex)
		return className + "@" + nameAndDescriptor
	case ConstantMethodRefInfo:
		classIndex := info.(ConstantMethodRefInfo).ClassIndex
		nameAndTypeIndex := info.(ConstantMethodRefInfo).NameAndTypeIndex
		className := constantStringify(constantInfos, classIndex)
		nameAndDescriptor := constantStringify(constantInfos, nameAndTypeIndex)
		return className + "." + nameAndDescriptor
	case ConstantInterfaceRefInfo:
		classIndex := info.(ConstantInterfaceRefInfo).ClassIndex
		nameAndTypeIndex := info.(ConstantInterfaceRefInfo).NameAndTypeIndex
		className := constantStringify(constantInfos, classIndex)
		nameAndDescriptor := constantStringify(constantInfos, nameAndTypeIndex)
		return className + "." + nameAndDescriptor
	case ConstantNameAndTypeInfo:
		nameIndex := info.(ConstantNameAndTypeInfo).NameIndex
		descriptorIndex := info.(ConstantNameAndTypeInfo).DescriptorIndex
		name := constantStringify(constantInfos, nameIndex)
		descriptor := constantStringify(constantInfos, descriptorIndex)
		return name + "#" + descriptor
	default:
		return fmt.Sprintf("%#v", info)
	}
}
