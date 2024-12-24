package qt

import (
	"unsafe"
)

type QColorDialog__ColorDialogOption int

const (
	QColorDialog__ShowAlphaChannel    QColorDialog__ColorDialogOption = 1
	QColorDialog__NoButtons           QColorDialog__ColorDialogOption = 2
	QColorDialog__DontUseNativeDialog QColorDialog__ColorDialogOption = 4
	QColorDialog__NoEyeDropperButton  QColorDialog__ColorDialogOption = 8
)

type QColorDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQColorDialog constructs a new QColorDialog object.
func NewQColorDialog(parent *QWidget) *QColorDialog {

	ret := newQColorDialog(QColorDialog_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQColorDialog2 constructs a new QColorDialog object.
func NewQColorDialog2() *QColorDialog {

	ret := newQColorDialog(QColorDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQColorDialog3 constructs a new QColorDialog object.
func NewQColorDialog3(initial *QColor) *QColorDialog {

	ret := newQColorDialog(QColorDialog_new3(initial.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQColorDialog4 constructs a new QColorDialog object.
func NewQColorDialog4(initial *QColor, parent *QWidget) *QColorDialog {

	ret := newQColorDialog(QColorDialog_new4(initial.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QColorDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QColorDialog_MetaObject(this.h))
}

func (this *QColorDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QColorDialog_Metacast(this.h, param1_Cstring))
}

func QColorDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QColorDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColorDialog) SetCurrentColor(color *QColor) {
	QColorDialog_SetCurrentColor(this.h, color.cPointer())
}

func (this *QColorDialog) CurrentColor() *QColor {
	_goptr := newQColor(QColorDialog_CurrentColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorDialog) SelectedColor() *QColor {
	_goptr := newQColor(QColorDialog_SelectedColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorDialog) SetOption(option ColorDialogOption) {
	QColorDialog_SetOption(this.h, option)
}

func (this *QColorDialog) TestOption(option ColorDialogOption) bool {
	return (bool)(QColorDialog_TestOption(this.h, option))
}

func (this *QColorDialog) SetOptions(options ColorDialogOptions) {
	QColorDialog_SetOptions(this.h, options)
}

func (this *QColorDialog) Options() ColorDialogOptions {
	xxxxxxxxx
}

func (this *QColorDialog) SetVisible(visible bool) {
	QColorDialog_SetVisible(this.h, (bool)(visible))
}

func QColorDialog_GetColor() *QColor {
	_goptr := newQColor(QColorDialog_GetColor())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_CustomCount() int {
	return (int)(QColorDialog_CustomCount())
}

func QColorDialog_CustomColor(index int) *QColor {
	_goptr := newQColor(QColorDialog_CustomColor((int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_SetCustomColor(index int, color QColor) {
	QColorDialog_SetCustomColor((int)(index), color.cPointer())
}

func QColorDialog_StandardColor(index int) *QColor {
	_goptr := newQColor(QColorDialog_StandardColor((int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_SetStandardColor(index int, color QColor) {
	QColorDialog_SetStandardColor((int)(index), color.cPointer())
}

func (this *QColorDialog) CurrentColorChanged(color *QColor) {
	QColorDialog_CurrentColorChanged(this.h, color.cPointer())
}
func (this *QColorDialog) OnCurrentColorChanged(slot func(color *QColor)) {
	QColorDialog_connect_CurrentColorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_CurrentColorChanged
func miqt_exec_callback_QColorDialog_CurrentColorChanged(cb intptr_t, color *QColor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(color *QColor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQColor(color)

	gofunc(slotval1)
}

func (this *QColorDialog) ColorSelected(color *QColor) {
	QColorDialog_ColorSelected(this.h, color.cPointer())
}
func (this *QColorDialog) OnColorSelected(slot func(color *QColor)) {
	QColorDialog_connect_ColorSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_ColorSelected
func miqt_exec_callback_QColorDialog_ColorSelected(cb intptr_t, color *QColor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(color *QColor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQColor(color)

	gofunc(slotval1)
}

func QColorDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QColorDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QColorDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QColorDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColorDialog) SetOption2(option ColorDialogOption, on bool) {
	QColorDialog_SetOption2(this.h, option, (bool)(on))
}

func QColorDialog_GetColor1(initial *QColor) *QColor {
	_goptr := newQColor(QColorDialog_GetColor1(initial.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_GetColor2(initial *QColor, parent *QWidget) *QColor {
	_goptr := newQColor(QColorDialog_GetColor2(initial.cPointer(), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_GetColor3(initial *QColor, parent *QWidget, title string) *QColor {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQColor(QColorDialog_GetColor3(initial.cPointer(), parent.cPointer(), title_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_GetColor4(initial *QColor, parent *QWidget, title string, options ColorDialogOptions) *QColor {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQColor(QColorDialog_GetColor4(initial.cPointer(), parent.cPointer(), title_ms, options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorDialog) callVirtualBase_SetVisible(visible bool) {

	QColorDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QColorDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_SetVisible
func miqt_exec_callback_QColorDialog_SetVisible(self QColorDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QColorDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QColorDialog) callVirtualBase_ChangeEvent(event *QEvent) {

	QColorDialog_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QColorDialog) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_ChangeEvent
func miqt_exec_callback_QColorDialog_ChangeEvent(self QColorDialog, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QColorDialog{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QColorDialog) callVirtualBase_Done(result int) {

	QColorDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(result))

}
func (this *QColorDialog) OnDone(slot func(super func(result int), result int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_Done
func miqt_exec_callback_QColorDialog_Done(self QColorDialog, cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(result int), result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc((&QColorDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QColorDialog) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QColorDialog_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColorDialog) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_SizeHint
func miqt_exec_callback_QColorDialog_SizeHint(self QColorDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColorDialog{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QColorDialog) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QColorDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QColorDialog) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_MinimumSizeHint
func miqt_exec_callback_QColorDialog_MinimumSizeHint(self QColorDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColorDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QColorDialog) callVirtualBase_Open() {

	QColorDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QColorDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_Open
func miqt_exec_callback_QColorDialog_Open(self QColorDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColorDialog{h: self}).callVirtualBase_Open)

}

func (this *QColorDialog) callVirtualBase_Exec() int {

	return (int)(QColorDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QColorDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_Exec
func miqt_exec_callback_QColorDialog_Exec(self QColorDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColorDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QColorDialog) callVirtualBase_Accept() {

	QColorDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QColorDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_Accept
func miqt_exec_callback_QColorDialog_Accept(self QColorDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColorDialog{h: self}).callVirtualBase_Accept)

}

func (this *QColorDialog) callVirtualBase_Reject() {

	QColorDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QColorDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_Reject
func miqt_exec_callback_QColorDialog_Reject(self QColorDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QColorDialog{h: self}).callVirtualBase_Reject)

}

func (this *QColorDialog) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QColorDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QColorDialog) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_KeyPressEvent
func miqt_exec_callback_QColorDialog_KeyPressEvent(self QColorDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QColorDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QColorDialog) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	QColorDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QColorDialog) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_CloseEvent
func miqt_exec_callback_QColorDialog_CloseEvent(self QColorDialog, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QColorDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QColorDialog) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QColorDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QColorDialog) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_ShowEvent
func miqt_exec_callback_QColorDialog_ShowEvent(self QColorDialog, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QColorDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QColorDialog) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QColorDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QColorDialog) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_ResizeEvent
func miqt_exec_callback_QColorDialog_ResizeEvent(self QColorDialog, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QColorDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QColorDialog) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QColorDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QColorDialog) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_ContextMenuEvent
func miqt_exec_callback_QColorDialog_ContextMenuEvent(self QColorDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QColorDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QColorDialog) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QColorDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QColorDialog) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_EventFilter
func miqt_exec_callback_QColorDialog_EventFilter(self QColorDialog, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QColorDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
