package main

import (
	"math"
	"strconv"
)

// reverse integer
// see: // see:https://leetcode.com/problems/reverse-integer/
func main() {
	println(reverse(123456789234567))
}

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	abs := x
	pi := 1
	if x < 0 {
		abs = int(math.Abs(float64(x)))
		pi = -1
	}

	var s string
	for {
		pop := abs % 10
		abs /= 10
		if pop == 0 {
			break
		}
		s += strconv.Itoa(pop)
	}
	i, _ := strconv.Atoi(s)

	return i * pi
}
