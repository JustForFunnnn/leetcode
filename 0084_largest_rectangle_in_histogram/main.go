package main

import (
	"fmt"
	"reflect"
)

/*
Given n non-negative integers representing the histogram's bar height where the width of each bar is 1

find the area of largest rectangle in the histogram.
*/

func largestRectangleArea(heights []int) (maxRectangle int) {
	if len(heights) == 0 {
		return 0
	}

	repeatedHeights := make([]bool, len(heights))

	n := len(heights)
	for i, nowHeight := range heights {
		if repeatedHeights[i] == true {
			continue
		}

		right := i + 1
		for ; right < n && nowHeight <= heights[right]; right++ {
			if heights[right] == nowHeight {
				repeatedHeights[right] = true
			}
		}

		left := i - 1
		for ; 0 <= left && nowHeight <= heights[left]; left-- {
		}

		rectangle := (right - left - 1) * nowHeight
		if rectangle > maxRectangle {
			maxRectangle = rectangle
		}
	}
	return maxRectangle
}

// solution II
func largestRectangleAreaII(heights []int) (maxRectangle int) {
	if len(heights) == 0 {
		return 0
	}

	heights = append(heights, 0)

	stack := make([]int, len(heights))
	stackLen := 0

	for i, height := range heights {

		for ; stackLen != 0 && heights[stack[stackLen-1]] > height; {
			currentHeight := heights[stack[stackLen-1]]
			stackLen--

			rightBound := i
			leftBound := -1
			if stackLen > 0 {
				leftBound = stack[stackLen-1]
			}
			rectangle := currentHeight * (rightBound - leftBound - 1)
			if rectangle > maxRectangle {
				maxRectangle = rectangle
			}
		}

		stack[stackLen] = i
		stackLen++
	}

	return maxRectangle
}

// test case
func main() {
	result := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	if reflect.DeepEqual(result, 10) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = largestRectangleAreaII([]int{2, 1, 5, 6, 2, 3})
	if reflect.DeepEqual(result, 10) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

}
