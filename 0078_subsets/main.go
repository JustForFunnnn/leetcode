package main

import (
	"fmt"
)

/*
Given a set of distinct integers, nums, return all possible subsets (the power set).

Note: The solution set must not contain duplicate subsets.
*/

func subsets(nums []int) (result [][]int) {
	travelNums(nums, []int{}, &result)
	return result
}

func travelNums(nums []int, subset []int, result *[][]int) {
	if len(nums) == 0 {
		temp := make([]int, len(subset))
		copy(temp, subset)
		*result = append(*result, temp)
		return
	}

	// without current number
	travelNums(nums[1:], subset, result)

	// with current number
	subset = append(subset, nums[0])
	travelNums(nums[1:], subset, result)
}

// test case
func main() {
	result := subsets([]int{1, 2, 3})
	fmt.Println(result)
}
