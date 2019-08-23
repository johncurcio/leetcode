/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func defaultVal(l *ListNode) int {
    if l != nil {
        return l.Val
    }
    return 0
}

func defaultNext(l *ListNode) *ListNode {
    if l != nil {
        return l.Next
    }
    return nil;
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry := 0
    dummy := &ListNode{0, nil}
    result := dummy
    
    for l1 != nil || l2 != nil {
        l1val := defaultVal(l1)
        l2val := defaultVal(l2)
        sum := l1val + l2val + carry
        carry = sum / 10
        dummy.Next = &ListNode{ sum % 10, nil }
        dummy = dummy.Next
        l1 = defaultNext(l1)
        l2 = defaultNext(l2)
    }

    if carry > 0 { 
        dummy.Next = &ListNode{ carry , nil }
    }
    
    return result.Next
}
