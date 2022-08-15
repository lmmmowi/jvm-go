package consts

import (
	"github.com/lmmmowi/jvm-go/instruction"
	"github.com/lmmmowi/jvm-go/rtda"
)

/**********************************************************/
/*****************    ipush系列指令共2条    *****************/
/**********************************************************/

type BIPush struct {
	val int8
}

func (inst *BIPush) FetchOperands(reader *instruction.BytecodeReader) {
	inst.val = reader.ReadInt8()
}

func (inst *BIPush) Execute(frame *rtda.Frame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}

type SIPush struct {
	val int16
}

func (inst *SIPush) FetchOperands(reader *instruction.BytecodeReader) {
	inst.val = reader.ReadInt16()
}

func (inst *SIPush) Execute(frame *rtda.Frame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}
