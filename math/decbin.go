package math

import "strconv"

func DecimalToBinary(n int) string {
	if n == 0 {
		return "0"
	}

	var binary string
	for n > 0 {
		rem := n % 2
		binary = strconv.Itoa(rem) + binary
		n = n / 2
	}

	return binary
}
