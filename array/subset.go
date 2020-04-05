package array

import (
	"fmt"
	"sort"
	"strings"
)

func SubsetsWithDup(a []int) [][]int {
	sort.Ints(a)

	var dupRes = [][]int{{}}
	dupRes = append(dupRes, subsetsWithDup(a, 0, []int{})...)

	var res [][]int
	exist := make(map[string]bool)
	for _, arr := range dupRes {
		key := strings.Trim(strings.Replace(fmt.Sprint(arr), " ", "", -1), "[]")
		if !exist[key] {
			exist[key] = true
			res = append(res, arr)
		}
	}

	return res
}

func subsetsWithDup(a []int, start int, pref []int) [][]int {
	var res [][]int
	for i := start; i < len(a); i++ {
		newPref := make([]int, len(pref))
		copy(newPref, pref)
		prefIter := append(newPref, a[i])
		res = append(res, prefIter)
		res = append(res, subsetsWithDup(a, i+1, prefIter)...)
	}
	return res
}
