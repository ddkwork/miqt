package qt

import (
	"unsafe"
)

type QAbstractItemView__SelectionMode int

const (
	QAbstractItemView__NoSelection         QAbstractItemView__SelectionMode = 0
	QAbstractItemView__SingleSelection     QAbstractItemView__SelectionMode = 1
	QAbstractItemView__MultiSelection      QAbstractItemView__SelectionMode = 2
	QAbstractItemView__ExtendedSelection   QAbstractItemView__SelectionMode = 3
	QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = 4
)

type QAbstractItemView__SelectionBehavior int

const (
	QAbstractItemView__SelectItems   QAbstractItemView__SelectionBehavior = 0
	QAbstractItemView__SelectRows    QAbstractItemView__SelectionBehavior = 1
	QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = 2
)

type QAbstractItemView__ScrollHint int

const (
	QAbstractItemView__EnsureVisible    QAbstractItemView__ScrollHint = 0
	QAbstractItemView__PositionAtTop    QAbstractItemView__ScrollHint = 1
	QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = 2
	QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = 3
)

type QAbstractItemView__EditTrigger int

const (
	QAbstractItemView__NoEditTriggers  QAbstractItemView__EditTrigger = 0
	QAbstractItemView__CurrentChanged  QAbstractItemView__EditTrigger = 1
	QAbstractItemView__DoubleClicked   QAbstractItemView__EditTrigger = 2
	QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = 4
	QAbstractItemView__EditKeyPressed  QAbstractItemView__EditTrigger = 8
	QAbstractItemView__AnyKeyPressed   QAbstractItemView__EditTrigger = 16
	QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = 31
)

type QAbstractItemView__ScrollMode int

const (
	QAbstractItemView__ScrollPerItem  QAbstractItemView__ScrollMode = 0
	QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = 1
)

type QAbstractItemView__DragDropMode int

const (
	QAbstractItemView__NoDragDrop   QAbstractItemView__DragDropMode = 0
	QAbstractItemView__DragOnly     QAbstractItemView__DragDropMode = 1
	QAbstractItemView__DropOnly     QAbstractItemView__DragDropMode = 2
	QAbstractItemView__DragDrop     QAbstractItemView__DragDropMode = 3
	QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = 4
)

type QAbstractItemView__CursorAction int

const (
	QAbstractItemView__MoveUp       QAbstractItemView__CursorAction = 0
	QAbstractItemView__MoveDown     QAbstractItemView__CursorAction = 1
	QAbstractItemView__MoveLeft     QAbstractItemView__CursorAction = 2
	QAbstractItemView__MoveRight    QAbstractItemView__CursorAction = 3
	QAbstractItemView__MoveHome     QAbstractItemView__CursorAction = 4
	QAbstractItemView__MoveEnd      QAbstractItemView__CursorAction = 5
	QAbstractItemView__MovePageUp   QAbstractItemView__CursorAction = 6
	QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = 7
	QAbstractItemView__MoveNext     QAbstractItemView__CursorAction = 8
	QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = 9
)

type QAbstractItemView__State int

const (
	QAbstractItemView__NoState            QAbstractItemView__State = 0
	QAbstractItemView__DraggingState      QAbstractItemView__State = 1
	QAbstractItemView__DragSelectingState QAbstractItemView__State = 2
	QAbstractItemView__EditingState       QAbstractItemView__State = 3
	QAbstractItemView__ExpandingState     QAbstractItemView__State = 4
	QAbstractItemView__CollapsingState    QAbstractItemView__State = 5
	QAbstractItemView__AnimatingState     QAbstractItemView__State = 6
)

type QAbstractItemView__DropIndicatorPosition int

const (
	QAbstractItemView__OnItem     QAbstractItemView__DropIndicatorPosition = 0
	QAbstractItemView__AboveItem  QAbstractItemView__DropIndicatorPosition = 1
	QAbstractItemView__BelowItem  QAbstractItemView__DropIndicatorPosition = 2
	QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = 3
)

type QAbstractItemView struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractItemView constructs a new QAbstractItemView object.
func NewQAbstractItemView(parent *QWidget) *QAbstractItemView {
	ret := newQAbstractItemView(QAbstractItemView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractItemView2 constructs a new QAbstractItemView object.
func NewQAbstractItemView2() *QAbstractItemView {
	ret := newQAbstractItemView(QAbstractItemView_new2())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractItemView) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractItemView_MetaObject(this.h))
}

func (this *QAbstractItemView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractItemView_Metacast(this.h, param1_Cstring))
}

func QAbstractItemView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractItemView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemView) SetModel(model *QAbstractItemModel) {
	QAbstractItemView_SetModel(this.h, model.cPointer())
}

func (this *QAbstractItemView) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QAbstractItemView_Model(this.h))
}

func (this *QAbstractItemView) SetSelectionModel(selectionModel *QItemSelectionModel) {
	QAbstractItemView_SetSelectionModel(this.h, selectionModel.cPointer())
}

func (this *QAbstractItemView) SelectionModel() *QItemSelectionModel {
	return newQItemSelectionModel(QAbstractItemView_SelectionModel(this.h))
}

func (this *QAbstractItemView) SetItemDelegate(delegate *QAbstractItemDelegate) {
	QAbstractItemView_SetItemDelegate(this.h, delegate.cPointer())
}

func (this *QAbstractItemView) ItemDelegate() *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QAbstractItemView_ItemDelegate(this.h))
}

func (this *QAbstractItemView) SetSelectionMode(mode QAbstractItemView__SelectionMode) {
	QAbstractItemView_SetSelectionMode(this.h, (int)(mode))
}

func (this *QAbstractItemView) SelectionMode() QAbstractItemView__SelectionMode {
	return (QAbstractItemView__SelectionMode)(QAbstractItemView_SelectionMode(this.h))
}

func (this *QAbstractItemView) SetSelectionBehavior(behavior QAbstractItemView__SelectionBehavior) {
	QAbstractItemView_SetSelectionBehavior(this.h, (int)(behavior))
}

func (this *QAbstractItemView) SelectionBehavior() QAbstractItemView__SelectionBehavior {
	return (QAbstractItemView__SelectionBehavior)(QAbstractItemView_SelectionBehavior(this.h))
}

func (this *QAbstractItemView) CurrentIndex() *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemView_CurrentIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) RootIndex() *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemView_RootIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SetEditTriggers(triggers EditTriggers) {
	QAbstractItemView_SetEditTriggers(this.h, triggers)
}

func (this *QAbstractItemView) EditTriggers() EditTriggers {
	xxxxxxxxx
}

func (this *QAbstractItemView) SetVerticalScrollMode(mode ScrollMode) {
	QAbstractItemView_SetVerticalScrollMode(this.h, mode)
}

func (this *QAbstractItemView) VerticalScrollMode() ScrollMode {
	xxxxxxxxx
}

func (this *QAbstractItemView) ResetVerticalScrollMode() {
	QAbstractItemView_ResetVerticalScrollMode(this.h)
}

func (this *QAbstractItemView) SetHorizontalScrollMode(mode ScrollMode) {
	QAbstractItemView_SetHorizontalScrollMode(this.h, mode)
}

func (this *QAbstractItemView) HorizontalScrollMode() ScrollMode {
	xxxxxxxxx
}

func (this *QAbstractItemView) ResetHorizontalScrollMode() {
	QAbstractItemView_ResetHorizontalScrollMode(this.h)
}

func (this *QAbstractItemView) SetAutoScroll(enable bool) {
	QAbstractItemView_SetAutoScroll(this.h, (bool)(enable))
}

func (this *QAbstractItemView) HasAutoScroll() bool {
	return (bool)(QAbstractItemView_HasAutoScroll(this.h))
}

func (this *QAbstractItemView) SetAutoScrollMargin(margin int) {
	QAbstractItemView_SetAutoScrollMargin(this.h, (int)(margin))
}

func (this *QAbstractItemView) AutoScrollMargin() int {
	return (int)(QAbstractItemView_AutoScrollMargin(this.h))
}

func (this *QAbstractItemView) SetTabKeyNavigation(enable bool) {
	QAbstractItemView_SetTabKeyNavigation(this.h, (bool)(enable))
}

func (this *QAbstractItemView) TabKeyNavigation() bool {
	return (bool)(QAbstractItemView_TabKeyNavigation(this.h))
}

func (this *QAbstractItemView) SetDropIndicatorShown(enable bool) {
	QAbstractItemView_SetDropIndicatorShown(this.h, (bool)(enable))
}

func (this *QAbstractItemView) ShowDropIndicator() bool {
	return (bool)(QAbstractItemView_ShowDropIndicator(this.h))
}

func (this *QAbstractItemView) SetDragEnabled(enable bool) {
	QAbstractItemView_SetDragEnabled(this.h, (bool)(enable))
}

func (this *QAbstractItemView) DragEnabled() bool {
	return (bool)(QAbstractItemView_DragEnabled(this.h))
}

func (this *QAbstractItemView) SetDragDropOverwriteMode(overwrite bool) {
	QAbstractItemView_SetDragDropOverwriteMode(this.h, (bool)(overwrite))
}

func (this *QAbstractItemView) DragDropOverwriteMode() bool {
	return (bool)(QAbstractItemView_DragDropOverwriteMode(this.h))
}

func (this *QAbstractItemView) SetDragDropMode(behavior DragDropMode) {
	QAbstractItemView_SetDragDropMode(this.h, behavior)
}

func (this *QAbstractItemView) DragDropMode() DragDropMode {
	xxxxxxxxx
}

func (this *QAbstractItemView) SetDefaultDropAction(dropAction DropAction) {
	QAbstractItemView_SetDefaultDropAction(this.h, (int)(dropAction))
}

func (this *QAbstractItemView) DefaultDropAction() DropAction {
	return (DropAction)(QAbstractItemView_DefaultDropAction(this.h))
}

func (this *QAbstractItemView) SetAlternatingRowColors(enable bool) {
	QAbstractItemView_SetAlternatingRowColors(this.h, (bool)(enable))
}

func (this *QAbstractItemView) AlternatingRowColors() bool {
	return (bool)(QAbstractItemView_AlternatingRowColors(this.h))
}

func (this *QAbstractItemView) SetIconSize(size *QSize) {
	QAbstractItemView_SetIconSize(this.h, size.cPointer())
}

func (this *QAbstractItemView) IconSize() *QSize {
	_goptr := newQSize(QAbstractItemView_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SetTextElideMode(mode TextElideMode) {
	QAbstractItemView_SetTextElideMode(this.h, (int)(mode))
}

func (this *QAbstractItemView) TextElideMode() TextElideMode {
	return (TextElideMode)(QAbstractItemView_TextElideMode(this.h))
}

func (this *QAbstractItemView) KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))
	QAbstractItemView_KeyboardSearch(this.h, search_ms)
}

func (this *QAbstractItemView) VisualRect(index *QModelIndex) *QRect {
	_goptr := newQRect(QAbstractItemView_VisualRect(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) ScrollTo(index *QModelIndex, hint ScrollHint) {
	QAbstractItemView_ScrollTo(this.h, index.cPointer(), hint)
}

func (this *QAbstractItemView) IndexAt(point *QPoint) *QModelIndex {
	_goptr := newQModelIndex(QAbstractItemView_IndexAt(this.h, point.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SizeHintForIndex(index *QModelIndex) *QSize {
	_goptr := newQSize(QAbstractItemView_SizeHintForIndex(this.h, index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) SizeHintForRow(row int) int {
	return (int)(QAbstractItemView_SizeHintForRow(this.h, (int)(row)))
}

func (this *QAbstractItemView) SizeHintForColumn(column int) int {
	return (int)(QAbstractItemView_SizeHintForColumn(this.h, (int)(column)))
}

func (this *QAbstractItemView) UpdateThreshold() uint32_t {
	return (uint32_t)(QAbstractItemView_UpdateThreshold(this.h))
}

func (this *QAbstractItemView) SetUpdateThreshold(threshold uint32_t) {
	QAbstractItemView_SetUpdateThreshold(this.h, (uint32_t)(threshold))
}

func (this *QAbstractItemView) OpenPersistentEditor(index *QModelIndex) {
	QAbstractItemView_OpenPersistentEditor(this.h, index.cPointer())
}

func (this *QAbstractItemView) ClosePersistentEditor(index *QModelIndex) {
	QAbstractItemView_ClosePersistentEditor(this.h, index.cPointer())
}

func (this *QAbstractItemView) IsPersistentEditorOpen(index *QModelIndex) bool {
	return (bool)(QAbstractItemView_IsPersistentEditorOpen(this.h, index.cPointer()))
}

func (this *QAbstractItemView) SetIndexWidget(index *QModelIndex, widget *QWidget) {
	QAbstractItemView_SetIndexWidget(this.h, index.cPointer(), widget.cPointer())
}

func (this *QAbstractItemView) IndexWidget(index *QModelIndex) *QWidget {
	return newQWidget(QAbstractItemView_IndexWidget(this.h, index.cPointer()))
}

func (this *QAbstractItemView) SetItemDelegateForRow(row int, delegate *QAbstractItemDelegate) {
	QAbstractItemView_SetItemDelegateForRow(this.h, (int)(row), delegate.cPointer())
}

func (this *QAbstractItemView) ItemDelegateForRow(row int) *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QAbstractItemView_ItemDelegateForRow(this.h, (int)(row)))
}

func (this *QAbstractItemView) SetItemDelegateForColumn(column int, delegate *QAbstractItemDelegate) {
	QAbstractItemView_SetItemDelegateForColumn(this.h, (int)(column), delegate.cPointer())
}

func (this *QAbstractItemView) ItemDelegateForColumn(column int) *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QAbstractItemView_ItemDelegateForColumn(this.h, (int)(column)))
}

func (this *QAbstractItemView) ItemDelegateWithIndex(index *QModelIndex) *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QAbstractItemView_ItemDelegateWithIndex(this.h, index.cPointer()))
}

func (this *QAbstractItemView) ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QAbstractItemView_ItemDelegateForIndex(this.h, index.cPointer()))
}

func (this *QAbstractItemView) InputMethodQuery(query InputMethodQuery) *QVariant {
	_goptr := newQVariant(QAbstractItemView_InputMethodQuery(this.h, (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractItemView) Reset() {
	QAbstractItemView_Reset(this.h)
}

func (this *QAbstractItemView) SetRootIndex(index *QModelIndex) {
	QAbstractItemView_SetRootIndex(this.h, index.cPointer())
}

func (this *QAbstractItemView) DoItemsLayout() {
	QAbstractItemView_DoItemsLayout(this.h)
}

func (this *QAbstractItemView) SelectAll() {
	QAbstractItemView_SelectAll(this.h)
}

func (this *QAbstractItemView) Edit(index *QModelIndex) {
	QAbstractItemView_Edit(this.h, index.cPointer())
}

func (this *QAbstractItemView) ClearSelection() {
	QAbstractItemView_ClearSelection(this.h)
}

func (this *QAbstractItemView) SetCurrentIndex(index *QModelIndex) {
	QAbstractItemView_SetCurrentIndex(this.h, index.cPointer())
}

func (this *QAbstractItemView) ScrollToTop() {
	QAbstractItemView_ScrollToTop(this.h)
}

func (this *QAbstractItemView) ScrollToBottom() {
	QAbstractItemView_ScrollToBottom(this.h)
}

func (this *QAbstractItemView) Update(index *QModelIndex) {
	QAbstractItemView_Update(this.h, index.cPointer())
}

func (this *QAbstractItemView) Pressed(index *QModelIndex) {
	QAbstractItemView_Pressed(this.h, index.cPointer())
}

func (this *QAbstractItemView) OnPressed(slot func(index *QModelIndex)) {
	QAbstractItemView_connect_Pressed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Pressed
func miqt_exec_callback_QAbstractItemView_Pressed(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QAbstractItemView) Clicked(index *QModelIndex) {
	QAbstractItemView_Clicked(this.h, index.cPointer())
}

func (this *QAbstractItemView) OnClicked(slot func(index *QModelIndex)) {
	QAbstractItemView_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Clicked
func miqt_exec_callback_QAbstractItemView_Clicked(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QAbstractItemView) DoubleClicked(index *QModelIndex) {
	QAbstractItemView_DoubleClicked(this.h, index.cPointer())
}

func (this *QAbstractItemView) OnDoubleClicked(slot func(index *QModelIndex)) {
	QAbstractItemView_connect_DoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DoubleClicked
func miqt_exec_callback_QAbstractItemView_DoubleClicked(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QAbstractItemView) Activated(index *QModelIndex) {
	QAbstractItemView_Activated(this.h, index.cPointer())
}

func (this *QAbstractItemView) OnActivated(slot func(index *QModelIndex)) {
	QAbstractItemView_connect_Activated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Activated
func miqt_exec_callback_QAbstractItemView_Activated(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QAbstractItemView) Entered(index *QModelIndex) {
	QAbstractItemView_Entered(this.h, index.cPointer())
}

func (this *QAbstractItemView) OnEntered(slot func(index *QModelIndex)) {
	QAbstractItemView_connect_Entered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Entered
func miqt_exec_callback_QAbstractItemView_Entered(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QAbstractItemView) ViewportEntered() {
	QAbstractItemView_ViewportEntered(this.h)
}

func (this *QAbstractItemView) OnViewportEntered(slot func()) {
	QAbstractItemView_connect_ViewportEntered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ViewportEntered
func miqt_exec_callback_QAbstractItemView_ViewportEntered(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractItemView) IconSizeChanged(size *QSize) {
	QAbstractItemView_IconSizeChanged(this.h, size.cPointer())
}

func (this *QAbstractItemView) OnIconSizeChanged(slot func(size *QSize)) {
	QAbstractItemView_connect_IconSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_IconSizeChanged
func miqt_exec_callback_QAbstractItemView_IconSizeChanged(cb intptr_t, size *QSize) {
	gofunc, ok := cgo.Handle(cb).Value().(func(size *QSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(size)

	gofunc(slotval1)
}

func QAbstractItemView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractItemView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractItemView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractItemView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractItemView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractItemView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractItemView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MetaObject
func miqt_exec_callback_QAbstractItemView_MetaObject(self QAbstractItemView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractItemView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractItemView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractItemView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Metacast
func miqt_exec_callback_QAbstractItemView_Metacast(self QAbstractItemView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
