/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func moveNode(dest **ListNode, source **ListNode) {
	newNode := *source
	if newNode != nil {
		*source = newNode.Next
		newNode.Next = *dest
		*dest = newNode
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	var result *ListNode = nil

	lastPtrRef := &result
	for {
		if list1 == nil {
			*lastPtrRef = list2
			break
		} else if list2 == nil {
			*lastPtrRef = list1
			break
		}

		if list1.Val <= list2.Val {
			moveNode(lastPtrRef, &list1)
		} else {
			moveNode(lastPtrRef, &list2)
		}

		lastPtrRef = &((*lastPtrRef).Next)
	}

	return result
}
