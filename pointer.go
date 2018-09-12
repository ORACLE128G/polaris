package main

import "fmt"

// Val transfer
func passVal(a int){
	a++
}

// Swap Val.
func swap (a, b *int){
	*a, *b = *b, *a
}
func main() {
	c, d := 1, 2
	swap(&c, &d)
	fmt.Println(c, d)
	a := 1
	passVal(a)
	fmt.Println(a)
}
