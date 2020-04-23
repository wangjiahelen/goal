package linkedlist

import "testing"

func TestSingleLinkedList_InsertToHead(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToHead(i + 1)
	}
	l.Print()
}

func TestSingleLinkedList_InsertToTail(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
}

func TestSingleLinkedList_FindByIndex(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 10; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.FindByIndex(0))
	t.Log(l.FindByIndex(9))
	t.Log(l.FindByIndex(5))
	t.Log(l.FindByIndex(11))
}

func TestSingleLinkedList_DeleteNode(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 3; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()

	t.Log(l.DeleteNode(l.head.next))
	l.Print()

	t.Log(l.DeleteNode(l.head.next.next))
	l.Print()
}

func TestSingleLinkedList_Reverse(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i + 1)
	}
	l.Print()
	l.Reverse()
	l.Print()
}

func TestSingleLinkedList_HasCycle(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i + 1)
	}
	t.Log(l.HasCycle())
	l.head.next.next.next.next.next.next = l.head.next.next.next
	t.Log(l.HasCycle())
}

func TestMergeSortedList(t *testing.T) {
	l1 := NewSingleLinkedList()
	for i := 0; i < 5; i++ {
		l1.InsertToTail(i*2 + 1)
	}
	l2 := NewSingleLinkedList()
	for i := 0; i < 5; i++ {
		l2.InsertToTail(i * 2)
	}
	MergeSortedList(l1, l2).Print()
}

func TestSingleLinkedList_DeleteBottomN(t *testing.T) {
	l := NewSingleLinkedList()
	for i := 0; i < 5; i++ {
		l.InsertToTail(i + 1)
	}
	l.DeleteBottomN(3)
	l.Print()
	l.DeleteBottomN(2)
	l.Print()
	t.Log(l.FindMiddleNode())
}
