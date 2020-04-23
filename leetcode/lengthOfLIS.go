package leetcode

//最长上升子序列,动态规划+贪心+二分
func lengthOfLIS(nums []int) int {
	//1将最长上升子序列放入tails中保存
	tails, res := make([]int, len(nums)), 0
	//2 遍历nums
	for _, v := range nums {
		//3 根据贪心，number比tails中都大，就直接加在后边。否则二分法寻找比当前number大的数中最小的，进行覆盖
		left, right := 0, res
		for left < right {
			mid := (left + right) >> 1
			if tails[mid] < v {
				left = mid + 1
			} else {
				right = mid
			}
		}
		tails[left] = v
		if right == res {
			res++
		}
	}

	return res
}
