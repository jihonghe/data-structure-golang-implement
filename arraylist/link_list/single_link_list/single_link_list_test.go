package single_link_list

import (
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
