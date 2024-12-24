package qt

import (
	"unsafe"
)

type QApplication struct {
	h          uintptr
	isSubclass bool
}

// NewQApplication constructs a new QApplication object.
func NewQApplication(args []string) *QApplication {
	argcPtr := uintptr(unsafe.Pointer(&argc))
	var argvPtr uintptr
	if argv != nil && len(argv) > 0 {
		argvPtr = uintptr(unsafe.Pointer(&argv[0])) // 获取argv的指针
	} else {
		argvPtr = 0 // 或者使用nil
	}
	g := newQApplication(QApplication_new())
	g.isSubclass = true
	return g
}

// NewQApplication2 constructs a new QApplication object.
func NewQApplication2(args []string, param3 int) *QApplication {
	argcPtr := uintptr(unsafe.Pointer(&argc))
	var argvPtr uintptr
	if argv != nil && len(argv) > 0 {
		argvPtr = uintptr(unsafe.Pointer(&argv[0])) // 获取argv的指针
	} else {
		argvPtr = 0 // 或者使用nil
	}
	g := newQApplication(QApplication_new2((int)(param3)))
	g.isSubclass = true
	return g
}

func (this *QApplication) MetaObject() *QMetaObject {
	return newQMetaObject(QApplication_MetaObject(this.h))
}

func (this *QApplication) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QApplication_Metacast(this.h, param1_Cstring))
}

func QApplication_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QApplication_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QApplication_Style() *QStyle {
	return newQStyle(QApplication_Style())
}

func QApplication_SetStyle(style *QStyle) {
	QApplication_SetStyle(style.cPointer())
}

func QApplication_SetStyleWithStyle(style string) *QStyle {
	style_ms := struct_miqt_string{}
	style_ms.data = CString(style)
	style_ms.len = size_t(len(style))
	defer free(unsafe.Pointer(style_ms.data))
	return newQStyle(QApplication_SetStyleWithStyle(style_ms))
}

func QApplication_Palette(param1 *QWidget) *QPalette {
	_goptr := newQPalette(QApplication_Palette(param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QApplication_PaletteWithClassName(className string) *QPalette {
	className_Cstring := CString(className)
	defer free(unsafe.Pointer(className_Cstring))
	_goptr := newQPalette(QApplication_PaletteWithClassName(className_Cstring))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QApplication_SetPalette(param1 *QPalette) {
	QApplication_SetPalette(param1.cPointer())
}

func QApplication_Font() *QFont {
	_goptr := newQFont(QApplication_Font())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QApplication_FontWithQWidget(param1 *QWidget) *QFont {
	_goptr := newQFont(QApplication_FontWithQWidget(param1.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QApplication_FontWithClassName(className string) *QFont {
	className_Cstring := CString(className)
	defer free(unsafe.Pointer(className_Cstring))
	_goptr := newQFont(QApplication_FontWithClassName(className_Cstring))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QApplication_SetFont(param1 *QFont) {
	QApplication_SetFont(param1.cPointer())
}

func QApplication_FontMetrics() *QFontMetrics {
	_goptr := newQFontMetrics(QApplication_FontMetrics())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QApplication_AllWidgets() []*QWidget {
	var _ma struct_miqt_array = QApplication_AllWidgets()
	_ret := make([]*QWidget, int(_ma.len))
	_outCast := (*[0xffff]*QWidget)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQWidget(_outCast[i])
	}
	return _ret
}

func QApplication_TopLevelWidgets() []*QWidget {
	var _ma struct_miqt_array = QApplication_TopLevelWidgets()
	_ret := make([]*QWidget, int(_ma.len))
	_outCast := (*[0xffff]*QWidget)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQWidget(_outCast[i])
	}
	return _ret
}

func QApplication_ActivePopupWidget() *QWidget {
	return newQWidget(QApplication_ActivePopupWidget())
}

func QApplication_ActiveModalWidget() *QWidget {
	return newQWidget(QApplication_ActiveModalWidget())
}

func QApplication_FocusWidget() *QWidget {
	return newQWidget(QApplication_FocusWidget())
}

func QApplication_ActiveWindow() *QWidget {
	return newQWidget(QApplication_ActiveWindow())
}

func QApplication_SetActiveWindow(act *QWidget) {
	QApplication_SetActiveWindow(act.cPointer())
}

func QApplication_WidgetAt(p *QPoint) *QWidget {
	return newQWidget(QApplication_WidgetAt(p.cPointer()))
}

func QApplication_WidgetAt2(x int, y int) *QWidget {
	return newQWidget(QApplication_WidgetAt2((int)(x), (int)(y)))
}

func QApplication_TopLevelAt(p *QPoint) *QWidget {
	return newQWidget(QApplication_TopLevelAt(p.cPointer()))
}

func QApplication_TopLevelAt2(x int, y int) *QWidget {
	return newQWidget(QApplication_TopLevelAt2((int)(x), (int)(y)))
}

func QApplication_Beep() {
	QApplication_Beep()
}

func QApplication_Alert(widget *QWidget) {
	QApplication_Alert(widget.cPointer())
}

func QApplication_SetCursorFlashTime(cursorFlashTime int) {
	QApplication_SetCursorFlashTime((int)(cursorFlashTime))
}

func QApplication_CursorFlashTime() int {
	return (int)(QApplication_CursorFlashTime())
}

func QApplication_SetDoubleClickInterval(doubleClickInterval int) {
	QApplication_SetDoubleClickInterval((int)(doubleClickInterval))
}

func QApplication_DoubleClickInterval() int {
	return (int)(QApplication_DoubleClickInterval())
}

func QApplication_SetKeyboardInputInterval(keyboardInputInterval int) {
	QApplication_SetKeyboardInputInterval((int)(keyboardInputInterval))
}

func QApplication_KeyboardInputInterval() int {
	return (int)(QApplication_KeyboardInputInterval())
}

func QApplication_SetWheelScrollLines(wheelScrollLines int) {
	QApplication_SetWheelScrollLines((int)(wheelScrollLines))
}

func QApplication_WheelScrollLines() int {
	return (int)(QApplication_WheelScrollLines())
}

func QApplication_SetStartDragTime(ms int) {
	QApplication_SetStartDragTime((int)(ms))
}

func QApplication_StartDragTime() int {
	return (int)(QApplication_StartDragTime())
}

func QApplication_SetStartDragDistance(l int) {
	QApplication_SetStartDragDistance((int)(l))
}

func QApplication_StartDragDistance() int {
	return (int)(QApplication_StartDragDistance())
}

func QApplication_IsEffectEnabled(param1 UIEffect) bool {
	return (bool)(QApplication_IsEffectEnabled((int)(param1)))
}

func QApplication_SetEffectEnabled(param1 UIEffect) {
	QApplication_SetEffectEnabled((int)(param1))
}

func QApplication_Exec() int {
	return (int)(QApplication_Exec())
}

func (this *QApplication) Notify(param1 *QObject, param2 *QEvent) bool {
	return (bool)(QApplication_Notify(this.h, param1.cPointer(), param2.cPointer()))
}

func (this *QApplication) FocusChanged(old *QWidget, now *QWidget) {
	QApplication_FocusChanged(this.h, old.cPointer(), now.cPointer())
}

func (this *QApplication) OnFocusChanged(slot func(old *QWidget, now *QWidget)) {
	QApplication_connect_FocusChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplication_FocusChanged
func miqt_exec_callback_QApplication_FocusChanged(cb intptr_t, old *QWidget, now *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(old *QWidget, now *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(old)

	slotval2 := newQWidget(now)

	gofunc(slotval1, slotval2)
}

func (this *QApplication) StyleSheet() string {
	var _ms struct_miqt_string = QApplication_StyleSheet(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QApplication) AutoSipEnabled() bool {
	return (bool)(QApplication_AutoSipEnabled(this.h))
}

func (this *QApplication) SetStyleSheet(sheet string) {
	sheet_ms := struct_miqt_string{}
	sheet_ms.data = CString(sheet)
	sheet_ms.len = size_t(len(sheet))
	defer free(unsafe.Pointer(sheet_ms.data))
	QApplication_SetStyleSheet(this.h, sheet_ms)
}

func (this *QApplication) SetAutoSipEnabled(enabled bool) {
	QApplication_SetAutoSipEnabled(this.h, (bool)(enabled))
}

func QApplication_CloseAllWindows() {
	QApplication_CloseAllWindows()
}

func QApplication_AboutQt() {
	QApplication_AboutQt()
}

func QApplication_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QApplication_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QApplication_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QApplication_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QApplication_SetPalette2(param1 *QPalette, className string) {
	className_Cstring := CString(className)
	defer free(unsafe.Pointer(className_Cstring))
	QApplication_SetPalette2(param1.cPointer(), className_Cstring)
}

func QApplication_SetFont2(param1 *QFont, className string) {
	className_Cstring := CString(className)
	defer free(unsafe.Pointer(className_Cstring))
	QApplication_SetFont2(param1.cPointer(), className_Cstring)
}

func QApplication_Alert2(widget *QWidget, duration int) {
	QApplication_Alert2(widget.cPointer(), (int)(duration))
}

func QApplication_SetEffectEnabled2(param1 UIEffect, enable bool) {
	QApplication_SetEffectEnabled2((int)(param1), (bool)(enable))
}

func (this *QApplication) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QApplication_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QApplication) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QApplication_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplication_MetaObject
func miqt_exec_callback_QApplication_MetaObject(self QApplication, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QApplication{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QApplication) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QApplication_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QApplication) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QApplication_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QApplication_Metacast
func miqt_exec_callback_QApplication_Metacast(self QApplication, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QApplication{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
