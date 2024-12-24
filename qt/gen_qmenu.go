package qt

import (
	"unsafe"
)

type QMenu struct {
	h          uintptr
	isSubclass bool
}

// NewQMenu constructs a new QMenu object.
func NewQMenu(parent *QWidget) *QMenu {
	g := newQMenu(QMenu_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQMenu2 constructs a new QMenu object.
func NewQMenu2() *QMenu {
	g := newQMenu(QMenu_new2())
	g.isSubclass = true
	return g
}

// NewQMenu3 constructs a new QMenu object.
func NewQMenu3(title string) *QMenu {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	g := newQMenu(QMenu_new3(title_ms))
	g.isSubclass = true
	return g
}

// NewQMenu4 constructs a new QMenu object.
func NewQMenu4(title string, parent *QWidget) *QMenu {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	g := newQMenu(QMenu_new4(title_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QMenu) MetaObject() *QMetaObject {
	return newQMetaObject(QMenu_MetaObject(this.h))
}

func (this *QMenu) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMenu_Metacast(this.h, param1_Cstring))
}

func QMenu_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMenu_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMenu) AddMenu(menu *QMenu) *QAction {
	return newQAction(QMenu_AddMenu(this.h, menu.cPointer()))
}

func (this *QMenu) AddMenuWithTitle(title string) *QMenu {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	return newQMenu(QMenu_AddMenuWithTitle(this.h, title_ms))
}

func (this *QMenu) AddMenu2(icon *QIcon, title string) *QMenu {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	return newQMenu(QMenu_AddMenu2(this.h, icon.cPointer(), title_ms))
}

func (this *QMenu) AddSeparator() *QAction {
	return newQAction(QMenu_AddSeparator(this.h))
}

func (this *QMenu) AddSection(text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QMenu_AddSection(this.h, text_ms))
}

func (this *QMenu) AddSection2(icon *QIcon, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QMenu_AddSection2(this.h, icon.cPointer(), text_ms))
}

func (this *QMenu) InsertMenu(before *QAction, menu *QMenu) *QAction {
	return newQAction(QMenu_InsertMenu(this.h, before.cPointer(), menu.cPointer()))
}

func (this *QMenu) InsertSeparator(before *QAction) *QAction {
	return newQAction(QMenu_InsertSeparator(this.h, before.cPointer()))
}

func (this *QMenu) InsertSection(before *QAction, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QMenu_InsertSection(this.h, before.cPointer(), text_ms))
}

func (this *QMenu) InsertSection2(before *QAction, icon *QIcon, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQAction(QMenu_InsertSection2(this.h, before.cPointer(), icon.cPointer(), text_ms))
}

func (this *QMenu) IsEmpty() bool {
	return (bool)(QMenu_IsEmpty(this.h))
}

func (this *QMenu) Clear() {
	QMenu_Clear(this.h)
}

func (this *QMenu) SetTearOffEnabled(tearOffEnabled bool) {
	QMenu_SetTearOffEnabled(this.h, (bool)(tearOffEnabled))
}

func (this *QMenu) IsTearOffEnabled() bool {
	return (bool)(QMenu_IsTearOffEnabled(this.h))
}

func (this *QMenu) IsTearOffMenuVisible() bool {
	return (bool)(QMenu_IsTearOffMenuVisible(this.h))
}

func (this *QMenu) ShowTearOffMenu() {
	QMenu_ShowTearOffMenu(this.h)
}

func (this *QMenu) ShowTearOffMenuWithPos(pos *QPoint) {
	QMenu_ShowTearOffMenuWithPos(this.h, pos.cPointer())
}

func (this *QMenu) HideTearOffMenu() {
	QMenu_HideTearOffMenu(this.h)
}

func (this *QMenu) SetDefaultAction(defaultAction *QAction) {
	QMenu_SetDefaultAction(this.h, defaultAction.cPointer())
}

func (this *QMenu) DefaultAction() *QAction {
	return newQAction(QMenu_DefaultAction(this.h))
}

func (this *QMenu) SetActiveAction(act *QAction) {
	QMenu_SetActiveAction(this.h, act.cPointer())
}

func (this *QMenu) ActiveAction() *QAction {
	return newQAction(QMenu_ActiveAction(this.h))
}

func (this *QMenu) Popup(pos *QPoint) {
	QMenu_Popup(this.h, pos.cPointer())
}

func (this *QMenu) Exec() *QAction {
	return newQAction(QMenu_Exec(this.h))
}

func (this *QMenu) ExecWithPos(pos *QPoint) *QAction {
	return newQAction(QMenu_ExecWithPos(this.h, pos.cPointer()))
}

func QMenu_Exec2(actions []*QAction, pos *QPoint) *QAction {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	return newQAction(QMenu_Exec2(actions_ma, pos.cPointer()))
}

func (this *QMenu) SizeHint() *QSize {
	_goptr := newQSize(QMenu_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMenu) ActionGeometry(param1 *QAction) *QRect {
	_goptr := newQRect(QMenu_ActionGeometry(this.h, param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMenu) ActionAt(param1 *QPoint) *QAction {
	return newQAction(QMenu_ActionAt(this.h, param1.cPointer()))
}

func (this *QMenu) MenuAction() *QAction {
	return newQAction(QMenu_MenuAction(this.h))
}

func QMenu_MenuInAction(action *QAction) *QMenu {
	return newQMenu(QMenu_MenuInAction(action.cPointer()))
}

func (this *QMenu) Title() string {
	var _ms struct_miqt_string = QMenu_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMenu) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QMenu_SetTitle(this.h, title_ms)
}

func (this *QMenu) Icon() *QIcon {
	_goptr := newQIcon(QMenu_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMenu) SetIcon(icon *QIcon) {
	QMenu_SetIcon(this.h, icon.cPointer())
}

func (this *QMenu) SetNoReplayFor(widget *QWidget) {
	QMenu_SetNoReplayFor(this.h, widget.cPointer())
}

func (this *QMenu) SeparatorsCollapsible() bool {
	return (bool)(QMenu_SeparatorsCollapsible(this.h))
}

func (this *QMenu) SetSeparatorsCollapsible(collapse bool) {
	QMenu_SetSeparatorsCollapsible(this.h, (bool)(collapse))
}

func (this *QMenu) ToolTipsVisible() bool {
	return (bool)(QMenu_ToolTipsVisible(this.h))
}

func (this *QMenu) SetToolTipsVisible(visible bool) {
	QMenu_SetToolTipsVisible(this.h, (bool)(visible))
}

func (this *QMenu) AboutToShow() {
	QMenu_AboutToShow(this.h)
}

func (this *QMenu) OnAboutToShow(slot func()) {
	QMenu_connect_AboutToShow(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenu_AboutToShow
func miqt_exec_callback_QMenu_AboutToShow(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMenu) AboutToHide() {
	QMenu_AboutToHide(this.h)
}

func (this *QMenu) OnAboutToHide(slot func()) {
	QMenu_connect_AboutToHide(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenu_AboutToHide
func miqt_exec_callback_QMenu_AboutToHide(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QMenu) Triggered(action *QAction) {
	QMenu_Triggered(this.h, action.cPointer())
}

func (this *QMenu) OnTriggered(slot func(action *QAction)) {
	QMenu_connect_Triggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenu_Triggered
func miqt_exec_callback_QMenu_Triggered(cb intptr_t, action *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(action)

	gofunc(slotval1)
}

func (this *QMenu) Hovered(action *QAction) {
	QMenu_Hovered(this.h, action.cPointer())
}

func (this *QMenu) OnHovered(slot func(action *QAction)) {
	QMenu_connect_Hovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenu_Hovered
func miqt_exec_callback_QMenu_Hovered(cb intptr_t, action *QAction) {
	gofunc, ok := cgo.Handle(cb).Value().(func(action *QAction))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAction(action)

	gofunc(slotval1)
}

func QMenu_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMenu_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMenu_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMenu_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMenu) Popup2(pos *QPoint, at *QAction) {
	QMenu_Popup2(this.h, pos.cPointer(), at.cPointer())
}

func (this *QMenu) Exec22(pos *QPoint, at *QAction) *QAction {
	return newQAction(QMenu_Exec22(this.h, pos.cPointer(), at.cPointer()))
}

func QMenu_Exec3(actions []*QAction, pos *QPoint, at *QAction) *QAction {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	return newQAction(QMenu_Exec3(actions_ma, pos.cPointer(), at.cPointer()))
}

func QMenu_Exec4(actions []*QAction, pos *QPoint, at *QAction, parent *QWidget) *QAction {
	actions_CArray := (*[0xffff]*QAction)(malloc(size_t(8 * len(actions))))
	defer free(unsafe.Pointer(actions_CArray))
	for i := range actions {
		actions_CArray[i] = actions[i].cPointer()
	}
	actions_ma := struct_miqt_array{len: size_t(len(actions)), data: unsafe.Pointer(actions_CArray)}
	return newQAction(QMenu_Exec4(actions_ma, pos.cPointer(), at.cPointer(), parent.cPointer()))
}

func (this *QMenu) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QMenu_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QMenu) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMenu_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenu_MetaObject
func miqt_exec_callback_QMenu_MetaObject(self QMenu, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMenu{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QMenu) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMenu_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMenu) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMenu_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMenu_Metacast
func miqt_exec_callback_QMenu_Metacast(self QMenu, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMenu{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
