package sequence_list

import (
	"errors"
	"fmt"
)

	type SequenceList struct {
	// 通过切片存储元素值
	elements []int
	// 顺序表元素个数
	length int
	// 顺序表的最大容量，不采用数组控制容量
	size int
}

const (
	// 顺序表的容量
	defaultSize = 10
)

// New: 创建并初始化顺序表
func New(elements ...int) *SequenceList {
	list := &SequenceList{}
	list.size = defaultSize

	elementsLength := len(elements)
	if elementsLength < list.size && elementsLength > 0 {
		for _, element := range elements {
			list.elements = append(list.elements, element)
		}
	}
	list.length = elementsLength

	return list
}

// Append: 添加元素。添加成功则返回元素位置，否则返回-1
func (l *SequenceList) Append(element int) int {
	if l.IsFull() {
		println("Append failed: the list is full.")
		return -1
	}

	index := l.length
	l.elements = append(l.elements, element)
	l.length++

	return index
}

// Insert: 按指定位置插入元素。插入成功则返回true，否则返回false
func (l *SequenceList) Insert(elementValue int, index int) bool {
	if l.IsFull() {
		println("Insert failed: the list is full.")

		return false
	}

	if index < 0 || index > l.length {
		println("Insert failed: index is invalid, the current length is", l.length)

		return false
	}

	if index == l.length {
		l.elements = append(l.elements, elementValue)
		l.length++

		return true
	}

	l.elements = append(l.elements, elementValue)
	copy(l.elements[index+1:], l.elements[index:l.length])
	l.elements[index] = elementValue
	l.length++

	return true
}

// DeleteLast: 删除表尾的一个元素
func (l *SequenceList) DeleteLast() bool {
	if l.IsEmpty() {
		println("DeleteLast failed: the list is empty.")
		return false
	}

	l.elements = l.elements[:l.length-1]
	l.length--

	return true
}

// Delete: 删除表中指定位置的元素
func (l *SequenceList) Delete(index int) bool {
	if l.IsEmpty() {
		println("Delete failed: the list is empty.")
		return false
	}

	if index < 0 || index >= l.length {
		println("Delete failed: the given index is invalid.")
		return false
	}

	copy(l.elements[index:], l.elements[index+1:])
	l.elements = l.elements[:l.length-1]
	l.length--

	return true
}

// Clear: 清除表中的所有元素
func (l *SequenceList) Clear() bool {
	if l.IsEmpty() {
		return true
	}

	l.elements = l.elements[:0]
	l.length = 0

	return true
}

// Set: 修改指定位置的元素的值
func (l *SequenceList) Set(elementValue int, index int) bool {
	if l.IsEmpty() {
		println("Set failed: the list is empty.")
		return false
	}
	if index < 0 || index >= l.length {
		println("Set failed: the given index is invalid.")
		return false
	}

	l.elements[index] = elementValue

	return true
}

// Get: 获取指定位置的元素值
func (l SequenceList) Get(index int) (int, error) {
	if l.IsEmpty() {
		return 0, errors.New("Get failed: the list is empty")
	}

	if index < 0 || index >= l.length {
		return 0, errors.New("Get failed: the given index is invalid")
	}

	return l.elements[index], nil
}

// Index: 获取给定值在表中第一次出现的位置
func (l SequenceList) Index(elementValue int) int {
	if l.IsEmpty() {
		return -1
	}

	for index, element := range l.elements {
		if elementValue == element {
			return index
		}
	}

	return -1
}

// Contains: 判断给出的元素值是否在表中
func (l SequenceList) Contains(elementValue int) bool {
	if l.IsEmpty() {
		return false
	}

	for _, element := range l.elements {
		if element == elementValue {
			return true
		}
	}

	return false
}

// Traverse: 遍历顺序表元素
func (l SequenceList) Traverse() {
	if l.IsEmpty() {
		return
	}

	print("Sequence list: ")
	elementsListString := "["
	for _, element := range l.elements {
		elementsListString += fmt.Sprintf("%d, ", element)
	}
	elementsListString += "]"
}

// IsEmpty: 判断表是否为空
func (l SequenceList) IsEmpty() bool {
	return (l.length == len(l.elements)) && (len(l.elements) == 0)
}

// IsFull: 判断表是否已满
func (l SequenceList) IsFull() bool {
	return l.length == l.size
}

// Length: 获取表长度
func (l SequenceList) Length() int {
	return l.length
}
