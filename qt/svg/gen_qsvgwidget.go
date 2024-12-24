package svg

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QSvgWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQSvgWidget constructs a new QSvgWidget object.
func NewQSvgWidget(parent *qt.QWidget) *QSvgWidget {

	ret := newQSvgWidget(QSvgWidget_new((*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

// NewQSvgWidget2 constructs a new QSvgWidget object.
func NewQSvgWidget2() *QSvgWidget {

	ret := newQSvgWidget(QSvgWidget_new2())
	ret.isSubclass = true
	return ret
}

// NewQSvgWidget3 constructs a new QSvgWidget object.
func NewQSvgWidget3(file string) *QSvgWidget {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))

	ret := newQSvgWidget(QSvgWidget_new3(file_ms))
	ret.isSubclass = true
	return ret
}

// NewQSvgWidget4 constructs a new QSvgWidget object.
func NewQSvgWidget4(file string, parent *qt.QWidget) *QSvgWidget {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))

	ret := newQSvgWidget(QSvgWidget_new4(file_ms, (*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QSvgWidget) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QSvgWidget_MetaObject(this.h)))
}

func (this *QSvgWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSvgWidget_Metacast(this.h, param1_Cstring))
}

func QSvgWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSvgWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgWidget) Renderer() *QSvgRenderer {
	return newQSvgRenderer(QSvgWidget_Renderer(this.h))
}

func (this *QSvgWidget) SizeHint() *qt.QSize {
	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgWidget_SizeHint(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QSvgWidget) Options() Option {
	return (Option)(QSvgWidget_Options(this.h))
}

func (this *QSvgWidget) SetOptions(options Option) {
	QSvgWidget_SetOptions(this.h, (int)(options))
}

func (this *QSvgWidget) Load(file string) {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	QSvgWidget_Load(this.h, file_ms)
}

func (this *QSvgWidget) LoadWithContents(contents []byte) {
	contents_alias := struct_miqt_string{}
	contents_alias.data = (char)(unsafe.Pointer(&contents[0]))
	contents_alias.len = size_t(len(contents))
	QSvgWidget_LoadWithContents(this.h, contents_alias)
}

func QSvgWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSvgWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSvgWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSvgWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSvgWidget) callVirtualBase_SizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgWidget_virtualbase_SizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSvgWidget) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_SizeHint
func miqt_exec_callback_QSvgWidget_SizeHint(self QSvgWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_SizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QSvgWidget) callVirtualBase_PaintEvent(event *qt.QPaintEvent) {

	QSvgWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), (*QPaintEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnPaintEvent(slot func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_PaintEvent
func miqt_exec_callback_QSvgWidget_PaintEvent(self QSvgWidget, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QPaintEvent), event *qt.QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPaintEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_DevType() int {

	return (int)(QSvgWidget_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QSvgWidget) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_DevType
func miqt_exec_callback_QSvgWidget_DevType(self QSvgWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QSvgWidget) callVirtualBase_SetVisible(visible bool) {

	QSvgWidget_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QSvgWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_SetVisible
func miqt_exec_callback_QSvgWidget_SetVisible(self QSvgWidget, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QSvgWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QSvgWidget) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QSvgWidget_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSvgWidget) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MinimumSizeHint
func miqt_exec_callback_QSvgWidget_MinimumSizeHint(self QSvgWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QSvgWidget) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QSvgWidget_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QSvgWidget) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_HeightForWidth
func miqt_exec_callback_QSvgWidget_HeightForWidth(self QSvgWidget, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QSvgWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QSvgWidget_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QSvgWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_HasHeightForWidth
func miqt_exec_callback_QSvgWidget_HasHeightForWidth(self QSvgWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QSvgWidget) callVirtualBase_PaintEngine() *qt.QPaintEngine {

	return qt.UnsafeNewQPaintEngine(unsafe.Pointer(QSvgWidget_virtualbase_PaintEngine(unsafe.Pointer(this.h))))

}
func (this *QSvgWidget) OnPaintEngine(slot func(super func() *qt.QPaintEngine) *qt.QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_PaintEngine
func miqt_exec_callback_QSvgWidget_PaintEngine(self QSvgWidget, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPaintEngine) *qt.QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_PaintEngine)

	return (*QPaintEngine)(virtualReturn.UnsafePointer())

}

func (this *QSvgWidget) callVirtualBase_Event(event *qt.QEvent) bool {

	return (bool)(QSvgWidget_virtualbase_Event(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer())))

}
func (this *QSvgWidget) OnEvent(slot func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_Event
func miqt_exec_callback_QSvgWidget_Event(self QSvgWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent) bool, event *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSvgWidget) callVirtualBase_MousePressEvent(event *qt.QMouseEvent) {

	QSvgWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnMousePressEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MousePressEvent
func miqt_exec_callback_QSvgWidget_MousePressEvent(self QSvgWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_MouseReleaseEvent(event *qt.QMouseEvent) {

	QSvgWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnMouseReleaseEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MouseReleaseEvent
func miqt_exec_callback_QSvgWidget_MouseReleaseEvent(self QSvgWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_MouseDoubleClickEvent(event *qt.QMouseEvent) {

	QSvgWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnMouseDoubleClickEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MouseDoubleClickEvent
func miqt_exec_callback_QSvgWidget_MouseDoubleClickEvent(self QSvgWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_MouseMoveEvent(event *qt.QMouseEvent) {

	QSvgWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), (*QMouseEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnMouseMoveEvent(slot func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MouseMoveEvent
func miqt_exec_callback_QSvgWidget_MouseMoveEvent(self QSvgWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMouseEvent), event *qt.QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMouseEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_WheelEvent(event *qt.QWheelEvent) {

	QSvgWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), (*QWheelEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnWheelEvent(slot func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_WheelEvent
func miqt_exec_callback_QSvgWidget_WheelEvent(self QSvgWidget, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QWheelEvent), event *qt.QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQWheelEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_KeyPressEvent(event *qt.QKeyEvent) {

	QSvgWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), (*QKeyEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnKeyPressEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_KeyPressEvent
func miqt_exec_callback_QSvgWidget_KeyPressEvent(self QSvgWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_KeyReleaseEvent(event *qt.QKeyEvent) {

	QSvgWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), (*QKeyEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnKeyReleaseEvent(slot func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_KeyReleaseEvent
func miqt_exec_callback_QSvgWidget_KeyReleaseEvent(self QSvgWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QKeyEvent), event *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_FocusInEvent(event *qt.QFocusEvent) {

	QSvgWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), (*QFocusEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnFocusInEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_FocusInEvent
func miqt_exec_callback_QSvgWidget_FocusInEvent(self QSvgWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_FocusOutEvent(event *qt.QFocusEvent) {

	QSvgWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), (*QFocusEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnFocusOutEvent(slot func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_FocusOutEvent
func miqt_exec_callback_QSvgWidget_FocusOutEvent(self QSvgWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QFocusEvent), event *qt.QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQFocusEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_EnterEvent(event *qt.QEnterEvent) {

	QSvgWidget_virtualbase_EnterEvent(unsafe.Pointer(this.h), (*QEnterEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnEnterEvent(slot func(super func(event *qt.QEnterEvent), event *qt.QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_EnterEvent
func miqt_exec_callback_QSvgWidget_EnterEvent(self QSvgWidget, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEnterEvent), event *qt.QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEnterEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_LeaveEvent(event *qt.QEvent) {

	QSvgWidget_virtualbase_LeaveEvent(unsafe.Pointer(this.h), (*QEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnLeaveEvent(slot func(super func(event *qt.QEvent), event *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_LeaveEvent
func miqt_exec_callback_QSvgWidget_LeaveEvent(self QSvgWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QEvent), event *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_MoveEvent(event *qt.QMoveEvent) {

	QSvgWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), (*QMoveEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnMoveEvent(slot func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_MoveEvent
func miqt_exec_callback_QSvgWidget_MoveEvent(self QSvgWidget, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QMoveEvent), event *qt.QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQMoveEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_ResizeEvent(event *qt.QResizeEvent) {

	QSvgWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), (*QResizeEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnResizeEvent(slot func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_ResizeEvent
func miqt_exec_callback_QSvgWidget_ResizeEvent(self QSvgWidget, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QResizeEvent), event *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_CloseEvent(event *qt.QCloseEvent) {

	QSvgWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), (*QCloseEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnCloseEvent(slot func(super func(event *qt.QCloseEvent), event *qt.QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_CloseEvent
func miqt_exec_callback_QSvgWidget_CloseEvent(self QSvgWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QCloseEvent), event *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_ContextMenuEvent(event *qt.QContextMenuEvent) {

	QSvgWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), (*QContextMenuEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnContextMenuEvent(slot func(super func(event *qt.QContextMenuEvent), event *qt.QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_ContextMenuEvent
func miqt_exec_callback_QSvgWidget_ContextMenuEvent(self QSvgWidget, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QContextMenuEvent), event *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_TabletEvent(event *qt.QTabletEvent) {

	QSvgWidget_virtualbase_TabletEvent(unsafe.Pointer(this.h), (*QTabletEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnTabletEvent(slot func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_TabletEvent
func miqt_exec_callback_QSvgWidget_TabletEvent(self QSvgWidget, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QTabletEvent), event *qt.QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQTabletEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_ActionEvent(event *qt.QActionEvent) {

	QSvgWidget_virtualbase_ActionEvent(unsafe.Pointer(this.h), (*QActionEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnActionEvent(slot func(super func(event *qt.QActionEvent), event *qt.QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_ActionEvent
func miqt_exec_callback_QSvgWidget_ActionEvent(self QSvgWidget, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QActionEvent), event *qt.QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQActionEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_DragEnterEvent(event *qt.QDragEnterEvent) {

	QSvgWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), (*QDragEnterEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnDragEnterEvent(slot func(super func(event *qt.QDragEnterEvent), event *qt.QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_DragEnterEvent
func miqt_exec_callback_QSvgWidget_DragEnterEvent(self QSvgWidget, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragEnterEvent), event *qt.QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragEnterEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_DragMoveEvent(event *qt.QDragMoveEvent) {

	QSvgWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), (*QDragMoveEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnDragMoveEvent(slot func(super func(event *qt.QDragMoveEvent), event *qt.QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_DragMoveEvent
func miqt_exec_callback_QSvgWidget_DragMoveEvent(self QSvgWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragMoveEvent), event *qt.QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragMoveEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_DragLeaveEvent(event *qt.QDragLeaveEvent) {

	QSvgWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), (*QDragLeaveEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnDragLeaveEvent(slot func(super func(event *qt.QDragLeaveEvent), event *qt.QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_DragLeaveEvent
func miqt_exec_callback_QSvgWidget_DragLeaveEvent(self QSvgWidget, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDragLeaveEvent), event *qt.QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_DropEvent(event *qt.QDropEvent) {

	QSvgWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), (*QDropEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnDropEvent(slot func(super func(event *qt.QDropEvent), event *qt.QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_DropEvent
func miqt_exec_callback_QSvgWidget_DropEvent(self QSvgWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QDropEvent), event *qt.QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQDropEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_ShowEvent(event *qt.QShowEvent) {

	QSvgWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), (*QShowEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnShowEvent(slot func(super func(event *qt.QShowEvent), event *qt.QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_ShowEvent
func miqt_exec_callback_QSvgWidget_ShowEvent(self QSvgWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QShowEvent), event *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_HideEvent(event *qt.QHideEvent) {

	QSvgWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), (*QHideEvent)(event.UnsafePointer()))

}
func (this *QSvgWidget) OnHideEvent(slot func(super func(event *qt.QHideEvent), event *qt.QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_HideEvent
func miqt_exec_callback_QSvgWidget_HideEvent(self QSvgWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *qt.QHideEvent), event *qt.QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQHideEvent(unsafe.Pointer(event))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QSvgWidget_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QSvgWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_NativeEvent
func miqt_exec_callback_QSvgWidget_NativeEvent(self QSvgWidget, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QSvgWidget) callVirtualBase_ChangeEvent(param1 *qt.QEvent) {

	QSvgWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), (*QEvent)(param1.UnsafePointer()))

}
func (this *QSvgWidget) OnChangeEvent(slot func(super func(param1 *qt.QEvent), param1 *qt.QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_ChangeEvent
func miqt_exec_callback_QSvgWidget_ChangeEvent(self QSvgWidget, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QEvent), param1 *qt.QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQEvent(unsafe.Pointer(param1))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QSvgWidget_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QSvgWidget) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_Metric
func miqt_exec_callback_QSvgWidget_Metric(self QSvgWidget, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QSvgWidget) callVirtualBase_InitPainter(painter *qt.QPainter) {

	QSvgWidget_virtualbase_InitPainter(unsafe.Pointer(this.h), (*QPainter)(painter.UnsafePointer()))

}
func (this *QSvgWidget) OnInitPainter(slot func(super func(painter *qt.QPainter), painter *qt.QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_InitPainter
func miqt_exec_callback_QSvgWidget_InitPainter(self QSvgWidget, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *qt.QPainter), painter *qt.QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPainter(unsafe.Pointer(painter))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QSvgWidget) callVirtualBase_Redirected(offset *qt.QPoint) *qt.QPaintDevice {

	return qt.UnsafeNewQPaintDevice(unsafe.Pointer(QSvgWidget_virtualbase_Redirected(unsafe.Pointer(this.h), (*QPoint)(offset.UnsafePointer()))))

}
func (this *QSvgWidget) OnRedirected(slot func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_Redirected
func miqt_exec_callback_QSvgWidget_Redirected(self QSvgWidget, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *qt.QPoint) *qt.QPaintDevice, offset *qt.QPoint) *qt.QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQPoint(unsafe.Pointer(offset))

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return (*QPaintDevice)(virtualReturn.UnsafePointer())

}

func (this *QSvgWidget) callVirtualBase_SharedPainter() *qt.QPainter {

	return qt.UnsafeNewQPainter(unsafe.Pointer(QSvgWidget_virtualbase_SharedPainter(unsafe.Pointer(this.h))))

}
func (this *QSvgWidget) OnSharedPainter(slot func(super func() *qt.QPainter) *qt.QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_SharedPainter
func miqt_exec_callback_QSvgWidget_SharedPainter(self QSvgWidget, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QPainter) *qt.QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_SharedPainter)

	return (*QPainter)(virtualReturn.UnsafePointer())

}

func (this *QSvgWidget) callVirtualBase_InputMethodEvent(param1 *qt.QInputMethodEvent) {

	QSvgWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), (*QInputMethodEvent)(param1.UnsafePointer()))

}
func (this *QSvgWidget) OnInputMethodEvent(slot func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_InputMethodEvent
func miqt_exec_callback_QSvgWidget_InputMethodEvent(self QSvgWidget, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QInputMethodEvent), param1 *qt.QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQInputMethodEvent(unsafe.Pointer(param1))

	gofunc((&QSvgWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QSvgWidget) callVirtualBase_InputMethodQuery(param1 qt.InputMethodQuery) *qt.QVariant {

	_goptr := qt.UnsafeNewQVariant(unsafe.Pointer(QSvgWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QSvgWidget) OnInputMethodQuery(slot func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_InputMethodQuery
func miqt_exec_callback_QSvgWidget_InputMethodQuery(self QSvgWidget, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 qt.InputMethodQuery) *qt.QVariant, param1 qt.InputMethodQuery) *qt.QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (qt.InputMethodQuery)(param1)

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return (*QVariant)(virtualReturn.UnsafePointer())

}

func (this *QSvgWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QSvgWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QSvgWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSvgWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSvgWidget_FocusNextPrevChild
func miqt_exec_callback_QSvgWidget_FocusNextPrevChild(self QSvgWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QSvgWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
