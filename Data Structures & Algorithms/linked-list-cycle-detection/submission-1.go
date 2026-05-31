/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    visited  := make(map[*ListNode]bool)

	curr := head

	for curr != nil{
		if _, ok := visited [curr.Next]; ok{
			return true
		}
		visited [curr] = true
		curr = curr.Next
	}
	return false
}
