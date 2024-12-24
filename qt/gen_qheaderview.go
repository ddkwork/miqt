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

func (this *QHeaderView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QHeaderView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QHeaderView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_MetaObject
func miqt_exec_callback_QHeaderView_MetaObject(self QHeaderView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QHeaderView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QHeaderView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QHeaderView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QHeaderView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QHeaderView_Metacast
func miqt_exec_callback_QHeaderView_Metacast(self QHeaderView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QHeaderView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
