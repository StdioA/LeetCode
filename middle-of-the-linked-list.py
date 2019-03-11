# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None
from made_list import ListNode


class Solution:
    def middleNode(self, head: ListNode) -> ListNode:
        slow = fast = head
        while fast.next and fast.next.next:
            slow = slow.next
            fast = fast.next.next
        if fast.next:
            slow = slow.next
        return slow


s = Solution()
print(s.middleNode(ListNode("1->2->3->4->5")))
print(s.middleNode(ListNode("1->2->3->4->5->6")))
