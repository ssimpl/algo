package strings

import (
	"sort"
	"strings"
)

// 14. Longest Common Prefix
func LongestCommonPrefix(a []string) string {
	if len(a) == 0 {
		return ""
	}
	if len(a) == 1 {
		return a[0]
	}

	min := len(a[0])
	for i := 1; i < len(a); i++ {
		if len(a[i]) < min {
			min = len(a[i])
		}
	}
	if min == 0 {
		return ""
	}

	var pref string
	for i := 0; i < min; i++ {
		for j := 1; j < len(a); j++ {
			if a[j][i] != a[0][i] {
				return pref
			}
		}
		pref += string(a[0][i])
	}

	return pref
}

func LongestCommonPrefix2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	sort.Strings(strs)

	var prefix string
	first, last := strs[0], strs[len(strs)-1]
	for i := 0; i < len(first) && i < len(last); i++ {
		if first[i] == last[i] {
			prefix = first[:i+1]
		} else {
			break
		}
	}

	return prefix
}

// 1455. Check If a Word Occurs As a Prefix of Any Word in a Sentence
func IsPrefixOfWord(sent string, word string) int {
	ss := strings.Split(sent, " ")

	for i, w := range ss {
		if len(w) < len(word) {
			continue
		}
		if strings.Compare(w[:len(word)], word) == 0 {
			return i + 1
		}
	}

	return -1
}
