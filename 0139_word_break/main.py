# Given a non-empty string s and a dictionary wordDict containing a list of non-empty words
# determine if s can be segmented into a space-separated sequence of one or more dictionary words.
#
# Note:
# The same word in the dictionary may be reused multiple times in the segmentation.
# You may assume the dictionary does not contain duplicate words.
#
# Example 1:
# Input: s = "leetcode", wordDict = ["leet", "code"]
# Output: true
# Explanation: Return true because "leetcode" can be segmented as "leet code".
#
# Example 2:
# Input: s = "applepenapple", wordDict = ["apple", "pen"]
# Output: true
# Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
#              Note that you are allowed to reuse a dictionary word.
#
# Example 3:
# Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
# Output: false


class Solution(object):
    def wordBreak(self, s, word_list):
        """
        :type s: str
        :type word_list: List[str]
        :rtype: bool
        """
        return self.word_break_helper(s, {word: True for word in word_list}, {})

    def word_break_helper(self, s, word_dict, cache):
        if s == "":
            return True

        if s in cache:
            return cache[s]

        for i in range(len(s)):
            if s[:i+1] in word_dict and self.word_break_helper(s[i+1:], word_dict, cache) is True:
                cache[s] = True
                return True

        cache[s] = False
        return False


if __name__ == "__main__":
    s = Solution()
    assert s.wordBreak("leetcode",  ["leet", "code"]) == True
    assert s.wordBreak("applepenapple",  ["apple", "pen"]) == True
    assert s.wordBreak("catsandog",  ["cats", "dog", "sand", "and", "cat"]) == False
    assert s.wordBreak("catsandog",  []) == False
    assert s.wordBreak("",  []) == True
    assert s.wordBreak("",  ["leet", "code"]) == True
