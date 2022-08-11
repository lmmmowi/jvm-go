package classfile

type SignatureAttribute struct {
	SignatureIndex uint16
	SignatureValue string
}

func readSignatureAttribute(reader *ClassReader) SignatureAttribute {
	index := reader.ReadUint16()
	value := reader.classFile.getConstant(index)
	return SignatureAttribute{
		SignatureIndex: index,
		SignatureValue: value,
	}
}

func (attr SignatureAttribute) print(placeHolder int) {
	_println(placeHolder, "-Index: %v", attr.SignatureIndex)
	_println(placeHolder, "-Value: %v", attr.SignatureValue)
}
