package classfile

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool []ConstantInfo
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []FieldInfo
	methods      []MethodInfo
	attributes   []AttributeInfo
}

func ParseClassFile(data []byte) *ClassFile {
	reader := newClassReader(data)
	cf := &ClassFile{}
	reader.bind(cf)
	cf.readMagic(&reader)
	cf.readMinorVersion(&reader)
	cf.readMajorVersion(&reader)
	cf.readConstantPool(&reader)
	cf.readAccessFlags(&reader)
	cf.readThisClass(&reader)
	cf.readSuperClass(&reader)
	cf.readInterfaces(&reader)
	cf.readFields(&reader)
	cf.readMethods(&reader)
	cf.readAttributes(&reader)
	return cf
}

func (cf *ClassFile) readMagic(reader *ClassReader) {
	cf.magic = reader.ReadUint32()
}

func (cf *ClassFile) readMinorVersion(reader *ClassReader) {
	cf.minorVersion = reader.ReadUint16()
}

func (cf *ClassFile) readMajorVersion(reader *ClassReader) {
	cf.majorVersion = reader.ReadUint16()
}

func (cf *ClassFile) readConstantPool(reader *ClassReader) {
	cf.constantPool = readConstantPool(reader)
}

func (cf *ClassFile) readAccessFlags(reader *ClassReader) {
	cf.accessFlags = reader.ReadUint16()
}

func (cf *ClassFile) readThisClass(reader *ClassReader) {
	cf.thisClass = reader.ReadUint16()
}

func (cf *ClassFile) readSuperClass(reader *ClassReader) {
	cf.superClass = reader.ReadUint16()
}

func (cf *ClassFile) readInterfaces(reader *ClassReader) {
	cf.interfaces = reader.readTable(func(r *ClassReader) uint16 {
		return r.ReadUint16()
	}).([]uint16)
}

func (cf *ClassFile) readFields(reader *ClassReader) {
	cf.fields = reader.readTable(readField).([]FieldInfo)
}

func (cf *ClassFile) readMethods(reader *ClassReader) {
	cf.methods = reader.readTable(readMethod).([]MethodInfo)
}

func (cf *ClassFile) readAttributes(reader *ClassReader) {
	cf.attributes = readAttributes(reader)
}

func (cf *ClassFile) getConstant(index uint16) string {
	return constantStringify(cf.constantPool, index)
}
