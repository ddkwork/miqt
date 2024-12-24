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

	ret := newQUndoCommand(QUndoCommand_new())
	ret.isSubclass = true
	return ret
}

// NewQUndoCommand2 constructs a new QUndoCommand object.
func NewQUndoCommand2(text string) *QUndoCommand {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQUndoCommand(QUndoCommand_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQUndoCommand3 constructs a new QUndoCommand object.
func NewQUndoCommand3(parent *QUndoCommand) *QUndoCommand {

	ret := newQUndoCommand(QUndoCommand_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQUndoCommand4 constructs a new QUndoCommand object.
func NewQUndoCommand4(text string, parent *QUndoCommand) *QUndoCommand {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQUndoCommand(QUndoCommand_new4(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QUndoCommand) callVirtualBase_Undo() {

	QUndoCommand_virtualbase_Undo(unsafe.Pointer(this.h))

}
func (this *QUndoCommand) OnUndo(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoCommand_override_virtual_Undo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoCommand_Undo
func miqt_exec_callback_QUndoCommand_Undo(self QUndoCommand, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUndoCommand{h: self}).callVirtualBase_Undo)

}

func (this *QUndoCommand) callVirtualBase_Redo() {

	QUndoCommand_virtualbase_Redo(unsafe.Pointer(this.h))

}
func (this *QUndoCommand) OnRedo(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoCommand_override_virtual_Redo(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoCommand_Redo
func miqt_exec_callback_QUndoCommand_Redo(self QUndoCommand, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QUndoCommand{h: self}).callVirtualBase_Redo)

}

func (this *QUndoCommand) callVirtualBase_Id() int {

	return (int)(QUndoCommand_virtualbase_Id(unsafe.Pointer(this.h)))

}
func (this *QUndoCommand) OnId(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoCommand_override_virtual_Id(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoCommand_Id
func miqt_exec_callback_QUndoCommand_Id(self QUndoCommand, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QUndoCommand{h: self}).callVirtualBase_Id)

	return (int)(virtualReturn)

}

func (this *QUndoCommand) callVirtualBase_MergeWith(other *QUndoCommand) bool {

	return (bool)(QUndoCommand_virtualbase_MergeWith(unsafe.Pointer(this.h), other.cPointer()))

}
func (this *QUndoCommand) OnMergeWith(slot func(super func(other *QUndoCommand) bool, other *QUndoCommand) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoCommand_override_virtual_MergeWith(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoCommand_MergeWith
func miqt_exec_callback_QUndoCommand_MergeWith(self QUndoCommand, cb intptr_t, other *QUndoCommand) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(other *QUndoCommand) bool, other *QUndoCommand) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUndoCommand(other)

	virtualReturn := gofunc((&QUndoCommand{h: self}).callVirtualBase_MergeWith, slotval1)

	return (bool)(virtualReturn)

}

type QUndoStack struct {
	h          uintptr
	isSubclass bool
}

// NewQUndoStack constructs a new QUndoStack object.
func NewQUndoStack() *QUndoStack {

	ret := newQUndoStack(QUndoStack_new())
	ret.isSubclass = true
	return ret
}

// NewQUndoStack2 constructs a new QUndoStack object.
func NewQUndoStack2(parent *QObject) *QUndoStack {

	ret := newQUndoStack(QUndoStack_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QUndoStack) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QUndoStack_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QUndoStack) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_Event
func miqt_exec_callback_QUndoStack_Event(self QUndoStack, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QUndoStack{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QUndoStack) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QUndoStack_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QUndoStack) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_EventFilter
func miqt_exec_callback_QUndoStack_EventFilter(self QUndoStack, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QUndoStack{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QUndoStack) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QUndoStack_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QUndoStack) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_TimerEvent
func miqt_exec_callback_QUndoStack_TimerEvent(self QUndoStack, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QUndoStack{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QUndoStack) callVirtualBase_ChildEvent(event *QChildEvent) {

	QUndoStack_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QUndoStack) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_ChildEvent
func miqt_exec_callback_QUndoStack_ChildEvent(self QUndoStack, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QUndoStack{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QUndoStack) callVirtualBase_CustomEvent(event *QEvent) {

	QUndoStack_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QUndoStack) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_CustomEvent
func miqt_exec_callback_QUndoStack_CustomEvent(self QUndoStack, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QUndoStack{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QUndoStack) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QUndoStack_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QUndoStack) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_ConnectNotify
func miqt_exec_callback_QUndoStack_ConnectNotify(self QUndoStack, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QUndoStack{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QUndoStack) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QUndoStack_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QUndoStack) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QUndoStack_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QUndoStack_DisconnectNotify
func miqt_exec_callback_QUndoStack_DisconnectNotify(self QUndoStack, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QUndoStack{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
