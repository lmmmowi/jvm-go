package classfile

import "fmt"

type SignatureAttribute struct {
	SignatureIndex uint16
	SignatureValue string
}

func readSignatureAttribute(reader *ClassReader) SignatureAttribute {
	index := reader.ReadUint16()
	value := reader.classFile.constantPool.stringify(index)
	return SignatureAttribute{
		SignatureIndex: index,
		SignatureValue: value,
	}
}

func (attr SignatureAttribute) getName() string {
	return Signature
}

func (attr SignatureAttribute) print() {
	fmt.Printf("\t\t-Index: %v\n", attr.SignatureIndex)
	fmt.Printf("\t\t-Value: %v\n", attr.SignatureValue)
}
