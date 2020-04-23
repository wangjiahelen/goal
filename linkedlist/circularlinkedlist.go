package linkedlist

type CircularLinkedNode struct {
	next  *CircularLinkedNode
	value interface{}
}

type CircularLinkedList struct {
	head   *CircularLinkedNode
	length uint
}

func NewCircularLinkedNode(v interface{}) *CircularLinkedNode {
	return &CircularLinkedNode{value: v}
}

func (this *CircularLinkedNode) GetNext() *CircularLinkedNode {
	return this.next
}

func (this *CircularLinkedNode) GetValue() interface{} {
	return this.value
}

func NewCircularLinkedList() *CircularLinkedList {
	return &CircularLinkedList{
		NewCircularLinkedNode(0),
		0,
	}
}
