package optimized_goroutine_k_way_merge_sort

import (
    "runtime"
    "sort"
    "sync"
)

// MergeSort performs the merge sort algorithm.
// Please supplement this function to accomplish the home work.
func MergeSort(src []int64) {
    srcLength := int64(len(src))
    
    // 获取由src分割排序所得的有序子数组
    sortedSubArrays := getSortedSubArrays(src)
    // 执行多路归并
    dst := make([]int64, srcLength)
    Merge(dst, sortedSubArrays)
    // 将结果复制到src中
    copy(src, dst)
}

func Merge(dst []int64, sortedArrays [][]int64) {
    k := len(sortedArrays)
    if k == 1 {
        copy(dst, sortedArrays[0])
        return
    }
    
    // 构建一个容量为k的最小堆
    minHeap := NewMinHeap(k)
    // 记录存在有效待处理元素的sortedArrays中的子数组的下标及子数组中的待处理的元素的下标
    arrayIdValidIndexMap := make(map[int]int)
    // 存储sortedArrays中的子数组的长度，用于校验arrayIdValidIndexMap中的子数组中的下标的有效性
    arrayLengthListInArrays := make([]int, k, k)
    // 向堆中添加k个节点
    for index := 0; index < k; index++ {
        // 保证要处理的序列非空
        arrayLength := len(sortedArrays[index])
        if arrayLength == 0 {
            continue
        }
        minHeap.AppendNode(sortedArrays[index][0], index)
        arrayIdValidIndexMap[index] = 1
        // 记录每一个序列的长度
        arrayLengthListInArrays[index] = arrayLength
    }
    // 合并
    dstLength := k * len(sortedArrays[0])
    for index := 0; index < dstLength; {
        // 若所有序列中的元素已处理或在堆中，则将堆中剩余的节点按照顺序放入finalArray中后退出循环，完成所有元素排序
        if len(arrayIdValidIndexMap) == 0 && !minHeap.IsEmpty() {
            for !minHeap.IsEmpty() {
                dst[index], _ = minHeap.PopRoot()
                index++
            }
            break
        }
        
        // 取出堆顶节点及其对应的序列id，将其插入finalArray中
        arrayIdOfNextValidIndex := minHeap.elementsArrayIds[0]
        rootNodeElement, _ := minHeap.PopRoot()
        dst[index] = rootNodeElement
        index++
        // 获取下一个待插入堆中的元素所在的序列
        if _, ok := arrayIdValidIndexMap[arrayIdOfNextValidIndex]; !ok {
            for arrayIndex := range arrayIdValidIndexMap {
                arrayIdOfNextValidIndex = arrayIndex
                break
            }
        }
        // 补充堆节点
        minHeap.AppendNode(
            sortedArrays[arrayIdOfNextValidIndex][arrayIdValidIndexMap[arrayIdOfNextValidIndex]],
            arrayIdOfNextValidIndex)
        // 修改对应列表中的处理元素的下标并校验下标是否有效，若无效则从map中删除该序列
        arrayIdValidIndexMap[arrayIdOfNextValidIndex]++
        if arrayIdValidIndexMap[arrayIdOfNextValidIndex] >= arrayLengthListInArrays[arrayIdOfNextValidIndex] {
            delete(arrayIdValidIndexMap, arrayIdOfNextValidIndex)
        }
    }
}

func getSortedSubArrays(array []int64) [][]int64 {
    arrayLength := len(array)
    // 以cpu核数作为子数组的个数
    subArrayAmount := runtime.NumCPU()
    // 获取每个子数组的最小长度(最后一个子数组需要将余下的数连接在一起)
    subArrayLength := arrayLength / subArrayAmount
    // 存储子数组
    subArrays := make([][]int64, subArrayAmount, subArrayAmount)
    // 设置等待组
    wg := new(sync.WaitGroup)
    wg.Add(subArrayAmount)
    
    subLeftIndex, subRightIndex := 0, subArrayLength
    for index := 0; index < subArrayAmount; index++ {
        if subRightIndex + subArrayLength > arrayLength {
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
    sort.Slice(array, func(i, j int) bool {
        return array[i] < array[j]
    })
}
