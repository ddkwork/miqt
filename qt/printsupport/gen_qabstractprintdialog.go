package printsupport

import (
	"github.com/mappu/miqt/qt"
	"unsafe"
)

type QAbstractPrintDialog__PrintRange int

const (
	QAbstractPrintDialog__AllPages    QAbstractPrintDialog__PrintRange = 0
	QAbstractPrintDialog__Selection   QAbstractPrintDialog__PrintRange = 1
	QAbstractPrintDialog__PageRange   QAbstractPrintDialog__PrintRange = 2
	QAbstractPrintDialog__CurrentPage QAbstractPrintDialog__PrintRange = 3
)

type QAbstractPrintDialog__PrintDialogOption int

const (
	QAbstractPrintDialog__PrintToFile        QAbstractPrintDialog__PrintDialogOption = 1
	QAbstractPrintDialog__PrintSelection     QAbstractPrintDialog__PrintDialogOption = 2
	QAbstractPrintDialog__PrintPageRange     QAbstractPrintDialog__PrintDialogOption = 4
	QAbstractPrintDialog__PrintShowPageSize  QAbstractPrintDialog__PrintDialogOption = 8
	QAbstractPrintDialog__PrintCollateCopies QAbstractPrintDialog__PrintDialogOption = 16
	QAbstractPrintDialog__PrintCurrentPage   QAbstractPrintDialog__PrintDialogOption = 64
)

type QAbstractPrintDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQAbstractPrintDialog constructs a new QAbstractPrintDialog object.
func NewQAbstractPrintDialog(printer *QPrinter) *QAbstractPrintDialog {

	ret := newQAbstractPrintDialog(QAbstractPrintDialog_new(printer.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAbstractPrintDialog2 constructs a new QAbstractPrintDialog object.
func NewQAbstractPrintDialog2(printer *QPrinter, parent *qt.QWidget) *QAbstractPrintDialog {

	ret := newQAbstractPrintDialog(QAbstractPrintDialog_new2(printer.cPointer(), (*QWidget)(parent.UnsafePointer())))
	ret.isSubclass = true
	return ret
}

func (this *QAbstractPrintDialog) MetaObject() *qt.QMetaObject {
	return qt.UnsafeNewQMetaObject(unsafe.Pointer(QAbstractPrintDialog_MetaObject(this.h)))
}

func (this *QAbstractPrintDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAbstractPrintDialog_Metacast(this.h, param1_Cstring))
}

func QAbstractPrintDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAbstractPrintDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractPrintDialog) SetOptionTabs(tabs []*qt.QWidget) {
	tabs_CArray := (*[0xffff]*QWidget)(malloc(size_t(8 * len(tabs))))
	defer free(unsafe.Pointer(tabs_CArray))
	for i := range tabs {
		tabs_CArray[i] = (*QWidget)(tabs[i].UnsafePointer())
	}
	tabs_ma := struct_miqt_array{len: size_t(len(tabs)), data: unsafe.Pointer(tabs_CArray)}
	QAbstractPrintDialog_SetOptionTabs(this.h, tabs_ma)
}

func (this *QAbstractPrintDialog) SetPrintRange(rangeVal PrintRange) {
	QAbstractPrintDialog_SetPrintRange(this.h, rangeVal)
}

func (this *QAbstractPrintDialog) PrintRange() PrintRange {
	xxxxxxxxx
}

func (this *QAbstractPrintDialog) SetMinMax(min int, max int) {
	QAbstractPrintDialog_SetMinMax(this.h, (int)(min), (int)(max))
}

func (this *QAbstractPrintDialog) MinPage() int {
	return (int)(QAbstractPrintDialog_MinPage(this.h))
}

func (this *QAbstractPrintDialog) MaxPage() int {
	return (int)(QAbstractPrintDialog_MaxPage(this.h))
}

func (this *QAbstractPrintDialog) SetFromTo(fromPage int, toPage int) {
	QAbstractPrintDialog_SetFromTo(this.h, (int)(fromPage), (int)(toPage))
}

func (this *QAbstractPrintDialog) FromPage() int {
	return (int)(QAbstractPrintDialog_FromPage(this.h))
}

func (this *QAbstractPrintDialog) ToPage() int {
	return (int)(QAbstractPrintDialog_ToPage(this.h))
}

func (this *QAbstractPrintDialog) Printer() *QPrinter {
	return newQPrinter(QAbstractPrintDialog_Printer(this.h))
}

func QAbstractPrintDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractPrintDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAbstractPrintDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAbstractPrintDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAbstractPrintDialog) callVirtualBase_SetVisible(visible bool) {

	QAbstractPrintDialog_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QAbstractPrintDialog) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_SetVisible
func miqt_exec_callback_QAbstractPrintDialog_SetVisible(self QAbstractPrintDialog, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_SizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QAbstractPrintDialog_virtualbase_SizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractPrintDialog) OnSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_SizeHint
func miqt_exec_callback_QAbstractPrintDialog_SizeHint(self QAbstractPrintDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_SizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QAbstractPrintDialog) callVirtualBase_MinimumSizeHint() *qt.QSize {

	_goptr := qt.UnsafeNewQSize(unsafe.Pointer(QAbstractPrintDialog_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QAbstractPrintDialog) OnMinimumSizeHint(slot func(super func() *qt.QSize) *qt.QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_MinimumSizeHint
func miqt_exec_callback_QAbstractPrintDialog_MinimumSizeHint(self QAbstractPrintDialog, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *qt.QSize) *qt.QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_MinimumSizeHint)

	return (*QSize)(virtualReturn.UnsafePointer())

}

func (this *QAbstractPrintDialog) callVirtualBase_Open() {

	QAbstractPrintDialog_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QAbstractPrintDialog) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_Open
func miqt_exec_callback_QAbstractPrintDialog_Open(self QAbstractPrintDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_Open)

}

func (this *QAbstractPrintDialog) callVirtualBase_Exec() int {

	return (int)(QAbstractPrintDialog_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QAbstractPrintDialog) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_Exec
func miqt_exec_callback_QAbstractPrintDialog_Exec(self QAbstractPrintDialog, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QAbstractPrintDialog) callVirtualBase_Done(param1 int) {

	QAbstractPrintDialog_virtualbase_Done(unsafe.Pointer(this.h), (int)(param1))

}
func (this *QAbstractPrintDialog) OnDone(slot func(super func(param1 int), param1 int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_Done
func miqt_exec_callback_QAbstractPrintDialog_Done(self QAbstractPrintDialog, cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int), param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_Accept() {

	QAbstractPrintDialog_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QAbstractPrintDialog) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_Accept
func miqt_exec_callback_QAbstractPrintDialog_Accept(self QAbstractPrintDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_Accept)

}

func (this *QAbstractPrintDialog) callVirtualBase_Reject() {

	QAbstractPrintDialog_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QAbstractPrintDialog) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_Reject
func miqt_exec_callback_QAbstractPrintDialog_Reject(self QAbstractPrintDialog, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_Reject)

}

func (this *QAbstractPrintDialog) callVirtualBase_KeyPressEvent(param1 *qt.QKeyEvent) {

	QAbstractPrintDialog_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), (*QKeyEvent)(param1.UnsafePointer()))

}
func (this *QAbstractPrintDialog) OnKeyPressEvent(slot func(super func(param1 *qt.QKeyEvent), param1 *qt.QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_KeyPressEvent
func miqt_exec_callback_QAbstractPrintDialog_KeyPressEvent(self QAbstractPrintDialog, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QKeyEvent), param1 *qt.QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQKeyEvent(unsafe.Pointer(param1))

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_CloseEvent(param1 *qt.QCloseEvent) {

	QAbstractPrintDialog_virtualbase_CloseEvent(unsafe.Pointer(this.h), (*QCloseEvent)(param1.UnsafePointer()))

}
func (this *QAbstractPrintDialog) OnCloseEvent(slot func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_CloseEvent
func miqt_exec_callback_QAbstractPrintDialog_CloseEvent(self QAbstractPrintDialog, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QCloseEvent), param1 *qt.QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQCloseEvent(unsafe.Pointer(param1))

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_ShowEvent(param1 *qt.QShowEvent) {

	QAbstractPrintDialog_virtualbase_ShowEvent(unsafe.Pointer(this.h), (*QShowEvent)(param1.UnsafePointer()))

}
func (this *QAbstractPrintDialog) OnShowEvent(slot func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_ShowEvent
func miqt_exec_callback_QAbstractPrintDialog_ShowEvent(self QAbstractPrintDialog, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QShowEvent), param1 *qt.QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQShowEvent(unsafe.Pointer(param1))

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_ResizeEvent(param1 *qt.QResizeEvent) {

	QAbstractPrintDialog_virtualbase_ResizeEvent(unsafe.Pointer(this.h), (*QResizeEvent)(param1.UnsafePointer()))

}
func (this *QAbstractPrintDialog) OnResizeEvent(slot func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_ResizeEvent
func miqt_exec_callback_QAbstractPrintDialog_ResizeEvent(self QAbstractPrintDialog, cb intptr_t, param1 *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QResizeEvent), param1 *qt.QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQResizeEvent(unsafe.Pointer(param1))

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_ContextMenuEvent(param1 *qt.QContextMenuEvent) {

	QAbstractPrintDialog_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), (*QContextMenuEvent)(param1.UnsafePointer()))

}
func (this *QAbstractPrintDialog) OnContextMenuEvent(slot func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_ContextMenuEvent
func miqt_exec_callback_QAbstractPrintDialog_ContextMenuEvent(self QAbstractPrintDialog, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QContextMenuEvent), param1 *qt.QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQContextMenuEvent(unsafe.Pointer(param1))

	gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QAbstractPrintDialog) callVirtualBase_EventFilter(param1 *qt.QObject, param2 *qt.QEvent) bool {

	return (bool)(QAbstractPrintDialog_virtualbase_EventFilter(unsafe.Pointer(this.h), (*QObject)(param1.UnsafePointer()), (*QEvent)(param2.UnsafePointer())))

}
func (this *QAbstractPrintDialog) OnEventFilter(slot func(super func(param1 *qt.QObject, param2 *qt.QEvent) bool, param1 *qt.QObject, param2 *qt.QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAbstractPrintDialog_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAbstractPrintDialog_EventFilter
func miqt_exec_callback_QAbstractPrintDialog_EventFilter(self QAbstractPrintDialog, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *qt.QObject, param2 *qt.QEvent) bool, param1 *qt.QObject, param2 *qt.QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := qt.UnsafeNewQObject(unsafe.Pointer(param1))

	slotval2 := qt.UnsafeNewQEvent(unsafe.Pointer(param2))

	virtualReturn := gofunc((&QAbstractPrintDialog{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}
