package stack

import "github.com/golang-collections/collections/stack"

func ParenthesesIsValid(a string) bool {
	stk := stack.New()
	for _, r := range a {
		c := string(r)
		if c == "(" || c == "{" || c == "[" {
			stk.Push(c)
			continue
		}

		if stk.Len() == 0 {
			return false
		}

		top := stk.Peek().(string)
		if top == "(" && c == ")" || top == "{" && c == "}" || top == "[" && c == "]" {
			stk.Pop()
		} else {
			return false
		}
	}
	return stk.Len() == 0
}
