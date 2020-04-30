package binarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func preOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	var stack []*TreeNode
	var res []int
	stack = append(stack, root)
	for len(stack) != 0 {
		e := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, e.Val)
		if e.Right != nil {
			stack = append(stack, e.Right)
		}
		if e.Left != nil {
			stack = append(stack, e.Left)
		}
	}
	return res
}

//中序遍历
func inOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := inOrderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inOrderTraversal(root.Right)...)

	return res
}

//后序遍历
func postOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		lres := postOrderTraversal(root.Left)
		if len(lres) > 0 {
			res = append(res, lres...)
		}
	}
	if root.Right != nil {
		rres := postOrderTraversal(root.Right)
		if len(rres) > 0 {
			res = append(res, rres...)
		}
	}
	res = append(res, root.Val)
	return res
}

//中序遍历
func inorder(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	inorderHelper(root, &res)
	return res
}
func inorderHelper(root *TreeNode, res *[]int) {
	if root.Left != nil {
		inorderHelper(root.Left, res)
	}
	*res = append(*res, root.Val)
	if root.Right != nil {
		inorderHelper(root.Right, res)
	}
}

//前序遍历
func preorder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	preorderHelper(root, &res)
	return res
}

func preorderHelper(root *TreeNode, res *[]int) {
	*res = append(*res, root.Val)
	if root.Left != nil {
		preorderHelper(root.Left, res)
	}
	if root.Right != nil {
		preorderHelper(root.Right, res)
	}
}

//后序遍历
func postorder(root *TreeNode) []int {
	var res []int
	postorderHelper(root, &res)
	return res
}

func postorderHelper(root *TreeNode, res *[]int) {
	if root != nil {
		if root.Left != nil {
			postorderHelper(root.Left, res)
		}
		if root.Right != nil {
			postorderHelper(root.Right, res)
		}
		*res = append(*res, root.Val)
	}
}

//层序遍历
