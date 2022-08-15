package consts

import (
	"github.com/lmmmowi/jvm-go/instruction/base"
	"github.com/lmmmowi/jvm-go/rtda"
)

/*********************************************************/
/*****************   const系列指令共15条   *****************/
/*********************************************************/

/*********************** object(1) ***********************/

type AconstNull struct {
	base.NoOperandsInstruction
}

func (inst *AconstNull) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

/*********************** int(7) ***********************/

type IConst0 struct {
	base.NoOperandsInstruction
}

func (inst *IConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type IConst1 struct {
	base.NoOperandsInstruction
}

func (inst *IConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

type IConst2 struct {
	base.NoOperandsInstruction
}

func (inst *IConst2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type IConst3 struct {
	base.NoOperandsInstruction
}

func (inst *IConst3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type IConst4 struct {
	base.NoOperandsInstruction
}

func (inst *IConst4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type IConst5 struct {
	base.NoOperandsInstruction
}

func (inst *IConst5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

type IConstM1 struct {
	base.NoOperandsInstruction
}

func (inst *IConstM1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

/*********************** long(2) ***********************/

type LConst0 struct {
	base.NoOperandsInstruction
}

func (inst *LConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LConst1 struct {
	base.NoOperandsInstruction
}

func (inst *LConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

/*********************** float(3) ***********************/

type FConst0 struct {
	base.NoOperandsInstruction
}

func (inst *FConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

type FConst1 struct {
	base.NoOperandsInstruction
}

func (inst *FConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

type FConst2 struct {
	base.NoOperandsInstruction
}

func (inst *FConst2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

/*********************** double(2) ***********************/

type DConst0 struct {
	base.NoOperandsInstruction
}

func (inst *DConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DConst1 struct {
	base.NoOperandsInstruction
}

func (inst *DConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}
