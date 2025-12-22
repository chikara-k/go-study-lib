package leetcode

func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    seen := make(map[int]bool)
    for _, n := range nums {
        seen[n] = true
    }

    maxLen := 0
    for n := range seen {
        if !seen[n-1] {
            cur := n
            length := 1
            for seen[cur+1] {
                cur++
                length++
            }
            if length > maxLen{
                maxLen = length
            }
        }
    }

    return maxLen
}