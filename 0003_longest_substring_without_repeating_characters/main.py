"""
Given a string, find the length of the longest substring without repeating characters.

Example 1:
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.

Example 2:
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.

Example 3:
Input: "pwwkew"
Output: 3
Explanation:
    The answer is "wke", with the length of 3.
    Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
"""

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        char_index_map = {}
        substring_strat_index = 0
        max_substring_len = 0

        for index, char in enumerate(list(s)):
            if char in char_index_map and char_index_map[char] >= substring_strat_index:
                substring_strat_index = char_index_map[char] + 1
            max_substring_len = max(index - substring_strat_index + 1, max_substring_len)
            char_index_map[char] = index

        return max_substring_len


solution = Solution()
assert solution.lengthOfLongestSubstring("abcabcbb") == 3
assert solution.lengthOfLongestSubstring("bbbbb") == 1
assert solution.lengthOfLongestSubstring("pwwkew") == 3