package main

import (
	"fmt"
	"sort"
)

/*
Given an array of strings, group anagrams together.

Example:
	Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
	Output:
	[
	  ["ate","eat","tea"],
	  ["nat","tan"],
	  ["bat"]
	]

Note:
	All inputs will be in lowercase.
	The order of your output does not matter.
*/

func groupAnagrams(strs []string) (anagrams [][]string) {
	anagramsMap := map[string][]string{}
	for _, str := range strs {
		charList := []byte(str)
		sort.Slice(charList, func(i int, j int) bool { return charList[i] < charList[j] })
		key := string(charList)
		anagramsMap[key] = append(anagramsMap[key], str)
	}
	for _, anagram := range anagramsMap {
		anagrams = append(anagrams, anagram)
	}
	return anagrams
}

// test case
func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("result: %+v\n", result)

}
