package classfile

type UnparsedAttribute struct {
	Name string
	Info []byte
}

func (attr UnparsedAttribute) print(placeHolder int) {
	_println(placeHolder, "-[%v]: %v", attr.Name, attr.Info)
}
