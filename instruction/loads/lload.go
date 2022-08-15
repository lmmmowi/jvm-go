package loads

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/**********************************************************/
/*****************    lload系列指令共5条    *****************/
/**********************************************************/

type LLoad struct {
	base.Index8Instruction
}

func (inst *LLoad) Execute(frame *rtda.Frame) {
	_lload(frame, inst.Index)
}

type LLoad0 struct {
	base.NoOperandsInstruction
}

func (inst *ILoad0) LLOAD_0(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLoad1 struct {
	base.NoOperandsInstruction
}

func (inst *LLoad1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLoad2 struct {
	base.NoOperandsInstruction
}

func (inst *LLoad2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLoad3 struct {
	base.NoOperandsInstruction
}

func (inst *LLoad3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	i := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(i)
}
