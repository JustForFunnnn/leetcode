package main

import (
	"fmt"
	"sort"
)

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
Find all unique triplets in the array which gives the sum of zero.

Note: The solution set must not contain duplicate triplets.

Example:
	Given array nums = [-1, 0, 1, 2, -1, -4],

	A solution set is:
	[
	  [-1, 0, 1],
	  [-1, -1, 2]
	]
*/

func threeSum(nums []int) (results [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)

	for i, num := range nums[:len(nums)-2] {
		if num > 0 {
			break
		}
		if i > 0 && nums[i-1] == num {
			continue
		}

		left, right := i+1, len(nums)-1
		for ; left < right; {
			sum := num + nums[left] + nums[right]

			if sum == 0 {
				results = append(results, []int{num, nums[left], nums[right]})
				for left++; left < right && nums[left] == nums[left-1]; left++ {
				}
				for right--; left < right && nums[right] == nums[right+1]; right-- {
				}
			} else if sum > 0 {
				right--
			} else {
				// sum < 0
				left++
			}
		}
	}

	return
}

/* can not use this solution, is too expense to remove the duplicated result
func threeSum(nums []int) (results [][]int) {
	if len(nums) < 3 {
		return
	}
	numTimesMap := make(map[int]int)

	results = make([][]int, 0, 10)
	for i, num := range nums[:len(nums)-2] {
		numTimesMap[num] += 1
		if numTimesMap[num] > 1 {
			continue
		}

		twoSumResults := twoSum(nums, i+1, 0-num)
		if len(twoSumResults) != 0 {
			for _, twoSumResult := range twoSumResults {
				twoSumResult = append(twoSumResult, num)
				results = append(results, twoSumResult)
			}
		}

	}

	return
}

func twoSum(nums []int, startIndex int, targetSum int) (result [][]int) {
	numIndexMap := make(map[int]int, len(nums))
	for i := startIndex; i < len(nums); i++ {
		numIndexMap[nums[i]] = i
	}

	for i := startIndex; i < len(nums); i++ {
		num := nums[i]
		restNum := targetSum - num
		if index, ok := numIndexMap[restNum]; ok && i < index {
			result = append(result, []int{num, restNum})
		}
	}

	return
}
*/

// test case
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
