package main

func checkInclusion(s1 string, s2 string) bool {

	need := make(map[byte]int)
	window := make(map[byte]int)
	for val := range s1 {
		need[s1[val]]++
	}
	left, right, valid := 0, 0, 0
	for right < len(s2) {
		c := s2[right]
		right++
		// 滑动窗口right 往右移动部分
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// 滑动窗口left 往右移动部分
		for right-left >= len(s1) {
			if valid == len(need) {
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
