package stack

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/*********************************************************/
/*****************    swap系列指令共1条    *****************/
/*********************************************************/

type Swap struct {
	base.NoOperandsInstruction
}

func (inst *Swap) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

