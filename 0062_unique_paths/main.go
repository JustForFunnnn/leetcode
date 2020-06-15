package main

import (
	"fmt"
	"reflect"
)

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time.

The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Example 1:
	Input: m = 3, n = 2
	Output: 3
	Explanation:
	From the top-left corner, there are a total of 3 ways to reach the bottom-right corner:
		1. Right -> Right -> Down
		2. Right -> Down -> Right
		3. Down -> Right -> Right

Example 2:
	Input: m = 7, n = 3
	Output: 28

Constraints:
	1 <= m, n <= 100
	It's guaranteed that the answer will be less than or equal to 2 * 10 ^ 9.

*/

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	toEndPointWays := make([][]int, m, m)
	for i, _ := range toEndPointWays {
		toEndPointWays[i] = make([]int, n, n)
	}

	return toEndPoint(0, 0, m-1, n-1, toEndPointWays)
}

func toEndPoint(nowX, nowY, endPointX, endPointY int, toEndPointWays [][]int) (ways int) {
	// illegal way
	if nowX > endPointX {
		return 0
	}
	// illegal way
	if nowY > endPointY {
		return 0
	}
	// cached way
	if toEndPointWays[nowX][nowY] != 0 {
		return toEndPointWays[nowX][nowY]
	}
	// end point
	if nowX == endPointX && nowY == endPointY {
		toEndPointWays[nowX][nowY] = 1
		return toEndPointWays[nowX][nowY]
	}

	// way to down
	toEndPointWays[nowX][nowY] += toEndPoint(nowX+1, nowY, endPointX, endPointY, toEndPointWays)
	// way to right
	toEndPointWays[nowX][nowY] += toEndPoint(nowX, nowY+1, endPointX, endPointY, toEndPointWays)

	return toEndPointWays[nowX][nowY]
}

// test case
func main() {
	result := uniquePaths(3, 2)
	if reflect.DeepEqual(result, 3) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = uniquePaths(7, 3)
	if reflect.DeepEqual(result, 28) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
