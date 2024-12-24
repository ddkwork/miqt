package qt

import (
	"unsafe"
)

type QComboBox__InsertPolicy int

const (
	QComboBox__NoInsert             QComboBox__InsertPolicy = 0
	QComboBox__InsertAtTop          QComboBox__InsertPolicy = 1
	QComboBox__InsertAtCurrent      QComboBox__InsertPolicy = 2
	QComboBox__InsertAtBottom       QComboBox__InsertPolicy = 3
	QComboBox__InsertAfterCurrent   QComboBox__InsertPolicy = 4
	QComboBox__InsertBeforeCurrent  QComboBox__InsertPolicy = 5
	QComboBox__InsertAlphabetically QComboBox__InsertPolicy = 6
)

type QComboBox__SizeAdjustPolicy int

const (
	QComboBox__AdjustToContents                      QComboBox__SizeAdjustPolicy = 0
	QComboBox__AdjustToContentsOnFirstShow           QComboBox__SizeAdjustPolicy = 1
	QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = 2
)

type QComboBox__LabelDrawingMode int

const (
	QComboBox__UseStyle    QComboBox__LabelDrawingMode = 0
	QComboBox__UseDelegate QComboBox__LabelDrawingMode = 1
)

type QComboBox struct {
	h          uintptr
	isSubclass bool
}

// NewQComboBox constructs a new QComboBox object.
func NewQComboBox(parent *QWidget) *QComboBox {
	g := newQComboBox(QComboBox_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQComboBox2 constructs a new QComboBox object.
func NewQComboBox2() *QComboBox {
	g := newQComboBox(QComboBox_new2())
	g.isSubclass = true
	return g
}

func (this *QComboBox) MetaObject() *QMetaObject {
	return newQMetaObject(QComboBox_MetaObject(this.h))
}

func (this *QComboBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QComboBox_Metacast(this.h, param1_Cstring))
}

func QComboBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QComboBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QComboBox) MaxVisibleItems() int {
	return (int)(QComboBox_MaxVisibleItems(this.h))
}

func (this *QComboBox) SetMaxVisibleItems(maxItems int) {
	QComboBox_SetMaxVisibleItems(this.h, (int)(maxItems))
}

func (this *QComboBox) Count() int {
	return (int)(QComboBox_Count(this.h))
}

func (this *QComboBox) SetMaxCount(max int) {
	QComboBox_SetMaxCount(this.h, (int)(max))
}

func (this *QComboBox) MaxCount() int {
	return (int)(QComboBox_MaxCount(this.h))
}

func (this *QComboBox) DuplicatesEnabled() bool {
	return (bool)(QComboBox_DuplicatesEnabled(this.h))
}

func (this *QComboBox) SetDuplicatesEnabled(enable bool) {
	QComboBox_SetDuplicatesEnabled(this.h, (bool)(enable))
}

func (this *QComboBox) SetFrame(frame bool) {
	QComboBox_SetFrame(this.h, (bool)(frame))
}

func (this *QComboBox) HasFrame() bool {
	return (bool)(QComboBox_HasFrame(this.h))
}

func (this *QComboBox) FindText(text string) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QComboBox_FindText(this.h, text_ms))
}

func (this *QComboBox) FindData(data *QVariant) int {
	return (int)(QComboBox_FindData(this.h, data.cPointer()))
}

func (this *QComboBox) InsertPolicy() InsertPolicy {
	xxxxxxxxx
}

func (this *QComboBox) SetInsertPolicy(policy InsertPolicy) {
	QComboBox_SetInsertPolicy(this.h, policy)
}

func (this *QComboBox) SizeAdjustPolicy() SizeAdjustPolicy {
	xxxxxxxxx
}

func (this *QComboBox) SetSizeAdjustPolicy(policy SizeAdjustPolicy) {
	QComboBox_SetSizeAdjustPolicy(this.h, policy)
}

func (this *QComboBox) MinimumContentsLength() int {
	return (int)(QComboBox_MinimumContentsLength(this.h))
}

func (this *QComboBox) SetMinimumContentsLength(characters int) {
	QComboBox_SetMinimumContentsLength(this.h, (int)(characters))
}

func (this *QComboBox) IconSize() *QSize {
	_goptr := newQSize(QComboBox_IconSize(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) SetIconSize(size *QSize) {
	QComboBox_SetIconSize(this.h, size.cPointer())
}

func (this *QComboBox) SetPlaceholderText(placeholderText string) {
	placeholderText_ms := struct_miqt_string{}
	placeholderText_ms.data = CString(placeholderText)
	placeholderText_ms.len = size_t(len(placeholderText))
	defer free(unsafe.Pointer(placeholderText_ms.data))
	QComboBox_SetPlaceholderText(this.h, placeholderText_ms)
}

func (this *QComboBox) PlaceholderText() string {
	var _ms struct_miqt_string = QComboBox_PlaceholderText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QComboBox) IsEditable() bool {
	return (bool)(QComboBox_IsEditable(this.h))
}

func (this *QComboBox) SetEditable(editable bool) {
	QComboBox_SetEditable(this.h, (bool)(editable))
}

func (this *QComboBox) SetLineEdit(edit *QLineEdit) {
	QComboBox_SetLineEdit(this.h, edit.cPointer())
}

func (this *QComboBox) LineEdit() *QLineEdit {
	return newQLineEdit(QComboBox_LineEdit(this.h))
}

func (this *QComboBox) SetValidator(v *QValidator) {
	QComboBox_SetValidator(this.h, v.cPointer())
}

func (this *QComboBox) Validator() *QValidator {
	return newQValidator(QComboBox_Validator(this.h))
}

func (this *QComboBox) SetCompleter(c *QCompleter) {
	QComboBox_SetCompleter(this.h, c.cPointer())
}

func (this *QComboBox) Completer() *QCompleter {
	return newQCompleter(QComboBox_Completer(this.h))
}

func (this *QComboBox) ItemDelegate() *QAbstractItemDelegate {
	return newQAbstractItemDelegate(QComboBox_ItemDelegate(this.h))
}

func (this *QComboBox) SetItemDelegate(delegate *QAbstractItemDelegate) {
	QComboBox_SetItemDelegate(this.h, delegate.cPointer())
}

func (this *QComboBox) Model() *QAbstractItemModel {
	return newQAbstractItemModel(QComboBox_Model(this.h))
}

func (this *QComboBox) SetModel(model *QAbstractItemModel) {
	QComboBox_SetModel(this.h, model.cPointer())
}

func (this *QComboBox) RootModelIndex() *QModelIndex {
	_goptr := newQModelIndex(QComboBox_RootModelIndex(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) SetRootModelIndex(index *QModelIndex) {
	QComboBox_SetRootModelIndex(this.h, index.cPointer())
}

func (this *QComboBox) ModelColumn() int {
	return (int)(QComboBox_ModelColumn(this.h))
}

func (this *QComboBox) SetModelColumn(visibleColumn int) {
	QComboBox_SetModelColumn(this.h, (int)(visibleColumn))
}

func (this *QComboBox) LabelDrawingMode() LabelDrawingMode {
	xxxxxxxxx
}

func (this *QComboBox) SetLabelDrawingMode(labelDrawing LabelDrawingMode) {
	QComboBox_SetLabelDrawingMode(this.h, labelDrawing)
}

func (this *QComboBox) CurrentIndex() int {
	return (int)(QComboBox_CurrentIndex(this.h))
}

func (this *QComboBox) CurrentText() string {
	var _ms struct_miqt_string = QComboBox_CurrentText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QComboBox) CurrentData() *QVariant {
	_goptr := newQVariant(QComboBox_CurrentData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) ItemText(index int) string {
	var _ms struct_miqt_string = QComboBox_ItemText(this.h, (int)(index))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QComboBox) ItemIcon(index int) *QIcon {
	_goptr := newQIcon(QComboBox_ItemIcon(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) ItemData(index int) *QVariant {
	_goptr := newQVariant(QComboBox_ItemData(this.h, (int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) AddItem(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_AddItem(this.h, text_ms)
}

func (this *QComboBox) AddItem2(icon *QIcon, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_AddItem2(this.h, icon.cPointer(), text_ms)
}

func (this *QComboBox) AddItems(texts []string) {
	texts_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(texts))))
	defer free(unsafe.Pointer(texts_CArray))
	for i := range texts {
		texts_i_ms := struct_miqt_string{}
		texts_i_ms.data = CString(texts[i])
		texts_i_ms.len = size_t(len(texts[i]))
		defer free(unsafe.Pointer(texts_i_ms.data))
		texts_CArray[i] = texts_i_ms
	}
	texts_ma := struct_miqt_array{len: size_t(len(texts)), data: unsafe.Pointer(texts_CArray)}
	QComboBox_AddItems(this.h, texts_ma)
}

func (this *QComboBox) InsertItem(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_InsertItem(this.h, (int)(index), text_ms)
}

func (this *QComboBox) InsertItem2(index int, icon *QIcon, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_InsertItem2(this.h, (int)(index), icon.cPointer(), text_ms)
}

func (this *QComboBox) InsertItems(index int, texts []string) {
	texts_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(texts))))
	defer free(unsafe.Pointer(texts_CArray))
	for i := range texts {
		texts_i_ms := struct_miqt_string{}
		texts_i_ms.data = CString(texts[i])
		texts_i_ms.len = size_t(len(texts[i]))
		defer free(unsafe.Pointer(texts_i_ms.data))
		texts_CArray[i] = texts_i_ms
	}
	texts_ma := struct_miqt_array{len: size_t(len(texts)), data: unsafe.Pointer(texts_CArray)}
	QComboBox_InsertItems(this.h, (int)(index), texts_ma)
}

func (this *QComboBox) InsertSeparator(index int) {
	QComboBox_InsertSeparator(this.h, (int)(index))
}

func (this *QComboBox) RemoveItem(index int) {
	QComboBox_RemoveItem(this.h, (int)(index))
}

func (this *QComboBox) SetItemText(index int, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_SetItemText(this.h, (int)(index), text_ms)
}

func (this *QComboBox) SetItemIcon(index int, icon *QIcon) {
	QComboBox_SetItemIcon(this.h, (int)(index), icon.cPointer())
}

func (this *QComboBox) SetItemData(index int, value *QVariant) {
	QComboBox_SetItemData(this.h, (int)(index), value.cPointer())
}

func (this *QComboBox) View() *QAbstractItemView {
	return newQAbstractItemView(QComboBox_View(this.h))
}

func (this *QComboBox) SetView(itemView *QAbstractItemView) {
	QComboBox_SetView(this.h, itemView.cPointer())
}

func (this *QComboBox) SizeHint() *QSize {
	_goptr := newQSize(QComboBox_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) MinimumSizeHint() *QSize {
	_goptr := newQSize(QComboBox_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) ShowPopup() {
	QComboBox_ShowPopup(this.h)
}

func (this *QComboBox) HidePopup() {
	QComboBox_HidePopup(this.h)
}

func (this *QComboBox) Event(event *QEvent) bool {
	return (bool)(QComboBox_Event(this.h, event.cPointer()))
}

func (this *QComboBox) InputMethodQuery(param1 InputMethodQuery) *QVariant {
	_goptr := newQVariant(QComboBox_InputMethodQuery(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) InputMethodQuery2(query InputMethodQuery, argument *QVariant) *QVariant {
	_goptr := newQVariant(QComboBox_InputMethodQuery2(this.h, (int)(query), argument.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) Clear() {
	QComboBox_Clear(this.h)
}

func (this *QComboBox) ClearEditText() {
	QComboBox_ClearEditText(this.h)
}

func (this *QComboBox) SetEditText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_SetEditText(this.h, text_ms)
}

func (this *QComboBox) SetCurrentIndex(index int) {
	QComboBox_SetCurrentIndex(this.h, (int)(index))
}

func (this *QComboBox) SetCurrentText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_SetCurrentText(this.h, text_ms)
}

func (this *QComboBox) EditTextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QComboBox_EditTextChanged(this.h, param1_ms)
}

func (this *QComboBox) OnEditTextChanged(slot func(param1 string)) {
	QComboBox_connect_EditTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_EditTextChanged
func miqt_exec_callback_QComboBox_EditTextChanged(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func (this *QComboBox) Activated(index int) {
	QComboBox_Activated(this.h, (int)(index))
}

func (this *QComboBox) OnActivated(slot func(index int)) {
	QComboBox_connect_Activated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_Activated
func miqt_exec_callback_QComboBox_Activated(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QComboBox) TextActivated(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QComboBox_TextActivated(this.h, param1_ms)
}

func (this *QComboBox) OnTextActivated(slot func(param1 string)) {
	QComboBox_connect_TextActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_TextActivated
func miqt_exec_callback_QComboBox_TextActivated(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func (this *QComboBox) Highlighted(index int) {
	QComboBox_Highlighted(this.h, (int)(index))
}

func (this *QComboBox) OnHighlighted(slot func(index int)) {
	QComboBox_connect_Highlighted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_Highlighted
func miqt_exec_callback_QComboBox_Highlighted(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QComboBox) TextHighlighted(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QComboBox_TextHighlighted(this.h, param1_ms)
}

func (this *QComboBox) OnTextHighlighted(slot func(param1 string)) {
	QComboBox_connect_TextHighlighted(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_TextHighlighted
func miqt_exec_callback_QComboBox_TextHighlighted(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func (this *QComboBox) CurrentIndexChanged(index int) {
	QComboBox_CurrentIndexChanged(this.h, (int)(index))
}

func (this *QComboBox) OnCurrentIndexChanged(slot func(index int)) {
	QComboBox_connect_CurrentIndexChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_CurrentIndexChanged
func miqt_exec_callback_QComboBox_CurrentIndexChanged(cb intptr_t, index int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(index int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(index)

	gofunc(slotval1)
}

func (this *QComboBox) CurrentTextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QComboBox_CurrentTextChanged(this.h, param1_ms)
}

func (this *QComboBox) OnCurrentTextChanged(slot func(param1 string)) {
	QComboBox_connect_CurrentTextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_CurrentTextChanged
func miqt_exec_callback_QComboBox_CurrentTextChanged(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func QComboBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QComboBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QComboBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QComboBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QComboBox) FindText2(text string, flags MatchFlag) int {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (int)(QComboBox_FindText2(this.h, text_ms, (int)(flags)))
}

func (this *QComboBox) FindData2(data *QVariant, role int) int {
	return (int)(QComboBox_FindData2(this.h, data.cPointer(), (int)(role)))
}

func (this *QComboBox) FindData3(data *QVariant, role int, flags MatchFlag) int {
	return (int)(QComboBox_FindData3(this.h, data.cPointer(), (int)(role), (int)(flags)))
}

func (this *QComboBox) CurrentData1(role int) *QVariant {
	_goptr := newQVariant(QComboBox_CurrentData1(this.h, (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) ItemData2(index int, role int) *QVariant {
	_goptr := newQVariant(QComboBox_ItemData2(this.h, (int)(index), (int)(role)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QComboBox) AddItem22(text string, userData *QVariant) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_AddItem22(this.h, text_ms, userData.cPointer())
}

func (this *QComboBox) AddItem3(icon *QIcon, text string, userData *QVariant) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_AddItem3(this.h, icon.cPointer(), text_ms, userData.cPointer())
}

func (this *QComboBox) InsertItem3(index int, text string, userData *QVariant) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_InsertItem3(this.h, (int)(index), text_ms, userData.cPointer())
}

func (this *QComboBox) InsertItem4(index int, icon *QIcon, text string, userData *QVariant) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QComboBox_InsertItem4(this.h, (int)(index), icon.cPointer(), text_ms, userData.cPointer())
}

func (this *QComboBox) SetItemData3(index int, value *QVariant, role int) {
	QComboBox_SetItemData3(this.h, (int)(index), value.cPointer(), (int)(role))
}

func (this *QComboBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QComboBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QComboBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MetaObject
func miqt_exec_callback_QComboBox_MetaObject(self QComboBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QComboBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QComboBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QComboBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_Metacast
func miqt_exec_callback_QComboBox_Metacast(self QComboBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
