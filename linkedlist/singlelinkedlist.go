package linkedlist

import (
	"fmt"
)

/*
单链表的定义和基本操作，支持增删操作
循环链表的定义和基本操作，支持增删操作
双向链表的定义和基本操作，支持增删操作
实现单链表的反转
实现两个有序链表合并为一个有序链表
实现求链表的中间节点
*/

//单链表节点
type SingleListNode struct {
	next  *SingleListNode
	value interface{}
}

//单链表
type SingleLinkedList struct {
	head   *SingleListNode
	length uint
}

func NewSingleListNode(v interface{}) *SingleListNode {
	return &SingleListNode{nil, v}
}

func (this *SingleListNode) GetNext() *SingleListNode {
	return this.next
}

func (this *SingleListNode) GetValue() interface{} {
	return this.value
}

func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{
		NewSingleListNode(0),
		0,
	}
}

//在某个节点后面插入节点
func (this *SingleLinkedList) InsertAfter(p *SingleListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewSingleListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

func (this *SingleLinkedList) InsertBefore(p *SingleListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}

	newNode := NewSingleListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

func (this *SingleLinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

func (this *SingleLinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for cur.next != nil {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

func (this *SingleLinkedList) FindByIndex(index uint) *SingleListNode {
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

func (this *SingleLinkedList) DeleteNode(p *SingleListNode) bool {
	if nil == p {
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
	if nil == cur {
		return false
	}
	pre.next = cur.next
	p = nil //空置删除的节点，防止内存泄漏
	this.length--
	return true
}

func (this *SingleLinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())

		if nil != cur.next {
			format += "->"
		}
		cur = cur.next
	}
	fmt.Println(format)
}

//实现单链表的反转，时间复杂度O(N)
func (this *SingleLinkedList) Reverse() {
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}

	var pre *SingleListNode = nil
	cur := this.head.next
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
}

//判断单链表是否有环
func (this *SingleLinkedList) HasCycle() bool {
	if nil != this.head {
		slow, fast := this.head, this.head
		for nil != fast && nil != fast.next {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}

//实现两个有序链表合并为一个有序链表
func MergeSortedList(l1, l2 *SingleLinkedList) *SingleLinkedList {
	if nil == l1 || nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2 || nil == l2.head || nil == l2.head.next {
		return l1
	}
	l := &SingleLinkedList{head: &SingleListNode{}}
	cur := l.head
	curl1 := l1.head.next
	curl2 := l2.head.next
	for nil != curl1 && nil != curl2 {
		if curl1.value.(int) > curl2.value.(int) {
			cur.next = curl2
			curl2 = curl2.next
		} else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}

	if nil != curl1 {
		cur.next = curl1
	} else if nil != curl2 {
		cur.next = curl2
	}
	return l
}

//删除倒数第N个结点
func (this *SingleLinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == this.head || nil == this.head.next {
		return
	}
	slow, fast := this.head, this.head
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.next
	}
	if nil == fast {
		return
	}

	for nil != fast.next {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
}

//实现求链表的中间节点
func (this *SingleLinkedList) FindMiddleNode() *SingleListNode {
	if nil == this.head || nil == this.head.next {
		return nil
	}
	if nil == this.head.next.next {
		return this.head.next
	}
	slow, fast := this.head, this.head
	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
