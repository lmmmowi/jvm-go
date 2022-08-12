package rtda

import "math"

type OperandStack struct {
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	return &OperandStack{
		size:  0,
		slots: make([]Slot, maxStack),
	}
}

func (stack *OperandStack) PushInt(val int32) {
	stack.slots[stack.size].num = val
	stack.size++
}

func (stack *OperandStack) PopInt() int32 {
	stack.size--
	return stack.slots[stack.size].num
}

func (stack *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	stack.slots[stack.size].num = int32(bits)
	stack.size++
}

func (stack *OperandStack) PopFloat() float32 {
	stack.size--
	bits := uint32(stack.slots[stack.size].num)
	return math.Float32frombits(bits)
}

func (stack *OperandStack) PushLong(val int64) {
	stack.slots[stack.size].num = int32(val)
	stack.slots[stack.size+1].num = int32(val >> 32)
	stack.size += 2
}

func (stack *OperandStack) PopLong() int64 {
	stack.size -= 2
	low := uint32(stack.slots[stack.size].num)
	high := uint32(stack.slots[stack.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (stack *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	stack.PushLong(int64(bits))
}

func (stack *OperandStack) PopDouble() float64 {
	bits := uint64(stack.PopLong())
	return math.Float64frombits(bits)
}

func (stack *OperandStack) PushRef(ref *Object) {
	stack.slots[stack.size].ref = ref
	stack.size++
}

func (stack *OperandStack) PopRef() *Object {
	stack.size--
	return stack.slots[stack.size].ref
}

func (stack *OperandStack) PushSlot(slot Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}

func (stack *OperandStack) PopSlot() Slot {
	stack.size--
	return stack.slots[stack.size]
}
