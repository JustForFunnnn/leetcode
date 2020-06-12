package main

import (
	"fmt"
	"reflect"
)

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:
	Input: [-2,1,-3,4,-1,2,1,-5,4],
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

Follow up:
	If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func maxSubArray(nums []int) (maxSum int) {
	if len(nums) == 0 {
		// do we really need to consider 0 as largest sum of Nil array?
		return 0
	}
	nowSum := nums[0]
	maxSum = nowSum

	for _, num := range nums[1:] {
		if nowSum < 0 {
			nowSum = num
		} else {
			nowSum += num
		}

		if maxSum < nowSum {
			maxSum = nowSum
		}

	}
	return maxSum
}

func maxSubArraySolutionII(nums []int) (maxSum int) {
	if len(nums) == 0 {
		// do we really need to consider 0 as largest sum of Nil array?
		return 0
	}
	nowSum := nums[0]
	for _, num := range nums[1:] {
		if nowSum < 0 {
			nowSum = num
		} else {
			nowSum += num
		}

		if maxSum < nowSum {
			maxSum = nowSum
		}

	}
	return maxSum
}

// test case
func main() {
	result := maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if reflect.DeepEqual(result, 6) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	// Follow up case
	result = maxSubArraySolutionII([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	if reflect.DeepEqual(result, 6) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
