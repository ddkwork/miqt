package qt

import (
	"unsafe"
)

type QImageIOHandler__ImageOption int

const (
	QImageIOHandler__Size                 QImageIOHandler__ImageOption = 0
	QImageIOHandler__ClipRect             QImageIOHandler__ImageOption = 1
	QImageIOHandler__Description          QImageIOHandler__ImageOption = 2
	QImageIOHandler__ScaledClipRect       QImageIOHandler__ImageOption = 3
	QImageIOHandler__ScaledSize           QImageIOHandler__ImageOption = 4
	QImageIOHandler__CompressionRatio     QImageIOHandler__ImageOption = 5
	QImageIOHandler__Gamma                QImageIOHandler__ImageOption = 6
	QImageIOHandler__Quality              QImageIOHandler__ImageOption = 7
	QImageIOHandler__Name                 QImageIOHandler__ImageOption = 8
	QImageIOHandler__SubType              QImageIOHandler__ImageOption = 9
	QImageIOHandler__IncrementalReading   QImageIOHandler__ImageOption = 10
	QImageIOHandler__Endianness           QImageIOHandler__ImageOption = 11
	QImageIOHandler__Animation            QImageIOHandler__ImageOption = 12
	QImageIOHandler__BackgroundColor      QImageIOHandler__ImageOption = 13
	QImageIOHandler__ImageFormat          QImageIOHandler__ImageOption = 14
	QImageIOHandler__SupportedSubTypes    QImageIOHandler__ImageOption = 15
	QImageIOHandler__OptimizedWrite       QImageIOHandler__ImageOption = 16
	QImageIOHandler__ProgressiveScanWrite QImageIOHandler__ImageOption = 17
	QImageIOHandler__ImageTransformation  QImageIOHandler__ImageOption = 18
)

type QImageIOHandler__Transformation int

const (
	QImageIOHandler__TransformationNone              QImageIOHandler__Transformation = 0
	QImageIOHandler__TransformationMirror            QImageIOHandler__Transformation = 1
	QImageIOHandler__TransformationFlip              QImageIOHandler__Transformation = 2
	QImageIOHandler__TransformationRotate180         QImageIOHandler__Transformation = 3
	QImageIOHandler__TransformationRotate90          QImageIOHandler__Transformation = 4
	QImageIOHandler__TransformationMirrorAndRotate90 QImageIOHandler__Transformation = 5
	QImageIOHandler__TransformationFlipAndRotate90   QImageIOHandler__Transformation = 6
	QImageIOHandler__TransformationRotate270         QImageIOHandler__Transformation = 7
)

type QImageIOPlugin__Capability int

const (
	QImageIOPlugin__CanRead            QImageIOPlugin__Capability = 1
	QImageIOPlugin__CanWrite           QImageIOPlugin__Capability = 2
	QImageIOPlugin__CanReadIncremental QImageIOPlugin__Capability = 4
)

type QImageIOHandler struct {
	h          uintptr
	isSubclass bool
}

// NewQImageIOHandler constructs a new QImageIOHandler object.
func NewQImageIOHandler() *QImageIOHandler {

	ret := newQImageIOHandler(QImageIOHandler_new())
	ret.isSubclass = true
	return ret
}

func (this *QImageIOHandler) SetDevice(device *QIODevice) {
	QImageIOHandler_SetDevice(this.h, device.cPointer())
}

func (this *QImageIOHandler) Device() *QIODevice {
	return newQIODevice(QImageIOHandler_Device(this.h))
}

func (this *QImageIOHandler) SetFormat(format []byte) {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	QImageIOHandler_SetFormat(this.h, format_alias)
}

func (this *QImageIOHandler) SetFormatWithFormat(format []byte) {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	QImageIOHandler_SetFormatWithFormat(this.h, format_alias)
}

func (this *QImageIOHandler) Format() []byte {
	var _bytearray struct_miqt_string = QImageIOHandler_Format(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QImageIOHandler) CanRead() bool {
	return (bool)(QImageIOHandler_CanRead(this.h))
}

func (this *QImageIOHandler) Read(image *QImage) bool {
	return (bool)(QImageIOHandler_Read(this.h, image.cPointer()))
}

func (this *QImageIOHandler) Write(image *QImage) bool {
	return (bool)(QImageIOHandler_Write(this.h, image.cPointer()))
}

func (this *QImageIOHandler) Option(option ImageOption) *QVariant {
	_goptr := newQVariant(QImageIOHandler_Option(this.h, option))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QImageIOHandler) SetOption(option ImageOption, value *QVariant) {
	QImageIOHandler_SetOption(this.h, option, value.cPointer())
}

func (this *QImageIOHandler) SupportsOption(option ImageOption) bool {
	return (bool)(QImageIOHandler_SupportsOption(this.h, option))
}

func (this *QImageIOHandler) JumpToNextImage() bool {
	return (bool)(QImageIOHandler_JumpToNextImage(this.h))
}

func (this *QImageIOHandler) JumpToImage(imageNumber int) bool {
	return (bool)(QImageIOHandler_JumpToImage(this.h, (int)(imageNumber)))
}

func (this *QImageIOHandler) LoopCount() int {
	return (int)(QImageIOHandler_LoopCount(this.h))
}

func (this *QImageIOHandler) ImageCount() int {
	return (int)(QImageIOHandler_ImageCount(this.h))
}

func (this *QImageIOHandler) NextImageDelay() int {
	return (int)(QImageIOHandler_NextImageDelay(this.h))
}

func (this *QImageIOHandler) CurrentImageNumber() int {
	return (int)(QImageIOHandler_CurrentImageNumber(this.h))
}

func (this *QImageIOHandler) CurrentImageRect() *QRect {
	_goptr := newQRect(QImageIOHandler_CurrentImageRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QImageIOHandler_AllocateImage(size QSize, format QImage__Format, image *QImage) bool {
	return (bool)(QImageIOHandler_AllocateImage(size.cPointer(), (int)(format), image.cPointer()))
}
func (this *QImageIOHandler) OnCanRead(slot func() bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_CanRead(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_CanRead
func miqt_exec_callback_QImageIOHandler_CanRead(self QImageIOHandler, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func() bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (bool)(virtualReturn)

}
func (this *QImageIOHandler) OnRead(slot func(image *QImage) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_Read(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_Read
func miqt_exec_callback_QImageIOHandler_Read(self QImageIOHandler, cb intptr_t, image *QImage) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(image *QImage) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQImage(image)

	virtualReturn := gofunc(slotval1)

	return (bool)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_Write(image *QImage) bool {

	return (bool)(QImageIOHandler_virtualbase_Write(unsafe.Pointer(this.h), image.cPointer()))

}
func (this *QImageIOHandler) OnWrite(slot func(super func(image *QImage) bool, image *QImage) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_Write(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_Write
func miqt_exec_callback_QImageIOHandler_Write(self QImageIOHandler, cb intptr_t, image *QImage) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(image *QImage) bool, image *QImage) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQImage(image)

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_Write, slotval1)

	return (bool)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_Option(option ImageOption) *QVariant {

	_goptr := newQVariant(QImageIOHandler_virtualbase_Option(unsafe.Pointer(this.h), option))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QImageIOHandler) OnOption(slot func(super func(option ImageOption) *QVariant, option ImageOption) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_Option(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_Option
func miqt_exec_callback_QImageIOHandler_Option(self QImageIOHandler, cb intptr_t, option ImageOption) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option ImageOption) *QVariant, option ImageOption) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_Option, slotval1)

	return virtualReturn.cPointer()

}

func (this *QImageIOHandler) callVirtualBase_SetOption(option ImageOption, value *QVariant) {

	QImageIOHandler_virtualbase_SetOption(unsafe.Pointer(this.h), option, value.cPointer())

}
func (this *QImageIOHandler) OnSetOption(slot func(super func(option ImageOption, value *QVariant), option ImageOption, value *QVariant)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_SetOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_SetOption
func miqt_exec_callback_QImageIOHandler_SetOption(self QImageIOHandler, cb intptr_t, option ImageOption, value *QVariant) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option ImageOption, value *QVariant), option ImageOption, value *QVariant))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQVariant(value)

	gofunc((&QImageIOHandler{h: self}).callVirtualBase_SetOption, slotval1, slotval2)

}

func (this *QImageIOHandler) callVirtualBase_SupportsOption(option ImageOption) bool {

	return (bool)(QImageIOHandler_virtualbase_SupportsOption(unsafe.Pointer(this.h), option))

}
func (this *QImageIOHandler) OnSupportsOption(slot func(super func(option ImageOption) bool, option ImageOption) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_SupportsOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_SupportsOption
func miqt_exec_callback_QImageIOHandler_SupportsOption(self QImageIOHandler, cb intptr_t, option ImageOption) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option ImageOption) bool, option ImageOption) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_SupportsOption, slotval1)

	return (bool)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_JumpToNextImage() bool {

	return (bool)(QImageIOHandler_virtualbase_JumpToNextImage(unsafe.Pointer(this.h)))

}
func (this *QImageIOHandler) OnJumpToNextImage(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_JumpToNextImage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_JumpToNextImage
func miqt_exec_callback_QImageIOHandler_JumpToNextImage(self QImageIOHandler, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_JumpToNextImage)

	return (bool)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_JumpToImage(imageNumber int) bool {

	return (bool)(QImageIOHandler_virtualbase_JumpToImage(unsafe.Pointer(this.h), (int)(imageNumber)))

}
func (this *QImageIOHandler) OnJumpToImage(slot func(super func(imageNumber int) bool, imageNumber int) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_JumpToImage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_JumpToImage
func miqt_exec_callback_QImageIOHandler_JumpToImage(self QImageIOHandler, cb intptr_t, imageNumber int) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(imageNumber int) bool, imageNumber int) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(imageNumber)

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_JumpToImage, slotval1)

	return (bool)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_LoopCount() int {

	return (int)(QImageIOHandler_virtualbase_LoopCount(unsafe.Pointer(this.h)))

}
func (this *QImageIOHandler) OnLoopCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_LoopCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_LoopCount
func miqt_exec_callback_QImageIOHandler_LoopCount(self QImageIOHandler, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_LoopCount)

	return (int)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_ImageCount() int {

	return (int)(QImageIOHandler_virtualbase_ImageCount(unsafe.Pointer(this.h)))

}
func (this *QImageIOHandler) OnImageCount(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_ImageCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_ImageCount
func miqt_exec_callback_QImageIOHandler_ImageCount(self QImageIOHandler, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_ImageCount)

	return (int)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_NextImageDelay() int {

	return (int)(QImageIOHandler_virtualbase_NextImageDelay(unsafe.Pointer(this.h)))

}
func (this *QImageIOHandler) OnNextImageDelay(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_NextImageDelay(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_NextImageDelay
func miqt_exec_callback_QImageIOHandler_NextImageDelay(self QImageIOHandler, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_NextImageDelay)

	return (int)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_CurrentImageNumber() int {

	return (int)(QImageIOHandler_virtualbase_CurrentImageNumber(unsafe.Pointer(this.h)))

}
func (this *QImageIOHandler) OnCurrentImageNumber(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_CurrentImageNumber(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_CurrentImageNumber
func miqt_exec_callback_QImageIOHandler_CurrentImageNumber(self QImageIOHandler, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_CurrentImageNumber)

	return (int)(virtualReturn)

}

func (this *QImageIOHandler) callVirtualBase_CurrentImageRect() *QRect {

	_goptr := newQRect(QImageIOHandler_virtualbase_CurrentImageRect(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QImageIOHandler) OnCurrentImageRect(slot func(super func() *QRect) *QRect) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOHandler_override_virtual_CurrentImageRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOHandler_CurrentImageRect
func miqt_exec_callback_QImageIOHandler_CurrentImageRect(self QImageIOHandler, cb intptr_t) *QRect {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRect) *QRect)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOHandler{h: self}).callVirtualBase_CurrentImageRect)

	return virtualReturn.cPointer()

}

type QImageIOPlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQImageIOPlugin constructs a new QImageIOPlugin object.
func NewQImageIOPlugin() *QImageIOPlugin {

	ret := newQImageIOPlugin(QImageIOPlugin_new())
	ret.isSubclass = true
	return ret
}

// NewQImageIOPlugin2 constructs a new QImageIOPlugin object.
func NewQImageIOPlugin2(parent *QObject) *QImageIOPlugin {

	ret := newQImageIOPlugin(QImageIOPlugin_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QImageIOPlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QImageIOPlugin_MetaObject(this.h))
}

func (this *QImageIOPlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QImageIOPlugin_Metacast(this.h, param1_Cstring))
}

func QImageIOPlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QImageIOPlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QImageIOPlugin) Capabilities(device *QIODevice, format []byte) Capabilities {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	xxxxxxxxx
}

func (this *QImageIOPlugin) Create(device *QIODevice, format []byte) *QImageIOHandler {
	format_alias := struct_miqt_string{}
	format_alias.data = (char)(unsafe.Pointer(&format[0]))
	format_alias.len = size_t(len(format))
	return newQImageIOHandler(QImageIOPlugin_Create(this.h, device.cPointer(), format_alias))
}

func QImageIOPlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QImageIOPlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QImageIOPlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QImageIOPlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QImageIOPlugin) OnCapabilities(slot func(device *QIODevice, format []byte) Capabilities) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_Capabilities(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_Capabilities
func miqt_exec_callback_QImageIOPlugin_Capabilities(self QImageIOPlugin, cb intptr_t, device *QIODevice, format struct_miqt_string) Capabilities {
	gofunc, ok := cgo.Handle(cb).Value().(func(device *QIODevice, format []byte) Capabilities)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQIODevice(device)

	var format_bytearray struct_miqt_string = format
	format_ret := GoBytes(unsafe.Pointer(format_bytearray.data), int(int64(format_bytearray.len)))
	free(unsafe.Pointer(format_bytearray.data))
	slotval2 := format_ret

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn

}
func (this *QImageIOPlugin) OnCreate(slot func(device *QIODevice, format []byte) *QImageIOHandler) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_Create(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_Create
func miqt_exec_callback_QImageIOPlugin_Create(self QImageIOPlugin, cb intptr_t, device *QIODevice, format struct_miqt_string) *QImageIOHandler {
	gofunc, ok := cgo.Handle(cb).Value().(func(device *QIODevice, format []byte) *QImageIOHandler)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQIODevice(device)

	var format_bytearray struct_miqt_string = format
	format_ret := GoBytes(unsafe.Pointer(format_bytearray.data), int(int64(format_bytearray.len)))
	free(unsafe.Pointer(format_bytearray.data))
	slotval2 := format_ret

	virtualReturn := gofunc(slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QImageIOPlugin) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QImageIOPlugin_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QImageIOPlugin) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_Event
func miqt_exec_callback_QImageIOPlugin_Event(self QImageIOPlugin, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QImageIOPlugin{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QImageIOPlugin) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QImageIOPlugin_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QImageIOPlugin) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_EventFilter
func miqt_exec_callback_QImageIOPlugin_EventFilter(self QImageIOPlugin, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QImageIOPlugin{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QImageIOPlugin) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QImageIOPlugin_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QImageIOPlugin) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_TimerEvent
func miqt_exec_callback_QImageIOPlugin_TimerEvent(self QImageIOPlugin, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QImageIOPlugin{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QImageIOPlugin) callVirtualBase_ChildEvent(event *QChildEvent) {

	QImageIOPlugin_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QImageIOPlugin) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_ChildEvent
func miqt_exec_callback_QImageIOPlugin_ChildEvent(self QImageIOPlugin, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QImageIOPlugin{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QImageIOPlugin) callVirtualBase_CustomEvent(event *QEvent) {

	QImageIOPlugin_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QImageIOPlugin) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_CustomEvent
func miqt_exec_callback_QImageIOPlugin_CustomEvent(self QImageIOPlugin, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QImageIOPlugin{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QImageIOPlugin) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QImageIOPlugin_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QImageIOPlugin) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_ConnectNotify
func miqt_exec_callback_QImageIOPlugin_ConnectNotify(self QImageIOPlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QImageIOPlugin{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QImageIOPlugin) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QImageIOPlugin_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QImageIOPlugin) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_DisconnectNotify
func miqt_exec_callback_QImageIOPlugin_DisconnectNotify(self QImageIOPlugin, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QImageIOPlugin{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
