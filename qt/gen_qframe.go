package qt

import (
	"unsafe"
)

type QFrame__Shape int

const (
	QFrame__NoFrame     QFrame__Shape = 0
	QFrame__Box         QFrame__Shape = 1
	QFrame__Panel       QFrame__Shape = 2
	QFrame__WinPanel    QFrame__Shape = 3
	QFrame__HLine       QFrame__Shape = 4
	QFrame__VLine       QFrame__Shape = 5
	QFrame__StyledPanel QFrame__Shape = 6
)

type QFrame__Shadow int

const (
	QFrame__Plain  QFrame__Shadow = 16
	QFrame__Raised QFrame__Shadow = 32
	QFrame__Sunken QFrame__Shadow = 48
)

type QFrame__StyleMask int

const (
	QFrame__Shadow_Mask QFrame__StyleMask = 240
	QFrame__Shape_Mask  QFrame__StyleMask = 15
)

type QFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQFrame constructs a new QFrame object.
func NewQFrame(parent *QWidget) *QFrame {
	ret := newQFrame(QFrame_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFrame2 constructs a new QFrame object.
func NewQFrame2() *QFrame {
	ret := newQFrame(QFrame_new2())
	ret.isSubclass = true
	return ret
}

// NewQFrame3 constructs a new QFrame object.
func NewQFrame3(parent *QWidget, f WindowType) *QFrame {
	ret := newQFrame(QFrame_new3(parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

func (this *QFrame) MetaObject() *QMetaObject {
	return newQMetaObject(QFrame_MetaObject(this.h))
}

func (this *QFrame) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFrame_Metacast(this.h, param1_Cstring))
}

func QFrame_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFrame_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFrame) FrameStyle() int {
	return (int)(QFrame_FrameStyle(this.h))
}

func (this *QFrame) SetFrameStyle(frameStyle int) {
	QFrame_SetFrameStyle(this.h, (int)(frameStyle))
}

func (this *QFrame) FrameWidth() int {
	return (int)(QFrame_FrameWidth(this.h))
}

func (this *QFrame) SizeHint() *QSize {
	_goptr := newQSize(QFrame_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFrame) FrameShape() Shape {
	xxxxxxxxx
}

func (this *QFrame) SetFrameShape(frameShape Shape) {
	QFrame_SetFrameShape(this.h, frameShape)
}

func (this *QFrame) FrameShadow() Shadow {
	xxxxxxxxx
}

func (this *QFrame) SetFrameShadow(frameShadow Shadow) {
	QFrame_SetFrameShadow(this.h, frameShadow)
}

func (this *QFrame) LineWidth() int {
	return (int)(QFrame_LineWidth(this.h))
}

func (this *QFrame) SetLineWidth(lineWidth int) {
	QFrame_SetLineWidth(this.h, (int)(lineWidth))
}

func (this *QFrame) MidLineWidth() int {
	return (int)(QFrame_MidLineWidth(this.h))
}

func (this *QFrame) SetMidLineWidth(midLineWidth int) {
	QFrame_SetMidLineWidth(this.h, (int)(midLineWidth))
}

func (this *QFrame) FrameRect() *QRect {
	_goptr := newQRect(QFrame_FrameRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFrame) SetFrameRect(frameRect *QRect) {
	QFrame_SetFrameRect(this.h, frameRect.cPointer())
}

func QFrame_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFrame_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFrame_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFrame_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFrame) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFrame_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFrame) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_MetaObject
func miqt_exec_callback_QFrame_MetaObject(self QFrame, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFrame) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFrame_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFrame) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFrame_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFrame_Metacast
func miqt_exec_callback_QFrame_Metacast(self QFrame, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QFrame{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
