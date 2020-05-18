package main

import (
	"fmt"
)

/*
Given n non-negative integers a1, a2, ..., an
where each represents a point at coordinate (i, ai)
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0)
Find two lines, which together with x-axis forms a container
such that the container contains the most water.

Note: You may not slant the container and n is at least 2.
*/

func maxArea(height []int) (maxContainer int) {
	left, right := 0, len(height)-1
	for ; left < right; {
		currentHeight := height[left]
		if height[right] <= currentHeight {
			currentHeight = height[right]
		}

		if maxContainer < (right-left)*currentHeight {
			maxContainer = (right - left) * currentHeight
		}

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return
}

// test case
func main() {
	result := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	if result != 49 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	fmt.Println("test case passed!")
}
