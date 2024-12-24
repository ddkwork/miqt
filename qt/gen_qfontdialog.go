package qt

import (
	"unsafe"
)

type QFontDialog__FontDialogOption int

const (
	QFontDialog__NoButtons           QFontDialog__FontDialogOption = 1
	QFontDialog__DontUseNativeDialog QFontDialog__FontDialogOption = 2
	QFontDialog__ScalableFonts       QFontDialog__FontDialogOption = 4
	QFontDialog__NonScalableFonts    QFontDialog__FontDialogOption = 8
	QFontDialog__MonospacedFonts     QFontDialog__FontDialogOption = 16
	QFontDialog__ProportionalFonts   QFontDialog__FontDialogOption = 32
)

type QFontDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQFontDialog constructs a new QFontDialog object.
func NewQFontDialog(parent *QWidget) *QFontDialog {

	ret := newQFontDialog(QFontDialog_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontDialog2 constructs a new QFontDialog object.
func NewQFontDialog2() *QFontDialog {

	ret := newQFontDialog(QFontDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQFontDialog3 constructs a new QFontDialog object.
func NewQFontDialog3(initial *QFont) *QFontDialog {

	ret := newQFontDialog(QFontDialog_new3(initial.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontDialog4 constructs a new QFontDialog object.
func NewQFontDialog4(initial *QFont, parent *QWidget) *QFontDialog {

	ret := newQFontDialog(QFontDialog_new4(initial.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QFontDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QFontDialog_MetaObject(this.h))
}

func (this *QFontDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFontDialog_Metacast(this.h, param1_Cstring))
}

func QFontDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFontDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontDialog) SetCurrentFont(font *QFont) {
	QFontDialog_SetCurrentFont(this.h, font.cPointer())
}

func (this *QFontDialog) CurrentFont() *QFont {
	_goptr := newQFont(QFontDialog_CurrentFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) SelectedFont() *QFont {
	_goptr := newQFont(QFontDialog_SelectedFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) SetOption(option FontDialogOption) {
	QFontDialog_SetOption(this.h, option)
}

func (this *QFontDialog) TestOption(option FontDialogOption) bool {
	return (bool)(QFontDialog_TestOption(this.h, option))
}

func (this *QFontDialog) SetOptions(options FontDialogOptions) {
	QFontDialog_SetOptions(this.h, options)
}

func (this *QFontDialog) Options() FontDialogOptions {
	xxxxxxxxx
}

func (this *QFontDialog) SetVisible(visible bool) {
	QFontDialog_SetVisible(this.h, (bool)(visible))
}

func QFontDialog_GetFont(ok *bool) *QFont {
	_goptr := newQFont(QFontDialog_GetFont((*bool)(unsafe.Pointer(ok))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont2(ok *bool, initial *QFont) *QFont {
	_goptr := newQFont(QFontDialog_GetFont2((*bool)(unsafe.Pointer(ok)), initial.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) CurrentFontChanged(font *QFont) {
	QFontDialog_CurrentFontChanged(this.h, font.cPointer())
}
func (this *QFontDialog) OnCurrentFontChanged(slot func(font *QFont)) {
	QFontDialog_connect_CurrentFontChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_CurrentFontChanged
func miqt_exec_callback_QFontDialog_CurrentFontChanged(cb intptr_t, font *QFont) {
	gofunc, ok := cgo.Handle(cb).Value().(func(font *QFont))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFont(font)

	gofunc(slotval1)
}

func (this *QFontDialog) FontSelected(font *QFont) {
	QFontDialog_FontSelected(this.h, font.cPointer())
}
func (this *QFontDialog) OnFontSelected(slot func(font *QFont)) {
	QFontDialog_connect_FontSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_FontSelected
func miqt_exec_callback_QFontDialog_FontSelected(cb intptr_t, font *QFont) {
	gofunc, ok := cgo.Handle(cb).Value().(func(font *QFont))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFont(font)

	gofunc(slotval1)
}

func QFontDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFontDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFontDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontDialog) SetOption2(option FontDialogOption, on bool) {
	QFontDialog_SetOption2(this.h, option, (bool)(on))
}

func QFontDialog_GetFont22(ok *bool, parent *QWidget) *QFont {
	_goptr := newQFont(QFontDialog_GetFont22((*bool)(unsafe.Pointer(ok)), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont3(ok *bool, initial *QFont, parent *QWidget) *QFont {
	_goptr := newQFont(QFontDialog_GetFont3((*bool)(unsafe.Pointer(ok)), initial.cPointer(), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont4(ok *bool, initial *QFont, parent *QWidget, title string) *QFont {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQFont(QFontDialog_GetFont4((*bool)(unsafe.Pointer(ok)), initial.cPointer(), parent.cPointer(), title_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont5(ok *bool, initial *QFont, parent *QWidget, title string, options FontDialogOptions) *QFont {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQFont(QFontDialog_GetFont5((*bool)(unsafe.Pointer(ok)), initial.cPointer(), parent.cPointer(), title_ms, options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) callVirtualBase_SetVisible(visible bool) {

	QFontDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QFontDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_SetVisible
func miqt_exec_callback_QFontDialog_SetVisible(self QFontDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QFontDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QFontDialog) callVirtualBase_ChangeEvent(event *QEvent) {

	QFontDialog_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QFontDialog) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_ChangeEvent
func miqt_exec_callback_QFontDialog_ChangeEvent(self QFontDialog, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QFontDialog{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QFontDialog) callVirtualBase_Done(result int) {

	QFontDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(result))

}
func (this *QFontDialog) OnDone(slot func(super func(result int), result int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_Done
func miqt_exec_callback_QFontDialog_Done(self QFontDialog, cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(result int), result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc((&QFontDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QFontDialog) callVirtualBase_EventFilter(object *QObject, event *QEvent) bool {

	return (bool)(QFontDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), object.cPointer(), event.cPointer()))

}
func (this *QFontDialog) OnEventFilter(slot func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_EventFilter
func miqt_exec_callback_QFontDialog_EventFilter(self QFontDialog, cb intptr_t, object *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(object *QObject, event *QEvent) bool, object *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(object)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QFontDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QFontDialog) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QFontDialog_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFontDialog) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_SizeHint
func miqt_exec_callback_QFontDialog_SizeHint(self QFontDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontDialog{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QFontDialog) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QFontDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFontDialog) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_MinimumSizeHint
func miqt_exec_callback_QFontDialog_MinimumSizeHint(self QFontDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QFontDialog) callVirtualBase_Open() {

	QFontDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QFontDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_Open
func miqt_exec_callback_QFontDialog_Open(self QFontDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFontDialog{h: self}).callVirtualBase_Open)

}

func (this *QFontDialog) callVirtualBase_Exec() int {

	return (int)(QFontDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QFontDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_Exec
func miqt_exec_callback_QFontDialog_Exec(self QFontDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QFontDialog) callVirtualBase_Accept() {

	QFontDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QFontDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_Accept
func miqt_exec_callback_QFontDialog_Accept(self QFontDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFontDialog{h: self}).callVirtualBase_Accept)

}

func (this *QFontDialog) callVirtualBase_Reject() {

	QFontDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QFontDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_Reject
func miqt_exec_callback_QFontDialog_Reject(self QFontDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFontDialog{h: self}).callVirtualBase_Reject)

}

func (this *QFontDialog) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QFontDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFontDialog) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_KeyPressEvent
func miqt_exec_callback_QFontDialog_KeyPressEvent(self QFontDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QFontDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QFontDialog) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	QFontDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFontDialog) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_CloseEvent
func miqt_exec_callback_QFontDialog_CloseEvent(self QFontDialog, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QFontDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QFontDialog) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QFontDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFontDialog) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_ShowEvent
func miqt_exec_callback_QFontDialog_ShowEvent(self QFontDialog, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QFontDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QFontDialog) callVirtualBase_ResizeEvent(param1 *QResizeEvent) {

	QFontDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFontDialog) OnResizeEvent(slot func(super func(param1 *QResizeEvent), param1 *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_ResizeEvent
func miqt_exec_callback_QFontDialog_ResizeEvent(self QFontDialog, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QResizeEvent), param1 *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(param1)

	gofunc((&QFontDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QFontDialog) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QFontDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFontDialog) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_ContextMenuEvent
func miqt_exec_callback_QFontDialog_ContextMenuEvent(self QFontDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QFontDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}
