package main

import (
	"fmt"
)

/*
Given a set of candidate numbers (candidates) (without duplicates) and a target number (target),

find all unique combinations in candidates where the candidate numbers sums to target.

The same repeated number may be chosen from candidates unlimited number of times.

Note:

All numbers (including target) will be positive integers.
The solution set must not contain duplicate combinations.

Example 1:
	Input: candidates = [2,3,6,7], target = 7,
	A solution set is:
	[
	  [7],
	  [2,2,3]
	]

Example 2:
	Input: candidates = [2,3,5], target = 8,
	A solution set is:
	[
	  [2,2,2,2],
	  [2,3,3],
	  [3,5]
	]
*/
func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	distributeCandidates([]int{}, target, candidates, &result)
	return result
}

func distributeCandidates(nowCombination []int, target int, candidates []int, result *[][]int) {
	// illegal parenthesis
	nowSum := SumCombination(nowCombination)
	if nowSum > target {
		return
	}

	if nowSum == target {
		temp := make([]int, len(nowCombination))
		copy(temp, nowCombination)
		*result = append(*result, temp)
		return
	}

	for index, candidate := range candidates {
		if nowSum+candidate <= target {
			nowCombination = append(nowCombination, candidate)
			distributeCandidates(nowCombination, target, candidates[index:], result)
			nowCombination = nowCombination[:len(nowCombination)-1]
		}
	}
}

func SumCombination(nowCombination []int) (sum int) {
	for _, num := range nowCombination {
		sum += num
	}
	return sum
}

// test case
func main() {
	result := combinationSum([]int{2, 3, 6, 7}, 7)
	fmt.Println("[]int{2, 3, 6, 7}, 7 -> ", result)

	result = combinationSum([]int{2, 3, 5}, 8)
	fmt.Println("[]int{2, 3, 5}, 8 -> ", result)
}
