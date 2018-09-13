package main

import "fmt"

func ops() {
	// Create a slice, but not array
	var s [] int
	fmt.Println(s)
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	fmt.Println("After append is:", s)

	s1 := []int{0, 1, 2, 4}

	// Make a s slice
	s2 := make([]int, 32, 32)

	fmt.Println(s1)
	fmt.Println(s2)
}

func copySlice() {
	s1 := []int{0, 1}
	s2 := []int{2, 3}
	copy(s1, s2)
	fmt.Println(s2)
}

func appendz () {
	s1 := []int{0, 1, 2, 3, 4, 5, 6,7}
	ints := append(s1[:3], s1[:8]...)
	fmt.Println("appendz is :",ints)
}

func popping() {
	s1 := []int{0, 1, 2, 3, 4, 5, 6,7}
	front := s1[:1]
	fmt.Println("Popping arr front is :", front)
	back := s1[len(s1) - 1:]
	fmt.Println("Popping arr back is :", back)
}
func main() {
	appendz()
	copySlice()
	ops()

	popping()
}
