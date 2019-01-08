package main

import "fmt"

func main() {
	nums := []int {1,3,5,6}
	fmt.Println(searchInsert(nums, 2))
}

// Search insert position
// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums)-1, 0
	for left<=right {
		mid = (left+right)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			return mid
		}
	}
	return right + 1
}
