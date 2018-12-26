package main

import "fmt"

func main() {
	fmt.Println(romanToInt("IIIVI"))
}

func romanToInt(s string) (result int) {
	for i, n := range s {
		if i != 0 && values[rune(s[i-1])] < values[n] {
			result -= 2 * values[rune(s[i-1])]
		}
		result += values[n]
	}
	return
}

var values = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
