package qt

import (
	"unsafe"
)

type QAction__MenuRole int

const (
	QAction__NoRole                  QAction__MenuRole = 0
	QAction__TextHeuristicRole       QAction__MenuRole = 1
	QAction__ApplicationSpecificRole QAction__MenuRole = 2
	QAction__AboutQtRole             QAction__MenuRole = 3
	QAction__AboutRole               QAction__MenuRole = 4
	QAction__PreferencesRole         QAction__MenuRole = 5
	QAction__QuitRole                QAction__MenuRole = 6
)

type QAction__Priority int

const (
	QAction__LowPriority    QAction__Priority = 0
	QAction__NormalPriority QAction__Priority = 128
	QAction__HighPriority   QAction__Priority = 256
)

type QAction__ActionEvent int

const (
	QAction__Trigger QAction__ActionEvent = 0
	QAction__Hover   QAction__ActionEvent = 1
)

type QAction struct {
	h          uintptr
	isSubclass bool
}

// NewQAction constructs a new QAction object.
func NewQAction() *QAction {
	ret := newQAction(QAction_new())
	ret.isSubclass = true
	return ret
}

// NewQAction2 constructs a new QAction object.
func NewQAction2(text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQAction(QAction_new2(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQAction3 constructs a new QAction object.
func NewQAction3(icon *QIcon, text string) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQAction(QAction_new3(icon.cPointer(), text_ms))
	ret.isSubclass = true
	return ret
}

// NewQAction4 constructs a new QAction object.
func NewQAction4(parent *QObject) *QAction {
	ret := newQAction(QAction_new4(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAction5 constructs a new QAction object.
func NewQAction5(text string, parent *QObject) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQAction(QAction_new5(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQAction6 constructs a new QAction object.
func NewQAction6(icon *QIcon, text string, parent *QObject) *QAction {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQAction(QAction_new6(icon.cPointer(), text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QAction) MetaObject() *QMetaObject {
	return newQMetaObject(QAction_MetaObject(this.h))
}

func (this *QAction) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QAction_Metacast(this.h, param1_Cstring))
}

func QAction_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QAction_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) AssociatedObjects() []*QObject {
	var _ma struct_miqt_array = QAction_AssociatedObjects(this.h)
	_ret := make([]*QObject, int(_ma.len))
	_outCast := (*[0xffff]*QObject)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQObject(_outCast[i])
	}
	return _ret
}

func (this *QAction) SetActionGroup(group *QActionGroup) {
	QAction_SetActionGroup(this.h, group.cPointer())
}

func (this *QAction) ActionGroup() *QActionGroup {
	return newQActionGroup(QAction_ActionGroup(this.h))
}

func (this *QAction) SetIcon(icon *QIcon) {
	QAction_SetIcon(this.h, icon.cPointer())
}

func (this *QAction) Icon() *QIcon {
	_goptr := newQIcon(QAction_Icon(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAction) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QAction_SetText(this.h, text_ms)
}

func (this *QAction) Text() string {
	var _ms struct_miqt_string = QAction_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) SetIconText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QAction_SetIconText(this.h, text_ms)
}

func (this *QAction) IconText() string {
	var _ms struct_miqt_string = QAction_IconText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) SetToolTip(tip string) {
	tip_ms := struct_miqt_string{}
	tip_ms.data = CString(tip)
	tip_ms.len = size_t(len(tip))
	defer free(unsafe.Pointer(tip_ms.data))
	QAction_SetToolTip(this.h, tip_ms)
}

func (this *QAction) ToolTip() string {
	var _ms struct_miqt_string = QAction_ToolTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) SetStatusTip(statusTip string) {
	statusTip_ms := struct_miqt_string{}
	statusTip_ms.data = CString(statusTip)
	statusTip_ms.len = size_t(len(statusTip))
	defer free(unsafe.Pointer(statusTip_ms.data))
	QAction_SetStatusTip(this.h, statusTip_ms)
}

func (this *QAction) StatusTip() string {
	var _ms struct_miqt_string = QAction_StatusTip(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) SetWhatsThis(what string) {
	what_ms := struct_miqt_string{}
	what_ms.data = CString(what)
	what_ms.len = size_t(len(what))
	defer free(unsafe.Pointer(what_ms.data))
	QAction_SetWhatsThis(this.h, what_ms)
}

func (this *QAction) WhatsThis() string {
	var _ms struct_miqt_string = QAction_WhatsThis(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) SetPriority(priority Priority) {
	QAction_SetPriority(this.h, priority)
}

func (this *QAction) Priority() Priority {
	xxxxxxxxx
}

func (this *QAction) SetSeparator(b bool) {
	QAction_SetSeparator(this.h, (bool)(b))
}

func (this *QAction) IsSeparator() bool {
	return (bool)(QAction_IsSeparator(this.h))
}

func (this *QAction) SetShortcut(shortcut *QKeySequence) {
	QAction_SetShortcut(this.h, shortcut.cPointer())
}

func (this *QAction) Shortcut() *QKeySequence {
	_goptr := newQKeySequence(QAction_Shortcut(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAction) SetShortcuts(shortcuts []QKeySequence) {
	shortcuts_CArray := (*[0xffff]*QKeySequence)(malloc(size_t(8 * len(shortcuts))))
	defer free(unsafe.Pointer(shortcuts_CArray))
	for i := range shortcuts {
		shortcuts_CArray[i] = shortcuts[i].cPointer()
	}
	shortcuts_ma := struct_miqt_array{len: size_t(len(shortcuts)), data: unsafe.Pointer(shortcuts_CArray)}
	QAction_SetShortcuts(this.h, shortcuts_ma)
}

func (this *QAction) SetShortcutsWithShortcuts(shortcuts QKeySequence__StandardKey) {
	QAction_SetShortcutsWithShortcuts(this.h, (int)(shortcuts))
}

func (this *QAction) Shortcuts() []QKeySequence {
	var _ma struct_miqt_array = QAction_Shortcuts(this.h)
	_ret := make([]QKeySequence, int(_ma.len))
	_outCast := (*[0xffff]*QKeySequence)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQKeySequence(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QAction) SetShortcutContext(context ShortcutContext) {
	QAction_SetShortcutContext(this.h, (int)(context))
}

func (this *QAction) ShortcutContext() ShortcutContext {
	return (ShortcutContext)(QAction_ShortcutContext(this.h))
}

func (this *QAction) SetAutoRepeat(autoRepeat bool) {
	QAction_SetAutoRepeat(this.h, (bool)(autoRepeat))
}

func (this *QAction) AutoRepeat() bool {
	return (bool)(QAction_AutoRepeat(this.h))
}

func (this *QAction) SetFont(font *QFont) {
	QAction_SetFont(this.h, font.cPointer())
}

func (this *QAction) Font() *QFont {
	_goptr := newQFont(QAction_Font(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAction) SetCheckable(checkable bool) {
	QAction_SetCheckable(this.h, (bool)(checkable))
}

func (this *QAction) IsCheckable() bool {
	return (bool)(QAction_IsCheckable(this.h))
}

func (this *QAction) Data() *QVariant {
	_goptr := newQVariant(QAction_Data(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QAction) SetData(varVal *QVariant) {
	QAction_SetData(this.h, varVal.cPointer())
}

func (this *QAction) IsChecked() bool {
	return (bool)(QAction_IsChecked(this.h))
}

func (this *QAction) IsEnabled() bool {
	return (bool)(QAction_IsEnabled(this.h))
}

func (this *QAction) IsVisible() bool {
	return (bool)(QAction_IsVisible(this.h))
}

func (this *QAction) Activate(event ActionEvent) {
	QAction_Activate(this.h, event)
}

func (this *QAction) SetMenuRole(menuRole MenuRole) {
	QAction_SetMenuRole(this.h, menuRole)
}

func (this *QAction) MenuRole() MenuRole {
	xxxxxxxxx
}

func (this *QAction) SetIconVisibleInMenu(visible bool) {
	QAction_SetIconVisibleInMenu(this.h, (bool)(visible))
}

func (this *QAction) IsIconVisibleInMenu() bool {
	return (bool)(QAction_IsIconVisibleInMenu(this.h))
}

func (this *QAction) SetShortcutVisibleInContextMenu(show bool) {
	QAction_SetShortcutVisibleInContextMenu(this.h, (bool)(show))
}

func (this *QAction) IsShortcutVisibleInContextMenu() bool {
	return (bool)(QAction_IsShortcutVisibleInContextMenu(this.h))
}

func (this *QAction) ShowStatusText() bool {
	return (bool)(QAction_ShowStatusText(this.h))
}

func (this *QAction) Trigger() {
	QAction_Trigger(this.h)
}

func (this *QAction) Hover() {
	QAction_Hover(this.h)
}

func (this *QAction) SetChecked(checked bool) {
	QAction_SetChecked(this.h, (bool)(checked))
}

func (this *QAction) Toggle() {
	QAction_Toggle(this.h)
}

func (this *QAction) SetEnabled(enabled bool) {
	QAction_SetEnabled(this.h, (bool)(enabled))
}

func (this *QAction) ResetEnabled() {
	QAction_ResetEnabled(this.h)
}

func (this *QAction) SetDisabled(b bool) {
	QAction_SetDisabled(this.h, (bool)(b))
}

func (this *QAction) SetVisible(visible bool) {
	QAction_SetVisible(this.h, (bool)(visible))
}

func (this *QAction) Changed() {
	QAction_Changed(this.h)
}

func (this *QAction) OnChanged(slot func()) {
	QAction_connect_Changed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_Changed
func miqt_exec_callback_QAction_Changed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAction) EnabledChanged(enabled bool) {
	QAction_EnabledChanged(this.h, (bool)(enabled))
}

func (this *QAction) OnEnabledChanged(slot func(enabled bool)) {
	QAction_connect_EnabledChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_EnabledChanged
func miqt_exec_callback_QAction_EnabledChanged(cb intptr_t, enabled bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(enabled bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(enabled)

	gofunc(slotval1)
}

func (this *QAction) CheckableChanged(checkable bool) {
	QAction_CheckableChanged(this.h, (bool)(checkable))
}

func (this *QAction) OnCheckableChanged(slot func(checkable bool)) {
	QAction_connect_CheckableChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_CheckableChanged
func miqt_exec_callback_QAction_CheckableChanged(cb intptr_t, checkable bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checkable bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checkable)

	gofunc(slotval1)
}

func (this *QAction) VisibleChanged() {
	QAction_VisibleChanged(this.h)
}

func (this *QAction) OnVisibleChanged(slot func()) {
	QAction_connect_VisibleChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_VisibleChanged
func miqt_exec_callback_QAction_VisibleChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAction) Triggered() {
	QAction_Triggered(this.h)
}

func (this *QAction) OnTriggered(slot func()) {
	QAction_connect_Triggered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_Triggered
func miqt_exec_callback_QAction_Triggered(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAction) Hovered() {
	QAction_Hovered(this.h)
}

func (this *QAction) OnHovered(slot func()) {
	QAction_connect_Hovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_Hovered
func miqt_exec_callback_QAction_Hovered(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QAction) Toggled(param1 bool) {
	QAction_Toggled(this.h, (bool)(param1))
}

func (this *QAction) OnToggled(slot func(param1 bool)) {
	QAction_connect_Toggled(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_Toggled
func miqt_exec_callback_QAction_Toggled(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func QAction_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAction_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QAction_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QAction_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QAction) ShowStatusText1(object *QObject) bool {
	return (bool)(QAction_ShowStatusText1(this.h, object.cPointer()))
}

func (this *QAction) Triggered1(checked bool) {
	QAction_Triggered1(this.h, (bool)(checked))
}

func (this *QAction) OnTriggered1(slot func(checked bool)) {
	QAction_connect_Triggered1(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_Triggered1
func miqt_exec_callback_QAction_Triggered1(cb intptr_t, checked bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(checked bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(checked)

	gofunc(slotval1)
}

func (this *QAction) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QAction_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QAction) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAction_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_MetaObject
func miqt_exec_callback_QAction_MetaObject(self QAction, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QAction{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QAction) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QAction_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QAction) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QAction_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QAction_Metacast
func miqt_exec_callback_QAction_Metacast(self QAction, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QAction{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
