package qt

import (
	"unsafe"
)

type QInputDialog__InputDialogOption int

const (
	QInputDialog__NoButtons                    QInputDialog__InputDialogOption = 1
	QInputDialog__UseListViewForComboBoxItems  QInputDialog__InputDialogOption = 2
	QInputDialog__UsePlainTextEditForTextInput QInputDialog__InputDialogOption = 4
)

type QInputDialog__InputMode int

const (
	QInputDialog__TextInput   QInputDialog__InputMode = 0
	QInputDialog__IntInput    QInputDialog__InputMode = 1
	QInputDialog__DoubleInput QInputDialog__InputMode = 2
)

type QInputDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQInputDialog constructs a new QInputDialog object.
func NewQInputDialog(parent *QWidget) *QInputDialog {
	ret := newQInputDialog(QInputDialog_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQInputDialog2 constructs a new QInputDialog object.
func NewQInputDialog2() *QInputDialog {
	ret := newQInputDialog(QInputDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQInputDialog3 constructs a new QInputDialog object.
func NewQInputDialog3(parent *QWidget, flags WindowType) *QInputDialog {
	ret := newQInputDialog(QInputDialog_new3(parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QInputDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QInputDialog_MetaObject(this.h))
}

func (this *QInputDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QInputDialog_Metacast(this.h, param1_Cstring))
}

func QInputDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QInputDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDialog) SetInputMode(mode InputMode) {
	QInputDialog_SetInputMode(this.h, mode)
}

func (this *QInputDialog) InputMode() InputMode {
	xxxxxxxxx
}

func (this *QInputDialog) SetLabelText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QInputDialog_SetLabelText(this.h, text_ms)
}

func (this *QInputDialog) LabelText() string {
	var _ms struct_miqt_string = QInputDialog_LabelText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDialog) SetOption(option InputDialogOption) {
	QInputDialog_SetOption(this.h, option)
}

func (this *QInputDialog) TestOption(option InputDialogOption) bool {
	return (bool)(QInputDialog_TestOption(this.h, option))
}

func (this *QInputDialog) SetOptions(options InputDialogOptions) {
	QInputDialog_SetOptions(this.h, options)
}

func (this *QInputDialog) Options() InputDialogOptions {
	xxxxxxxxx
}

func (this *QInputDialog) SetTextValue(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QInputDialog_SetTextValue(this.h, text_ms)
}

func (this *QInputDialog) TextValue() string {
	var _ms struct_miqt_string = QInputDialog_TextValue(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDialog) SetTextEchoMode(mode QLineEdit__EchoMode) {
	QInputDialog_SetTextEchoMode(this.h, (int)(mode))
}

func (this *QInputDialog) TextEchoMode() QLineEdit__EchoMode {
	return (QLineEdit__EchoMode)(QInputDialog_TextEchoMode(this.h))
}

func (this *QInputDialog) SetComboBoxEditable(editable bool) {
	QInputDialog_SetComboBoxEditable(this.h, (bool)(editable))
}

func (this *QInputDialog) IsComboBoxEditable() bool {
	return (bool)(QInputDialog_IsComboBoxEditable(this.h))
}

func (this *QInputDialog) SetComboBoxItems(items []string) {
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	QInputDialog_SetComboBoxItems(this.h, items_ma)
}

func (this *QInputDialog) ComboBoxItems() []string {
	var _ma struct_miqt_array = QInputDialog_ComboBoxItems(this.h)
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

func (this *QInputDialog) SetIntValue(value int) {
	QInputDialog_SetIntValue(this.h, (int)(value))
}

func (this *QInputDialog) IntValue() int {
	return (int)(QInputDialog_IntValue(this.h))
}

func (this *QInputDialog) SetIntMinimum(min int) {
	QInputDialog_SetIntMinimum(this.h, (int)(min))
}

func (this *QInputDialog) IntMinimum() int {
	return (int)(QInputDialog_IntMinimum(this.h))
}

func (this *QInputDialog) SetIntMaximum(max int) {
	QInputDialog_SetIntMaximum(this.h, (int)(max))
}

func (this *QInputDialog) IntMaximum() int {
	return (int)(QInputDialog_IntMaximum(this.h))
}

func (this *QInputDialog) SetIntRange(min int, max int) {
	QInputDialog_SetIntRange(this.h, (int)(min), (int)(max))
}

func (this *QInputDialog) SetIntStep(step int) {
	QInputDialog_SetIntStep(this.h, (int)(step))
}

func (this *QInputDialog) IntStep() int {
	return (int)(QInputDialog_IntStep(this.h))
}

func (this *QInputDialog) SetDoubleValue(value float64) {
	QInputDialog_SetDoubleValue(this.h, (double)(value))
}

func (this *QInputDialog) DoubleValue() float64 {
	return (float64)(QInputDialog_DoubleValue(this.h))
}

func (this *QInputDialog) SetDoubleMinimum(min float64) {
	QInputDialog_SetDoubleMinimum(this.h, (double)(min))
}

func (this *QInputDialog) DoubleMinimum() float64 {
	return (float64)(QInputDialog_DoubleMinimum(this.h))
}

func (this *QInputDialog) SetDoubleMaximum(max float64) {
	QInputDialog_SetDoubleMaximum(this.h, (double)(max))
}

func (this *QInputDialog) DoubleMaximum() float64 {
	return (float64)(QInputDialog_DoubleMaximum(this.h))
}

func (this *QInputDialog) SetDoubleRange(min float64, max float64) {
	QInputDialog_SetDoubleRange(this.h, (double)(min), (double)(max))
}

func (this *QInputDialog) SetDoubleDecimals(decimals int) {
	QInputDialog_SetDoubleDecimals(this.h, (int)(decimals))
}

func (this *QInputDialog) DoubleDecimals() int {
	return (int)(QInputDialog_DoubleDecimals(this.h))
}

func (this *QInputDialog) SetOkButtonText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QInputDialog_SetOkButtonText(this.h, text_ms)
}

func (this *QInputDialog) OkButtonText() string {
	var _ms struct_miqt_string = QInputDialog_OkButtonText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDialog) SetCancelButtonText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QInputDialog_SetCancelButtonText(this.h, text_ms)
}

func (this *QInputDialog) CancelButtonText() string {
	var _ms struct_miqt_string = QInputDialog_CancelButtonText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDialog) MinimumSizeHint() *QSize {
	_goptr := newQSize(QInputDialog_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputDialog) SizeHint() *QSize {
	_goptr := newQSize(QInputDialog_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QInputDialog) SetVisible(visible bool) {
	QInputDialog_SetVisible(this.h, (bool)(visible))
}

func QInputDialog_GetText(parent *QWidget, title string, label string) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetText(parent.cPointer(), title_ms, label_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetMultiLineText(parent *QWidget, title string, label string) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetMultiLineText(parent.cPointer(), title_ms, label_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetItem(parent *QWidget, title string, label string, items []string) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	var _ms struct_miqt_string = QInputDialog_GetItem(parent.cPointer(), title_ms, label_ms, items_ma)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetInt(parent *QWidget, title string, label string) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt(parent.cPointer(), title_ms, label_ms))
}

func QInputDialog_GetDouble(parent *QWidget, title string, label string) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble(parent.cPointer(), title_ms, label_ms))
}

func (this *QInputDialog) SetDoubleStep(step float64) {
	QInputDialog_SetDoubleStep(this.h, (double)(step))
}

func (this *QInputDialog) DoubleStep() float64 {
	return (float64)(QInputDialog_DoubleStep(this.h))
}

func (this *QInputDialog) TextValueChanged(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QInputDialog_TextValueChanged(this.h, text_ms)
}

func (this *QInputDialog) OnTextValueChanged(slot func(text string)) {
	QInputDialog_connect_TextValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_TextValueChanged
func miqt_exec_callback_QInputDialog_TextValueChanged(cb intptr_t, text struct_miqt_string) {
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

func (this *QInputDialog) TextValueSelected(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QInputDialog_TextValueSelected(this.h, text_ms)
}

func (this *QInputDialog) OnTextValueSelected(slot func(text string)) {
	QInputDialog_connect_TextValueSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_TextValueSelected
func miqt_exec_callback_QInputDialog_TextValueSelected(cb intptr_t, text struct_miqt_string) {
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

func (this *QInputDialog) IntValueChanged(value int) {
	QInputDialog_IntValueChanged(this.h, (int)(value))
}

func (this *QInputDialog) OnIntValueChanged(slot func(value int)) {
	QInputDialog_connect_IntValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_IntValueChanged
func miqt_exec_callback_QInputDialog_IntValueChanged(cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc(slotval1)
}

func (this *QInputDialog) IntValueSelected(value int) {
	QInputDialog_IntValueSelected(this.h, (int)(value))
}

func (this *QInputDialog) OnIntValueSelected(slot func(value int)) {
	QInputDialog_connect_IntValueSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_IntValueSelected
func miqt_exec_callback_QInputDialog_IntValueSelected(cb intptr_t, value int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(value)

	gofunc(slotval1)
}

func (this *QInputDialog) DoubleValueChanged(value float64) {
	QInputDialog_DoubleValueChanged(this.h, (double)(value))
}

func (this *QInputDialog) OnDoubleValueChanged(slot func(value float64)) {
	QInputDialog_connect_DoubleValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_DoubleValueChanged
func miqt_exec_callback_QInputDialog_DoubleValueChanged(cb intptr_t, value double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(value)

	gofunc(slotval1)
}

func (this *QInputDialog) DoubleValueSelected(value float64) {
	QInputDialog_DoubleValueSelected(this.h, (double)(value))
}

func (this *QInputDialog) OnDoubleValueSelected(slot func(value float64)) {
	QInputDialog_connect_DoubleValueSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_DoubleValueSelected
func miqt_exec_callback_QInputDialog_DoubleValueSelected(cb intptr_t, value double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(value float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(value)

	gofunc(slotval1)
}

func (this *QInputDialog) Done(result int) {
	QInputDialog_Done(this.h, (int)(result))
}

func QInputDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QInputDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QInputDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QInputDialog) SetOption2(option InputDialogOption, on bool) {
	QInputDialog_SetOption2(this.h, option, (bool)(on))
}

func QInputDialog_GetText4(parent *QWidget, title string, label string, echo QLineEdit__EchoMode) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetText4(parent.cPointer(), title_ms, label_ms, (int)(echo))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetText5(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetText5(parent.cPointer(), title_ms, label_ms, (int)(echo), text_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetText6(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string, ok *bool) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetText6(parent.cPointer(), title_ms, label_ms, (int)(echo), text_ms, (*bool)(unsafe.Pointer(ok)))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetText7(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string, ok *bool, flags WindowType) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetText7(parent.cPointer(), title_ms, label_ms, (int)(echo), text_ms, (*bool)(unsafe.Pointer(ok)), (int)(flags))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetText8(parent *QWidget, title string, label string, echo QLineEdit__EchoMode, text string, ok *bool, flags WindowType, inputMethodHints InputMethodHint) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetText8(parent.cPointer(), title_ms, label_ms, (int)(echo), text_ms, (*bool)(unsafe.Pointer(ok)), (int)(flags), (int)(inputMethodHints))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetMultiLineText4(parent *QWidget, title string, label string, text string) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetMultiLineText4(parent.cPointer(), title_ms, label_ms, text_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetMultiLineText5(parent *QWidget, title string, label string, text string, ok *bool) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetMultiLineText5(parent.cPointer(), title_ms, label_ms, text_ms, (*bool)(unsafe.Pointer(ok)))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetMultiLineText6(parent *QWidget, title string, label string, text string, ok *bool, flags WindowType) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetMultiLineText6(parent.cPointer(), title_ms, label_ms, text_ms, (*bool)(unsafe.Pointer(ok)), (int)(flags))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetMultiLineText7(parent *QWidget, title string, label string, text string, ok *bool, flags WindowType, inputMethodHints InputMethodHint) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	var _ms struct_miqt_string = QInputDialog_GetMultiLineText7(parent.cPointer(), title_ms, label_ms, text_ms, (*bool)(unsafe.Pointer(ok)), (int)(flags), (int)(inputMethodHints))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetItem5(parent *QWidget, title string, label string, items []string, current int) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	var _ms struct_miqt_string = QInputDialog_GetItem5(parent.cPointer(), title_ms, label_ms, items_ma, (int)(current))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetItem6(parent *QWidget, title string, label string, items []string, current int, editable bool) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	var _ms struct_miqt_string = QInputDialog_GetItem6(parent.cPointer(), title_ms, label_ms, items_ma, (int)(current), (bool)(editable))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetItem7(parent *QWidget, title string, label string, items []string, current int, editable bool, ok *bool) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	var _ms struct_miqt_string = QInputDialog_GetItem7(parent.cPointer(), title_ms, label_ms, items_ma, (int)(current), (bool)(editable), (*bool)(unsafe.Pointer(ok)))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetItem8(parent *QWidget, title string, label string, items []string, current int, editable bool, ok *bool, flags WindowType) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	var _ms struct_miqt_string = QInputDialog_GetItem8(parent.cPointer(), title_ms, label_ms, items_ma, (int)(current), (bool)(editable), (*bool)(unsafe.Pointer(ok)), (int)(flags))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetItem9(parent *QWidget, title string, label string, items []string, current int, editable bool, ok *bool, flags WindowType, inputMethodHints InputMethodHint) string {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	items_CArray := (*[0xffff]struct_miqt_string)(malloc(size_t(int(unsafe.Sizeof(struct_miqt_string{})) * len(items))))
	defer free(unsafe.Pointer(items_CArray))
	for i := range items {
		items_i_ms := struct_miqt_string{}
		items_i_ms.data = CString(items[i])
		items_i_ms.len = size_t(len(items[i]))
		defer free(unsafe.Pointer(items_i_ms.data))
		items_CArray[i] = items_i_ms
	}
	items_ma := struct_miqt_array{len: size_t(len(items)), data: unsafe.Pointer(items_CArray)}
	var _ms struct_miqt_string = QInputDialog_GetItem9(parent.cPointer(), title_ms, label_ms, items_ma, (int)(current), (bool)(editable), (*bool)(unsafe.Pointer(ok)), (int)(flags), (int)(inputMethodHints))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QInputDialog_GetInt4(parent *QWidget, title string, label string, value int) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt4(parent.cPointer(), title_ms, label_ms, (int)(value)))
}

func QInputDialog_GetInt5(parent *QWidget, title string, label string, value int, minValue int) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt5(parent.cPointer(), title_ms, label_ms, (int)(value), (int)(minValue)))
}

func QInputDialog_GetInt6(parent *QWidget, title string, label string, value int, minValue int, maxValue int) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt6(parent.cPointer(), title_ms, label_ms, (int)(value), (int)(minValue), (int)(maxValue)))
}

func QInputDialog_GetInt7(parent *QWidget, title string, label string, value int, minValue int, maxValue int, step int) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt7(parent.cPointer(), title_ms, label_ms, (int)(value), (int)(minValue), (int)(maxValue), (int)(step)))
}

func QInputDialog_GetInt8(parent *QWidget, title string, label string, value int, minValue int, maxValue int, step int, ok *bool) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt8(parent.cPointer(), title_ms, label_ms, (int)(value), (int)(minValue), (int)(maxValue), (int)(step), (*bool)(unsafe.Pointer(ok))))
}

func QInputDialog_GetInt9(parent *QWidget, title string, label string, value int, minValue int, maxValue int, step int, ok *bool, flags WindowType) int {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (int)(QInputDialog_GetInt9(parent.cPointer(), title_ms, label_ms, (int)(value), (int)(minValue), (int)(maxValue), (int)(step), (*bool)(unsafe.Pointer(ok)), (int)(flags)))
}

func QInputDialog_GetDouble4(parent *QWidget, title string, label string, value float64) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble4(parent.cPointer(), title_ms, label_ms, (double)(value)))
}

func QInputDialog_GetDouble5(parent *QWidget, title string, label string, value float64, minValue float64) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble5(parent.cPointer(), title_ms, label_ms, (double)(value), (double)(minValue)))
}

func QInputDialog_GetDouble6(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble6(parent.cPointer(), title_ms, label_ms, (double)(value), (double)(minValue), (double)(maxValue)))
}

func QInputDialog_GetDouble7(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble7(parent.cPointer(), title_ms, label_ms, (double)(value), (double)(minValue), (double)(maxValue), (int)(decimals)))
}

func QInputDialog_GetDouble8(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble8(parent.cPointer(), title_ms, label_ms, (double)(value), (double)(minValue), (double)(maxValue), (int)(decimals), (*bool)(unsafe.Pointer(ok))))
}

func QInputDialog_GetDouble9(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool, flags WindowType) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble9(parent.cPointer(), title_ms, label_ms, (double)(value), (double)(minValue), (double)(maxValue), (int)(decimals), (*bool)(unsafe.Pointer(ok)), (int)(flags)))
}

func QInputDialog_GetDouble10(parent *QWidget, title string, label string, value float64, minValue float64, maxValue float64, decimals int, ok *bool, flags WindowType, step float64) float64 {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	label_ms := struct_miqt_string{}
	label_ms.data = CString(label)
	label_ms.len = size_t(len(label))
	defer free(unsafe.Pointer(label_ms.data))
	return (float64)(QInputDialog_GetDouble10(parent.cPointer(), title_ms, label_ms, (double)(value), (double)(minValue), (double)(maxValue), (int)(decimals), (*bool)(unsafe.Pointer(ok)), (int)(flags), (double)(step)))
}

func (this *QInputDialog) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QInputDialog_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QInputDialog) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputDialog_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_MetaObject
func miqt_exec_callback_QInputDialog_MetaObject(self QInputDialog, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QInputDialog{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QInputDialog) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QInputDialog_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QInputDialog) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QInputDialog_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QInputDialog_Metacast
func miqt_exec_callback_QInputDialog_Metacast(self QInputDialog, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QInputDialog{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
