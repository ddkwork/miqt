package svg

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSvgGenerator__SvgVersion int

const (
	QSvgGenerator__SvgTiny12 QSvgGenerator__SvgVersion = 0
	QSvgGenerator__Svg11     QSvgGenerator__SvgVersion = 1
)

type QSvgGenerator struct {
	h          uintptr
	isSubclass bool
}

// NewQSvgGenerator constructs a new QSvgGenerator object.
func NewQSvgGenerator() *QSvgGenerator {

	ret := newQSvgGenerator(QSvgGenerator_new())
	ret.isSubclass = true
	return ret
}

// NewQSvgGenerator2 constructs a new QSvgGenerator object.
func NewQSvgGenerator2(version SvgVersion) *QSvgGenerator {

	ret := newQSvgGenerator(QSvgGenerator_new2(version))
	ret.isSubclass = true
	return ret
}

func (this *QSvgGenerator) Title() string {
	var _ms struct_miqt_string = QSvgGenerator_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgGenerator) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QSvgGenerator_SetTitle(this.h, title_ms)
}

func (this *QSvgGenerator) Description() string {
	var _ms struct_miqt_string = QSvgGenerator_Description(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgGenerator) SetDescription(description string) {
	description_ms := struct_miqt_string{}
	description_ms.data = CString(description)
	description_ms.len = size_t(len(description))
	defer free(unsafe.Pointer(description_ms.data))
	QSvgGenerator_SetDescription(this.h, description_ms)
}

func (this *QSvgGenerator) Size() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgGenerator_Size(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgGenerator) SetSize(size *qt.QSize) {
	QSvgGenerator_SetSize(this.h, (*QSize)(size.UnsafePointer()))
}

func (this *QSvgGenerator) ViewBox() *qt.QRect {
	_goptr := qt.UnsafeNewQRect(unsafe.Pointer(QSvgGenerator_ViewBox(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgGenerator) ViewBoxF() *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QSvgGenerator_ViewBoxF(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgGenerator) SetViewBox(viewBox *qt.QRect) {
	QSvgGenerator_SetViewBox(this.h, (*QRect)(viewBox.UnsafePointer()))
}

func (this *QSvgGenerator) SetViewBoxWithViewBox(viewBox *qt.QRectF) {
	QSvgGenerator_SetViewBoxWithViewBox(this.h, (*QRectF)(viewBox.UnsafePointer()))
}

func (this *QSvgGenerator) FileName() string {
	var _ms struct_miqt_string = QSvgGenerator_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgGenerator) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QSvgGenerator_SetFileName(this.h, fileName_ms)
}

func (this *QSvgGenerator) OutputDevice() *qt.QIODevice {
	return qt.UnsafeNewQIODevice(unsafe.Pointer(QSvgGenerator_OutputDevice(this.h)))
}

func (this *QSvgGenerator) SetOutputDevice(outputDevice *qt.QIODevice) {
	QSvgGenerator_SetOutputDevice(this.h, (*QIODevice)(outputDevice.UnsafePointer()))
}

func (this *QSvgGenerator) SetResolution(dpi int) {
	QSvgGenerator_SetResolution(this.h, (int)(dpi))
}

func (this *QSvgGenerator) Resolution() int {
	return (int)(QSvgGenerator_Resolution(this.h))
}

func (this *QSvgGenerator) SvgVersion() SvgVersion {
	xxxxxxxxx
}

func (this *QSvgGenerator) callVirtualBase_PaintEngine() *qt.QPaintEngine {

	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(QSvgGenerator_virtualbase_PaintEngine(unsafe.Pointer(this.h))))

}
func (this *QSvgGenerator) OnPaintEngine(slot func(super func() *qt.QPaintEngine) *qt.QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgGenerator_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgGenerator_PaintEngine
func miqt_exec_callback_QSvgGenerator_PaintEngine(self QSvgGenerator, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPaintEngine) *qt.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgGenerator{h: self}).callVirtualBase_PaintEngine)

	return (*QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QSvgGenerator) callVirtualBase_Metric(metric qt.QPaintDevice__PaintDeviceMetric) int {

	return (int)(QSvgGenerator_virtualbase_Metric(unsafe.Pointer(this.h), (int)(metric)))

}
func (this *QSvgGenerator) OnMetric(slot func(super func(metric qt.QPaintDevice__PaintDeviceMetric) int, metric qt.QPaintDevice__PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgGenerator_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgGenerator_Metric
func miqt_exec_callback_QSvgGenerator_Metric(self QSvgGenerator, cb intptr_t, metric int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metric qt.QPaintDevice__PaintDeviceMetric) int, metric qt.QPaintDevice__PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.QPaintDevice__PaintDeviceMetric)(metric)

	virtualReturn := gofunc((&QSvgGenerator{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QSvgGenerator) callVirtualBase_DevType() int {

	return (int)(QSvgGenerator_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QSvgGenerator) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgGenerator_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgGenerator_DevType
func miqt_exec_callback_QSvgGenerator_DevType(self QSvgGenerator, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgGenerator{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QSvgGenerator) callVirtualBase_InitPainter(painter *qt.QPainter) {

	QSvgGenerator_virtualbase_InitPainter(unsafe.Pointer(this.h), (*QPainter)(painter.UnsafePointer()))

}
func (this *QSvgGenerator) OnInitPainter(slot func(super func(painter *qt.QPainter), painter *qt.QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgGenerator_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgGenerator_InitPainter
func miqt_exec_callback_QSvgGenerator_InitPainter(self QSvgGenerator, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter), painter *qt.QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	gofunc((&QSvgGenerator{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QSvgGenerator) callVirtualBase_Redirected(offset *qt.QPoint) *qt.QPaintDevice {

	return qt.UnsafeNewQPaintDevice(unsafe.Pointer(QSvgGenerator_virtualbase_Redirected(unsafe.Pointer(this.h), (*QPoint)(offset.UnsafePointer()))))

}
func (this *QSvgGenerator) OnRedirected(slot func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgGenerator_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgGenerator_Redirected
func miqt_exec_callback_QSvgGenerator_Redirected(self QSvgGenerator, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPoint(unsafe.Pointer(offset))

	virtualReturn := gofunc((&QSvgGenerator{h: self}).callVirtualBase_Redirected, slotval1)

	return (*QPaintDevice)(virtualReturn.UnsafePointer())

}

func (this *QSvgGenerator) callVirtualBase_SharedPainter() *qt.QPainter {

	return qt.UnsafeNewQPainter(unsafe.Pointer(QSvgGenerator_virtualbase_SharedPainter(unsafe.Pointer(this.h))))

}
func (this *QSvgGenerator) OnSharedPainter(slot func(super func() *qt.QPainter) *qt.QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgGenerator_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgGenerator_SharedPainter
func miqt_exec_callback_QSvgGenerator_SharedPainter(self QSvgGenerator, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainter) *qt.QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgGenerator{h: self}).callVirtualBase_SharedPainter)

	return (*QPainter)(virtualReturn.UnsafePointer())

}
