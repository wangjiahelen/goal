package tree

type BST struct {
	*BinaryTree
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BST {
	if nil == compareFunc {
		return nil
	}
	return &BST{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (this *BST) Find(v interface{}) *Node {
	p := this.root
	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 {
			p = p.right
		} else {
			p = p.left
		}
	}
	return nil
}

func (this *BST) Insert(v interface{}) bool {
	p := this.root
	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return false
		} else if compareResult > 0 {
			if p.right == nil {
				p.right = NewNode(v)
				break
			}
			p = p.right
		} else {
			if p.left == nil {
				p.left = NewNode(v)
				break
			}
			p = p.left
		}
	}
	return true
}

func (this *BST) Delete(v interface{}) bool {
	var pp *Node = nil
	p := this.root
	deleteLeft := false
	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult > 0 {
			pp = p
			p = p.right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.left
			deleteLeft = true
		} else {
			break
		}
	}

	if p == nil { //要删除的节点不存在
		return false
	} else if p.left == nil && p.right == nil { //删除的是一个叶子节点
		if pp != nil {
			if deleteLeft {
				pp.left = nil
			} else {
				pp.right = nil
			}
		} else { //根节点
			this.root = nil
		}
	} else if p.right != nil { //删除的是一个有右孩子，不一定有左孩子的节点
		//找到p节点右孩子的最小节点
		pq := p
		q := p.right
		fromRight := true
		for q.left != nil {
			pq = q
			q = q.left
			fromRight = false
		}
		if fromRight {
			pq.right = nil
		} else {
			pq.left = nil
		}
		q.left = p.left
		q.right = p.right
		if pp == nil { //根节点被删除
			this.root = q
		} else {
			if deleteLeft {
				pq.left = q
			} else {
				pq.right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if pp != nil {
			if deleteLeft {
				pp.left = p.left
			} else {
				pp.right = p.left
			}
		} else {
			if deleteLeft {
				this.root.left = p.left
			} else {
				this.root.right = p.left
			}
		}
	}
	return true
}

func (this *BST) Min() *Node {
	p := this.root
	for p.left != nil {
		p = p.left
	}
	return p
}

func (this *BST) Max() *Node {
	p := this.root
	for p.right != nil {
		p = p.right
	}
	return p
}
