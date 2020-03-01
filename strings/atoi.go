package strings

import (
	"math"
)

func Atoi(a string) int {
	const (
		space = ' '
		minus = '-'
		plus  = '+'
		zero  = '0'
		nine  = '9'
	)

	var (
		isNegative bool
		number     int
		start      bool
		digits     []int
	)
	for _, r := range a {
		if r != space && r != minus && r != plus && (r < zero || r > nine) {
			break
		}

		if start && (r < zero || r > nine) {
			break
		}

		if r == space {
			continue
		}

		if r == minus {
			isNegative = true
			start = true
			continue
		}

		if r == plus {
			start = true
			continue
		}

		start = true
		digits = append(digits, int(r)-zero)
	}

	n := len(digits) - 1
	for _, d := range digits {
		add := int(math.Pow(10, float64(n))) * d
		if isNegative && math.MinInt32+number > -1*add {
			return math.MinInt32
		} else if math.MaxInt32-number < add {
			return math.MaxInt32
		}
		number += add
		n--
	}

	if isNegative {
		number = -1 * number
	}

	return number
}
