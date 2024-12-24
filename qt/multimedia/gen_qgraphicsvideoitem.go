package multimedia

import (
	"unsafe"

	"github.com/mappu/miqt/qt"
)

type QGraphicsVideoItem__ int

const (
	QGraphicsVideoItem__Type QGraphicsVideoItem__ = 14
)

type QGraphicsVideoItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsVideoItem constructs a new QGraphicsVideoItem object.
func NewQGraphicsVideoItem() *QGraphicsVideoItem {
	ret := newQGraphicsVideoItem(QGraphicsVideoItem_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsVideoItem2 constructs a new QGraphicsVideoItem object.
func NewQGraphicsVideoItem2(parent *qt.QGraphicsItem) *QGraphicsVideoItem {
	ret := newQGraphicsVideoItem(QGraphicsVideoItem_new2((*QGraphicsItem)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsVideoItem) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QGraphicsVideoItem_MetaObject(this.h)))
}

func (this *QGraphicsVideoItem) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsVideoItem_Metacast(this.h, param1_Cstring))
}

func QGraphicsVideoItem_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsVideoItem_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsVideoItem) VideoSink() *QVideoSink {
	return newQVideoSink(QGraphicsVideoItem_VideoSink(this.h))
}

func (this *QGraphicsVideoItem) AspectRatioMode() qt.AspectRatioMode {
	return (qt.AspectRatioMode)(QGraphicsVideoItem_AspectRatioMode(this.h))
}

func (this *QGraphicsVideoItem) SetAspectRatioMode(mode qt.AspectRatioMode) {
	QGraphicsVideoItem_SetAspectRatioMode(this.h, (int)(mode))
}

func (this *QGraphicsVideoItem) Offset() *qt.QPointF {
	_goptr := qt.UnsafeNewQPointF(unsafe.Pointer(QGraphicsVideoItem_Offset(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsVideoItem) SetOffset(offset *qt.QPointF) {
	QGraphicsVideoItem_SetOffset(this.h, (*QPointF)(offset.UnsafePointer()))
}

func (this *QGraphicsVideoItem) Size() *qt.QSizeF {
	_goptr := qt.UnsafeNewQSizeF(unsafe.Pointer(QGraphicsVideoItem_Size(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsVideoItem) SetSize(size *qt.QSizeF) {
	QGraphicsVideoItem_SetSize(this.h, (*QSizeF)(size.UnsafePointer()))
}

func (this *QGraphicsVideoItem) NativeSize() *qt.QSizeF {
	_goptr := qt.UnsafeNewQSizeF(unsafe.Pointer(QGraphicsVideoItem_NativeSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsVideoItem) BoundingRect() *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QGraphicsVideoItem_BoundingRect(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsVideoItem) Paint(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget) {
	QGraphicsVideoItem_Paint(this.h, (*QPainter)(painter.UnsafePointer()), (*QStyleOptionGraphicsItem)(option.UnsafePointer()), (*QWidget)(widget.UnsafePointer()))
}

func (this *QGraphicsVideoItem) Type() int {
	return (int)(QGraphicsVideoItem_Type(this.h))
}

func (this *QGraphicsVideoItem) NativeSizeChanged(size *qt.QSizeF) {
	QGraphicsVideoItem_NativeSizeChanged(this.h, (*QSizeF)(size.UnsafePointer()))
}

func (this *QGraphicsVideoItem) OnNativeSizeChanged(slot func(size *qt.QSizeF)) {
	QGraphicsVideoItem_connect_NativeSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsVideoItem_NativeSizeChanged
func miqt_exec_callback_QGraphicsVideoItem_NativeSizeChanged(cb intptr_t, size *QSizeF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(size *qt.QSizeF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQSizeF(unsafe.Pointer(size))

	gofunc(slotval1)
}

func QGraphicsVideoItem_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsVideoItem_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsVideoItem_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsVideoItem_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsVideoItem) callVirtualBase_MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QGraphicsVideoItem_virtualbase_MetaObject(unsafe.Pointer(this.h))))
}

func (this *QGraphicsVideoItem) OnMetaObject(slot func(super func() *qt.QMetaObject) *qt.QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsVideoItem_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsVideoItem_MetaObject
func miqt_exec_callback_QGraphicsVideoItem_MetaObject(self QGraphicsVideoItem, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QMetaObject) *qt.QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsVideoItem{h: self}).callVirtualBase_MetaObject)

	return (*QMetaObject)(virtualReturn.UnsafePointer())
}

func (this *QGraphicsVideoItem) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsVideoItem_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsVideoItem) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsVideoItem_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsVideoItem_Metacast
func miqt_exec_callback_QGraphicsVideoItem_Metacast(self QGraphicsVideoItem, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QGraphicsVideoItem{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
