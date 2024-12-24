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
	g := newQTextBrowser(QTextBrowser_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQTextBrowser2 constructs a new QTextBrowser object.
func NewQTextBrowser2() *QTextBrowser {
	g := newQTextBrowser(QTextBrowser_new2())
	g.isSubclass = true
	return g
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

func (this *QTextBrowser) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTextBrowser_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTextBrowser) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_MetaObject
func miqt_exec_callback_QTextBrowser_MetaObject(self QTextBrowser, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTextBrowser) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTextBrowser_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTextBrowser) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextBrowser_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextBrowser_Metacast
func miqt_exec_callback_QTextBrowser_Metacast(self QTextBrowser, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QTextBrowser{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
