package qt

import (
	"unsafe"
)

type QAbstractTextDocumentLayout struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractTextDocumentLayout constructs a new QAbstractTextDocumentLayout object.
func NewQAbstractTextDocumentLayout(doc *QTextDocument) *QAbstractTextDocumentLayout {

	ret := newQAbstractTextDocumentLayout(QAbstractTextDocumentLayout_new(doc.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractTextDocumentLayout) MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractTextDocumentLayout_MetaObject(this.h))
}

func (this *QAbstractTextDocumentLayout) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractTextDocumentLayout_Metacast(this.h, param1_Cstring))
}

func QAbstractTextDocumentLayout_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractTextDocumentLayout_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTextDocumentLayout) Draw(painter *QPainter, context *PaintContext) {
	QAbstractTextDocumentLayout_Draw(this.h, painter.cPointer(), context)
}

func (this *QAbstractTextDocumentLayout) HitTest(point *QPointF, accuracy HitTestAccuracy) int {
	return (int)(QAbstractTextDocumentLayout_HitTest(this.h, point.cPointer(), (int)(accuracy)))
}

func (this *QAbstractTextDocumentLayout) AnchorAt(pos *QPointF) string {
	var _ms struct_miqt_string = QAbstractTextDocumentLayout_AnchorAt(this.h, pos.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTextDocumentLayout) ImageAt(pos *QPointF) string {
	var _ms struct_miqt_string = QAbstractTextDocumentLayout_ImageAt(this.h, pos.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTextDocumentLayout) FormatAt(pos *QPointF) *QTextFormat {
	_goptr := newQTextFormat(QAbstractTextDocumentLayout_FormatAt(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTextDocumentLayout) BlockWithMarkerAt(pos *QPointF) *QTextBlock {
	_goptr := newQTextBlock(QAbstractTextDocumentLayout_BlockWithMarkerAt(this.h, pos.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTextDocumentLayout) PageCount() int {
	return (int)(QAbstractTextDocumentLayout_PageCount(this.h))
}

func (this *QAbstractTextDocumentLayout) DocumentSize() *QSizeF {
	_goptr := newQSizeF(QAbstractTextDocumentLayout_DocumentSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTextDocumentLayout) FrameBoundingRect(frame *QTextFrame) *QRectF {
	_goptr := newQRectF(QAbstractTextDocumentLayout_FrameBoundingRect(this.h, frame.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTextDocumentLayout) BlockBoundingRect(block *QTextBlock) *QRectF {
	_goptr := newQRectF(QAbstractTextDocumentLayout_BlockBoundingRect(this.h, block.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAbstractTextDocumentLayout) SetPaintDevice(device *QPaintDevice) {
	QAbstractTextDocumentLayout_SetPaintDevice(this.h, device.cPointer())
}

func (this *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice {
	return newQPaintDevice(QAbstractTextDocumentLayout_PaintDevice(this.h))
}

func (this *QAbstractTextDocumentLayout) Document() *QTextDocument {
	return newQTextDocument(QAbstractTextDocumentLayout_Document(this.h))
}

func (this *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component *QObject) {
	QAbstractTextDocumentLayout_RegisterHandler(this.h, (int)(objectType), component.cPointer())
}

func (this *QAbstractTextDocumentLayout) UnregisterHandler(objectType int) {
	QAbstractTextDocumentLayout_UnregisterHandler(this.h, (int)(objectType))
}

func (this *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface {
	return newQTextObjectInterface(QAbstractTextDocumentLayout_HandlerForObject(this.h, (int)(objectType)))
}

func (this *QAbstractTextDocumentLayout) Update() {
	QAbstractTextDocumentLayout_Update(this.h)
}
func (this *QAbstractTextDocumentLayout) OnUpdate(slot func()) {
	QAbstractTextDocumentLayout_connect_Update(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_Update
func miqt_exec_callback_QAbstractTextDocumentLayout_Update(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAbstractTextDocumentLayout) UpdateBlock(block *QTextBlock) {
	QAbstractTextDocumentLayout_UpdateBlock(this.h, block.cPointer())
}
func (this *QAbstractTextDocumentLayout) OnUpdateBlock(slot func(block *QTextBlock)) {
	QAbstractTextDocumentLayout_connect_UpdateBlock(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_UpdateBlock
func miqt_exec_callback_QAbstractTextDocumentLayout_UpdateBlock(cb intptr_t, block *QTextBlock) {
	gofunc, ok := cgo.Handle(cb).Value().(func(block *QTextBlock))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextBlock(block)

	gofunc(slotval1)
}

func (this *QAbstractTextDocumentLayout) DocumentSizeChanged(newSize *QSizeF) {
	QAbstractTextDocumentLayout_DocumentSizeChanged(this.h, newSize.cPointer())
}
func (this *QAbstractTextDocumentLayout) OnDocumentSizeChanged(slot func(newSize *QSizeF)) {
	QAbstractTextDocumentLayout_connect_DocumentSizeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_DocumentSizeChanged
func miqt_exec_callback_QAbstractTextDocumentLayout_DocumentSizeChanged(cb intptr_t, newSize *QSizeF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newSize *QSizeF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQSizeF(newSize)

	gofunc(slotval1)
}

func (this *QAbstractTextDocumentLayout) PageCountChanged(newPages int) {
	QAbstractTextDocumentLayout_PageCountChanged(this.h, (int)(newPages))
}
func (this *QAbstractTextDocumentLayout) OnPageCountChanged(slot func(newPages int)) {
	QAbstractTextDocumentLayout_connect_PageCountChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_PageCountChanged
func miqt_exec_callback_QAbstractTextDocumentLayout_PageCountChanged(cb intptr_t, newPages int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newPages int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(newPages)

	gofunc(slotval1)
}

func QAbstractTextDocumentLayout_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractTextDocumentLayout_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractTextDocumentLayout_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractTextDocumentLayout_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractTextDocumentLayout) UnregisterHandler2(objectType int, component *QObject) {
	QAbstractTextDocumentLayout_UnregisterHandler2(this.h, (int)(objectType), component.cPointer())
}

func (this *QAbstractTextDocumentLayout) Update1(param1 *QRectF) {
	QAbstractTextDocumentLayout_Update1(this.h, param1.cPointer())
}
func (this *QAbstractTextDocumentLayout) OnUpdate1(slot func(param1 *QRectF)) {
	QAbstractTextDocumentLayout_connect_Update1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_Update1
func miqt_exec_callback_QAbstractTextDocumentLayout_Update1(cb intptr_t, param1 *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(param1)

	gofunc(slotval1)
}

func (this *QAbstractTextDocumentLayout) OnDraw(slot func(painter *QPainter, context *PaintContext)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_Draw(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_Draw
func miqt_exec_callback_QAbstractTextDocumentLayout_Draw(self QAbstractTextDocumentLayout, cb intptr_t, painter *QPainter, context *PaintContext) {
	gofunc, ok := cgo.Handle(cb).Value().(func(painter *QPainter, context *PaintContext))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	xxxxxxxxx

	gofunc(slotval1, slotval2)

}
func (this *QAbstractTextDocumentLayout) OnHitTest(slot func(point *QPointF, accuracy HitTestAccuracy) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_HitTest(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_HitTest
func miqt_exec_callback_QAbstractTextDocumentLayout_HitTest(self QAbstractTextDocumentLayout, cb intptr_t, point *QPointF, accuracy int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(point *QPointF, accuracy HitTestAccuracy) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(point)

	slotval2 := (HitTestAccuracy)(accuracy)

	virtualReturn := gofunc(slotval1, slotval2)

	return (int)(virtualReturn)

}
func (this *QAbstractTextDocumentLayout) OnPageCount(slot func() int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_PageCount(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_PageCount
func miqt_exec_callback_QAbstractTextDocumentLayout_PageCount(self QAbstractTextDocumentLayout, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func() int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return (int)(virtualReturn)

}
func (this *QAbstractTextDocumentLayout) OnDocumentSize(slot func() *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_DocumentSize(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_DocumentSize
func miqt_exec_callback_QAbstractTextDocumentLayout_DocumentSize(self QAbstractTextDocumentLayout, cb intptr_t) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func() *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc()

	return virtualReturn.cPointer()

}
func (this *QAbstractTextDocumentLayout) OnFrameBoundingRect(slot func(frame *QTextFrame) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_FrameBoundingRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_FrameBoundingRect
func miqt_exec_callback_QAbstractTextDocumentLayout_FrameBoundingRect(self QAbstractTextDocumentLayout, cb intptr_t, frame *QTextFrame) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(frame *QTextFrame) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextFrame(frame)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}
func (this *QAbstractTextDocumentLayout) OnBlockBoundingRect(slot func(block *QTextBlock) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_BlockBoundingRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_BlockBoundingRect
func miqt_exec_callback_QAbstractTextDocumentLayout_BlockBoundingRect(self QAbstractTextDocumentLayout, cb intptr_t, block *QTextBlock) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(block *QTextBlock) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextBlock(block)

	virtualReturn := gofunc(slotval1)

	return virtualReturn.cPointer()

}
func (this *QAbstractTextDocumentLayout) OnDocumentChanged(slot func(from int, charsRemoved int, charsAdded int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_DocumentChanged(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_DocumentChanged
func miqt_exec_callback_QAbstractTextDocumentLayout_DocumentChanged(self QAbstractTextDocumentLayout, cb intptr_t, from int, charsRemoved int, charsAdded int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(from int, charsRemoved int, charsAdded int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(from)

	slotval2 := (int)(charsRemoved)

	slotval3 := (int)(charsAdded)

	gofunc(slotval1, slotval2, slotval3)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_ResizeInlineObject(item QTextInlineObject, posInDocument int, format *QTextFormat) {

	QAbstractTextDocumentLayout_virtualbase_ResizeInlineObject(unsafe.Pointer(this.h), item.cPointer(), (int)(posInDocument), format.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnResizeInlineObject(slot func(super func(item QTextInlineObject, posInDocument int, format *QTextFormat), item QTextInlineObject, posInDocument int, format *QTextFormat)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_ResizeInlineObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_ResizeInlineObject
func miqt_exec_callback_QAbstractTextDocumentLayout_ResizeInlineObject(self QAbstractTextDocumentLayout, cb intptr_t, item *QTextInlineObject, posInDocument int, format *QTextFormat) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(item QTextInlineObject, posInDocument int, format *QTextFormat), item QTextInlineObject, posInDocument int, format *QTextFormat))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	item_goptr := newQTextInlineObject(item)
	item_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *item_goptr

	slotval2 := (int)(posInDocument)

	slotval3 := newQTextFormat(format)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_ResizeInlineObject, slotval1, slotval2, slotval3)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_PositionInlineObject(item QTextInlineObject, posInDocument int, format *QTextFormat) {

	QAbstractTextDocumentLayout_virtualbase_PositionInlineObject(unsafe.Pointer(this.h), item.cPointer(), (int)(posInDocument), format.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnPositionInlineObject(slot func(super func(item QTextInlineObject, posInDocument int, format *QTextFormat), item QTextInlineObject, posInDocument int, format *QTextFormat)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_PositionInlineObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_PositionInlineObject
func miqt_exec_callback_QAbstractTextDocumentLayout_PositionInlineObject(self QAbstractTextDocumentLayout, cb intptr_t, item *QTextInlineObject, posInDocument int, format *QTextFormat) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(item QTextInlineObject, posInDocument int, format *QTextFormat), item QTextInlineObject, posInDocument int, format *QTextFormat))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	item_goptr := newQTextInlineObject(item)
	item_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *item_goptr

	slotval2 := (int)(posInDocument)

	slotval3 := newQTextFormat(format)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_PositionInlineObject, slotval1, slotval2, slotval3)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_DrawInlineObject(painter *QPainter, rect *QRectF, object QTextInlineObject, posInDocument int, format *QTextFormat) {

	QAbstractTextDocumentLayout_virtualbase_DrawInlineObject(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), object.cPointer(), (int)(posInDocument), format.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnDrawInlineObject(slot func(super func(painter *QPainter, rect *QRectF, object QTextInlineObject, posInDocument int, format *QTextFormat), painter *QPainter, rect *QRectF, object QTextInlineObject, posInDocument int, format *QTextFormat)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_DrawInlineObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_DrawInlineObject
func miqt_exec_callback_QAbstractTextDocumentLayout_DrawInlineObject(self QAbstractTextDocumentLayout, cb intptr_t, painter *QPainter, rect *QRectF, object *QTextInlineObject, posInDocument int, format *QTextFormat) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRectF, object QTextInlineObject, posInDocument int, format *QTextFormat), painter *QPainter, rect *QRectF, object QTextInlineObject, posInDocument int, format *QTextFormat))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRectF(rect)

	object_goptr := newQTextInlineObject(object)
	object_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval3 := *object_goptr

	slotval4 := (int)(posInDocument)

	slotval5 := newQTextFormat(format)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_DrawInlineObject, slotval1, slotval2, slotval3, slotval4, slotval5)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QAbstractTextDocumentLayout_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QAbstractTextDocumentLayout) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_Event
func miqt_exec_callback_QAbstractTextDocumentLayout_Event(self QAbstractTextDocumentLayout, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QAbstractTextDocumentLayout_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QAbstractTextDocumentLayout) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_EventFilter
func miqt_exec_callback_QAbstractTextDocumentLayout_EventFilter(self QAbstractTextDocumentLayout, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QAbstractTextDocumentLayout_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_TimerEvent
func miqt_exec_callback_QAbstractTextDocumentLayout_TimerEvent(self QAbstractTextDocumentLayout, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_ChildEvent(event *QChildEvent) {

	QAbstractTextDocumentLayout_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_ChildEvent
func miqt_exec_callback_QAbstractTextDocumentLayout_ChildEvent(self QAbstractTextDocumentLayout, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_CustomEvent(event *QEvent) {

	QAbstractTextDocumentLayout_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_CustomEvent
func miqt_exec_callback_QAbstractTextDocumentLayout_CustomEvent(self QAbstractTextDocumentLayout, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QAbstractTextDocumentLayout_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_ConnectNotify
func miqt_exec_callback_QAbstractTextDocumentLayout_ConnectNotify(self QAbstractTextDocumentLayout, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QAbstractTextDocumentLayout) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QAbstractTextDocumentLayout_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QAbstractTextDocumentLayout) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_DisconnectNotify
func miqt_exec_callback_QAbstractTextDocumentLayout_DisconnectNotify(self QAbstractTextDocumentLayout, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}

type QTextObjectInterface struct {
	h          uintptr
	isSubclass bool
}

func (this *QTextObjectInterface) IntrinsicSize(doc *QTextDocument, posInDocument int, format *QTextFormat) *QSizeF {
	_goptr := newQSizeF(QTextObjectInterface_IntrinsicSize(this.h, doc.cPointer(), (int)(posInDocument), format.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextObjectInterface) DrawObject(painter *QPainter, rect *QRectF, doc *QTextDocument, posInDocument int, format *QTextFormat) {
	QTextObjectInterface_DrawObject(this.h, painter.cPointer(), rect.cPointer(), doc.cPointer(), (int)(posInDocument), format.cPointer())
}

func (this *QTextObjectInterface) OperatorAssign(param1 *QTextObjectInterface) {
	QTextObjectInterface_OperatorAssign(this.h, param1.cPointer())
}

type QAbstractTextDocumentLayout__Selection struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractTextDocumentLayout__Selection constructs a new QAbstractTextDocumentLayout::Selection object.
func NewQAbstractTextDocumentLayout__Selection() *QAbstractTextDocumentLayout__Selection {

	ret := newQAbstractTextDocumentLayout__Selection(QAbstractTextDocumentLayout__Selection_new())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractTextDocumentLayout__Selection) OperatorAssign(param1 *Selection) {
	QAbstractTextDocumentLayout__Selection_OperatorAssign(this.h, param1)
}

type QAbstractTextDocumentLayout__PaintContext struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractTextDocumentLayout__PaintContext constructs a new QAbstractTextDocumentLayout::PaintContext object.
func NewQAbstractTextDocumentLayout__PaintContext() *QAbstractTextDocumentLayout__PaintContext {

	ret := newQAbstractTextDocumentLayout__PaintContext(QAbstractTextDocumentLayout__PaintContext_new())
	ret.isSubclass = true
	return ret
}

func (this *QAbstractTextDocumentLayout__PaintContext) OperatorAssign(param1 *PaintContext) {
	QAbstractTextDocumentLayout__PaintContext_OperatorAssign(this.h, param1)
}
