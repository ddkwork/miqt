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

func (this *QRhiWidget) callVirtualBase_Initialize(cb *QRhiCommandBuffer) {

	QRhiWidget_virtualbase_Initialize(unsafe.Pointer(this.h), cb)

}
func (this *QRhiWidget) OnInitialize(slot func(super func(cb *QRhiCommandBuffer), cb *QRhiCommandBuffer)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_Initialize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_Initialize
func miqt_exec_callback_QRhiWidget_Initialize(self QRhiWidget, cb intptr_t, cb *QRhiCommandBuffer) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cb *QRhiCommandBuffer), cb *QRhiCommandBuffer))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QRhiWidget{h: self}).callVirtualBase_Initialize, slotval1)

}

func (this *QRhiWidget) callVirtualBase_Render(cb *QRhiCommandBuffer) {

	QRhiWidget_virtualbase_Render(unsafe.Pointer(this.h), cb)

}
func (this *QRhiWidget) OnRender(slot func(super func(cb *QRhiCommandBuffer), cb *QRhiCommandBuffer)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_Render(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_Render
func miqt_exec_callback_QRhiWidget_Render(self QRhiWidget, cb intptr_t, cb *QRhiCommandBuffer) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cb *QRhiCommandBuffer), cb *QRhiCommandBuffer))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	gofunc((&QRhiWidget{h: self}).callVirtualBase_Render, slotval1)

}

func (this *QRhiWidget) callVirtualBase_ReleaseResources() {

	QRhiWidget_virtualbase_ReleaseResources(unsafe.Pointer(this.h))

}
func (this *QRhiWidget) OnReleaseResources(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_ReleaseResources(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ReleaseResources
func miqt_exec_callback_QRhiWidget_ReleaseResources(self QRhiWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QRhiWidget{h: self}).callVirtualBase_ReleaseResources)

}

func (this *QRhiWidget) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QRhiWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QRhiWidget) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ResizeEvent
func miqt_exec_callback_QRhiWidget_ResizeEvent(self QRhiWidget, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QRhiWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QRhiWidget) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_PaintEvent
func miqt_exec_callback_QRhiWidget_PaintEvent(self QRhiWidget, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QRhiWidget_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QRhiWidget) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_Event
func miqt_exec_callback_QRhiWidget_Event(self QRhiWidget, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QRhiWidget) callVirtualBase_DevType() int {

	return (int)(QRhiWidget_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QRhiWidget) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_DevType
func miqt_exec_callback_QRhiWidget_DevType(self QRhiWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QRhiWidget) callVirtualBase_SetVisible(visible bool) {

	QRhiWidget_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QRhiWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_SetVisible
func miqt_exec_callback_QRhiWidget_SetVisible(self QRhiWidget, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QRhiWidget) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QRhiWidget_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QRhiWidget) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_SizeHint
func miqt_exec_callback_QRhiWidget_SizeHint(self QRhiWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QRhiWidget) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QRhiWidget_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QRhiWidget) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MinimumSizeHint
func miqt_exec_callback_QRhiWidget_MinimumSizeHint(self QRhiWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QRhiWidget) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QRhiWidget_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QRhiWidget) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_HeightForWidth
func miqt_exec_callback_QRhiWidget_HeightForWidth(self QRhiWidget, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QRhiWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QRhiWidget_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QRhiWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_HasHeightForWidth
func miqt_exec_callback_QRhiWidget_HasHeightForWidth(self QRhiWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QRhiWidget) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QRhiWidget_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QRhiWidget) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_PaintEngine
func miqt_exec_callback_QRhiWidget_PaintEngine(self QRhiWidget, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QRhiWidget) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QRhiWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MousePressEvent
func miqt_exec_callback_QRhiWidget_MousePressEvent(self QRhiWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QRhiWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MouseReleaseEvent
func miqt_exec_callback_QRhiWidget_MouseReleaseEvent(self QRhiWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QRhiWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MouseDoubleClickEvent
func miqt_exec_callback_QRhiWidget_MouseDoubleClickEvent(self QRhiWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QRhiWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MouseMoveEvent
func miqt_exec_callback_QRhiWidget_MouseMoveEvent(self QRhiWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QRhiWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_WheelEvent
func miqt_exec_callback_QRhiWidget_WheelEvent(self QRhiWidget, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QRhiWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_KeyPressEvent
func miqt_exec_callback_QRhiWidget_KeyPressEvent(self QRhiWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QRhiWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_KeyReleaseEvent
func miqt_exec_callback_QRhiWidget_KeyReleaseEvent(self QRhiWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QRhiWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_FocusInEvent
func miqt_exec_callback_QRhiWidget_FocusInEvent(self QRhiWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QRhiWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_FocusOutEvent
func miqt_exec_callback_QRhiWidget_FocusOutEvent(self QRhiWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QRhiWidget_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_EnterEvent
func miqt_exec_callback_QRhiWidget_EnterEvent(self QRhiWidget, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_LeaveEvent(event *QEvent) {

	QRhiWidget_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_LeaveEvent
func miqt_exec_callback_QRhiWidget_LeaveEvent(self QRhiWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QRhiWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_MoveEvent
func miqt_exec_callback_QRhiWidget_MoveEvent(self QRhiWidget, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QRhiWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_CloseEvent
func miqt_exec_callback_QRhiWidget_CloseEvent(self QRhiWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QRhiWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ContextMenuEvent
func miqt_exec_callback_QRhiWidget_ContextMenuEvent(self QRhiWidget, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QRhiWidget_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_TabletEvent
func miqt_exec_callback_QRhiWidget_TabletEvent(self QRhiWidget, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_ActionEvent(event *QActionEvent) {

	QRhiWidget_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ActionEvent
func miqt_exec_callback_QRhiWidget_ActionEvent(self QRhiWidget, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QRhiWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_DragEnterEvent
func miqt_exec_callback_QRhiWidget_DragEnterEvent(self QRhiWidget, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QRhiWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_DragMoveEvent
func miqt_exec_callback_QRhiWidget_DragMoveEvent(self QRhiWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QRhiWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_DragLeaveEvent
func miqt_exec_callback_QRhiWidget_DragLeaveEvent(self QRhiWidget, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_DropEvent(event *QDropEvent) {

	QRhiWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_DropEvent
func miqt_exec_callback_QRhiWidget_DropEvent(self QRhiWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_ShowEvent(event *QShowEvent) {

	QRhiWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ShowEvent
func miqt_exec_callback_QRhiWidget_ShowEvent(self QRhiWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_HideEvent(event *QHideEvent) {

	QRhiWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRhiWidget) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_HideEvent
func miqt_exec_callback_QRhiWidget_HideEvent(self QRhiWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QRhiWidget_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QRhiWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_NativeEvent
func miqt_exec_callback_QRhiWidget_NativeEvent(self QRhiWidget, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QRhiWidget) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QRhiWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRhiWidget) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_ChangeEvent
func miqt_exec_callback_QRhiWidget_ChangeEvent(self QRhiWidget, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QRhiWidget_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QRhiWidget) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_Metric
func miqt_exec_callback_QRhiWidget_Metric(self QRhiWidget, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QRhiWidget) callVirtualBase_InitPainter(painter *QPainter) {

	QRhiWidget_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QRhiWidget) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_InitPainter
func miqt_exec_callback_QRhiWidget_InitPainter(self QRhiWidget, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QRhiWidget) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QRhiWidget_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QRhiWidget) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_Redirected
func miqt_exec_callback_QRhiWidget_Redirected(self QRhiWidget, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QRhiWidget) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QRhiWidget_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QRhiWidget) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_SharedPainter
func miqt_exec_callback_QRhiWidget_SharedPainter(self QRhiWidget, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QRhiWidget) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QRhiWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRhiWidget) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_InputMethodEvent
func miqt_exec_callback_QRhiWidget_InputMethodEvent(self QRhiWidget, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QRhiWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QRhiWidget) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QRhiWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QRhiWidget) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_InputMethodQuery
func miqt_exec_callback_QRhiWidget_InputMethodQuery(self QRhiWidget, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QRhiWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QRhiWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QRhiWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRhiWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRhiWidget_FocusNextPrevChild
func miqt_exec_callback_QRhiWidget_FocusNextPrevChild(self QRhiWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QRhiWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
