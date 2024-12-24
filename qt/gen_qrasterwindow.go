package qt

import (
	"unsafe"
)

type QRasterWindow struct {
	h          uintptr
	isSubclass bool
}

// NewQRasterWindow constructs a new QRasterWindow object.
func NewQRasterWindow() *QRasterWindow {

	ret := newQRasterWindow(QRasterWindow_new())
	ret.isSubclass = true
	return ret
}

// NewQRasterWindow2 constructs a new QRasterWindow object.
func NewQRasterWindow2(parent *QWindow) *QRasterWindow {

	ret := newQRasterWindow(QRasterWindow_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRasterWindow) MetaObject() *QMetaObject {
	return newQMetaObject(QRasterWindow_MetaObject(this.h))
}

func (this *QRasterWindow) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRasterWindow_Metacast(this.h, param1_Cstring))
}

func QRasterWindow_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRasterWindow_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRasterWindow_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRasterWindow_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRasterWindow_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRasterWindow) callVirtualBase_Metric(metric PaintDeviceMetric) int {

	return (int)(QRasterWindow_virtualbase_Metric(unsafe.Pointer(this.h), metric))

}
func (this *QRasterWindow) OnMetric(slot func(super func(metric PaintDeviceMetric) int, metric PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Metric
func miqt_exec_callback_QRasterWindow_Metric(self QRasterWindow, cb intptr_t, metric PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(metric PaintDeviceMetric) int, metric PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QRasterWindow) callVirtualBase_Redirected(param1 *QPoint) *QPaintDevice {

	return newQPaintDevice(QRasterWindow_virtualbase_Redirected(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QRasterWindow) OnRedirected(slot func(super func(param1 *QPoint) *QPaintDevice, param1 *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Redirected
func miqt_exec_callback_QRasterWindow_Redirected(self QRasterWindow, cb intptr_t, param1 *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QPoint) *QPaintDevice, param1 *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(param1)

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QRasterWindow) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QRasterWindow_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRasterWindow) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_ResizeEvent
func miqt_exec_callback_QRasterWindow_ResizeEvent(self QRasterWindow, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QRasterWindow{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QRasterWindow) callVirtualBase_ExposeEvent(param1 *QExposeEvent) {

	QRasterWindow_virtualbase_ExposeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QRasterWindow) OnExposeEvent(slot func(super func(param1 *QExposeEvent), param1 *QExposeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_ExposeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_ExposeEvent
func miqt_exec_callback_QRasterWindow_ExposeEvent(self QRasterWindow, cb intptr_t, param1 *QExposeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QExposeEvent), param1 *QExposeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQExposeEvent(param1)

	gofunc((&QRasterWindow{h: self}).callVirtualBase_ExposeEvent, slotval1)

}

func (this *QRasterWindow) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QRasterWindow_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QRasterWindow) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_PaintEvent
func miqt_exec_callback_QRasterWindow_PaintEvent(self QRasterWindow, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QRasterWindow{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QRasterWindow) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QRasterWindow_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QRasterWindow) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRasterWindow_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRasterWindow_Event
func miqt_exec_callback_QRasterWindow_Event(self QRasterWindow, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QRasterWindow{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}
