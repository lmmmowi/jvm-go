package instruction

import (
	"github.com/lmmmowi/jvm-go/rtda"
)

/**********************************************************/
/*****************    ipush系列指令共2条    *****************/
/**********************************************************/

type BIPUSH struct {
	val int8
}

func (inst *BIPUSH) FetchOperands(reader *BytecodeReader) {
	inst.val = reader.ReadInt8()
}

func (inst *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (inst *SIPUSH) FetchOperands(reader *BytecodeReader) {
	inst.val = reader.ReadInt16()
}

func (inst *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(inst.val)
	frame.OperandStack().PushInt(i)
}
