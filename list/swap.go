package list

func SwapPairs(a *node) *node {
	if a.next == nil {
		return a
	}

	var (
		cur, head  = a, a.next
		prev, next *node
	)
	for cur != nil {
		if prev == nil {
			prev = cur
			cur = cur.next
			continue
		}

		next = cur.next
		cur.next = prev
		if next != nil {
			if next.next == nil {
				prev.next = next
			} else {
				prev.next = next.next
			}
		} else {
			prev.next = nil
		}
		prev = nil
		cur = next
	}

	return head
}
