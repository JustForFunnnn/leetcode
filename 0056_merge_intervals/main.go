package main

import (
	"fmt"
	"reflect"
	"sort"
)

/*
Given a collection of intervals, merge all overlapping intervals.

Example 1:
	Input: [[1,3],[2,6],[8,10],[15,18]]
	Output: [[1,6],[8,10],[15,18]]
	Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].

Example 2:
	Input: [[1,4],[4,5]]
	Output: [[1,5]]
	Explanation: Intervals [1,4] and [4,5] are considered overlapping.

*/

func merge(intervals [][]int) (result [][]int) {
	if len(intervals) == 0 {
		return nil
	}
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] <= intervals[j][0] })

	for _, interval := range intervals {
		if len(result) != 0 {
			lastStackIndex := len(result) - 1
			// overlap
			if interval[0] <= result[lastStackIndex][1] {
				newLeft := result[lastStackIndex][0]
				newRight := interval[1]
				if newRight <= result[lastStackIndex][1] {
					newRight = result[lastStackIndex][1]
				}
				result[lastStackIndex] = []int{newLeft, newRight}
				continue
			}
		}

		result = append(result, interval)

	}
	return result
}

// test case
func main() {
	result := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	if reflect.DeepEqual(result, [][]int{{1, 6}, {8, 10}, {15, 18}}) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = merge([][]int{{1, 4}, {4, 5}})
	if reflect.DeepEqual(result, [][]int{{1, 5}}) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
