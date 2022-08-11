package main

import (
	"fmt"
	"github.com/lmmmowi/jvm-go/rtda"
)

func main() {
	localVars := rtda.NewLocalVars(10)
	localVars.SetLong(0, -12452345453333)
	a := localVars.GetLong(0)
	fmt.Println(a)

	ops := rtda.NewOperandStack(10)
	ops.PushFloat(323.44)
	fmt.Println(ops.PopFloat())
}
