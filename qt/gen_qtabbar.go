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

	ret := newQTabBar(QTabBar_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTabBar2 constructs a new QTabBar object.
func NewQTabBar2() *QTabBar {

	ret := newQTabBar(QTabBar_new2())
	ret.isSubclass = true
	return ret
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

func (this *QTabBar) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QTabBar_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabBar) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_SizeHint
func miqt_exec_callback_QTabBar_SizeHint(self QTabBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QTabBar_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabBar) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MinimumSizeHint
func miqt_exec_callback_QTabBar_MinimumSizeHint(self QTabBar, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_TabSizeHint(index int) *QSize {

	_goptr := newQSize(QTabBar_virtualbase_TabSizeHint(unsafe.Pointer(this.h), (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabBar) OnTabSizeHint(slot func(super func(index int) *QSize, index int) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_TabSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabSizeHint
func miqt_exec_callback_QTabBar_TabSizeHint(self QTabBar, cb intptr_t, index int) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QSize, index int) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_TabSizeHint, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_MinimumTabSizeHint(index int) *QSize {

	_goptr := newQSize(QTabBar_virtualbase_MinimumTabSizeHint(unsafe.Pointer(this.h), (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabBar) OnMinimumTabSizeHint(slot func(super func(index int) *QSize, index int) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MinimumTabSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MinimumTabSizeHint
func miqt_exec_callback_QTabBar_MinimumTabSizeHint(self QTabBar, cb intptr_t, index int) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int) *QSize, index int) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_MinimumTabSizeHint, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_TabInserted(index int) {

	QTabBar_virtualbase_TabInserted(unsafe.Pointer(this.h), (int)(index))

}
func (this *QTabBar) OnTabInserted(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_TabInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabInserted
func miqt_exec_callback_QTabBar_TabInserted(self QTabBar, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QTabBar{h: self}).callVirtualBase_TabInserted, slotval1)

}

func (this *QTabBar) callVirtualBase_TabRemoved(index int) {

	QTabBar_virtualbase_TabRemoved(unsafe.Pointer(this.h), (int)(index))

}
func (this *QTabBar) OnTabRemoved(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_TabRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabRemoved
func miqt_exec_callback_QTabBar_TabRemoved(self QTabBar, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QTabBar{h: self}).callVirtualBase_TabRemoved, slotval1)

}

func (this *QTabBar) callVirtualBase_TabLayoutChange() {

	QTabBar_virtualbase_TabLayoutChange(unsafe.Pointer(this.h))

}
func (this *QTabBar) OnTabLayoutChange(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_TabLayoutChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabLayoutChange
func miqt_exec_callback_QTabBar_TabLayoutChange(self QTabBar, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTabBar{h: self}).callVirtualBase_TabLayoutChange)

}

func (this *QTabBar) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QTabBar_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QTabBar) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_Event
func miqt_exec_callback_QTabBar_Event(self QTabBar, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTabBar) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QTabBar_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_ResizeEvent
func miqt_exec_callback_QTabBar_ResizeEvent(self QTabBar, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QTabBar_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_ShowEvent
func miqt_exec_callback_QTabBar_ShowEvent(self QTabBar, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_HideEvent(param1 *QHideEvent) {

	QTabBar_virtualbase_HideEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnHideEvent(slot func(super func(param1 *QHideEvent), param1 *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_HideEvent
func miqt_exec_callback_QTabBar_HideEvent(self QTabBar, cb intptr_t, param1 *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QHideEvent), param1 *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QTabBar_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_PaintEvent
func miqt_exec_callback_QTabBar_PaintEvent(self QTabBar, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_MousePressEvent(param1 *QMouseEvent) {

	QTabBar_virtualbase_MousePressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnMousePressEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MousePressEvent
func miqt_exec_callback_QTabBar_MousePressEvent(self QTabBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_MouseMoveEvent(param1 *QMouseEvent) {

	QTabBar_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnMouseMoveEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MouseMoveEvent
func miqt_exec_callback_QTabBar_MouseMoveEvent(self QTabBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_MouseReleaseEvent(param1 *QMouseEvent) {

	QTabBar_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnMouseReleaseEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MouseReleaseEvent
func miqt_exec_callback_QTabBar_MouseReleaseEvent(self QTabBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_MouseDoubleClickEvent(param1 *QMouseEvent) {

	QTabBar_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnMouseDoubleClickEvent(slot func(super func(param1 *QMouseEvent), param1 *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MouseDoubleClickEvent
func miqt_exec_callback_QTabBar_MouseDoubleClickEvent(self QTabBar, cb intptr_t, param1 *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QMouseEvent), param1 *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QTabBar_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_WheelEvent
func miqt_exec_callback_QTabBar_WheelEvent(self QTabBar, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QTabBar_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_KeyPressEvent
func miqt_exec_callback_QTabBar_KeyPressEvent(self QTabBar, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QTabBar_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_ChangeEvent
func miqt_exec_callback_QTabBar_ChangeEvent(self QTabBar, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTabBar_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TimerEvent
func miqt_exec_callback_QTabBar_TimerEvent(self QTabBar, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_InitStyleOption(option *QStyleOptionTab, tabIndex int) {

	QTabBar_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer(), (int)(tabIndex))

}
func (this *QTabBar) OnInitStyleOption(slot func(super func(option *QStyleOptionTab, tabIndex int), option *QStyleOptionTab, tabIndex int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_InitStyleOption
func miqt_exec_callback_QTabBar_InitStyleOption(self QTabBar, cb intptr_t, option *QStyleOptionTab, tabIndex int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionTab, tabIndex int), option *QStyleOptionTab, tabIndex int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionTab(option)

	slotval2 := (int)(tabIndex)

	gofunc((&QTabBar{h: self}).callVirtualBase_InitStyleOption, slotval1, slotval2)

}

func (this *QTabBar) callVirtualBase_DevType() int {

	return (int)(QTabBar_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QTabBar) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_DevType
func miqt_exec_callback_QTabBar_DevType(self QTabBar, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QTabBar) callVirtualBase_SetVisible(visible bool) {

	QTabBar_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QTabBar) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_SetVisible
func miqt_exec_callback_QTabBar_SetVisible(self QTabBar, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QTabBar{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QTabBar) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QTabBar_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QTabBar) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_HeightForWidth
func miqt_exec_callback_QTabBar_HeightForWidth(self QTabBar, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QTabBar) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QTabBar_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QTabBar) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_HasHeightForWidth
func miqt_exec_callback_QTabBar_HasHeightForWidth(self QTabBar, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QTabBar) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QTabBar_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QTabBar) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_PaintEngine
func miqt_exec_callback_QTabBar_PaintEngine(self QTabBar, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QTabBar_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_KeyReleaseEvent
func miqt_exec_callback_QTabBar_KeyReleaseEvent(self QTabBar, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QTabBar_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_FocusInEvent
func miqt_exec_callback_QTabBar_FocusInEvent(self QTabBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QTabBar_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_FocusOutEvent
func miqt_exec_callback_QTabBar_FocusOutEvent(self QTabBar, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QTabBar_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_EnterEvent
func miqt_exec_callback_QTabBar_EnterEvent(self QTabBar, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_LeaveEvent(event *QEvent) {

	QTabBar_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_LeaveEvent
func miqt_exec_callback_QTabBar_LeaveEvent(self QTabBar, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QTabBar_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_MoveEvent
func miqt_exec_callback_QTabBar_MoveEvent(self QTabBar, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QTabBar_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_CloseEvent
func miqt_exec_callback_QTabBar_CloseEvent(self QTabBar, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QTabBar_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_ContextMenuEvent
func miqt_exec_callback_QTabBar_ContextMenuEvent(self QTabBar, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QTabBar_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_TabletEvent
func miqt_exec_callback_QTabBar_TabletEvent(self QTabBar, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_ActionEvent(event *QActionEvent) {

	QTabBar_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_ActionEvent
func miqt_exec_callback_QTabBar_ActionEvent(self QTabBar, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QTabBar_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_DragEnterEvent
func miqt_exec_callback_QTabBar_DragEnterEvent(self QTabBar, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QTabBar_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_DragMoveEvent
func miqt_exec_callback_QTabBar_DragMoveEvent(self QTabBar, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QTabBar_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_DragLeaveEvent
func miqt_exec_callback_QTabBar_DragLeaveEvent(self QTabBar, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_DropEvent(event *QDropEvent) {

	QTabBar_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabBar) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_DropEvent
func miqt_exec_callback_QTabBar_DropEvent(self QTabBar, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QTabBar{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QTabBar_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QTabBar) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_NativeEvent
func miqt_exec_callback_QTabBar_NativeEvent(self QTabBar, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QTabBar) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QTabBar_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QTabBar) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_Metric
func miqt_exec_callback_QTabBar_Metric(self QTabBar, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QTabBar) callVirtualBase_InitPainter(painter *QPainter) {

	QTabBar_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QTabBar) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_InitPainter
func miqt_exec_callback_QTabBar_InitPainter(self QTabBar, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QTabBar{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QTabBar) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QTabBar_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QTabBar) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_Redirected
func miqt_exec_callback_QTabBar_Redirected(self QTabBar, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QTabBar_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QTabBar) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_SharedPainter
func miqt_exec_callback_QTabBar_SharedPainter(self QTabBar, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QTabBar_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabBar) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_InputMethodEvent
func miqt_exec_callback_QTabBar_InputMethodEvent(self QTabBar, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QTabBar{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTabBar) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QTabBar_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabBar) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_InputMethodQuery
func miqt_exec_callback_QTabBar_InputMethodQuery(self QTabBar, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTabBar) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTabBar_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTabBar) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabBar_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabBar_FocusNextPrevChild
func miqt_exec_callback_QTabBar_FocusNextPrevChild(self QTabBar, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTabBar{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
