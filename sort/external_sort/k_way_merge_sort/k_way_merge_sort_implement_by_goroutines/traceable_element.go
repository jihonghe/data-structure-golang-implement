package k_way_merge_sort_implement_by_goroutines

type TraceableElement struct {
	value int64
	srcArray int
}

func NewTraceableElement(value int64, srcArray int) *TraceableElement {
	return &TraceableElement{
		value:    value,
		srcArray: srcArray,
	}
}
