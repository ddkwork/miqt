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

func (this *QTabWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTabWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTabWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_MetaObject
func miqt_exec_callback_QTabWidget_MetaObject(self QTabWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTabWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTabWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTabWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTabWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTabWidget_Metacast
func miqt_exec_callback_QTabWidget_Metacast(self QTabWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTabWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
