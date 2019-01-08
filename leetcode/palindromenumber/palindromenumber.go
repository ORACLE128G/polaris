package palindromenumber

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(121))
}


func isPalindrome(x int) bool {

	xs := strconv.Itoa(x)
	// m := make([]rune, len(xs))
	var m []rune
	for _, n := range xs{
		m = append(m, n)
	}

	rs := []rune(xs)
	count := 0
	for i := len(rs) -1; i>=0;i--{
		if rs[i] != m[count]{
			return false
		}
		count++
	}
	return true
}