package main

import (
	"fmt"
	"strings"
)

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:
	[
	  "((()))",
	  "(()())",
	  "(())()",
	  "()(())",
	  "()()()"
	]
*/
func generateParenthesis(n int) (result []string) {
	distributeBracket(0, 0, n, []byte{}, &result)
	return result
}

func distributeBracket(distributedLeft, distributedRight, total int, nowParenthesis []byte, result *[]string) {
	// illegal parenthesis
	if distributedLeft < distributedRight || total < distributedLeft || total < distributedRight {
		return
	}
	if distributedLeft == total && distributedRight == total {
		var strBuild strings.Builder
		strBuild.Write(nowParenthesis)
		*result = append(*result, strBuild.String())
	}

	if distributedLeft < total {
		nowParenthesis = append(nowParenthesis, '(')
		distributeBracket(distributedLeft+1, distributedRight, total, nowParenthesis, result)
		// high performance, cuz nowParenthesis and nowParenthesis[:] share the same memory
		nowParenthesis = nowParenthesis[:len(nowParenthesis)-1]
	}

	if distributedRight < total && distributedRight < distributedLeft {
		nowParenthesis = append(nowParenthesis, ')')
		distributeBracket(distributedLeft, distributedRight+1, total, nowParenthesis, result)
	}
}

// test case
func main() {
	result := generateParenthesis(3)
	fmt.Println(result)
}
