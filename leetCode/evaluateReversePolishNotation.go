package leetcode

import "strconv"

func evalRPN(tokens []string) int {
    stack := []int{}

    for _, t := range tokens {
        switch t {
        case "+", "-", "*", "/":
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]

            var res int
            switch t {
            case "+":
                res = a + b
            case "-":
                res = a - b
            case "*":
                res = a * b
            case "/":
                res = a / b 
            }

            stack = append(stack, res)

        default:
            num, _ := strconv.Atoi(t)
            stack = append(stack, num)
        }
    }

    return stack[0]
}