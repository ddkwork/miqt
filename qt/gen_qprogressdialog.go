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

func (this *QProgressDialog) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QProgressDialog_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QProgressDialog) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_MetaObject
func miqt_exec_callback_QProgressDialog_MetaObject(self QProgressDialog, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QProgressDialog{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QProgressDialog) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QProgressDialog_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QProgressDialog) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QProgressDialog_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QProgressDialog_Metacast
func miqt_exec_callback_QProgressDialog_Metacast(self QProgressDialog, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QProgressDialog{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
