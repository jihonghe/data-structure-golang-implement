package sequence_list

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
	for leftIndex := l.length; leftIndex < index; {
		l.elements[leftIndex - 1], l.elements[leftIndex] = l.elements[leftIndex], l.elements[leftIndex - 1]
		leftIndex--
	}
	l.length++

	return true
}

// IsEmpty: 判断表是否为空
func (l SequenceList) IsEmpty() bool {
	return l.length == 0
}

// IsFull: 判断表是否已满
func (l SequenceList) IsFull() bool {
	return l.length == l.size
}

// Length: 获取表长度
func (l SequenceList) Length() int {
	return l.length
}
