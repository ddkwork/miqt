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

func (this *QAbstractItemView) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QAbstractItemView_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QAbstractItemView) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SetModel
func miqt_exec_callback_QAbstractItemView_SetModel(self QAbstractItemView, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_SetSelectionModel(selectionModel *QItemSelectionModel) {

	QAbstractItemView_virtualbase_SetSelectionModel(unsafe.Pointer(this.h), selectionModel.cPointer())

}
func (this *QAbstractItemView) OnSetSelectionModel(slot func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SetSelectionModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SetSelectionModel
func miqt_exec_callback_QAbstractItemView_SetSelectionModel(self QAbstractItemView, cb intptr_t, selectionModel *QItemSelectionModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selectionModel *QItemSelectionModel), selectionModel *QItemSelectionModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelectionModel(selectionModel)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_SetSelectionModel, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_KeyboardSearch(search string) {
	search_ms := struct_miqt_string{}
	search_ms.data = CString(search)
	search_ms.len = size_t(len(search))
	defer free(unsafe.Pointer(search_ms.data))

	QAbstractItemView_virtualbase_KeyboardSearch(unsafe.Pointer(this.h), search_ms)

}
func (this *QAbstractItemView) OnKeyboardSearch(slot func(super func(search string), search string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_KeyboardSearch(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_KeyboardSearch
func miqt_exec_callback_QAbstractItemView_KeyboardSearch(self QAbstractItemView, cb intptr_t, search struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(search string), search string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var search_ms struct_miqt_string = search
	search_ret := GoStringN(search_ms.data, int(int64(search_ms.len)))
	free(unsafe.Pointer(search_ms.data))
	slotval1 := search_ret

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_KeyboardSearch, slotval1)

}
func (this *QAbstractItemView) OnVisualRect(slot func(index *QModelIndex) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_VisualRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_VisualRect
func miqt_exec_callback_QAbstractItemView_VisualRect(self QAbstractItemView, cb intptr_t, index *QModelIndex) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}
func (this *QAbstractItemView) OnScrollTo(slot func(index *QModelIndex, hint ScrollHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ScrollTo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ScrollTo
func miqt_exec_callback_QAbstractItemView_ScrollTo(self QAbstractItemView, cb intptr_t, index *QModelIndex, hint ScrollHint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex, hint ScrollHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx

	gofunc(slotval1, slotval2)

}
func (this *QAbstractItemView) OnIndexAt(slot func(point *QPoint) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_IndexAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_IndexAt
func miqt_exec_callback_QAbstractItemView_IndexAt(self QAbstractItemView, cb intptr_t, point *QPoint) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(point *QPoint) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(point)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_SizeHintForRow(row int) int {

	return (int)(QAbstractItemView_virtualbase_SizeHintForRow(unsafe.Pointer(this.h), (int)(row)))

}
func (this *QAbstractItemView) OnSizeHintForRow(slot func(super func(row int) int, row int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SizeHintForRow(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SizeHintForRow
func miqt_exec_callback_QAbstractItemView_SizeHintForRow(self QAbstractItemView, cb intptr_t, row int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(row int) int, row int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(row)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_SizeHintForRow, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_SizeHintForColumn(column int) int {

	return (int)(QAbstractItemView_virtualbase_SizeHintForColumn(unsafe.Pointer(this.h), (int)(column)))

}
func (this *QAbstractItemView) OnSizeHintForColumn(slot func(super func(column int) int, column int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SizeHintForColumn(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SizeHintForColumn
func miqt_exec_callback_QAbstractItemView_SizeHintForColumn(self QAbstractItemView, cb intptr_t, column int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(column int) int, column int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(column)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_SizeHintForColumn, slotval1)

	return (int)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_ItemDelegateForIndex(index *QModelIndex) *QAbstractItemDelegate {

	return newQAbstractItemDelegate(QAbstractItemView_virtualbase_ItemDelegateForIndex(unsafe.Pointer(this.h), index.cPointer()))

}
func (this *QAbstractItemView) OnItemDelegateForIndex(slot func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ItemDelegateForIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ItemDelegateForIndex
func miqt_exec_callback_QAbstractItemView_ItemDelegateForIndex(self QAbstractItemView, cb intptr_t, index *QModelIndex) *QAbstractItemDelegate {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex) *QAbstractItemDelegate, index *QModelIndex) *QAbstractItemDelegate)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_ItemDelegateForIndex, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QAbstractItemView_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemView) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_InputMethodQuery
func miqt_exec_callback_QAbstractItemView_InputMethodQuery(self QAbstractItemView, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_Reset() {

	QAbstractItemView_virtualbase_Reset(unsafe.Pointer(this.h))

}
func (this *QAbstractItemView) OnReset(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_Reset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Reset
func miqt_exec_callback_QAbstractItemView_Reset(self QAbstractItemView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_Reset)

}

func (this *QAbstractItemView) callVirtualBase_SetRootIndex(index *QModelIndex) {

	QAbstractItemView_virtualbase_SetRootIndex(unsafe.Pointer(this.h), index.cPointer())

}
func (this *QAbstractItemView) OnSetRootIndex(slot func(super func(index *QModelIndex), index *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SetRootIndex(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SetRootIndex
func miqt_exec_callback_QAbstractItemView_SetRootIndex(self QAbstractItemView, cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex), index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_SetRootIndex, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_DoItemsLayout() {

	QAbstractItemView_virtualbase_DoItemsLayout(unsafe.Pointer(this.h))

}
func (this *QAbstractItemView) OnDoItemsLayout(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_DoItemsLayout(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DoItemsLayout
func miqt_exec_callback_QAbstractItemView_DoItemsLayout(self QAbstractItemView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_DoItemsLayout)

}

func (this *QAbstractItemView) callVirtualBase_SelectAll() {

	QAbstractItemView_virtualbase_SelectAll(unsafe.Pointer(this.h))

}
func (this *QAbstractItemView) OnSelectAll(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SelectAll(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SelectAll
func miqt_exec_callback_QAbstractItemView_SelectAll(self QAbstractItemView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_SelectAll)

}

func (this *QAbstractItemView) callVirtualBase_DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_CArray := (*[0xffff]int)(malloc(size_t(8 * len(roles))))
	defer free(unsafe.Pointer(roles_CArray))
	for i := range roles {
		roles_CArray[i] = (int)(roles[i])
	}
	roles_ma := struct_miqt_array{len: size_t(len(roles)), data: unsafe.Pointer(roles_CArray)}

	QAbstractItemView_virtualbase_DataChanged(unsafe.Pointer(this.h), topLeft.cPointer(), bottomRight.cPointer(), roles_ma)

}
func (this *QAbstractItemView) OnDataChanged(slot func(super func(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int), topLeft *QModelIndex, bottomRight *QModelIndex, roles []int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_DataChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DataChanged
func miqt_exec_callback_QAbstractItemView_DataChanged(self QAbstractItemView, cb intptr_t, topLeft *QModelIndex, bottomRight *QModelIndex, roles struct_miqt_array) {
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

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_DataChanged, slotval1, slotval2, slotval3)

}

func (this *QAbstractItemView) callVirtualBase_RowsInserted(parent *QModelIndex, start int, end int) {

	QAbstractItemView_virtualbase_RowsInserted(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QAbstractItemView) OnRowsInserted(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_RowsInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_RowsInserted
func miqt_exec_callback_QAbstractItemView_RowsInserted(self QAbstractItemView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_RowsInserted, slotval1, slotval2, slotval3)

}

func (this *QAbstractItemView) callVirtualBase_RowsAboutToBeRemoved(parent *QModelIndex, start int, end int) {

	QAbstractItemView_virtualbase_RowsAboutToBeRemoved(unsafe.Pointer(this.h), parent.cPointer(), (int)(start), (int)(end))

}
func (this *QAbstractItemView) OnRowsAboutToBeRemoved(slot func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_RowsAboutToBeRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_RowsAboutToBeRemoved
func miqt_exec_callback_QAbstractItemView_RowsAboutToBeRemoved(self QAbstractItemView, cb intptr_t, parent *QModelIndex, start int, end int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QModelIndex, start int, end int), parent *QModelIndex, start int, end int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(parent)

	slotval2 := (int)(start)

	slotval3 := (int)(end)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_RowsAboutToBeRemoved, slotval1, slotval2, slotval3)

}

func (this *QAbstractItemView) callVirtualBase_SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {

	QAbstractItemView_virtualbase_SelectionChanged(unsafe.Pointer(this.h), selected.cPointer(), deselected.cPointer())

}
func (this *QAbstractItemView) OnSelectionChanged(slot func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SelectionChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SelectionChanged
func miqt_exec_callback_QAbstractItemView_SelectionChanged(self QAbstractItemView, cb intptr_t, selected *QItemSelection, deselected *QItemSelection) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(selected *QItemSelection, deselected *QItemSelection), selected *QItemSelection, deselected *QItemSelection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selected)

	slotval2 := newQItemSelection(deselected)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_SelectionChanged, slotval1, slotval2)

}

func (this *QAbstractItemView) callVirtualBase_CurrentChanged(current *QModelIndex, previous *QModelIndex) {

	QAbstractItemView_virtualbase_CurrentChanged(unsafe.Pointer(this.h), current.cPointer(), previous.cPointer())

}
func (this *QAbstractItemView) OnCurrentChanged(slot func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_CurrentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_CurrentChanged
func miqt_exec_callback_QAbstractItemView_CurrentChanged(self QAbstractItemView, cb intptr_t, current *QModelIndex, previous *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(current *QModelIndex, previous *QModelIndex), current *QModelIndex, previous *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(current)

	slotval2 := newQModelIndex(previous)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_CurrentChanged, slotval1, slotval2)

}

func (this *QAbstractItemView) callVirtualBase_UpdateEditorData() {

	QAbstractItemView_virtualbase_UpdateEditorData(unsafe.Pointer(this.h))

}
func (this *QAbstractItemView) OnUpdateEditorData(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_UpdateEditorData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_UpdateEditorData
func miqt_exec_callback_QAbstractItemView_UpdateEditorData(self QAbstractItemView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_UpdateEditorData)

}

func (this *QAbstractItemView) callVirtualBase_UpdateEditorGeometries() {

	QAbstractItemView_virtualbase_UpdateEditorGeometries(unsafe.Pointer(this.h))

}
func (this *QAbstractItemView) OnUpdateEditorGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_UpdateEditorGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_UpdateEditorGeometries
func miqt_exec_callback_QAbstractItemView_UpdateEditorGeometries(self QAbstractItemView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_UpdateEditorGeometries)

}

func (this *QAbstractItemView) callVirtualBase_UpdateGeometries() {

	QAbstractItemView_virtualbase_UpdateGeometries(unsafe.Pointer(this.h))

}
func (this *QAbstractItemView) OnUpdateGeometries(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_UpdateGeometries(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_UpdateGeometries
func miqt_exec_callback_QAbstractItemView_UpdateGeometries(self QAbstractItemView, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_UpdateGeometries)

}

func (this *QAbstractItemView) callVirtualBase_VerticalScrollbarAction(action int) {

	QAbstractItemView_virtualbase_VerticalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QAbstractItemView) OnVerticalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_VerticalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_VerticalScrollbarAction
func miqt_exec_callback_QAbstractItemView_VerticalScrollbarAction(self QAbstractItemView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_VerticalScrollbarAction, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_HorizontalScrollbarAction(action int) {

	QAbstractItemView_virtualbase_HorizontalScrollbarAction(unsafe.Pointer(this.h), (int)(action))

}
func (this *QAbstractItemView) OnHorizontalScrollbarAction(slot func(super func(action int), action int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_HorizontalScrollbarAction(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_HorizontalScrollbarAction
func miqt_exec_callback_QAbstractItemView_HorizontalScrollbarAction(self QAbstractItemView, cb intptr_t, action int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(action int), action int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(action)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_HorizontalScrollbarAction, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_VerticalScrollbarValueChanged(value int) {

	QAbstractItemView_virtualbase_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QAbstractItemView) OnVerticalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_VerticalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_VerticalScrollbarValueChanged
func miqt_exec_callback_QAbstractItemView_VerticalScrollbarValueChanged(self QAbstractItemView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_VerticalScrollbarValueChanged, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_HorizontalScrollbarValueChanged(value int) {

	QAbstractItemView_virtualbase_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), (int)(value))

}
func (this *QAbstractItemView) OnHorizontalScrollbarValueChanged(slot func(super func(value int), value int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_HorizontalScrollbarValueChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_HorizontalScrollbarValueChanged
func miqt_exec_callback_QAbstractItemView_HorizontalScrollbarValueChanged(self QAbstractItemView, cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(value int), value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_HorizontalScrollbarValueChanged, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_CloseEditor(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {

	QAbstractItemView_virtualbase_CloseEditor(unsafe.Pointer(this.h), editor.cPointer(), (int)(hint))

}
func (this *QAbstractItemView) OnCloseEditor(slot func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_CloseEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_CloseEditor
func miqt_exec_callback_QAbstractItemView_CloseEditor(self QAbstractItemView, cb intptr_t, editor *QWidget, hint int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget, hint QAbstractItemDelegate__EndEditHint), editor *QWidget, hint QAbstractItemDelegate__EndEditHint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	slotval2 := (QAbstractItemDelegate__EndEditHint)(hint)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_CloseEditor, slotval1, slotval2)

}

func (this *QAbstractItemView) callVirtualBase_CommitData(editor *QWidget) {

	QAbstractItemView_virtualbase_CommitData(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QAbstractItemView) OnCommitData(slot func(super func(editor *QWidget), editor *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_CommitData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_CommitData
func miqt_exec_callback_QAbstractItemView_CommitData(self QAbstractItemView, cb intptr_t, editor *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QWidget), editor *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(editor)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_CommitData, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_EditorDestroyed(editor *QObject) {

	QAbstractItemView_virtualbase_EditorDestroyed(unsafe.Pointer(this.h), editor.cPointer())

}
func (this *QAbstractItemView) OnEditorDestroyed(slot func(super func(editor *QObject), editor *QObject)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_EditorDestroyed(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_EditorDestroyed
func miqt_exec_callback_QAbstractItemView_EditorDestroyed(self QAbstractItemView, cb intptr_t, editor *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(editor *QObject), editor *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(editor)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_EditorDestroyed, slotval1)

}
func (this *QAbstractItemView) OnMoveCursor(slot func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MoveCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MoveCursor
func miqt_exec_callback_QAbstractItemView_MoveCursor(self QAbstractItemView, cb intptr_t, cursorAction CursorAction, modifiers int) *QModelIndex {
	gofunc, ok := cgo.Handle(cb).Value().(func(cursorAction CursorAction, modifiers KeyboardModifier) *QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := (KeyboardModifier)(modifiers)

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}
func (this *QAbstractItemView) OnHorizontalOffset(slot func() int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_HorizontalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_HorizontalOffset
func miqt_exec_callback_QAbstractItemView_HorizontalOffset(self QAbstractItemView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}
func (this *QAbstractItemView) OnVerticalOffset(slot func() int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_VerticalOffset(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_VerticalOffset
func miqt_exec_callback_QAbstractItemView_VerticalOffset(self QAbstractItemView, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}
func (this *QAbstractItemView) OnIsIndexHidden(slot func(index *QModelIndex) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_IsIndexHidden(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_IsIndexHidden
func miqt_exec_callback_QAbstractItemView_IsIndexHidden(self QAbstractItemView, cb intptr_t, index *QModelIndex) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	virtualReturn := gofunc(slotval1)

	return (bool)(virtualReturn)

}
func (this *QAbstractItemView) OnSetSelection(slot func(rect *QRect, command SelectionFlag)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SetSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SetSelection
func miqt_exec_callback_QAbstractItemView_SetSelection(self QAbstractItemView, cb intptr_t, rect *QRect, command int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(rect *QRect, command SelectionFlag))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRect(rect)

	slotval2 := (SelectionFlag)(command)

	gofunc(slotval1, slotval2)

}
func (this *QAbstractItemView) OnVisualRegionForSelection(slot func(selection *QItemSelection) *QRegion) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_VisualRegionForSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_VisualRegionForSelection
func miqt_exec_callback_QAbstractItemView_VisualRegionForSelection(self QAbstractItemView, cb intptr_t, selection *QItemSelection) *QRegion {
	gofunc, ok := cgo.Handle(cb).Value().(func(selection *QItemSelection) *QRegion)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQItemSelection(selection)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_SelectedIndexes() []QModelIndex {

	var _ma struct_miqt_array = QAbstractItemView_virtualbase_SelectedIndexes(unsafe.Pointer(this.h))
	_ret := make([]QModelIndex, int(_ma.len))
	_outCast := (*[0xffff]*QModelIndex)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQModelIndex(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret

}
func (this *QAbstractItemView) OnSelectedIndexes(slot func(super func() []QModelIndex) []QModelIndex) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SelectedIndexes(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SelectedIndexes
func miqt_exec_callback_QAbstractItemView_SelectedIndexes(self QAbstractItemView, cb intptr_t) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() []QModelIndex) []QModelIndex)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_SelectedIndexes)
	virtualReturn_CArray := (*[0xffff]*QModelIndex)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i].cPointer()
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}

func (this *QAbstractItemView) callVirtualBase_Edit2(index *QModelIndex, trigger EditTrigger, event *QEvent) bool {

	return (bool)(QAbstractItemView_virtualbase_Edit2(unsafe.Pointer(this.h), index.cPointer(), trigger, event.cPointer()))

}
func (this *QAbstractItemView) OnEdit2(slot func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_Edit2(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Edit2
func miqt_exec_callback_QAbstractItemView_Edit2(self QAbstractItemView, cb intptr_t, index *QModelIndex, trigger EditTrigger, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, trigger EditTrigger, event *QEvent) bool, index *QModelIndex, trigger EditTrigger, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	xxxxxxxxx
	slotval3 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_Edit2, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_SelectionCommand(index *QModelIndex, event *QEvent) SelectionFlag {

	return (SelectionFlag)(QAbstractItemView_virtualbase_SelectionCommand(unsafe.Pointer(this.h), index.cPointer(), event.cPointer()))

}
func (this *QAbstractItemView) OnSelectionCommand(slot func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SelectionCommand(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SelectionCommand
func miqt_exec_callback_QAbstractItemView_SelectionCommand(self QAbstractItemView, cb intptr_t, index *QModelIndex, event *QEvent) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index *QModelIndex, event *QEvent) SelectionFlag, index *QModelIndex, event *QEvent) SelectionFlag)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_SelectionCommand, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_StartDrag(supportedActions DropAction) {

	QAbstractItemView_virtualbase_StartDrag(unsafe.Pointer(this.h), (int)(supportedActions))

}
func (this *QAbstractItemView) OnStartDrag(slot func(super func(supportedActions DropAction), supportedActions DropAction)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_StartDrag(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_StartDrag
func miqt_exec_callback_QAbstractItemView_StartDrag(self QAbstractItemView, cb intptr_t, supportedActions int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(supportedActions DropAction), supportedActions DropAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (DropAction)(supportedActions)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_StartDrag, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_InitViewItemOption(option *QStyleOptionViewItem) {

	QAbstractItemView_virtualbase_InitViewItemOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QAbstractItemView) OnInitViewItemOption(slot func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_InitViewItemOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_InitViewItemOption
func miqt_exec_callback_QAbstractItemView_InitViewItemOption(self QAbstractItemView, cb intptr_t, option *QStyleOptionViewItem) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionViewItem), option *QStyleOptionViewItem))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionViewItem(option)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_InitViewItemOption, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QAbstractItemView_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QAbstractItemView) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_FocusNextPrevChild
func miqt_exec_callback_QAbstractItemView_FocusNextPrevChild(self QAbstractItemView, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QAbstractItemView_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAbstractItemView) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_Event
func miqt_exec_callback_QAbstractItemView_Event(self QAbstractItemView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_ViewportEvent(event *QEvent) bool {

	return (bool)(QAbstractItemView_virtualbase_ViewportEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAbstractItemView) OnViewportEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ViewportEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ViewportEvent
func miqt_exec_callback_QAbstractItemView_ViewportEvent(self QAbstractItemView, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_ViewportEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QAbstractItemView_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MousePressEvent
func miqt_exec_callback_QAbstractItemView_MousePressEvent(self QAbstractItemView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QAbstractItemView_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MouseMoveEvent
func miqt_exec_callback_QAbstractItemView_MouseMoveEvent(self QAbstractItemView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QAbstractItemView_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MouseReleaseEvent
func miqt_exec_callback_QAbstractItemView_MouseReleaseEvent(self QAbstractItemView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QAbstractItemView_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MouseDoubleClickEvent
func miqt_exec_callback_QAbstractItemView_MouseDoubleClickEvent(self QAbstractItemView, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QAbstractItemView_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DragEnterEvent
func miqt_exec_callback_QAbstractItemView_DragEnterEvent(self QAbstractItemView, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QAbstractItemView_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DragMoveEvent
func miqt_exec_callback_QAbstractItemView_DragMoveEvent(self QAbstractItemView, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QAbstractItemView_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DragLeaveEvent
func miqt_exec_callback_QAbstractItemView_DragLeaveEvent(self QAbstractItemView, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_DropEvent(event *QDropEvent) {

	QAbstractItemView_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_DropEvent
func miqt_exec_callback_QAbstractItemView_DropEvent(self QAbstractItemView, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QAbstractItemView_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_FocusInEvent
func miqt_exec_callback_QAbstractItemView_FocusInEvent(self QAbstractItemView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QAbstractItemView_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_FocusOutEvent
func miqt_exec_callback_QAbstractItemView_FocusOutEvent(self QAbstractItemView, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QAbstractItemView_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_KeyPressEvent
func miqt_exec_callback_QAbstractItemView_KeyPressEvent(self QAbstractItemView, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QAbstractItemView_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ResizeEvent
func miqt_exec_callback_QAbstractItemView_ResizeEvent(self QAbstractItemView, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QAbstractItemView_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_TimerEvent
func miqt_exec_callback_QAbstractItemView_TimerEvent(self QAbstractItemView, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QAbstractItemView_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractItemView) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_InputMethodEvent
func miqt_exec_callback_QAbstractItemView_InputMethodEvent(self QAbstractItemView, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QAbstractItemView_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QAbstractItemView) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_EventFilter
func miqt_exec_callback_QAbstractItemView_EventFilter(self QAbstractItemView, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractItemView) callVirtualBase_ViewportSizeHint() *QSize {

	_goptr := newQSize(QAbstractItemView_virtualbase_ViewportSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemView) OnViewportSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ViewportSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ViewportSizeHint
func miqt_exec_callback_QAbstractItemView_ViewportSizeHint(self QAbstractItemView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_ViewportSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QAbstractItemView_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemView) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_MinimumSizeHint
func miqt_exec_callback_QAbstractItemView_MinimumSizeHint(self QAbstractItemView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QAbstractItemView_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractItemView) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SizeHint
func miqt_exec_callback_QAbstractItemView_SizeHint(self QAbstractItemView, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractItemView{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QAbstractItemView) callVirtualBase_SetupViewport(viewport *QWidget) {

	QAbstractItemView_virtualbase_SetupViewport(unsafe.Pointer(this.h), viewport.cPointer())

}
func (this *QAbstractItemView) OnSetupViewport(slot func(super func(viewport *QWidget), viewport *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_SetupViewport(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_SetupViewport
func miqt_exec_callback_QAbstractItemView_SetupViewport(self QAbstractItemView, cb intptr_t, viewport *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(viewport *QWidget), viewport *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(viewport)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_SetupViewport, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QAbstractItemView_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractItemView) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_PaintEvent
func miqt_exec_callback_QAbstractItemView_PaintEvent(self QAbstractItemView, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_WheelEvent(param1 *QWheelEvent) {

	QAbstractItemView_virtualbase_WheelEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractItemView) OnWheelEvent(slot func(super func(param1 *QWheelEvent), param1 *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_WheelEvent
func miqt_exec_callback_QAbstractItemView_WheelEvent(self QAbstractItemView, cb intptr_t, param1 *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QWheelEvent), param1 *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(param1)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QAbstractItemView_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QAbstractItemView) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ContextMenuEvent
func miqt_exec_callback_QAbstractItemView_ContextMenuEvent(self QAbstractItemView, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QAbstractItemView) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QAbstractItemView_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QAbstractItemView) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractItemView_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractItemView_ScrollContentsBy
func miqt_exec_callback_QAbstractItemView_ScrollContentsBy(self QAbstractItemView, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QAbstractItemView{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}
