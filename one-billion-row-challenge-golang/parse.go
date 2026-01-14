package main

import (
	"strconv"
)

func parseDigitsFromString(digits string) int {
	temperature10, err := strconv.Atoi(digits[:len(digits)-2])
	if err != nil {
		panic(err)
	}

	temperature1 := digits[len(digits)-1] - '0'
	return 10*temperature10 + int(temperature1)
}

func parseDigitsFromBytes(digits []byte) int {
	switch len(digits) {
	case 3:
		return 10*int(digits[0]-'0') + int(digits[2]-'0')
	case 4:
		return 100*int(digits[0]-'0') + 10*int(digits[1]-'0') + int(digits[3]-'0')
	default:
		panic(string(digits))
	}
}
