package main

import "fmt"

func main() {
	var slice = []int{1, 1, 2}
	length := removeDuplicates(slice)
	for i := 0; i < length; i++ {
		fmt.Printf("%v, ", slice[i])
	}

}

// Remove duplicates from sorted array.
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	length := 1
	for i, e := range nums {
		if i > 0 && nums[length-1] != e {
			nums[length] = e
			length++
		}
	}
	return length
}
