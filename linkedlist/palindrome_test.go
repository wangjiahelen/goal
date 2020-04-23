package linkedlist

import (
	"fmt"
	"testing"
)

func TestPalindrome1(t *testing.T) {
	strs := []string{"heooeh", "hello", "heoeh", "a", ""}
	for _, str1 := range strs {
		l := NewSingleLinkedList()
		for _, c := range str1 {
			l.InsertToTail(string(c))
		}
		l.Print()
		fmt.Println(isPalindrome1(l))
		fmt.Println(isPalindrome2(l))
		fmt.Println(isPalindrome(l.head))
		//t.Log(isPalindrome1(l))
		//t.Log(isPalindrome2(l))
		//t.Log(isPalindrome(l.head))
	}
}
