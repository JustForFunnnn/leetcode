package main

import (
	"fmt"
	"reflect"
)

/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:
	Input:
	[
	  [1,3,1],
	  [1,5,1],
	  [4,2,1]
	]
	Output: 7
	Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

const UnReach = -1

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		// illegal grid
		return UnReach
	}

	minCost := make([][]int, len(grid), len(grid))
	for i, _ := range minCost {
		minCost[i] = make([]int, len(grid[0]), len(grid[0]))
		for j := range minCost[i] {
			minCost[i][j] = UnReach
		}
	}
	return toEndPointMinCost(0, 0, grid, minCost)
}

func toEndPointMinCost(nowX, nowY int, grid [][]int, minCost [][]int) (cost int) {
	// cached way
	if minCost[nowX][nowY] >= 0 {
		return minCost[nowX][nowY]
	}

	maxX := len(grid) - 1
	maxY := len(grid[0]) - 1

	// end point
	if nowX == maxX && nowY == maxY {
		minCost[nowX][nowY] = grid[nowX][nowY]
		return minCost[nowX][nowY]
	}

	// way to down
	if nowX < maxX {
		downWayCost := toEndPointMinCost(nowX+1, nowY, grid, minCost)
		minCost[nowX][nowY] = grid[nowX][nowY] + downWayCost
		if downWayCost == 1 {
			return minCost[nowX][nowY]
		}
	}

	// way to right
	if nowY < maxY {
		rightWayCost := toEndPointMinCost(nowX, nowY+1, grid, minCost)
		if grid[nowX][nowY]+rightWayCost < minCost[nowX][nowY] || minCost[nowX][nowY] == UnReach {
			minCost[nowX][nowY] = grid[nowX][nowY] + rightWayCost
		}
	}

	return minCost[nowX][nowY]
}

// test case
func main() {
	result := minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	})
	if reflect.DeepEqual(result, 7) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	result = minPathSum([][]int{
		{9,9,0,8,9,0,5,7,2,2,7,0,8,0,2,4,8},{4,4,2,7,6,0,9,7,3,2,5,4,6,5,4,8,7},{4,9,7,0,7,9,2,4,0,2,4,4,6,2,8,0,7},{7,7,9,6,6,4,8,4,8,7,9,4,7,6,9,6,5},{1,3,7,5,7,9,7,3,3,3,8,3,6,5,0,3,6},{7,1,0,7,5,0,6,6,5,3,2,6,0,0,9,5,7},{6,5,6,3,8,1,8,6,4,4,3,4,9,9,3,3,1},{1,0,2,9,7,9,3,1,7,5,1,8,2,8,4,7,6},{9,6,7,7,4,1,4,0,6,5,1,9,0,3,2,1,7},{2,0,8,7,1,7,4,3,5,6,1,9,4,0,0,2,7},{9,8,1,3,8,7,1,2,8,3,7,3,4,6,7,6,6},{4,8,3,8,1,0,4,4,1,0,4,1,4,4,0,3,5},{6,3,4,7,5,4,2,2,7,9,8,4,5,6,0,3,9},{0,4,9,7,1,0,7,7,3,2,1,4,7,6,0,0,0},
	})
	if reflect.DeepEqual(result, 77) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
