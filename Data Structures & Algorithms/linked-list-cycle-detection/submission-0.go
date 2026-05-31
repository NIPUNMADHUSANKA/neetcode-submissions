/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    result := make(map[*ListNode]bool)

	curr := head

	for curr != nil{
		if _, ok := result[curr.Next]; ok{
			return true
		}
		result[curr.Next] = true
		curr = curr.Next
	}
	return false
}
