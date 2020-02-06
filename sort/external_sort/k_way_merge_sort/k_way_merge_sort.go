package k_way_merge_sort

import (
	"fmt"
	"math/rand"
	"sort"
)

const (
	baseArrayLength = 1000
)

func KWayMergeSort(array []int) []int {
	arrayLength := len(array)
	if arrayLength <= baseArrayLength {
		sort.Ints(array)
		return array
	}
	arraysCount := GetGroupCount(arrayLength, baseArrayLength)

	var subArrays []([]int)
	for index := 1; index <= arraysCount; index++ {
		if index * baseArrayLength > arrayLength {
			subArray := array[(index - 1) * baseArrayLength: arrayLength]
			subArrays = append(subArrays, subArray)
			break
		}
		subArray := array[(index - 1) * baseArrayLength: index * baseArrayLength]
		subArrays = append(subArrays, subArray)
	}

	for i, subArray := range subArrays {
		subArrays[i] = KWayMergeSort(subArray)
	}
	return KWayMerge(arraysCount, subArrays)
}

func KWayMerge(k int, arrays []([]int)) []int {
	// 用arrayIndexList存放的元素的下标对应的元素构造堆，容量为k
	minHeap := NewMinHeap(k)
	// arrayIdValidIndexMap: key为序列在arrays中的下标，value为下一个待插入元素在对应的序列中的下标
	arrayIdValidIndexMap := make(map[int]int)
	arrayLengthListInArrays := make([]int, 0)
	for index := 0; index < k; index++ {
		// 保证要处理的序列非空
		arrayLength := len(arrays[index])
		if arrayLength == 0 {
			continue
		}
		minHeap.AppendNode(*NewTraceableElement(arrays[index][0], index))
		arrayIdValidIndexMap[index] = 1
		// 记录每一个序列的长度
		arrayLengthListInArrays = append(arrayLengthListInArrays, arrayLength)
	}

	finalArray := make([]int, 0)
	for {
		// 若所有序列中的元素已处理或在堆中，则将堆中剩余的节点按照顺序放入finalArray中后退出循环，完成所有元素排序
		if len(arrayIdValidIndexMap) == 0 && !minHeap.IsEmpty() {
			for !minHeap.IsEmpty() {
				finalArray = append(finalArray, minHeap.PopRoot().value)
			}
			break
		}

		// 取出堆顶节点，将其插入finalArray中
		rootNodeElement := minHeap.PopRoot()
		finalArray = append(finalArray, rootNodeElement.GetElementValue())
		// 获取下一个待插入堆中的元素所在的序列
		arrayIdOfNextValidIndex := rootNodeElement.GetElementFromArray()
		if _, ok := arrayIdValidIndexMap[arrayIdOfNextValidIndex]; !ok {
			for arrayIndex, _ := range arrayIdValidIndexMap {
				arrayIdOfNextValidIndex = arrayIndex
				break
			}
		}
		// 补充堆节点
		minHeap.AppendNode(*NewTraceableElement(
			arrays[arrayIdOfNextValidIndex][arrayIdValidIndexMap[arrayIdOfNextValidIndex]],
			arrayIdOfNextValidIndex))
		// 修改对应列表中的处理元素的下标并校验下标是否有效，若无效则从map中删除该序列
		arrayIdValidIndexMap[arrayIdOfNextValidIndex]++
		if arrayIdValidIndexMap[arrayIdOfNextValidIndex] >= arrayLengthListInArrays[arrayIdOfNextValidIndex] {
			delete(arrayIdValidIndexMap, arrayIdOfNextValidIndex)
		}
	}

	return finalArray
}

func GetGroupCount(num, baseNum int) int {
	groupCount := 1
	for biggerNum := baseNum * groupCount; biggerNum - num <= 0; biggerNum = baseNum * groupCount {
		groupCount++
	}

	return groupCount
}

func TraverseIntArray(array []int) {
	for _, value := range array {
		fmt.Printf("%d, ", value)
	}
	println()
}

func GetRandomSlice(length , shuffleTime int) []int {
	var nums []int
	for index := 0; index < length; index++ {
		nums = append(nums, index)
	}

	for index := 0; index < shuffleTime; index++ {
		rand.Shuffle(length, func(i, j int) {
			nums[i], nums[j] = nums[j], nums[i]
		})
	}

	return nums
}
