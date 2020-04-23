package array

import "testing"

/**
*实现一个支持动态扩容的数组TestArray_Add
*实现一个大小固定的有序数组，支持动态增删改操作TestArray_Add,TestArray_Set,TestArray_Delete
*实现两个有序数组合并为一个有序数组TestArray_Merge
 */
//测试实例化
func TestNewArray(t *testing.T) {
	array := NewArray(2)
	for i := 0; i < 10; i++ {
		if err := array.Add(i, i+1); err != nil {
			t.Error(err)
			break
		}
	}

	array.Print()
	if err := array.Add(10, 11); err != nil {
		t.Error(err)
	}

	if err := array.Add(10, 12); err != nil {
		t.Error(err)
	}

	array.Print()
}

//获取数组容量
func TestArray_GetCapacity(t *testing.T) {
	array := NewArray()
	array.AddTail("我爱学习")
	array.Print()
	t.Logf("array cap:%d \n", array.GetCapacity())
}

//获取数组长度
func TestArray_GetSize(t *testing.T) {
	array := NewArray()
	t.Logf("array len: %d \n", array.GetSize())
}

//判断是否为空
func TestArray_IsEmpty(t *testing.T) {
	array := NewArray()
	t.Logf("array empty: %t \n", array.IsEmpty())
}

//向数组头部添加元素
func TestArray_AddHead(t *testing.T) {
	array := NewArray()
	array.AddHead("array add head")
	array.AddTail("array add tail")
	array.Add(1, "array add value")
	array.Print()
}

//向数组尾部添加元素
func TestArray_AddTail(t *testing.T) {
	array := NewArray()
	array.Add(0, 23)
	array.Add(1, 23)
	array.Add(2, 23)
	array.AddTail(20)
	array.Print()
}

//动态添加元素
func TestArray_Add(t *testing.T) {
	strs := []string{"A", "B", "C", "D", "E", "F"}
	array := NewArray()
	for i := 0; i < len(strs); i++ {
		if err := array.Add(i, strs[i]); err != nil {
			t.Error(err)
			break
		}
	}
	array.Print()
}

//动态添加元素性能测试
func BenchmarkArray_Add(b *testing.B) {
	array := NewArray()

	//重置时间点
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		array.Add(i, i+1)
	}

	array.Print()
}

//根据索引获取某个值
func TestArray_Get(t *testing.T) {
	array := NewArray(10)
	array.AddHead(10)
	if value, err := array.Get(0); err != nil {
		t.Error(err)
	} else {
		t.Log(value)
	}
}

//根据索引修改某个值
func TestArray_Set(t *testing.T) {
	array := NewArray()
	array.AddHead("array add head")
	array.Print()
	array.Set(0, "array set value")
	array.Print()
}

//判断数组是否存在某个值
func TestArray_Contains(t *testing.T) {
	array := NewArray()
	if array.Contains("我爱学习") {
		t.Log("找到了")
	}
	array.AddHead("我爱学习")
	if array.Contains("我爱学习") {
		t.Log("找到了")
	}
	array.Print()
}

//查询一个值的索引
func TestArray_Find(t *testing.T) {
	array := NewArray(10)
	array.AddHead("Monday")
	array.AddHead("Tuesday")
	array.AddTail("Wednesday")
	find := array.Find("Monday")
	t.Log(find)
	array.Print()
}

//根据索引删除元素
func TestArray_Delete(t *testing.T) {
	array := NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}
	array.Print()
	array.Delete(2)
	array.Print()
	array.DeleteHead()
	array.Print()
	array.DeleteTail()
	array.Print()
	array.DeleteElement("E")
	array.Print()
}

//清空元素
func TestArray_Clear(t *testing.T) {
	array := NewArray()
	strs := []string{"A", "B", "C", "D", "E", "F"}
	for index, str := range strs {
		array.Add(index, str)
	}
	array.Print()
	array.Clear()
	array.Print()
}

//实现两个有序数组合并为一个有序数组
func TestArray_Merge(t *testing.T) {
	array1 := NewArray()
	array2 := NewArray()
	for i := 0; i < 5; i++ {
		array1.Add(i, 2*i+1)
	}
	array1.Print()
	for j := 0; j < 3; j++ {
		array2.Add(j, 2*j)
	}
	array2.Print()
	m := array1.GetSize()
	n := array2.GetSize()
	array1.size = m + n//Important to resize before merge

	for n > 0 {
		array2Last, _ := array2.Get(n - 1)
		array1Last, _ := array1.Get(m - 1)
		num2, _ := array2Last.(int)
		num1, _ := array1Last.(int)
		if m <= 0 || num2 > num1 {
			if err := array1.Set(m+n-1, array2Last); err != nil {
				t.Error(err)
			}
			n--
		} else {
			if err := array1.Set(m+n-1, array1Last); err != nil {
				t.Error(err)
			}
			m--
		}
	}
	array1.Print()
	t.Log(array1.data)
}
