package main

import (
	"fmt"
	"reflect"
)

/*
Given two words word1 and word2, find the minimum number of operations required to convert word1 to word2.

You have the following 3 operations permitted on a word:
	Insert a character
	Delete a character
	Replace a character

Example 1:
	Input: word1 = "horse", word2 = "ros"
	Output: 3
	Explanation:
		horse -> rorse (replace 'h' with 'r')
		rorse -> rose (remove 'r')
		rose -> ros (remove 'e')

Example 2:
	Input: word1 = "intention", word2 = "execution"
	Output: 5
	Explanation:
		intention -> inention (remove 't')
		inention -> enention (replace 'i' with 'e')
		enention -> exention (replace 'n' with 'x')
		exention -> exection (replace 'n' with 'c')
		exection -> execution (insert 'u')
*/

const UnKnow = -1

func minDistance(word1 string, word2 string) (min int) {
	mem := make([][]int, len(word1))
	for i := range mem {
		mem[i] = make([]int, len(word2))
		for j := range mem[i] {
			mem[i][j] = UnKnow
		}
	}
	return minDistanceHelper(word1, 0, word2, 0, mem)
}

func minDistanceHelper(word1 string, word1Idx int, word2 string, word2Idx int, mem [][]int) (min int) {
	if len(word1) == word1Idx {
		return len(word2) - word2Idx
	}

	if len(word2) == word2Idx {
		return len(word1) - word1Idx
	}

	if mem[word1Idx][word2Idx] != UnKnow {
		return mem[word1Idx][word2Idx]
	}

	if word1[word1Idx] == word2[word2Idx] {
		mem[word1Idx][word2Idx] = minDistanceHelper(word1, word1Idx+1, word2, word2Idx+1, mem)
		return mem[word1Idx][word2Idx]
	}

	// Insert a character
	min = 1 + minDistanceHelper(word1, word1Idx, word2, word2Idx+1, mem)
	// Delete a character
	nowMin := 1 + minDistanceHelper(word1, word1Idx+1, word2, word2Idx, mem)
	if nowMin < min {
		min = nowMin
	}
	// Replace a character
	nowMin = 1 + minDistanceHelper(word1, word1Idx+1, word2, word2Idx+1, mem)
	if nowMin < min {
		min = nowMin
	}

	mem[word1Idx][word2Idx] = min
	return mem[word1Idx][word2Idx]
}

func minDistanceII(word1 string, word2 string) (min int) {
	// minStep[i][j] means the mini distance from word1[0:i+1] become word2[0:j+1]
	minStep := make([][]int, len(word1)+1)
	for i := range minStep {
		minStep[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		minStep[i][0] = i
	}

	for j := 0; j <= len(word2); j++ {
		minStep[0][j] = j
	}

	for i := 0; i <= len(word1); i++ {
		for j := 0; j <= len(word2); j++ {
			if i == 0 || j == 0 {
				continue
			}
			if word1[i-1] == word2[j-1] {
				minStep[i][j] = minStep[i-1][j-1]
			} else {
				minStep[i][j] = minStep[i-1][j-1] + 1
				if minStep[i][j-1]+1 < minStep[i][j] {
					minStep[i][j] = minStep[i][j-1] + 1
				}
				if minStep[i-1][j]+1 < minStep[i][j] {
					minStep[i][j] = minStep[i-1][j] + 1
				}
			}
		}
	}

	return minStep[len(word1)][len(word2)]
}

// test case
func main() {
	result := minDistance("horse", "ros")
	if reflect.DeepEqual(result, 3) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = minDistance("intention", "execution")
	if reflect.DeepEqual(result, 5) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = minDistanceII("horse", "ros")
	if reflect.DeepEqual(result, 3) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = minDistanceII("intention", "execution")
	if reflect.DeepEqual(result, 5) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
