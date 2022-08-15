package base

import "github.com/lmmmowi/jvm-go/rtda"

type Nop struct {
	NoOperandsInstruction
}

func (inst *Nop) Execute(frame *rtda.Frame) {
	// do nothing
}
