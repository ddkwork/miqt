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
	g := newQImageIOHandler(QImageIOHandler_new())
	g.isSubclass = true
	return g
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

type QImageIOPlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQImageIOPlugin constructs a new QImageIOPlugin object.
func NewQImageIOPlugin() *QImageIOPlugin {
	g := newQImageIOPlugin(QImageIOPlugin_new())
	g.isSubclass = true
	return g
}

// NewQImageIOPlugin2 constructs a new QImageIOPlugin object.
func NewQImageIOPlugin2(parent *QObject) *QImageIOPlugin {
	g := newQImageIOPlugin(QImageIOPlugin_new2(parent.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QImageIOPlugin) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QImageIOPlugin_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QImageIOPlugin) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_MetaObject
func miqt_exec_callback_QImageIOPlugin_MetaObject(self QImageIOPlugin, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QImageIOPlugin{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QImageIOPlugin) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QImageIOPlugin_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QImageIOPlugin) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QImageIOPlugin_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QImageIOPlugin_Metacast
func miqt_exec_callback_QImageIOPlugin_Metacast(self QImageIOPlugin, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QImageIOPlugin{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
