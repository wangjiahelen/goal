package design

import (
	"fmt"
	"sync"
	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()
	if instance1 != instance2 {
		fmt.Println("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*Singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			fmt.Println("instance is not equal too")
		}
	}
}

func TestParallelSingleton2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(100)
	instances := [100]*Singleton{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < 100; i++ {
		if instances[i] != instances[i-1] {
			fmt.Println("instance is not equal ya")
		}
	}
}

/*
//并发实现获取20个ID对应的基础信息和订单信息，获取成功后返回。或者规定20s内获取成功后返回，否则到时就把已经获取成功的返回
func GetInfo(arr []int) []string {
	wg := sync.WaitGroup{}
	wg.Add(20)
	results := [20]string{}
	for i := 0; i < 20; i++ {
		go func(id int) {
			results[i] = append(results[i], getBasic(id), getMore(id))
			wg.Done()
		}(i)
	}
	wg.Wait()
	return results
}
*/
