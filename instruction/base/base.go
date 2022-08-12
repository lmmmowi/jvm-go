package base

import (
	"github.com/lmmmowi/jvm-go/instruction"
	"github.com/lmmmowi/jvm-go/rtda"
)

type Instruction interface {
	FetchOperands(reader *instruction.BytecodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (inst *NoOperandsInstruction) FetchOperands() {
	// do nothing
}

type BranchInstruction struct {
	Offset int
}

func (inst *BranchInstruction) FetchOperands(reader *instruction.BytecodeReader) {
	inst.Offset = int(reader.ReadUint16())
}

// Index8Instruction 用于通过索引访问局部变量表
type Index8Instruction struct {
	Index uint
}

func (inst *Index8Instruction) FetchOperands(reader *instruction.BytecodeReader) {
	inst.Index = uint(reader.ReadUint8())
}

// Index16Instruction 用于通过索引访问运行时常量池
type Index16Instruction struct {
	Index uint
}

func (inst *Index16Instruction) FetchOperands(reader *instruction.BytecodeReader) {
	inst.Index = uint(reader.ReadUint16())
}
