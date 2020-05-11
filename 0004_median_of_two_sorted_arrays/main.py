"""
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:
nums1 = [1, 3]
nums2 = [2]
The median is 2.0

Example 2:
nums1 = [1, 2]
nums2 = [3, 4]
The median is (2 + 3)/2 = 2.5
"""

from typing import List

class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        if not nums1 and not nums2:
            return None

        nums1_index, nums2_index = 0, 0
        nums1_len, nums2_len = len(nums1), len(nums2)
        new_array = []
        while nums1_index < nums1_len or nums2_index < nums2_len:
            if nums2_index >= nums2_len or (nums1_index < nums1_len and nums1[nums1_index] <= nums2[nums2_index]):
                new_array.append(nums1[nums1_index])
                nums1_index += 1
            else:
                new_array.append(nums2[nums2_index])
                nums2_index += 1

        median = new_array[(nums1_len + nums2_len) // 2]
        if (nums1_len + nums2_len) % 2 == 0:
            median += new_array[(nums1_len + nums2_len) // 2 - 1]
            median /= 2.0

        return median



solution = Solution()
assert solution.findMedianSortedArrays([1, 3], [2, ]) == 2.0
assert solution.findMedianSortedArrays([1, 2], [3, 4]) == 2.5
