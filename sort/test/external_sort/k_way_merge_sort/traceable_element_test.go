package k_way_merge_sort_test

import (
	"data-structure-golang-implement/sort/external_sort/k_way_merge_sort"
	"testing"
)

func TestNewTraceableElement(t *testing.T) {
	element := k_way_merge_sort.NewTraceableElement(9, 1)

	if element.GetElementValue() != 9  || element.GetElementFromArray() != 1 {
		t.Error("k_way_merge_sort.NewTraceableElement() error: initialize component failed.")
	}
}

func TestSetElementValue(t *testing.T) {
	element := k_way_merge_sort.NewTraceableElement(9, 1)

	element.SetElementValue(3)
	if element.GetElementValue() != 3 {
		t.Error("k_way_merge_sort.NewTraceableElement()　error: set element value failed.")
	}
}

func TestSetElementFromArray(t *testing.T) {
	element := k_way_merge_sort.NewTraceableElement(9, 1)

	element.SetElementFromArray(4)
	if element.GetElementFromArray() != 4 {
		t.Error("k_way_merge_sort.NewTraceableElement() error: set elementFromArray failed.")
	}
}
