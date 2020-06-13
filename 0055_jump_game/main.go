package main

import (
	"fmt"
	"reflect"
)

/*
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.



Example 1:
	Input: nums = [2,3,1,1,4]
	Output: true
	Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

Example 2:
	Input: nums = [3,2,1,0,4]
	Output: false
	Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
*/

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	maxReachIndex := 0
	finalIndex := len(nums) - 1

	for i := 0; i <= maxReachIndex; i++ {
		if (i + nums[i]) > maxReachIndex {
			maxReachIndex = i + nums[i]
		}
		if maxReachIndex >= finalIndex {
			return true
		}
	}

	return false
}

// test case
func main() {
	result := canJump([]int{2, 3, 1, 1, 4})
	if reflect.DeepEqual(result, true) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = canJump([]int{3, 2, 1, 0, 4})
	if reflect.DeepEqual(result, false) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
