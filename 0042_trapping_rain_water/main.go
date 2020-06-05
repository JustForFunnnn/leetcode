package main

import (
	"fmt"
)

/*
Given n non-negative integers representing an elevation map where the width of each bar is 1,

compute how much water it is able to trap after raining.

Example:
	Input: [0,1,0,2,1,0,1,3,2,1,2,1]
	Output: 6
*/
func trap(height []int) (result int) {
	MaxHeightCanTrap := make([]int, len(height), len(height))
	leftHighest := 0
	for i, currentHeight := range height {
		MaxHeightCanTrap[i] = leftHighest
		if currentHeight > leftHighest {
			leftHighest = currentHeight
		}
	}

	rightHighest := 0
	for i := len(height) - 1; i >= 0; i-- {
		currentHeight := height[i]
		if rightHighest < MaxHeightCanTrap[i] {
			MaxHeightCanTrap[i] = rightHighest
		}
		if currentHeight > rightHighest {
			rightHighest = currentHeight
		}
	}

	for i, _ := range height {
		if MaxHeightCanTrap[i]-height[i] > 0 {
			result += MaxHeightCanTrap[i] - height[i]
		}
	}

	return result
}

// test case
func main() {
	result := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	if result != 6 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
