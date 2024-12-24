package qt

import (
	"unsafe"
)

type QSlider__TickPosition int

const (
	QSlider__NoTicks        QSlider__TickPosition = 0
	QSlider__TicksAbove     QSlider__TickPosition = 1
	QSlider__TicksLeft      QSlider__TickPosition = 1
	QSlider__TicksBelow     QSlider__TickPosition = 2
	QSlider__TicksRight     QSlider__TickPosition = 2
	QSlider__TicksBothSides QSlider__TickPosition = 3
)

type QSlider struct {
	h          uintptr
	isSubclass bool
}

// NewQSlider constructs a new QSlider object.
func NewQSlider(parent *QWidget) *QSlider {
	g := newQSlider(QSlider_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSlider2 constructs a new QSlider object.
func NewQSlider2() *QSlider {
	g := newQSlider(QSlider_new2())
	g.isSubclass = true
	return g
}

// NewQSlider3 constructs a new QSlider object.
func NewQSlider3(orientation Orientation) *QSlider {
	g := newQSlider(QSlider_new3((int)(orientation)))
	g.isSubclass = true
	return g
}

// NewQSlider4 constructs a new QSlider object.
func NewQSlider4(orientation Orientation, parent *QWidget) *QSlider {
	g := newQSlider(QSlider_new4((int)(orientation), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSlider) MetaObject() *QMetaObject {
	return newQMetaObject(QSlider_MetaObject(this.h))
}

func (this *QSlider) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSlider_Metacast(this.h, param1_Cstring))
}

func QSlider_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSlider_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSlider) SizeHint() *QSize {
	_goptr := newQSize(QSlider_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSlider) MinimumSizeHint() *QSize {
	_goptr := newQSize(QSlider_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSlider) SetTickPosition(position TickPosition) {
	QSlider_SetTickPosition(this.h, position)
}

func (this *QSlider) TickPosition() TickPosition {
	xxxxxxxxx
}

func (this *QSlider) SetTickInterval(ti int) {
	QSlider_SetTickInterval(this.h, (int)(ti))
}

func (this *QSlider) TickInterval() int {
	return (int)(QSlider_TickInterval(this.h))
}

func (this *QSlider) Event(event *QEvent) bool {
	return (bool)(QSlider_Event(this.h, event.cPointer()))
}

func QSlider_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSlider_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSlider_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSlider_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSlider) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSlider_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSlider) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_MetaObject
func miqt_exec_callback_QSlider_MetaObject(self QSlider, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSlider{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSlider) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSlider_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSlider) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSlider_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSlider_Metacast
func miqt_exec_callback_QSlider_Metacast(self QSlider, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSlider{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
