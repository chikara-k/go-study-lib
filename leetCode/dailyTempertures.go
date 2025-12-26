package leetcode

func dailyTemperatures(temperatures []int) []int {
    n := len(temperatures)
    ans := make([]int, n)

    // スタックには「まだ次の暖かい日が見つかっていない index」を積む
    // 温度は stack の上に行くほど小さくなる（単調減少）
    stack := []int{}

    for i := 0; i < n; i++ {
        // 今日の温度が、スタック上の過去の日より暖かいなら、その過去の日は解決できる
        for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
            prev := stack[len(stack)-1]       // 解決する日
            stack = stack[:len(stack)-1]      // pop
            ans[prev] = i - prev              // 何日後に暖かくなるか
        }
        // 今日(i)は未来のどこかで解決されるかもしれないので積む
        stack = append(stack, i)
    }

    return ans
}