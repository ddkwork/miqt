package svg

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSvgRenderer struct {
	h          uintptr
	isSubclass bool
}

// NewQSvgRenderer constructs a new QSvgRenderer object.
func NewQSvgRenderer() *QSvgRenderer {

	ret := newQSvgRenderer(QSvgRenderer_new())
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer2 constructs a new QSvgRenderer object.
func NewQSvgRenderer2(filename string) *QSvgRenderer {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))

	ret := newQSvgRenderer(QSvgRenderer_new2(filename_ms))
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer3 constructs a new QSvgRenderer object.
func NewQSvgRenderer3(contents []byte) *QSvgRenderer {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))

	ret := newQSvgRenderer(QSvgRenderer_new3(contents_alias))
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer4 constructs a new QSvgRenderer object.
func NewQSvgRenderer4(contents *qt.QXmlStreamReader) *QSvgRenderer {

	ret := newQSvgRenderer(QSvgRenderer_new4((*QXmlStreamReader)(contents.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer5 constructs a new QSvgRenderer object.
func NewQSvgRenderer5(parent *qt.QObject) *QSvgRenderer {

	ret := newQSvgRenderer(QSvgRenderer_new5((*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer6 constructs a new QSvgRenderer object.
func NewQSvgRenderer6(filename string, parent *qt.QObject) *QSvgRenderer {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))

	ret := newQSvgRenderer(QSvgRenderer_new6(filename_ms, (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer7 constructs a new QSvgRenderer object.
func NewQSvgRenderer7(contents []byte, parent *qt.QObject) *QSvgRenderer {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))

	ret := newQSvgRenderer(QSvgRenderer_new7(contents_alias, (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSvgRenderer8 constructs a new QSvgRenderer object.
func NewQSvgRenderer8(contents *qt.QXmlStreamReader, parent *qt.QObject) *QSvgRenderer {

	ret := newQSvgRenderer(QSvgRenderer_new8((*QXmlStreamReader)(contents.UnsafePointer()), (*QObject)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QSvgRenderer) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSvgRenderer_MetaObject(this.h)))
}

func (this *QSvgRenderer) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSvgRenderer_Metacast(this.h, param1_Cstring))
}

func QSvgRenderer_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSvgRenderer_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgRenderer) IsValid() bool {
	return (bool)(QSvgRenderer_IsValid(this.h))
}

func (this *QSvgRenderer) DefaultSize() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgRenderer_DefaultSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgRenderer) ViewBox() *qt.QRect {
	_goptr := qt.UnsafeNewQRect(unsafe.Pointer(QSvgRenderer_ViewBox(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgRenderer) ViewBoxF() *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QSvgRenderer_ViewBoxF(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgRenderer) SetViewBox(viewbox *qt.QRect) {
	QSvgRenderer_SetViewBox(this.h, (*QRect)(viewbox.UnsafePointer()))
}

func (this *QSvgRenderer) SetViewBoxWithViewbox(viewbox *qt.QRectF) {
	QSvgRenderer_SetViewBoxWithViewbox(this.h, (*QRectF)(viewbox.UnsafePointer()))
}

func (this *QSvgRenderer) AspectRatioMode() qt.AspectRatioMode {
	return (qt.AspectRatioMode)(QSvgRenderer_AspectRatioMode(this.h))
}

func (this *QSvgRenderer) SetAspectRatioMode(mode qt.AspectRatioMode) {
	QSvgRenderer_SetAspectRatioMode(this.h, (int)(mode))
}

func (this *QSvgRenderer) Options() Option {
	return (Option)(QSvgRenderer_Options(this.h))
}

func (this *QSvgRenderer) SetOptions(flags Option) {
	QSvgRenderer_SetOptions(this.h, (int)(flags))
}

func (this *QSvgRenderer) Animated() bool {
	return (bool)(QSvgRenderer_Animated(this.h))
}

func (this *QSvgRenderer) FramesPerSecond() int {
	return (int)(QSvgRenderer_FramesPerSecond(this.h))
}

func (this *QSvgRenderer) SetFramesPerSecond(num int) {
	QSvgRenderer_SetFramesPerSecond(this.h, (int)(num))
}

func (this *QSvgRenderer) CurrentFrame() int {
	return (int)(QSvgRenderer_CurrentFrame(this.h))
}

func (this *QSvgRenderer) SetCurrentFrame(currentFrame int) {
	QSvgRenderer_SetCurrentFrame(this.h, (int)(currentFrame))
}

func (this *QSvgRenderer) AnimationDuration() int {
	return (int)(QSvgRenderer_AnimationDuration(this.h))
}

func (this *QSvgRenderer) IsAnimationEnabled() bool {
	return (bool)(QSvgRenderer_IsAnimationEnabled(this.h))
}

func (this *QSvgRenderer) SetAnimationEnabled(enable bool) {
	QSvgRenderer_SetAnimationEnabled(this.h, (bool)(enable))
}

func (this *QSvgRenderer) BoundsOnElement(id string) *qt.QRectF {
	id_ms := struct_miqt_string{}
	id_ms.data = CString(id)
	id_ms.len = size_t(len(id))
	defer free(unsafe.Pointer(id_ms.data))
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QSvgRenderer_BoundsOnElement(this.h, id_ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgRenderer) ElementExists(id string) bool {
	id_ms := struct_miqt_string{}
	id_ms.data = CString(id)
	id_ms.len = size_t(len(id))
	defer free(unsafe.Pointer(id_ms.data))
	return (bool)(QSvgRenderer_ElementExists(this.h, id_ms))
}

func (this *QSvgRenderer) TransformForElement(id string) *qt.QTransform {
	id_ms := struct_miqt_string{}
	id_ms.data = CString(id)
	id_ms.len = size_t(len(id))
	defer free(unsafe.Pointer(id_ms.data))
	_goptr := qt.UnsafeNewQTransform(unsafe.Pointer(QSvgRenderer_TransformForElement(this.h, id_ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QSvgRenderer_SetDefaultOptions(flags Option) {
	QSvgRenderer_SetDefaultOptions((int)(flags))
}

func (this *QSvgRenderer) Load(filename string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QSvgRenderer_Load(this.h, filename_ms))
}

func (this *QSvgRenderer) LoadWithContents(contents []byte) bool {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))
	return (bool)(QSvgRenderer_LoadWithContents(this.h, contents_alias))
}

func (this *QSvgRenderer) Load2(contents *qt.QXmlStreamReader) bool {
	return (bool)(QSvgRenderer_Load2(this.h, (*QXmlStreamReader)(contents.UnsafePointer())))
}

func (this *QSvgRenderer) Render(p *qt.QPainter) {
	QSvgRenderer_Render(this.h, (*QPainter)(p.UnsafePointer()))
}

func (this *QSvgRenderer) Render2(p *qt.QPainter, bounds *qt.QRectF) {
	QSvgRenderer_Render2(this.h, (*QPainter)(p.UnsafePointer()), (*QRectF)(bounds.UnsafePointer()))
}

func (this *QSvgRenderer) Render3(p *qt.QPainter, elementId string) {
	elementId_ms := struct_miqt_string{}
	elementId_ms.data = CString(elementId)
	elementId_ms.len = size_t(len(elementId))
	defer free(unsafe.Pointer(elementId_ms.data))
	QSvgRenderer_Render3(this.h, (*QPainter)(p.UnsafePointer()), elementId_ms)
}

func (this *QSvgRenderer) RepaintNeeded() {
	QSvgRenderer_RepaintNeeded(this.h)
}
func (this *QSvgRenderer) OnRepaintNeeded(slot func()) {
	QSvgRenderer_connect_RepaintNeeded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_RepaintNeeded
func miqt_exec_callback_QSvgRenderer_RepaintNeeded(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QSvgRenderer_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSvgRenderer_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSvgRenderer_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSvgRenderer_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgRenderer) Render32(p *qt.QPainter, elementId string, bounds *qt.QRectF) {
	elementId_ms := struct_miqt_string{}
	elementId_ms.data = CString(elementId)
	elementId_ms.len = size_t(len(elementId))
	defer free(unsafe.Pointer(elementId_ms.data))
	QSvgRenderer_Render32(this.h, (*QPainter)(p.UnsafePointer()), elementId_ms, (*QRectF)(bounds.UnsafePointer()))
}

func (this *QSvgRenderer) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QSvgRenderer_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QSvgRenderer) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_Event
func miqt_exec_callback_QSvgRenderer_Event(self QSvgRenderer, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QSvgRenderer{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSvgRenderer) callVirtualBase_EventFilter(watched *qt.QObject, event *qt.QEvent) bool {

	return (bool)(QSvgRenderer_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(watched.UnsafePointer()), (*QEvent)(event.UnsafePointer())))

}
func (this *QSvgRenderer) OnEventFilter(slot func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_EventFilter
func miqt_exec_callback_QSvgRenderer_EventFilter(self QSvgRenderer, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *qt.QObject, event *qt.QEvent) bool, watched *qt.QObject, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(watched))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QSvgRenderer{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSvgRenderer) callVirtualBase_TimerEvent(event *qt.QTimerEvent) {

	QSvgRenderer_virtualbase_TimerEvent(unsafe.Pointer(this.h), (*QTimerEvent)(event.UnsafePointer()))

}
func (this *QSvgRenderer) OnTimerEvent(slot func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_TimerEvent
func miqt_exec_callback_QSvgRenderer_TimerEvent(self QSvgRenderer, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTimerEvent), event *qt.QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTimerEvent(unsafe.Pointer(event))

	gofunc((&QSvgRenderer{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSvgRenderer) callVirtualBase_ChildEvent(event *qt.QChildEvent) {

	QSvgRenderer_virtualbase_ChildEvent(unsafe.Pointer(this.h), (*QChildEvent)(event.UnsafePointer()))

}
func (this *QSvgRenderer) OnChildEvent(slot func(super func(event *qt.QChildEvent), event *qt.QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_ChildEvent
func miqt_exec_callback_QSvgRenderer_ChildEvent(self QSvgRenderer, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QChildEvent), event *qt.QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQChildEvent(unsafe.Pointer(event))

	gofunc((&QSvgRenderer{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSvgRenderer) callVirtualBase_CustomEvent(event *qt.QEvent) {

	QSvgRenderer_virtualbase_CustomEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QSvgRenderer) OnCustomEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_CustomEvent
func miqt_exec_callback_QSvgRenderer_CustomEvent(self QSvgRenderer, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QSvgRenderer{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSvgRenderer) callVirtualBase_ConnectNotify(signal *qt.QMetaMethod) {

	QSvgRenderer_virtualbase_ConnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QSvgRenderer) OnConnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_ConnectNotify
func miqt_exec_callback_QSvgRenderer_ConnectNotify(self QSvgRenderer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QSvgRenderer{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSvgRenderer) callVirtualBase_DisconnectNotify(signal *qt.QMetaMethod) {

	QSvgRenderer_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), (*QMetaMethod)(signal.UnsafePointer()))

}
func (this *QSvgRenderer) OnDisconnectNotify(slot func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_DisconnectNotify
func miqt_exec_callback_QSvgRenderer_DisconnectNotify(self QSvgRenderer, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *qt.QMetaMethod), signal *qt.QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMetaMethod(unsafe.Pointer(signal))

	gofunc((&QSvgRenderer{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
