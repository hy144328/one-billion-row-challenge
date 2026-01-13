package main

func parseDigits(digits []byte) int {
	res := 0
	sgn := 1

	if digits[0] == '-' {
		sgn = -1
		digits = digits[1:]
	}

	switch len(digits) {
	case 3:
		res = 10 * int(digits[0] - '0') + int(digits[2] - '0')
	case 4:
		res = 100 * int(digits[0] - '0') + 10 * int(digits[1] - '0') + int(digits[3] - '0')
	}

	return sgn * res
}
