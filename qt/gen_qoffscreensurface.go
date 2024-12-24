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
	g := newQOffscreenSurface(QOffscreenSurface_new())
	g.isSubclass = true
	return g
}

// NewQOffscreenSurface2 constructs a new QOffscreenSurface object.
func NewQOffscreenSurface2(screen *QScreen) *QOffscreenSurface {
	g := newQOffscreenSurface(QOffscreenSurface_new2(screen.cPointer()))
	g.isSubclass = true
	return g
}

// NewQOffscreenSurface3 constructs a new QOffscreenSurface object.
func NewQOffscreenSurface3(screen *QScreen, parent *QObject) *QOffscreenSurface {
	g := newQOffscreenSurface(QOffscreenSurface_new3(screen.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QOffscreenSurface) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QOffscreenSurface_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QOffscreenSurface) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_MetaObject
func miqt_exec_callback_QOffscreenSurface_MetaObject(self QOffscreenSurface, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QOffscreenSurface) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QOffscreenSurface_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QOffscreenSurface) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QOffscreenSurface_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QOffscreenSurface_Metacast
func miqt_exec_callback_QOffscreenSurface_Metacast(self QOffscreenSurface, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QOffscreenSurface{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
