package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	length := removeElement(nums, 3)

	for i := 0; i < length; i++ {
		fmt.Printf("%d ", nums[i])
	}
}

// Remove element
// https://leetcode.com/problems/remove-element/
func removeElement(nums []int, val int) int {

	length := 0
	for i, v := range nums {
		if v != val {
			temp := nums[i]
			nums[i] = nums[length]
			nums[length] = temp
			length++
		}
	}
	return length
}
