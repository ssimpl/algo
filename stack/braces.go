package stack

import "github.com/golang-collections/collections/stack"

func Braces(a string) bool {
	stk := stack.New()
	for _, r := range a {
		c := string(r)
		if c == "+" || c == "-" || c == "/" || c == "*" || c == " " {
			continue
		}

		if c != ")" {
			stk.Push(c)
			continue
		}

		if stk.Peek().(string) == "(" {
			return true
		} else {
			stk.Pop()                       //delete second operand
			if stk.Peek().(string) == "(" { //if the second operand was the only one
				return true
			} else {
				stk.Pop()     //delete first operand
				stk.Pop()     //delete brace
				stk.Push("a") //add operand instead deleted expression
			}
		}
	}

	return false
}
