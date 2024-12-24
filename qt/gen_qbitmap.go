package qt

import (
	"unsafe"
)

type QBitmap struct {
	h          uintptr
	isSubclass bool
}

// NewQBitmap constructs a new QBitmap object.
func NewQBitmap() *QBitmap {

	ret := newQBitmap(QBitmap_new())
	ret.isSubclass = true
	return ret
}

// NewQBitmap2 constructs a new QBitmap object.
func NewQBitmap2(param1 *QPixmap) *QBitmap {

	ret := newQBitmap(QBitmap_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBitmap3 constructs a new QBitmap object.
func NewQBitmap3(w int, h int) *QBitmap {

	ret := newQBitmap(QBitmap_new3((int)(w), (int)(h)))
	ret.isSubclass = true
	return ret
}

// NewQBitmap4 constructs a new QBitmap object.
func NewQBitmap4(param1 *QSize) *QBitmap {

	ret := newQBitmap(QBitmap_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQBitmap5 constructs a new QBitmap object.
func NewQBitmap5(fileName string) *QBitmap {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQBitmap(QBitmap_new5(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQBitmap6 constructs a new QBitmap object.
func NewQBitmap6(fileName string, format string) *QBitmap {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	format_Cstring := CString(format)
	defer free(unsafe.Pointer(format_Cstring))

	ret := newQBitmap(QBitmap_new6(fileName_ms, format_Cstring))
	ret.isSubclass = true
	return ret
}

func (this *QBitmap) OperatorAssign(param1 *QPixmap) {
	QBitmap_OperatorAssign(this.h, param1.cPointer())
}

func (this *QBitmap) Swap(other *QBitmap) {
	QBitmap_Swap(this.h, other.cPointer())
}

func (this *QBitmap) Clear() {
	QBitmap_Clear(this.h)
}

func QBitmap_FromImage(image *QImage) *QBitmap {
	_goptr := newQBitmap(QBitmap_FromImage(image.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QBitmap_FromData(size *QSize, bits *byte) *QBitmap {
	_goptr := newQBitmap(QBitmap_FromData(size.cPointer(), (*uchar)(unsafe.Pointer(bits))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QBitmap_FromPixmap(pixmap *QPixmap) *QBitmap {
	_goptr := newQBitmap(QBitmap_FromPixmap(pixmap.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBitmap) Transformed(matrix *QTransform) *QBitmap {
	_goptr := newQBitmap(QBitmap_Transformed(this.h, matrix.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBitmap) OperatorAssignWithQBitmap(param1 *QBitmap) {
	QBitmap_OperatorAssignWithQBitmap(this.h, param1.cPointer())
}

func QBitmap_FromImage2(image *QImage, flags ImageConversionFlag) *QBitmap {
	_goptr := newQBitmap(QBitmap_FromImage2(image.cPointer(), (int)(flags)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QBitmap_FromData3(size *QSize, bits *byte, monoFormat QImage__Format) *QBitmap {
	_goptr := newQBitmap(QBitmap_FromData3(size.cPointer(), (*uchar)(unsafe.Pointer(bits)), (int)(monoFormat)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QBitmap) callVirtualBase_DevType() int {

	return (int)(QBitmap_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QBitmap) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBitmap_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBitmap_DevType
func miqt_exec_callback_QBitmap_DevType(self QBitmap, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBitmap{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QBitmap) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QBitmap_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QBitmap) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBitmap_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBitmap_PaintEngine
func miqt_exec_callback_QBitmap_PaintEngine(self QBitmap, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QBitmap{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QBitmap) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QBitmap_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QBitmap) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QBitmap_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QBitmap_Metric
func miqt_exec_callback_QBitmap_Metric(self QBitmap, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QBitmap{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}
