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

	ret := newQComboBox(QComboBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQComboBox2 constructs a new QComboBox object.
func NewQComboBox2() *QComboBox {

	ret := newQComboBox(QComboBox_new2())
	ret.isSubclass = true
	return ret
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

func (this *QComboBox) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QComboBox_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QComboBox) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_SetModel
func miqt_exec_callback_QComboBox_SetModel(self QComboBox, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QComboBox{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QComboBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QComboBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QComboBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_SizeHint
func miqt_exec_callback_QComboBox_SizeHint(self QComboBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QComboBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QComboBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QComboBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MinimumSizeHint
func miqt_exec_callback_QComboBox_MinimumSizeHint(self QComboBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QComboBox) callVirtualBase_ShowPopup() {

	QComboBox_virtualbase_ShowPopup(unsafe.Pointer(this.h))

}
func (this *QComboBox) OnShowPopup(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_ShowPopup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_ShowPopup
func miqt_exec_callback_QComboBox_ShowPopup(self QComboBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QComboBox{h: self}).callVirtualBase_ShowPopup)

}

func (this *QComboBox) callVirtualBase_HidePopup() {

	QComboBox_virtualbase_HidePopup(unsafe.Pointer(this.h))

}
func (this *QComboBox) OnHidePopup(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_HidePopup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_HidePopup
func miqt_exec_callback_QComboBox_HidePopup(self QComboBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QComboBox{h: self}).callVirtualBase_HidePopup)

}

func (this *QComboBox) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QComboBox_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QComboBox) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_Event
func miqt_exec_callback_QComboBox_Event(self QComboBox, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QComboBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QComboBox_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QComboBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_InputMethodQuery
func miqt_exec_callback_QComboBox_InputMethodQuery(self QComboBox, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QComboBox) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	QComboBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_FocusInEvent
func miqt_exec_callback_QComboBox_FocusInEvent(self QComboBox, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_FocusOutEvent(e *QFocusEvent) {

	QComboBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnFocusOutEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_FocusOutEvent
func miqt_exec_callback_QComboBox_FocusOutEvent(self QComboBox, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_ChangeEvent(e *QEvent) {

	QComboBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_ChangeEvent
func miqt_exec_callback_QComboBox_ChangeEvent(self QComboBox, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QComboBox_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_ResizeEvent
func miqt_exec_callback_QComboBox_ResizeEvent(self QComboBox, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QComboBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_PaintEvent
func miqt_exec_callback_QComboBox_PaintEvent(self QComboBox, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_ShowEvent(e *QShowEvent) {

	QComboBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnShowEvent(slot func(super func(e *QShowEvent), e *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_ShowEvent
func miqt_exec_callback_QComboBox_ShowEvent(self QComboBox, cb intptr_t, e *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QShowEvent), e *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_HideEvent(e *QHideEvent) {

	QComboBox_virtualbase_HideEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnHideEvent(slot func(super func(e *QHideEvent), e *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_HideEvent
func miqt_exec_callback_QComboBox_HideEvent(self QComboBox, cb intptr_t, e *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QHideEvent), e *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	QComboBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MousePressEvent
func miqt_exec_callback_QComboBox_MousePressEvent(self QComboBox, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QComboBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MouseReleaseEvent
func miqt_exec_callback_QComboBox_MouseReleaseEvent(self QComboBox, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_KeyPressEvent(e *QKeyEvent) {

	QComboBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnKeyPressEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_KeyPressEvent
func miqt_exec_callback_QComboBox_KeyPressEvent(self QComboBox, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	QComboBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_KeyReleaseEvent
func miqt_exec_callback_QComboBox_KeyReleaseEvent(self QComboBox, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QComboBox_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_WheelEvent
func miqt_exec_callback_QComboBox_WheelEvent(self QComboBox, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_ContextMenuEvent(e *QContextMenuEvent) {

	QComboBox_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QComboBox) OnContextMenuEvent(slot func(super func(e *QContextMenuEvent), e *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_ContextMenuEvent
func miqt_exec_callback_QComboBox_ContextMenuEvent(self QComboBox, cb intptr_t, e *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QContextMenuEvent), e *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(e)

	gofunc((&QComboBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QComboBox_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QComboBox) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_InputMethodEvent
func miqt_exec_callback_QComboBox_InputMethodEvent(self QComboBox, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QComboBox{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_InitStyleOption(option *QStyleOptionComboBox) {

	QComboBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QComboBox) OnInitStyleOption(slot func(super func(option *QStyleOptionComboBox), option *QStyleOptionComboBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_InitStyleOption
func miqt_exec_callback_QComboBox_InitStyleOption(self QComboBox, cb intptr_t, option *QStyleOptionComboBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionComboBox), option *QStyleOptionComboBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionComboBox(option)

	gofunc((&QComboBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QComboBox) callVirtualBase_DevType() int {

	return (int)(QComboBox_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QComboBox) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_DevType
func miqt_exec_callback_QComboBox_DevType(self QComboBox, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QComboBox) callVirtualBase_SetVisible(visible bool) {

	QComboBox_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QComboBox) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_SetVisible
func miqt_exec_callback_QComboBox_SetVisible(self QComboBox, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QComboBox{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QComboBox) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QComboBox_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QComboBox) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_HeightForWidth
func miqt_exec_callback_QComboBox_HeightForWidth(self QComboBox, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QComboBox) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QComboBox_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QComboBox) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_HasHeightForWidth
func miqt_exec_callback_QComboBox_HasHeightForWidth(self QComboBox, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QComboBox) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QComboBox_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QComboBox) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_PaintEngine
func miqt_exec_callback_QComboBox_PaintEngine(self QComboBox, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QComboBox) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QComboBox_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MouseDoubleClickEvent
func miqt_exec_callback_QComboBox_MouseDoubleClickEvent(self QComboBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QComboBox_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MouseMoveEvent
func miqt_exec_callback_QComboBox_MouseMoveEvent(self QComboBox, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QComboBox_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_EnterEvent
func miqt_exec_callback_QComboBox_EnterEvent(self QComboBox, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_LeaveEvent(event *QEvent) {

	QComboBox_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_LeaveEvent
func miqt_exec_callback_QComboBox_LeaveEvent(self QComboBox, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QComboBox_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_MoveEvent
func miqt_exec_callback_QComboBox_MoveEvent(self QComboBox, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QComboBox_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_CloseEvent
func miqt_exec_callback_QComboBox_CloseEvent(self QComboBox, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QComboBox_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_TabletEvent
func miqt_exec_callback_QComboBox_TabletEvent(self QComboBox, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_ActionEvent(event *QActionEvent) {

	QComboBox_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_ActionEvent
func miqt_exec_callback_QComboBox_ActionEvent(self QComboBox, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QComboBox_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_DragEnterEvent
func miqt_exec_callback_QComboBox_DragEnterEvent(self QComboBox, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QComboBox_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_DragMoveEvent
func miqt_exec_callback_QComboBox_DragMoveEvent(self QComboBox, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QComboBox_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_DragLeaveEvent
func miqt_exec_callback_QComboBox_DragLeaveEvent(self QComboBox, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_DropEvent(event *QDropEvent) {

	QComboBox_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QComboBox) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_DropEvent
func miqt_exec_callback_QComboBox_DropEvent(self QComboBox, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QComboBox{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QComboBox) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QComboBox_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QComboBox) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_NativeEvent
func miqt_exec_callback_QComboBox_NativeEvent(self QComboBox, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QComboBox) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QComboBox_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QComboBox) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_Metric
func miqt_exec_callback_QComboBox_Metric(self QComboBox, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QComboBox) callVirtualBase_InitPainter(painter *QPainter) {

	QComboBox_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QComboBox) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_InitPainter
func miqt_exec_callback_QComboBox_InitPainter(self QComboBox, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QComboBox{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QComboBox) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QComboBox_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QComboBox) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_Redirected
func miqt_exec_callback_QComboBox_Redirected(self QComboBox, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QComboBox) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QComboBox_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QComboBox) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_SharedPainter
func miqt_exec_callback_QComboBox_SharedPainter(self QComboBox, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QComboBox) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QComboBox_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QComboBox) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QComboBox_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QComboBox_FocusNextPrevChild
func miqt_exec_callback_QComboBox_FocusNextPrevChild(self QComboBox, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QComboBox{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
