package qt

import (
	"unsafe"
)

type QSplashScreen struct {
	h          uintptr
	isSubclass bool
}

// NewQSplashScreen constructs a new QSplashScreen object.
func NewQSplashScreen() *QSplashScreen {
	g := newQSplashScreen(QSplashScreen_new())
	g.isSubclass = true
	return g
}

// NewQSplashScreen2 constructs a new QSplashScreen object.
func NewQSplashScreen2(screen *QScreen) *QSplashScreen {
	g := newQSplashScreen(QSplashScreen_new2(screen.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSplashScreen3 constructs a new QSplashScreen object.
func NewQSplashScreen3(pixmap *QPixmap) *QSplashScreen {
	g := newQSplashScreen(QSplashScreen_new3(pixmap.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSplashScreen4 constructs a new QSplashScreen object.
func NewQSplashScreen4(pixmap *QPixmap, f WindowType) *QSplashScreen {
	g := newQSplashScreen(QSplashScreen_new4(pixmap.cPointer(), (int)(f)))
	g.isSubclass = true
	return g
}

// NewQSplashScreen5 constructs a new QSplashScreen object.
func NewQSplashScreen5(screen *QScreen, pixmap *QPixmap) *QSplashScreen {
	g := newQSplashScreen(QSplashScreen_new5(screen.cPointer(), pixmap.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSplashScreen6 constructs a new QSplashScreen object.
func NewQSplashScreen6(screen *QScreen, pixmap *QPixmap, f WindowType) *QSplashScreen {
	g := newQSplashScreen(QSplashScreen_new6(screen.cPointer(), pixmap.cPointer(), (int)(f)))
	g.isSubclass = true
	return g
}

func (this *QSplashScreen) MetaObject() *QMetaObject {
	return newQMetaObject(QSplashScreen_MetaObject(this.h))
}

func (this *QSplashScreen) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSplashScreen_Metacast(this.h, param1_Cstring))
}

func QSplashScreen_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSplashScreen_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplashScreen) SetPixmap(pixmap *QPixmap) {
	QSplashScreen_SetPixmap(this.h, pixmap.cPointer())
}

func (this *QSplashScreen) Pixmap() *QPixmap {
	_goptr := newQPixmap(QSplashScreen_Pixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSplashScreen) Finish(w *QWidget) {
	QSplashScreen_Finish(this.h, w.cPointer())
}

func (this *QSplashScreen) Repaint() {
	QSplashScreen_Repaint(this.h)
}

func (this *QSplashScreen) Message() string {
	var _ms struct_miqt_string = QSplashScreen_Message(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplashScreen) ShowMessage(message string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QSplashScreen_ShowMessage(this.h, message_ms)
}

func (this *QSplashScreen) ClearMessage() {
	QSplashScreen_ClearMessage(this.h)
}

func (this *QSplashScreen) MessageChanged(message string) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QSplashScreen_MessageChanged(this.h, message_ms)
}

func (this *QSplashScreen) OnMessageChanged(slot func(message string)) {
	QSplashScreen_connect_MessageChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplashScreen_MessageChanged
func miqt_exec_callback_QSplashScreen_MessageChanged(cb intptr_t, message struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(message string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var message_ms struct_miqt_string = message
	message_ret := GoStringN(message_ms.data, int(int64(message_ms.len)))
	free(unsafe.Pointer(message_ms.data))
	slotval1 := message_ret

	gofunc(slotval1)
}

func QSplashScreen_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplashScreen_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSplashScreen_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSplashScreen_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSplashScreen) ShowMessage2(message string, alignment int) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QSplashScreen_ShowMessage2(this.h, message_ms, (int)(alignment))
}

func (this *QSplashScreen) ShowMessage3(message string, alignment int, color *QColor) {
	message_ms := struct_miqt_string{}
	message_ms.data = CString(message)
	message_ms.len = size_t(len(message))
	defer free(unsafe.Pointer(message_ms.data))
	QSplashScreen_ShowMessage3(this.h, message_ms, (int)(alignment), color.cPointer())
}

func (this *QSplashScreen) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSplashScreen_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSplashScreen) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplashScreen_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplashScreen_MetaObject
func miqt_exec_callback_QSplashScreen_MetaObject(self QSplashScreen, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSplashScreen{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSplashScreen) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSplashScreen_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSplashScreen) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSplashScreen_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSplashScreen_Metacast
func miqt_exec_callback_QSplashScreen_Metacast(self QSplashScreen, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSplashScreen{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
