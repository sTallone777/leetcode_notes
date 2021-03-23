/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    h := head
    i := h
    for h != nil{
        if h.Next != nil{
            if h.Val != h.Next.Val{
                i.Next = h.Next
                i = i.Next
            }
        }else{
            if i.Val == h.Val {
                i.Next = h.Next
            }
        }
        
        h = h.Next
    }
    return head
}