package leetcode

/**
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
func isPalindrome(head *ListNode) bool {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}

	l, r := 0, len(vals)-1
	for l < r {
		if vals[l] != vals[r] {
			return false
		}
		l++
		r--
	}

	return true
}
*/

// Other Answer
func isPalindrome(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var prev *ListNode
	curr := slow
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev, head = prev.Next, head.Next
	}
	return true
}
