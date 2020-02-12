package k_way_merge_sort_implement_by_goroutines

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

const (
	baseArrayLength = int64(5000)
	mergeBaseLength = int64(100)
)

func MergeSort(src []int64) {
	srcLength := int64(len(src))

	// 获取由src分割排序所得的有序子数组
	start := time.Now().UnixNano()
	sortedSubArrays := getSortedSubArrays(src, srcLength, baseArrayLength)
	end := time.Now().UnixNano()
	fmt.Printf("sortedSubArraysLength: %d, time-costs: %3f\n", len(sortedSubArrays), float64(end-start)/1e9)
	// 执行多路归并
	start = time.Now().UnixNano()
	dst := make([]int64, srcLength)
	//Merge(dst, sortedSubArrays)
	dst = getMergedArray(sortedSubArrays, mergeBaseLength)
	end = time.Now().UnixNano()
	fmt.Printf("merge-time-costs: %3f\n", float64(end-start)/1e9)
	// 将结果复制到src中
	copy(src, dst)
}

func getMergedArray(sortedArrays [][]int64, mergeBaseLength int64) []int64 {
	arraysAmount := int64(len(sortedArrays))
	wg := new(sync.WaitGroup)
	var subSortedArrays [][]int64
	baseLength := mergeBaseLength
	for arraysAmount >= baseLength {
		println("newSortedArrayLength:", arraysAmount, "\tbaseLength:", baseLength)
		unitArrayLength := int64(len(sortedArrays[0]))
		// 获取子数组长度
		subArraysAmount := getGroupLen(arraysAmount, baseLength)
		subSortedArrays = make([][]int64, subArraysAmount, subArraysAmount)
		wg.Add(int(subArraysAmount))
		subLeftIndex, subRightIndex := int64(0), baseLength
		for index := 0; int64(index) < subArraysAmount; index++ {
			if subRightIndex > arraysAmount {
				subRightIndex = arraysAmount
			}
			subArrays := sortedArrays[subLeftIndex: subRightIndex]
			subLeftIndex, subRightIndex = subRightIndex, subRightIndex + baseLength
			dstLength := (subRightIndex - subLeftIndex) * unitArrayLength
			dst := make([]int64, dstLength, dstLength)
			go Merge2(dst, subArrays, wg)
			subSortedArrays[index] = dst
		}
		wg.Wait()
		sortedArrays = subSortedArrays
		arraysAmount = int64(len(sortedArrays))
		if baseLength / 10 < 10 {
			baseLength = 10
		} else {
			baseLength /= 10
		}
	}
	wg.Add(1)
	dst := make([]int64, len(sortedArrays) * len(sortedArrays[0]))
	go Merge2(dst, sortedArrays, wg)
	wg.Wait()

	return dst
}

func Merge2(dst []int64, sortedArrays [][]int64, wg *sync.WaitGroup) {
	defer wg.Done()
	k := len(sortedArrays)
	if k == 1 {
		copy(dst, sortedArrays[0])
		return
	}

	dstLength := k * len(sortedArrays[0])
	// 构建一个容量为k的最小堆
	minHeap := NewMinHeap(k)
	// 向堆中添加k个节点
	arrayIdValidIndexMap := make(map[int]int)
	arrayLengthListInArrays := make([]int, 0)
	for index := 0; index < k; index++ {
		// 保证要处理的序列非空
		arrayLength := len(sortedArrays[index])
		if arrayLength == 0 {
			continue
		}
		minHeap.AppendNode(*NewTraceableElement(sortedArrays[index][0], index))
		arrayIdValidIndexMap[index] = 1
		// 记录每一个序列的长度
		arrayLengthListInArrays = append(arrayLengthListInArrays, arrayLength)
	}
	// 合并
	start := time.Now().UnixNano()
	for index := 0; index < dstLength; {
		// 若所有序列中的元素已处理或在堆中，则将堆中剩余的节点按照顺序放入finalArray中后退出循环，完成所有元素排序
		if len(arrayIdValidIndexMap) == 0 && !minHeap.IsEmpty() {
			for !minHeap.IsEmpty() {
				dst[index] = minHeap.PopRoot().value
				index++
			}
			break
		}

		// 取出堆顶节点，将其插入finalArray中
		rootNodeElement := minHeap.PopRoot()
		dst[index] = rootNodeElement.value
		index++
		// 获取下一个待插入堆中的元素所在的序列
		arrayIdOfNextValidIndex := rootNodeElement.srcArray
		if _, ok := arrayIdValidIndexMap[arrayIdOfNextValidIndex]; !ok {
			for arrayIndex := range arrayIdValidIndexMap {
				arrayIdOfNextValidIndex = arrayIndex
				break
			}
		}
		// 补充堆节点
		minHeap.AppendNode(*NewTraceableElement(
			sortedArrays[arrayIdOfNextValidIndex][arrayIdValidIndexMap[arrayIdOfNextValidIndex]],
			arrayIdOfNextValidIndex))
		// 修改对应列表中的处理元素的下标并校验下标是否有效，若无效则从map中删除该序列
		arrayIdValidIndexMap[arrayIdOfNextValidIndex]++
		if arrayIdValidIndexMap[arrayIdOfNextValidIndex] >= arrayLengthListInArrays[arrayIdOfNextValidIndex] {
			delete(arrayIdValidIndexMap, arrayIdOfNextValidIndex)
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("merge-for-time-cost: %3f, dstLength: %d\n", float64(end - start)/1e9, len(dst))
}

func Merge(dst []int64, sortedArrays [][]int64) {
	k := len(sortedArrays)

	// 构建一个容量为k的最小堆
	minHeap := NewMinHeap(k)
	// 向堆中添加k个节点
	arrayIdValidIndexMap := make(map[int]int)
	arrayLengthListInArrays := make([]int, 0)
	for index := 0; index < k; index++ {
		// 保证要处理的序列非空
		arrayLength := len(sortedArrays[index])
		if arrayLength == 0 {
			continue
		}
		minHeap.AppendNode(*NewTraceableElement(sortedArrays[index][0], index))
		arrayIdValidIndexMap[index] = 1
		// 记录每一个序列的长度
		arrayLengthListInArrays = append(arrayLengthListInArrays, arrayLength)
	}
	// 合并
	start := time.Now().UnixNano()
	for index := 0; ; {
		// 若所有序列中的元素已处理或在堆中，则将堆中剩余的节点按照顺序放入finalArray中后退出循环，完成所有元素排序
		if len(arrayIdValidIndexMap) == 0 && !minHeap.IsEmpty() {
			for !minHeap.IsEmpty() {
				dst[index] = minHeap.PopRoot().value
				index++
			}
			break
		}

		// 取出堆顶节点，将其插入finalArray中
		rootNodeElement := minHeap.PopRoot()
		dst[index] = rootNodeElement.value
		index++
		// 获取下一个待插入堆中的元素所在的序列
		arrayIdOfNextValidIndex := rootNodeElement.srcArray
		if _, ok := arrayIdValidIndexMap[arrayIdOfNextValidIndex]; !ok {
			for arrayIndex := range arrayIdValidIndexMap {
				arrayIdOfNextValidIndex = arrayIndex
				break
			}
		}
		// 补充堆节点
		minHeap.AppendNode(*NewTraceableElement(
			sortedArrays[arrayIdOfNextValidIndex][arrayIdValidIndexMap[arrayIdOfNextValidIndex]],
			arrayIdOfNextValidIndex))
		// 修改对应列表中的处理元素的下标并校验下标是否有效，若无效则从map中删除该序列
		arrayIdValidIndexMap[arrayIdOfNextValidIndex]++
		if arrayIdValidIndexMap[arrayIdOfNextValidIndex] >= arrayLengthListInArrays[arrayIdOfNextValidIndex] {
			delete(arrayIdValidIndexMap, arrayIdOfNextValidIndex)
		}
	}
	end := time.Now().UnixNano()
	fmt.Printf("merge-for-time-cost: %3f, dstLength: %d\n", float64(end - start)/1e9, len(dst))
}

func getSortedSubArrays(array []int64, arrayLength, subArrayLength int64) [][]int64 {
	wg := new(sync.WaitGroup)
	// 获取子数组长度
	subArrayAmount := getGroupLen(arrayLength, subArrayLength)
	// 获取有序子数组
	subArrays := make([][]int64, subArrayAmount, subArrayAmount)
	subLeftIndex, subRightIndex := int64(0), subArrayLength
	for index := 0; int64(index) < subArrayAmount; index++ {
		wg.Add(1)
		if subRightIndex > arrayLength {
			subRightIndex = arrayLength
		}
		subArray := array[subLeftIndex: subRightIndex]
		subArrays[index] = subArray
		go inMemorySort(subArray, wg)
		subLeftIndex, subRightIndex = subRightIndex, subRightIndex + subArrayLength
	}
	wg.Wait()

	return subArrays
}

func inMemorySort(array []int64, wg *sync.WaitGroup) {
	defer wg.Done()
	int64Array := Int64Array(array)
	sort.Sort(int64Array)
}

func getGroupLen(totalLength, baseLength int64) int64 {
	if totalLength % baseLength == 0 {
		return totalLength / baseLength
	} else {
		return totalLength / baseLength + 1
	}
}

