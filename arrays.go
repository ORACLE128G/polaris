package main

import "fmt"

var arr [4]int
var arr1 = [3]int{1, 2, 3}
var arr2 = [...]int{1, 2, 3, 4, 5}
var arr3 [2][3]int

func foreachArr() {
	for i, v := range arr2 {
		fmt.Printf("arr1 values is %d, index is %d \n",v, i)
	}
	for _, i:= range arr3 {
		fmt.Printf("i is : %d\n", i)
	}
}

func PrintArr (arr [5]int) {
	for _, v := range arr {
		fmt.Printf("v is : %d \n" ,v)
	}
}
func main() {
	PrintArr(arr2)
	foreachArr()
	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
}
