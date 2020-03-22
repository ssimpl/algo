package stack

import (
	"strconv"

	"github.com/golang-collections/collections/stack"
)

func EvalRPN(a []string) int {
	if len(a) == 1 {
		num, _ := strconv.Atoi(a[0])
		return num
	}

	stk := stack.New()
	for i := 0; i < len(a); i++ {
		s := a[i]
		if s == "+" || s == "-" || s == "/" || s == "*" {
			y := stk.Pop().(int)
			x := stk.Pop().(int)
			var res int
			switch s {
			case "+":
				res = x + y
			case "-":
				res = x - y
			case "*":
				res = x * y
			case "/":
				res = x / y
			}

			if stk.Len() == 0 && i == len(a)-1 {
				return res
			} else {
				stk.Push(res)
			}
		} else {
			num, _ := strconv.Atoi(s)
			stk.Push(num)
		}
	}

	return 0
}
