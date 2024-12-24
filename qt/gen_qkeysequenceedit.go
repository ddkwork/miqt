package qt

import (
	"unsafe"
)

type QKeySequenceEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQKeySequenceEdit constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit(parent *QWidget) *QKeySequenceEdit {

	ret := newQKeySequenceEdit(QKeySequenceEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQKeySequenceEdit2 constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit2() *QKeySequenceEdit {

	ret := newQKeySequenceEdit(QKeySequenceEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQKeySequenceEdit3 constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit3(keySequence *QKeySequence) *QKeySequenceEdit {

	ret := newQKeySequenceEdit(QKeySequenceEdit_new3(keySequence.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQKeySequenceEdit4 constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit4(keySequence *QKeySequence, parent *QWidget) *QKeySequenceEdit {

	ret := newQKeySequenceEdit(QKeySequenceEdit_new4(keySequence.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QKeySequenceEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QKeySequenceEdit_MetaObject(this.h))
}

func (this *QKeySequenceEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QKeySequenceEdit_Metacast(this.h, param1_Cstring))
}

func QKeySequenceEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QKeySequenceEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QKeySequenceEdit) KeySequence() *QKeySequence {
	_goptr := newQKeySequence(QKeySequenceEdit_KeySequence(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QKeySequenceEdit) MaximumSequenceLength() int64 {
	return (int64)(QKeySequenceEdit_MaximumSequenceLength(this.h))
}

func (this *QKeySequenceEdit) SetClearButtonEnabled(enable bool) {
	QKeySequenceEdit_SetClearButtonEnabled(this.h, (bool)(enable))
}

func (this *QKeySequenceEdit) IsClearButtonEnabled() bool {
	return (bool)(QKeySequenceEdit_IsClearButtonEnabled(this.h))
}

func (this *QKeySequenceEdit) SetFinishingKeyCombinations(finishingKeyCombinations []QKeyCombination) {
	finishingKeyCombinations_CArray := (*[0xffff]*QKeyCombination)(malloc(size_t(8 * len(finishingKeyCombinations))))
	defer free(unsafe.Pointer(finishingKeyCombinations_CArray))
	for i := range finishingKeyCombinations {
		finishingKeyCombinations_CArray[i] = finishingKeyCombinations[i].cPointer()
	}
	finishingKeyCombinations_ma := struct_miqt_array{len: size_t(len(finishingKeyCombinations)), data: unsafe.Pointer(finishingKeyCombinations_CArray)}
	QKeySequenceEdit_SetFinishingKeyCombinations(this.h, finishingKeyCombinations_ma)
}

func (this *QKeySequenceEdit) FinishingKeyCombinations() []QKeyCombination {
	var _ma struct_miqt_array = QKeySequenceEdit_FinishingKeyCombinations(this.h)
	_ret := make([]QKeyCombination, int(_ma.len))
	_outCast := (*[0xffff]*QKeyCombination)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQKeyCombination(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QKeySequenceEdit) SetKeySequence(keySequence *QKeySequence) {
	QKeySequenceEdit_SetKeySequence(this.h, keySequence.cPointer())
}

func (this *QKeySequenceEdit) Clear() {
	QKeySequenceEdit_Clear(this.h)
}

func (this *QKeySequenceEdit) SetMaximumSequenceLength(count int64) {
	QKeySequenceEdit_SetMaximumSequenceLength(this.h, (ptrdiff_t)(count))
}

func (this *QKeySequenceEdit) EditingFinished() {
	QKeySequenceEdit_EditingFinished(this.h)
}
func (this *QKeySequenceEdit) OnEditingFinished(slot func()) {
	QKeySequenceEdit_connect_EditingFinished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_EditingFinished
func miqt_exec_callback_QKeySequenceEdit_EditingFinished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QKeySequenceEdit) KeySequenceChanged(keySequence *QKeySequence) {
	QKeySequenceEdit_KeySequenceChanged(this.h, keySequence.cPointer())
}
func (this *QKeySequenceEdit) OnKeySequenceChanged(slot func(keySequence *QKeySequence)) {
	QKeySequenceEdit_connect_KeySequenceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_KeySequenceChanged
func miqt_exec_callback_QKeySequenceEdit_KeySequenceChanged(cb intptr_t, keySequence *QKeySequence) {
	gofunc, ok := cgo.Handle(cb).Value().(func(keySequence *QKeySequence))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeySequence(keySequence)

	gofunc(slotval1)
}

func QKeySequenceEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QKeySequenceEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QKeySequenceEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QKeySequenceEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QKeySequenceEdit) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QKeySequenceEdit_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QKeySequenceEdit) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_Event
func miqt_exec_callback_QKeySequenceEdit_Event(self QKeySequenceEdit, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QKeySequenceEdit) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QKeySequenceEdit_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QKeySequenceEdit) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_KeyPressEvent
func miqt_exec_callback_QKeySequenceEdit_KeyPressEvent(self QKeySequenceEdit, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_KeyReleaseEvent(param1 *QKeyEvent) {

	QKeySequenceEdit_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QKeySequenceEdit) OnKeyReleaseEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_KeyReleaseEvent
func miqt_exec_callback_QKeySequenceEdit_KeyReleaseEvent(self QKeySequenceEdit, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_TimerEvent(param1 *QTimerEvent) {

	QKeySequenceEdit_virtualbase_TimerEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QKeySequenceEdit) OnTimerEvent(slot func(super func(param1 *QTimerEvent), param1 *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_TimerEvent
func miqt_exec_callback_QKeySequenceEdit_TimerEvent(self QKeySequenceEdit, cb intptr_t, param1 *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QTimerEvent), param1 *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(param1)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_FocusOutEvent(param1 *QFocusEvent) {

	QKeySequenceEdit_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QKeySequenceEdit) OnFocusOutEvent(slot func(super func(param1 *QFocusEvent), param1 *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_FocusOutEvent
func miqt_exec_callback_QKeySequenceEdit_FocusOutEvent(self QKeySequenceEdit, cb intptr_t, param1 *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QFocusEvent), param1 *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(param1)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_DevType() int {

	return (int)(QKeySequenceEdit_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QKeySequenceEdit) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_DevType
func miqt_exec_callback_QKeySequenceEdit_DevType(self QKeySequenceEdit, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QKeySequenceEdit) callVirtualBase_SetVisible(visible bool) {

	QKeySequenceEdit_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QKeySequenceEdit) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_SetVisible
func miqt_exec_callback_QKeySequenceEdit_SetVisible(self QKeySequenceEdit, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QKeySequenceEdit_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QKeySequenceEdit) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_SizeHint
func miqt_exec_callback_QKeySequenceEdit_SizeHint(self QKeySequenceEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QKeySequenceEdit) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QKeySequenceEdit_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QKeySequenceEdit) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MinimumSizeHint
func miqt_exec_callback_QKeySequenceEdit_MinimumSizeHint(self QKeySequenceEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QKeySequenceEdit) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QKeySequenceEdit_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QKeySequenceEdit) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_HeightForWidth
func miqt_exec_callback_QKeySequenceEdit_HeightForWidth(self QKeySequenceEdit, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QKeySequenceEdit) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QKeySequenceEdit_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QKeySequenceEdit) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_HasHeightForWidth
func miqt_exec_callback_QKeySequenceEdit_HasHeightForWidth(self QKeySequenceEdit, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QKeySequenceEdit) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QKeySequenceEdit_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QKeySequenceEdit) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_PaintEngine
func miqt_exec_callback_QKeySequenceEdit_PaintEngine(self QKeySequenceEdit, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QKeySequenceEdit) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QKeySequenceEdit_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MousePressEvent
func miqt_exec_callback_QKeySequenceEdit_MousePressEvent(self QKeySequenceEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QKeySequenceEdit_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MouseReleaseEvent
func miqt_exec_callback_QKeySequenceEdit_MouseReleaseEvent(self QKeySequenceEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QKeySequenceEdit_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MouseDoubleClickEvent
func miqt_exec_callback_QKeySequenceEdit_MouseDoubleClickEvent(self QKeySequenceEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QKeySequenceEdit_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MouseMoveEvent
func miqt_exec_callback_QKeySequenceEdit_MouseMoveEvent(self QKeySequenceEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QKeySequenceEdit_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_WheelEvent
func miqt_exec_callback_QKeySequenceEdit_WheelEvent(self QKeySequenceEdit, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QKeySequenceEdit_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_FocusInEvent
func miqt_exec_callback_QKeySequenceEdit_FocusInEvent(self QKeySequenceEdit, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QKeySequenceEdit_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_EnterEvent
func miqt_exec_callback_QKeySequenceEdit_EnterEvent(self QKeySequenceEdit, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_LeaveEvent(event *QEvent) {

	QKeySequenceEdit_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_LeaveEvent
func miqt_exec_callback_QKeySequenceEdit_LeaveEvent(self QKeySequenceEdit, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QKeySequenceEdit_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_PaintEvent
func miqt_exec_callback_QKeySequenceEdit_PaintEvent(self QKeySequenceEdit, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QKeySequenceEdit_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MoveEvent
func miqt_exec_callback_QKeySequenceEdit_MoveEvent(self QKeySequenceEdit, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QKeySequenceEdit_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_ResizeEvent
func miqt_exec_callback_QKeySequenceEdit_ResizeEvent(self QKeySequenceEdit, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QKeySequenceEdit_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_CloseEvent
func miqt_exec_callback_QKeySequenceEdit_CloseEvent(self QKeySequenceEdit, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QKeySequenceEdit_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_ContextMenuEvent
func miqt_exec_callback_QKeySequenceEdit_ContextMenuEvent(self QKeySequenceEdit, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QKeySequenceEdit_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_TabletEvent
func miqt_exec_callback_QKeySequenceEdit_TabletEvent(self QKeySequenceEdit, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_ActionEvent(event *QActionEvent) {

	QKeySequenceEdit_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_ActionEvent
func miqt_exec_callback_QKeySequenceEdit_ActionEvent(self QKeySequenceEdit, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QKeySequenceEdit_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_DragEnterEvent
func miqt_exec_callback_QKeySequenceEdit_DragEnterEvent(self QKeySequenceEdit, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QKeySequenceEdit_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_DragMoveEvent
func miqt_exec_callback_QKeySequenceEdit_DragMoveEvent(self QKeySequenceEdit, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QKeySequenceEdit_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_DragLeaveEvent
func miqt_exec_callback_QKeySequenceEdit_DragLeaveEvent(self QKeySequenceEdit, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_DropEvent(event *QDropEvent) {

	QKeySequenceEdit_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_DropEvent
func miqt_exec_callback_QKeySequenceEdit_DropEvent(self QKeySequenceEdit, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_ShowEvent(event *QShowEvent) {

	QKeySequenceEdit_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_ShowEvent
func miqt_exec_callback_QKeySequenceEdit_ShowEvent(self QKeySequenceEdit, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_HideEvent(event *QHideEvent) {

	QKeySequenceEdit_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QKeySequenceEdit) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_HideEvent
func miqt_exec_callback_QKeySequenceEdit_HideEvent(self QKeySequenceEdit, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QKeySequenceEdit_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QKeySequenceEdit) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_NativeEvent
func miqt_exec_callback_QKeySequenceEdit_NativeEvent(self QKeySequenceEdit, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
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

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QKeySequenceEdit) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QKeySequenceEdit_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QKeySequenceEdit) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_ChangeEvent
func miqt_exec_callback_QKeySequenceEdit_ChangeEvent(self QKeySequenceEdit, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QKeySequenceEdit_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QKeySequenceEdit) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_Metric
func miqt_exec_callback_QKeySequenceEdit_Metric(self QKeySequenceEdit, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QKeySequenceEdit) callVirtualBase_InitPainter(painter *QPainter) {

	QKeySequenceEdit_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QKeySequenceEdit) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_InitPainter
func miqt_exec_callback_QKeySequenceEdit_InitPainter(self QKeySequenceEdit, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QKeySequenceEdit_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QKeySequenceEdit) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_Redirected
func miqt_exec_callback_QKeySequenceEdit_Redirected(self QKeySequenceEdit, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QKeySequenceEdit) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QKeySequenceEdit_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QKeySequenceEdit) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_SharedPainter
func miqt_exec_callback_QKeySequenceEdit_SharedPainter(self QKeySequenceEdit, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QKeySequenceEdit) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QKeySequenceEdit_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QKeySequenceEdit) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_InputMethodEvent
func miqt_exec_callback_QKeySequenceEdit_InputMethodEvent(self QKeySequenceEdit, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QKeySequenceEdit) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QKeySequenceEdit_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QKeySequenceEdit) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_InputMethodQuery
func miqt_exec_callback_QKeySequenceEdit_InputMethodQuery(self QKeySequenceEdit, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QKeySequenceEdit) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QKeySequenceEdit_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QKeySequenceEdit) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_FocusNextPrevChild
func miqt_exec_callback_QKeySequenceEdit_FocusNextPrevChild(self QKeySequenceEdit, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
