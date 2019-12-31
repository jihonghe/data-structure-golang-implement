package sequence_list

import (
	"testing"
)

func TestListNew(t *testing.T) {
	list1 := New()
	if !list1.IsEmpty() {
		t.Error("Error: sequence list is not empty.")
	}

	list2 := New(1, true, "hello", 3.14)
	if list2.Length() != 4 {
		t.Error("Error: sequence list length is invalid.")
	}
}
