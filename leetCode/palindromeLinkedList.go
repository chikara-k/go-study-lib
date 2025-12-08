/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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