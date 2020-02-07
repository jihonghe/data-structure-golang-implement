package k_way_merge_sort

import (
	"sort"
)

const (
	baseLength = 1000
	leftIndexStr = "leftIndex"
	rightIndexStr = "rightIndex"
)

// OptimizedKWayMergeSort:
// arrayLeft和arrayRight是都是有效下标，对应的是bufferArray的待存储区间及array的待排序区间
func OptimizedKWayMergeSort(array, bufferArray []int, arrayLeft, arrayRight int) {
	arrayLength := arrayRight - arrayLeft + 1
	// 当序列长度在baseLength范围内时，采用原地排序算法，避免递归调用
	if arrayLength <= baseLength {
		copy(bufferArray[arrayLeft: arrayRight + 1], array)
		sort.Ints(bufferArray[arrayLeft: arrayRight + 1])
		return
	}

	// 根据array长度划分子序列个数
	arraysCount := GetGroupCount(arrayLength, baseLength)
	/**
	arraysRangeMap: 用于存放subArrays中每个子序列在bufferArray上对应的下标范围
	外层map的key(int)对应subArrays的下标，表示每个子序列在subArrays中的下标
	外层map的value(map[string]int)用于存放对应的子序列在bufferArray上计算出的下标范围，其key只有"leftIndex"和"rightIndex"
	注意："leftIndex"和"rightIndex"形成的区间是闭区间
	*/
	arraysRangeMap := make(map[int](map[string]int))
	// 用于存储子序列
	subArrays := make([]([]int), arraysCount, arraysCount + 1)
	for index := 0; index < arraysCount; index++ {
		subArrayLeftIndex, subArrayRightIndex := index * baseLength, (index + 1) * baseLength
		if subArrayRightIndex > arrayLength {
			subArrays[index] = array[subArrayLeftIndex: arrayLength]
			arraysRangeMap[index] = make(map[string]int)
			arraysRangeMap[index][leftIndexStr], arraysRangeMap[index][rightIndexStr] = subArrayLeftIndex, arrayLength - 1
			break
		}
		subArrays[index] = array[subArrayLeftIndex: subArrayRightIndex]
		arraysRangeMap[index] = make(map[string]int)
		arraysRangeMap[index][leftIndexStr], arraysRangeMap[index][rightIndexStr] = subArrayLeftIndex, subArrayRightIndex - 1
	}

	// 多路递归
	for i, subArray := range subArrays {
		OptimizedKWayMergeSort(subArray, bufferArray, arraysRangeMap[i][leftIndexStr], arraysRangeMap[i][rightIndexStr])
	}
	// 合并多个有序子序列
	OptimizedKWayMerge(arraysCount, subArrays, bufferArray, arraysRangeMap)
}

func OptimizedKWayMerge(k int, arrays []([]int), bufferArray []int, arraysRangeMap map[int](map[string]int)) {
	// 用arrayIndexList存放的元素的下标对应的元素构造堆，容量为k
	minHeap := NewMinHeap(k)
	// arrayIdValidIndexMap: key为序列在arrays中的下标，value为下一个待插入元素在对应的序列中的下标
	arrayIdValidIndexMap := make(map[int]int)
	arrayLengthListInArrays := make([]int, k)
	for index := 0; index < k; index++ {
		// 保证要处理的序列非空
		arrayLength := len(arrays[index])
		if arrayLength == 0 {
			continue
		}
		minHeap.AppendNode(*NewTraceableElement(arrays[index][0], index))
		arrayIdValidIndexMap[index] = 1
		// 记录每一个序列的长度
		arrayLengthListInArrays[index] = arrayLength
	}

	bufferArrayLeftIndex, bufferArrayRightIndex := arraysRangeMap[0][leftIndexStr], arraysRangeMap[k - 1][rightIndexStr] + 1
	finalArray := bufferArray[bufferArrayLeftIndex: bufferArrayRightIndex]
	for index := 0; index < bufferArrayRightIndex; {
		// 若所有序列中的元素已处理或在堆中，则将堆中剩余的节点按照顺序放入finalArray中后退出循环，完成所有元素排序
		if len(arrayIdValidIndexMap) == 0 && !minHeap.IsEmpty() {
			for !minHeap.IsEmpty() {
				finalArray[index] = minHeap.PopRoot().value
				index++
				if index >= bufferArrayRightIndex {
					break
				}
			}
			break
		}

		// 取出堆顶节点，将其插入finalArray中
		rootNodeElement := minHeap.PopRoot()
		finalArray[index] = rootNodeElement.GetElementValue()
		index++
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
}
//
//func GetGroupCount(num, baseNum int) int {
//	groupCount := 1
//	for biggerNum := baseNum * groupCount; biggerNum - num <= 0; biggerNum = baseNum * groupCount {
//		groupCount++
//	}
//
//	return groupCount
//}
//
//func TraverseIntArray(array []int) {
//	for _, value := range array {
//		fmt.Printf("%d, ", value)
//	}
//	println()
//}
//
//func GetRandomSlice(length , shuffleTime int) []int {
//	var nums []int
//	for index := 0; index < length; index++ {
//		nums = append(nums, index)
//	}
//
//	for index := 0; index < shuffleTime; index++ {
//		rand.Shuffle(length, func(i, j int) {
//			nums[i], nums[j] = nums[j], nums[i]
//		})
//	}
//
//	return nums
//}
