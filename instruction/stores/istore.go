package stores

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/***********************************************************/
/*****************    istore系列指令共5条    *****************/
/***********************************************************/

type IStore struct {
	base.Index8Instruction
}

func (inst *IStore) Execute(frame *rtda.Frame) {
	_istore(frame, inst.Index)
}

type IStore0 struct {
	base.NoOperandsInstruction
}

func (inst *IStore0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

type IStore1 struct {
	base.NoOperandsInstruction
}

func (inst *IStore1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

type IStore2 struct {
	base.NoOperandsInstruction
}

func (inst *IStore2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

type IStore3 struct {
	base.NoOperandsInstruction
}

func (inst *IStore3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}

func _istore(frame *rtda.Frame, index uint) {
	i := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, i)
}
