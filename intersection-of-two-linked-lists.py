# coding: utf-8

# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

# 从A和B分别开始，求出链表的长度，假设差了N
# 那么再从头开始，先让长的那个走N格
# 然后两端一起推进

class Solution:
    # @param two ListNodes
    # @return the intersected ListNode
    def getIntersectionNode(self, headA, headB):
        lenA = lenB = 0

        temp = headA
        while temp != None:
        	lenA += 1
        	temp = temp.next
        temp = headB
        while temp != None:
        	lenB += 1
        	temp = temp.next

        posA, posB = headA, headB
        if lenA > lenB:
        	minus = lenA - lenB
        	for i in range(minus):
        		posA = posA.next
        elif lenA < lenB:
        	minus = lenB - lenA
        	for i in range(minus):
        		posB = posB.next

        while posA != None:
        	if posA == posB:
        		return posA
        	posA, posB = posA.next, posB.next

        return None

