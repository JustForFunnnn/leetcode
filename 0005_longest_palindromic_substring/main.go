package main

import "fmt"

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/

// solution 1
func longestPalindrome(s string) (result string) {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	charList := []byte(s)
	for i := 0; i < sLen; i++ {
		if subString := getLongestPalindromeBaseOnIndex(charList, i, i, sLen); len(subString) > len(result) {
			result = subString
		}

		if (i + 1) < sLen {
			if subString := getLongestPalindromeBaseOnIndex(charList, i, i+1, sLen); len(subString) > len(result) {
				result = subString
			}
		}
	}

	return result
}

func getLongestPalindromeBaseOnIndex(charList []byte, left, right, sLen int) string {
	for ; 0 <= left && right < sLen; {
		if charList[left] != charList[right] {
			break
		}

		left -= 1
		right += 1
	}

	return string(charList[left+1 : right])
}

// solution 2
func longestPalindrome2(s string) (result string) {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}
	palindromeCache := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		palindromeCache[i] = make([]bool, sLen)
	}

	charList := []byte(s)
	for i := 0; i < sLen; i++ {
		for j := 0; j <= i; j++ {
			if charList[i] == charList[j] &&
				(i == j || j+1 == i || j+2 == i || palindromeCache[i-1][j+1]) {
				palindromeCache[i][j] = true

				if (i + 1 - j) > len(result) {
					result = string(charList[j : i+1])
				}
				continue
			}

			palindromeCache[i][j] = false
		}
	}
	return result
}

// test case
func main() {
	result := longestPalindrome("babad")
	if result != "aba" && result != "bab" {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	result = longestPalindrome("cbbd")
	if result != "bb" {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	result = longestPalindrome2("babad")
	if result != "aba" && result != "bab" {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	result = longestPalindrome2("cbbd")
	if result != "bb" {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	fmt.Println("test case passed!")
}
