package twosum

import (
	"errors"
	"fmt"
)

func main() {
	v := twoSum([]int{5, 7, 12, 15}, 12)
	fmt.Printf("two sum solution: %v \n", v)
}

func twoSum(nums []int, target int) [] int {
	m := make(map[int]int)
	for i := range nums {
		complement := target - nums[i]
		if _, ok := m[complement]; ok {
			return []int{m[complement], i}
		}

		m[nums[i]] = i
	}
	panic(errors.New("no two sum solution"))
}
