package sequence_list

type SequenceList struct {
	// 通过切片存储元素值
	elements []interface{}
	// 顺序表元素个数
	length int
	// 顺序表的最大容量，不采用数组控制容量
	size int
}

const (
	// 顺序表的容量
	defaultSize = 10
)

func New(elements ...interface{}) *SequenceList {
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

func (l SequenceList) IsEmpty() bool {
	return l.length == 0
}

func (l SequenceList) Length() int {
	return l.length
}
