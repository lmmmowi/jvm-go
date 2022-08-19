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

func (frame *Frame) Pop() Slot {
	return frame.operandStack.PopSlot()
}

func (frame *Frame) PopInt() int32 {
	return frame.operandStack.PopInt()
}

func (frame *Frame) PushInt(val int32) {
	frame.operandStack.PushInt(val)
}

func (frame *Frame) GetIntVar(index uint) int32 {
	return frame.localVars[index].num
}

func (frame *Frame) SetIntVar(index uint, val int32) {
	frame.localVars[index].num = val
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}

func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}
