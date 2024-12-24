package qt

import (
	"unsafe"
)

type QSpinBox struct {
	h          uintptr
	isSubclass bool
}

// NewQSpinBox constructs a new QSpinBox object.
func NewQSpinBox(parent *QWidget) *QSpinBox {
	g := newQSpinBox(QSpinBox_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQSpinBox2 constructs a new QSpinBox object.
func NewQSpinBox2() *QSpinBox {
	g := newQSpinBox(QSpinBox_new2())
	g.isSubclass = true
	return g
}

func (this *QSpinBox) MetaObject() *QMetaObject {
	return newQMetaObject(QSpinBox_MetaObject(this.h))
}

func (this *QSpinBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSpinBox_Metacast(this.h, param1_Cstring))
}

func QSpinBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSpinBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) Value() int {
	return (int)(QSpinBox_Value(this.h))
}

func (this *QSpinBox) Prefix() string {
	var _ms struct_miqt_string = QSpinBox_Prefix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) SetPrefix(prefix string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	QSpinBox_SetPrefix(this.h, prefix_ms)
}

func (this *QSpinBox) Suffix() string {
	var _ms struct_miqt_string = QSpinBox_Suffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) SetSuffix(suffix string) {
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	QSpinBox_SetSuffix(this.h, suffix_ms)
}

func (this *QSpinBox) CleanText() string {
	var _ms struct_miqt_string = QSpinBox_CleanText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) SingleStep() int {
	return (int)(QSpinBox_SingleStep(this.h))
}

func (this *QSpinBox) SetSingleStep(val int) {
	QSpinBox_SetSingleStep(this.h, (int)(val))
}

func (this *QSpinBox) Minimum() int {
	return (int)(QSpinBox_Minimum(this.h))
}

func (this *QSpinBox) SetMinimum(min int) {
	QSpinBox_SetMinimum(this.h, (int)(min))
}

func (this *QSpinBox) Maximum() int {
	return (int)(QSpinBox_Maximum(this.h))
}

func (this *QSpinBox) SetMaximum(max int) {
	QSpinBox_SetMaximum(this.h, (int)(max))
}

func (this *QSpinBox) SetRange(min int, max int) {
	QSpinBox_SetRange(this.h, (int)(min), (int)(max))
}

func (this *QSpinBox) StepType() StepType {
	xxxxxxxxx
}

func (this *QSpinBox) SetStepType(stepType StepType) {
	QSpinBox_SetStepType(this.h, stepType)
}

func (this *QSpinBox) DisplayIntegerBase() int {
	return (int)(QSpinBox_DisplayIntegerBase(this.h))
}

func (this *QSpinBox) SetDisplayIntegerBase(base int) {
	QSpinBox_SetDisplayIntegerBase(this.h, (int)(base))
}

func (this *QSpinBox) SetValue(val int) {
	QSpinBox_SetValue(this.h, (int)(val))
}

func (this *QSpinBox) ValueChanged(param1 int) {
	QSpinBox_ValueChanged(this.h, (int)(param1))
}

func (this *QSpinBox) OnValueChanged(slot func(param1 int)) {
	QSpinBox_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_ValueChanged
func miqt_exec_callback_QSpinBox_ValueChanged(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QSpinBox) TextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QSpinBox_TextChanged(this.h, param1_ms)
}

func (this *QSpinBox) OnTextChanged(slot func(param1 string)) {
	QSpinBox_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_TextChanged
func miqt_exec_callback_QSpinBox_TextChanged(cb intptr_t, param1 struct_miqt_string) {
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

func QSpinBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSpinBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSpinBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSpinBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSpinBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSpinBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSpinBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_MetaObject
func miqt_exec_callback_QSpinBox_MetaObject(self QSpinBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSpinBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSpinBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSpinBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSpinBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSpinBox_Metacast
func miqt_exec_callback_QSpinBox_Metacast(self QSpinBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSpinBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QDoubleSpinBox struct {
	h          uintptr
	isSubclass bool
}

// NewQDoubleSpinBox constructs a new QDoubleSpinBox object.
func NewQDoubleSpinBox(parent *QWidget) *QDoubleSpinBox {
	g := newQDoubleSpinBox(QDoubleSpinBox_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQDoubleSpinBox2 constructs a new QDoubleSpinBox object.
func NewQDoubleSpinBox2() *QDoubleSpinBox {
	g := newQDoubleSpinBox(QDoubleSpinBox_new2())
	g.isSubclass = true
	return g
}

func (this *QDoubleSpinBox) MetaObject() *QMetaObject {
	return newQMetaObject(QDoubleSpinBox_MetaObject(this.h))
}

func (this *QDoubleSpinBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDoubleSpinBox_Metacast(this.h, param1_Cstring))
}

func QDoubleSpinBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDoubleSpinBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) Value() float64 {
	return (float64)(QDoubleSpinBox_Value(this.h))
}

func (this *QDoubleSpinBox) Prefix() string {
	var _ms struct_miqt_string = QDoubleSpinBox_Prefix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) SetPrefix(prefix string) {
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	QDoubleSpinBox_SetPrefix(this.h, prefix_ms)
}

func (this *QDoubleSpinBox) Suffix() string {
	var _ms struct_miqt_string = QDoubleSpinBox_Suffix(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) SetSuffix(suffix string) {
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	QDoubleSpinBox_SetSuffix(this.h, suffix_ms)
}

func (this *QDoubleSpinBox) CleanText() string {
	var _ms struct_miqt_string = QDoubleSpinBox_CleanText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) SingleStep() float64 {
	return (float64)(QDoubleSpinBox_SingleStep(this.h))
}

func (this *QDoubleSpinBox) SetSingleStep(val float64) {
	QDoubleSpinBox_SetSingleStep(this.h, (double)(val))
}

func (this *QDoubleSpinBox) Minimum() float64 {
	return (float64)(QDoubleSpinBox_Minimum(this.h))
}

func (this *QDoubleSpinBox) SetMinimum(min float64) {
	QDoubleSpinBox_SetMinimum(this.h, (double)(min))
}

func (this *QDoubleSpinBox) Maximum() float64 {
	return (float64)(QDoubleSpinBox_Maximum(this.h))
}

func (this *QDoubleSpinBox) SetMaximum(max float64) {
	QDoubleSpinBox_SetMaximum(this.h, (double)(max))
}

func (this *QDoubleSpinBox) SetRange(min float64, max float64) {
	QDoubleSpinBox_SetRange(this.h, (double)(min), (double)(max))
}

func (this *QDoubleSpinBox) StepType() StepType {
	xxxxxxxxx
}

func (this *QDoubleSpinBox) SetStepType(stepType StepType) {
	QDoubleSpinBox_SetStepType(this.h, stepType)
}

func (this *QDoubleSpinBox) Decimals() int {
	return (int)(QDoubleSpinBox_Decimals(this.h))
}

func (this *QDoubleSpinBox) SetDecimals(prec int) {
	QDoubleSpinBox_SetDecimals(this.h, (int)(prec))
}

func (this *QDoubleSpinBox) Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	return (QValidator__State)(QDoubleSpinBox_Validate(this.h, input_ms, (*int)(unsafe.Pointer(pos))))
}

func (this *QDoubleSpinBox) ValueFromText(text string) float64 {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return (float64)(QDoubleSpinBox_ValueFromText(this.h, text_ms))
}

func (this *QDoubleSpinBox) TextFromValue(val float64) string {
	var _ms struct_miqt_string = QDoubleSpinBox_TextFromValue(this.h, (double)(val))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) Fixup(str string) {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	QDoubleSpinBox_Fixup(this.h, str_ms)
}

func (this *QDoubleSpinBox) SetValue(val float64) {
	QDoubleSpinBox_SetValue(this.h, (double)(val))
}

func (this *QDoubleSpinBox) ValueChanged(param1 float64) {
	QDoubleSpinBox_ValueChanged(this.h, (double)(param1))
}

func (this *QDoubleSpinBox) OnValueChanged(slot func(param1 float64)) {
	QDoubleSpinBox_connect_ValueChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_ValueChanged
func miqt_exec_callback_QDoubleSpinBox_ValueChanged(cb intptr_t, param1 double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(param1)

	gofunc(slotval1)
}

func (this *QDoubleSpinBox) TextChanged(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QDoubleSpinBox_TextChanged(this.h, param1_ms)
}

func (this *QDoubleSpinBox) OnTextChanged(slot func(param1 string)) {
	QDoubleSpinBox_connect_TextChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_TextChanged
func miqt_exec_callback_QDoubleSpinBox_TextChanged(cb intptr_t, param1 struct_miqt_string) {
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

func QDoubleSpinBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDoubleSpinBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDoubleSpinBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDoubleSpinBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleSpinBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDoubleSpinBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDoubleSpinBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_MetaObject
func miqt_exec_callback_QDoubleSpinBox_MetaObject(self QDoubleSpinBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDoubleSpinBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDoubleSpinBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDoubleSpinBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleSpinBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleSpinBox_Metacast
func miqt_exec_callback_QDoubleSpinBox_Metacast(self QDoubleSpinBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QDoubleSpinBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
