package classfile

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
	return RuntimeVisibleAnnotationsAttribute{
		Annotations: reader.readTable(readAnnotation).([]Annotation),
	}
}

func readAnnotation(reader *ClassReader) Annotation {
	return Annotation{
		TypeName:          reader.classFile.getConstant(reader.ReadUint16()),
		ElementValuePairs: reader.readTable(readElementValuePair).([]ElementValuePair),
	}
}

func readElementValuePair(reader *ClassReader) ElementValuePair {
	nameIndex := reader.ReadUint16()
	name := reader.classFile.getConstant(nameIndex)

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
			Elements: reader.readTable(readElementValue).([]ElementValue),
		}
	default:
		return ConstValue{
			Tag:             tag,
			ConstValueIndex: reader.ReadUint16(),
		}
	}
}

func (attr RuntimeVisibleAnnotationsAttribute) print(placeHolder int) {
	for _, annotation := range attr.Annotations {
		_println(placeHolder, "-[%v]:", annotation.TypeName)

		for _, pair := range annotation.ElementValuePairs {
			_println(placeHolder+1, "-%v: %v", pair.Name, pair.Value)
		}
	}
}
