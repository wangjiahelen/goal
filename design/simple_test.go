package design

import (
	"fmt"
	"testing"
)

func TestType1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Tom")
	if s != "Hi, Tom" {
		fmt.Println("Type1 test fail")
	}
}

func TestType2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Tom")
	if s != "Hello, Tom" {
		fmt.Println("Type2 test fail")
	}
}
