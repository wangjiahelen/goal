package leetcode

/*滑动窗口*/

//最小覆盖子串
func minWindow(s string, t string) string {
	//put target into needs记录t中每个字母出现的次数
	var minLen, left, right, resL, resR, match, l int
	needs := make(map[byte]int, len(t))
	windows := make(map[byte]int, len(t))
	tt := []byte(t)
	for _, v := range tt {
		needs[v]++
	}
	//set left and right pointer,sliding window
	//1 only right++ unless match==len(needs)
	for right < len(s) {
		if _, ok := needs[s[right]]; ok {
			windows[s[right]]++
			if windows[s[right]] == needs[s[right]] {
				match++
			}
		}
		//2 only left++ unless match < len(needs)
		for match == len(needs) {
			//update res
			l = right - left + 1
			if minLen == 0 || l < minLen {
				resL = left
				resR = right
				minLen = l
			}
			if _, ok := needs[s[left]]; ok {
				windows[s[left]]--
				if windows[s[left]] < needs[s[left]] {
					match--
				}
			}
			left++
		}
		right++
	}
	if minLen == 0 {
		return ""
	}

	return s[resL : resR+1]

}

//找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	needs, window := make(map[byte]int, len(p)), make(map[byte]int, len(p))
	var left, right, match int
	var res []int
	pp := []byte(p)
	//1 set needs
	for _, v := range pp {
		needs[v]++
	}
	//2 right++
	for right < len(s) {
		if _, ok := needs[s[right]]; ok {
			window[s[right]]++
			if window[s[right]] == needs[s[right]] {
				match++
			}
		}
		right++
		//3 while len(p)==right-left, add left into res && left++
		for len(needs) == match {
			if len(p) == right-left {
				res = append(res, left)
			}
			if _, ok := needs[s[left]]; ok {
				window[s[left]]--
				if window[s[left]] < needs[s[left]] {
					match--
				}
			}
			left++
		}
	}
	return res
}

//无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	windows := make(map[byte]int, len(s))
	var left, right, res int
	for right < len(s) {
		windows[s[right]]++
		right++

		for windows[s[right-1]] > 1 {
			windows[s[left]]--
			left++
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}
