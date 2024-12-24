package qt

import (
	"unsafe"
)

type QOffscreenSurface struct {
	h          uintptr
	isSubclass bool
}

// NewQOffscreenSurface constructs a new QOffscreenSurface object.
func NewQOffscreenSurface() *QOffscreenSurface {

	ret := newQOffscreenSurface(QOffscreenSurface_new())
	ret.isSubclass = true
	return ret
}

// NewQOffscreenSurface2 constructs a new QOffscreenSurface object.
func NewQOffscreenSurface2(screen *QScreen) *QOffscreenSurface {

	ret := newQOffscreenSurface(QOffscreenSurface_new2(screen.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQOffscreenSurface3 constructs a new QOffscreenSurface object.
func NewQOffscreenSurface3(screen *QScreen, parent *QObject) *QOffscreenSurface {

	ret := newQOffscreenSurface(QOffscreenSurface_new3(screen.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QOffscreenSurface) MetaObject() *QMetaObject {
	return newQMetaObject(QOffscreenSurface_MetaObject(this.h))
}

func (this *QOffscreenSurface) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QOffscreenSurface_Metacast(this.h, param1_Cstring))
}

func QOffscreenSurface_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QOffscreenSurface_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QOffscreenSurface) SurfaceType() SurfaceType {
	xxxxxxxxx
}

func (this *QOffscreenSurface) Create() {
	QOffscreenSurface_Create(this.h)
}

func (this *QOffscreenSurface) Destroy() {
	QOffscreenSurface_Destroy(this.h)
}

func (this *QOffscreenSurface) IsValid() bool {
	return (bool)(QOffscreenSurface_IsValid(this.h))
}

func (this *QOffscreenSurface) SetFormat(format *QSurfaceFormat) {
	QOffscreenSurface_SetFormat(this.h, format.cPointer())
}

func (this *QOffscreenSurface) Format() *QSurfaceFormat {
	_goptr := newQSurfaceFormat(QOffscreenSurface_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOffscreenSurface) RequestedFormat() *QSurfaceFormat {
	_goptr := newQSurfaceFormat(QOffscreenSurface_RequestedFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOffscreenSurface) Size() *QSize {
	_goptr := newQSize(QOffscreenSurface_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QOffscreenSurface) Screen() *QScreen {
	return newQScreen(QOffscreenSurface_Screen(this.h))
}

func (this *QOffscreenSurface) SetScreen(screen *QScreen) {
	QOffscreenSurface_SetScreen(this.h, screen.cPointer())
}

func (this *QOffscreenSurface) ScreenChanged(screen *QScreen) {
	QOffscreenSurface_ScreenChanged(this.h, screen.cPointer())
}
func (this *QOffscreenSurface) OnScreenChanged(slot func(screen *QScreen)) {
	QOffscreenSurface_connect_ScreenChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_ScreenChanged
func miqt_exec_callback_QOffscreenSurface_ScreenChanged(cb intptr_t, screen *QScreen) {
	gofunc, ok := cgo.Handle(cb).Value().(func(screen *QScreen))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQScreen(screen)

	gofunc(slotval1)
}

func QOffscreenSurface_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QOffscreenSurface_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QOffscreenSurface_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QOffscreenSurface_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QOffscreenSurface) callVirtualBase_SurfaceType() SurfaceType {

	xxxxxxxxx
}
func (this *QOffscreenSurface) OnSurfaceType(slot func(super func() SurfaceType) SurfaceType) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_SurfaceType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_SurfaceType
func miqt_exec_callback_QOffscreenSurface_SurfaceType(self QOffscreenSurface, cb intptr_t) SurfaceType {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() SurfaceType) SurfaceType)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_SurfaceType)

	return virtualReturn

}

func (this *QOffscreenSurface) callVirtualBase_Format() *QSurfaceFormat {

	_goptr := newQSurfaceFormat(QOffscreenSurface_virtualbase_Format(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QOffscreenSurface) OnFormat(slot func(super func() *QSurfaceFormat) *QSurfaceFormat) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_Format(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Format
func miqt_exec_callback_QOffscreenSurface_Format(self QOffscreenSurface, cb intptr_t) *QSurfaceFormat {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSurfaceFormat) *QSurfaceFormat)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Format)

	return virtualReturn.cPointer()

}

func (this *QOffscreenSurface) callVirtualBase_Size() *QSize {

	_goptr := newQSize(QOffscreenSurface_virtualbase_Size(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QOffscreenSurface) OnSize(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_Size(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Size
func miqt_exec_callback_QOffscreenSurface_Size(self QOffscreenSurface, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Size)

	return virtualReturn.cPointer()

}

func (this *QOffscreenSurface) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QOffscreenSurface_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QOffscreenSurface) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Event
func miqt_exec_callback_QOffscreenSurface_Event(self QOffscreenSurface, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QOffscreenSurface) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QOffscreenSurface_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QOffscreenSurface) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_EventFilter
func miqt_exec_callback_QOffscreenSurface_EventFilter(self QOffscreenSurface, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QOffscreenSurface) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QOffscreenSurface_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QOffscreenSurface) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_TimerEvent
func miqt_exec_callback_QOffscreenSurface_TimerEvent(self QOffscreenSurface, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_ChildEvent(event *QChildEvent) {

	QOffscreenSurface_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QOffscreenSurface) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_ChildEvent
func miqt_exec_callback_QOffscreenSurface_ChildEvent(self QOffscreenSurface, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_CustomEvent(event *QEvent) {

	QOffscreenSurface_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QOffscreenSurface) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_CustomEvent
func miqt_exec_callback_QOffscreenSurface_CustomEvent(self QOffscreenSurface, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QOffscreenSurface_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QOffscreenSurface) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_ConnectNotify
func miqt_exec_callback_QOffscreenSurface_ConnectNotify(self QOffscreenSurface, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QOffscreenSurface) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QOffscreenSurface_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QOffscreenSurface) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_DisconnectNotify
func miqt_exec_callback_QOffscreenSurface_DisconnectNotify(self QOffscreenSurface, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QOffscreenSurface{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
