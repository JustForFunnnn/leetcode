package main

import (
	"fmt"
)

/*
Given an unsorted integer array, find the smallest missing positive integer.

Example 1:
	Input: [1,2,0]
	Output: 3

Example 2:
	Input: [3,4,-1,1]
	Output: 2

Example 3:
	Input: [7,8,9,11,12]
	Output: 1

Note:
Your algorithm should run in O(n) time and uses constant extra space.
*/
func firstMissingPositive(nums []int) int {
	for i, _ := range nums {
		for ;0 < nums[i] && nums[i] <= len(nums) && nums[nums[i] - 1] != nums[i]; {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, num := range nums {
		if num != (i + 1) {
			return i + 1
		}
	}

	return len(nums) + 1
}

// test case
func main() {
	result := firstMissingPositive([]int{1, 2, 0})
	if result != 3 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = firstMissingPositive([]int{3, 4, -1, 1})
	if result != 2 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = firstMissingPositive([]int{7, 8, 9, 11, 12})
	if result != 1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = firstMissingPositive([]int{1})
	if result != 2 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = firstMissingPositive([]int{2})
	if result != 1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = firstMissingPositive([]int{-1})
	if result != 1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = firstMissingPositive([]int{7, 12})
	if result != 1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
