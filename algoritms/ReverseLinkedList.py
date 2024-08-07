class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def reverseList(self, head: ListNode) -> ListNode:
        if not head:
            return 
        temp = head
        tail = ListNode(val=temp.val)
        while temp.next:
            temp = temp.next
            tail = ListNode(val=temp.val, next=tail)
        return tail
