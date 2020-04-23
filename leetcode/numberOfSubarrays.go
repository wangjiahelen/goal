package leetcode

func numberOfSubarrays(nums []int, k int) int {
	//记录odd的index
	odd := make([]int, len(nums)+2)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			count++
			odd[count] = i
		}
	}

	count++
	odd[0] = -1
	odd[count] = len(nums)
	//循环查找
	for i := 1; i+k <= count; i++ {
		ans += (odd[i] - odd[i-1]) * (odd[i+k] - odd[i+k-1])
	}
	return ans
}
