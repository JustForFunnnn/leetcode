package main

import "fmt"

/*
Given a string containing just the characters "(", ")", "{", "}", "[" and "]", determine if the input string is valid.

An input string is valid if:
	1. Open brackets must be closed by the same type of brackets.
	2. Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:
	Input: "()"
	Output: true
Example 2:
	Input: "()[]{}"
	Output: true
Example 3:
	Input: "(]"
	Output: false
Example 4:
	Input: "([)]"
	Output: false
Example 5:
	Input: "{[]}"
	Output: true
*/

var BracketPairMap = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	bracketsStack := make([]rune, len(s), len(s))
	stackLen := 0

	for _, bracket := range s {
		switch bracket {
		case '(', '[', '{':
			bracketsStack[stackLen] = bracket
			stackLen += 1
		case ')', ']', '}':
			if stackLen == 0 || bracketsStack[stackLen-1] != BracketPairMap[bracket] {
				return false
			}
			stackLen -= 1
		default:
			panic(fmt.Sprintf("unexecpted character: %c", bracket))
		}
	}

	return stackLen == 0
}

// test case
func main() {
	result := isValid("()")
	if result != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = isValid("()[]{}")
	if result != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = isValid("(]")
	if result != false {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = isValid("([)]")
	if result != false {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = isValid("{[]}")
	if result != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	fmt.Println("test case passed!")
}
