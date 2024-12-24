package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QScreenCapture__Error int

const (
	QScreenCapture__NoError               QScreenCapture__Error = 0
	QScreenCapture__InternalError         QScreenCapture__Error = 1
	QScreenCapture__CapturingNotSupported QScreenCapture__Error = 2
	QScreenCapture__CaptureFailed         QScreenCapture__Error = 4
	QScreenCapture__NotFound              QScreenCapture__Error = 5
)

type QScreenCapture struct {
	h          uintptr
	isSubclass bool
}

// NewQScreenCapture constructs a new QScreenCapture object.
func NewQScreenCapture() *QScreenCapture {
	g := newQScreenCapture(QScreenCapture_new())
	g.isSubclass = true
	return g
}

// NewQScreenCapture2 constructs a new QScreenCapture object.
func NewQScreenCapture2(parent *qt.QObject) *QScreenCapture {
	g := newQScreenCapture(QScreenCapture_new2((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

func (this *QScreenCapture) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QScreenCapture_MetaObject(this.h)))
}

func (this *QScreenCapture) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QScreenCapture_Metacast(this.h, param1_Cstring))
}

func QScreenCapture_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QScreenCapture_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreenCapture) CaptureSession() *QMediaCaptureSession {
	return newQMediaCaptureSession(QScreenCapture_CaptureSession(this.h))
}

func (this *QScreenCapture) SetScreen(screen *qt.QScreen) {
	QScreenCapture_SetScreen(this.h, (*QScreen)(screen.UnsafePointer()))
}

func (this *QScreenCapture) Screen() *qt.QScreen {
	return qt.UnsafeNewQScreen(unsafe.Pointer(QScreenCapture_Screen(this.h)))
}

func (this *QScreenCapture) IsActive() bool {
	return (bool)(QScreenCapture_IsActive(this.h))
}

func (this *QScreenCapture) Error() Error {
	xxxxxxxxx
}

func (this *QScreenCapture) ErrorString() string {
	var _ms struct_miqt_string = QScreenCapture_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreenCapture) SetActive(active bool) {
	QScreenCapture_SetActive(this.h, (bool)(active))
}

func (this *QScreenCapture) Start() {
	QScreenCapture_Start(this.h)
}

func (this *QScreenCapture) Stop() {
	QScreenCapture_Stop(this.h)
}

func (this *QScreenCapture) ActiveChanged(param1 bool) {
	QScreenCapture_ActiveChanged(this.h, (bool)(param1))
}

func (this *QScreenCapture) OnActiveChanged(slot func(param1 bool)) {
	QScreenCapture_connect_ActiveChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenCapture_ActiveChanged
func miqt_exec_callback_QScreenCapture_ActiveChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QScreenCapture) ErrorChanged() {
	QScreenCapture_ErrorChanged(this.h)
}

func (this *QScreenCapture) OnErrorChanged(slot func()) {
	QScreenCapture_connect_ErrorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenCapture_ErrorChanged
func miqt_exec_callback_QScreenCapture_ErrorChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QScreenCapture) ScreenChanged(param1 *qt.QScreen) {
	QScreenCapture_ScreenChanged(this.h, (*QScreen)(param1.UnsafePointer()))
}

func (this *QScreenCapture) OnScreenChanged(slot func(param1 *qt.QScreen)) {
	QScreenCapture_connect_ScreenChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenCapture_ScreenChanged
func miqt_exec_callback_QScreenCapture_ScreenChanged(cb intptr_t, param1 *QScreen) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *qt.QScreen))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQScreen(unsafe.Pointer(param1))

	gofunc(slotval1)
}

func (this *QScreenCapture) ErrorOccurred(error QScreenCapture__Error, errorString string) {
	errorString_ms := struct_miqt_string{}
	errorString_ms.data = CString(errorString)
	errorString_ms.len = size_t(len(errorString))
	defer free(unsafe.Pointer(errorString_ms.data))
	QScreenCapture_ErrorOccurred(this.h, (int)(error), errorString_ms)
}

func (this *QScreenCapture) OnErrorOccurred(slot func(error QScreenCapture__Error, errorString string)) {
	QScreenCapture_connect_ErrorOccurred(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenCapture_ErrorOccurred
func miqt_exec_callback_QScreenCapture_ErrorOccurred(cb intptr_t, error int, errorString struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(error QScreenCapture__Error, errorString string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QScreenCapture__Error)(error)

	var errorString_ms struct_miqt_string = errorString
	errorString_ret := GoStringN(errorString_ms.data, int(int64(errorString_ms.len)))
	free(unsafe.Pointer(errorString_ms.data))
	slotval2 := errorString_ret

	gofunc(slotval1, slotval2)
}

func QScreenCapture_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScreenCapture_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QScreenCapture_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QScreenCapture_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QScreenCapture) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QScreenCapture_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QScreenCapture) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScreenCapture_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenCapture_MetaObject
func miqt_exec_callback_QScreenCapture_MetaObject(self QScreenCapture, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QScreenCapture{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QScreenCapture) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QScreenCapture_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QScreenCapture) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QScreenCapture_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QScreenCapture_Metacast
func miqt_exec_callback_QScreenCapture_Metacast(self QScreenCapture, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QScreenCapture{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
