package qt

import (
	"unsafe"
)

type QTabBar__Shape int

const (
	QTabBar__RoundedNorth    QTabBar__Shape = 0
	QTabBar__RoundedSouth    QTabBar__Shape = 1
	QTabBar__RoundedWest     QTabBar__Shape = 2
	QTabBar__RoundedEast     QTabBar__Shape = 3
	QTabBar__TriangularNorth QTabBar__Shape = 4
	QTabBar__TriangularSouth QTabBar__Shape = 5
	QTabBar__TriangularWest  QTabBar__Shape = 6
	QTabBar__TriangularEast  QTabBar__Shape = 7
)

type QTabBar__ButtonPosition int

const (
	QTabBar__LeftSide  QTabBar__ButtonPosition = 0
	QTabBar__RightSide QTabBar__ButtonPosition = 1
)

type QTabBar__SelectionBehavior int

const (
	QTabBar__SelectLeftTab     QTabBar__SelectionBehavior = 0
	QTabBar__SelectRightTab    QTabBar__SelectionBehavior = 1
	QTabBar__SelectPreviousTab QTabBar__SelectionBehavior = 2
)

type QTabBar struct {
	h          uintptr
	isSubclass bool
}

// NewQTabBar constructs a new QTabBar object.
func NewQTabBar(parent *QWidget) *QTabBar {
	g := newQTabBar(QTabBar_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQTabBar2 constructs a new QTabBar object.
func NewQTabBar2() *QTabBar {
	g := newQTabBar(QTabBar_new2())
	g.isSubclass = true
	return g
}

func (this *QTabBar) MetaObject() *QMetaObject {
	return newQMetaObject(QTabBar_MetaObject(this.h))
}

func (this *QTabBar) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTabBar_Metacast(this.h, param1_Cstring))
}

func QTabBar_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTabBar_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabBar) Shape() Shape {
	xxxxxxxxx
}

func (this *QTabBar) SetShape(shape Shape) {
	QTabBar_SetShape(this.h, shape)
}

func (this *QTabBar) AddTab(text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QTabBar_AddTab(this.h, text_ms))
}

func (this *QTabBar) AddTab2(icon *QIcon, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QTabBar_AddTab2(this.h, icon.cPointer(), text_ms))
}

func (this *QTabBar) InsertTab(index int, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QTabBar_InsertTab(this.h, (int)(index), text_ms))
}

func (this *QTabBar) InsertTab2(index int, icon *QIcon, text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QTabBar_InsertTab2(this.h, (int)(index), icon.cPointer(), text_ms))
}

func (this *QTabBar) RemoveTab(index int) {
	QTabBar_RemoveTab(this.h, (int)(index))
}

func (this *QTabBar) MoveTab(from int, to int) {
	QTabBar_MoveTab(this.h, (int)(from), (int)(to))
}

func (this *QTabBar) IsTabEnabled(index int) bool {
	return (bool)(QTabBar_IsTabEnabled(this.h, (int)(index)))
}

func (this *QTabBar) SetTabEnabled(index int, enabled bool) {
	QTabBar_SetTabEnabled(this.h, (int)(index), (bool)(enabled))
}

func (this *QTabBar) IsTabVisible(index int) bool {
	return (bool)(QTabBar_IsTabVisible(this.h, (int)(index)))
}

func (this *QTabBar) SetTabVisible(index int, visible bool) {
	QTabBar_SetTabVisible(this.h, (int)(index), (bool)(visible))
}

func (this *QTabBar) TabText(index int) string {
	var _ms struct_miqt_string = QTabBar_TabText(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabBar) SetTabText(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTabBar_SetTabText(this.h, (int)(index), text_ms)
}

func (this *QTabBar) TabTextColor(index int) *QColor {
	_goptr := newQColor(QTabBar_TabTextColor(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) SetTabTextColor(index int, color *QColor) {
	QTabBar_SetTabTextColor(this.h, (int)(index), color.cPointer())
}

func (this *QTabBar) TabIcon(index int) *QIcon {
	_goptr := newQIcon(QTabBar_TabIcon(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) SetTabIcon(index int, icon *QIcon) {
	QTabBar_SetTabIcon(this.h, (int)(index), icon.cPointer())
}

func (this *QTabBar) ElideMode() TextElideMode {
	return (TextElideMode)(QTabBar_ElideMode(this.h))
}

func (this *QTabBar) SetElideMode(mode TextElideMode) {
	QTabBar_SetElideMode(this.h, (int)(mode))
}

func (this *QTabBar) SetTabToolTip(index int, tip string) {
	tip_ms := struct_miqt_string{}
	tip_ms.data = CString(tip)
	tip_ms.len = size_t(len(tip))
	defer free(unsafe.Pointer(tip_ms.data))
	QTabBar_SetTabToolTip(this.h, (int)(index), tip_ms)
}

func (this *QTabBar) TabToolTip(index int) string {
	var _ms struct_miqt_string = QTabBar_TabToolTip(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabBar) SetTabWhatsThis(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTabBar_SetTabWhatsThis(this.h, (int)(index), text_ms)
}

func (this *QTabBar) TabWhatsThis(index int) string {
	var _ms struct_miqt_string = QTabBar_TabWhatsThis(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabBar) SetTabData(index int, data *QVariant) {
	QTabBar_SetTabData(this.h, (int)(index), data.cPointer())
}

func (this *QTabBar) TabData(index int) *QVariant {
	_goptr := newQVariant(QTabBar_TabData(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) TabRect(index int) *QRect {
	_goptr := newQRect(QTabBar_TabRect(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) TabAt(pos *QPoint) int {
	return (int)(QTabBar_TabAt(this.h, pos.cPointer()))
}

func (this *QTabBar) CurrentIndex() int {
	return (int)(QTabBar_CurrentIndex(this.h))
}

func (this *QTabBar) Count() int {
	return (int)(QTabBar_Count(this.h))
}

func (this *QTabBar) SizeHint() *QSize {
	_goptr := newQSize(QTabBar_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) MinimumSizeHint() *QSize {
	_goptr := newQSize(QTabBar_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) SetDrawBase(drawTheBase bool) {
	QTabBar_SetDrawBase(this.h, (bool)(drawTheBase))
}

func (this *QTabBar) DrawBase() bool {
	return (bool)(QTabBar_DrawBase(this.h))
}

func (this *QTabBar) IconSize() *QSize {
	_goptr := newQSize(QTabBar_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabBar) SetIconSize(size *QSize) {
	QTabBar_SetIconSize(this.h, size.cPointer())
}

func (this *QTabBar) UsesScrollButtons() bool {
	return (bool)(QTabBar_UsesScrollButtons(this.h))
}

func (this *QTabBar) SetUsesScrollButtons(useButtons bool) {
	QTabBar_SetUsesScrollButtons(this.h, (bool)(useButtons))
}

func (this *QTabBar) TabsClosable() bool {
	return (bool)(QTabBar_TabsClosable(this.h))
}

func (this *QTabBar) SetTabsClosable(closable bool) {
	QTabBar_SetTabsClosable(this.h, (bool)(closable))
}

func (this *QTabBar) SetTabButton(index int, position ButtonPosition, widget *QWidget) {
	QTabBar_SetTabButton(this.h, (int)(index), position, widget.cPointer())
}

func (this *QTabBar) TabButton(index int, position ButtonPosition) *QWidget {
	return newQWidget(QTabBar_TabButton(this.h, (int)(index), position))
}

func (this *QTabBar) SelectionBehaviorOnRemove() SelectionBehavior {
	xxxxxxxxx
}

func (this *QTabBar) SetSelectionBehaviorOnRemove(behavior SelectionBehavior) {
	QTabBar_SetSelectionBehaviorOnRemove(this.h, behavior)
}

func (this *QTabBar) Expanding() bool {
	return (bool)(QTabBar_Expanding(this.h))
}

func (this *QTabBar) SetExpanding(enabled bool) {
	QTabBar_SetExpanding(this.h, (bool)(enabled))
}

func (this *QTabBar) IsMovable() bool {
	return (bool)(QTabBar_IsMovable(this.h))
}

func (this *QTabBar) SetMovable(movable bool) {
	QTabBar_SetMovable(this.h, (bool)(movable))
}

func (this *QTabBar) DocumentMode() bool {
	return (bool)(QTabBar_DocumentMode(this.h))
}

func (this *QTabBar) SetDocumentMode(set bool) {
	QTabBar_SetDocumentMode(this.h, (bool)(set))
}

func (this *QTabBar) AutoHide() bool {
	return (bool)(QTabBar_AutoHide(this.h))
}

func (this *QTabBar) SetAutoHide(hide bool) {
	QTabBar_SetAutoHide(this.h, (bool)(hide))
}

func (this *QTabBar) ChangeCurrentOnDrag() bool {
	return (bool)(QTabBar_ChangeCurrentOnDrag(this.h))
}

func (this *QTabBar) SetChangeCurrentOnDrag(change bool) {
	QTabBar_SetChangeCurrentOnDrag(this.h, (bool)(change))
}

func (this *QTabBar) SetCurrentIndex(index int) {
	QTabBar_SetCurrentIndex(this.h, (int)(index))
}

func (this *QTabBar) CurrentChanged(index int) {
	QTabBar_CurrentChanged(this.h, (int)(index))
}

func (this *QTabBar) OnCurrentChanged(slot func(index int)) {
	QTabBar_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_CurrentChanged
func miqt_exec_callback_QTabBar_CurrentChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QTabBar) TabCloseRequested(index int) {
	QTabBar_TabCloseRequested(this.h, (int)(index))
}

func (this *QTabBar) OnTabCloseRequested(slot func(index int)) {
	QTabBar_connect_TabCloseRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabCloseRequested
func miqt_exec_callback_QTabBar_TabCloseRequested(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QTabBar) TabMoved(from int, to int) {
	QTabBar_TabMoved(this.h, (int)(from), (int)(to))
}

func (this *QTabBar) OnTabMoved(slot func(from int, to int)) {
	QTabBar_connect_TabMoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabMoved
func miqt_exec_callback_QTabBar_TabMoved(cb intptr_t, from int, to int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(from int, to int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(from)

	slotval2 := (int)(to)

	gofunc(slotval1, slotval2)
}

func (this *QTabBar) TabBarClicked(index int) {
	QTabBar_TabBarClicked(this.h, (int)(index))
}

func (this *QTabBar) OnTabBarClicked(slot func(index int)) {
	QTabBar_connect_TabBarClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabBarClicked
func miqt_exec_callback_QTabBar_TabBarClicked(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QTabBar) TabBarDoubleClicked(index int) {
	QTabBar_TabBarDoubleClicked(this.h, (int)(index))
}

func (this *QTabBar) OnTabBarDoubleClicked(slot func(index int)) {
	QTabBar_connect_TabBarDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabBarDoubleClicked
func miqt_exec_callback_QTabBar_TabBarDoubleClicked(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func QTabBar_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTabBar_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTabBar_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTabBar_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabBar) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTabBar_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTabBar) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MetaObject
func miqt_exec_callback_QTabBar_MetaObject(self QTabBar, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTabBar) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTabBar_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTabBar) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_Metacast
func miqt_exec_callback_QTabBar_Metacast(self QTabBar, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
