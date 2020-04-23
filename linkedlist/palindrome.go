package linkedlist

//回文字符串
func isPalindrome(head *SingleListNode) bool {
	var slow *SingleListNode = head.next
	var fast *SingleListNode = head.next
	var prev *SingleListNode = nil
	var temp *SingleListNode = nil

	if (head == nil || head.next == nil) {
		return false
	}
	if head.next.next == nil {
		return true
	}

	for (fast != nil && fast.next != nil) {
		fast = fast.next.next
		temp = slow.next
		slow.next = prev
		prev = slow
		slow = temp
	} // 快的先跑完,同时反转了一半链表,剪短

	if fast != nil {
		slow = slow.next // 处理余数,跨过中位数
		// prev 增加中 2->1->nil
	}

	var l1 *SingleListNode = prev
	var l2 *SingleListNode = slow

	for (l1 != nil && l2 != nil && l1.value == l2.value) {
		l1 = l1.next
		l2 = l2.next
	}

	return (l1 == nil && l2 == nil)
}

/*
思路1：开一个栈存放链表前半段
时间复杂度：O(N)
空间复杂度：O(N)
*/
func isPalindrome1(l *SingleLinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}
	if lLen == 1 {
		return true
	}

	s := make([]string, 0, lLen/2)
	cur := l.head
	for i := uint(1); i <= lLen; i++ {
		cur = cur.next
		if lLen%2 != 0 && i == (lLen/2+1) { //如果链表有奇数个节点，中间的直接忽略
			continue
		}

		if i <= lLen/2 { //前一半入栈
			s = append(s, cur.GetValue().(string))
		} else { //后一半与前一半进行对比
			if s[lLen-i] != cur.GetValue().(string) {
				return false
			}
		}
	}

	return true
}

/*
思路2
找到链表中间节点，将前半部分转置，再从中间向左右遍历对比
时间复杂度：O(N)
*/

func isPalindrome2(l *SingleLinkedList) bool {
	lLen := l.length
	if lLen == 0 {
		return false
	}

	if lLen == 1 {
		return true
	}

	var isPalindrome = true
	step := lLen / 2
	var pre *SingleListNode = nil
	cur := l.head.next
	next := l.head.next.next

	for i := uint(1); i <= step; i++ {
		tmp := cur.GetNext()
		cur.next = pre
		pre = cur
		cur = tmp
		next = cur.GetNext()
	}

	mid := cur

	var left, right *SingleListNode = pre, nil
	if lLen%2 != 0 {
		right = mid.GetNext()
	} else {
		right = mid
	}

	for left != nil && right != nil {
		if left.GetValue().(string) != right.GetValue().(string) {
			isPalindrome = false
			break
		}
		left = left.GetNext()
		right = right.GetNext()
	}

	cur = pre
	pre = mid
	for cur != nil {
		next = cur.GetNext()
		cur.next = pre
		pre = cur
		cur = next
	}
	l.head.next = pre

	return isPalindrome
}
