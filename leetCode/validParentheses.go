package leetcode

func isValid(s string) bool {
    pairs := map[byte]byte{
        ')': '(',
        '}': '{',
        ']': '[',
    }

    stack := make([]byte, 0, len(s))

    for i := 0; i < len(s); i++ {
        ch := s[i]
        if open, isClose := pairs[ch]; isClose {
            if len(stack) == 0 {
                return false
            }

            if stack[len(stack)-1] != open {
                return false
            }

            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, ch)
        }
    }

    return len(stack) == 0
}