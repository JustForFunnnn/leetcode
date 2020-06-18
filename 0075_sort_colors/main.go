package main

import (
	"fmt"
	"reflect"
)

/*
Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example:
	Input: [2,0,2,1,1,0]
	Output: [0,0,1,1,2,2]
*/

const (
	Red   = 0
	White = 1
	BLUE  = 2
)

func sortColors(nums []int) {
	redIdx, blueIdx := 0, len(nums)-1
	for i := 0; i <= blueIdx; {
		switch nums[i] {
		case Red:
			nums[redIdx], nums[i] = nums[i], nums[redIdx]
			redIdx++
			i++
		case BLUE:
			nums[blueIdx], nums[i] = nums[i], nums[blueIdx]
			blueIdx--
		case White:
			i++
		default:
			panic("illegal color")
		}
	}
}

// test case
func main() {
	colors := []int{2, 0, 2, 1, 1, 0}
	sortColors(colors)
	if reflect.DeepEqual(colors, []int{0, 0, 1, 1, 2, 2}) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", colors))
	}
}
