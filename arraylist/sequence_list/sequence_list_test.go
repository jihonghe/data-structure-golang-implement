package sequence_list

import (
	"testing"
)

func TestListNew(t *testing.T) {
	// 校验关键点：
	// 1. 初始化空表时，判断长度是否为0
	// 2. 初始化非空表时，判断长度是否合理

	list1 := New()
	if !list1.IsEmpty() {
		t.Error("New() Error: sequence list is not empty.")
	}

	list2 := New(1, 2, 3, 4)
	if list2.Length() != 4 {
		t.Error("New() Error: sequence list length is invalid.")
	}
}

func TestSequenceList_Append(t *testing.T) {
	// 校验关键点：
	// 1. 校验插入后元素的位置与值是否相符
	// 2. 校验表长度是否合理

	list := New(1, 3, 5, 7, 9, 2, 4, 6)

	list.Append(8)
	if list.IsFull() {
		t.Error("Append() Error: sequence list should not be full.")
	}

	index := list.Append(10)
	if index != 9 || list.elements[index] != 10 || list.Length() != 10 {
		t.Error("Append() Error: Append() failed.")
	}

	indexEOF := list.Append(11)
	if indexEOF != -1 {
		t.Error("Append() Error: sequence list is full, but can append element.")
	}
}

func TestSequenceList_Insert(t *testing.T) {
	// 校验关键点：
	// 1. 能否正确处理插入位置不合理的操作
	// 2. 插入的元素是否在指定的位置上
	// 3. 表长度是否合理
	// 4. 校验插入后，其后的子切片是否正确移动(省了)

	list := New()

	elementValue, index := 3, 9
	finished := list.Insert(elementValue, index)
	if finished {
		t.Error("Insert() Error: the invalid index should not be insert successfully.")
	}


	print("Before Insert(): ")
	for _, element := range list.elements {
		print(element, ",\t")
	}
	println()
	elementValue, index = 5, 0
	lengthBeforeInsert := list.length
	finished = list.Insert(elementValue, index)
	if !finished {
		t.Error("Insert() Error: it should be finished, however it was not.")
	}
	if list.elements[index] != elementValue {
		t.Error("Insert() Error: the inserted value is not at the index that should be inserted.")
	}
	if list.Length() - lengthBeforeInsert != 1 {
		t.Error("Insert() Error: the length of list after Insert() done is invalid.")
	}
	print("After Insert(): ")
	for _, element := range list.elements {
		print(element, ",\t")
	}
	println()
}

func TestSequenceList_DeleteLast(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，元素是否正确删除，表的长度是否正确

	list := New()
	deleted := list.DeleteLast()
	if deleted {
		t.Error("DeleteLast() Error: the empty list should not finish delete.")
	}

	list.Append(4)
	list.Append(3)
	list.Append(3)
	deleted = list.DeleteLast()
	if !deleted || len(list.elements) != 2 {
		t.Error("DeleteLast() Error: the length of list who has only one element after function DeleteLast() is not equal 0.")
	}
}

func TestSequenceList_Delete(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，能否正确处理边界条件：删除首尾元素
	// 3. 表不为空时，能否正确处理删除非首尾元素

	list := New()
	deleted := list.Delete(3)
	if deleted {
		t.Error("Delete() Error: empty list cannot delete element.")
	}

	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Append(6)
	list.Append(7)
	deleted = list.Delete(0)
	if !deleted || len(list.elements) != 4 {
		t.Error("Delete() Error: the length of list is error after deleting an element.")
	}

	list.Delete(3)
	if !deleted || len(list.elements) != 3 {
		t.Error("Delete() Error: the length of list is error after deleting an element.")
	}
	list.Delete(1)
	if !deleted || len(list.elements) != 2 {
		t.Error("Delete() Error: the length of list is error after deleting an element.")
	}
}

func TestSequenceList_Clear(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时
	// 2. 表不为空时，经过Clear()后表长度是否为0

	list := New()
	cleared := list.Clear()
	if !cleared {
		t.Error("Clear() Error: empty list clear should return true.")
	}
	list.Append(9)
	list.Append(9)
	list.Append(9)
	list.Append(9)
	cleared = list.Clear()
	if !cleared {
		t.Error("Clear() Error: the list with elements dose not clear successfully.")
	}
}

func TestSequenceList_Set(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，能否正确修改对应位置的元素

	list := New()
	updated := list.Set(3, 1)
	if updated {
		t.Error("Set() Error: empty list should not return true.")
	}

	list.Append(3)
	list.Append(3)
	list.Append(3)
	updated = list.Set(9, 1)
	if !updated || list.elements[1] != 9 {
		t.Error("Set Error: Set() failed.")
	}
}

func TestSequenceList_Get(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，若所给的index不合理能否正确处理，若index合理，校验返回值与在表中的实际值是否一致

	list := New()
	value, err := list.Get(3)
	if err == nil {
		t.Error("Get() Error: empty list should returns error.")
	}

	list.Append(3)
	list.Append(4)
	list.Append(5)
	value, err = list.Get(5)
	if err == nil {
		t.Error("Get() Error: invalid index should returns error.")
	}

	value, err = list.Get(1)
	if err != nil {
		t.Error("Get() Error: the index is valid, it should not return error.")
	}
	if value != list.elements[1] {
		t.Error("Get() Error: the returned value is not equal to the element`s value in the list.")
	}
}

func TestSequenceList_Index(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时，能否正确处理
	// 2. 表不为空时，查找不在表中的元素能否正确返回-1
	// 3. 表不为空时，查找在表中的元素能否正确返回相应的index

	list := New()
	index := list.Index(9)
	if index != -1 {
		t.Error("Index() Error: empty list should return -1")
	}

	list02 := New(3, 1, 4, 5, 0, 1)
	index = list02.Index(9)
	if index != -1 {
		t.Error("Index() Error: 9 is not in list, Index() should return -1")
	}

	index = list02.Index(1)
	if index != 1 {
		t.Error("Index() Error: the returned index is not equal to the index that the element in list.")
	}
}

func TestSequenceList_Contains(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，所得结果是否与实际相符

	list := New()
	in := list.Contains(3)
	if in {
		t.Error("Contains() Error: empty list should return false")
	}

	list02 := New(3, 1, 9)
	in = list02.Contains(0)
	if in {
		t.Error("Contains() Error: element 0 is not in the list")
	}
}
