package qt

import (
	"unsafe"
)

type QHeaderView__ResizeMode int

const (
	QHeaderView__Interactive      QHeaderView__ResizeMode = 0
	QHeaderView__Stretch          QHeaderView__ResizeMode = 1
	QHeaderView__Fixed            QHeaderView__ResizeMode = 2
	QHeaderView__ResizeToContents QHeaderView__ResizeMode = 3
	QHeaderView__Custom           QHeaderView__ResizeMode = 2
)

type QHeaderView struct {
	h          uintptr
	isSubclass bool
}

// NewQHeaderView constructs a new QHeaderView object.
func NewQHeaderView(orientation Orientation) *QHeaderView {

	ret := newQHeaderView(QHeaderView_new((int)(orientation)))
	ret.isSubclass = true
	return ret
}

// NewQHeaderView2 constructs a new QHeaderView object.
func NewQHeaderView2(orientation Orientation, parent *QWidget) *QHeaderView {

	ret := newQHeaderView(QHeaderView_new2((int)(orientation), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QHeaderView) MetaObject() *QMetaObject {
	return newQMetaObject(QHeaderView_MetaObject(this.h))
}

func (this *QHeaderView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QHeaderView_Metacast(this.h, param1_Cstring))
}

func QHeaderView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QHeaderView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHeaderView) SetModel(model *QAbstractItemModel) {
	QHeaderView_SetModel(this.h, model.cPointer())
}

func (this *QHeaderView) Orientation() Orientation {
	return (Orientation)(QHeaderView_Orientation(this.h))
}

func (this *QHeaderView) Offset() int {
	return (int)(QHeaderView_Offset(this.h))
}

func (this *QHeaderView) Length() int {
	return (int)(QHeaderView_Length(this.h))
}

func (this *QHeaderView) SizeHint() *QSize {
	_goptr := newQSize(QHeaderView_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QHeaderView) SetVisible(v bool) {
	QHeaderView_SetVisible(this.h, (bool)(v))
}

func (this *QHeaderView) SectionSizeHint(logicalIndex int) int {
	return (int)(QHeaderView_SectionSizeHint(this.h, (int)(logicalIndex)))
}

func (this *QHeaderView) VisualIndexAt(position int) int {
	return (int)(QHeaderView_VisualIndexAt(this.h, (int)(position)))
}

func (this *QHeaderView) LogicalIndexAt(position int) int {
	return (int)(QHeaderView_LogicalIndexAt(this.h, (int)(position)))
}

func (this *QHeaderView) LogicalIndexAt2(x int, y int) int {
	return (int)(QHeaderView_LogicalIndexAt2(this.h, (int)(x), (int)(y)))
}

func (this *QHeaderView) LogicalIndexAtWithPos(pos *QPoint) int {
	return (int)(QHeaderView_LogicalIndexAtWithPos(this.h, pos.cPointer()))
}

func (this *QHeaderView) SectionSize(logicalIndex int) int {
	return (int)(QHeaderView_SectionSize(this.h, (int)(logicalIndex)))
}

func (this *QHeaderView) SectionPosition(logicalIndex int) int {
	return (int)(QHeaderView_SectionPosition(this.h, (int)(logicalIndex)))
}

func (this *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	return (int)(QHeaderView_SectionViewportPosition(this.h, (int)(logicalIndex)))
}

func (this *QHeaderView) MoveSection(from int, to int) {
	QHeaderView_MoveSection(this.h, (int)(from), (int)(to))
}

func (this *QHeaderView) SwapSections(first int, second int) {
	QHeaderView_SwapSections(this.h, (int)(first), (int)(second))
}

func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	QHeaderView_ResizeSection(this.h, (int)(logicalIndex), (int)(size))
}

func (this *QHeaderView) ResizeSections(mode QHeaderView__ResizeMode) {
	QHeaderView_ResizeSections(this.h, (int)(mode))
}

func (this *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	return (bool)(QHeaderView_IsSectionHidden(this.h, (int)(logicalIndex)))
}

func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	QHeaderView_SetSectionHidden(this.h, (int)(logicalIndex), (bool)(hide))
}

func (this *QHeaderView) HiddenSectionCount() int {
	return (int)(QHeaderView_HiddenSectionCount(this.h))
}

func (this *QHeaderView) HideSection(logicalIndex int) {
	QHeaderView_HideSection(this.h, (int)(logicalIndex))
}

func (this *QHeaderView) ShowSection(logicalIndex int) {
	QHeaderView_ShowSection(this.h, (int)(logicalIndex))
}

func (this *QHeaderView) Count() int {
	return (int)(QHeaderView_Count(this.h))
}

func (this *QHeaderView) VisualIndex(logicalIndex int) int {
	return (int)(QHeaderView_VisualIndex(this.h, (int)(logicalIndex)))
}

func (this *QHeaderView) LogicalIndex(visualIndex int) int {
	return (int)(QHeaderView_LogicalIndex(this.h, (int)(visualIndex)))
}

func (this *QHeaderView) SetSectionsMovable(movable bool) {
	QHeaderView_SetSectionsMovable(this.h, (bool)(movable))
}

func (this *QHeaderView) SectionsMovable() bool {
	return (bool)(QHeaderView_SectionsMovable(this.h))
}

func (this *QHeaderView) SetFirstSectionMovable(movable bool) {
	QHeaderView_SetFirstSectionMovable(this.h, (bool)(movable))
}

func (this *QHeaderView) IsFirstSectionMovable() bool {
	return (bool)(QHeaderView_IsFirstSectionMovable(this.h))
}

func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	QHeaderView_SetSectionsClickable(this.h, (bool)(clickable))
}

func (this *QHeaderView) SectionsClickable() bool {
	return (bool)(QHeaderView_SectionsClickable(this.h))
}

func (this *QHeaderView) SetHighlightSections(highlight bool) {
	QHeaderView_SetHighlightSections(this.h, (bool)(highlight))
}

func (this *QHeaderView) HighlightSections() bool {
	return (bool)(QHeaderView_HighlightSections(this.h))
}

func (this *QHeaderView) SectionResizeMode(logicalIndex int) ResizeMode {
	xxxxxxxxx
}

func (this *QHeaderView) SetSectionResizeMode(mode ResizeMode) {
	QHeaderView_SetSectionResizeMode(this.h, mode)
}

func (this *QHeaderView) SetSectionResizeMode2(logicalIndex int, mode ResizeMode) {
	QHeaderView_SetSectionResizeMode2(this.h, (int)(logicalIndex), mode)
}

func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	QHeaderView_SetResizeContentsPrecision(this.h, (int)(precision))
}

func (this *QHeaderView) ResizeContentsPrecision() int {
	return (int)(QHeaderView_ResizeContentsPrecision(this.h))
}

func (this *QHeaderView) StretchSectionCount() int {
	return (int)(QHeaderView_StretchSectionCount(this.h))
}

func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	QHeaderView_SetSortIndicatorShown(this.h, (bool)(show))
}

func (this *QHeaderView) IsSortIndicatorShown() bool {
	return (bool)(QHeaderView_IsSortIndicatorShown(this.h))
}

func (this *QHeaderView) SetSortIndicator(logicalIndex int, order SortOrder) {
	QHeaderView_SetSortIndicator(this.h, (int)(logicalIndex), (int)(order))
}

func (this *QHeaderView) SortIndicatorSection() int {
	return (int)(QHeaderView_SortIndicatorSection(this.h))
}

func (this *QHeaderView) SortIndicatorOrder() SortOrder {
	return (SortOrder)(QHeaderView_SortIndicatorOrder(this.h))
}

func (this *QHeaderView) SetSortIndicatorClearable(clearable bool) {
	QHeaderView_SetSortIndicatorClearable(this.h, (bool)(clearable))
}

func (this *QHeaderView) IsSortIndicatorClearable() bool {
	return (bool)(QHeaderView_IsSortIndicatorClearable(this.h))
}

func (this *QHeaderView) StretchLastSection() bool {
	return (bool)(QHeaderView_StretchLastSection(this.h))
}

func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	QHeaderView_SetStretchLastSection(this.h, (bool)(stretch))
}

func (this *QHeaderView) CascadingSectionResizes() bool {
	return (bool)(QHeaderView_CascadingSectionResizes(this.h))
}

func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	QHeaderView_SetCascadingSectionResizes(this.h, (bool)(enable))
}

func (this *QHeaderView) DefaultSectionSize() int {
	return (int)(QHeaderView_DefaultSectionSize(this.h))
}

func (this *QHeaderView) SetDefaultSectionSize(size int) {
	QHeaderView_SetDefaultSectionSize(this.h, (int)(size))
}

func (this *QHeaderView) ResetDefaultSectionSize() {
	QHeaderView_ResetDefaultSectionSize(this.h)
}

func (this *QHeaderView) MinimumSectionSize() int {
	return (int)(QHeaderView_MinimumSectionSize(this.h))
}

func (this *QHeaderView) SetMinimumSectionSize(size int) {
	QHeaderView_SetMinimumSectionSize(this.h, (int)(size))
}

func (this *QHeaderView) MaximumSectionSize() int {
	return (int)(QHeaderView_MaximumSectionSize(this.h))
}

func (this *QHeaderView) SetMaximumSectionSize(size int) {
	QHeaderView_SetMaximumSectionSize(this.h, (int)(size))
}

func (this *QHeaderView) DefaultAlignment() AlignmentFlag {
	return (AlignmentFlag)(QHeaderView_DefaultAlignment(this.h))
}

func (this *QHeaderView) SetDefaultAlignment(alignment AlignmentFlag) {
	QHeaderView_SetDefaultAlignment(this.h, (int)(alignment))
}

func (this *QHeaderView) DoItemsLayout() {
	QHeaderView_DoItemsLayout(this.h)
}

func (this *QHeaderView) SectionsMoved() bool {
	return (bool)(QHeaderView_SectionsMoved(this.h))
}

func (this *QHeaderView) SectionsHidden() bool {
	return (bool)(QHeaderView_SectionsHidden(this.h))
}

func (this *QHeaderView) SaveState() []byte {
	var _bytearray struct_miqt_string = QHeaderView_SaveState(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QHeaderView) RestoreState(state []byte) bool {
	state_alias := struct_miqt_string{}
	state_alias.data = (char)(unsafe.Pointer(&state[0]))
	state_alias.len = size_t(len(state))
	return (bool)(QHeaderView_RestoreState(this.h, state_alias))
}

func (this *QHeaderView) Reset() {
	QHeaderView_Reset(this.h)
}

func (this *QHeaderView) SetOffset(offset int) {
	QHeaderView_SetOffset(this.h, (int)(offset))
}

func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	QHeaderView_SetOffsetToSectionPosition(this.h, (int)(visualIndex))
}

func (this *QHeaderView) SetOffsetToLastSection() {
	QHeaderView_SetOffsetToLastSection(this.h)
}

func (this *QHeaderView) HeaderDataChanged(orientation Orientation, logicalFirst int, logicalLast int) {
	QHeaderView_HeaderDataChanged(this.h, (int)(orientation), (int)(logicalFirst), (int)(logicalLast))
}

func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	QHeaderView_SectionMoved(this.h, (int)(logicalIndex), (int)(oldVisualIndex), (int)(newVisualIndex))
}
func (this *QHeaderView) OnSectionMoved(slot func(logicalIndex int, oldVisualIndex int, newVisualIndex int)) {
	QHeaderView_connect_SectionMoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionMoved
func miqt_exec_callback_QHeaderView_SectionMoved(cb intptr_t, logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int, oldVisualIndex int, newVisualIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	slotval2 := (int)(oldVisualIndex)

	slotval3 := (int)(newVisualIndex)

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	QHeaderView_SectionResized(this.h, (int)(logicalIndex), (int)(oldSize), (int)(newSize))
}
func (this *QHeaderView) OnSectionResized(slot func(logicalIndex int, oldSize int, newSize int)) {
	QHeaderView_connect_SectionResized(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionResized
func miqt_exec_callback_QHeaderView_SectionResized(cb intptr_t, logicalIndex int, oldSize int, newSize int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int, oldSize int, newSize int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	slotval2 := (int)(oldSize)

	slotval3 := (int)(newSize)

	gofunc(slotval1, slotval2, slotval3)
}

func (this *QHeaderView) SectionPressed(logicalIndex int) {
	QHeaderView_SectionPressed(this.h, (int)(logicalIndex))
}
func (this *QHeaderView) OnSectionPressed(slot func(logicalIndex int)) {
	QHeaderView_connect_SectionPressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionPressed
func miqt_exec_callback_QHeaderView_SectionPressed(cb intptr_t, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	gofunc(slotval1)
}

func (this *QHeaderView) SectionClicked(logicalIndex int) {
	QHeaderView_SectionClicked(this.h, (int)(logicalIndex))
}
func (this *QHeaderView) OnSectionClicked(slot func(logicalIndex int)) {
	QHeaderView_connect_SectionClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionClicked
func miqt_exec_callback_QHeaderView_SectionClicked(cb intptr_t, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	gofunc(slotval1)
}

func (this *QHeaderView) SectionEntered(logicalIndex int) {
	QHeaderView_SectionEntered(this.h, (int)(logicalIndex))
}
func (this *QHeaderView) OnSectionEntered(slot func(logicalIndex int)) {
	QHeaderView_connect_SectionEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionEntered
func miqt_exec_callback_QHeaderView_SectionEntered(cb intptr_t, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	gofunc(slotval1)
}

func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	QHeaderView_SectionDoubleClicked(this.h, (int)(logicalIndex))
}
func (this *QHeaderView) OnSectionDoubleClicked(slot func(logicalIndex int)) {
	QHeaderView_connect_SectionDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionDoubleClicked
func miqt_exec_callback_QHeaderView_SectionDoubleClicked(cb intptr_t, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	gofunc(slotval1)
}

func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	QHeaderView_SectionCountChanged(this.h, (int)(oldCount), (int)(newCount))
}
func (this *QHeaderView) OnSectionCountChanged(slot func(oldCount int, newCount int)) {
	QHeaderView_connect_SectionCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionCountChanged
func miqt_exec_callback_QHeaderView_SectionCountChanged(cb intptr_t, oldCount int, newCount int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(oldCount int, newCount int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(oldCount)

	slotval2 := (int)(newCount)

	gofunc(slotval1, slotval2)
}

func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	QHeaderView_SectionHandleDoubleClicked(this.h, (int)(logicalIndex))
}
func (this *QHeaderView) OnSectionHandleDoubleClicked(slot func(logicalIndex int)) {
	QHeaderView_connect_SectionHandleDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionHandleDoubleClicked
func miqt_exec_callback_QHeaderView_SectionHandleDoubleClicked(cb intptr_t, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	gofunc(slotval1)
}

func (this *QHeaderView) GeometriesChanged() {
	QHeaderView_GeometriesChanged(this.h)
}
func (this *QHeaderView) OnGeometriesChanged(slot func()) {
	QHeaderView_connect_GeometriesChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_GeometriesChanged
func miqt_exec_callback_QHeaderView_GeometriesChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order SortOrder) {
	QHeaderView_SortIndicatorChanged(this.h, (int)(logicalIndex), (int)(order))
}
func (this *QHeaderView) OnSortIndicatorChanged(slot func(logicalIndex int, order SortOrder)) {
	QHeaderView_connect_SortIndicatorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SortIndicatorChanged
func miqt_exec_callback_QHeaderView_SortIndicatorChanged(cb intptr_t, logicalIndex int, order int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(logicalIndex int, order SortOrder))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	slotval2 := (SortOrder)(order)

	gofunc(slotval1, slotval2)
}

func (this *QHeaderView) SortIndicatorClearableChanged(clearable bool) {
	QHeaderView_SortIndicatorClearableChanged(this.h, (bool)(clearable))
}
func (this *QHeaderView) OnSortIndicatorClearableChanged(slot func(clearable bool)) {
	QHeaderView_connect_SortIndicatorClearableChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SortIndicatorClearableChanged
func miqt_exec_callback_QHeaderView_SortIndicatorClearableChanged(cb intptr_t, clearable bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(clearable bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(clearable)

	gofunc(slotval1)
}

func QHeaderView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHeaderView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QHeaderView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QHeaderView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QHeaderView) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QHeaderView_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QHeaderView) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SetModel
func miqt_exec_callback_QHeaderView_SetModel(self QHeaderView, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QHeaderView{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QHeaderView) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QHeaderView_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SizeHint
func miqt_exec_callback_QHeaderView_SizeHint(self QHeaderView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_SetVisible(v bool) {

	QHeaderView_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(v))

}
func (this *QHeaderView) OnSetVisible(slot func(super func(v bool), v bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SetVisible
func miqt_exec_callback_QHeaderView_SetVisible(self QHeaderView, cb intptr_t, v bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(v bool), v bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(v)

	gofunc((&QHeaderView{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QHeaderView) callVirtualBase_DoItemsLayout() {

	QHeaderView_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QHeaderView) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_DoItemsLayout
func miqt_exec_callback_QHeaderView_DoItemsLayout(self QHeaderView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHeaderView{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QHeaderView) callVirtualBase_Reset() {

	QHeaderView_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QHeaderView) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_Reset
func miqt_exec_callback_QHeaderView_Reset(self QHeaderView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHeaderView{h: self}).callVirtualBase_Reset)

}

func (this *QHeaderView) callVirtualBase_CurrentChanged(current *QModelIndex, old *QModelIndex) {

	QHeaderView_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), old.cPointer())

}
func (this *QHeaderView) OnCurrentChanged(slot func(super func(current *QModelIndex, old *QModelIndex), current *QModelIndex, old *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_CurrentChanged
func miqt_exec_callback_QHeaderView_CurrentChanged(self QHeaderView, cb intptr_t, current *QModelIndex, old *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, old *QModelIndex), current *QModelIndex, old *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(old)

	gofunc((&QHeaderView{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QHeaderView_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QHeaderView) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_Event
func miqt_exec_callback_QHeaderView_Event(self QHeaderView, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QHeaderView_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QHeaderView) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_PaintEvent
func miqt_exec_callback_QHeaderView_PaintEvent(self QHeaderView, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QHeaderView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	QHeaderView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QHeaderView) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_MousePressEvent
func miqt_exec_callback_QHeaderView_MousePressEvent(self QHeaderView, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QHeaderView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_MouseMoveEvent(e *QMouseEvent) {

	QHeaderView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QHeaderView) OnMouseMoveEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_MouseMoveEvent
func miqt_exec_callback_QHeaderView_MouseMoveEvent(self QHeaderView, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QHeaderView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QHeaderView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QHeaderView) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_MouseReleaseEvent
func miqt_exec_callback_QHeaderView_MouseReleaseEvent(self QHeaderView, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QHeaderView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_MouseDoubleClickEvent(e *QMouseEvent) {

	QHeaderView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QHeaderView) OnMouseDoubleClickEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_MouseDoubleClickEvent
func miqt_exec_callback_QHeaderView_MouseDoubleClickEvent(self QHeaderView, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QHeaderView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_ViewportEvent(e *QEvent) bool {

	return (bool)(QHeaderView_virtualbase_ViewportEvent(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QHeaderView) OnViewportEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_ViewportEvent
func miqt_exec_callback_QHeaderView_ViewportEvent(self QHeaderView, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_PaintSection(painter *QPainter, rect *QRect, logicalIndex int) {

	QHeaderView_virtualbase_PaintSection(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), (int)(logicalIndex))

}
func (this *QHeaderView) OnPaintSection(slot func(super func(painter *QPainter, rect *QRect, logicalIndex int), painter *QPainter, rect *QRect, logicalIndex int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_PaintSection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_PaintSection
func miqt_exec_callback_QHeaderView_PaintSection(self QHeaderView, cb intptr_t, painter *QPainter, rect *QRect, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, logicalIndex int), painter *QPainter, rect *QRect, logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	slotval3 := (int)(logicalIndex)

	gofunc((&QHeaderView{h: self}).callVirtualBase_PaintSection, slotval1, slotval2, slotval3)

}

func (this *QHeaderView) callVirtualBase_SectionSizeFromContents(logicalIndex int) *QSize {

	_goptr := newQSize(QHeaderView_virtualbase_SectionSizeFromContents(unsafe.Pointer(this.h), (int)(logicalIndex)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnSectionSizeFromContents(slot func(super func(logicalIndex int) *QSize, logicalIndex int) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SectionSizeFromContents(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SectionSizeFromContents
func miqt_exec_callback_QHeaderView_SectionSizeFromContents(self QHeaderView, cb intptr_t, logicalIndex int) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(logicalIndex int) *QSize, logicalIndex int) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(logicalIndex)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_SectionSizeFromContents, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_HorizontalOffset() int {

	return (int)(QHeaderView_virtualbase_HorizontalOffset(unsafe.Pointer(this.h)))

}
func (this *QHeaderView) OnHorizontalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_HorizontalOffset
func miqt_exec_callback_QHeaderView_HorizontalOffset(self QHeaderView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_HorizontalOffset)

	return (int)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_VerticalOffset() int {

	return (int)(QHeaderView_virtualbase_VerticalOffset(unsafe.Pointer(this.h)))

}
func (this *QHeaderView) OnVerticalOffset(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_VerticalOffset
func miqt_exec_callback_QHeaderView_VerticalOffset(self QHeaderView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_VerticalOffset)

	return (int)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_UpdateGeometries() {

	QHeaderView_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QHeaderView) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_UpdateGeometries
func miqt_exec_callback_QHeaderView_UpdateGeometries(self QHeaderView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHeaderView{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QHeaderView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QHeaderView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QHeaderView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_ScrollContentsBy
func miqt_exec_callback_QHeaderView_ScrollContentsBy(self QHeaderView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QHeaderView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QHeaderView_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QHeaderView) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_DataChanged
func miqt_exec_callback_QHeaderView_DataChanged(self QHeaderView, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(topLeft)

	slotval2 := newQModelIndex(bottomRight)

	var roles_ma struct_miqt_array = roles
	roles_ret := make([]int, int(roles_ma.len))
	roles_outCast := (*[0xffff]int)(unsafe.Pointer(roles_ma.data)) // hey ya
	for i := 0; i < int(roles_ma.len); i++ {
		roles_ret[i] = (int)(roles_outCast[i])
	}
	slotval3 := roles_ret

	gofunc((&QHeaderView{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QHeaderView) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QHeaderView_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QHeaderView) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_RowsInserted
func miqt_exec_callback_QHeaderView_RowsInserted(self QHeaderView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QHeaderView{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QHeaderView) callVirtualBase_VisualRect(index *QModelIndex) *QRect {

	_goptr := newQRect(QHeaderView_virtualbase_VisualRect(unsafe.Pointer(this.h), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnVisualRect(slot func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_VisualRect
func miqt_exec_callback_QHeaderView_VisualRect(self QHeaderView, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QRect, index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_VisualRect, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_ScrollTo(index *QModelIndex, hint ScrollHint) {

	QHeaderView_virtualbase_ScrollTo(unsafe.Pointer(this.h), index.cPointer(), hint)

}
func (this *QHeaderView) OnScrollTo(slot func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_ScrollTo
func miqt_exec_callback_QHeaderView_ScrollTo(self QHeaderView, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, hint ScrollHint), index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc((&QHeaderView{h: self}).callVirtualBase_ScrollTo, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_IndexAt(p *QPoint) *QModelIndex {

	_goptr := newQModelIndex(QHeaderView_virtualbase_IndexAt(unsafe.Pointer(this.h), p.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnIndexAt(slot func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_IndexAt
func miqt_exec_callback_QHeaderView_IndexAt(self QHeaderView, cb intptr_t, p *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(p *QPoint) *QModelIndex, p *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(p)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_IndexAt, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_IsIndexHidden(index *QModelIndex) bool {

	return (bool)(QHeaderView_virtualbase_IsIndexHidden(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QHeaderView) OnIsIndexHidden(slot func(super func(index *QModelIndex) bool, index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_IsIndexHidden
func miqt_exec_callback_QHeaderView_IsIndexHidden(self QHeaderView, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) bool, index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_IsIndexHidden, slotval1)

	return (bool)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_MoveCursor(param1 CursorAction, param2 KeyboardModifier) *QModelIndex {

	_goptr := newQModelIndex(QHeaderView_virtualbase_MoveCursor(unsafe.Pointer(this.h), param1, (int)(param2)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnMoveCursor(slot func(super func(param1 CursorAction, param2 KeyboardModifier) *QModelIndex, param1 CursorAction, param2 KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_MoveCursor
func miqt_exec_callback_QHeaderView_MoveCursor(self QHeaderView, cb intptr_t, param1 CursorAction, param2 int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 CursorAction, param2 KeyboardModifier) *QModelIndex, param1 CursorAction, param2 KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(param2)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_MoveCursor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_SetSelection(rect *QRect, flags SelectionFlag) {

	QHeaderView_virtualbase_SetSelection(unsafe.Pointer(this.h), rect.cPointer(), (int)(flags))

}
func (this *QHeaderView) OnSetSelection(slot func(super func(rect *QRect, flags SelectionFlag), rect *QRect, flags SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SetSelection
func miqt_exec_callback_QHeaderView_SetSelection(self QHeaderView, cb intptr_t, rect *QRect, flags int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRect, flags SelectionFlag), rect *QRect, flags SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(flags)

	gofunc((&QHeaderView{h: self}).callVirtualBase_SetSelection, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_VisualRegionForSelection(selection *QItemSelection) *QRegion {

	_goptr := newQRegion(QHeaderView_virtualbase_VisualRegionForSelection(unsafe.Pointer(this.h), selection.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnVisualRegionForSelection(slot func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_VisualRegionForSelection
func miqt_exec_callback_QHeaderView_VisualRegionForSelection(self QHeaderView, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selection *QItemSelection) *QRegion, selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_VisualRegionForSelection, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_InitStyleOptionForIndex(option *QStyleOptionHeader, logicalIndex int) {

	QHeaderView_virtualbase_InitStyleOptionForIndex(unsafe.Pointer(this.h), option.cPointer(), (int)(logicalIndex))

}
func (this *QHeaderView) OnInitStyleOptionForIndex(slot func(super func(option *QStyleOptionHeader, logicalIndex int), option *QStyleOptionHeader, logicalIndex int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_InitStyleOptionForIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_InitStyleOptionForIndex
func miqt_exec_callback_QHeaderView_InitStyleOptionForIndex(self QHeaderView, cb intptr_t, option *QStyleOptionHeader, logicalIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionHeader, logicalIndex int), option *QStyleOptionHeader, logicalIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionHeader(option)

	slotval2 := (int)(logicalIndex)

	gofunc((&QHeaderView{h: self}).callVirtualBase_InitStyleOptionForIndex, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_InitStyleOption(option *QStyleOptionHeader) {

	QHeaderView_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QHeaderView) OnInitStyleOption(slot func(super func(option *QStyleOptionHeader), option *QStyleOptionHeader)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_InitStyleOption
func miqt_exec_callback_QHeaderView_InitStyleOption(self QHeaderView, cb intptr_t, option *QStyleOptionHeader) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionHeader), option *QStyleOptionHeader))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionHeader(option)

	gofunc((&QHeaderView{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QHeaderView) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QHeaderView_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QHeaderView) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SetSelectionModel
func miqt_exec_callback_QHeaderView_SetSelectionModel(self QHeaderView, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QHeaderView{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QHeaderView) callVirtualBase_KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))

	QHeaderView_virtualbase_KeyboardSearch(unsafe.Pointer(this.h), search_ms)

}
func (this *QHeaderView) OnKeyboardSearch(slot func(super func(search string), search string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_KeyboardSearch(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_KeyboardSearch
func miqt_exec_callback_QHeaderView_KeyboardSearch(self QHeaderView, cb intptr_t, search struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(search string), search string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var search_ms struct_miqt_string = search
	search_ret := GoStringN(search_ms.data, int(int64(search_ms.len)))
	free(unsafe.Pointer(search_ms.data))
	slotval1 := search_ret

	gofunc((&QHeaderView{h: self}).callVirtualBase_KeyboardSearch, slotval1)

}

func (this *QHeaderView) callVirtualBase_SizeHintForRow(row int) int {

	return (int)(QHeaderView_virtualbase_SizeHintForRow(unsafe.Pointer(this.h), (int)(row)))

}
func (this *QHeaderView) OnSizeHintForRow(slot func(super func(row int) int, row int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SizeHintForRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SizeHintForRow
func miqt_exec_callback_QHeaderView_SizeHintForRow(self QHeaderView, cb intptr_t, row int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int) int, row int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_SizeHintForRow, slotval1)

	return (int)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_SizeHintForColumn(column int) int {

	return (int)(QHeaderView_virtualbase_SizeHintForColumn(unsafe.Pointer(this.h), (int)(column)))

}
func (this *QHeaderView) OnSizeHintForColumn(slot func(super func(column int) int, column int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SizeHintForColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SizeHintForColumn
func miqt_exec_callback_QHeaderView_SizeHintForColumn(self QHeaderView, cb intptr_t, column int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int) int, column int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_SizeHintForColumn, slotval1)

	return (int)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {

	return newQAbstractItemDelegate(QHeaderView_virtualbase_ItemDelegateForIndex(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QHeaderView) OnItemDelegateForIndex(slot func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_ItemDelegateForIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_ItemDelegateForIndex
func miqt_exec_callback_QHeaderView_ItemDelegateForIndex(self QHeaderView, cb intptr_t, index *QModelIndex) *QAbstractItemDelegate {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_ItemDelegateForIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QHeaderView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_InputMethodQuery
func miqt_exec_callback_QHeaderView_InputMethodQuery(self QHeaderView, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QHeaderView) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QHeaderView_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QHeaderView) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SetRootIndex
func miqt_exec_callback_QHeaderView_SetRootIndex(self QHeaderView, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QHeaderView{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QHeaderView) callVirtualBase_SelectAll() {

	QHeaderView_virtualbase_SelectAll(unsafe.Pointer(this.h))

}
func (this *QHeaderView) OnSelectAll(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SelectAll(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SelectAll
func miqt_exec_callback_QHeaderView_SelectAll(self QHeaderView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHeaderView{h: self}).callVirtualBase_SelectAll)

}

func (this *QHeaderView) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QHeaderView_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QHeaderView) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_RowsAboutToBeRemoved
func miqt_exec_callback_QHeaderView_RowsAboutToBeRemoved(self QHeaderView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QHeaderView{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QHeaderView) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QHeaderView_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QHeaderView) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SelectionChanged
func miqt_exec_callback_QHeaderView_SelectionChanged(self QHeaderView, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QHeaderView{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_UpdateEditorData() {

	QHeaderView_virtualbase_UpdateEditorData(unsafe.Pointer(this.h))

}
func (this *QHeaderView) OnUpdateEditorData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_UpdateEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_UpdateEditorData
func miqt_exec_callback_QHeaderView_UpdateEditorData(self QHeaderView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHeaderView{h: self}).callVirtualBase_UpdateEditorData)

}

func (this *QHeaderView) callVirtualBase_UpdateEditorGeometries() {

	QHeaderView_virtualbase_UpdateEditorGeometries(unsafe.Pointer(this.h))

}
func (this *QHeaderView) OnUpdateEditorGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_UpdateEditorGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_UpdateEditorGeometries
func miqt_exec_callback_QHeaderView_UpdateEditorGeometries(self QHeaderView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QHeaderView{h: self}).callVirtualBase_UpdateEditorGeometries)

}

func (this *QHeaderView) callVirtualBase_VerticalScrollbarAction(action int) {

	QHeaderView_virtualbase_VerticalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QHeaderView) OnVerticalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_VerticalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_VerticalScrollbarAction
func miqt_exec_callback_QHeaderView_VerticalScrollbarAction(self QHeaderView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QHeaderView{h: self}).callVirtualBase_VerticalScrollbarAction, slotval1)

}

func (this *QHeaderView) callVirtualBase_HorizontalScrollbarAction(action int) {

	QHeaderView_virtualbase_HorizontalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QHeaderView) OnHorizontalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_HorizontalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_HorizontalScrollbarAction
func miqt_exec_callback_QHeaderView_HorizontalScrollbarAction(self QHeaderView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QHeaderView{h: self}).callVirtualBase_HorizontalScrollbarAction, slotval1)

}

func (this *QHeaderView) callVirtualBase_VerticalScrollbarValueChanged(value int) {

	QHeaderView_virtualbase_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QHeaderView) OnVerticalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_VerticalScrollbarValueChanged
func miqt_exec_callback_QHeaderView_VerticalScrollbarValueChanged(self QHeaderView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QHeaderView{h: self}).callVirtualBase_VerticalScrollbarValueChanged, slotval1)

}

func (this *QHeaderView) callVirtualBase_HorizontalScrollbarValueChanged(value int) {

	QHeaderView_virtualbase_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QHeaderView) OnHorizontalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_HorizontalScrollbarValueChanged
func miqt_exec_callback_QHeaderView_HorizontalScrollbarValueChanged(self QHeaderView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QHeaderView{h: self}).callVirtualBase_HorizontalScrollbarValueChanged, slotval1)

}

func (this *QHeaderView) callVirtualBase_CloseEditor(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {

	QHeaderView_virtualbase_CloseEditor(unsafe.Pointer(this.h), editor.cPointer(), (int)(hint))

}
func (this *QHeaderView) OnCloseEditor(slot func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_CloseEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_CloseEditor
func miqt_exec_callback_QHeaderView_CloseEditor(self QHeaderView, cb intptr_t, editor *QWidget, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := (QAbstractItemDelegate__EndEditHint)(hint)

	gofunc((&QHeaderView{h: self}).callVirtualBase_CloseEditor, slotval1, slotval2)

}

func (this *QHeaderView) callVirtualBase_CommitData(editor *QWidget) {

	QHeaderView_virtualbase_CommitData(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QHeaderView) OnCommitData(slot func(super func(editor *QWidget), editor *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_CommitData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_CommitData
func miqt_exec_callback_QHeaderView_CommitData(self QHeaderView, cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget), editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc((&QHeaderView{h: self}).callVirtualBase_CommitData, slotval1)

}

func (this *QHeaderView) callVirtualBase_EditorDestroyed(editor *QObject) {

	QHeaderView_virtualbase_EditorDestroyed(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QHeaderView) OnEditorDestroyed(slot func(super func(editor *QObject), editor *QObject)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_EditorDestroyed(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_EditorDestroyed
func miqt_exec_callback_QHeaderView_EditorDestroyed(self QHeaderView, cb intptr_t, editor *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QObject), editor *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(editor)

	gofunc((&QHeaderView{h: self}).callVirtualBase_EditorDestroyed, slotval1)

}

func (this *QHeaderView) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QHeaderView_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QHeaderView) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SelectedIndexes
func miqt_exec_callback_QHeaderView_SelectedIndexes(self QHeaderView, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QHeaderView) callVirtualBase_Edit2(index *QModelIndex, trigger EditTrigger, event *QEvent) bool {

	return (bool)(QHeaderView_virtualbase_Edit2(unsafe.Pointer(this.h), index.cPointer(), trigger, event.cPointer()))

}
func (this *QHeaderView) OnEdit2(slot func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_Edit2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_Edit2
func miqt_exec_callback_QHeaderView_Edit2(self QHeaderView, cb intptr_t, index *QModelIndex, trigger EditTrigger, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx
	slotval3 := newQEvent(event)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_Edit2, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_SelectionCommand(index *QModelIndex, event *QEvent) SelectionFlag {

	return (SelectionFlag)(QHeaderView_virtualbase_SelectionCommand(unsafe.Pointer(this.h), index.cPointer(), event.cPointer()))

}
func (this *QHeaderView) OnSelectionCommand(slot func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_SelectionCommand(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_SelectionCommand
func miqt_exec_callback_QHeaderView_SelectionCommand(self QHeaderView, cb intptr_t, index *QModelIndex, event *QEvent) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_SelectionCommand, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_StartDrag(supportedActions DropAction) {

	QHeaderView_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QHeaderView) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_StartDrag
func miqt_exec_callback_QHeaderView_StartDrag(self QHeaderView, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QHeaderView{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QHeaderView) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QHeaderView_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QHeaderView) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_InitViewItemOption
func miqt_exec_callback_QHeaderView_InitViewItemOption(self QHeaderView, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QHeaderView{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QHeaderView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QHeaderView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QHeaderView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_FocusNextPrevChild
func miqt_exec_callback_QHeaderView_FocusNextPrevChild(self QHeaderView, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QHeaderView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_DragEnterEvent
func miqt_exec_callback_QHeaderView_DragEnterEvent(self QHeaderView, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QHeaderView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_DragMoveEvent
func miqt_exec_callback_QHeaderView_DragMoveEvent(self QHeaderView, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QHeaderView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_DragLeaveEvent
func miqt_exec_callback_QHeaderView_DragLeaveEvent(self QHeaderView, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_DropEvent(event *QDropEvent) {

	QHeaderView_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_DropEvent
func miqt_exec_callback_QHeaderView_DropEvent(self QHeaderView, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QHeaderView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_FocusInEvent
func miqt_exec_callback_QHeaderView_FocusInEvent(self QHeaderView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QHeaderView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_FocusOutEvent
func miqt_exec_callback_QHeaderView_FocusOutEvent(self QHeaderView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QHeaderView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_KeyPressEvent
func miqt_exec_callback_QHeaderView_KeyPressEvent(self QHeaderView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QHeaderView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_ResizeEvent
func miqt_exec_callback_QHeaderView_ResizeEvent(self QHeaderView, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QHeaderView_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_TimerEvent
func miqt_exec_callback_QHeaderView_TimerEvent(self QHeaderView, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QHeaderView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QHeaderView) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_InputMethodEvent
func miqt_exec_callback_QHeaderView_InputMethodEvent(self QHeaderView, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QHeaderView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QHeaderView) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QHeaderView_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QHeaderView) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_EventFilter
func miqt_exec_callback_QHeaderView_EventFilter(self QHeaderView, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QHeaderView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QHeaderView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QHeaderView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_ViewportSizeHint
func miqt_exec_callback_QHeaderView_ViewportSizeHint(self QHeaderView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}
