package internal_sort

/**
	在内部排序中，我们常说的归并排序叫二路归并排序。
 */

func MergeSort(array []int) []int {
	if len(array) == 1 {
		return array
	}

	arrayLength := len(array)
	middleIndex := arrayLength / 2

	leftArray := MergeSort(array[0: middleIndex])
	rightArray := MergeSort(array[middleIndex: arrayLength])

	return merge(leftArray, rightArray)
}

func merge(array1, array2 []int) []int {
	array1Index, array1Length := 0, len(array1)
	array2Index, array2Length := 0, len(array2)
	sortedArray := make([]int, 0)

	for array1Index < array1Length && array2Index < array2Length {
		if array1[array1Index] <= array2[array2Index] {
			sortedArray = append(sortedArray, array1[array1Index])
			array1Index++
		} else {
			sortedArray = append(sortedArray, array2[array2Index])
			array2Index++
		}
	}
	if array1Index < array1Length {
		sortedArray = append(sortedArray, array1[array1Index:]...)
	}
	if array2Index < array2Length {
		sortedArray = append(sortedArray, array2[array2Index:]...)
	}


	return sortedArray
}
