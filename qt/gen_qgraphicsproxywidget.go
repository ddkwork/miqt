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

	ret := newQGraphicsProxyWidget(QGraphicsProxyWidget_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsProxyWidget2 constructs a new QGraphicsProxyWidget object.
func NewQGraphicsProxyWidget2(parent *QGraphicsItem) *QGraphicsProxyWidget {

	ret := newQGraphicsProxyWidget(QGraphicsProxyWidget_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsProxyWidget3 constructs a new QGraphicsProxyWidget object.
func NewQGraphicsProxyWidget3(parent *QGraphicsItem, wFlags WindowType) *QGraphicsProxyWidget {

	ret := newQGraphicsProxyWidget(QGraphicsProxyWidget_new3(parent.cPointer(), (int)(wFlags)))
	ret.isSubclass = true
	return ret
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

func (this *QGraphicsProxyWidget) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsProxyWidget_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsProxyWidget) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_SetGeometry
func miqt_exec_callback_QGraphicsProxyWidget_SetGeometry(self QGraphicsProxyWidget, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {

	QGraphicsProxyWidget_virtualbase_Paint(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), widget.cPointer())

}
func (this *QGraphicsProxyWidget) OnPaint(slot func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_Paint
func miqt_exec_callback_QGraphicsProxyWidget_Paint(self QGraphicsProxyWidget, cb intptr_t, painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionGraphicsItem(option)

	slotval3 := newQWidget(widget)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_Paint, slotval1, slotval2, slotval3)

}

func (this *QGraphicsProxyWidget) callVirtualBase_Type() int {

	return (int)(QGraphicsProxyWidget_virtualbase_Type(unsafe.Pointer(this.h)))

}
func (this *QGraphicsProxyWidget) OnType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_Type
func miqt_exec_callback_QGraphicsProxyWidget_Type(self QGraphicsProxyWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_Type)

	return (int)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_ItemChange(change GraphicsItemChange, value *QVariant) *QVariant {

	_goptr := newQVariant(QGraphicsProxyWidget_virtualbase_ItemChange(unsafe.Pointer(this.h), change, value.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsProxyWidget) OnItemChange(slot func(super func(change GraphicsItemChange, value *QVariant) *QVariant, change GraphicsItemChange, value *QVariant) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_ItemChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_ItemChange
func miqt_exec_callback_QGraphicsProxyWidget_ItemChange(self QGraphicsProxyWidget, cb intptr_t, change GraphicsItemChange, value *QVariant) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change GraphicsItemChange, value *QVariant) *QVariant, change GraphicsItemChange, value *QVariant) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQVariant(value)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_ItemChange, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsProxyWidget) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGraphicsProxyWidget_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsProxyWidget) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_Event
func miqt_exec_callback_QGraphicsProxyWidget_Event(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QGraphicsProxyWidget_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QGraphicsProxyWidget) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_EventFilter
func miqt_exec_callback_QGraphicsProxyWidget_EventFilter(self QGraphicsProxyWidget, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_ShowEvent(event *QShowEvent) {

	QGraphicsProxyWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_ShowEvent
func miqt_exec_callback_QGraphicsProxyWidget_ShowEvent(self QGraphicsProxyWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_HideEvent(event *QHideEvent) {

	QGraphicsProxyWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_HideEvent
func miqt_exec_callback_QGraphicsProxyWidget_HideEvent(self QGraphicsProxyWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_ContextMenuEvent(event *QGraphicsSceneContextMenuEvent) {

	QGraphicsProxyWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnContextMenuEvent(slot func(super func(event *QGraphicsSceneContextMenuEvent), event *QGraphicsSceneContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_ContextMenuEvent
func miqt_exec_callback_QGraphicsProxyWidget_ContextMenuEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneContextMenuEvent), event *QGraphicsSceneContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneContextMenuEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_DragEnterEvent(event *QGraphicsSceneDragDropEvent) {

	QGraphicsProxyWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnDragEnterEvent(slot func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_DragEnterEvent
func miqt_exec_callback_QGraphicsProxyWidget_DragEnterEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneDragDropEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_DragLeaveEvent(event *QGraphicsSceneDragDropEvent) {

	QGraphicsProxyWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnDragLeaveEvent(slot func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_DragLeaveEvent
func miqt_exec_callback_QGraphicsProxyWidget_DragLeaveEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneDragDropEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_DragMoveEvent(event *QGraphicsSceneDragDropEvent) {

	QGraphicsProxyWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnDragMoveEvent(slot func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_DragMoveEvent
func miqt_exec_callback_QGraphicsProxyWidget_DragMoveEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneDragDropEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_DropEvent(event *QGraphicsSceneDragDropEvent) {

	QGraphicsProxyWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnDropEvent(slot func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_DropEvent
func miqt_exec_callback_QGraphicsProxyWidget_DropEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneDragDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneDragDropEvent), event *QGraphicsSceneDragDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneDragDropEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_HoverEnterEvent(event *QGraphicsSceneHoverEvent) {

	QGraphicsProxyWidget_virtualbase_HoverEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnHoverEnterEvent(slot func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_HoverEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_HoverEnterEvent
func miqt_exec_callback_QGraphicsProxyWidget_HoverEnterEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneHoverEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_HoverEnterEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_HoverLeaveEvent(event *QGraphicsSceneHoverEvent) {

	QGraphicsProxyWidget_virtualbase_HoverLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnHoverLeaveEvent(slot func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_HoverLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_HoverLeaveEvent
func miqt_exec_callback_QGraphicsProxyWidget_HoverLeaveEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneHoverEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_HoverLeaveEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_HoverMoveEvent(event *QGraphicsSceneHoverEvent) {

	QGraphicsProxyWidget_virtualbase_HoverMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnHoverMoveEvent(slot func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_HoverMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_HoverMoveEvent
func miqt_exec_callback_QGraphicsProxyWidget_HoverMoveEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneHoverEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_HoverMoveEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_GrabMouseEvent(event *QEvent) {

	QGraphicsProxyWidget_virtualbase_GrabMouseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnGrabMouseEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_GrabMouseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_GrabMouseEvent
func miqt_exec_callback_QGraphicsProxyWidget_GrabMouseEvent(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_GrabMouseEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_UngrabMouseEvent(event *QEvent) {

	QGraphicsProxyWidget_virtualbase_UngrabMouseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnUngrabMouseEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_UngrabMouseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_UngrabMouseEvent
func miqt_exec_callback_QGraphicsProxyWidget_UngrabMouseEvent(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_UngrabMouseEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_MouseMoveEvent(event *QGraphicsSceneMouseEvent) {

	QGraphicsProxyWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnMouseMoveEvent(slot func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_MouseMoveEvent
func miqt_exec_callback_QGraphicsProxyWidget_MouseMoveEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneMouseEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_MousePressEvent(event *QGraphicsSceneMouseEvent) {

	QGraphicsProxyWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnMousePressEvent(slot func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_MousePressEvent
func miqt_exec_callback_QGraphicsProxyWidget_MousePressEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneMouseEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_MouseReleaseEvent(event *QGraphicsSceneMouseEvent) {

	QGraphicsProxyWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnMouseReleaseEvent(slot func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_MouseReleaseEvent
func miqt_exec_callback_QGraphicsProxyWidget_MouseReleaseEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneMouseEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_MouseDoubleClickEvent(event *QGraphicsSceneMouseEvent) {

	QGraphicsProxyWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnMouseDoubleClickEvent(slot func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_MouseDoubleClickEvent
func miqt_exec_callback_QGraphicsProxyWidget_MouseDoubleClickEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneMouseEvent), event *QGraphicsSceneMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneMouseEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_WheelEvent(event *QGraphicsSceneWheelEvent) {

	QGraphicsProxyWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnWheelEvent(slot func(super func(event *QGraphicsSceneWheelEvent), event *QGraphicsSceneWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_WheelEvent
func miqt_exec_callback_QGraphicsProxyWidget_WheelEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneWheelEvent), event *QGraphicsSceneWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneWheelEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QGraphicsProxyWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_KeyPressEvent
func miqt_exec_callback_QGraphicsProxyWidget_KeyPressEvent(self QGraphicsProxyWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QGraphicsProxyWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_KeyReleaseEvent
func miqt_exec_callback_QGraphicsProxyWidget_KeyReleaseEvent(self QGraphicsProxyWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QGraphicsProxyWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_FocusInEvent
func miqt_exec_callback_QGraphicsProxyWidget_FocusInEvent(self QGraphicsProxyWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QGraphicsProxyWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_FocusOutEvent
func miqt_exec_callback_QGraphicsProxyWidget_FocusOutEvent(self QGraphicsProxyWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QGraphicsProxyWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QGraphicsProxyWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_FocusNextPrevChild
func miqt_exec_callback_QGraphicsProxyWidget_FocusNextPrevChild(self QGraphicsProxyWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_InputMethodQuery(query InputMethodQuery) *QVariant {

	_goptr := newQVariant(QGraphicsProxyWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(query)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsProxyWidget) OnInputMethodQuery(slot func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_InputMethodQuery
func miqt_exec_callback_QGraphicsProxyWidget_InputMethodQuery(self QGraphicsProxyWidget, cb intptr_t, query int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(query InputMethodQuery) *QVariant, query InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(query)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QGraphicsProxyWidget) callVirtualBase_InputMethodEvent(event *QInputMethodEvent) {

	QGraphicsProxyWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnInputMethodEvent(slot func(super func(event *QInputMethodEvent), event *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_InputMethodEvent
func miqt_exec_callback_QGraphicsProxyWidget_InputMethodEvent(self QGraphicsProxyWidget, cb intptr_t, event *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QInputMethodEvent), event *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {

	_goptr := newQSizeF(QGraphicsProxyWidget_virtualbase_SizeHint(unsafe.Pointer(this.h), (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsProxyWidget) OnSizeHint(slot func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_SizeHint
func miqt_exec_callback_QGraphicsProxyWidget_SizeHint(self QGraphicsProxyWidget, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (SizeHint)(which)

	slotval2 := newQSizeF(constraint)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsProxyWidget) callVirtualBase_ResizeEvent(event *QGraphicsSceneResizeEvent) {

	QGraphicsProxyWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnResizeEvent(slot func(super func(event *QGraphicsSceneResizeEvent), event *QGraphicsSceneResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_ResizeEvent
func miqt_exec_callback_QGraphicsProxyWidget_ResizeEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneResizeEvent), event *QGraphicsSceneResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneResizeEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsProxyWidget_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsProxyWidget) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_GetContentsMargins
func miqt_exec_callback_QGraphicsProxyWidget_GetContentsMargins(self QGraphicsProxyWidget, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsProxyWidget) callVirtualBase_PaintWindowFrame(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {

	QGraphicsProxyWidget_virtualbase_PaintWindowFrame(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), widget.cPointer())

}
func (this *QGraphicsProxyWidget) OnPaintWindowFrame(slot func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_PaintWindowFrame(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_PaintWindowFrame
func miqt_exec_callback_QGraphicsProxyWidget_PaintWindowFrame(self QGraphicsProxyWidget, cb intptr_t, painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionGraphicsItem(option)

	slotval3 := newQWidget(widget)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_PaintWindowFrame, slotval1, slotval2, slotval3)

}

func (this *QGraphicsProxyWidget) callVirtualBase_BoundingRect() *QRectF {

	_goptr := newQRectF(QGraphicsProxyWidget_virtualbase_BoundingRect(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsProxyWidget) OnBoundingRect(slot func(super func() *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_BoundingRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_BoundingRect
func miqt_exec_callback_QGraphicsProxyWidget_BoundingRect(self QGraphicsProxyWidget, cb intptr_t) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_BoundingRect)

	return virtualReturn.cPointer()

}

func (this *QGraphicsProxyWidget) callVirtualBase_Shape() *QPainterPath {

	_goptr := newQPainterPath(QGraphicsProxyWidget_virtualbase_Shape(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsProxyWidget) OnShape(slot func(super func() *QPainterPath) *QPainterPath) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_Shape(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_Shape
func miqt_exec_callback_QGraphicsProxyWidget_Shape(self QGraphicsProxyWidget, cb intptr_t) *QPainterPath {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainterPath) *QPainterPath)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_Shape)

	return virtualReturn.cPointer()

}

func (this *QGraphicsProxyWidget) callVirtualBase_InitStyleOption(option *QStyleOption) {

	QGraphicsProxyWidget_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QGraphicsProxyWidget) OnInitStyleOption(slot func(super func(option *QStyleOption), option *QStyleOption)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_InitStyleOption
func miqt_exec_callback_QGraphicsProxyWidget_InitStyleOption(self QGraphicsProxyWidget, cb intptr_t, option *QStyleOption) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOption), option *QStyleOption))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOption(option)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_UpdateGeometry() {

	QGraphicsProxyWidget_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsProxyWidget) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_UpdateGeometry
func miqt_exec_callback_QGraphicsProxyWidget_UpdateGeometry(self QGraphicsProxyWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsProxyWidget) callVirtualBase_PropertyChange(propertyName string, value *QVariant) *QVariant {
	propertyName_ms := struct_miqt_string{}
	propertyName_ms.data = CString(propertyName)
	propertyName_ms.len = size_t(len(propertyName))
	defer free(unsafe.Pointer(propertyName_ms.data))

	_goptr := newQVariant(QGraphicsProxyWidget_virtualbase_PropertyChange(unsafe.Pointer(this.h), propertyName_ms, value.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsProxyWidget) OnPropertyChange(slot func(super func(propertyName string, value *QVariant) *QVariant, propertyName string, value *QVariant) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_PropertyChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_PropertyChange
func miqt_exec_callback_QGraphicsProxyWidget_PropertyChange(self QGraphicsProxyWidget, cb intptr_t, propertyName struct_miqt_string, value *QVariant) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(propertyName string, value *QVariant) *QVariant, propertyName string, value *QVariant) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var propertyName_ms struct_miqt_string = propertyName
	propertyName_ret := GoStringN(propertyName_ms.data, int(int64(propertyName_ms.len)))
	free(unsafe.Pointer(propertyName_ms.data))
	slotval1 := propertyName_ret
	slotval2 := newQVariant(value)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_PropertyChange, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsProxyWidget) callVirtualBase_SceneEvent(event *QEvent) bool {

	return (bool)(QGraphicsProxyWidget_virtualbase_SceneEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsProxyWidget) OnSceneEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_SceneEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_SceneEvent
func miqt_exec_callback_QGraphicsProxyWidget_SceneEvent(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_SceneEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_WindowFrameEvent(e *QEvent) bool {

	return (bool)(QGraphicsProxyWidget_virtualbase_WindowFrameEvent(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QGraphicsProxyWidget) OnWindowFrameEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_WindowFrameEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_WindowFrameEvent
func miqt_exec_callback_QGraphicsProxyWidget_WindowFrameEvent(self QGraphicsProxyWidget, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_WindowFrameEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_WindowFrameSectionAt(pos *QPointF) WindowFrameSection {

	return (WindowFrameSection)(QGraphicsProxyWidget_virtualbase_WindowFrameSectionAt(unsafe.Pointer(this.h), pos.cPointer()))

}
func (this *QGraphicsProxyWidget) OnWindowFrameSectionAt(slot func(super func(pos *QPointF) WindowFrameSection, pos *QPointF) WindowFrameSection) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_WindowFrameSectionAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_WindowFrameSectionAt
func miqt_exec_callback_QGraphicsProxyWidget_WindowFrameSectionAt(self QGraphicsProxyWidget, cb intptr_t, pos *QPointF) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos *QPointF) WindowFrameSection, pos *QPointF) WindowFrameSection)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(pos)

	virtualReturn := gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_WindowFrameSectionAt, slotval1)

	return (int)(virtualReturn)

}

func (this *QGraphicsProxyWidget) callVirtualBase_ChangeEvent(event *QEvent) {

	QGraphicsProxyWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_ChangeEvent
func miqt_exec_callback_QGraphicsProxyWidget_ChangeEvent(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QGraphicsProxyWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_CloseEvent
func miqt_exec_callback_QGraphicsProxyWidget_CloseEvent(self QGraphicsProxyWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_MoveEvent(event *QGraphicsSceneMoveEvent) {

	QGraphicsProxyWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnMoveEvent(slot func(super func(event *QGraphicsSceneMoveEvent), event *QGraphicsSceneMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_MoveEvent
func miqt_exec_callback_QGraphicsProxyWidget_MoveEvent(self QGraphicsProxyWidget, cb intptr_t, event *QGraphicsSceneMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneMoveEvent), event *QGraphicsSceneMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneMoveEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_PolishEvent() {

	QGraphicsProxyWidget_virtualbase_PolishEvent(unsafe.Pointer(this.h))

}
func (this *QGraphicsProxyWidget) OnPolishEvent(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_PolishEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_PolishEvent
func miqt_exec_callback_QGraphicsProxyWidget_PolishEvent(self QGraphicsProxyWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_PolishEvent)

}

func (this *QGraphicsProxyWidget) callVirtualBase_GrabKeyboardEvent(event *QEvent) {

	QGraphicsProxyWidget_virtualbase_GrabKeyboardEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnGrabKeyboardEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_GrabKeyboardEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_GrabKeyboardEvent
func miqt_exec_callback_QGraphicsProxyWidget_GrabKeyboardEvent(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_GrabKeyboardEvent, slotval1)

}

func (this *QGraphicsProxyWidget) callVirtualBase_UngrabKeyboardEvent(event *QEvent) {

	QGraphicsProxyWidget_virtualbase_UngrabKeyboardEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsProxyWidget) OnUngrabKeyboardEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsProxyWidget_override_virtual_UngrabKeyboardEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsProxyWidget_UngrabKeyboardEvent
func miqt_exec_callback_QGraphicsProxyWidget_UngrabKeyboardEvent(self QGraphicsProxyWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsProxyWidget{h: self}).callVirtualBase_UngrabKeyboardEvent, slotval1)

}
