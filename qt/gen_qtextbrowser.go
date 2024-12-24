package qt

import (
	"unsafe"
)

type QTextBrowser struct {
	h          uintptr
	isSubclass bool
}

// NewQTextBrowser constructs a new QTextBrowser object.
func NewQTextBrowser(parent *QWidget) *QTextBrowser {

	ret := newQTextBrowser(QTextBrowser_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextBrowser2 constructs a new QTextBrowser object.
func NewQTextBrowser2() *QTextBrowser {

	ret := newQTextBrowser(QTextBrowser_new2())
	ret.isSubclass = true
	return ret
}

func (this *QTextBrowser) MetaObject() *QMetaObject {
	return newQMetaObject(QTextBrowser_MetaObject(this.h))
}

func (this *QTextBrowser) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextBrowser_Metacast(this.h, param1_Cstring))
}

func QTextBrowser_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextBrowser_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextBrowser) Source() *QUrl {
	_goptr := newQUrl(QTextBrowser_Source(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBrowser) SourceType() QTextDocument__ResourceType {
	return (QTextDocument__ResourceType)(QTextBrowser_SourceType(this.h))
}

func (this *QTextBrowser) SearchPaths() []string {
	var _ma struct_miqt_array = QTextBrowser_SearchPaths(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QTextBrowser) SetSearchPaths(paths []string) {
	paths_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(paths))))
	defer free(unsafe.Pointer(paths_CArray))
	for i := range paths {
		paths_i_ms := struct_miqt_string{}
		paths_i_ms.data = CString(paths[i])
		paths_i_ms.len = size_t(len(paths[i]))
		defer free(unsafe.Pointer(paths_i_ms.data))
		paths_CArray[i] = paths_i_ms
	}
	paths_ma := struct_miqt_array{len: size_t(len(paths)), data: unsafe.Pointer(paths_CArray)}
	QTextBrowser_SetSearchPaths(this.h, paths_ma)
}

func (this *QTextBrowser) LoadResource(typeVal int, name *QUrl) *QVariant {
	_goptr := newQVariant(QTextBrowser_LoadResource(this.h, (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBrowser) IsBackwardAvailable() bool {
	return (bool)(QTextBrowser_IsBackwardAvailable(this.h))
}

func (this *QTextBrowser) IsForwardAvailable() bool {
	return (bool)(QTextBrowser_IsForwardAvailable(this.h))
}

func (this *QTextBrowser) ClearHistory() {
	QTextBrowser_ClearHistory(this.h)
}

func (this *QTextBrowser) HistoryTitle(param1 int) string {
	var _ms struct_miqt_string = QTextBrowser_HistoryTitle(this.h, (int)(param1))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextBrowser) HistoryUrl(param1 int) *QUrl {
	_goptr := newQUrl(QTextBrowser_HistoryUrl(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBrowser) BackwardHistoryCount() int {
	return (int)(QTextBrowser_BackwardHistoryCount(this.h))
}

func (this *QTextBrowser) ForwardHistoryCount() int {
	return (int)(QTextBrowser_ForwardHistoryCount(this.h))
}

func (this *QTextBrowser) OpenExternalLinks() bool {
	return (bool)(QTextBrowser_OpenExternalLinks(this.h))
}

func (this *QTextBrowser) SetOpenExternalLinks(open bool) {
	QTextBrowser_SetOpenExternalLinks(this.h, (bool)(open))
}

func (this *QTextBrowser) OpenLinks() bool {
	return (bool)(QTextBrowser_OpenLinks(this.h))
}

func (this *QTextBrowser) SetOpenLinks(open bool) {
	QTextBrowser_SetOpenLinks(this.h, (bool)(open))
}

func (this *QTextBrowser) SetSource(name *QUrl) {
	QTextBrowser_SetSource(this.h, name.cPointer())
}

func (this *QTextBrowser) Backward() {
	QTextBrowser_Backward(this.h)
}

func (this *QTextBrowser) Forward() {
	QTextBrowser_Forward(this.h)
}

func (this *QTextBrowser) Home() {
	QTextBrowser_Home(this.h)
}

func (this *QTextBrowser) Reload() {
	QTextBrowser_Reload(this.h)
}

func (this *QTextBrowser) BackwardAvailable(param1 bool) {
	QTextBrowser_BackwardAvailable(this.h, (bool)(param1))
}
func (this *QTextBrowser) OnBackwardAvailable(slot func(param1 bool)) {
	QTextBrowser_connect_BackwardAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_BackwardAvailable
func miqt_exec_callback_QTextBrowser_BackwardAvailable(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QTextBrowser) ForwardAvailable(param1 bool) {
	QTextBrowser_ForwardAvailable(this.h, (bool)(param1))
}
func (this *QTextBrowser) OnForwardAvailable(slot func(param1 bool)) {
	QTextBrowser_connect_ForwardAvailable(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_ForwardAvailable
func miqt_exec_callback_QTextBrowser_ForwardAvailable(cb intptr_t, param1 bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(param1)

	gofunc(slotval1)
}

func (this *QTextBrowser) HistoryChanged() {
	QTextBrowser_HistoryChanged(this.h)
}
func (this *QTextBrowser) OnHistoryChanged(slot func()) {
	QTextBrowser_connect_HistoryChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_HistoryChanged
func miqt_exec_callback_QTextBrowser_HistoryChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QTextBrowser) SourceChanged(param1 *QUrl) {
	QTextBrowser_SourceChanged(this.h, param1.cPointer())
}
func (this *QTextBrowser) OnSourceChanged(slot func(param1 *QUrl)) {
	QTextBrowser_connect_SourceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_SourceChanged
func miqt_exec_callback_QTextBrowser_SourceChanged(cb intptr_t, param1 *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(param1)

	gofunc(slotval1)
}

func (this *QTextBrowser) Highlighted(param1 *QUrl) {
	QTextBrowser_Highlighted(this.h, param1.cPointer())
}
func (this *QTextBrowser) OnHighlighted(slot func(param1 *QUrl)) {
	QTextBrowser_connect_Highlighted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Highlighted
func miqt_exec_callback_QTextBrowser_Highlighted(cb intptr_t, param1 *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(param1)

	gofunc(slotval1)
}

func (this *QTextBrowser) AnchorClicked(param1 *QUrl) {
	QTextBrowser_AnchorClicked(this.h, param1.cPointer())
}
func (this *QTextBrowser) OnAnchorClicked(slot func(param1 *QUrl)) {
	QTextBrowser_connect_AnchorClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_AnchorClicked
func miqt_exec_callback_QTextBrowser_AnchorClicked(cb intptr_t, param1 *QUrl) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QUrl))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(param1)

	gofunc(slotval1)
}

func QTextBrowser_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextBrowser_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextBrowser_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextBrowser_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextBrowser) SetSource2(name *QUrl, typeVal QTextDocument__ResourceType) {
	QTextBrowser_SetSource2(this.h, name.cPointer(), (int)(typeVal))
}

func (this *QTextBrowser) callVirtualBase_LoadResource(typeVal int, name *QUrl) *QVariant {

	_goptr := newQVariant(QTextBrowser_virtualbase_LoadResource(unsafe.Pointer(this.h), (int)(typeVal), name.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextBrowser) OnLoadResource(slot func(super func(typeVal int, name *QUrl) *QVariant, typeVal int, name *QUrl) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_LoadResource(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_LoadResource
func miqt_exec_callback_QTextBrowser_LoadResource(self QTextBrowser, cb intptr_t, typeVal int, name *QUrl) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(typeVal int, name *QUrl) *QVariant, typeVal int, name *QUrl) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(typeVal)

	slotval2 := newQUrl(name)

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_LoadResource, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QTextBrowser) callVirtualBase_Backward() {

	QTextBrowser_virtualbase_Backward(unsafe.Pointer(this.h))

}
func (this *QTextBrowser) OnBackward(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_Backward(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Backward
func miqt_exec_callback_QTextBrowser_Backward(self QTextBrowser, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTextBrowser{h: self}).callVirtualBase_Backward)

}

func (this *QTextBrowser) callVirtualBase_Forward() {

	QTextBrowser_virtualbase_Forward(unsafe.Pointer(this.h))

}
func (this *QTextBrowser) OnForward(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_Forward(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Forward
func miqt_exec_callback_QTextBrowser_Forward(self QTextBrowser, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTextBrowser{h: self}).callVirtualBase_Forward)

}

func (this *QTextBrowser) callVirtualBase_Home() {

	QTextBrowser_virtualbase_Home(unsafe.Pointer(this.h))

}
func (this *QTextBrowser) OnHome(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_Home(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Home
func miqt_exec_callback_QTextBrowser_Home(self QTextBrowser, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTextBrowser{h: self}).callVirtualBase_Home)

}

func (this *QTextBrowser) callVirtualBase_Reload() {

	QTextBrowser_virtualbase_Reload(unsafe.Pointer(this.h))

}
func (this *QTextBrowser) OnReload(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_Reload(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Reload
func miqt_exec_callback_QTextBrowser_Reload(self QTextBrowser, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTextBrowser{h: self}).callVirtualBase_Reload)

}

func (this *QTextBrowser) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QTextBrowser_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QTextBrowser) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Event
func miqt_exec_callback_QTextBrowser_Event(self QTextBrowser, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextBrowser) callVirtualBase_KeyPressEvent(ev *QKeyEvent) {

	QTextBrowser_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QTextBrowser) OnKeyPressEvent(slot func(super func(ev *QKeyEvent), ev *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_KeyPressEvent
func miqt_exec_callback_QTextBrowser_KeyPressEvent(self QTextBrowser, cb intptr_t, ev *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QKeyEvent), ev *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(ev)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_MouseMoveEvent(ev *QMouseEvent) {

	QTextBrowser_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QTextBrowser) OnMouseMoveEvent(slot func(super func(ev *QMouseEvent), ev *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_MouseMoveEvent
func miqt_exec_callback_QTextBrowser_MouseMoveEvent(self QTextBrowser, cb intptr_t, ev *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QMouseEvent), ev *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(ev)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_MousePressEvent(ev *QMouseEvent) {

	QTextBrowser_virtualbase_MousePressEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QTextBrowser) OnMousePressEvent(slot func(super func(ev *QMouseEvent), ev *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_MousePressEvent
func miqt_exec_callback_QTextBrowser_MousePressEvent(self QTextBrowser, cb intptr_t, ev *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QMouseEvent), ev *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(ev)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_MouseReleaseEvent(ev *QMouseEvent) {

	QTextBrowser_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QTextBrowser) OnMouseReleaseEvent(slot func(super func(ev *QMouseEvent), ev *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_MouseReleaseEvent
func miqt_exec_callback_QTextBrowser_MouseReleaseEvent(self QTextBrowser, cb intptr_t, ev *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QMouseEvent), ev *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(ev)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_FocusOutEvent(ev *QFocusEvent) {

	QTextBrowser_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), ev.cPointer())

}
func (this *QTextBrowser) OnFocusOutEvent(slot func(super func(ev *QFocusEvent), ev *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_FocusOutEvent
func miqt_exec_callback_QTextBrowser_FocusOutEvent(self QTextBrowser, cb intptr_t, ev *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(ev *QFocusEvent), ev *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(ev)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTextBrowser_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTextBrowser) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_FocusNextPrevChild
func miqt_exec_callback_QTextBrowser_FocusNextPrevChild(self QTextBrowser, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextBrowser) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QTextBrowser_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_PaintEvent
func miqt_exec_callback_QTextBrowser_PaintEvent(self QTextBrowser, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_DoSetSource(name *QUrl, typeVal QTextDocument__ResourceType) {

	QTextBrowser_virtualbase_DoSetSource(unsafe.Pointer(this.h), name.cPointer(), (int)(typeVal))

}
func (this *QTextBrowser) OnDoSetSource(slot func(super func(name *QUrl, typeVal QTextDocument__ResourceType), name *QUrl, typeVal QTextDocument__ResourceType)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_DoSetSource(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_DoSetSource
func miqt_exec_callback_QTextBrowser_DoSetSource(self QTextBrowser, cb intptr_t, name *QUrl, typeVal int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(name *QUrl, typeVal QTextDocument__ResourceType), name *QUrl, typeVal QTextDocument__ResourceType))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQUrl(name)

	slotval2 := (QTextDocument__ResourceType)(typeVal)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_DoSetSource, slotval1, slotval2)

}

func (this *QTextBrowser) callVirtualBase_InputMethodQuery(property InputMethodQuery) *QVariant {

	_goptr := newQVariant(QTextBrowser_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(property)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTextBrowser) OnInputMethodQuery(slot func(super func(property InputMethodQuery) *QVariant, property InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_InputMethodQuery
func miqt_exec_callback_QTextBrowser_InputMethodQuery(self QTextBrowser, cb intptr_t, property int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(property InputMethodQuery) *QVariant, property InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(property)

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTextBrowser) callVirtualBase_TimerEvent(e *QTimerEvent) {

	QTextBrowser_virtualbase_TimerEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnTimerEvent(slot func(super func(e *QTimerEvent), e *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_TimerEvent
func miqt_exec_callback_QTextBrowser_TimerEvent(self QTextBrowser, cb intptr_t, e *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QTimerEvent), e *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	QTextBrowser_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_KeyReleaseEvent
func miqt_exec_callback_QTextBrowser_KeyReleaseEvent(self QTextBrowser, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QTextBrowser_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_ResizeEvent
func miqt_exec_callback_QTextBrowser_ResizeEvent(self QTextBrowser, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_MouseDoubleClickEvent(e *QMouseEvent) {

	QTextBrowser_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnMouseDoubleClickEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_MouseDoubleClickEvent
func miqt_exec_callback_QTextBrowser_MouseDoubleClickEvent(self QTextBrowser, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_ContextMenuEvent(e *QContextMenuEvent) {

	QTextBrowser_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnContextMenuEvent(slot func(super func(e *QContextMenuEvent), e *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_ContextMenuEvent
func miqt_exec_callback_QTextBrowser_ContextMenuEvent(self QTextBrowser, cb intptr_t, e *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QContextMenuEvent), e *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_DragEnterEvent(e *QDragEnterEvent) {

	QTextBrowser_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnDragEnterEvent(slot func(super func(e *QDragEnterEvent), e *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_DragEnterEvent
func miqt_exec_callback_QTextBrowser_DragEnterEvent(self QTextBrowser, cb intptr_t, e *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragEnterEvent), e *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_DragLeaveEvent(e *QDragLeaveEvent) {

	QTextBrowser_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnDragLeaveEvent(slot func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_DragLeaveEvent
func miqt_exec_callback_QTextBrowser_DragLeaveEvent(self QTextBrowser, cb intptr_t, e *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragLeaveEvent), e *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_DragMoveEvent(e *QDragMoveEvent) {

	QTextBrowser_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnDragMoveEvent(slot func(super func(e *QDragMoveEvent), e *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_DragMoveEvent
func miqt_exec_callback_QTextBrowser_DragMoveEvent(self QTextBrowser, cb intptr_t, e *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDragMoveEvent), e *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_DropEvent(e *QDropEvent) {

	QTextBrowser_virtualbase_DropEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnDropEvent(slot func(super func(e *QDropEvent), e *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_DropEvent
func miqt_exec_callback_QTextBrowser_DropEvent(self QTextBrowser, cb intptr_t, e *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QDropEvent), e *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	QTextBrowser_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_FocusInEvent
func miqt_exec_callback_QTextBrowser_FocusInEvent(self QTextBrowser, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QTextBrowser_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTextBrowser) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_ShowEvent
func miqt_exec_callback_QTextBrowser_ShowEvent(self QTextBrowser, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_ChangeEvent(e *QEvent) {

	QTextBrowser_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_ChangeEvent
func miqt_exec_callback_QTextBrowser_ChangeEvent(self QTextBrowser, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QTextBrowser_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QTextBrowser) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_WheelEvent
func miqt_exec_callback_QTextBrowser_WheelEvent(self QTextBrowser, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_CreateMimeDataFromSelection() *QMimeData {

	return newQMimeData(QTextBrowser_virtualbase_CreateMimeDataFromSelection(unsafe.Pointer(this.h)))

}
func (this *QTextBrowser) OnCreateMimeDataFromSelection(slot func(super func() *QMimeData) *QMimeData) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_CreateMimeDataFromSelection(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_CreateMimeDataFromSelection
func miqt_exec_callback_QTextBrowser_CreateMimeDataFromSelection(self QTextBrowser, cb intptr_t) *QMimeData {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMimeData) *QMimeData)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_CreateMimeDataFromSelection)

	return virtualReturn.cPointer()

}

func (this *QTextBrowser) callVirtualBase_CanInsertFromMimeData(source *QMimeData) bool {

	return (bool)(QTextBrowser_virtualbase_CanInsertFromMimeData(unsafe.Pointer(this.h), source.cPointer()))

}
func (this *QTextBrowser) OnCanInsertFromMimeData(slot func(super func(source *QMimeData) bool, source *QMimeData) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_CanInsertFromMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_CanInsertFromMimeData
func miqt_exec_callback_QTextBrowser_CanInsertFromMimeData(self QTextBrowser, cb intptr_t, source *QMimeData) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(source *QMimeData) bool, source *QMimeData) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(source)

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_CanInsertFromMimeData, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTextBrowser) callVirtualBase_InsertFromMimeData(source *QMimeData) {

	QTextBrowser_virtualbase_InsertFromMimeData(unsafe.Pointer(this.h), source.cPointer())

}
func (this *QTextBrowser) OnInsertFromMimeData(slot func(super func(source *QMimeData), source *QMimeData)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_InsertFromMimeData(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_InsertFromMimeData
func miqt_exec_callback_QTextBrowser_InsertFromMimeData(self QTextBrowser, cb intptr_t, source *QMimeData) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(source *QMimeData), source *QMimeData))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMimeData(source)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_InsertFromMimeData, slotval1)

}

func (this *QTextBrowser) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QTextBrowser_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QTextBrowser) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_InputMethodEvent
func miqt_exec_callback_QTextBrowser_InputMethodEvent(self QTextBrowser, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QTextBrowser) callVirtualBase_ScrollContentsBy(dx int, dy int) {

	QTextBrowser_virtualbase_ScrollContentsBy(unsafe.Pointer(this.h), (int)(dx), (int)(dy))

}
func (this *QTextBrowser) OnScrollContentsBy(slot func(super func(dx int, dy int), dx int, dy int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_ScrollContentsBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_ScrollContentsBy
func miqt_exec_callback_QTextBrowser_ScrollContentsBy(self QTextBrowser, cb intptr_t, dx int, dy int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dx int, dy int), dx int, dy int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(dx)

	slotval2 := (int)(dy)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_ScrollContentsBy, slotval1, slotval2)

}

func (this *QTextBrowser) callVirtualBase_DoSetTextCursor(cursor *QTextCursor) {

	QTextBrowser_virtualbase_DoSetTextCursor(unsafe.Pointer(this.h), cursor.cPointer())

}
func (this *QTextBrowser) OnDoSetTextCursor(slot func(super func(cursor *QTextCursor), cursor *QTextCursor)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_DoSetTextCursor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_DoSetTextCursor
func miqt_exec_callback_QTextBrowser_DoSetTextCursor(self QTextBrowser, cb intptr_t, cursor *QTextCursor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(cursor *QTextCursor), cursor *QTextCursor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTextCursor(cursor)

	gofunc((&QTextBrowser{h: self}).callVirtualBase_DoSetTextCursor, slotval1)

}
