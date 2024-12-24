package qt

import (
	"unsafe"
)

type QGraphicsProxyWidget__ int

const (
	QGraphicsProxyWidget__Type QGraphicsProxyWidget__ = 12
)

type QGraphicsProxyWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsProxyWidget constructs a new QGraphicsProxyWidget object.
func NewQGraphicsProxyWidget() *QGraphicsProxyWidget {
	g := newQGraphicsProxyWidget(QGraphicsProxyWidget_new())
	g.isSubclass = true
	return g
}

// NewQGraphicsProxyWidget2 constructs a new QGraphicsProxyWidget object.
func NewQGraphicsProxyWidget2(parent *QGraphicsItem) *QGraphicsProxyWidget {
	g := newQGraphicsProxyWidget(QGraphicsProxyWidget_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQGraphicsProxyWidget3 constructs a new QGraphicsProxyWidget object.
func NewQGraphicsProxyWidget3(parent *QGraphicsItem, wFlags WindowType) *QGraphicsProxyWidget {
	g := newQGraphicsProxyWidget(QGraphicsProxyWidget_new3(parent.cPointer(), (int)(wFlags)))
	g.isSubclass = true
	return g
}

func (this *QGraphicsProxyWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsProxyWidget_MetaObject(this.h))
}

func (this *QGraphicsProxyWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsProxyWidget_Metacast(this.h, param1_Cstring))
}

func QGraphicsProxyWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsProxyWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsProxyWidget) SetWidget(widget *QWidget) {
	QGraphicsProxyWidget_SetWidget(this.h, widget.cPointer())
}

func (this *QGraphicsProxyWidget) Widget() *QWidget {
	return newQWidget(QGraphicsProxyWidget_Widget(this.h))
}

func (this *QGraphicsProxyWidget) SubWidgetRect(widget *QWidget) *QRectF {
	_goptr := newQRectF(QGraphicsProxyWidget_SubWidgetRect(this.h, widget.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsProxyWidget) SetGeometry(rect *QRectF) {
	QGraphicsProxyWidget_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsProxyWidget) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsProxyWidget_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsProxyWidget) Type() int {
	return (int)(QGraphicsProxyWidget_Type(this.h))
}

func (this *QGraphicsProxyWidget) CreateProxyForChildWidget(child *QWidget) *QGraphicsProxyWidget {
	return newQGraphicsProxyWidget(QGraphicsProxyWidget_CreateProxyForChildWidget(this.h, child.cPointer()))
}

func QGraphicsProxyWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsProxyWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsProxyWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsProxyWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsProxyWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsProxyWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGraphicsProxyWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_MetaObject
func miqt_exec_callback_QGraphicsProxyWidget_MetaObject(self QGraphicsProxyWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGraphicsProxyWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGraphicsProxyWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGraphicsProxyWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_Metacast
func miqt_exec_callback_QGraphicsProxyWidget_Metacast(self QGraphicsProxyWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
