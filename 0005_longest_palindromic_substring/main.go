package main

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

func longestPalindrome(s string) (result string) {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	charList := []byte(s)
	for i := 0; i < sLen; i++ {
		if subString := getLongestPalindromeBaseOnIndex(charList, i, i, sLen); len(subString) > len(result){
			result = subString
		}

		if (i + 1) < sLen{
			if subString := getLongestPalindromeBaseOnIndex(charList, i, i+1, sLen); len(subString) > len(result){
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

	return string(charList[left+1:right])
}

func longestPalindrome1(s string) (result string) {
	sLen := len(s)
	if sLen <= 1 {
		return s
	}

	palindromeCache := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		palindromeCache[i] = make([]bool, sLen)
		palindromeCache[i][i] = true
	}

	charList := []byte(s)
	for i := sLen - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if isPalindrome(charList, i, j, palindromeCache) == true {
				if len(result) <= (j - i + 1) {
					result = string(charList[i : j+1])
				}
			}
		}
	}

	return result
}

func isPalindrome(charList []byte, start, end int, palindromeCache [][]bool) bool {
	if start == end {
		return true
	}

	palindromeCache[start][end] = false
	if charList[start] == charList[end] && palindromeCache[start+1][end-1] == true {
		palindromeCache[start][end] = true
	}

	return palindromeCache[start][end]

}
func main() {
	if !(longestPalindrome("babad") == "bab" || longestPalindrome("babad") == "aba") {
		panic("not expected result")
	}
	if !(longestPalindrome("cbbd") == "bb") {
		panic("not expected result")
	}
}
