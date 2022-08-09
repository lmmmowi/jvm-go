package classfile

import (
	"encoding/binary"
	"github.com/lmmmowi/jvm-go/utils"
)

type ClassReader struct {
	utils.BytesReader
}

func newClassReader(data []byte) ClassReader {
	br := utils.NewBytesReader(data, binary.BigEndian)
	return ClassReader{
		BytesReader: br,
	}
}
