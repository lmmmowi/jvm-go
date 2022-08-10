package classfile

import "fmt"

type RuntimeVisibleAnnotationsAttribute struct {
	Annotations []Annotation
}

type Annotation struct {
	TypeName          string
	ElementValuePairs []ElementValuePair
}

type ElementValuePair struct {
	Name  string
	Value ElementValue
}

type ElementValue interface{}

type ConstValue struct {
	Tag             uint8
	ConstValueIndex uint16
}

type EnumConstValue struct {
	TypeNameIndex  uint16
	ConstNameIndex uint16
}

type ClassInfoValue struct {
	ClassInfoIndex uint16
}

type AnnotationValue struct {
	NestedAnnotation Annotation
}

type ArrayValue struct {
	Elements []ElementValue
}

func readRuntimeVisibleAnnotationsAttribute(reader *ClassReader) RuntimeVisibleAnnotationsAttribute {
	count := reader.ReadUint16()
	annotations := make([]Annotation, count)
	for i := range annotations {
		annotations[i] = readAnnotation(reader)
	}

	return RuntimeVisibleAnnotationsAttribute{
		Annotations: annotations,
	}
}

func readAnnotation(reader *ClassReader) Annotation {
	typeIndex := reader.ReadUint16()
	typeName := reader.classFile.constantPool.stringify(typeIndex)

	count := reader.ReadUint16()
	elementValuePairs := make([]ElementValuePair, count)
	for i := range elementValuePairs {
		elementValuePairs[i] = readElementValuePair(reader)
	}

	return Annotation{
		TypeName:          typeName,
		ElementValuePairs: elementValuePairs,
	}
}

func readElementValuePair(reader *ClassReader) ElementValuePair {
	nameIndex := reader.ReadUint16()
	name := reader.classFile.constantPool.stringify(nameIndex)

	return ElementValuePair{
		Name:  name,
		Value: readElementValue(reader),
	}
}

func readElementValue(reader *ClassReader) ElementValue {
	tag := reader.ReadUint8()
	switch tag {
	case 'e':
		return EnumConstValue{
			TypeNameIndex:  reader.ReadUint16(),
			ConstNameIndex: reader.ReadUint16(),
		}
	case 'c':
		return ClassInfoValue{
			ClassInfoIndex: reader.ReadUint16(),
		}
	case '@':
		return AnnotationValue{
			NestedAnnotation: readAnnotation(reader),
		}
	case '[':
		return ArrayValue{
			Elements: readElementValues(reader),
		}
	default:
		return ConstValue{
			Tag:             tag,
			ConstValueIndex: reader.ReadUint16(),
		}
	}
}

func readElementValues(reader *ClassReader) []ElementValue {
	count := reader.ReadUint16()
	elementValues := make([]ElementValue, count)
	for i := range elementValues {
		elementValues[i] = readElementValue(reader)
	}
	return elementValues
}

func (attr RuntimeVisibleAnnotationsAttribute) getName() string {
	return RuntimeVisibleAnnotations
}

func (attr RuntimeVisibleAnnotationsAttribute) print() {
	for _, annotation := range attr.Annotations {
		fmt.Printf("\t\t-[%v]:\n", annotation.TypeName)

		for _, pair := range annotation.ElementValuePairs {
			fmt.Printf("\t\t\t-%v: %v\n", pair.Name, pair.Value)
		}
	}
}
