package qt

import (
	"unsafe"
)

type QRhiWidget__Api int

const (
	QRhiWidget__Null       QRhiWidget__Api = 0
	QRhiWidget__OpenGL     QRhiWidget__Api = 1
	QRhiWidget__Metal      QRhiWidget__Api = 2
	QRhiWidget__Vulkan     QRhiWidget__Api = 3
	QRhiWidget__Direct3D11 QRhiWidget__Api = 4
	QRhiWidget__Direct3D12 QRhiWidget__Api = 5
)

type QRhiWidget__TextureFormat int

const (
	QRhiWidget__RGBA8   QRhiWidget__TextureFormat = 0
	QRhiWidget__RGBA16F QRhiWidget__TextureFormat = 1
	QRhiWidget__RGBA32F QRhiWidget__TextureFormat = 2
	QRhiWidget__RGB10A2 QRhiWidget__TextureFormat = 3
)

type QRhiWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQRhiWidget constructs a new QRhiWidget object.
func NewQRhiWidget(parent *QWidget) *QRhiWidget {
	ret := newQRhiWidget(QRhiWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRhiWidget2 constructs a new QRhiWidget object.
func NewQRhiWidget2() *QRhiWidget {
	ret := newQRhiWidget(QRhiWidget_new2())
	ret.isSubclass = true
	return ret
}

// NewQRhiWidget3 constructs a new QRhiWidget object.
func NewQRhiWidget3(parent *QWidget, f WindowType) *QRhiWidget {
	ret := newQRhiWidget(QRhiWidget_new3(parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

func (this *QRhiWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QRhiWidget_MetaObject(this.h))
}

func (this *QRhiWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRhiWidget_Metacast(this.h, param1_Cstring))
}

func QRhiWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRhiWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRhiWidget) Api() Api {
	xxxxxxxxx
}

func (this *QRhiWidget) SetApi(api Api) {
	QRhiWidget_SetApi(this.h, api)
}

func (this *QRhiWidget) IsDebugLayerEnabled() bool {
	return (bool)(QRhiWidget_IsDebugLayerEnabled(this.h))
}

func (this *QRhiWidget) SetDebugLayerEnabled(enable bool) {
	QRhiWidget_SetDebugLayerEnabled(this.h, (bool)(enable))
}

func (this *QRhiWidget) SampleCount() int {
	return (int)(QRhiWidget_SampleCount(this.h))
}

func (this *QRhiWidget) SetSampleCount(samples int) {
	QRhiWidget_SetSampleCount(this.h, (int)(samples))
}

func (this *QRhiWidget) ColorBufferFormat() TextureFormat {
	xxxxxxxxx
}

func (this *QRhiWidget) SetColorBufferFormat(format TextureFormat) {
	QRhiWidget_SetColorBufferFormat(this.h, format)
}

func (this *QRhiWidget) FixedColorBufferSize() *QSize {
	_goptr := newQSize(QRhiWidget_FixedColorBufferSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRhiWidget) SetFixedColorBufferSize(pixelSize QSize) {
	QRhiWidget_SetFixedColorBufferSize(this.h, pixelSize.cPointer())
}

func (this *QRhiWidget) SetFixedColorBufferSize2(w int, h int) {
	QRhiWidget_SetFixedColorBufferSize2(this.h, (int)(w), (int)(h))
}

func (this *QRhiWidget) IsMirrorVerticallyEnabled() bool {
	return (bool)(QRhiWidget_IsMirrorVerticallyEnabled(this.h))
}

func (this *QRhiWidget) SetMirrorVertically(enabled bool) {
	QRhiWidget_SetMirrorVertically(this.h, (bool)(enabled))
}

func (this *QRhiWidget) GrabFramebuffer() *QImage {
	_goptr := newQImage(QRhiWidget_GrabFramebuffer(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRhiWidget) FrameSubmitted() {
	QRhiWidget_FrameSubmitted(this.h)
}

func (this *QRhiWidget) OnFrameSubmitted(slot func()) {
	QRhiWidget_connect_FrameSubmitted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_FrameSubmitted
func miqt_exec_callback_QRhiWidget_FrameSubmitted(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QRhiWidget) RenderFailed() {
	QRhiWidget_RenderFailed(this.h)
}

func (this *QRhiWidget) OnRenderFailed(slot func()) {
	QRhiWidget_connect_RenderFailed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_RenderFailed
func miqt_exec_callback_QRhiWidget_RenderFailed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QRhiWidget) SampleCountChanged(samples int) {
	QRhiWidget_SampleCountChanged(this.h, (int)(samples))
}

func (this *QRhiWidget) OnSampleCountChanged(slot func(samples int)) {
	QRhiWidget_connect_SampleCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_SampleCountChanged
func miqt_exec_callback_QRhiWidget_SampleCountChanged(cb intptr_t, samples int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(samples int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(samples)

	gofunc(slotval1)
}

func (this *QRhiWidget) ColorBufferFormatChanged(format TextureFormat) {
	QRhiWidget_ColorBufferFormatChanged(this.h, format)
}

func (this *QRhiWidget) OnColorBufferFormatChanged(slot func(format TextureFormat)) {
	QRhiWidget_connect_ColorBufferFormatChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ColorBufferFormatChanged
func miqt_exec_callback_QRhiWidget_ColorBufferFormatChanged(cb intptr_t, format TextureFormat) {
	gofunc, ok := cgo.Handle(cb).Value().(func(format TextureFormat))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc(slotval1)
}

func (this *QRhiWidget) FixedColorBufferSizeChanged(pixelSize *QSize) {
	QRhiWidget_FixedColorBufferSizeChanged(this.h, pixelSize.cPointer())
}

func (this *QRhiWidget) OnFixedColorBufferSizeChanged(slot func(pixelSize *QSize)) {
	QRhiWidget_connect_FixedColorBufferSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_FixedColorBufferSizeChanged
func miqt_exec_callback_QRhiWidget_FixedColorBufferSizeChanged(cb intptr_t, pixelSize *QSize) {
	gofunc, ok := cgo.Handle(cb).Value().(func(pixelSize *QSize))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSize(pixelSize)

	gofunc(slotval1)
}

func (this *QRhiWidget) MirrorVerticallyChanged(enabled bool) {
	QRhiWidget_MirrorVerticallyChanged(this.h, (bool)(enabled))
}

func (this *QRhiWidget) OnMirrorVerticallyChanged(slot func(enabled bool)) {
	QRhiWidget_connect_MirrorVerticallyChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MirrorVerticallyChanged
func miqt_exec_callback_QRhiWidget_MirrorVerticallyChanged(cb intptr_t, enabled bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(enabled bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(enabled)

	gofunc(slotval1)
}

func QRhiWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRhiWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRhiWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRhiWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRhiWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QRhiWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QRhiWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MetaObject
func miqt_exec_callback_QRhiWidget_MetaObject(self QRhiWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QRhiWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QRhiWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QRhiWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_Metacast
func miqt_exec_callback_QRhiWidget_Metacast(self QRhiWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
