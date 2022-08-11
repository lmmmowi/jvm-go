package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.size {
		panic("java.lang.StackOverflowError")
	} else {
		stack.size++
	}

	if stack._top == nil {
		stack._top = frame
	} else {
		stack._top.next = frame
	}
}

func (stack *Stack) pop() *Frame {
	if stack.size == 0 {
		panic("thread stack is empty")
	} else {
		stack.size--
	}

	frame := stack._top
	stack._top = frame.next
	return frame
}

func (stack *Stack) top() *Frame {
	if stack.size == 0 {
		panic("thread stack is empty")
	}

	return stack._top
}
