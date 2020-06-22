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

// test case
func main() {
	result := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	if reflect.DeepEqual(result, 10) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

}
