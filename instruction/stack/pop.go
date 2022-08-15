package stack

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/********************************************************/
/*****************    pop系列指令共2条    *****************/
/********************************************************/

type Pop struct {
	base.NoOperandsInstruction
}

func (inst *Pop) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type Pop2 struct {
	base.NoOperandsInstruction
}

func (inst *Pop2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
