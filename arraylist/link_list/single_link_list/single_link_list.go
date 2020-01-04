package single_link_list

import "fmt"

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

// IsEmpty: 判断单链表是否为空
func (l SingleLinkList) IsEmpty() bool {
	return l.length == 0
}
