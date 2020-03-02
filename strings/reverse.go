package strings

func ReverseWords(s string) string {
	var (
		words []string
		word  string
	)
	for _, c := range s {
		if c == ' ' {
			if len(word) == 0 {
				continue
			} else {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(c)
		}
	}

	if len(word) != 0 {
		words = append(words, word)
	}

	var res string
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i != 0 {
			res += " "
		}
	}

	return res
}
