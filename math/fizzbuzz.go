package math

import "strconv"

func FizzBuzz(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		var item string
		if i%3 == 0 {
			item += "Fizz"
		}
		if i%5 == 0 {
			item += "Buzz"
		}
		if len(item) == 0 {
			item = strconv.Itoa(i)
		}
		res = append(res, item)
	}
	return res
}
