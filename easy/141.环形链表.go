/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    m := make(map[*ListNode]int)
    pos := -1
    idx := 0
    h := head
    for h != nil{
        if _, ok := m[h]; ok{
            pos = m[h]
            break
        }else{
            m[h] = idx
        }
        idx++
        h = h.Next
    }
    if pos > -1 {
        return true
    }
    return false
}