package arraytest

import "errors"

/**
* 1)数组的插入、删除、按照下标随机访问操作；
* 2)数组中的数据是interface类型；
 */
type Array struct {
	data []interface{}
	size int
}

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

func (array *Array) IsIndexOutOfRange(index int) bool {
	if index < 0 || index >= array.size {
		return true
	}
	return false
}

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

func (array *Array) IsEmpty() bool {
	return array.size == 0
}

func (array *Array) Add(index int, value interface{}) (err error) {
	if index < 0 || index > array.size {
		err = errors.New("Add Failed")
		return
	}
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

func (array *Array) Delete(index int) (value interface{}, err error) {
	if array.IsIndexOutOfRange(index) {
		err = errors.New("Delete Failed")
		return
	}

	value = array.data[index]
	for i := index + 1; i < array.size; i++ {
		array.data[i-1] = array.data[i]
	}
	array.size--
	array.data[array.size] = nil

	capacity := array.GetCapacity()
	//数组元素个数为当前容量 1/4 时，缩容为当前容量的一半。
	if array.size == capacity/4 && capacity/2 != 0 {
		array.resize(capacity / 2)
	}
	return
}
