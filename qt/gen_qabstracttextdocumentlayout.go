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
	g := newQAbstractTextDocumentLayout(QAbstractTextDocumentLayout_new(doc.cPointer()))
	g.isSubclass = true
	return g
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

func (this *QAbstractTextDocumentLayout) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAbstractTextDocumentLayout_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAbstractTextDocumentLayout) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_MetaObject
func miqt_exec_callback_QAbstractTextDocumentLayout_MetaObject(self QAbstractTextDocumentLayout, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAbstractTextDocumentLayout) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAbstractTextDocumentLayout_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAbstractTextDocumentLayout) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractTextDocumentLayout_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractTextDocumentLayout_Metacast
func miqt_exec_callback_QAbstractTextDocumentLayout_Metacast(self QAbstractTextDocumentLayout, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QAbstractTextDocumentLayout{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
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
	g := newQAbstractTextDocumentLayout__Selection(QAbstractTextDocumentLayout__Selection_new())
	g.isSubclass = true
	return g
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
	g := newQAbstractTextDocumentLayout__PaintContext(QAbstractTextDocumentLayout__PaintContext_new())
	g.isSubclass = true
	return g
}

func (this *QAbstractTextDocumentLayout__PaintContext) OperatorAssign(param1 *PaintContext) {
	QAbstractTextDocumentLayout__PaintContext_OperatorAssign(this.h, param1)
}
