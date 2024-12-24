package qt

import (
	"unsafe"
)

type QUndoGroup struct {
	h          uintptr
	isSubclass bool
}

// NewQUndoGroup constructs a new QUndoGroup object.
func NewQUndoGroup() *QUndoGroup {
	g := newQUndoGroup(QUndoGroup_new())
	g.isSubclass = true
	return g
}

// NewQUndoGroup2 constructs a new QUndoGroup object.
func NewQUndoGroup2(parent *QObject) *QUndoGroup {
	g := newQUndoGroup(QUndoGroup_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QUndoGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QUndoGroup_MetaObject(this.h))
}

func (this *QUndoGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QUndoGroup_Metacast(this.h, param1_Cstring))
}

func QUndoGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QUndoGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoGroup) AddStack(stack *QUndoStack) {
	QUndoGroup_AddStack(this.h, stack.cPointer())
}

func (this *QUndoGroup) RemoveStack(stack *QUndoStack) {
	QUndoGroup_RemoveStack(this.h, stack.cPointer())
}

func (this *QUndoGroup) Stacks() []*QUndoStack {
	var _ma struct_miqt_array = QUndoGroup_Stacks(this.h)
	_ret := make([]*QUndoStack, int(_ma.len))
	_outCast := (*[0xffff]*QUndoStack)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQUndoStack(_outCast[i])
	}
	return _ret
}

func (this *QUndoGroup) ActiveStack() *QUndoStack {
	return newQUndoStack(QUndoGroup_ActiveStack(this.h))
}

func (this *QUndoGroup) CreateUndoAction(parent *QObject) *QAction {
	return newQAction(QUndoGroup_CreateUndoAction(this.h, parent.cPointer()))
}

func (this *QUndoGroup) CreateRedoAction(parent *QObject) *QAction {
	return newQAction(QUndoGroup_CreateRedoAction(this.h, parent.cPointer()))
}

func (this *QUndoGroup) CanUndo() bool {
	return (bool)(QUndoGroup_CanUndo(this.h))
}

func (this *QUndoGroup) CanRedo() bool {
	return (bool)(QUndoGroup_CanRedo(this.h))
}

func (this *QUndoGroup) UndoText() string {
	var _ms struct_miqt_string = QUndoGroup_UndoText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoGroup) RedoText() string {
	var _ms struct_miqt_string = QUndoGroup_RedoText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoGroup) IsClean() bool {
	return (bool)(QUndoGroup_IsClean(this.h))
}

func (this *QUndoGroup) Undo() {
	QUndoGroup_Undo(this.h)
}

func (this *QUndoGroup) Redo() {
	QUndoGroup_Redo(this.h)
}

func (this *QUndoGroup) SetActiveStack(stack *QUndoStack) {
	QUndoGroup_SetActiveStack(this.h, stack.cPointer())
}

func (this *QUndoGroup) ActiveStackChanged(stack *QUndoStack) {
	QUndoGroup_ActiveStackChanged(this.h, stack.cPointer())
}

func (this *QUndoGroup) OnActiveStackChanged(slot func(stack *QUndoStack)) {
	QUndoGroup_connect_ActiveStackChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_ActiveStackChanged
func miqt_exec_callback_QUndoGroup_ActiveStackChanged(cb intptr_t, stack *QUndoStack) {
	gofunc, ok := cgo.Handle(cb).Value().(func(stack *QUndoStack))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUndoStack(stack)

	gofunc(slotval1)
}

func (this *QUndoGroup) IndexChanged(idx int) {
	QUndoGroup_IndexChanged(this.h, (int)(idx))
}

func (this *QUndoGroup) OnIndexChanged(slot func(idx int)) {
	QUndoGroup_connect_IndexChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_IndexChanged
func miqt_exec_callback_QUndoGroup_IndexChanged(cb intptr_t, idx int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(idx int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(idx)

	gofunc(slotval1)
}

func (this *QUndoGroup) CleanChanged(clean bool) {
	QUndoGroup_CleanChanged(this.h, (bool)(clean))
}

func (this *QUndoGroup) OnCleanChanged(slot func(clean bool)) {
	QUndoGroup_connect_CleanChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_CleanChanged
func miqt_exec_callback_QUndoGroup_CleanChanged(cb intptr_t, clean bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(clean bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(clean)

	gofunc(slotval1)
}

func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	QUndoGroup_CanUndoChanged(this.h, (bool)(canUndo))
}

func (this *QUndoGroup) OnCanUndoChanged(slot func(canUndo bool)) {
	QUndoGroup_connect_CanUndoChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_CanUndoChanged
func miqt_exec_callback_QUndoGroup_CanUndoChanged(cb intptr_t, canUndo bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(canUndo bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(canUndo)

	gofunc(slotval1)
}

func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	QUndoGroup_CanRedoChanged(this.h, (bool)(canRedo))
}

func (this *QUndoGroup) OnCanRedoChanged(slot func(canRedo bool)) {
	QUndoGroup_connect_CanRedoChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_CanRedoChanged
func miqt_exec_callback_QUndoGroup_CanRedoChanged(cb intptr_t, canRedo bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(canRedo bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(canRedo)

	gofunc(slotval1)
}

func (this *QUndoGroup) UndoTextChanged(undoText string) {
	undoText_ms := struct_miqt_string{}
	undoText_ms.data = CString(undoText)
	undoText_ms.len = size_t(len(undoText))
	defer free(unsafe.Pointer(undoText_ms.data))
	QUndoGroup_UndoTextChanged(this.h, undoText_ms)
}

func (this *QUndoGroup) OnUndoTextChanged(slot func(undoText string)) {
	QUndoGroup_connect_UndoTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_UndoTextChanged
func miqt_exec_callback_QUndoGroup_UndoTextChanged(cb intptr_t, undoText struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(undoText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var undoText_ms struct_miqt_string = undoText
	undoText_ret := GoStringN(undoText_ms.data, int(int64(undoText_ms.len)))
	free(unsafe.Pointer(undoText_ms.data))
	slotval1 := undoText_ret

	gofunc(slotval1)
}

func (this *QUndoGroup) RedoTextChanged(redoText string) {
	redoText_ms := struct_miqt_string{}
	redoText_ms.data = CString(redoText)
	redoText_ms.len = size_t(len(redoText))
	defer free(unsafe.Pointer(redoText_ms.data))
	QUndoGroup_RedoTextChanged(this.h, redoText_ms)
}

func (this *QUndoGroup) OnRedoTextChanged(slot func(redoText string)) {
	QUndoGroup_connect_RedoTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_RedoTextChanged
func miqt_exec_callback_QUndoGroup_RedoTextChanged(cb intptr_t, redoText struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(redoText string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var redoText_ms struct_miqt_string = redoText
	redoText_ret := GoStringN(redoText_ms.data, int(int64(redoText_ms.len)))
	free(unsafe.Pointer(redoText_ms.data))
	slotval1 := redoText_ret

	gofunc(slotval1)
}

func QUndoGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUndoGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoGroup) CreateUndoAction2(parent *QObject, prefix string) *QAction {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	return newQAction(QUndoGroup_CreateUndoAction2(this.h, parent.cPointer(), prefix_ms))
}

func (this *QUndoGroup) CreateRedoAction2(parent *QObject, prefix string) *QAction {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	return newQAction(QUndoGroup_CreateRedoAction2(this.h, parent.cPointer(), prefix_ms))
}

func (this *QUndoGroup) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QUndoGroup_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QUndoGroup) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoGroup_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_MetaObject
func miqt_exec_callback_QUndoGroup_MetaObject(self QUndoGroup, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoGroup{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QUndoGroup) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QUndoGroup_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QUndoGroup) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoGroup_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoGroup_Metacast
func miqt_exec_callback_QUndoGroup_Metacast(self QUndoGroup, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QUndoGroup{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
