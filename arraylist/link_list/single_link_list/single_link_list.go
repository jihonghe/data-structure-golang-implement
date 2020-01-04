package single_link_list

import (
	"errors"
	"fmt"
)

type element struct {
	value interface{}
	next *element
}

type SingleLinkList struct {
	first *element
	last *element
	length int
}

// NewElement: 创建新节点
func NewElement(elementValue interface{}) *element {
	return &element{
		value: elementValue,
		next: nil,
	}
}

// New: 创建并初始化单链表
func New(elements ...interface{}) *SingleLinkList {
	if len(elements) == 0 {
		return &SingleLinkList{}
	}

	singleLinkList := &SingleLinkList{}
	for index, elementValue := range elements {
		if index == 0 {
			singleLinkList.first = NewElement(elementValue)
			singleLinkList.last = singleLinkList.first
		} else {
			singleLinkList.last.next = NewElement(elementValue)
			singleLinkList.last = singleLinkList.last.next
		}
	}
	singleLinkList.length = len(elements)

	return singleLinkList
}

// Traverse: 遍历单链表
func (l SingleLinkList) Traverse() {
	if l.first == nil && l.length == 0 {
		return
	}

	listElementsStr := "list: ["
	for element := l.first; element != nil; element = element.next {
		listElementsStr += fmt.Sprintf("%v, ", element.value)
	}
	listElementsStr += "]"
	println(listElementsStr)
}

// Append: 在表尾添加节点
func (l *SingleLinkList) Append(elementValue interface{}) {
	newElement := NewElement(elementValue)

	if l.IsEmpty() {
		l.first = newElement
		l.last = newElement
		l.length++

		return
	}

	l.last.next = &element{
		value: elementValue,
		next: nil,
	}
	l.last = l.last.next
	l.length++
}

// Prepend: 在表头添加元素
func (l *SingleLinkList) Prepend(elementValue interface{}) {
	newElement := NewElement(elementValue)

	if l.IsEmpty() {
		l.first = newElement
		l.last = newElement
		l.length++
		return
	}

	newElement.next = l.first
	l.first = newElement
	l.length++
}

// BulkAppend: 批量在表尾添加元素
func (l *SingleLinkList) BulkAppend(elementValues ...interface{}) {
	if l.IsEmpty() {
		newList := New(elementValues...)
		l.first = newList.first
		l.last = newList.last
		l.length = newList.length

		return
	}

	for _, elementValue := range elementValues {
		l.Append(elementValue)
	}
}

// Insert: 在表中的指定位置添加元素
func (l *SingleLinkList) Insert(elementValue interface{}, index int) {
	if index < 0 {
		return
	}

	if l.IsEmpty() {
		if index == 0 {
			l.Append(elementValue)
			return
		}

		return
	}

	if index == 0 {
		l.Prepend(elementValue)
	} else if index == l.length {
		l.Append(elementValue)
	} else {
		preElement := l.first
		for tmpIndex := 1; tmpIndex < index; tmpIndex++ {
			preElement = preElement.next
		}

		newElement := NewElement(elementValue)
		newElement.next = preElement.next
		preElement.next = newElement
		l.length++
	}
}

// DeleteFirst: 从头部删除一个元素
func (l *SingleLinkList) DeleteFirst() interface{} {
	if l.IsEmpty() {
		println("The list is empty.")
		return nil
	}

	deletedElement := l.first
	l.first = l.first.next
	l.length--

	return deletedElement.value
}

// DeleteLast: 删除标为元素
func (l *SingleLinkList) DeleteLast() interface{} {
	if l.IsEmpty() {
		println("The list is empty.")
		return nil
	}

	deletedElement := l.last
	newLast := l.first
	for index := 1; index < l.length - 1; index++ {
		newLast = newLast.next
	}
	for ; newLast.next != deletedElement; newLast = newLast.next {}
	l.last = newLast
	l.last.next = nil
	l.length--

	return deletedElement.value
}

// Delete: 删除指定位置的元素
func (l *SingleLinkList) Delete(index int) interface{} {
	if l.IsEmpty() {
		println("The list is empty.")
		return nil
	}

	if index < 0 || index >= l.length {
		println("The index is invalid.")
		return nil
	}

	if index == 0 {
		return l.DeleteFirst()
	} else if index == l.length - 1 {
		return l.DeleteLast()
	} else {
		tmpIndex := 0
		preElement := l.first
		for ; tmpIndex != index - 1; tmpIndex++ {
			preElement = preElement.next
		}
		deletedElement := preElement.next
		preElement.next = preElement.next.next
		l.length--

		return deletedElement.value
	}
}

// Clear: 删除链表所有元素
func (l *SingleLinkList) Clear() int {
	if l.IsEmpty() {
		return 0
	}

	l.first, l.last = nil, nil
	deleteCount := l.length
	l.length = 0

	return deleteCount
}

// Get: 获取指定位置的元素
// range of index: [0, length - 1]
func (l SingleLinkList) Get(index int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("the list is empty")
	} else if index < 0 || index >= l.Length() {
		return nil, errors.New("the given index is invalid")
	} else {
		if index == 0 {
			return l.first.value, nil
		} else if index == l.Length() - 1 {
			return l.last.value, nil
		} else {
			targetElement := l.first
			for tmpIndex := 0; tmpIndex < index; tmpIndex++ {
				targetElement = targetElement.next
			}
			return targetElement.value, nil
		}
	}
}

// Index: 获取链表中某元素值第一次出现的索引(如果有)
// param elementValue: 索要查找的元素的值
// return: 如果元素在链表中，则返回对应元素值的索引；否则，返回-1
func (l SingleLinkList) Index(elementValue interface{}) int {
	if l.IsEmpty() {
		println("The list is empty.")
		return -1
	}

	index := 0
	for tmpElement := l.first; tmpElement != nil; tmpElement = tmpElement.next {
		if tmpElement.value == elementValue {
			break
		}
		index++
	}

	if index >= l.Length() {
		return -1
	}

	return index
}

// Contains: 判断一个元素值是否在链表中
// param elementValue: 待查找的元素值
// return: 如果元素值在链表中则返回true, 否则返回false
func (l SingleLinkList) Contains(elementValue interface{}) bool {
	if l.Index(elementValue) == -1 {
		return false
	} else {
		return true
	}
}

// Set: 修改单链表的某一个元素的值
func (l *SingleLinkList) Set(elementValue interface{}, index int) (interface{}, error) {
	if l.IsEmpty() {
		return nil, errors.New("the list is empty")
	} else if index < 0 || index >= l.Length() {
		return nil, errors.New("the given index is invalid")
	} else {
		if index == 0 {
			oldVal := l.first.value
			l.first.value = elementValue
			return oldVal, nil
		} else if index == l.Length() - 1 {
			oldVal := l.last.value
			l.last.value = elementValue
			return oldVal, nil
		} else {
			targetElement := l.first
			for ; index != 0; index-- {
				targetElement = targetElement.next
			}
			oldVal := targetElement.value
			targetElement.value = elementValue

			return oldVal, nil
		}
	}
}

// Length: 获取单链表的长度
func (l SingleLinkList) Length() int {
	return l.length
}

// IsEmpty: 判断单链表是否为空
func (l SingleLinkList) IsEmpty() bool {
	return l.length == 0
}

