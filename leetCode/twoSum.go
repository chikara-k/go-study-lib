package leetcode

func twoSum(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        firstNum := nums[i]
        for a := 0;  a < len(nums); a++ {
            if i == a {
                continue
            }
            secondNum := nums[a]
            if target == (firstNum + secondNum) {
                return []int{i, a}
            }
        }
    }

    return []int{}
}