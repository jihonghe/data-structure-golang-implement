package single_link_list

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	// 校验关键点：
	// 1. 校验初始化空表时表的长度及第一个元素
	// 2. 校验初始化的非空表时，表达的长度及表是否真的不为空

	list := New()

	if list.first != nil || list.length != 0 {
		t.Error("New() Error: single link list without elements initialize should be empty.")
	}

	list = New(3, 1, true)
	if list.length != 3 || list.first == list.last {
		t.Error("New() Error: single link list with elements initialize failed.")
	}
	list.Traverse()
}

func TestSingleLinkList_Append(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能够正确添加元素
	// 2. 表不为空时，能正确的添加到表尾
	// 3. 元素添加后长度要正确

	list := New()

	list.Append(3)
	if list.length != 1 || list.first.value != 3 {
		t.Error("Append() Error: the length or the added element value error.")
	}

	list.Append(true)
	if list.length != 2 || list.last.value != true {
		t.Error("Append() Error: the length or the added element value error.")
	}
}

func TestSingleLinkList_Prepend(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能正确的添加元素
	// 2. 表不为空时，能正确添加元素到表头
	// 3. 添加元素后，表长度要正确

	list := New()
	list.Prepend(false)
	if list.length != 1 || list.last != list.first || list.first.value != false {
		t.Error("Prepend() Error: add element to the head of empty list failed.")
	}

	list.Prepend("Hello, 2020")
	if list.length != 2 || list.first.value != "Hello, 2020" {
		t.Error("Prepend() Error: add element to the head of list failed.")
	}
}

func TestSingleLinkList_BulkAppend(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确添加元素
	// 2.表不为空时，能否在表尾正确添加元素

	list := New()
	list.BulkAppend(3, "a", false)
	if list.length != 3 {
		t.Error("BulkAppend() Error: add elements to empty list failed.")
	}

	list.BulkAppend(-1, 9)
	if list.length != 5 || list.last.value != 9 {
		t.Error("BulkAppend() Error: add elements to list failed.")
	}
}

func TestSingleLinkList_Insert(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理插入位置为0和其他位置的情况
	// 2. 表不为空时，能否正确处理特殊位置的插入(表头与表尾)
	// 3. 表不为空时，能否正确处理非首尾位置的插入

	list := New()
	list.Insert(3, 9)
	list.Insert(-1, -1)
	list.Insert(9, 0)
	if list.length != 1 || list.first.value != 9 {
		t.Error("Insert() Error: add element to empty list failed.")
	}

	list.Insert("first", 0)
	list.Insert("last", list.length)
	list.Insert("middle", list.length/2)
	if list.length != 4 || list.first.value != "first" || list.last.value != "last" {
		t.Error("Insert() Error: add element to list failed.")
	}
}

func TestSingleLinkList_DeleteFirst(t *testing.T) {
	// 校验关键点：
	// 1. 当表为空时能否正确处理
	// 2. 当表不为空时能否正确删除：表长、第一个元素

	list := New()
	deletedValue := list.DeleteFirst()
	if deletedValue != nil {
		t.Error("DeleteFirst() Error(): the empty list should delete operation should return nil.")
	}

	list.BulkAppend(3, true, "hello")
	deletedValue = list.DeleteFirst()
	if deletedValue != 3 || list.length != 2 {
		t.Error("DeleteFirst() Error: list delete the first element failed.")
	}
}

func TestSingleLinkList_DeleteLast(t *testing.T) {
	// 校验关键点：
	// 1. 当表为空时能否正确处理
	// 2. 当表不为空时是否有效删除表尾元素

	list := New()
	deletedValue := list.DeleteLast()
	if deletedValue != nil {
		t.Error("DeleteLast() Error: the empty list delete operation should return nil.")
	}

	list.BulkAppend(2, 1, "i", nil, nil)
	deletedValue = list.DeleteLast()
	if deletedValue != nil || list.length != 4 {
		t.Error("DeleteLast() Error: list delete the last element failed.")
	}
	list.Traverse()
}

func TestSingleLinkList_Delete(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，能否正确处理特殊位置的删除：表尾和表首
	// 3. 删除中间位置的元素
	// 4. 删除元素后，表长是否正确

	list := New()
	deletedValue := list.Delete(9)
	if deletedValue != nil || list.length != 0 {
		t.Error("Delete() Error: the empty list delete operation should return nil.")
	}

	list.BulkAppend(3, 1, 9, true, "hello", false, nil)
	deletedValue = list.Delete(0)
	if deletedValue != 3 || list.length != 6 {
		t.Error("Delete() Error: list delete operation failed.")
	}
	list.Traverse()
	deletedValue = list.Delete(5)
	if deletedValue != nil || list.length != 5 {
		t.Error("Delete() Error: list delete operation failed.")
	}
	list.Traverse()

	deletedValue = list.Delete(2)
	if elemVal := deletedValue.(bool); !elemVal || list.length != 4 {
		t.Error("Delete() Error: list delete operation failed.")
	}
	list.Traverse()
}

func TestSingleLinkList_Clear(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时，清除链表后的表长度及表的首尾指针的指向

	list := New()
	deleteCount := list.Clear()
	if deleteCount != 0 {
		t.Error("Clear() Error: the empty list clear operation should return 0.")
	}

	list.BulkAppend(3, true, "hello")
	deleteCount = list.Clear()
	if deleteCount != 3 || list.first != nil || list.last != list.last {
		t.Error("Clear() Error: list clear operation failed.")
	}
}

func TestSingleLinkList_Get(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时能否正确返回对应的值

	list := New()
	val, err := list.Get(9)
	if err == nil || val != nil {
		t.Error("Get() Error: the empty list return errors should not be nil.")
	}

	list.BulkAppend(3, true, "hello", 3.14)
	val, err = list.Get(2)
	if err != nil || val != "hello" {
		t.Error("Get() Error: the list get operation failed.")
	}
	fmt.Printf("%v\n", val)
}

func TestSingleLinkList_Set(t *testing.T) {
	// 校验关键点：
	// 1. 表为空时能否正确处理
	// 2. 表不为空时能否正确处理不合理的index
	// 3. 表不为空时，能否正确修改合理的index的值

	list := New()
	oldVal, err := list.Set(false, 3)
	if err == nil || oldVal != nil {
		t.Error("Set() Error: the empty list set operation return error should not be nil.")
	}

	list.BulkAppend(3, 2, true, "hello")
	oldVal, err = list.Set(9, 7)
	if err == nil || oldVal != nil {
		t.Error("Set() Error: the invalid index of list set operation failed.")
	}
	oldVal, err = list.Set(1, 1)
	val, err1 := list.Get(1)
	if err != nil || err1 != nil || val != 1 {
		t.Error("Get() Error: the list get operation failed.")
	}
}
