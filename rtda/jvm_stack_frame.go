package rtda

type Frame struct {
	next         *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

type Slot struct {
	num int32
	ref *Object
}

func newFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
	}
}
