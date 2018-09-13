package main

import "fmt"

func updateSlices(s [] int) {
	s[0] = 100
}
func slices() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6,}
	fmt.Println(arr[2:6])
	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])
}

func sliceExten() []int{
	arr := [...]int {0, 1, 2, 3, 4, 5, 6, 7}
	// s1[2, 3, 4, 5]
	s1 := arr[2:6]
	fmt.Println("s1 is :" , s1)
	//
	s2 := s1[3:len(arr) - 2 - 1]
	fmt.Println("s2 is :" , s2)
	return s2
}

func appends () {
	arr := [...]int {0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[:]
	ss := append(s1, 10, 111, 123)
	fmt.Println("ss is :", ss)
	fmt.Println("arr is :", arr)

}


func main() {

	appends()
	fmt.Println("sliceExten return val is : " ,sliceExten())


	arr := [...]int{0, 1, 2, 3, 4, 5, 6,}
	ints := arr[2:6]

	fmt.Println(arr[2:])
	fmt.Println(arr[:6])
	fmt.Println(arr[:])
	fmt.Println(ints)
	fmt.Println("After updateSlice.")
	updateSlices(ints)
	fmt.Println(ints)
	fmt.Println("arr is " , arr)
	fmt.Println("Reslice :")
	fmt.Println(arr[:][1:][2:])

}
