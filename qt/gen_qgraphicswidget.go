package qt

import (
	"unsafe"
)

type QGraphicsWidget__ int

const (
	QGraphicsWidget__Type QGraphicsWidget__ = 11
)

type QGraphicsWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQGraphicsWidget constructs a new QGraphicsWidget object.
func NewQGraphicsWidget() *QGraphicsWidget {

	ret := newQGraphicsWidget(QGraphicsWidget_new())
	ret.isSubclass = true
	return ret
}

// NewQGraphicsWidget2 constructs a new QGraphicsWidget object.
func NewQGraphicsWidget2(parent *QGraphicsItem) *QGraphicsWidget {

	ret := newQGraphicsWidget(QGraphicsWidget_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQGraphicsWidget3 constructs a new QGraphicsWidget object.
func NewQGraphicsWidget3(parent *QGraphicsItem, wFlags WindowType) *QGraphicsWidget {

	ret := newQGraphicsWidget(QGraphicsWidget_new3(parent.cPointer(), (int)(wFlags)))
	ret.isSubclass = true
	return ret
}

func (this *QGraphicsWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QGraphicsWidget_MetaObject(this.h))
}

func (this *QGraphicsWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGraphicsWidget_Metacast(this.h, param1_Cstring))
}

func QGraphicsWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGraphicsWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsWidget) Layout() *QGraphicsLayout {
	return newQGraphicsLayout(QGraphicsWidget_Layout(this.h))
}

func (this *QGraphicsWidget) SetLayout(layout *QGraphicsLayout) {
	QGraphicsWidget_SetLayout(this.h, layout.cPointer())
}

func (this *QGraphicsWidget) AdjustSize() {
	QGraphicsWidget_AdjustSize(this.h)
}

func (this *QGraphicsWidget) LayoutDirection() LayoutDirection {
	return (LayoutDirection)(QGraphicsWidget_LayoutDirection(this.h))
}

func (this *QGraphicsWidget) SetLayoutDirection(direction LayoutDirection) {
	QGraphicsWidget_SetLayoutDirection(this.h, (int)(direction))
}

func (this *QGraphicsWidget) UnsetLayoutDirection() {
	QGraphicsWidget_UnsetLayoutDirection(this.h)
}

func (this *QGraphicsWidget) Style() *QStyle {
	return newQStyle(QGraphicsWidget_Style(this.h))
}

func (this *QGraphicsWidget) SetStyle(style *QStyle) {
	QGraphicsWidget_SetStyle(this.h, style.cPointer())
}

func (this *QGraphicsWidget) Font() *QFont {
	_goptr := newQFont(QGraphicsWidget_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) SetFont(font *QFont) {
	QGraphicsWidget_SetFont(this.h, font.cPointer())
}

func (this *QGraphicsWidget) Palette() *QPalette {
	_goptr := newQPalette(QGraphicsWidget_Palette(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) SetPalette(palette *QPalette) {
	QGraphicsWidget_SetPalette(this.h, palette.cPointer())
}

func (this *QGraphicsWidget) AutoFillBackground() bool {
	return (bool)(QGraphicsWidget_AutoFillBackground(this.h))
}

func (this *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	QGraphicsWidget_SetAutoFillBackground(this.h, (bool)(enabled))
}

func (this *QGraphicsWidget) Resize(size *QSizeF) {
	QGraphicsWidget_Resize(this.h, size.cPointer())
}

func (this *QGraphicsWidget) Resize2(w float64, h float64) {
	QGraphicsWidget_Resize2(this.h, (double)(w), (double)(h))
}

func (this *QGraphicsWidget) Size() *QSizeF {
	_goptr := newQSizeF(QGraphicsWidget_Size(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) SetGeometry(rect *QRectF) {
	QGraphicsWidget_SetGeometry(this.h, rect.cPointer())
}

func (this *QGraphicsWidget) SetGeometry2(x float64, y float64, w float64, h float64) {
	QGraphicsWidget_SetGeometry2(this.h, (double)(x), (double)(y), (double)(w), (double)(h))
}

func (this *QGraphicsWidget) Rect() *QRectF {
	_goptr := newQRectF(QGraphicsWidget_Rect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	QGraphicsWidget_SetContentsMargins(this.h, (double)(left), (double)(top), (double)(right), (double)(bottom))
}

func (this *QGraphicsWidget) SetContentsMarginsWithMargins(margins QMarginsF) {
	QGraphicsWidget_SetContentsMarginsWithMargins(this.h, margins.cPointer())
}

func (this *QGraphicsWidget) GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {
	QGraphicsWidget_GetContentsMargins(this.h, (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))
}

func (this *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	QGraphicsWidget_SetWindowFrameMargins(this.h, (double)(left), (double)(top), (double)(right), (double)(bottom))
}

func (this *QGraphicsWidget) SetWindowFrameMarginsWithMargins(margins QMarginsF) {
	QGraphicsWidget_SetWindowFrameMarginsWithMargins(this.h, margins.cPointer())
}

func (this *QGraphicsWidget) GetWindowFrameMargins(left *float64, top *float64, right *float64, bottom *float64) {
	QGraphicsWidget_GetWindowFrameMargins(this.h, (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))
}

func (this *QGraphicsWidget) UnsetWindowFrameMargins() {
	QGraphicsWidget_UnsetWindowFrameMargins(this.h)
}

func (this *QGraphicsWidget) WindowFrameGeometry() *QRectF {
	_goptr := newQRectF(QGraphicsWidget_WindowFrameGeometry(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) WindowFrameRect() *QRectF {
	_goptr := newQRectF(QGraphicsWidget_WindowFrameRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) WindowFlags() WindowType {
	return (WindowType)(QGraphicsWidget_WindowFlags(this.h))
}

func (this *QGraphicsWidget) WindowType() WindowType {
	return (WindowType)(QGraphicsWidget_WindowType(this.h))
}

func (this *QGraphicsWidget) SetWindowFlags(wFlags WindowType) {
	QGraphicsWidget_SetWindowFlags(this.h, (int)(wFlags))
}

func (this *QGraphicsWidget) IsActiveWindow() bool {
	return (bool)(QGraphicsWidget_IsActiveWindow(this.h))
}

func (this *QGraphicsWidget) SetWindowTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QGraphicsWidget_SetWindowTitle(this.h, title_ms)
}

func (this *QGraphicsWidget) WindowTitle() string {
	var _ms struct_miqt_string = QGraphicsWidget_WindowTitle(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsWidget) FocusPolicy() FocusPolicy {
	return (FocusPolicy)(QGraphicsWidget_FocusPolicy(this.h))
}

func (this *QGraphicsWidget) SetFocusPolicy(policy FocusPolicy) {
	QGraphicsWidget_SetFocusPolicy(this.h, (int)(policy))
}

func QGraphicsWidget_SetTabOrder(first *QGraphicsWidget, second *QGraphicsWidget) {
	QGraphicsWidget_SetTabOrder(first.cPointer(), second.cPointer())
}

func (this *QGraphicsWidget) FocusWidget() *QGraphicsWidget {
	return newQGraphicsWidget(QGraphicsWidget_FocusWidget(this.h))
}

func (this *QGraphicsWidget) GrabShortcut(sequence *QKeySequence) int {
	return (int)(QGraphicsWidget_GrabShortcut(this.h, sequence.cPointer()))
}

func (this *QGraphicsWidget) ReleaseShortcut(id int) {
	QGraphicsWidget_ReleaseShortcut(this.h, (int)(id))
}

func (this *QGraphicsWidget) SetShortcutEnabled(id int) {
	QGraphicsWidget_SetShortcutEnabled(this.h, (int)(id))
}

func (this *QGraphicsWidget) SetShortcutAutoRepeat(id int) {
	QGraphicsWidget_SetShortcutAutoRepeat(this.h, (int)(id))
}

func (this *QGraphicsWidget) AddAction(action *QAction) {
	QGraphicsWidget_AddAction(this.h, action.cPointer())
}

func (this *QGraphicsWidget) AddActions(actions []*QAction) {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	QGraphicsWidget_AddActions(this.h, actions_ma)
}

func (this *QGraphicsWidget) InsertActions(before *QAction, actions []*QAction) {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	QGraphicsWidget_InsertActions(this.h, before.cPointer(), actions_ma)
}

func (this *QGraphicsWidget) InsertAction(before *QAction, action *QAction) {
	QGraphicsWidget_InsertAction(this.h, before.cPointer(), action.cPointer())
}

func (this *QGraphicsWidget) RemoveAction(action *QAction) {
	QGraphicsWidget_RemoveAction(this.h, action.cPointer())
}

func (this *QGraphicsWidget) Actions() []*QAction {
	var _ma struct_miqt_array = QGraphicsWidget_Actions(this.h)
	_ret := make([]*QAction, int(_ma.len))
	_outCast := (*[0xffff]*QAction)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQAction(_outCast[i])
	}
	return _ret
}

func (this *QGraphicsWidget) SetAttribute(attribute WidgetAttribute) {
	QGraphicsWidget_SetAttribute(this.h, (int)(attribute))
}

func (this *QGraphicsWidget) TestAttribute(attribute WidgetAttribute) bool {
	return (bool)(QGraphicsWidget_TestAttribute(this.h, (int)(attribute)))
}

func (this *QGraphicsWidget) Type() int {
	return (int)(QGraphicsWidget_Type(this.h))
}

func (this *QGraphicsWidget) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsWidget_Paint(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsWidget) PaintWindowFrame(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	QGraphicsWidget_PaintWindowFrame(this.h, painter.cPointer(), option.cPointer(), widget.cPointer())
}

func (this *QGraphicsWidget) BoundingRect() *QRectF {
	_goptr := newQRectF(QGraphicsWidget_BoundingRect(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) Shape() *QPainterPath {
	_goptr := newQPainterPath(QGraphicsWidget_Shape(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QGraphicsWidget) GeometryChanged() {
	QGraphicsWidget_GeometryChanged(this.h)
}
func (this *QGraphicsWidget) OnGeometryChanged(slot func()) {
	QGraphicsWidget_connect_GeometryChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_GeometryChanged
func miqt_exec_callback_QGraphicsWidget_GeometryChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsWidget) LayoutChanged() {
	QGraphicsWidget_LayoutChanged(this.h)
}
func (this *QGraphicsWidget) OnLayoutChanged(slot func()) {
	QGraphicsWidget_connect_LayoutChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_LayoutChanged
func miqt_exec_callback_QGraphicsWidget_LayoutChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QGraphicsWidget) Close() bool {
	return (bool)(QGraphicsWidget_Close(this.h))
}

func QGraphicsWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGraphicsWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGraphicsWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGraphicsWidget) GrabShortcut2(sequence *QKeySequence, context ShortcutContext) int {
	return (int)(QGraphicsWidget_GrabShortcut2(this.h, sequence.cPointer(), (int)(context)))
}

func (this *QGraphicsWidget) SetShortcutEnabled2(id int, enabled bool) {
	QGraphicsWidget_SetShortcutEnabled2(this.h, (int)(id), (bool)(enabled))
}

func (this *QGraphicsWidget) SetShortcutAutoRepeat2(id int, enabled bool) {
	QGraphicsWidget_SetShortcutAutoRepeat2(this.h, (int)(id), (bool)(enabled))
}

func (this *QGraphicsWidget) SetAttribute2(attribute WidgetAttribute, on bool) {
	QGraphicsWidget_SetAttribute2(this.h, (int)(attribute), (bool)(on))
}

func (this *QGraphicsWidget) callVirtualBase_SetGeometry(rect *QRectF) {

	QGraphicsWidget_virtualbase_SetGeometry(unsafe.Pointer(this.h), rect.cPointer())

}
func (this *QGraphicsWidget) OnSetGeometry(slot func(super func(rect *QRectF), rect *QRectF)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_SetGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_SetGeometry
func miqt_exec_callback_QGraphicsWidget_SetGeometry(self QGraphicsWidget, cb intptr_t, rect *QRectF) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(rect *QRectF), rect *QRectF))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRectF(rect)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_SetGeometry, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_GetContentsMargins(left *float64, top *float64, right *float64, bottom *float64) {

	QGraphicsWidget_virtualbase_GetContentsMargins(unsafe.Pointer(this.h), (*double)(unsafe.Pointer(left)), (*double)(unsafe.Pointer(top)), (*double)(unsafe.Pointer(right)), (*double)(unsafe.Pointer(bottom)))

}
func (this *QGraphicsWidget) OnGetContentsMargins(slot func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_GetContentsMargins(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_GetContentsMargins
func miqt_exec_callback_QGraphicsWidget_GetContentsMargins(self QGraphicsWidget, cb intptr_t, left *double, top *double, right *double, bottom *double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(left *float64, top *float64, right *float64, bottom *float64), left *float64, top *float64, right *float64, bottom *float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*float64)(unsafe.Pointer(left))

	slotval2 := (*float64)(unsafe.Pointer(top))

	slotval3 := (*float64)(unsafe.Pointer(right))

	slotval4 := (*float64)(unsafe.Pointer(bottom))

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_GetContentsMargins, slotval1, slotval2, slotval3, slotval4)

}

func (this *QGraphicsWidget) callVirtualBase_Type() int {

	return (int)(QGraphicsWidget_virtualbase_Type(unsafe.Pointer(this.h)))

}
func (this *QGraphicsWidget) OnType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_Type(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_Type
func miqt_exec_callback_QGraphicsWidget_Type(self QGraphicsWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_Type)

	return (int)(virtualReturn)

}

func (this *QGraphicsWidget) callVirtualBase_Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {

	QGraphicsWidget_virtualbase_Paint(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), widget.cPointer())

}
func (this *QGraphicsWidget) OnPaint(slot func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_Paint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_Paint
func miqt_exec_callback_QGraphicsWidget_Paint(self QGraphicsWidget, cb intptr_t, painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionGraphicsItem(option)

	slotval3 := newQWidget(widget)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_Paint, slotval1, slotval2, slotval3)

}

func (this *QGraphicsWidget) callVirtualBase_PaintWindowFrame(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {

	QGraphicsWidget_virtualbase_PaintWindowFrame(unsafe.Pointer(this.h), painter.cPointer(), option.cPointer(), widget.cPointer())

}
func (this *QGraphicsWidget) OnPaintWindowFrame(slot func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_PaintWindowFrame(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_PaintWindowFrame
func miqt_exec_callback_QGraphicsWidget_PaintWindowFrame(self QGraphicsWidget, cb intptr_t, painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget), painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQStyleOptionGraphicsItem(option)

	slotval3 := newQWidget(widget)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_PaintWindowFrame, slotval1, slotval2, slotval3)

}

func (this *QGraphicsWidget) callVirtualBase_BoundingRect() *QRectF {

	_goptr := newQRectF(QGraphicsWidget_virtualbase_BoundingRect(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsWidget) OnBoundingRect(slot func(super func() *QRectF) *QRectF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_BoundingRect(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_BoundingRect
func miqt_exec_callback_QGraphicsWidget_BoundingRect(self QGraphicsWidget, cb intptr_t) *QRectF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QRectF) *QRectF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_BoundingRect)

	return virtualReturn.cPointer()

}

func (this *QGraphicsWidget) callVirtualBase_Shape() *QPainterPath {

	_goptr := newQPainterPath(QGraphicsWidget_virtualbase_Shape(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsWidget) OnShape(slot func(super func() *QPainterPath) *QPainterPath) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_Shape(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_Shape
func miqt_exec_callback_QGraphicsWidget_Shape(self QGraphicsWidget, cb intptr_t) *QPainterPath {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainterPath) *QPainterPath)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_Shape)

	return virtualReturn.cPointer()

}

func (this *QGraphicsWidget) callVirtualBase_InitStyleOption(option *QStyleOption) {

	QGraphicsWidget_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QGraphicsWidget) OnInitStyleOption(slot func(super func(option *QStyleOption), option *QStyleOption)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_InitStyleOption
func miqt_exec_callback_QGraphicsWidget_InitStyleOption(self QGraphicsWidget, cb intptr_t, option *QStyleOption) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOption), option *QStyleOption))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOption(option)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_SizeHint(which SizeHint, constraint *QSizeF) *QSizeF {

	_goptr := newQSizeF(QGraphicsWidget_virtualbase_SizeHint(unsafe.Pointer(this.h), (int)(which), constraint.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsWidget) OnSizeHint(slot func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_SizeHint
func miqt_exec_callback_QGraphicsWidget_SizeHint(self QGraphicsWidget, cb intptr_t, which int, constraint *QSizeF) *QSizeF {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(which SizeHint, constraint *QSizeF) *QSizeF, which SizeHint, constraint *QSizeF) *QSizeF)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (SizeHint)(which)

	slotval2 := newQSizeF(constraint)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_SizeHint, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsWidget) callVirtualBase_UpdateGeometry() {

	QGraphicsWidget_virtualbase_UpdateGeometry(unsafe.Pointer(this.h))

}
func (this *QGraphicsWidget) OnUpdateGeometry(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_UpdateGeometry(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_UpdateGeometry
func miqt_exec_callback_QGraphicsWidget_UpdateGeometry(self QGraphicsWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_UpdateGeometry)

}

func (this *QGraphicsWidget) callVirtualBase_ItemChange(change GraphicsItemChange, value *QVariant) *QVariant {

	_goptr := newQVariant(QGraphicsWidget_virtualbase_ItemChange(unsafe.Pointer(this.h), change, value.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsWidget) OnItemChange(slot func(super func(change GraphicsItemChange, value *QVariant) *QVariant, change GraphicsItemChange, value *QVariant) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_ItemChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_ItemChange
func miqt_exec_callback_QGraphicsWidget_ItemChange(self QGraphicsWidget, cb intptr_t, change GraphicsItemChange, value *QVariant) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(change GraphicsItemChange, value *QVariant) *QVariant, change GraphicsItemChange, value *QVariant) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx
	slotval2 := newQVariant(value)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_ItemChange, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsWidget) callVirtualBase_PropertyChange(propertyName string, value *QVariant) *QVariant {
	propertyName_ms := struct_miqt_string{}
	propertyName_ms.data = CString(propertyName)
	propertyName_ms.len = size_t(len(propertyName))
	defer free(unsafe.Pointer(propertyName_ms.data))

	_goptr := newQVariant(QGraphicsWidget_virtualbase_PropertyChange(unsafe.Pointer(this.h), propertyName_ms, value.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QGraphicsWidget) OnPropertyChange(slot func(super func(propertyName string, value *QVariant) *QVariant, propertyName string, value *QVariant) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_PropertyChange(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_PropertyChange
func miqt_exec_callback_QGraphicsWidget_PropertyChange(self QGraphicsWidget, cb intptr_t, propertyName struct_miqt_string, value *QVariant) *QVariant {
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

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_PropertyChange, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QGraphicsWidget) callVirtualBase_SceneEvent(event *QEvent) bool {

	return (bool)(QGraphicsWidget_virtualbase_SceneEvent(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsWidget) OnSceneEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_SceneEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_SceneEvent
func miqt_exec_callback_QGraphicsWidget_SceneEvent(self QGraphicsWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_SceneEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsWidget) callVirtualBase_WindowFrameEvent(e *QEvent) bool {

	return (bool)(QGraphicsWidget_virtualbase_WindowFrameEvent(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QGraphicsWidget) OnWindowFrameEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_WindowFrameEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_WindowFrameEvent
func miqt_exec_callback_QGraphicsWidget_WindowFrameEvent(self QGraphicsWidget, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_WindowFrameEvent, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsWidget) callVirtualBase_WindowFrameSectionAt(pos *QPointF) WindowFrameSection {

	return (WindowFrameSection)(QGraphicsWidget_virtualbase_WindowFrameSectionAt(unsafe.Pointer(this.h), pos.cPointer()))

}
func (this *QGraphicsWidget) OnWindowFrameSectionAt(slot func(super func(pos *QPointF) WindowFrameSection, pos *QPointF) WindowFrameSection) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_WindowFrameSectionAt(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_WindowFrameSectionAt
func miqt_exec_callback_QGraphicsWidget_WindowFrameSectionAt(self QGraphicsWidget, cb intptr_t, pos *QPointF) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(pos *QPointF) WindowFrameSection, pos *QPointF) WindowFrameSection)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPointF(pos)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_WindowFrameSectionAt, slotval1)

	return (int)(virtualReturn)

}

func (this *QGraphicsWidget) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QGraphicsWidget_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QGraphicsWidget) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_Event
func miqt_exec_callback_QGraphicsWidget_Event(self QGraphicsWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsWidget) callVirtualBase_ChangeEvent(event *QEvent) {

	QGraphicsWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_ChangeEvent
func miqt_exec_callback_QGraphicsWidget_ChangeEvent(self QGraphicsWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QGraphicsWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_CloseEvent
func miqt_exec_callback_QGraphicsWidget_CloseEvent(self QGraphicsWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QGraphicsWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_FocusInEvent
func miqt_exec_callback_QGraphicsWidget_FocusInEvent(self QGraphicsWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QGraphicsWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QGraphicsWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_FocusNextPrevChild
func miqt_exec_callback_QGraphicsWidget_FocusNextPrevChild(self QGraphicsWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QGraphicsWidget) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QGraphicsWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_FocusOutEvent
func miqt_exec_callback_QGraphicsWidget_FocusOutEvent(self QGraphicsWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_HideEvent(event *QHideEvent) {

	QGraphicsWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_HideEvent
func miqt_exec_callback_QGraphicsWidget_HideEvent(self QGraphicsWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_MoveEvent(event *QGraphicsSceneMoveEvent) {

	QGraphicsWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnMoveEvent(slot func(super func(event *QGraphicsSceneMoveEvent), event *QGraphicsSceneMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_MoveEvent
func miqt_exec_callback_QGraphicsWidget_MoveEvent(self QGraphicsWidget, cb intptr_t, event *QGraphicsSceneMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneMoveEvent), event *QGraphicsSceneMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneMoveEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_PolishEvent() {

	QGraphicsWidget_virtualbase_PolishEvent(unsafe.Pointer(this.h))

}
func (this *QGraphicsWidget) OnPolishEvent(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_PolishEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_PolishEvent
func miqt_exec_callback_QGraphicsWidget_PolishEvent(self QGraphicsWidget, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_PolishEvent)

}

func (this *QGraphicsWidget) callVirtualBase_ResizeEvent(event *QGraphicsSceneResizeEvent) {

	QGraphicsWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnResizeEvent(slot func(super func(event *QGraphicsSceneResizeEvent), event *QGraphicsSceneResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_ResizeEvent
func miqt_exec_callback_QGraphicsWidget_ResizeEvent(self QGraphicsWidget, cb intptr_t, event *QGraphicsSceneResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneResizeEvent), event *QGraphicsSceneResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneResizeEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_ShowEvent(event *QShowEvent) {

	QGraphicsWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_ShowEvent
func miqt_exec_callback_QGraphicsWidget_ShowEvent(self QGraphicsWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_HoverMoveEvent(event *QGraphicsSceneHoverEvent) {

	QGraphicsWidget_virtualbase_HoverMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnHoverMoveEvent(slot func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_HoverMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_HoverMoveEvent
func miqt_exec_callback_QGraphicsWidget_HoverMoveEvent(self QGraphicsWidget, cb intptr_t, event *QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneHoverEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_HoverMoveEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_HoverLeaveEvent(event *QGraphicsSceneHoverEvent) {

	QGraphicsWidget_virtualbase_HoverLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnHoverLeaveEvent(slot func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_HoverLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_HoverLeaveEvent
func miqt_exec_callback_QGraphicsWidget_HoverLeaveEvent(self QGraphicsWidget, cb intptr_t, event *QGraphicsSceneHoverEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QGraphicsSceneHoverEvent), event *QGraphicsSceneHoverEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQGraphicsSceneHoverEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_HoverLeaveEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_GrabMouseEvent(event *QEvent) {

	QGraphicsWidget_virtualbase_GrabMouseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnGrabMouseEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_GrabMouseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_GrabMouseEvent
func miqt_exec_callback_QGraphicsWidget_GrabMouseEvent(self QGraphicsWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_GrabMouseEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_UngrabMouseEvent(event *QEvent) {

	QGraphicsWidget_virtualbase_UngrabMouseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnUngrabMouseEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_UngrabMouseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_UngrabMouseEvent
func miqt_exec_callback_QGraphicsWidget_UngrabMouseEvent(self QGraphicsWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_UngrabMouseEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_GrabKeyboardEvent(event *QEvent) {

	QGraphicsWidget_virtualbase_GrabKeyboardEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnGrabKeyboardEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_GrabKeyboardEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_GrabKeyboardEvent
func miqt_exec_callback_QGraphicsWidget_GrabKeyboardEvent(self QGraphicsWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_GrabKeyboardEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_UngrabKeyboardEvent(event *QEvent) {

	QGraphicsWidget_virtualbase_UngrabKeyboardEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QGraphicsWidget) OnUngrabKeyboardEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_UngrabKeyboardEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_UngrabKeyboardEvent
func miqt_exec_callback_QGraphicsWidget_UngrabKeyboardEvent(self QGraphicsWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QGraphicsWidget{h: self}).callVirtualBase_UngrabKeyboardEvent, slotval1)

}

func (this *QGraphicsWidget) callVirtualBase_IsEmpty() bool {

	return (bool)(QGraphicsWidget_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QGraphicsWidget) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGraphicsWidget_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGraphicsWidget_IsEmpty
func miqt_exec_callback_QGraphicsWidget_IsEmpty(self QGraphicsWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGraphicsWidget{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}
