package main

import (
	"bytes"
	"fmt"
	"sort"
)
// https://leetcode.cn/problems/sort-characters-by-frequency/solution/gen-ju-zi-fu-chu-xian-pin-lu-pai-xu-by-l-zmvy/
func frequencySort(s string) string {
	cnt := map[byte]int{}
	for i := range s {
		cnt[s[i]]++
	}

	type pair struct {
		ch  byte
		cnt int
	}
	pairs := make([]pair, 0, len(cnt))
	for k, v := range cnt {
		pairs = append(pairs, pair{k, v})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].cnt > pairs[j].cnt })

	ans := make([]byte, 0, len(s))
	for _, p := range pairs {
		ans = append(ans, bytes.Repeat([]byte{p.ch}, p.cnt)...)
	}
	return string(ans)
}

func frequencySortNew(s string) string {
    cnt := map[byte]int{}
    maxFreq := 0
    for i := range s {
        cnt[s[i]]++
        maxFreq = max(maxFreq, cnt[s[i]])
    }

    buckets := make([][]byte, maxFreq+1)
    for ch, c := range cnt {
        buckets[c] = append(buckets[c], ch)
    }

    ans := make([]byte, 0, len(s))
    for i := maxFreq; i > 0; i-- {
        for _, ch := range buckets[i] {
            ans = append(ans, bytes.Repeat([]byte{ch}, i)...)
        }
    }
    return string(ans)
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}


func main() {
	test := "cccbaaa"
	fmt.Println(frequencySort(test))
	fmt.Println(frequencySortNew(test))
}
