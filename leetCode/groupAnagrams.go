package leetcode

import "sort"

func groupAnagrams(strs []string) [][]string {
    strMap := make(map[string][]string)
    for _, str := range strs {
        chars := []rune(str)
        sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })
        key := string(chars)
        strMap[key] = append(strMap[key], str)
    }

    res := [][]string{}
    for _, v := range strMap {
        res = append(res, v)
    }
    return res
}
