package main

import "fmt"

// Count the stairs
func main() {
	n := 7
	for !((n%2 == 1) &&
		(n%3 == 2) &&
		(n%5 == 4) &&
		(n%6 == 5) &&
		(n%7 == 0)) {
		n = n + 7
	}

	fmt.Printf("count the stairs= %d\n", n)
}
