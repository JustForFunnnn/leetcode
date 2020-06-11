package main

import (
	"fmt"
	"sort"
	"strings"
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

const TotalLowercaseCharCount = 26

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

func groupAnagramsSolution2(strs []string) (anagrams [][]string) {
	anagramsMap := map[string][]string{}
	for _, str := range strs {
		charCount := make([]int, TotalLowercaseCharCount, TotalLowercaseCharCount)
		for _, char := range []byte(str) {
			charCount[char-'a'] += 1
		}
		key := listJoin(charCount, "|")
		anagramsMap[key] = append(anagramsMap[key], str)
	}
	for _, anagram := range anagramsMap {
		anagrams = append(anagrams, anagram)
	}
	return anagrams
}

func listJoin(list []int, spe string) string {
	s := strings.Builder{}
	for _, item := range list {
		s.WriteString(spe)
		s.WriteString(fmt.Sprintf("%d", item))
	}
	return s.String()
}

// test case
func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("result: %+v\n", result)

	result = groupAnagramsSolution2([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("result: %+v\n", result)
}
