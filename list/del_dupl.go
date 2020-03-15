package list

func DeleteDuplicates(a *node) *node {
	head := a
	prev := a
	a = a.next
	for a != nil {
		if prev.value != a.value {
			prev.next = a
			prev = a
		}
		a = a.next
	}
	prev.next = nil
	return head
}
