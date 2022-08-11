package classfile

import (
	"encoding/binary"
	"github.com/lmmmowi/jvm-go/utils"
	"reflect"
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

func (reader *ClassReader) readTable(readFn interface{}) interface{} {
	n := int(reader.ReadUint16())

	itemType := reflect.TypeOf(readFn).Out(0)
	sliceType := reflect.SliceOf(itemType)
	arr := reflect.MakeSlice(sliceType, n, n)

	readFnVal := reflect.ValueOf(readFn)
	args := []reflect.Value{reflect.ValueOf(reader)}

	for i := 0; i < n; i++ {
		x := readFnVal.Call(args)[0]
		arr.Index(i).Set(x) // s[i] = x
	}

	return arr.Interface()
}