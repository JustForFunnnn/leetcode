package main

import (
	"fmt"
	"reflect"
)

/*
Given a 2D binary matrix filled with 0's and 1's, find the largest rectangle containing only 1's and return its area.

Example:

	Input:
	[
	  ["1","0","1","0","0"],
	  ["1","0","1","1","1"],
	  ["1","1","1","1","1"],
	  ["1","0","0","1","0"]
	]
	Output: 6
*/

func maximalRectangle(matrix [][]byte) (maxArea int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0])+1)

	for row := range matrix {
		for j, char := range matrix[row] {
			if char == '0' {
				heights[j] = 0
			} else {
				heights[j] = heights[j] + 1
			}

		}
		nowArea := maxRectangleArea(heights)
		if maxArea < nowArea {
			maxArea = nowArea
		}
	}

	return maxArea
}

func maxRectangleArea(heights []int) (maxArea int) {
	heights = append(heights, 0)

	stock := make([]int, len(heights))
	stockLen := 0
	for i, height := range heights {

		for ; stockLen != 0 && height < heights[stock[stockLen-1]]; {
			rectAngleHeight := heights[stock[stockLen-1]]
			stockLen--
			rightBoundIdx := i
			leftBoundIdx := -1
			if stockLen != 0 {
				leftBoundIdx = stock[stockLen-1]
			}

			nowArea := rectAngleHeight * (rightBoundIdx - leftBoundIdx - 1)
			if maxArea < nowArea {
				maxArea = nowArea
			}
		}

		stock[stockLen] = i
		stockLen++
	}
	return maxArea
}

// test case
func main() {
	result := maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	})
	if reflect.DeepEqual(result, 6) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
