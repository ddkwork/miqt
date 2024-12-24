package qt

import (
	"unsafe"
)

type QCompleter__CompletionMode int

const (
	QCompleter__PopupCompletion           QCompleter__CompletionMode = 0
	QCompleter__UnfilteredPopupCompletion QCompleter__CompletionMode = 1
	QCompleter__InlineCompletion          QCompleter__CompletionMode = 2
)

type QCompleter__ModelSorting int

const (
	QCompleter__UnsortedModel                QCompleter__ModelSorting = 0
	QCompleter__CaseSensitivelySortedModel   QCompleter__ModelSorting = 1
	QCompleter__CaseInsensitivelySortedModel QCompleter__ModelSorting = 2
)

type QCompleter struct {
	h          uintptr
	isSubclass bool
}

// NewQCompleter constructs a new QCompleter object.
func NewQCompleter() *QCompleter {
	g := newQCompleter(QCompleter_new())
	g.isSubclass = true
	return g
}

// NewQCompleter2 constructs a new QCompleter object.
func NewQCompleter2(model *QAbstractItemModel) *QCompleter {
	g := newQCompleter(QCompleter_new2(model.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCompleter3 constructs a new QCompleter object.
func NewQCompleter3(completions []string) *QCompleter {
	completions_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(completions))))
	defer free(unsafe.Pointer(completions_CArray))
	for i := range completions {
		completions_i_ms := struct_miqt_string{}
		completions_i_ms.data = CString(completions[i])
		completions_i_ms.len = size_t(len(completions[i]))
		defer free(unsafe.Pointer(completions_i_ms.data))
		completions_CArray[i] = completions_i_ms
	}
	completions_ma := struct_miqt_array{len: size_t(len(completions)), data: unsafe.Pointer(completions_CArray)}
	g := newQCompleter(QCompleter_new3(completions_ma))
	g.isSubclass = true
	return g
}

// NewQCompleter4 constructs a new QCompleter object.
func NewQCompleter4(parent *QObject) *QCompleter {
	g := newQCompleter(QCompleter_new4(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCompleter5 constructs a new QCompleter object.
func NewQCompleter5(model *QAbstractItemModel, parent *QObject) *QCompleter {
	g := newQCompleter(QCompleter_new5(model.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCompleter6 constructs a new QCompleter object.
func NewQCompleter6(completions []string, parent *QObject) *QCompleter {
	completions_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(completions))))
	defer free(unsafe.Pointer(completions_CArray))
	for i := range completions {
		completions_i_ms := struct_miqt_string{}
		completions_i_ms.data = CString(completions[i])
		completions_i_ms.len = size_t(len(completions[i]))
		defer free(unsafe.Pointer(completions_i_ms.data))
		completions_CArray[i] = completions_i_ms
	}
	completions_ma := struct_miqt_array{len: size_t(len(completions)), data: unsafe.Pointer(completions_CArray)}
	g := newQCompleter(QCompleter_new6(completions_ma, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QCompleter) MetaObject() *QMetaObject {
	return newQMetaObject(QCompleter_MetaObject(this.h))
}

func (this *QCompleter) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCompleter_Metacast(this.h, param1_Cstring))
}

func QCompleter_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCompleter_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) SetWidget(widget *QWidget) {
	QCompleter_SetWidget(this.h, widget.cPointer())
}

func (this *QCompleter) Widget() *QWidget {
	return newQWidget(QCompleter_Widget(this.h))
}

func (this *QCompleter) SetModel(c *QAbstractItemModel) {
	QCompleter_SetModel(this.h, c.cPointer())
}

func (this *QCompleter) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QCompleter_Model(this.h))
}

func (this *QCompleter) SetCompletionMode(mode CompletionMode) {
	QCompleter_SetCompletionMode(this.h, mode)
}

func (this *QCompleter) CompletionMode() CompletionMode {
	xxxxxxxxx
}

func (this *QCompleter) SetFilterMode(filterMode MatchFlag) {
	QCompleter_SetFilterMode(this.h, (int)(filterMode))
}

func (this *QCompleter) FilterMode() MatchFlag {
	return (MatchFlag)(QCompleter_FilterMode(this.h))
}

func (this *QCompleter) Popup() *QAbstractItemView {
	return newQAbstractItemView(QCompleter_Popup(this.h))
}

func (this *QCompleter) SetPopup(popup *QAbstractItemView) {
	QCompleter_SetPopup(this.h, popup.cPointer())
}

func (this *QCompleter) SetCaseSensitivity(caseSensitivity CaseSensitivity) {
	QCompleter_SetCaseSensitivity(this.h, (int)(caseSensitivity))
}

func (this *QCompleter) CaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(QCompleter_CaseSensitivity(this.h))
}

func (this *QCompleter) SetModelSorting(sorting ModelSorting) {
	QCompleter_SetModelSorting(this.h, sorting)
}

func (this *QCompleter) ModelSorting() ModelSorting {
	xxxxxxxxx
}

func (this *QCompleter) SetCompletionColumn(column int) {
	QCompleter_SetCompletionColumn(this.h, (int)(column))
}

func (this *QCompleter) CompletionColumn() int {
	return (int)(QCompleter_CompletionColumn(this.h))
}

func (this *QCompleter) SetCompletionRole(role int) {
	QCompleter_SetCompletionRole(this.h, (int)(role))
}

func (this *QCompleter) CompletionRole() int {
	return (int)(QCompleter_CompletionRole(this.h))
}

func (this *QCompleter) WrapAround() bool {
	return (bool)(QCompleter_WrapAround(this.h))
}

func (this *QCompleter) MaxVisibleItems() int {
	return (int)(QCompleter_MaxVisibleItems(this.h))
}

func (this *QCompleter) SetMaxVisibleItems(maxItems int) {
	QCompleter_SetMaxVisibleItems(this.h, (int)(maxItems))
}

func (this *QCompleter) CompletionCount() int {
	return (int)(QCompleter_CompletionCount(this.h))
}

func (this *QCompleter) SetCurrentRow(row int) bool {
	return (bool)(QCompleter_SetCurrentRow(this.h, (int)(row)))
}

func (this *QCompleter) CurrentRow() int {
	return (int)(QCompleter_CurrentRow(this.h))
}

func (this *QCompleter) CurrentIndex() *QModelIndex {
	_goptr := newQModelIndex(QCompleter_CurrentIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCompleter) CurrentCompletion() string {
	var _ms struct_miqt_string = QCompleter_CurrentCompletion(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) CompletionModel() *QAbstractItemModel {
	return newQAbstractItemModel(QCompleter_CompletionModel(this.h))
}

func (this *QCompleter) CompletionPrefix() string {
	var _ms struct_miqt_string = QCompleter_CompletionPrefix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) SetCompletionPrefix(prefix string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	QCompleter_SetCompletionPrefix(this.h, prefix_ms)
}

func (this *QCompleter) Complete() {
	QCompleter_Complete(this.h)
}

func (this *QCompleter) SetWrapAround(wrap bool) {
	QCompleter_SetWrapAround(this.h, (bool)(wrap))
}

func (this *QCompleter) PathFromIndex(index *QModelIndex) string {
	var _ms struct_miqt_string = QCompleter_PathFromIndex(this.h, index.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) SplitPath(path string) []string {
	path_ms := struct_miqt_string{}
	path_ms.data = CString(path)
	path_ms.len = size_t(len(path))
	defer free(unsafe.Pointer(path_ms.data))
	var _ma struct_miqt_array = QCompleter_SplitPath(this.h, path_ms)
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

func (this *QCompleter) Activated(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QCompleter_Activated(this.h, text_ms)
}

func (this *QCompleter) OnActivated(slot func(text string)) {
	QCompleter_connect_Activated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_Activated
func miqt_exec_callback_QCompleter_Activated(cb intptr_t, text struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func (this *QCompleter) ActivatedWithIndex(index *QModelIndex) {
	QCompleter_ActivatedWithIndex(this.h, index.cPointer())
}

func (this *QCompleter) OnActivatedWithIndex(slot func(index *QModelIndex)) {
	QCompleter_connect_ActivatedWithIndex(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_ActivatedWithIndex
func miqt_exec_callback_QCompleter_ActivatedWithIndex(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func (this *QCompleter) Highlighted(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QCompleter_Highlighted(this.h, text_ms)
}

func (this *QCompleter) OnHighlighted(slot func(text string)) {
	QCompleter_connect_Highlighted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_Highlighted
func miqt_exec_callback_QCompleter_Highlighted(cb intptr_t, text struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)
}

func (this *QCompleter) HighlightedWithIndex(index *QModelIndex) {
	QCompleter_HighlightedWithIndex(this.h, index.cPointer())
}

func (this *QCompleter) OnHighlightedWithIndex(slot func(index *QModelIndex)) {
	QCompleter_connect_HighlightedWithIndex(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_HighlightedWithIndex
func miqt_exec_callback_QCompleter_HighlightedWithIndex(cb intptr_t, index *QModelIndex) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index *QModelIndex))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQModelIndex(index)

	gofunc(slotval1)
}

func QCompleter_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCompleter_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCompleter_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCompleter_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCompleter) Complete1(rect *QRect) {
	QCompleter_Complete1(this.h, rect.cPointer())
}

func (this *QCompleter) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QCompleter_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QCompleter) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCompleter_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_MetaObject
func miqt_exec_callback_QCompleter_MetaObject(self QCompleter, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCompleter{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QCompleter) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QCompleter_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QCompleter) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCompleter_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCompleter_Metacast
func miqt_exec_callback_QCompleter_Metacast(self QCompleter, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QCompleter{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
