package stack

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/********************************************************/
/*****************    pop系列指令共2条    *****************/
/********************************************************/

type POP struct {
	base.NoOperandsInstruction
}

func (inst *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (inst *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
