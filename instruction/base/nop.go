package base

import "github.com/lmmmowi/jvm-go/rtda"

type NOP struct {
	NoOperandsInstruction
}

func (inst *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}
