package leetcode

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	//1 special conditions
	len1, len2 := len(s1), len(s2)
	var index1, index2, ans int
	if len1 == 0 || len2 == 0 || len1*n1 < len2*n2 {
		return ans
	}
	map1, map2 := make(map[int]int), make(map[int]int)
	//2 index1 为s1*n1下标，index2为s2*n2下标，循环
	for index1/len1 < n1 {
		//3 查找循环节，使用map记录循环节，如果找到，则快速跳过循环
		if index1%len1 == len1-1 {
			if val, ok := map1[index2%len2]; ok {
				cycleLen := (index1 - val) / len1
				cycleNum := (n1 - 1 - index1/len1) / cycleLen
				cycleS2Num := (index2 - map2[index2%len2]) / len2

				index1 += cycleLen * cycleNum * len1
				ans += cycleNum * cycleS2Num
			} else {
				map1[index2%len2] = index1
				map2[index2%len2] = index2
			}
		}
		//4 如果没有找到，则穷举
		if s1[index1%len1] == s2[index2%len2] {
			if index2%len2 == len2-1 {
				ans++
			}
			index2++
		}
		index1++
	}
	return ans / n2
}
