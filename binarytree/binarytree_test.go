package binarytree

import (
	"fmt"
	"testing"
)

var tcs = []struct {
	pre, in, post []int
}{
	{
		[]int{1, 2, 3},
		[]int{1, 3, 2},
		[]int{3, 2, 1},
	},
	{
		[]int{1, 2, 4, 5, 3, 6, 7},
		[]int{4, 2, 5, 1, 6, 3, 7},
		[]int{4, 5, 2, 6, 7, 3, 1},
	},
}

//根据前序和中序构建二叉树
func PreIn2Tree(pre, in []int) *TreeNode {
	if len(pre) != len(in) {
		panic("preIn2Tree 中两个切片的长度不相等")
	}
	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = PreIn2Tree(pre[1:idx+1], in[:idx])
	res.Right = PreIn2Tree(pre[idx+1:], in[idx+1:])

	return res
}

func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}
	return 0
}

func Test_OrderTraversal(t *testing.T) {
	for _, tc := range tcs {
		root := PreIn2Tree(tc.pre, tc.in)
		fmt.Println("前序遍历1结果： ", preOrderTraversal(root))
		fmt.Println("前序遍历2结果： ", preorder(root))

		fmt.Println("中序遍历1结果： ", inOrderTraversal(root))
		fmt.Println("中序遍历2结果： ", inorder(root))

		fmt.Println("后序遍历1结果： ", postOrderTraversal(root))
		fmt.Println("后序遍历2结果： ", postorder(root))
	}
}
