package main

import (
	"fmt"
)

/*
Given a collection of distinct integers, return all possible permutations.

Example:
	Input: [1,2,3]
	Output:
	[
	  [1,2,3],
	  [1,3,2],
	  [2,1,3],
	  [2,3,1],
	  [3,1,2],
	  [3,2,1]
	]
*/

func permute(nums []int) (result [][]int) {
	result = make([][]int, 0)
	if len(nums) == 0 {
		return result
	}

	distributeNum(nums, []int{}, &result)

	return result
}

func distributeNum(nums []int, permutation []int, result *[][]int) {
	if len(nums) == 0 && len(permutation) != 0 {
		tempPermutation := make([]int, len(permutation))
		copy(tempPermutation, permutation)
		*result = append(*result, tempPermutation)
		return
	}

	for i, num := range nums {
		permutation = append(permutation, num)
		newNums := make([]int, 0, len(nums) - 1)
		newNums = append(newNums, nums[:i]...)
		newNums = append(newNums, nums[i+1:]...)
		distributeNum(newNums, permutation, result)
		permutation = permutation[:len(permutation)-1]
	}
}

// test case
func main() {
	result := permute([]int{1, 2, 3})
	fmt.Println(result)

	result = permute([]int{1})
	fmt.Println(result)
}
