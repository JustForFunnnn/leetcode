package main

import (
	"fmt"
)

/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:
s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
*/

func isMatch(s string, p string) bool {
	// assume p is validated patter, otherwise we need to add a check
	// root case
	if p == "" {
		return s == ""
	}
	if s == "" {
		return canBeEmptyPattern(p)
	}

	if len(p) > 1 && p[1] == '*' {
		// with *
		if p[0] != '.' && s[0] != p[0] {
			return isMatch(s, p[2:])
		}
		return isMatch(s[1:], p) || isMatch(s, p[2:])
	} else {
		// without *
		if s[0] != p[0] && p[0] != '.' {
			return false
		}

		return isMatch(s[1:], p[1:])
	}
}

func canBeEmptyPattern(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	for index, char := range s {
		if index%2 == 1 && char != '*' {
			return false
		}
	}
	return true
}

// test case
func main() {
	// "a" does not match the entire string "aa".
	result := isMatch("aa", "a")
	if result != false {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	// '*' means zero or more of the preceding element, 'a'.
	// Therefore, by repeating 'a' once, it becomes "aa".
	result = isMatch("aa", "a*")
	if result != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	// ".*" means "zero or more (*) of any character (.)".
	result = isMatch("ab", ".*")
	if result != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	// c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
	result = isMatch("aab", "c*a*b")
	if result != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	result = isMatch("mississippi", "mis*is*p*.")
	if result != false {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
	fmt.Println("test case passed!")
}
