package main

import (
	"fmt"
	"reflect"
)

/*
Given an unsorted array of integers, find the length of the longest consecutive elements sequence.

Your algorithm should run in O(n) complexity.

Example:
	Input: [100, 4, 200, 1, 3, 2]
	Output: 4
	Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
*/

func longestConsecutive(nums []int) (maxCons int) {
	if len(nums) == 0 {
		return 0
	}

	seqMap := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		seqMap[num] = struct{}{}
	}

	for num, _ := range seqMap {
		left := num - 1
		right := num + 1

		for {
			_, ok := seqMap[left]
			if !ok {
				break
			}
			delete(seqMap, left)
			left--
		}

		for {
			_, ok := seqMap[right]
			if !ok {
				break
			}
			delete(seqMap, right)
			right++
		}

		if (right-left)-1 > maxCons {
			maxCons = (right - left) - 1
		}
		delete(seqMap, num)
	}

	return maxCons
}

// test case
func main() {
	result := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	if reflect.DeepEqual(result, 4) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = longestConsecutive([]int{100, 1})
	if reflect.DeepEqual(result, 1) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
