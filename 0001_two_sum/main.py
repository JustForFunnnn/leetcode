"""
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
"""

from typing import List

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        num_index_map = {}
        for index, num in enumerate(nums):
            num_index_map[num] = index

        for index, num in enumerate(nums):
            rest_num = target - num
            if rest_num in num_index_map and num_index_map[rest_num] > index:
                return [index, num_index_map[rest_num]]

        return None


solution = Solution()
assert solution.twoSum(nums=[2, 7, 11, 15], target=9) == [0, 1]
