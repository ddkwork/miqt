package qt

import (
	"unsafe"
)

type QPicture struct {
	h          uintptr
	isSubclass bool
}

// NewQPicture constructs a new QPicture object.
func NewQPicture() *QPicture {

	ret := newQPicture(QPicture_new())
	ret.isSubclass = true
	return ret
}

// NewQPicture2 constructs a new QPicture object.
func NewQPicture2(param1 *QPicture) *QPicture {

	ret := newQPicture(QPicture_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPicture3 constructs a new QPicture object.
func NewQPicture3(formatVersion int) *QPicture {

	ret := newQPicture(QPicture_new3((int)(formatVersion)))
	ret.isSubclass = true
	return ret
}

func (this *QPicture) IsNull() bool {
	return (bool)(QPicture_IsNull(this.h))
}

func (this *QPicture) DevType() int {
	return (int)(QPicture_DevType(this.h))
}

func (this *QPicture) Size() uint {
	return (uint)(QPicture_Size(this.h))
}

func (this *QPicture) Data() string {
	_ret := QPicture_Data(this.h)
	return GoString(_ret)
}

func (this *QPicture) SetData(data string, size uint) {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))
	QPicture_SetData(this.h, data_Cstring, (uint)(size))
}

func (this *QPicture) Play(p *QPainter) bool {
	return (bool)(QPicture_Play(this.h, p.cPointer()))
}

func (this *QPicture) Load(dev *QIODevice) bool {
	return (bool)(QPicture_Load(this.h, dev.cPointer()))
}

func (this *QPicture) LoadWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QPicture_LoadWithFileName(this.h, fileName_ms))
}

func (this *QPicture) Save(dev *QIODevice) bool {
	return (bool)(QPicture_Save(this.h, dev.cPointer()))
}

func (this *QPicture) SaveWithFileName(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QPicture_SaveWithFileName(this.h, fileName_ms))
}

func (this *QPicture) BoundingRect() *QRect {
	_goptr := newQRect(QPicture_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPicture) SetBoundingRect(r *QRect) {
	QPicture_SetBoundingRect(this.h, r.cPointer())
}

func (this *QPicture) OperatorAssign(p *QPicture) {
	QPicture_OperatorAssign(this.h, p.cPointer())
}

func (this *QPicture) Swap(other *QPicture) {
	QPicture_Swap(this.h, other.cPointer())
}

func (this *QPicture) Detach() {
	QPicture_Detach(this.h)
}

func (this *QPicture) IsDetached() bool {
	return (bool)(QPicture_IsDetached(this.h))
}

func (this *QPicture) PaintEngine() *QPaintEngine {
	return newQPaintEngine(QPicture_PaintEngine(this.h))
}

func (this *QPicture) DataPtr() *DataPtr {
	xxxxxxxxx
}

func (this *QPicture) callVirtualBase_DevType() int {

	return (int)(QPicture_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QPicture) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_DevType
func miqt_exec_callback_QPicture_DevType(self QPicture, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPicture{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QPicture) callVirtualBase_SetData(data string, size uint) {
	data_Cstring := CString(data)
	defer free(unsafe.Pointer(data_Cstring))

	QPicture_virtualbase_SetData(unsafe.Pointer(this.h), data_Cstring, (uint)(size))

}
func (this *QPicture) OnSetData(slot func(super func(data string, size uint), data string, size uint)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_SetData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_SetData
func miqt_exec_callback_QPicture_SetData(self QPicture, cb intptr_t, data *const_char, size uint) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(data string, size uint), data string, size uint))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	data_ret := data
	slotval1 := GoString(data_ret)

	slotval2 := (uint)(size)

	gofunc((&QPicture{h: self}).callVirtualBase_SetData, slotval1, slotval2)

}

func (this *QPicture) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QPicture_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QPicture) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_PaintEngine
func miqt_exec_callback_QPicture_PaintEngine(self QPicture, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPicture{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QPicture) callVirtualBase_Metric(m PaintDeviceMetric) int {

	return (int)(QPicture_virtualbase_Metric(unsafe.Pointer(this.h), m))

}
func (this *QPicture) OnMetric(slot func(super func(m PaintDeviceMetric) int, m PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_Metric
func miqt_exec_callback_QPicture_Metric(self QPicture, cb intptr_t, m PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(m PaintDeviceMetric) int, m PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QPicture{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QPicture) callVirtualBase_InitPainter(painter *QPainter) {

	QPicture_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QPicture) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_InitPainter
func miqt_exec_callback_QPicture_InitPainter(self QPicture, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QPicture{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QPicture) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QPicture_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QPicture) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_Redirected
func miqt_exec_callback_QPicture_Redirected(self QPicture, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QPicture{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QPicture) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QPicture_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QPicture) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPicture_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPicture_SharedPainter
func miqt_exec_callback_QPicture_SharedPainter(self QPicture, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPicture{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}
