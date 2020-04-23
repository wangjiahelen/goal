package leetcode

import (
	"github.com/sirupsen/logrus"
)

//最小覆盖子串
func TestMinWindow(s string, t string) string {
	//1 set our needs and windows
	needs, windows := make(map[byte]int, len(t)), make(map[byte]int, len(t))
	var left, right, resL, resR, match, minLen int
	tt := []byte(t)
	for _, v := range tt {
		needs[v]++
	}
	//2 do the process right++ unless match==len(needs),sliding window
	for right < len(s) {
		if _, ok := needs[s[right]]; ok {
			windows[s[right]]++
			if windows[s[right]] == needs[s[right]] {
				match++
			}
		}
		right++
		//3 while match==len(needs),update res, left++
		for match == len(needs) {
			//update res
			if minLen == 0 || right-left+1 < minLen {
				resL = left
				resR = right
				minLen = right - left + 1
			}

			if _, ok := needs[s[left]]; ok {
				windows[s[left]]--
				if windows[s[left]] < needs[s[left]] {
					match--
				}
			}
			left++
		}

	}
	if minLen == 0 {
		return ""
	}
	return s[resL : resR+1]

}

//斐波那契数
func FList(N int) int {
	if N < 2 {
		return N
	}
	if N == 2 {
		return 1
	}
	pre, cur := 1, 1
	for i := 3; i <= N; i++ {
		sum := pre + cur
		pre = cur
		cur = sum
	}
	return cur
}

func test() {
	queue := make([]int, 0)
	queue = append(queue, 1, 2)
	logrus.Info(queue)
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func rightSideView(root *TreeNode) []int {
	res = []int{}
	dfs(root, 1)
	return res
}
func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}
	if level > len(res) {
		res = append(res, root.Val)
	}
	dfs(root.Right, level+1)
	dfs(root.Left, level+1)
}

func rightSideViewBFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{}
	var cur, next = []*TreeNode{root}, []*TreeNode{}
	for len(cur) != 0 {
		node := cur[0]
		cur = cur[1:]
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
		if len(cur) == 0 {
			result = append(result, node.Val)
			cur = next
			next = []*TreeNode{}
		}
	}
	return result
}

//层次遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	var queue = []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res = append(res, []int{})
		for 0 < counter {
			counter--
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			}
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val)
			queue = queue[1:]
		}
		level++
	}
	return res
}
