package qt

import (
	"unsafe"
)

type QRubberBand__Shape int

const (
	QRubberBand__Line      QRubberBand__Shape = 0
	QRubberBand__Rectangle QRubberBand__Shape = 1
)

type QRubberBand struct {
	h          uintptr
	isSubclass bool
}

// NewQRubberBand constructs a new QRubberBand object.
func NewQRubberBand(param1 Shape) *QRubberBand {
	g := newQRubberBand(QRubberBand_new(param1))
	g.isSubclass = true
	return g
}

// NewQRubberBand2 constructs a new QRubberBand object.
func NewQRubberBand2(param1 Shape, param2 *QWidget) *QRubberBand {
	g := newQRubberBand(QRubberBand_new2(param1, param2.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QRubberBand) MetaObject() *QMetaObject {
	return newQMetaObject(QRubberBand_MetaObject(this.h))
}

func (this *QRubberBand) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRubberBand_Metacast(this.h, param1_Cstring))
}

func QRubberBand_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRubberBand_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRubberBand) Shape() Shape {
	xxxxxxxxx
}

func (this *QRubberBand) SetGeometry(r *QRect) {
	QRubberBand_SetGeometry(this.h, r.cPointer())
}

func (this *QRubberBand) SetGeometry2(x int, y int, w int, h int) {
	QRubberBand_SetGeometry2(this.h, (int)(x), (int)(y), (int)(w), (int)(h))
}

func (this *QRubberBand) Move(x int, y int) {
	QRubberBand_Move(this.h, (int)(x), (int)(y))
}

func (this *QRubberBand) MoveWithQPoint(p *QPoint) {
	QRubberBand_MoveWithQPoint(this.h, p.cPointer())
}

func (this *QRubberBand) Resize(w int, h int) {
	QRubberBand_Resize(this.h, (int)(w), (int)(h))
}

func (this *QRubberBand) ResizeWithQSize(s *QSize) {
	QRubberBand_ResizeWithQSize(this.h, s.cPointer())
}

func QRubberBand_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRubberBand_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRubberBand_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRubberBand_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRubberBand) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QRubberBand_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QRubberBand) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_MetaObject
func miqt_exec_callback_QRubberBand_MetaObject(self QRubberBand, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QRubberBand) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QRubberBand_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QRubberBand) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRubberBand_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRubberBand_Metacast
func miqt_exec_callback_QRubberBand_Metacast(self QRubberBand, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QRubberBand{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
