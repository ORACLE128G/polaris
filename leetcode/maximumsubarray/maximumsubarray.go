package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

// Maximum subarray
// https://leetcode.com/problems/maximum-subarray/submissions/
func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// Kadane's Algorithm
	maxSoFar, maxEndingHere := ^int(^uint(0) >> 1), 0
	for _, v := range nums {
		maxEndingHere = maxEndingHere + v

		if maxSoFar < maxEndingHere {
			maxSoFar = maxEndingHere
		}
		if maxEndingHere < 0 {
			maxEndingHere = 0
		}
	}
	return maxSoFar
}
