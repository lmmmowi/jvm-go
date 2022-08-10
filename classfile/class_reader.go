package classfile

import (
	"encoding/binary"
	"github.com/lmmmowi/jvm-go/utils"
)

type ClassReader struct {
	utils.BytesReader
	classFile *ClassFile
}

func newClassReader(data []byte) ClassReader {
	br := utils.NewBytesReader(data, binary.BigEndian)
	return ClassReader{
		BytesReader: br,
	}
}

func (reader *ClassReader) bind(classFile *ClassFile) {
	reader.classFile = classFile
}
