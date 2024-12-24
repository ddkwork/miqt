package svg

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QSvgRenderer struct {
	h          uintptr
	isSubclass bool
}

// NewQSvgRenderer constructs a new QSvgRenderer object.
func NewQSvgRenderer() *QSvgRenderer {
	g := newQSvgRenderer(QSvgRenderer_new())
	g.isSubclass = true
	return g
}

// NewQSvgRenderer2 constructs a new QSvgRenderer object.
func NewQSvgRenderer2(filename string) *QSvgRenderer {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	g := newQSvgRenderer(QSvgRenderer_new2(filename_ms))
	g.isSubclass = true
	return g
}

// NewQSvgRenderer3 constructs a new QSvgRenderer object.
func NewQSvgRenderer3(contents []byte) *QSvgRenderer {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))
	g := newQSvgRenderer(QSvgRenderer_new3(contents_alias))
	g.isSubclass = true
	return g
}

// NewQSvgRenderer4 constructs a new QSvgRenderer object.
func NewQSvgRenderer4(contents *qt.QXmlStreamReader) *QSvgRenderer {
	g := newQSvgRenderer(QSvgRenderer_new4((*QXmlStreamReader)(contents.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQSvgRenderer5 constructs a new QSvgRenderer object.
func NewQSvgRenderer5(parent *qt.QObject) *QSvgRenderer {
	g := newQSvgRenderer(QSvgRenderer_new5((*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQSvgRenderer6 constructs a new QSvgRenderer object.
func NewQSvgRenderer6(filename string, parent *qt.QObject) *QSvgRenderer {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	g := newQSvgRenderer(QSvgRenderer_new6(filename_ms, (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQSvgRenderer7 constructs a new QSvgRenderer object.
func NewQSvgRenderer7(contents []byte, parent *qt.QObject) *QSvgRenderer {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))
	g := newQSvgRenderer(QSvgRenderer_new7(contents_alias, (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
}

// NewQSvgRenderer8 constructs a new QSvgRenderer object.
func NewQSvgRenderer8(contents *qt.QXmlStreamReader, parent *qt.QObject) *QSvgRenderer {
	g := newQSvgRenderer(QSvgRenderer_new8((*QXmlStreamReader)(contents.UnsafePointer()), (*QObject)(parent.UnsafePointer())))
	g.isSubclass = true
	return g
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

func (this *QSvgRenderer) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSvgRenderer_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QSvgRenderer) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_MetaObject
func miqt_exec_callback_QSvgRenderer_MetaObject(self QSvgRenderer, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgRenderer{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QSvgRenderer) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSvgRenderer_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSvgRenderer) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgRenderer_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgRenderer_Metacast
func miqt_exec_callback_QSvgRenderer_Metacast(self QSvgRenderer, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSvgRenderer{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
