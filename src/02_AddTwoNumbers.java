/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode(int x) { val = x; }
 * }
 */
class Solution {
    
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        // if the numbers are the same size
        int carry = 0, sum = 0;
        ListNode result = new ListNode(0);
        ListNode res_ptr = result;
        int l1val = 0, l2val = 0;
        while (l1 != null || l2 != null){
            l1val = l1 == null ? 0 : l1.val;
            l2val = l2 == null ? 0 : l2.val;
            sum = l1val + l2val + carry;
            carry = sum / 10;
            sum = sum % 10;
            result.next = new ListNode(sum);
            if (l1 != null) l1 = l1.next;
            if (l2 != null) l2 = l2.next;
            result = result.next;
        }

        if (carry > 0){
            result.next = new ListNode(carry); //10  
        }
        
        return res_ptr.next;
    }
}