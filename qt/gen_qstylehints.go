package qt

import (
	"unsafe"
)

type QStyleHints struct {
	h          uintptr
	isSubclass bool
}

func (this *QStyleHints) MetaObject() *QMetaObject {
	return newQMetaObject(QStyleHints_MetaObject(this.h))
}

func (this *QStyleHints) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStyleHints_Metacast(this.h, param1_Cstring))
}

func QStyleHints_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStyleHints_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStyleHints) SetMouseDoubleClickInterval(mouseDoubleClickInterval int) {
	QStyleHints_SetMouseDoubleClickInterval(this.h, (int)(mouseDoubleClickInterval))
}

func (this *QStyleHints) MouseDoubleClickInterval() int {
	return (int)(QStyleHints_MouseDoubleClickInterval(this.h))
}

func (this *QStyleHints) MouseDoubleClickDistance() int {
	return (int)(QStyleHints_MouseDoubleClickDistance(this.h))
}

func (this *QStyleHints) TouchDoubleTapDistance() int {
	return (int)(QStyleHints_TouchDoubleTapDistance(this.h))
}

func (this *QStyleHints) SetMousePressAndHoldInterval(mousePressAndHoldInterval int) {
	QStyleHints_SetMousePressAndHoldInterval(this.h, (int)(mousePressAndHoldInterval))
}

func (this *QStyleHints) MousePressAndHoldInterval() int {
	return (int)(QStyleHints_MousePressAndHoldInterval(this.h))
}

func (this *QStyleHints) SetStartDragDistance(startDragDistance int) {
	QStyleHints_SetStartDragDistance(this.h, (int)(startDragDistance))
}

func (this *QStyleHints) StartDragDistance() int {
	return (int)(QStyleHints_StartDragDistance(this.h))
}

func (this *QStyleHints) SetStartDragTime(startDragTime int) {
	QStyleHints_SetStartDragTime(this.h, (int)(startDragTime))
}

func (this *QStyleHints) StartDragTime() int {
	return (int)(QStyleHints_StartDragTime(this.h))
}

func (this *QStyleHints) StartDragVelocity() int {
	return (int)(QStyleHints_StartDragVelocity(this.h))
}

func (this *QStyleHints) SetKeyboardInputInterval(keyboardInputInterval int) {
	QStyleHints_SetKeyboardInputInterval(this.h, (int)(keyboardInputInterval))
}

func (this *QStyleHints) KeyboardInputInterval() int {
	return (int)(QStyleHints_KeyboardInputInterval(this.h))
}

func (this *QStyleHints) KeyboardAutoRepeatRate() int {
	return (int)(QStyleHints_KeyboardAutoRepeatRate(this.h))
}

func (this *QStyleHints) KeyboardAutoRepeatRateF() float64 {
	return (float64)(QStyleHints_KeyboardAutoRepeatRateF(this.h))
}

func (this *QStyleHints) SetCursorFlashTime(cursorFlashTime int) {
	QStyleHints_SetCursorFlashTime(this.h, (int)(cursorFlashTime))
}

func (this *QStyleHints) CursorFlashTime() int {
	return (int)(QStyleHints_CursorFlashTime(this.h))
}

func (this *QStyleHints) ShowIsFullScreen() bool {
	return (bool)(QStyleHints_ShowIsFullScreen(this.h))
}

func (this *QStyleHints) ShowIsMaximized() bool {
	return (bool)(QStyleHints_ShowIsMaximized(this.h))
}

func (this *QStyleHints) ShowShortcutsInContextMenus() bool {
	return (bool)(QStyleHints_ShowShortcutsInContextMenus(this.h))
}

func (this *QStyleHints) SetShowShortcutsInContextMenus(showShortcutsInContextMenus bool) {
	QStyleHints_SetShowShortcutsInContextMenus(this.h, (bool)(showShortcutsInContextMenus))
}

func (this *QStyleHints) ContextMenuTrigger() ContextMenuTrigger {
	return (ContextMenuTrigger)(QStyleHints_ContextMenuTrigger(this.h))
}

func (this *QStyleHints) SetContextMenuTrigger(contextMenuTrigger ContextMenuTrigger) {
	QStyleHints_SetContextMenuTrigger(this.h, (int)(contextMenuTrigger))
}

func (this *QStyleHints) PasswordMaskDelay() int {
	return (int)(QStyleHints_PasswordMaskDelay(this.h))
}

func (this *QStyleHints) PasswordMaskCharacter() *QChar {
	_goptr := newQChar(QStyleHints_PasswordMaskCharacter(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStyleHints) FontSmoothingGamma() float64 {
	return (float64)(QStyleHints_FontSmoothingGamma(this.h))
}

func (this *QStyleHints) UseRtlExtensions() bool {
	return (bool)(QStyleHints_UseRtlExtensions(this.h))
}

func (this *QStyleHints) SetFocusOnTouchRelease() bool {
	return (bool)(QStyleHints_SetFocusOnTouchRelease(this.h))
}

func (this *QStyleHints) TabFocusBehavior() TabFocusBehavior {
	return (TabFocusBehavior)(QStyleHints_TabFocusBehavior(this.h))
}

func (this *QStyleHints) SetTabFocusBehavior(tabFocusBehavior TabFocusBehavior) {
	QStyleHints_SetTabFocusBehavior(this.h, (int)(tabFocusBehavior))
}

func (this *QStyleHints) SingleClickActivation() bool {
	return (bool)(QStyleHints_SingleClickActivation(this.h))
}

func (this *QStyleHints) UseHoverEffects() bool {
	return (bool)(QStyleHints_UseHoverEffects(this.h))
}

func (this *QStyleHints) SetUseHoverEffects(useHoverEffects bool) {
	QStyleHints_SetUseHoverEffects(this.h, (bool)(useHoverEffects))
}

func (this *QStyleHints) WheelScrollLines() int {
	return (int)(QStyleHints_WheelScrollLines(this.h))
}

func (this *QStyleHints) SetWheelScrollLines(scrollLines int) {
	QStyleHints_SetWheelScrollLines(this.h, (int)(scrollLines))
}

func (this *QStyleHints) SetMouseQuickSelectionThreshold(threshold int) {
	QStyleHints_SetMouseQuickSelectionThreshold(this.h, (int)(threshold))
}

func (this *QStyleHints) MouseQuickSelectionThreshold() int {
	return (int)(QStyleHints_MouseQuickSelectionThreshold(this.h))
}

func (this *QStyleHints) ColorScheme() ColorScheme {
	return (ColorScheme)(QStyleHints_ColorScheme(this.h))
}

func (this *QStyleHints) SetColorScheme(scheme ColorScheme) {
	QStyleHints_SetColorScheme(this.h, (int)(scheme))
}

func (this *QStyleHints) UnsetColorScheme() {
	QStyleHints_UnsetColorScheme(this.h)
}

func (this *QStyleHints) CursorFlashTimeChanged(cursorFlashTime int) {
	QStyleHints_CursorFlashTimeChanged(this.h, (int)(cursorFlashTime))
}

func (this *QStyleHints) OnCursorFlashTimeChanged(slot func(cursorFlashTime int)) {
	QStyleHints_connect_CursorFlashTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_CursorFlashTimeChanged
func miqt_exec_callback_QStyleHints_CursorFlashTimeChanged(cb intptr_t, cursorFlashTime int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(cursorFlashTime int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(cursorFlashTime)

	gofunc(slotval1)
}

func (this *QStyleHints) KeyboardInputIntervalChanged(keyboardInputInterval int) {
	QStyleHints_KeyboardInputIntervalChanged(this.h, (int)(keyboardInputInterval))
}

func (this *QStyleHints) OnKeyboardInputIntervalChanged(slot func(keyboardInputInterval int)) {
	QStyleHints_connect_KeyboardInputIntervalChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_KeyboardInputIntervalChanged
func miqt_exec_callback_QStyleHints_KeyboardInputIntervalChanged(cb intptr_t, keyboardInputInterval int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(keyboardInputInterval int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(keyboardInputInterval)

	gofunc(slotval1)
}

func (this *QStyleHints) MouseDoubleClickIntervalChanged(mouseDoubleClickInterval int) {
	QStyleHints_MouseDoubleClickIntervalChanged(this.h, (int)(mouseDoubleClickInterval))
}

func (this *QStyleHints) OnMouseDoubleClickIntervalChanged(slot func(mouseDoubleClickInterval int)) {
	QStyleHints_connect_MouseDoubleClickIntervalChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_MouseDoubleClickIntervalChanged
func miqt_exec_callback_QStyleHints_MouseDoubleClickIntervalChanged(cb intptr_t, mouseDoubleClickInterval int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(mouseDoubleClickInterval int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(mouseDoubleClickInterval)

	gofunc(slotval1)
}

func (this *QStyleHints) MousePressAndHoldIntervalChanged(mousePressAndHoldInterval int) {
	QStyleHints_MousePressAndHoldIntervalChanged(this.h, (int)(mousePressAndHoldInterval))
}

func (this *QStyleHints) OnMousePressAndHoldIntervalChanged(slot func(mousePressAndHoldInterval int)) {
	QStyleHints_connect_MousePressAndHoldIntervalChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_MousePressAndHoldIntervalChanged
func miqt_exec_callback_QStyleHints_MousePressAndHoldIntervalChanged(cb intptr_t, mousePressAndHoldInterval int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(mousePressAndHoldInterval int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(mousePressAndHoldInterval)

	gofunc(slotval1)
}

func (this *QStyleHints) StartDragDistanceChanged(startDragDistance int) {
	QStyleHints_StartDragDistanceChanged(this.h, (int)(startDragDistance))
}

func (this *QStyleHints) OnStartDragDistanceChanged(slot func(startDragDistance int)) {
	QStyleHints_connect_StartDragDistanceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_StartDragDistanceChanged
func miqt_exec_callback_QStyleHints_StartDragDistanceChanged(cb intptr_t, startDragDistance int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(startDragDistance int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(startDragDistance)

	gofunc(slotval1)
}

func (this *QStyleHints) StartDragTimeChanged(startDragTime int) {
	QStyleHints_StartDragTimeChanged(this.h, (int)(startDragTime))
}

func (this *QStyleHints) OnStartDragTimeChanged(slot func(startDragTime int)) {
	QStyleHints_connect_StartDragTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_StartDragTimeChanged
func miqt_exec_callback_QStyleHints_StartDragTimeChanged(cb intptr_t, startDragTime int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(startDragTime int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(startDragTime)

	gofunc(slotval1)
}

func (this *QStyleHints) TabFocusBehaviorChanged(tabFocusBehavior TabFocusBehavior) {
	QStyleHints_TabFocusBehaviorChanged(this.h, (int)(tabFocusBehavior))
}

func (this *QStyleHints) OnTabFocusBehaviorChanged(slot func(tabFocusBehavior TabFocusBehavior)) {
	QStyleHints_connect_TabFocusBehaviorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_TabFocusBehaviorChanged
func miqt_exec_callback_QStyleHints_TabFocusBehaviorChanged(cb intptr_t, tabFocusBehavior int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(tabFocusBehavior TabFocusBehavior))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (TabFocusBehavior)(tabFocusBehavior)

	gofunc(slotval1)
}

func (this *QStyleHints) UseHoverEffectsChanged(useHoverEffects bool) {
	QStyleHints_UseHoverEffectsChanged(this.h, (bool)(useHoverEffects))
}

func (this *QStyleHints) OnUseHoverEffectsChanged(slot func(useHoverEffects bool)) {
	QStyleHints_connect_UseHoverEffectsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_UseHoverEffectsChanged
func miqt_exec_callback_QStyleHints_UseHoverEffectsChanged(cb intptr_t, useHoverEffects bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(useHoverEffects bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(useHoverEffects)

	gofunc(slotval1)
}

func (this *QStyleHints) ShowShortcutsInContextMenusChanged(param1 bool) {
	QStyleHints_ShowShortcutsInContextMenusChanged(this.h, (bool)(param1))
}

func (this *QStyleHints) OnShowShortcutsInContextMenusChanged(slot func(param1 bool)) {
	QStyleHints_connect_ShowShortcutsInContextMenusChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_ShowShortcutsInContextMenusChanged
func miqt_exec_callback_QStyleHints_ShowShortcutsInContextMenusChanged(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QStyleHints) ContextMenuTriggerChanged(contextMenuTrigger ContextMenuTrigger) {
	QStyleHints_ContextMenuTriggerChanged(this.h, (int)(contextMenuTrigger))
}

func (this *QStyleHints) OnContextMenuTriggerChanged(slot func(contextMenuTrigger ContextMenuTrigger)) {
	QStyleHints_connect_ContextMenuTriggerChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_ContextMenuTriggerChanged
func miqt_exec_callback_QStyleHints_ContextMenuTriggerChanged(cb intptr_t, contextMenuTrigger int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(contextMenuTrigger ContextMenuTrigger))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ContextMenuTrigger)(contextMenuTrigger)

	gofunc(slotval1)
}

func (this *QStyleHints) WheelScrollLinesChanged(scrollLines int) {
	QStyleHints_WheelScrollLinesChanged(this.h, (int)(scrollLines))
}

func (this *QStyleHints) OnWheelScrollLinesChanged(slot func(scrollLines int)) {
	QStyleHints_connect_WheelScrollLinesChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_WheelScrollLinesChanged
func miqt_exec_callback_QStyleHints_WheelScrollLinesChanged(cb intptr_t, scrollLines int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(scrollLines int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(scrollLines)

	gofunc(slotval1)
}

func (this *QStyleHints) MouseQuickSelectionThresholdChanged(threshold int) {
	QStyleHints_MouseQuickSelectionThresholdChanged(this.h, (int)(threshold))
}

func (this *QStyleHints) OnMouseQuickSelectionThresholdChanged(slot func(threshold int)) {
	QStyleHints_connect_MouseQuickSelectionThresholdChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_MouseQuickSelectionThresholdChanged
func miqt_exec_callback_QStyleHints_MouseQuickSelectionThresholdChanged(cb intptr_t, threshold int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(threshold int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(threshold)

	gofunc(slotval1)
}

func (this *QStyleHints) ColorSchemeChanged(colorScheme ColorScheme) {
	QStyleHints_ColorSchemeChanged(this.h, (int)(colorScheme))
}

func (this *QStyleHints) OnColorSchemeChanged(slot func(colorScheme ColorScheme)) {
	QStyleHints_connect_ColorSchemeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyleHints_ColorSchemeChanged
func miqt_exec_callback_QStyleHints_ColorSchemeChanged(cb intptr_t, colorScheme int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(colorScheme ColorScheme))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (ColorScheme)(colorScheme)

	gofunc(slotval1)
}

func QStyleHints_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStyleHints_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStyleHints_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStyleHints_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
