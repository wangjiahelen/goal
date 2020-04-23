package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//层次遍历
func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	queue := []*TreeNode{root}
	var level int
	for len(queue) > 0 {
		counter := len(queue)
		res[level] = append(res[level], []int{})
		for counter > 0 {
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

//前序+中序 构造二叉树
func buildTree1(preorder []int, inorder []int) *TreeNode {
	for k := range inorder {
		//前序第0个元素 与 中序第k个元素相等时
		if inorder[k] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree1(preorder[1:k+1], inorder[0:k]),
				Right: buildTree1(preorder[k+1:]),inorder[k+1:]),
			}
		}
	}
	return nil
}
//中序+后序 构造二叉树
func buildTree2(inorder []int,postorder []int) *TreeNode {
	l := len(postorder)-1
	for k:=range inorder {
		//中序第k个元素 与 后序第l个元素相等时
		if inorder[k]==postorder[l] {
			return &TreeNode{
				Val:postorder[l],
				Left:buildTree2(inorder[0:k],postorder[0:k]),
				Right:buildTree2(inorder[k+1:],postorder[k:l]),
			}
		}
	}
	return nil
}
//前序+后序 构造二叉树
func buildTree3(preorder []int,postorder []int) *TreeNode {
	l := len(postorder)-1
	if preorder[0] == postorder[l] {
		return &TreeNode{
			Val:preorder[0],
			Left:buildTree3(preorder[]),
			Right:buildTree3(),
		}
	}
	return nil
}