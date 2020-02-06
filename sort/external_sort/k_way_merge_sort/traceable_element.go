package k_way_merge_sort

type TraceableElement struct {
	value int
	elementFromArray int
}

func NewTraceableElement(elementValue int, elementFromArray int) *TraceableElement {
	return &TraceableElement{
		value: elementValue,
		elementFromArray: elementFromArray,
	}
}

func (traceableElement TraceableElement) GetElementValue() int {
	return traceableElement.value
}

func (traceableElement *TraceableElement) SetElementValue(elementValue int) {
	traceableElement.value = elementValue
}

func (traceableElement TraceableElement) GetElementFromArray() int {
	return traceableElement.elementFromArray
}

func (traceableElement *TraceableElement) SetElementFromArray(elementFromArray int) {
	traceableElement.elementFromArray = elementFromArray
}
