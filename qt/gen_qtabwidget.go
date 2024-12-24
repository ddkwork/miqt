package qt

import (
	"unsafe"
)

type QTabWidget__TabPosition int

const (
	QTabWidget__North QTabWidget__TabPosition = 0
	QTabWidget__South QTabWidget__TabPosition = 1
	QTabWidget__West  QTabWidget__TabPosition = 2
	QTabWidget__East  QTabWidget__TabPosition = 3
)

type QTabWidget__TabShape int

const (
	QTabWidget__Rounded    QTabWidget__TabShape = 0
	QTabWidget__Triangular QTabWidget__TabShape = 1
)

type QTabWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQTabWidget constructs a new QTabWidget object.
func NewQTabWidget(parent *QWidget) *QTabWidget {

	ret := newQTabWidget(QTabWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTabWidget2 constructs a new QTabWidget object.
func NewQTabWidget2() *QTabWidget {

	ret := newQTabWidget(QTabWidget_new2())
	ret.isSubclass = true
	return ret
}

func (this *QTabWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QTabWidget_MetaObject(this.h))
}

func (this *QTabWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTabWidget_Metacast(this.h, param1_Cstring))
}

func QTabWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTabWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabWidget) AddTab(widget *QWidget, param2 string) int {
	param2_ms := struct_miqt_string{}
	param2_ms.data = CString(param2)
	param2_ms.len = size_t(len(param2))
	defer free(unsafe.Pointer(param2_ms.data))
	return (int)(QTabWidget_AddTab(this.h, widget.cPointer(), param2_ms))
}

func (this *QTabWidget) AddTab2(widget *QWidget, icon *QIcon, label string) int {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QTabWidget_AddTab2(this.h, widget.cPointer(), icon.cPointer(), label_ms))
}

func (this *QTabWidget) InsertTab(index int, widget *QWidget, param3 string) int {
	param3_ms := struct_miqt_string{}
	param3_ms.data = CString(param3)
	param3_ms.len = size_t(len(param3))
	defer free(unsafe.Pointer(param3_ms.data))
	return (int)(QTabWidget_InsertTab(this.h, (int)(index), widget.cPointer(), param3_ms))
}

func (this *QTabWidget) InsertTab2(index int, widget *QWidget, icon *QIcon, label string) int {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QTabWidget_InsertTab2(this.h, (int)(index), widget.cPointer(), icon.cPointer(), label_ms))
}

func (this *QTabWidget) RemoveTab(index int) {
	QTabWidget_RemoveTab(this.h, (int)(index))
}

func (this *QTabWidget) IsTabEnabled(index int) bool {
	return (bool)(QTabWidget_IsTabEnabled(this.h, (int)(index)))
}

func (this *QTabWidget) SetTabEnabled(index int, enabled bool) {
	QTabWidget_SetTabEnabled(this.h, (int)(index), (bool)(enabled))
}

func (this *QTabWidget) IsTabVisible(index int) bool {
	return (bool)(QTabWidget_IsTabVisible(this.h, (int)(index)))
}

func (this *QTabWidget) SetTabVisible(index int, visible bool) {
	QTabWidget_SetTabVisible(this.h, (int)(index), (bool)(visible))
}

func (this *QTabWidget) TabText(index int) string {
	var _ms struct_miqt_string = QTabWidget_TabText(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabWidget) SetTabText(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTabWidget_SetTabText(this.h, (int)(index), text_ms)
}

func (this *QTabWidget) TabIcon(index int) *QIcon {
	_goptr := newQIcon(QTabWidget_TabIcon(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabWidget) SetTabIcon(index int, icon *QIcon) {
	QTabWidget_SetTabIcon(this.h, (int)(index), icon.cPointer())
}

func (this *QTabWidget) SetTabToolTip(index int, tip string) {
	tip_ms := struct_miqt_string{}
	tip_ms.data = CString(tip)
	tip_ms.len = size_t(len(tip))
	defer free(unsafe.Pointer(tip_ms.data))
	QTabWidget_SetTabToolTip(this.h, (int)(index), tip_ms)
}

func (this *QTabWidget) TabToolTip(index int) string {
	var _ms struct_miqt_string = QTabWidget_TabToolTip(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabWidget) SetTabWhatsThis(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QTabWidget_SetTabWhatsThis(this.h, (int)(index), text_ms)
}

func (this *QTabWidget) TabWhatsThis(index int) string {
	var _ms struct_miqt_string = QTabWidget_TabWhatsThis(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabWidget) CurrentIndex() int {
	return (int)(QTabWidget_CurrentIndex(this.h))
}

func (this *QTabWidget) CurrentWidget() *QWidget {
	return newQWidget(QTabWidget_CurrentWidget(this.h))
}

func (this *QTabWidget) Widget(index int) *QWidget {
	return newQWidget(QTabWidget_Widget(this.h, (int)(index)))
}

func (this *QTabWidget) IndexOf(widget *QWidget) int {
	return (int)(QTabWidget_IndexOf(this.h, widget.cPointer()))
}

func (this *QTabWidget) Count() int {
	return (int)(QTabWidget_Count(this.h))
}

func (this *QTabWidget) TabPosition() TabPosition {
	xxxxxxxxx
}

func (this *QTabWidget) SetTabPosition(position TabPosition) {
	QTabWidget_SetTabPosition(this.h, position)
}

func (this *QTabWidget) TabsClosable() bool {
	return (bool)(QTabWidget_TabsClosable(this.h))
}

func (this *QTabWidget) SetTabsClosable(closeable bool) {
	QTabWidget_SetTabsClosable(this.h, (bool)(closeable))
}

func (this *QTabWidget) IsMovable() bool {
	return (bool)(QTabWidget_IsMovable(this.h))
}

func (this *QTabWidget) SetMovable(movable bool) {
	QTabWidget_SetMovable(this.h, (bool)(movable))
}

func (this *QTabWidget) TabShape() TabShape {
	xxxxxxxxx
}

func (this *QTabWidget) SetTabShape(s TabShape) {
	QTabWidget_SetTabShape(this.h, s)
}

func (this *QTabWidget) SizeHint() *QSize {
	_goptr := newQSize(QTabWidget_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabWidget) MinimumSizeHint() *QSize {
	_goptr := newQSize(QTabWidget_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabWidget) HeightForWidth(width int) int {
	return (int)(QTabWidget_HeightForWidth(this.h, (int)(width)))
}

func (this *QTabWidget) HasHeightForWidth() bool {
	return (bool)(QTabWidget_HasHeightForWidth(this.h))
}

func (this *QTabWidget) SetCornerWidget(w *QWidget) {
	QTabWidget_SetCornerWidget(this.h, w.cPointer())
}

func (this *QTabWidget) CornerWidget() *QWidget {
	return newQWidget(QTabWidget_CornerWidget(this.h))
}

func (this *QTabWidget) ElideMode() TextElideMode {
	return (TextElideMode)(QTabWidget_ElideMode(this.h))
}

func (this *QTabWidget) SetElideMode(mode TextElideMode) {
	QTabWidget_SetElideMode(this.h, (int)(mode))
}

func (this *QTabWidget) IconSize() *QSize {
	_goptr := newQSize(QTabWidget_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTabWidget) SetIconSize(size *QSize) {
	QTabWidget_SetIconSize(this.h, size.cPointer())
}

func (this *QTabWidget) UsesScrollButtons() bool {
	return (bool)(QTabWidget_UsesScrollButtons(this.h))
}

func (this *QTabWidget) SetUsesScrollButtons(useButtons bool) {
	QTabWidget_SetUsesScrollButtons(this.h, (bool)(useButtons))
}

func (this *QTabWidget) DocumentMode() bool {
	return (bool)(QTabWidget_DocumentMode(this.h))
}

func (this *QTabWidget) SetDocumentMode(set bool) {
	QTabWidget_SetDocumentMode(this.h, (bool)(set))
}

func (this *QTabWidget) TabBarAutoHide() bool {
	return (bool)(QTabWidget_TabBarAutoHide(this.h))
}

func (this *QTabWidget) SetTabBarAutoHide(enabled bool) {
	QTabWidget_SetTabBarAutoHide(this.h, (bool)(enabled))
}

func (this *QTabWidget) Clear() {
	QTabWidget_Clear(this.h)
}

func (this *QTabWidget) TabBar() *QTabBar {
	return newQTabBar(QTabWidget_TabBar(this.h))
}

func (this *QTabWidget) SetCurrentIndex(index int) {
	QTabWidget_SetCurrentIndex(this.h, (int)(index))
}

func (this *QTabWidget) SetCurrentWidget(widget *QWidget) {
	QTabWidget_SetCurrentWidget(this.h, widget.cPointer())
}

func (this *QTabWidget) CurrentChanged(index int) {
	QTabWidget_CurrentChanged(this.h, (int)(index))
}
func (this *QTabWidget) OnCurrentChanged(slot func(index int)) {
	QTabWidget_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_CurrentChanged
func miqt_exec_callback_QTabWidget_CurrentChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QTabWidget) TabCloseRequested(index int) {
	QTabWidget_TabCloseRequested(this.h, (int)(index))
}
func (this *QTabWidget) OnTabCloseRequested(slot func(index int)) {
	QTabWidget_connect_TabCloseRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_TabCloseRequested
func miqt_exec_callback_QTabWidget_TabCloseRequested(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QTabWidget) TabBarClicked(index int) {
	QTabWidget_TabBarClicked(this.h, (int)(index))
}
func (this *QTabWidget) OnTabBarClicked(slot func(index int)) {
	QTabWidget_connect_TabBarClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_TabBarClicked
func miqt_exec_callback_QTabWidget_TabBarClicked(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QTabWidget) TabBarDoubleClicked(index int) {
	QTabWidget_TabBarDoubleClicked(this.h, (int)(index))
}
func (this *QTabWidget) OnTabBarDoubleClicked(slot func(index int)) {
	QTabWidget_connect_TabBarDoubleClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_TabBarDoubleClicked
func miqt_exec_callback_QTabWidget_TabBarDoubleClicked(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func QTabWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTabWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTabWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTabWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTabWidget) SetCornerWidget2(w *QWidget, corner Corner) {
	QTabWidget_SetCornerWidget2(this.h, w.cPointer(), (int)(corner))
}

func (this *QTabWidget) CornerWidget1(corner Corner) *QWidget {
	return newQWidget(QTabWidget_CornerWidget1(this.h, (int)(corner)))
}

func (this *QTabWidget) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QTabWidget_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabWidget) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_SizeHint
func miqt_exec_callback_QTabWidget_SizeHint(self QTabWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QTabWidget) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QTabWidget_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabWidget) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MinimumSizeHint
func miqt_exec_callback_QTabWidget_MinimumSizeHint(self QTabWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QTabWidget) callVirtualBase_HeightForWidth(width int) int {

	return (int)(QTabWidget_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(width)))

}
func (this *QTabWidget) OnHeightForWidth(slot func(super func(width int) int, width int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_HeightForWidth
func miqt_exec_callback_QTabWidget_HeightForWidth(self QTabWidget, cb intptr_t, width int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(width int) int, width int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(width)

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QTabWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QTabWidget_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QTabWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_HasHeightForWidth
func miqt_exec_callback_QTabWidget_HasHeightForWidth(self QTabWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QTabWidget) callVirtualBase_TabInserted(index int) {

	QTabWidget_virtualbase_TabInserted(unsafe.Pointer(this.h), (int)(index))

}
func (this *QTabWidget) OnTabInserted(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_TabInserted(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_TabInserted
func miqt_exec_callback_QTabWidget_TabInserted(self QTabWidget, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QTabWidget{h: self}).callVirtualBase_TabInserted, slotval1)

}

func (this *QTabWidget) callVirtualBase_TabRemoved(index int) {

	QTabWidget_virtualbase_TabRemoved(unsafe.Pointer(this.h), (int)(index))

}
func (this *QTabWidget) OnTabRemoved(slot func(super func(index int), index int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_TabRemoved(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_TabRemoved
func miqt_exec_callback_QTabWidget_TabRemoved(self QTabWidget, cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(index int), index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc((&QTabWidget{h: self}).callVirtualBase_TabRemoved, slotval1)

}

func (this *QTabWidget) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QTabWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabWidget) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_ShowEvent
func miqt_exec_callback_QTabWidget_ShowEvent(self QTabWidget, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QTabWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QTabWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabWidget) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_ResizeEvent
func miqt_exec_callback_QTabWidget_ResizeEvent(self QTabWidget, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QTabWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QTabWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabWidget) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_KeyPressEvent
func miqt_exec_callback_QTabWidget_KeyPressEvent(self QTabWidget, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QTabWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_PaintEvent(param1 *QPaintEvent) {

	QTabWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabWidget) OnPaintEvent(slot func(super func(param1 *QPaintEvent), param1 *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_PaintEvent
func miqt_exec_callback_QTabWidget_PaintEvent(self QTabWidget, cb intptr_t, param1 *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPaintEvent), param1 *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(param1)

	gofunc((&QTabWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QTabWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabWidget) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_ChangeEvent
func miqt_exec_callback_QTabWidget_ChangeEvent(self QTabWidget, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QTabWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QTabWidget_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QTabWidget) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_Event
func miqt_exec_callback_QTabWidget_Event(self QTabWidget, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTabWidget) callVirtualBase_InitStyleOption(option *QStyleOptionTabWidgetFrame) {

	QTabWidget_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QTabWidget) OnInitStyleOption(slot func(super func(option *QStyleOptionTabWidgetFrame), option *QStyleOptionTabWidgetFrame)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_InitStyleOption
func miqt_exec_callback_QTabWidget_InitStyleOption(self QTabWidget, cb intptr_t, option *QStyleOptionTabWidgetFrame) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionTabWidgetFrame), option *QStyleOptionTabWidgetFrame))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionTabWidgetFrame(option)

	gofunc((&QTabWidget{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QTabWidget) callVirtualBase_DevType() int {

	return (int)(QTabWidget_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QTabWidget) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_DevType
func miqt_exec_callback_QTabWidget_DevType(self QTabWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QTabWidget) callVirtualBase_SetVisible(visible bool) {

	QTabWidget_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QTabWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_SetVisible
func miqt_exec_callback_QTabWidget_SetVisible(self QTabWidget, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QTabWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QTabWidget) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QTabWidget_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QTabWidget) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_PaintEngine
func miqt_exec_callback_QTabWidget_PaintEngine(self QTabWidget, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QTabWidget) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QTabWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MousePressEvent
func miqt_exec_callback_QTabWidget_MousePressEvent(self QTabWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QTabWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MouseReleaseEvent
func miqt_exec_callback_QTabWidget_MouseReleaseEvent(self QTabWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QTabWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MouseDoubleClickEvent
func miqt_exec_callback_QTabWidget_MouseDoubleClickEvent(self QTabWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QTabWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MouseMoveEvent
func miqt_exec_callback_QTabWidget_MouseMoveEvent(self QTabWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QTabWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_WheelEvent
func miqt_exec_callback_QTabWidget_WheelEvent(self QTabWidget, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QTabWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_KeyReleaseEvent
func miqt_exec_callback_QTabWidget_KeyReleaseEvent(self QTabWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QTabWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_FocusInEvent
func miqt_exec_callback_QTabWidget_FocusInEvent(self QTabWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QTabWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_FocusOutEvent
func miqt_exec_callback_QTabWidget_FocusOutEvent(self QTabWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QTabWidget_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_EnterEvent
func miqt_exec_callback_QTabWidget_EnterEvent(self QTabWidget, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_LeaveEvent(event *QEvent) {

	QTabWidget_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_LeaveEvent
func miqt_exec_callback_QTabWidget_LeaveEvent(self QTabWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QTabWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MoveEvent
func miqt_exec_callback_QTabWidget_MoveEvent(self QTabWidget, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QTabWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_CloseEvent
func miqt_exec_callback_QTabWidget_CloseEvent(self QTabWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QTabWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_ContextMenuEvent
func miqt_exec_callback_QTabWidget_ContextMenuEvent(self QTabWidget, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QTabWidget_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_TabletEvent
func miqt_exec_callback_QTabWidget_TabletEvent(self QTabWidget, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_ActionEvent(event *QActionEvent) {

	QTabWidget_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_ActionEvent
func miqt_exec_callback_QTabWidget_ActionEvent(self QTabWidget, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QTabWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_DragEnterEvent
func miqt_exec_callback_QTabWidget_DragEnterEvent(self QTabWidget, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QTabWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_DragMoveEvent
func miqt_exec_callback_QTabWidget_DragMoveEvent(self QTabWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QTabWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_DragLeaveEvent
func miqt_exec_callback_QTabWidget_DragLeaveEvent(self QTabWidget, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_DropEvent(event *QDropEvent) {

	QTabWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_DropEvent
func miqt_exec_callback_QTabWidget_DropEvent(self QTabWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_HideEvent(event *QHideEvent) {

	QTabWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTabWidget) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_HideEvent
func miqt_exec_callback_QTabWidget_HideEvent(self QTabWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QTabWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QTabWidget_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QTabWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_NativeEvent
func miqt_exec_callback_QTabWidget_NativeEvent(self QTabWidget, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QTabWidget) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QTabWidget_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QTabWidget) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_Metric
func miqt_exec_callback_QTabWidget_Metric(self QTabWidget, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QTabWidget) callVirtualBase_InitPainter(painter *QPainter) {

	QTabWidget_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QTabWidget) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_InitPainter
func miqt_exec_callback_QTabWidget_InitPainter(self QTabWidget, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QTabWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QTabWidget) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QTabWidget_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QTabWidget) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_Redirected
func miqt_exec_callback_QTabWidget_Redirected(self QTabWidget, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTabWidget) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QTabWidget_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QTabWidget) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_SharedPainter
func miqt_exec_callback_QTabWidget_SharedPainter(self QTabWidget, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QTabWidget) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QTabWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTabWidget) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_InputMethodEvent
func miqt_exec_callback_QTabWidget_InputMethodEvent(self QTabWidget, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QTabWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTabWidget) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QTabWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTabWidget) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_InputMethodQuery
func miqt_exec_callback_QTabWidget_InputMethodQuery(self QTabWidget, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTabWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTabWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTabWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_FocusNextPrevChild
func miqt_exec_callback_QTabWidget_FocusNextPrevChild(self QTabWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
