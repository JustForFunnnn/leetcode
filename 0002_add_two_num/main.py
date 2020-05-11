"""
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
"""

# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        flag = 0
        result_head = ListNode(-1)
        now_node = result_head
        while l1 or l2 or flag:
            new_node = ListNode(flag)
            if l1:
                new_node.val += l1.val
                l1 = l1.next

            if l2:
                new_node.val += l2.val
                l2 = l2.next

            flag = new_node.val // 10
            new_node.val %= 10
            now_node.next = new_node
            now_node = now_node.next

        return result_head.next
