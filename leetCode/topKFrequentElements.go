package leetcode

import "sort"

func topKFrequent(nums []int, k int) []int {
    countByElement := make(map[int]int)
    for _, n := range nums {
        countByElement[n]++
    }

    type pair struct {
        num   int
        count int
    }

    pairs := make([]pair, 0, len(countByElement))
    for num, c := range countByElement {
        pairs = append(pairs, pair{num: num, count: c})
    }

    sort.Slice(pairs, func(i, j int) bool {
        return pairs[i].count > pairs[j].count
    })

    res := make([]int, 0, k)
    for i := 0; i < k && i < len(pairs); i++ {
        res = append(res, pairs[i].num)
    }

    return res
}
