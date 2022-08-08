package classfile

import (
	"encoding/binary"
	"github.com/lmmmowi/jvm-go/utils"
)

type ClassReader struct {
	utils.BytesReader
}

func ParseClassFile(data []byte) *ClassFile {
	cr := newClassReader(data)
	cf := &ClassFile{}
	cf.readMagic(&cr)
	cf.readMinorVersion(&cr)
	cf.readMajorVersion(&cr)
	cf.readConstantPool(&cr)
	return cf
}

func newClassReader(data []byte) ClassReader {
	br := utils.NewBytesReader(data, binary.BigEndian)
	return ClassReader{
		BytesReader: br,
	}
}
