package array

import (
	"sort"
	"strings"
)

func Anagrams(a []string) [][]int {
	var res [][]int

	orderedWords := make(map[string][]int)
	for i := 0; i < len(a); i++ {
		wordSlice := strings.Split(a[i], "")
		sort.Strings(wordSlice)
		sortedWord := strings.Join(wordSlice, "")
		orderedWords[sortedWord] = append(orderedWords[sortedWord], i+1)
	}

	for _, w := range a {
		wordSlice := strings.Split(w, "")
		sort.Strings(wordSlice)
		sortedWord := strings.Join(wordSlice, "")
		if ixs, ok := orderedWords[sortedWord]; ok {
			res = append(res, ixs)
			delete(orderedWords, sortedWord)
		}
	}

	return res
}
