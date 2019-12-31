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
