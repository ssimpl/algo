package stack

import "github.com/golang-collections/collections/stack"

func ReverseString(a string) string {
	stk := stack.New()
	for _, c := range a {
		stk.Push(string(c))
	}

	var res string
	for stk.Len() != 0 {
		res += stk.Pop().(string)
	}

	return res
}
