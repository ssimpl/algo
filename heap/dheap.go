package heap

const (
	dHeapMaxN = 101
	dHeapMod  = 1000000007
)

type dHeap struct {
	dp   []int
	nck  [][]int
	log2 []int
}

func DHeap(n int) int {
	dp := make([]int, dHeapMaxN)
	for i := 0; i <= n; i++ {
		dp[i] = -1
	}

	nck := make([][]int, dHeapMaxN)
	for i := 0; i <= n; i++ {
		if nck[i] == nil {
			nck[i] = make([]int, dHeapMaxN)
		}
		for j := 0; j <= n; j++ {
			nck[i][j] = -1
		}
	}

	currLog2 := -1
	currPower2 := 1
	log2 := make([]int, dHeapMaxN)
	for i := 1; i <= n; i++ {
		if currPower2 == i {
			currLog2++
			currPower2 *= 2
		}
		log2[i] = currLog2
	}

	dHeap := &dHeap{
		dp:   dp,
		nck:  nck,
		log2: log2,
	}

	return dHeap.numberOfHeaps(n)
}

func (dh *dHeap) numberOfHeaps(n int) int {
	if n <= 1 {
		return 1
	}

	if dh.dp[n] != -1 {
		return dh.dp[n]
	}

	left := dh.getLeft(n)
	ans := (dh.choose(n-1, left) * dh.numberOfHeaps(left)) % dHeapMod * dh.numberOfHeaps(n-1-left) % dHeapMod

	dh.dp[n] = ans

	return ans
}

func (dh *dHeap) getLeft(n int) int {
	if n == 1 {
		return 0
	}

	h := dh.log2[n]
	numH := 1 << h
	last := n - ((1 << h) - 1)

	if last >= numH/2 {
		return numH - 1
	} else {
		return numH - 1 - ((numH / 2) - last)
	}
}

func (dh *dHeap) choose(n int, k int) int {
	if k > n {
		return 0
	}

	if n <= 1 || k == 0 {
		return 1
	}

	if dh.nck[n][k] != -1 {
		return dh.nck[n][k]
	}

	answer := (dh.choose(n-1, k-1) + dh.choose(n-1, k)) % dHeapMod
	dh.nck[n][k] = answer

	return answer
}
