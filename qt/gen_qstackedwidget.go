package qt

import (
	"unsafe"
)

type QStackedWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQStackedWidget constructs a new QStackedWidget object.
func NewQStackedWidget(parent *QWidget) *QStackedWidget {
	g := newQStackedWidget(QStackedWidget_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQStackedWidget2 constructs a new QStackedWidget object.
func NewQStackedWidget2() *QStackedWidget {
	g := newQStackedWidget(QStackedWidget_new2())
	g.isSubclass = true
	return g
}

func (this *QStackedWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QStackedWidget_MetaObject(this.h))
}

func (this *QStackedWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStackedWidget_Metacast(this.h, param1_Cstring))
}

func QStackedWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStackedWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStackedWidget) AddWidget(w *QWidget) int {
	return (int)(QStackedWidget_AddWidget(this.h, w.cPointer()))
}

func (this *QStackedWidget) InsertWidget(index int, w *QWidget) int {
	return (int)(QStackedWidget_InsertWidget(this.h, (int)(index), w.cPointer()))
}

func (this *QStackedWidget) RemoveWidget(w *QWidget) {
	QStackedWidget_RemoveWidget(this.h, w.cPointer())
}

func (this *QStackedWidget) CurrentWidget() *QWidget {
	return newQWidget(QStackedWidget_CurrentWidget(this.h))
}

func (this *QStackedWidget) CurrentIndex() int {
	return (int)(QStackedWidget_CurrentIndex(this.h))
}

func (this *QStackedWidget) IndexOf(param1 *QWidget) int {
	return (int)(QStackedWidget_IndexOf(this.h, param1.cPointer()))
}

func (this *QStackedWidget) Widget(param1 int) *QWidget {
	return newQWidget(QStackedWidget_Widget(this.h, (int)(param1)))
}

func (this *QStackedWidget) Count() int {
	return (int)(QStackedWidget_Count(this.h))
}

func (this *QStackedWidget) SetCurrentIndex(index int) {
	QStackedWidget_SetCurrentIndex(this.h, (int)(index))
}

func (this *QStackedWidget) SetCurrentWidget(w *QWidget) {
	QStackedWidget_SetCurrentWidget(this.h, w.cPointer())
}

func (this *QStackedWidget) CurrentChanged(param1 int) {
	QStackedWidget_CurrentChanged(this.h, (int)(param1))
}

func (this *QStackedWidget) OnCurrentChanged(slot func(param1 int)) {
	QStackedWidget_connect_CurrentChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedWidget_CurrentChanged
func miqt_exec_callback_QStackedWidget_CurrentChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QStackedWidget) WidgetRemoved(index int) {
	QStackedWidget_WidgetRemoved(this.h, (int)(index))
}

func (this *QStackedWidget) OnWidgetRemoved(slot func(index int)) {
	QStackedWidget_connect_WidgetRemoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedWidget_WidgetRemoved
func miqt_exec_callback_QStackedWidget_WidgetRemoved(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QStackedWidget) WidgetAdded(index int) {
	QStackedWidget_WidgetAdded(this.h, (int)(index))
}

func (this *QStackedWidget) OnWidgetAdded(slot func(index int)) {
	QStackedWidget_connect_WidgetAdded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedWidget_WidgetAdded
func miqt_exec_callback_QStackedWidget_WidgetAdded(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func QStackedWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStackedWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStackedWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStackedWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStackedWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStackedWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStackedWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedWidget_MetaObject
func miqt_exec_callback_QStackedWidget_MetaObject(self QStackedWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStackedWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStackedWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStackedWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStackedWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStackedWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStackedWidget_Metacast
func miqt_exec_callback_QStackedWidget_Metacast(self QStackedWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QStackedWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
