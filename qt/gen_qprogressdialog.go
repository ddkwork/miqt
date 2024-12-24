package qt

import (
	"unsafe"
)

type QProgressDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQProgressDialog constructs a new QProgressDialog object.
func NewQProgressDialog(parent *QWidget) *QProgressDialog {

	ret := newQProgressDialog(QProgressDialog_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQProgressDialog2 constructs a new QProgressDialog object.
func NewQProgressDialog2() *QProgressDialog {

	ret := newQProgressDialog(QProgressDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQProgressDialog3 constructs a new QProgressDialog object.
func NewQProgressDialog3(labelText string, cancelButtonText string, minimum int, maximum int) *QProgressDialog {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	cancelButtonText_ms := struct_miqt_string{}
	cancelButtonText_ms.data = CString(cancelButtonText)
	cancelButtonText_ms.len = size_t(len(cancelButtonText))
	defer free(unsafe.Pointer(cancelButtonText_ms.data))

	ret := newQProgressDialog(QProgressDialog_new3(labelText_ms, cancelButtonText_ms, (int)(minimum), (int)(maximum)))
	ret.isSubclass = true
	return ret
}

// NewQProgressDialog4 constructs a new QProgressDialog object.
func NewQProgressDialog4(parent *QWidget, flags WindowType) *QProgressDialog {

	ret := newQProgressDialog(QProgressDialog_new4(parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

// NewQProgressDialog5 constructs a new QProgressDialog object.
func NewQProgressDialog5(labelText string, cancelButtonText string, minimum int, maximum int, parent *QWidget) *QProgressDialog {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	cancelButtonText_ms := struct_miqt_string{}
	cancelButtonText_ms.data = CString(cancelButtonText)
	cancelButtonText_ms.len = size_t(len(cancelButtonText))
	defer free(unsafe.Pointer(cancelButtonText_ms.data))

	ret := newQProgressDialog(QProgressDialog_new5(labelText_ms, cancelButtonText_ms, (int)(minimum), (int)(maximum), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQProgressDialog6 constructs a new QProgressDialog object.
func NewQProgressDialog6(labelText string, cancelButtonText string, minimum int, maximum int, parent *QWidget, flags WindowType) *QProgressDialog {
	labelText_ms := struct_miqt_string{}
	labelText_ms.data = CString(labelText)
	labelText_ms.len = size_t(len(labelText))
	defer free(unsafe.Pointer(labelText_ms.data))
	cancelButtonText_ms := struct_miqt_string{}
	cancelButtonText_ms.data = CString(cancelButtonText)
	cancelButtonText_ms.len = size_t(len(cancelButtonText))
	defer free(unsafe.Pointer(cancelButtonText_ms.data))

	ret := newQProgressDialog(QProgressDialog_new6(labelText_ms, cancelButtonText_ms, (int)(minimum), (int)(maximum), parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QProgressDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QProgressDialog_MetaObject(this.h))
}

func (this *QProgressDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QProgressDialog_Metacast(this.h, param1_Cstring))
}

func QProgressDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QProgressDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressDialog) SetLabel(label *QLabel) {
	QProgressDialog_SetLabel(this.h, label.cPointer())
}

func (this *QProgressDialog) SetCancelButton(button *QPushButton) {
	QProgressDialog_SetCancelButton(this.h, button.cPointer())
}

func (this *QProgressDialog) SetBar(bar *QProgressBar) {
	QProgressDialog_SetBar(this.h, bar.cPointer())
}

func (this *QProgressDialog) WasCanceled() bool {
	return (bool)(QProgressDialog_WasCanceled(this.h))
}

func (this *QProgressDialog) Minimum() int {
	return (int)(QProgressDialog_Minimum(this.h))
}

func (this *QProgressDialog) Maximum() int {
	return (int)(QProgressDialog_Maximum(this.h))
}

func (this *QProgressDialog) Value() int {
	return (int)(QProgressDialog_Value(this.h))
}

func (this *QProgressDialog) SizeHint() *QSize {
	_goptr := newQSize(QProgressDialog_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QProgressDialog) LabelText() string {
	var _ms struct_miqt_string = QProgressDialog_LabelText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressDialog) MinimumDuration() int {
	return (int)(QProgressDialog_MinimumDuration(this.h))
}

func (this *QProgressDialog) SetAutoReset(reset bool) {
	QProgressDialog_SetAutoReset(this.h, (bool)(reset))
}

func (this *QProgressDialog) AutoReset() bool {
	return (bool)(QProgressDialog_AutoReset(this.h))
}

func (this *QProgressDialog) SetAutoClose(close bool) {
	QProgressDialog_SetAutoClose(this.h, (bool)(close))
}

func (this *QProgressDialog) AutoClose() bool {
	return (bool)(QProgressDialog_AutoClose(this.h))
}

func (this *QProgressDialog) Cancel() {
	QProgressDialog_Cancel(this.h)
}

func (this *QProgressDialog) Reset() {
	QProgressDialog_Reset(this.h)
}

func (this *QProgressDialog) SetMaximum(maximum int) {
	QProgressDialog_SetMaximum(this.h, (int)(maximum))
}

func (this *QProgressDialog) SetMinimum(minimum int) {
	QProgressDialog_SetMinimum(this.h, (int)(minimum))
}

func (this *QProgressDialog) SetRange(minimum int, maximum int) {
	QProgressDialog_SetRange(this.h, (int)(minimum), (int)(maximum))
}

func (this *QProgressDialog) SetValue(progress int) {
	QProgressDialog_SetValue(this.h, (int)(progress))
}

func (this *QProgressDialog) SetLabelText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QProgressDialog_SetLabelText(this.h, text_ms)
}

func (this *QProgressDialog) SetCancelButtonText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QProgressDialog_SetCancelButtonText(this.h, text_ms)
}

func (this *QProgressDialog) SetMinimumDuration(ms int) {
	QProgressDialog_SetMinimumDuration(this.h, (int)(ms))
}

func (this *QProgressDialog) Canceled() {
	QProgressDialog_Canceled(this.h)
}
func (this *QProgressDialog) OnCanceled(slot func()) {
	QProgressDialog_connect_Canceled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Canceled
func miqt_exec_callback_QProgressDialog_Canceled(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QProgressDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProgressDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QProgressDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QProgressDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QProgressDialog) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QProgressDialog_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProgressDialog) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_SizeHint
func miqt_exec_callback_QProgressDialog_SizeHint(self QProgressDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressDialog{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QProgressDialog) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QProgressDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressDialog) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_ResizeEvent
func miqt_exec_callback_QProgressDialog_ResizeEvent(self QProgressDialog, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QProgressDialog) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QProgressDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressDialog) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_CloseEvent
func miqt_exec_callback_QProgressDialog_CloseEvent(self QProgressDialog, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QProgressDialog) callVirtualBase_ChangeEvent(event *QEvent) {

	QProgressDialog_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressDialog) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_ChangeEvent
func miqt_exec_callback_QProgressDialog_ChangeEvent(self QProgressDialog, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QProgressDialog) callVirtualBase_ShowEvent(event *QShowEvent) {

	QProgressDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QProgressDialog) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_ShowEvent
func miqt_exec_callback_QProgressDialog_ShowEvent(self QProgressDialog, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QProgressDialog) callVirtualBase_SetVisible(visible bool) {

	QProgressDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QProgressDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_SetVisible
func miqt_exec_callback_QProgressDialog_SetVisible(self QProgressDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QProgressDialog) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QProgressDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QProgressDialog) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_MinimumSizeHint
func miqt_exec_callback_QProgressDialog_MinimumSizeHint(self QProgressDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QProgressDialog) callVirtualBase_Open() {

	QProgressDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QProgressDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Open
func miqt_exec_callback_QProgressDialog_Open(self QProgressDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QProgressDialog{h: self}).callVirtualBase_Open)

}

func (this *QProgressDialog) callVirtualBase_Exec() int {

	return (int)(QProgressDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QProgressDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Exec
func miqt_exec_callback_QProgressDialog_Exec(self QProgressDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QProgressDialog) callVirtualBase_Done(param1 int) {

	QProgressDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QProgressDialog) OnDone(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Done
func miqt_exec_callback_QProgressDialog_Done(self QProgressDialog, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QProgressDialog) callVirtualBase_Accept() {

	QProgressDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QProgressDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Accept
func miqt_exec_callback_QProgressDialog_Accept(self QProgressDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QProgressDialog{h: self}).callVirtualBase_Accept)

}

func (this *QProgressDialog) callVirtualBase_Reject() {

	QProgressDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QProgressDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Reject
func miqt_exec_callback_QProgressDialog_Reject(self QProgressDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QProgressDialog{h: self}).callVirtualBase_Reject)

}

func (this *QProgressDialog) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QProgressDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QProgressDialog) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_KeyPressEvent
func miqt_exec_callback_QProgressDialog_KeyPressEvent(self QProgressDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QProgressDialog) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QProgressDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QProgressDialog) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_ContextMenuEvent
func miqt_exec_callback_QProgressDialog_ContextMenuEvent(self QProgressDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QProgressDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QProgressDialog) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QProgressDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QProgressDialog) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_EventFilter
func miqt_exec_callback_QProgressDialog_EventFilter(self QProgressDialog, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QProgressDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
