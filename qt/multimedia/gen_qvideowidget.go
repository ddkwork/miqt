package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QVideoWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQVideoWidget constructs a new QVideoWidget object.
func NewQVideoWidget(parent *qt.QWidget) *QVideoWidget {
	g := newQVideoWidget(QVideoWidget_new((*QWidget)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQVideoWidget2 constructs a new QVideoWidget object.
func NewQVideoWidget2() *QVideoWidget {
	g := newQVideoWidget(QVideoWidget_new2())
	g.isSubclass = true
	return g
}

func (this *QVideoWidget) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QVideoWidget_MetaObject(this.h)))
}

func (this *QVideoWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QVideoWidget_Metacast(this.h, param1_Cstring))
}

func QVideoWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QVideoWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoWidget) VideoSink() *QVideoSink {
	return newQVideoSink(QVideoWidget_VideoSink(this.h))
}

func (this *QVideoWidget) AspectRatioMode() qt.AspectRatioMode {
	return (qt.AspectRatioMode)(QVideoWidget_AspectRatioMode(this.h))
}

func (this *QVideoWidget) SizeHint() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QVideoWidget_SizeHint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QVideoWidget) SetFullScreen(fullScreen bool) {
	QVideoWidget_SetFullScreen(this.h, (bool)(fullScreen))
}

func (this *QVideoWidget) SetAspectRatioMode(mode qt.AspectRatioMode) {
	QVideoWidget_SetAspectRatioMode(this.h, (int)(mode))
}

func (this *QVideoWidget) FullScreenChanged(fullScreen bool) {
	QVideoWidget_FullScreenChanged(this.h, (bool)(fullScreen))
}

func (this *QVideoWidget) OnFullScreenChanged(slot func(fullScreen bool)) {
	QVideoWidget_connect_FullScreenChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoWidget_FullScreenChanged
func miqt_exec_callback_QVideoWidget_FullScreenChanged(cb intptr_t, fullScreen bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(fullScreen bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(fullScreen)

	gofunc(slotval1)
}

func (this *QVideoWidget) AspectRatioModeChanged(mode qt.AspectRatioMode) {
	QVideoWidget_AspectRatioModeChanged(this.h, (int)(mode))
}

func (this *QVideoWidget) OnAspectRatioModeChanged(slot func(mode qt.AspectRatioMode)) {
	QVideoWidget_connect_AspectRatioModeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoWidget_AspectRatioModeChanged
func miqt_exec_callback_QVideoWidget_AspectRatioModeChanged(cb intptr_t, mode int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(mode qt.AspectRatioMode))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.AspectRatioMode)(mode)

	gofunc(slotval1)
}

func QVideoWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVideoWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QVideoWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QVideoWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QVideoWidget) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QVideoWidget_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QVideoWidget) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVideoWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoWidget_MetaObject
func miqt_exec_callback_QVideoWidget_MetaObject(self QVideoWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QVideoWidget{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QVideoWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QVideoWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QVideoWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QVideoWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QVideoWidget_Metacast
func miqt_exec_callback_QVideoWidget_Metacast(self QVideoWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QVideoWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
