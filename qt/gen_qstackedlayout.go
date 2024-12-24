package qt

import (
	"unsafe"
)

type QStackedLayout__StackingMode int

const (
	QStackedLayout__StackOne QStackedLayout__StackingMode = 0
	QStackedLayout__StackAll QStackedLayout__StackingMode = 1
)

type QStackedLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQStackedLayout constructs a new QStackedLayout object.
func NewQStackedLayout(parent *QWidget) *QStackedLayout {
	ret := newQStackedLayout(QStackedLayout_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStackedLayout2 constructs a new QStackedLayout object.
func NewQStackedLayout2() *QStackedLayout {
	ret := newQStackedLayout(QStackedLayout_new2())
	ret.isSubclass = true
	return ret
}

// NewQStackedLayout3 constructs a new QStackedLayout object.
func NewQStackedLayout3(parentLayout *QLayout) *QStackedLayout {
	ret := newQStackedLayout(QStackedLayout_new3(parentLayout.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QStackedLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QStackedLayout_MetaObject(this.h))
}

func (this *QStackedLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStackedLayout_Metacast(this.h, param1_Cstring))
}

func QStackedLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStackedLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStackedLayout) AddWidget(w *QWidget) int {
	return (int)(QStackedLayout_AddWidget(this.h, w.cPointer()))
}

func (this *QStackedLayout) InsertWidget(index int, w *QWidget) int {
	return (int)(QStackedLayout_InsertWidget(this.h, (int)(index), w.cPointer()))
}

func (this *QStackedLayout) CurrentWidget() *QWidget {
	return newQWidget(QStackedLayout_CurrentWidget(this.h))
}

func (this *QStackedLayout) CurrentIndex() int {
	return (int)(QStackedLayout_CurrentIndex(this.h))
}

func (this *QStackedLayout) Widget(param1 int) *QWidget {
	return newQWidget(QStackedLayout_Widget(this.h, (int)(param1)))
}

func (this *QStackedLayout) Count() int {
	return (int)(QStackedLayout_Count(this.h))
}

func (this *QStackedLayout) StackingMode() StackingMode {
	xxxxxxxxx
}

func (this *QStackedLayout) SetStackingMode(stackingMode StackingMode) {
	QStackedLayout_SetStackingMode(this.h, stackingMode)
}

func (this *QStackedLayout) AddItem(item *QLayoutItem) {
	QStackedLayout_AddItem(this.h, item.cPointer())
}

func (this *QStackedLayout) SizeHint() *QSize {
	_goptr := newQSize(QStackedLayout_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStackedLayout) MinimumSize() *QSize {
	_goptr := newQSize(QStackedLayout_MinimumSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStackedLayout) ItemAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QStackedLayout_ItemAt(this.h, (int)(param1)))
}

func (this *QStackedLayout) TakeAt(param1 int) *QLayoutItem {
	return newQLayoutItem(QStackedLayout_TakeAt(this.h, (int)(param1)))
}

func (this *QStackedLayout) SetGeometry(rect *QRect) {
	QStackedLayout_SetGeometry(this.h, rect.cPointer())
}

func (this *QStackedLayout) HasHeightForWidth() bool {
	return (bool)(QStackedLayout_HasHeightForWidth(this.h))
}

func (this *QStackedLayout) HeightForWidth(width int) int {
	return (int)(QStackedLayout_HeightForWidth(this.h, (int)(width)))
}

func (this *QStackedLayout) WidgetRemoved(index int) {
	QStackedLayout_WidgetRemoved(this.h, (int)(index))
}

func (this *QStackedLayout) OnWidgetRemoved(slot func(index int)) {
	QStackedLayout_connect_WidgetRemoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_WidgetRemoved
func miqt_exec_callback_QStackedLayout_WidgetRemoved(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedLayout) CurrentChanged(index int) {
	QStackedLayout_CurrentChanged(this.h, (int)(index))
}

func (this *QStackedLayout) OnCurrentChanged(slot func(index int)) {
	QStackedLayout_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_CurrentChanged
func miqt_exec_callback_QStackedLayout_CurrentChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedLayout) WidgetAdded(index int) {
	QStackedLayout_WidgetAdded(this.h, (int)(index))
}

func (this *QStackedLayout) OnWidgetAdded(slot func(index int)) {
	QStackedLayout_connect_WidgetAdded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_WidgetAdded
func miqt_exec_callback_QStackedLayout_WidgetAdded(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedLayout) SetCurrentIndex(index int) {
	QStackedLayout_SetCurrentIndex(this.h, (int)(index))
}

func (this *QStackedLayout) SetCurrentWidget(w *QWidget) {
	QStackedLayout_SetCurrentWidget(this.h, w.cPointer())
}

func QStackedLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStackedLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStackedLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStackedLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStackedLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStackedLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStackedLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_MetaObject
func miqt_exec_callback_QStackedLayout_MetaObject(self QStackedLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStackedLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStackedLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStackedLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedLayout_Metacast
func miqt_exec_callback_QStackedLayout_Metacast(self QStackedLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QStackedLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
