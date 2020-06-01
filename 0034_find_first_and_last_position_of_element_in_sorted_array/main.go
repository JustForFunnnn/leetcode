package main

import (
	"fmt"
)

/*
Given an array of integers nums sorted in ascending order,
find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:
	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]

Example 2:
	Input: nums = [5,7,7,8,8,10], target = 6
	Output: [-1,-1]
*/
const SearchNotFound = -1

type SearchType int

const
(
	SearchLeftest SearchType = iota
	SearchRightest
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{SearchNotFound, SearchNotFound}
	}
	return []int{
		binSearch(nums, target, 0, len(nums)-1, SearchLeftest),
		binSearch(nums, target, 0, len(nums)-1, SearchRightest),
	}
}

func binSearch(nums []int, target, left, right int, searchType SearchType) int {
	if left > right {
		return SearchNotFound
	}

	mid := (left + right) / 2
	if nums[mid] == target {
		// the leftest matched element
		if searchType == SearchLeftest && (left == mid || nums[mid-1] != target) {
			return mid
		}

		// the rightest matched element
		if searchType == SearchRightest && (right == mid || nums[mid+1] != target) {
			return mid
		}
	}

	if nums[mid] == target {
		if searchType == SearchLeftest {
			return binSearch(nums, target, left, mid-1, searchType)
		} else {
			return binSearch(nums, target, mid+1, right, searchType)
		}
	} else if nums[mid] < target {
		return binSearch(nums, target, mid+1, right, searchType)
	} else {
		return binSearch(nums, target, left, mid-1, searchType)
	}
}

// test case
func main() {
	result := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	fmt.Println("[]int{5,7,7,8,8,10}, 8 -> ", result)

	result = searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	fmt.Println("[]int{5,7,7,8,8,10}, 6 -> ", result)
}
