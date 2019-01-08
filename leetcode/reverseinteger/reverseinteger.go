package reverseinteger

import (
	"math"
	"strconv"
)

// reverse integer
// see: // see:https://leetcode.com/problems/reverse-integer/
func main() {
	println(reverse(1534236469))
}

func reverse(x int) int {
	if x == 0 || x > 2147483647 || x < -2147483647 {
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
		if abs == 0 && pop == 0 {
			break
		}
		s += strconv.Itoa(pop)
	}
	i, _ := strconv.Atoi(s)

	r := i * pi
	if r > 2147483647 || r < -2147483647 {
		return 0
	}
	return r
}
