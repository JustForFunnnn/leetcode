package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:
	Input: S = "ADOBECODEBANC", T = "ABC"
	Output: "BANC"

Note:
	If there is no such window in S that covers all characters in T, return the empty string "".
	If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) (result string) {
	if s == "" || t == "" {
		return ""
	}

	sLen, tLen := len(s), len(t)
	left, right := 0, 0

	tCountMap := make(map[rune]int)
	sCountMap := make(map[rune]int)
	for _, char := range t {
		tCountMap[char] ++
	}

	coveredCount := 0
	minWindowLen := math.MaxInt64
	for ; right < sLen; {
		rightChar := rune(s[right])
		sCountMap[rightChar] += 1
		if tCountMap[rightChar] > 0 && sCountMap[rightChar] <= tCountMap[rightChar] {
			coveredCount += 1
		}

		for ; left <= right && coveredCount == tLen; {
			if right-left+1 < minWindowLen {
				minWindowLen = right - left + 1
				result = s[left : right+1]
			}
			leftChar := rune(s[left])
			sCountMap[leftChar] -= 1

			if tCountMap[leftChar] > 0 && sCountMap[leftChar] < tCountMap[leftChar] {
				coveredCount -= 1
			}

			left++
		}

		right++
	}

	return result
}

// test case
func main() {
	result := minWindow("ADOBECODEBANC", "ABC")
	if reflect.DeepEqual(result, "BANC") != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = minWindow("ABC", "ABC")
	if reflect.DeepEqual(result, "ABC") != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = minWindow("ABCD", "ABC")
	if reflect.DeepEqual(result, "ABC") != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = minWindow("BC", "ABC")
	if reflect.DeepEqual(result, "") != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
