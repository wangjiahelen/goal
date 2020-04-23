package arraytest

/*
单链表的定义和基本操作，支持增删操作
单链表反转
链表中环的检测
两个有序的链表合并
删除链表倒数第 n 个结点
求链表的中间结点
*/
type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length uint
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (this *Node) GetNext() *Node {
	return this.next
}

func (this *Node) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		NewNode(0),
		0,
	}
}

func (this *LinkedList) InsertAfter(p *Node, v interface{}) bool {
	if p == nil {
		return false
	}
	n1 := NewNode(v)
	oldNext := p.next
	p.next = n1
	n1.next = oldNext

	this.length++
	return true
}

func (this *LinkedList) InsertBefore(p *Node, v interface{}) bool {
	if p == nil || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	n1 := NewNode(v)
	pre.next = n1
	n1.next = cur
	this.length++
	return true
}

func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

func (this *LinkedList) FindByIndex(index uint) *Node {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

func (this *LinkedList) DeleteNode(p *Node) bool {
	if p == nil {
		return false
	}
}
