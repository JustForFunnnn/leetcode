package main

import (
	"fmt"
)

/*
Given a string containing just the characters '(' and ')',
find the length of the longest valid (well-formed) parentheses substring.

Example 1:
	Input:
		"(()"
	Output:
		2
	Explanation: The longest valid parentheses substring is "()"

Example 2:
	Input:
		")()())"
	Output:
		4
	Explanation: The longest valid parentheses substring is "()()"
*/

func longestValidParentheses(s string) (maxLen int) {
	if len(s) == 0 {
		return 0
	}

	stack := make([]int, len(s), len(s))
	stackLen := 0

	for i, char := range s {
		if char == ')' && stackLen != 0 && s[stack[stackLen-1]] == '(' {
			// pop '('
			stackLen -= 1

			lastStringEnd := stackLen - 1
			if lastStringEnd >= 0 {
				if i-stack[lastStringEnd] > maxLen {
					maxLen = i - stack[lastStringEnd]
				}
			} else {
				if i+1 > maxLen {
					maxLen = i + 1
				}
			}
			continue
		}

		stack[stackLen] = i
		stackLen++
	}

	return maxLen
}

// solution 2
//func longestValidParentheses(s string) (maxLen int) {
//	if len(s) == 0 {
//		return 0
//	}
//	// count of '(' for substring
//	leftCount := 0
//	// count of ')' for substring
//	rightCount := 0
//
//	for nowIndex := 0; nowIndex < len(s); nowIndex++ {
//		if s[nowIndex] == '(' {
//			leftCount++
//		}
//		if s[nowIndex] == ')' {
//			rightCount++
//
//			// illegal case
//			if leftCount < rightCount {
//				// find another '(' to start a new substring
//				for ; nowIndex < len(s) && s[nowIndex] == ')'; nowIndex++ {
//				}
//				leftCount = 1
//				rightCount = 0
//				continue
//			}
//
//			// full matched case
//			if leftCount == rightCount && 2*rightCount > maxLen {
//				maxLen = 2 * rightCount
//				continue
//			}
//
//			// partial matched case
//			if leftCount > rightCount {
//				parenthesesLen := findValidParenthesesLenBackToFrontByIndex(s, nowIndex)
//				if parenthesesLen > maxLen {
//					maxLen = parenthesesLen
//				}
//			}
//
//		}
//	}
//
//	return maxLen
//}
//
//func findValidParenthesesLenBackToFrontByIndex(s string, index int) (parenthesesLen int) {
//	leftCount := 0
//	rightCount := 0
//	for ; 0 <= index; index-- {
//		if s[index] == '(' {
//			leftCount++
//			if leftCount > rightCount {
//				return
//			}
//			if leftCount == rightCount {
//				parenthesesLen = leftCount * 2
//			}
//		}
//		if s[index] == ')' {
//			rightCount++
//		}
//	}
//	return parenthesesLen
//}

// test case
func main() {
	result := longestValidParentheses("(()")
	if result != 2 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = longestValidParentheses(")()())")
	if result != 4 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = longestValidParentheses("")
	if result != 0 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = longestValidParentheses(")(")
	if result != 0 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	fmt.Println("test case passed!")
}
