package main

import (
	"fmt"
	"reflect"
)

/*
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly.
DO NOT allocate another 2D matrix and do the rotation.

Example 1:
	Given input matrix =
	[
	  [1,2,3],
	  [4,5,6],
	  [7,8,9]
	],

	rotate the input matrix in-place such that it becomes:
	[
	  [7,4,1],
	  [8,5,2],
	  [9,6,3]
	]

Example 2:
	Given input matrix =
	[
	  [ 5, 1, 9,11],
	  [ 2, 4, 8,10],
	  [13, 3, 6, 7],
	  [15,14,12,16]
	],

	rotate the input matrix in-place such that it becomes:
	[
	  [15,13, 2, 5],
	  [14, 3, 4, 1],
	  [12, 6, 8, 9],
	  [16, 7,10,11]
	]
*/

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	if len(matrix) != len(matrix[0]) {
		panic("illegal matrix")
	}

	matrixWidth := len(matrix)
	for layer := 0; layer <= (matrixWidth-1)/2; layer++ {
		leftBorder := layer
		rightBorder := matrixWidth - 1 - layer
		upBorder := layer
		downBorder := matrixWidth - 1 - layer

		// single node
		if leftBorder == rightBorder {
			continue
		}

		for i := 0; i < (rightBorder - leftBorder); i++ {
			tempLeftUp := matrix[upBorder][leftBorder+i]

			// leftTop = leftDown
			matrix[upBorder][leftBorder+i] = matrix[downBorder-i][leftBorder]

			// leftDown = rightDown
			matrix[downBorder-i][leftBorder] = matrix[downBorder][rightBorder-i]

			// rightDown = rightUp
			matrix[downBorder][rightBorder-i] = matrix[upBorder+i][rightBorder]

			// rightUp = leftUp
			matrix[upBorder+i][rightBorder] = tempLeftUp
		}
	}

}

// test case
func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	if reflect.DeepEqual(matrix, [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", matrix))
	}

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	if reflect.DeepEqual(matrix, [][]int{
		{15, 13, 2, 5},
		{14, 3, 4, 1},
		{12, 6, 8, 9},
		{16, 7, 10, 11},
	}) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", matrix))
	}
}
