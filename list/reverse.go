package list

func Reverse(a *node) *node {
	head := a

	var reverse func(a *node)
	reverse = func(a *node) {
		if a.next == nil {
			head = a
			return
		}
		reverse(a.next)
		tmp := a.next
		tmp.next = a
		a.next = nil
	}
	reverse(a)

	return head
}

func Reverse2(a *node, prev *node) *node {
	next := a.next
	a.next = prev
	if next == nil {
		return a
	}
	return Reverse2(next, a)
}
