package main

import (
	"fmt"
	"reflect"
)

/*
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell

where "adjacent" cells are those horizontally or vertically neighboring.

The same letter cell may not be used more than once.

Example:
	board =
	[
	  ['A','B','C','E'],
	  ['S','F','C','S'],
	  ['A','D','E','E']
	]

	Given word = "ABCCED", return true.
	Given word = "SEE", return true.
	Given word = "ABCB", return false.


Constraints:
	board and word consists only of lowercase and uppercase English letters.
	1 <= board.length <= 200
	1 <= board[i].length <= 200
	1 <= word.length <= 10^3
*/

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
		return false
	}
	for x := range board {
		for y := range board[x] {
			if containWord(board, x, y, word) {
				return true
			}
		}
	}
	return false
}

func containWord(board [][]byte, nowX, nowY int, word string) bool {
	if len(word) == 0 {
		return true
	}

	maxX := len(board) - 1
	maxY := len(board[0]) - 1

	if maxX < nowX || nowX < 0 {
		return false
	}

	if maxY < nowY || nowY < 0 {
		return false
	}

	if board[nowX][nowY] == '#' {
		return false
	}

	if board[nowX][nowY] != word[0] {
		return false
	}

	temp := board[nowX][nowY]
	board[nowX][nowY] = '#'

	// up
	if containWord(board, nowX-1, nowY, word[1:]) {
		return true
	}

	// down
	if containWord(board, nowX+1, nowY, word[1:]) {
		return true
	}

	// right
	if containWord(board, nowX, nowY+1, word[1:]) {
		return true
	}

	// left
	if containWord(board, nowX, nowY-1, word[1:]) {
		return true
	}

	board[nowX][nowY] = temp

	return false
}

// test case
func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	result := exist(board, "ABCCED")
	if reflect.DeepEqual(result, true) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}


	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	result = exist(board, "SEE")
	if reflect.DeepEqual(result, true) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	board = [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	result = exist(board, "ABCB")
	if reflect.DeepEqual(result, false) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
