package svg

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QGraphicsSvgItem__ int

const (
	QGraphicsSvgItem__Type QGraphicsSvgItem__ = 13
)

type QGraphicsSvgItem struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsSvgItem constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem() *QGraphicsSvgItem {

	ret := newQGraphicsSvgItem(QGraphicsSvgItem_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSvgItem2 constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem2(fileName string) *QGraphicsSvgItem {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQGraphicsSvgItem(QGraphicsSvgItem_new2(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSvgItem3 constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem3(parentItem *qt.QGraphicsItem) *QGraphicsSvgItem {

	ret := newQGraphicsSvgItem(QGraphicsSvgItem_new3((*QGraphicsItem)(parentItem.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsSvgItem4 constructs a new QGraphicsSvgItem object.
func NewQGraphicsSvgItem4(fileName string, parentItem *qt.QGraphicsItem) *QGraphicsSvgItem {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQGraphicsSvgItem(QGraphicsSvgItem_new4(fileName_ms, (*QGraphicsItem)(parentItem.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsSvgItem) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QGraphicsSvgItem_MetaObject(this.h)))
}

func (this *QGraphicsSvgItem) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsSvgItem_Metacast(this.h, param1_Cstring))
}

func QGraphicsSvgItem_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsSvgItem_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsSvgItem) SetSharedRenderer(renderer *QSvgRenderer) {
	QGraphicsSvgItem_SetSharedRenderer(this.h, renderer.cPointer())
}

func (this *QGraphicsSvgItem) Renderer() *QSvgRenderer {
	return newQSvgRenderer(QGraphicsSvgItem_Renderer(this.h))
}

func (this *QGraphicsSvgItem) SetElementId(id string) {
	id_ms := struct_miqt_string{}
	id_ms.data = CString(id)
	id_ms.len = size_t(len(id))
	defer free(unsafe.Pointer(id_ms.data))
	QGraphicsSvgItem_SetElementId(this.h, id_ms)
}

func (this *QGraphicsSvgItem) ElementId() string {
	var _ms struct_miqt_string = QGraphicsSvgItem_ElementId(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsSvgItem) SetCachingEnabled(cachingEnabled bool) {
	QGraphicsSvgItem_SetCachingEnabled(this.h, (bool)(cachingEnabled))
}

func (this *QGraphicsSvgItem) IsCachingEnabled() bool {
	return (bool)(QGraphicsSvgItem_IsCachingEnabled(this.h))
}

func (this *QGraphicsSvgItem) SetMaximumCacheSize(size *qt.QSize) {
	QGraphicsSvgItem_SetMaximumCacheSize(this.h, (*QSize)(size.UnsafePointer()))
}

func (this *QGraphicsSvgItem) MaximumCacheSize() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QGraphicsSvgItem_MaximumCacheSize(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSvgItem) BoundingRect() *qt.QRectF {
	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QGraphicsSvgItem_BoundingRect(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsSvgItem) Paint(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget) {
	QGraphicsSvgItem_Paint(this.h, (*QPainter)(painter.UnsafePointer()), (*QStyleOptionGraphicsItem)(option.UnsafePointer()), (*QWidget)(widget.UnsafePointer()))
}

func (this *QGraphicsSvgItem) Type() int {
	return (int)(QGraphicsSvgItem_Type(this.h))
}

func QGraphicsSvgItem_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsSvgItem_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsSvgItem_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsSvgItem_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsSvgItem) callVirtualBase_BoundingRect() *qt.QRectF {

	_goptr := qt.UnsafeNewQRectF(unsafe.Pointer(QGraphicsSvgItem_virtualbase_BoundingRect(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsSvgItem) OnBoundingRect(slot func(super func() *qt.QRectF) *qt.QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsSvgItem_override_virtual_BoundingRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsSvgItem_BoundingRect
func miqt_exec_callback_QGraphicsSvgItem_BoundingRect(self QGraphicsSvgItem, cb intptr_t) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QRectF) *qt.QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_BoundingRect)

	return (*QRectF)(virtualReturn.UnsafePointer())

}

func (this *QGraphicsSvgItem) callVirtualBase_Paint(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget) {

	QGraphicsSvgItem_virtualbase_Paint(unsafe.Pointer(this.h), (*QPainter)(painter.UnsafePointer()), (*QStyleOptionGraphicsItem)(option.UnsafePointer()), (*QWidget)(widget.UnsafePointer()))

}
func (this *QGraphicsSvgItem) OnPaint(slot func(super func(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget), painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsSvgItem_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsSvgItem_Paint
func miqt_exec_callback_QGraphicsSvgItem_Paint(self QGraphicsSvgItem, cb intptr_t, painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget), painter *qt.QPainter, option *qt.QStyleOptionGraphicsItem, widget *qt.QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	slotval2 := qt.UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option))

	slotval3 := qt.UnsafeNewQWidget(unsafe.Pointer(widget))

	gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Paint, slotval1, slotval2, slotval3)

}

func (this *QGraphicsSvgItem) callVirtualBase_Type() int {

	return (int)(QGraphicsSvgItem_virtualbase_Type(unsafe.Pointer(this.h)))

}
func (this *QGraphicsSvgItem) OnType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsSvgItem_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsSvgItem_Type
func miqt_exec_callback_QGraphicsSvgItem_Type(self QGraphicsSvgItem, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Type)

	return (int)(virtualReturn)

}

func (this *QGraphicsSvgItem) callVirtualBase_Event(ev *qt.QEvent) bool {

	return (bool)(QGraphicsSvgItem_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(ev.UnsafePointer())))

}
func (this *QGraphicsSvgItem) OnEvent(slot func(super func(ev *qt.QEvent) bool, ev *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsSvgItem_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsSvgItem_Event
func miqt_exec_callback_QGraphicsSvgItem_Event(self QGraphicsSvgItem, cb intptr_t, ev *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *qt.QEvent) bool, ev *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(ev))

	virtualReturn := gofunc((&QGraphicsSvgItem{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}
