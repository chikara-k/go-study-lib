package leetcode

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    anamap := make(map[rune]int)

    for _, char := range s {
        anamap[char]++
    }

    for _, char := range t {
        anamap[char]--
        if anamap[char] < 0 {
            return false
        }
    }

    return true
}