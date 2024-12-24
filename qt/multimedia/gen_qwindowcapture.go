package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QWindowCapture__Error int

const (
	QWindowCapture__NoError               QWindowCapture__Error = 0
	QWindowCapture__InternalError         QWindowCapture__Error = 1
	QWindowCapture__CapturingNotSupported QWindowCapture__Error = 2
	QWindowCapture__CaptureFailed         QWindowCapture__Error = 4
	QWindowCapture__NotFound              QWindowCapture__Error = 5
)

type QWindowCapture struct {
	h          uintptr
	isSubclass bool
}

// NewQWindowCapture constructs a new QWindowCapture object.
func NewQWindowCapture() *QWindowCapture {
	g := newQWindowCapture(QWindowCapture_new())
	g.isSubclass = true
	return g
}

// NewQWindowCapture2 constructs a new QWindowCapture object.
func NewQWindowCapture2(parent *qt.QObject) *QWindowCapture {
	g := newQWindowCapture(QWindowCapture_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QWindowCapture) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWindowCapture_MetaObject(this.h)))
}

func (this *QWindowCapture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWindowCapture_Metacast(this.h, param1_Cstring))
}

func QWindowCapture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWindowCapture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWindowCapture_CapturableWindows() []QCapturableWindow {
	var _ma struct_miqt_array = QWindowCapture_CapturableWindows()
	_ret := make([]QCapturableWindow, int(_ma.len))
	_outCast := (*[0xffff]*QCapturableWindow)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQCapturableWindow(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QWindowCapture) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QWindowCapture_CaptureSession(this.h))
}

func (this *QWindowCapture) SetWindow(window QCapturableWindow) {
	QWindowCapture_SetWindow(this.h, window.cPointer())
}

func (this *QWindowCapture) Window() *QCapturableWindow {
	_goptr := newQCapturableWindow(QWindowCapture_Window(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindowCapture) IsActive() bool {
	return (bool)(QWindowCapture_IsActive(this.h))
}

func (this *QWindowCapture) Error() Error {
	xxxxxxxxx
}

func (this *QWindowCapture) ErrorString() string {
	var _ms struct_miqt_string = QWindowCapture_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWindowCapture) SetActive(active bool) {
	QWindowCapture_SetActive(this.h, (bool)(active))
}

func (this *QWindowCapture) Start() {
	QWindowCapture_Start(this.h)
}

func (this *QWindowCapture) Stop() {
	QWindowCapture_Stop(this.h)
}

func (this *QWindowCapture) ActiveChanged(param1 bool) {
	QWindowCapture_ActiveChanged(this.h, (bool)(param1))
}

func (this *QWindowCapture) OnActiveChanged(slot func(param1 bool)) {
	QWindowCapture_connect_ActiveChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowCapture_ActiveChanged
func miqt_exec_callback_QWindowCapture_ActiveChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QWindowCapture) WindowChanged(window QCapturableWindow) {
	QWindowCapture_WindowChanged(this.h, window.cPointer())
}

func (this *QWindowCapture) OnWindowChanged(slot func(window QCapturableWindow)) {
	QWindowCapture_connect_WindowChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowCapture_WindowChanged
func miqt_exec_callback_QWindowCapture_WindowChanged(cb intptr_t, window *QCapturableWindow) {
	gofunc, ok := cgo.Handle(cb).Value().(func(window QCapturableWindow))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	window_goptr := newQCapturableWindow(window)
	window_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *window_goptr

	gofunc(slotval1)
}

func (this *QWindowCapture) ErrorChanged() {
	QWindowCapture_ErrorChanged(this.h)
}

func (this *QWindowCapture) OnErrorChanged(slot func()) {
	QWindowCapture_connect_ErrorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowCapture_ErrorChanged
func miqt_exec_callback_QWindowCapture_ErrorChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWindowCapture) ErrorOccurred(error QWindowCapture__Error, errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QWindowCapture_ErrorOccurred(this.h, (int)(error), errorString_ms)
}

func (this *QWindowCapture) OnErrorOccurred(slot func(error QWindowCapture__Error, errorString string)) {
	QWindowCapture_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowCapture_ErrorOccurred
func miqt_exec_callback_QWindowCapture_ErrorOccurred(cb intptr_t, error int, errorString struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QWindowCapture__Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QWindowCapture__Error)(error)

	var errorString_ms struct_miqt_string = errorString
	errorString_ret := GoStringN(errorString_ms.data, int(int64(errorString_ms.len)))
	free(unsafe.Pointer(errorString_ms.data))
	slotval2 := errorString_ret

	gofunc(slotval1, slotval2)
}

func QWindowCapture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWindowCapture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWindowCapture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWindowCapture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWindowCapture) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QWindowCapture_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QWindowCapture) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowCapture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowCapture_MetaObject
func miqt_exec_callback_QWindowCapture_MetaObject(self QWindowCapture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWindowCapture{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QWindowCapture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWindowCapture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWindowCapture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowCapture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowCapture_Metacast
func miqt_exec_callback_QWindowCapture_Metacast(self QWindowCapture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QWindowCapture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
