package main

import (
	"fmt"
)

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:
	Input: nums = [4,5,6,7,0,1,2], target = 0
	Output: 4

Example 2:
	Input: nums = [4,5,6,7,0,1,2], target = 3
	Output: -1
*/

const SearchNotFound = -1

func search(nums []int, target int) int {
	// edge case
	if len(nums) == 0 {
		return SearchNotFound
	}

	for left, right := 0, len(nums)-1; left <= right; {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		// mid in left ascending part
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// nums[left] > nums[mid] -> mid in right ascending part
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return SearchNotFound
}

// test case
func main() {
	result := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	if result != 4 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	if result != -1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = search([]int{1, 2, 3, 4}, 3)
	if result != 2 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = search([]int{1, 2, 3, 4}, 0)
	if result != -1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = search([]int{1}, 0)
	if result != -1 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = search([]int{1}, 1)
	if result != 0 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	fmt.Println("test case passed!")
}
