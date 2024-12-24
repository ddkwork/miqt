package qt

import (
	"unsafe"
)

type QInputMethod__Action int

const (
	QInputMethod__Click       QInputMethod__Action = 0
	QInputMethod__ContextMenu QInputMethod__Action = 1
)

type QInputMethod struct {
	h          uintptr
	isSubclass bool
}

func (this *QInputMethod) MetaObject() *QMetaObject {
	return newQMetaObject(QInputMethod_MetaObject(this.h))
}

func (this *QInputMethod) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QInputMethod_Metacast(this.h, param1_Cstring))
}

func QInputMethod_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QInputMethod_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputMethod) InputItemTransform() *QTransform {
	_goptr := newQTransform(QInputMethod_InputItemTransform(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) SetInputItemTransform(transform *QTransform) {
	QInputMethod_SetInputItemTransform(this.h, transform.cPointer())
}

func (this *QInputMethod) InputItemRectangle() *QRectF {
	_goptr := newQRectF(QInputMethod_InputItemRectangle(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) SetInputItemRectangle(rect *QRectF) {
	QInputMethod_SetInputItemRectangle(this.h, rect.cPointer())
}

func (this *QInputMethod) CursorRectangle() *QRectF {
	_goptr := newQRectF(QInputMethod_CursorRectangle(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) AnchorRectangle() *QRectF {
	_goptr := newQRectF(QInputMethod_AnchorRectangle(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) KeyboardRectangle() *QRectF {
	_goptr := newQRectF(QInputMethod_KeyboardRectangle(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) InputItemClipRectangle() *QRectF {
	_goptr := newQRectF(QInputMethod_InputItemClipRectangle(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) IsVisible() bool {
	return (bool)(QInputMethod_IsVisible(this.h))
}

func (this *QInputMethod) SetVisible(visible bool) {
	QInputMethod_SetVisible(this.h, (bool)(visible))
}

func (this *QInputMethod) IsAnimating() bool {
	return (bool)(QInputMethod_IsAnimating(this.h))
}

func (this *QInputMethod) Locale() *QLocale {
	_goptr := newQLocale(QInputMethod_Locale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) InputDirection() LayoutDirection {
	return (LayoutDirection)(QInputMethod_InputDirection(this.h))
}

func QInputMethod_QueryFocusObject(query InputMethodQuery, argument *QVariant) *QVariant {
	_goptr := newQVariant(QInputMethod_QueryFocusObject((int)(query), argument.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputMethod) Show() {
	QInputMethod_Show(this.h)
}

func (this *QInputMethod) Hide() {
	QInputMethod_Hide(this.h)
}

func (this *QInputMethod) Update(queries InputMethodQuery) {
	QInputMethod_Update(this.h, (int)(queries))
}

func (this *QInputMethod) Reset() {
	QInputMethod_Reset(this.h)
}

func (this *QInputMethod) Commit() {
	QInputMethod_Commit(this.h)
}

func (this *QInputMethod) InvokeAction(a Action, cursorPosition int) {
	QInputMethod_InvokeAction(this.h, a, (int)(cursorPosition))
}

func (this *QInputMethod) CursorRectangleChanged() {
	QInputMethod_CursorRectangleChanged(this.h)
}

func (this *QInputMethod) OnCursorRectangleChanged(slot func()) {
	QInputMethod_connect_CursorRectangleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_CursorRectangleChanged
func miqt_exec_callback_QInputMethod_CursorRectangleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) AnchorRectangleChanged() {
	QInputMethod_AnchorRectangleChanged(this.h)
}

func (this *QInputMethod) OnAnchorRectangleChanged(slot func()) {
	QInputMethod_connect_AnchorRectangleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_AnchorRectangleChanged
func miqt_exec_callback_QInputMethod_AnchorRectangleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) KeyboardRectangleChanged() {
	QInputMethod_KeyboardRectangleChanged(this.h)
}

func (this *QInputMethod) OnKeyboardRectangleChanged(slot func()) {
	QInputMethod_connect_KeyboardRectangleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_KeyboardRectangleChanged
func miqt_exec_callback_QInputMethod_KeyboardRectangleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) InputItemClipRectangleChanged() {
	QInputMethod_InputItemClipRectangleChanged(this.h)
}

func (this *QInputMethod) OnInputItemClipRectangleChanged(slot func()) {
	QInputMethod_connect_InputItemClipRectangleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_InputItemClipRectangleChanged
func miqt_exec_callback_QInputMethod_InputItemClipRectangleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) VisibleChanged() {
	QInputMethod_VisibleChanged(this.h)
}

func (this *QInputMethod) OnVisibleChanged(slot func()) {
	QInputMethod_connect_VisibleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_VisibleChanged
func miqt_exec_callback_QInputMethod_VisibleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) AnimatingChanged() {
	QInputMethod_AnimatingChanged(this.h)
}

func (this *QInputMethod) OnAnimatingChanged(slot func()) {
	QInputMethod_connect_AnimatingChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_AnimatingChanged
func miqt_exec_callback_QInputMethod_AnimatingChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) LocaleChanged() {
	QInputMethod_LocaleChanged(this.h)
}

func (this *QInputMethod) OnLocaleChanged(slot func()) {
	QInputMethod_connect_LocaleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_LocaleChanged
func miqt_exec_callback_QInputMethod_LocaleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QInputMethod) InputDirectionChanged(newDirection LayoutDirection) {
	QInputMethod_InputDirectionChanged(this.h, (int)(newDirection))
}

func (this *QInputMethod) OnInputDirectionChanged(slot func(newDirection LayoutDirection)) {
	QInputMethod_connect_InputDirectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputMethod_InputDirectionChanged
func miqt_exec_callback_QInputMethod_InputDirectionChanged(cb intptr_t, newDirection int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(newDirection LayoutDirection))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (LayoutDirection)(newDirection)

	gofunc(slotval1)
}

func QInputMethod_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QInputMethod_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputMethod_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QInputMethod_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
