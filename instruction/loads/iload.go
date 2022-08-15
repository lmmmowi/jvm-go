package loads

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/**********************************************************/
/*****************    iload系列指令共5条    *****************/
/**********************************************************/

type ILoad struct {
	base.Index8Instruction
}

func (inst *ILoad) Execute(frame *rtda.Frame) {
	_iload(frame, inst.Index)
}

type ILoad0 struct {
	base.NoOperandsInstruction
}

func (inst *ILoad0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILoad1 struct {
	base.NoOperandsInstruction
}

func (inst *ILoad1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILoad2 struct {
	base.NoOperandsInstruction
}

func (inst *ILoad2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILoad3 struct {
	base.NoOperandsInstruction
}

func (inst *ILoad3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

func _iload(frame *rtda.Frame, index uint) {
	i := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(i)
}
