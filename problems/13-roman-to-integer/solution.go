// Problem: https://leetcode.com/problems/roman-to-integer/
// Submission: https://leetcode.com/problems/roman-to-integer/submissions/852125590/

package romanToInteger

// Given a roman input, returns the corresponding number
func RomanToInt(s string) int {
	runes := []rune(s)
	l := len(runes)
	total := 0

	for i := 0; i < l; i++ {
		var next *rune

		current := &runes[i]
		if i != l-1 {
			next = &runes[i+1]
		}

		value, skip := convert(current, next)
		total += value
		if skip {
			i += 1
		}
	}

	return total
}

func convert(current *rune, next *rune) (int, bool) {
	if current == nil {
		return 0, false
	}

	switch *current {
	case 'I':
		if next != nil {
			switch *next {
			case 'V':
				return 4, true

			case 'X':
				return 9, true
			}
		}
		return 1, false

	case 'V':
		return 5, false

	case 'X':
		if next != nil {
			switch *next {
			case 'L':
				return 40, true

			case 'C':
				return 90, true
			}
		}
		return 10, false

	case 'L':
		return 50, false

	case 'C':
		if next != nil {
			switch *next {
			case 'D':
				return 400, true

			case 'M':
				return 900, true
			}
		}

		return 100, false

	case 'D':
		return 500, false

	case 'M':
		return 1000, false
	}

	return 0, false
}
