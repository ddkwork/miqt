package qt

import (
	"unsafe"
)

type QUndoCommand struct {
	h          uintptr
	isSubclass bool
}

// NewQUndoCommand constructs a new QUndoCommand object.
func NewQUndoCommand() *QUndoCommand {
	g := newQUndoCommand(QUndoCommand_new())
	g.isSubclass = true
	return g
}

// NewQUndoCommand2 constructs a new QUndoCommand object.
func NewQUndoCommand2(text string) *QUndoCommand {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQUndoCommand(QUndoCommand_new2(text_ms))
	g.isSubclass = true
	return g
}

// NewQUndoCommand3 constructs a new QUndoCommand object.
func NewQUndoCommand3(parent *QUndoCommand) *QUndoCommand {
	g := newQUndoCommand(QUndoCommand_new3(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQUndoCommand4 constructs a new QUndoCommand object.
func NewQUndoCommand4(text string, parent *QUndoCommand) *QUndoCommand {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	g := newQUndoCommand(QUndoCommand_new4(text_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QUndoCommand) Undo() {
	QUndoCommand_Undo(this.h)
}

func (this *QUndoCommand) Redo() {
	QUndoCommand_Redo(this.h)
}

func (this *QUndoCommand) Text() string {
	var _ms struct_miqt_string = QUndoCommand_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoCommand) ActionText() string {
	var _ms struct_miqt_string = QUndoCommand_ActionText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoCommand) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QUndoCommand_SetText(this.h, text_ms)
}

func (this *QUndoCommand) IsObsolete() bool {
	return (bool)(QUndoCommand_IsObsolete(this.h))
}

func (this *QUndoCommand) SetObsolete(obsolete bool) {
	QUndoCommand_SetObsolete(this.h, (bool)(obsolete))
}

func (this *QUndoCommand) Id() int {
	return (int)(QUndoCommand_Id(this.h))
}

func (this *QUndoCommand) MergeWith(other *QUndoCommand) bool {
	return (bool)(QUndoCommand_MergeWith(this.h, other.cPointer()))
}

func (this *QUndoCommand) ChildCount() int {
	return (int)(QUndoCommand_ChildCount(this.h))
}

func (this *QUndoCommand) Child(index int) *QUndoCommand {
	return newQUndoCommand(QUndoCommand_Child(this.h, (int)(index)))
}

type QUndoStack struct {
	h          uintptr
	isSubclass bool
}

// NewQUndoStack constructs a new QUndoStack object.
func NewQUndoStack() *QUndoStack {
	g := newQUndoStack(QUndoStack_new())
	g.isSubclass = true
	return g
}

// NewQUndoStack2 constructs a new QUndoStack object.
func NewQUndoStack2(parent *QObject) *QUndoStack {
	g := newQUndoStack(QUndoStack_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QUndoStack) MetaObject() *QMetaObject {
	return newQMetaObject(QUndoStack_MetaObject(this.h))
}

func (this *QUndoStack) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QUndoStack_Metacast(this.h, param1_Cstring))
}

func QUndoStack_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QUndoStack_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoStack) Clear() {
	QUndoStack_Clear(this.h)
}

func (this *QUndoStack) Push(cmd *QUndoCommand) {
	QUndoStack_Push(this.h, cmd.cPointer())
}

func (this *QUndoStack) CanUndo() bool {
	return (bool)(QUndoStack_CanUndo(this.h))
}

func (this *QUndoStack) CanRedo() bool {
	return (bool)(QUndoStack_CanRedo(this.h))
}

func (this *QUndoStack) UndoText() string {
	var _ms struct_miqt_string = QUndoStack_UndoText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoStack) RedoText() string {
	var _ms struct_miqt_string = QUndoStack_RedoText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoStack) Count() int {
	return (int)(QUndoStack_Count(this.h))
}

func (this *QUndoStack) Index() int {
	return (int)(QUndoStack_Index(this.h))
}

func (this *QUndoStack) Text(idx int) string {
	var _ms struct_miqt_string = QUndoStack_Text(this.h, (int)(idx))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoStack) CreateUndoAction(parent *QObject) *QAction {
	return newQAction(QUndoStack_CreateUndoAction(this.h, parent.cPointer()))
}

func (this *QUndoStack) CreateRedoAction(parent *QObject) *QAction {
	return newQAction(QUndoStack_CreateRedoAction(this.h, parent.cPointer()))
}

func (this *QUndoStack) IsActive() bool {
	return (bool)(QUndoStack_IsActive(this.h))
}

func (this *QUndoStack) IsClean() bool {
	return (bool)(QUndoStack_IsClean(this.h))
}

func (this *QUndoStack) CleanIndex() int {
	return (int)(QUndoStack_CleanIndex(this.h))
}

func (this *QUndoStack) BeginMacro(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QUndoStack_BeginMacro(this.h, text_ms)
}

func (this *QUndoStack) EndMacro() {
	QUndoStack_EndMacro(this.h)
}

func (this *QUndoStack) SetUndoLimit(limit int) {
	QUndoStack_SetUndoLimit(this.h, (int)(limit))
}

func (this *QUndoStack) UndoLimit() int {
	return (int)(QUndoStack_UndoLimit(this.h))
}

func (this *QUndoStack) Command(index int) *QUndoCommand {
	return newQUndoCommand(QUndoStack_Command(this.h, (int)(index)))
}

func (this *QUndoStack) SetClean() {
	QUndoStack_SetClean(this.h)
}

func (this *QUndoStack) ResetClean() {
	QUndoStack_ResetClean(this.h)
}

func (this *QUndoStack) SetIndex(idx int) {
	QUndoStack_SetIndex(this.h, (int)(idx))
}

func (this *QUndoStack) Undo() {
	QUndoStack_Undo(this.h)
}

func (this *QUndoStack) Redo() {
	QUndoStack_Redo(this.h)
}

func (this *QUndoStack) SetActive() {
	QUndoStack_SetActive(this.h)
}

func (this *QUndoStack) IndexChanged(idx int) {
	QUndoStack_IndexChanged(this.h, (int)(idx))
}

func (this *QUndoStack) OnIndexChanged(slot func(idx int)) {
	QUndoStack_connect_IndexChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_IndexChanged
func miqt_exec_callback_QUndoStack_IndexChanged(cb intptr_t, idx int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(idx int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(idx)

	gofunc(slotval1)
}

func (this *QUndoStack) CleanChanged(clean bool) {
	QUndoStack_CleanChanged(this.h, (bool)(clean))
}

func (this *QUndoStack) OnCleanChanged(slot func(clean bool)) {
	QUndoStack_connect_CleanChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_CleanChanged
func miqt_exec_callback_QUndoStack_CleanChanged(cb intptr_t, clean bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(clean bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(clean)

	gofunc(slotval1)
}

func (this *QUndoStack) CanUndoChanged(canUndo bool) {
	QUndoStack_CanUndoChanged(this.h, (bool)(canUndo))
}

func (this *QUndoStack) OnCanUndoChanged(slot func(canUndo bool)) {
	QUndoStack_connect_CanUndoChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_CanUndoChanged
func miqt_exec_callback_QUndoStack_CanUndoChanged(cb intptr_t, canUndo bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(canUndo bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(canUndo)

	gofunc(slotval1)
}

func (this *QUndoStack) CanRedoChanged(canRedo bool) {
	QUndoStack_CanRedoChanged(this.h, (bool)(canRedo))
}

func (this *QUndoStack) OnCanRedoChanged(slot func(canRedo bool)) {
	QUndoStack_connect_CanRedoChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_CanRedoChanged
func miqt_exec_callback_QUndoStack_CanRedoChanged(cb intptr_t, canRedo bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(canRedo bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(canRedo)

	gofunc(slotval1)
}

func (this *QUndoStack) UndoTextChanged(undoText string) {
	undoText_ms := struct_miqt_string{}
	undoText_ms.data = CString(undoText)
	undoText_ms.len = size_t(len(undoText))
	defer free(unsafe.Pointer(undoText_ms.data))
	QUndoStack_UndoTextChanged(this.h, undoText_ms)
}

func (this *QUndoStack) OnUndoTextChanged(slot func(undoText string)) {
	QUndoStack_connect_UndoTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_UndoTextChanged
func miqt_exec_callback_QUndoStack_UndoTextChanged(cb intptr_t, undoText struct_miqt_string) {
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

func (this *QUndoStack) RedoTextChanged(redoText string) {
	redoText_ms := struct_miqt_string{}
	redoText_ms.data = CString(redoText)
	redoText_ms.len = size_t(len(redoText))
	defer free(unsafe.Pointer(redoText_ms.data))
	QUndoStack_RedoTextChanged(this.h, redoText_ms)
}

func (this *QUndoStack) OnRedoTextChanged(slot func(redoText string)) {
	QUndoStack_connect_RedoTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_RedoTextChanged
func miqt_exec_callback_QUndoStack_RedoTextChanged(cb intptr_t, redoText struct_miqt_string) {
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

func QUndoStack_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoStack_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QUndoStack_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QUndoStack_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QUndoStack) CreateUndoAction2(parent *QObject, prefix string) *QAction {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	return newQAction(QUndoStack_CreateUndoAction2(this.h, parent.cPointer(), prefix_ms))
}

func (this *QUndoStack) CreateRedoAction2(parent *QObject, prefix string) *QAction {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	return newQAction(QUndoStack_CreateRedoAction2(this.h, parent.cPointer(), prefix_ms))
}

func (this *QUndoStack) SetActive1(active bool) {
	QUndoStack_SetActive1(this.h, (bool)(active))
}

func (this *QUndoStack) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QUndoStack_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QUndoStack) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_MetaObject
func miqt_exec_callback_QUndoStack_MetaObject(self QUndoStack, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoStack{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QUndoStack) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QUndoStack_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QUndoStack) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_Metacast
func miqt_exec_callback_QUndoStack_Metacast(self QUndoStack, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QUndoStack{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
