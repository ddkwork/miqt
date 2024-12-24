package qt

import (
	"unsafe"
)

type QUndoView struct {
	h          uintptr
	isSubclass bool
}

// NewQUndoView constructs a new QUndoView object.
func NewQUndoView(parent *QWidget) *QUndoView {
	ret := newQUndoView(QUndoView_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView2 constructs a new QUndoView object.
func NewQUndoView2() *QUndoView {
	ret := newQUndoView(QUndoView_new2())
	ret.isSubclass = true
	return ret
}

// NewQUndoView3 constructs a new QUndoView object.
func NewQUndoView3(stack *QUndoStack) *QUndoView {
	ret := newQUndoView(QUndoView_new3(stack.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView4 constructs a new QUndoView object.
func NewQUndoView4(group *QUndoGroup) *QUndoView {
	ret := newQUndoView(QUndoView_new4(group.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView5 constructs a new QUndoView object.
func NewQUndoView5(stack *QUndoStack, parent *QWidget) *QUndoView {
	ret := newQUndoView(QUndoView_new5(stack.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoView6 constructs a new QUndoView object.
func NewQUndoView6(group *QUndoGroup, parent *QWidget) *QUndoView {
	ret := newQUndoView(QUndoView_new6(group.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QUndoView) MetaObject() *QMetaObject {
	return newQMetaObject(QUndoView_MetaObject(this.h))
}

func (this *QUndoView) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QUndoView_Metacast(this.h, param1_Cstring))
}

func QUndoView_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QUndoView_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoView) Stack() *QUndoStack {
	return newQUndoStack(QUndoView_Stack(this.h))
}

func (this *QUndoView) Group() *QUndoGroup {
	return newQUndoGroup(QUndoView_Group(this.h))
}

func (this *QUndoView) SetEmptyLabel(label string) {
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	QUndoView_SetEmptyLabel(this.h, label_ms)
}

func (this *QUndoView) EmptyLabel() string {
	var _ms struct_miqt_string = QUndoView_EmptyLabel(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoView) SetCleanIcon(icon *QIcon) {
	QUndoView_SetCleanIcon(this.h, icon.cPointer())
}

func (this *QUndoView) CleanIcon() *QIcon {
	_goptr := newQIcon(QUndoView_CleanIcon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QUndoView) SetStack(stack *QUndoStack) {
	QUndoView_SetStack(this.h, stack.cPointer())
}

func (this *QUndoView) SetGroup(group *QUndoGroup) {
	QUndoView_SetGroup(this.h, group.cPointer())
}

func QUndoView_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoView_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUndoView_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoView_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoView) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QUndoView_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QUndoView) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_MetaObject
func miqt_exec_callback_QUndoView_MetaObject(self QUndoView, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QUndoView) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QUndoView_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QUndoView) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoView_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoView_Metacast
func miqt_exec_callback_QUndoView_Metacast(self QUndoView, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QUndoView{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
