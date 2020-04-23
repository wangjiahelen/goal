package array

import (
	"errors"
	"fmt"
)

/**
* 1)数组的插入、删除、按照下标随机访问操作；
* 2)数组中的数据是interface类型；
 */

type Array struct {
	data []interface{}
	size int
}

//init array with default size of 10
func NewArray(capacity ...int) (array *Array) {
	if len(capacity) >= 1 && capacity[0] != 0 {
		array = &Array{
			data: make([]interface{}, capacity[0]),
			size: 0,
		}
	} else {
		array = &Array{
			data: make([]interface{}, 10),
			size: 0,
		}
	}
	return
}

//is index out of range
func (array *Array) IsIndexOutOfRange(index int) bool {
	if index < 0 || index >= array.size {
		return true
	}
	return false
}

//resize array
func (array *Array) resize(capacity int) {
	newArray := make([]interface{}, capacity)
	for i := 0; i < array.size; i++ {
		newArray[i] = array.data[i]
	}

	array.data = newArray
	newArray = nil
}

func (array *Array) GetCapacity() int {
	return cap(array.data)
}

func (array *Array) GetSize() int {
	return array.size
}

func (array *Array) IsEmpty() bool {
	return array.size == 0
}

// add value at index, O(m+n)
func (array *Array) Add(index int, value interface{}) (err error) {
	if index < 0 || index > array.size {
		err = errors.New("Add failed. Index out of range")
		return
	}

	//如果当前元素个数等于数组容量，则将数组扩容为原来的2倍
	capacity := array.GetCapacity()
	if array.size == capacity {
		array.resize(capacity * 2)
	}

	for i := array.size - 1; i >= index; i-- {
		array.data[i+1] = array.data[i]
	}

	array.data[index] = value
	array.size++
	return
}

func (array *Array) AddHead(value interface{}) error {
	return array.Add(0, value)
}

func (array *Array) AddTail(value interface{}) error {
	return array.Add(array.size, value)
}

//获取对应INDEX位置的元素
func (array *Array) Get(index int) (value interface{}, err error) {
	if array.IsIndexOutOfRange(index) {
		err = errors.New("Get failed. Index is illegal.")
		return
	}

	value = array.data[index]
	return
}

//修改index位置的元素
func (array *Array) Set(index int, value interface{}) (err error) {
	if array.IsIndexOutOfRange(index) {
		err = errors.New("Set failed. Index is illegal.")
		return
	}

	array.data[index] = value
	return
}

func (array *Array) Contains(value interface{}) bool {
	for i := 0; i < array.size; i++ {
		if array.data[i] == value {
			return true
		}
	}
	return false
}

//通过值查找索引，索引范围[0,N-1]（未找到，返回－1）
func (array *Array) Find(value interface{}) int {
	for i := 0; i < array.size; i++ {
		if array.data[i] == value {
			return i
		}
	}
	return -1
}

func (array *Array) Delete(index int) (value interface{}, err error) {
	if array.IsIndexOutOfRange(index) {
		err = errors.New("Delete failed. Index is illegal.")
		return
	}

	value = array.data[index]
	for i := index + 1; i < array.size; i++ {
		array.data[i-1] = array.data[i]
	}

	array.size--
	array.data[array.size] = nil //loitering objects != memory leak

	//数组元素个数为当前容量 1/4 时，缩容为当前容量的一半。
	capacity := array.GetCapacity()
	if array.size == capacity/4 && capacity/2 != 0 {
		array.resize(capacity / 2)
	}

	return
}

func (array *Array) DeleteHead() (interface{}, error) {
	return array.Delete(0)
}

func (array *Array) DeleteTail() (interface{}, error) {
	return array.Delete(int(array.size - 1))
}

func (array *Array) DeleteElement(value interface{}) (e interface{}, err error) {
	index := array.Find(value)
	if index != -1 {
		e, err = array.Delete(index)
	}
	return
}

func (array *Array) Clear() {
	array.data = make([]interface{}, array.size)
	array.size = 0
}

func (array *Array) Print() {
	var format string
	format = fmt.Sprintf("Array: size = %d , capacity = %d\n", array.GetSize(), array.GetCapacity())
	format += "["
	for i := 0; i < array.size; i++ {
		format += fmt.Sprintf("%+v", array.data[i])
		if i != array.size-1 {
			format += ", "
		}
	}
	format += "]"
	fmt.Println(format)
}
