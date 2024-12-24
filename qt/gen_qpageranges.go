package qt

import (
	"unsafe"
)

type QPageRanges struct {
	h          uintptr
	isSubclass bool
}

// NewQPageRanges constructs a new QPageRanges object.
func NewQPageRanges() *QPageRanges {
	g := newQPageRanges(QPageRanges_new())
	g.isSubclass = true
	return g
}

// NewQPageRanges2 constructs a new QPageRanges object.
func NewQPageRanges2(other *QPageRanges) *QPageRanges {
	g := newQPageRanges(QPageRanges_new2(other.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QPageRanges) OperatorAssign(other *QPageRanges) {
	QPageRanges_OperatorAssign(this.h, other.cPointer())
}

func (this *QPageRanges) Swap(other *QPageRanges) {
	QPageRanges_Swap(this.h, other.cPointer())
}

func (this *QPageRanges) AddPage(pageNumber int) {
	QPageRanges_AddPage(this.h, (int)(pageNumber))
}

func (this *QPageRanges) AddRange(from int, to int) {
	QPageRanges_AddRange(this.h, (int)(from), (int)(to))
}

func (this *QPageRanges) ToRangeList() []Range {
	var _ma struct_miqt_array = QPageRanges_ToRangeList(this.h)
	_ret := make([]Range, int(_ma.len))
	_outCast := (*[0xffff]Range)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QPageRanges) Clear() {
	QPageRanges_Clear(this.h)
}

func (this *QPageRanges) ToString() string {
	var _ms struct_miqt_string = QPageRanges_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPageRanges_FromString(ranges string) *QPageRanges {
	ranges_ms := struct_miqt_string{}
	ranges_ms.data = CString(ranges)
	ranges_ms.len = size_t(len(ranges))
	defer free(unsafe.Pointer(ranges_ms.data))
	_goptr := newQPageRanges(QPageRanges_FromString(ranges_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPageRanges) Contains(pageNumber int) bool {
	return (bool)(QPageRanges_Contains(this.h, (int)(pageNumber)))
}

func (this *QPageRanges) IsEmpty() bool {
	return (bool)(QPageRanges_IsEmpty(this.h))
}

func (this *QPageRanges) FirstPage() int {
	return (int)(QPageRanges_FirstPage(this.h))
}

func (this *QPageRanges) LastPage() int {
	return (int)(QPageRanges_LastPage(this.h))
}

func (this *QPageRanges) Detach() {
	QPageRanges_Detach(this.h)
}

type QPageRanges__Range struct {
	h          uintptr
	isSubclass bool
}

// NewQPageRanges__Range constructs a new QPageRanges::Range object.
func NewQPageRanges__Range() *QPageRanges__Range {
	g := newQPageRanges__Range(QPageRanges__Range_new())
	g.isSubclass = true
	return g
}

// NewQPageRanges__Range2 constructs a new QPageRanges::Range object.
func NewQPageRanges__Range2(param1 *Range) *QPageRanges__Range {
	g := newQPageRanges__Range(QPageRanges__Range_new2(param1))
	g.isSubclass = true
	return g
}

func (this *QPageRanges__Range) Contains(pageNumber int) bool {
	return (bool)(QPageRanges__Range_Contains(this.h, (int)(pageNumber)))
}
