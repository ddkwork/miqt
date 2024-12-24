package qt

import (
	"unsafe"
)

type QValidator__State int

const (
	QValidator__Invalid      QValidator__State = 0
	QValidator__Intermediate QValidator__State = 1
	QValidator__Acceptable   QValidator__State = 2
)

type QDoubleValidator__Notation int

const (
	QDoubleValidator__StandardNotation   QDoubleValidator__Notation = 0
	QDoubleValidator__ScientificNotation QDoubleValidator__Notation = 1
)

type QValidator struct {
	h          uintptr
	isSubclass bool
}

// NewQValidator constructs a new QValidator object.
func NewQValidator() *QValidator {
	ret := newQValidator(QValidator_new())
	ret.isSubclass = true
	return ret
}

// NewQValidator2 constructs a new QValidator object.
func NewQValidator2(parent *QObject) *QValidator {
	ret := newQValidator(QValidator_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QValidator) MetaObject() *QMetaObject {
	return newQMetaObject(QValidator_MetaObject(this.h))
}

func (this *QValidator) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QValidator_Metacast(this.h, param1_Cstring))
}

func QValidator_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QValidator_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QValidator) SetLocale(locale *QLocale) {
	QValidator_SetLocale(this.h, locale.cPointer())
}

func (this *QValidator) Locale() *QLocale {
	_goptr := newQLocale(QValidator_Locale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QValidator) Validate(param1 string, param2 *int) State {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	xxxxxxxxx
}

func (this *QValidator) Fixup(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QValidator_Fixup(this.h, param1_ms)
}

func (this *QValidator) Changed() {
	QValidator_Changed(this.h)
}

func (this *QValidator) OnChanged(slot func()) {
	QValidator_connect_Changed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QValidator_Changed
func miqt_exec_callback_QValidator_Changed(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QValidator_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QValidator_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QValidator_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QValidator_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QValidator) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QValidator_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QValidator) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QValidator_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QValidator_MetaObject
func miqt_exec_callback_QValidator_MetaObject(self QValidator, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QValidator{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QValidator) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QValidator_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QValidator) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QValidator_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QValidator_Metacast
func miqt_exec_callback_QValidator_Metacast(self QValidator, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QValidator{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QIntValidator struct {
	h          uintptr
	isSubclass bool
}

// NewQIntValidator constructs a new QIntValidator object.
func NewQIntValidator() *QIntValidator {
	ret := newQIntValidator(QIntValidator_new())
	ret.isSubclass = true
	return ret
}

// NewQIntValidator2 constructs a new QIntValidator object.
func NewQIntValidator2(bottom int, top int) *QIntValidator {
	ret := newQIntValidator(QIntValidator_new2((int)(bottom), (int)(top)))
	ret.isSubclass = true
	return ret
}

// NewQIntValidator3 constructs a new QIntValidator object.
func NewQIntValidator3(parent *QObject) *QIntValidator {
	ret := newQIntValidator(QIntValidator_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQIntValidator4 constructs a new QIntValidator object.
func NewQIntValidator4(bottom int, top int, parent *QObject) *QIntValidator {
	ret := newQIntValidator(QIntValidator_new4((int)(bottom), (int)(top), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QIntValidator) MetaObject() *QMetaObject {
	return newQMetaObject(QIntValidator_MetaObject(this.h))
}

func (this *QIntValidator) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QIntValidator_Metacast(this.h, param1_Cstring))
}

func QIntValidator_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QIntValidator_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIntValidator) Validate(param1 string, param2 *int) QValidator__State {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	return (QValidator__State)(QIntValidator_Validate(this.h, param1_ms, (*int)(unsafe.Pointer(param2))))
}

func (this *QIntValidator) Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	QIntValidator_Fixup(this.h, input_ms)
}

func (this *QIntValidator) SetBottom(bottom int) {
	QIntValidator_SetBottom(this.h, (int)(bottom))
}

func (this *QIntValidator) SetTop(top int) {
	QIntValidator_SetTop(this.h, (int)(top))
}

func (this *QIntValidator) SetRange(bottom int, top int) {
	QIntValidator_SetRange(this.h, (int)(bottom), (int)(top))
}

func (this *QIntValidator) Bottom() int {
	return (int)(QIntValidator_Bottom(this.h))
}

func (this *QIntValidator) Top() int {
	return (int)(QIntValidator_Top(this.h))
}

func (this *QIntValidator) BottomChanged(bottom int) {
	QIntValidator_BottomChanged(this.h, (int)(bottom))
}

func (this *QIntValidator) OnBottomChanged(slot func(bottom int)) {
	QIntValidator_connect_BottomChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIntValidator_BottomChanged
func miqt_exec_callback_QIntValidator_BottomChanged(cb intptr_t, bottom int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(bottom int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(bottom)

	gofunc(slotval1)
}

func (this *QIntValidator) TopChanged(top int) {
	QIntValidator_TopChanged(this.h, (int)(top))
}

func (this *QIntValidator) OnTopChanged(slot func(top int)) {
	QIntValidator_connect_TopChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIntValidator_TopChanged
func miqt_exec_callback_QIntValidator_TopChanged(cb intptr_t, top int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(top int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(top)

	gofunc(slotval1)
}

func QIntValidator_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIntValidator_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIntValidator_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIntValidator_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIntValidator) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QIntValidator_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QIntValidator) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIntValidator_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIntValidator_MetaObject
func miqt_exec_callback_QIntValidator_MetaObject(self QIntValidator, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIntValidator{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QIntValidator) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QIntValidator_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QIntValidator) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIntValidator_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIntValidator_Metacast
func miqt_exec_callback_QIntValidator_Metacast(self QIntValidator, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QIntValidator{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QDoubleValidator struct {
	h          uintptr
	isSubclass bool
}

// NewQDoubleValidator constructs a new QDoubleValidator object.
func NewQDoubleValidator() *QDoubleValidator {
	ret := newQDoubleValidator(QDoubleValidator_new())
	ret.isSubclass = true
	return ret
}

// NewQDoubleValidator2 constructs a new QDoubleValidator object.
func NewQDoubleValidator2(bottom float64, top float64, decimals int) *QDoubleValidator {
	ret := newQDoubleValidator(QDoubleValidator_new2((double)(bottom), (double)(top), (int)(decimals)))
	ret.isSubclass = true
	return ret
}

// NewQDoubleValidator3 constructs a new QDoubleValidator object.
func NewQDoubleValidator3(parent *QObject) *QDoubleValidator {
	ret := newQDoubleValidator(QDoubleValidator_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDoubleValidator4 constructs a new QDoubleValidator object.
func NewQDoubleValidator4(bottom float64, top float64, decimals int, parent *QObject) *QDoubleValidator {
	ret := newQDoubleValidator(QDoubleValidator_new4((double)(bottom), (double)(top), (int)(decimals), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDoubleValidator) MetaObject() *QMetaObject {
	return newQMetaObject(QDoubleValidator_MetaObject(this.h))
}

func (this *QDoubleValidator) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDoubleValidator_Metacast(this.h, param1_Cstring))
}

func QDoubleValidator_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDoubleValidator_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleValidator) Validate(param1 string, param2 *int) QValidator__State {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	return (QValidator__State)(QDoubleValidator_Validate(this.h, param1_ms, (*int)(unsafe.Pointer(param2))))
}

func (this *QDoubleValidator) Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	QDoubleValidator_Fixup(this.h, input_ms)
}

func (this *QDoubleValidator) SetRange(bottom float64, top float64, decimals int) {
	QDoubleValidator_SetRange(this.h, (double)(bottom), (double)(top), (int)(decimals))
}

func (this *QDoubleValidator) SetRange2(bottom float64, top float64) {
	QDoubleValidator_SetRange2(this.h, (double)(bottom), (double)(top))
}

func (this *QDoubleValidator) SetBottom(bottom float64) {
	QDoubleValidator_SetBottom(this.h, (double)(bottom))
}

func (this *QDoubleValidator) SetTop(top float64) {
	QDoubleValidator_SetTop(this.h, (double)(top))
}

func (this *QDoubleValidator) SetDecimals(decimals int) {
	QDoubleValidator_SetDecimals(this.h, (int)(decimals))
}

func (this *QDoubleValidator) SetNotation(notation Notation) {
	QDoubleValidator_SetNotation(this.h, notation)
}

func (this *QDoubleValidator) Bottom() float64 {
	return (float64)(QDoubleValidator_Bottom(this.h))
}

func (this *QDoubleValidator) Top() float64 {
	return (float64)(QDoubleValidator_Top(this.h))
}

func (this *QDoubleValidator) Decimals() int {
	return (int)(QDoubleValidator_Decimals(this.h))
}

func (this *QDoubleValidator) Notation() Notation {
	xxxxxxxxx
}

func (this *QDoubleValidator) BottomChanged(bottom float64) {
	QDoubleValidator_BottomChanged(this.h, (double)(bottom))
}

func (this *QDoubleValidator) OnBottomChanged(slot func(bottom float64)) {
	QDoubleValidator_connect_BottomChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleValidator_BottomChanged
func miqt_exec_callback_QDoubleValidator_BottomChanged(cb intptr_t, bottom double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(bottom float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(bottom)

	gofunc(slotval1)
}

func (this *QDoubleValidator) TopChanged(top float64) {
	QDoubleValidator_TopChanged(this.h, (double)(top))
}

func (this *QDoubleValidator) OnTopChanged(slot func(top float64)) {
	QDoubleValidator_connect_TopChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleValidator_TopChanged
func miqt_exec_callback_QDoubleValidator_TopChanged(cb intptr_t, top double) {
	gofunc, ok := cgo.Handle(cb).Value().(func(top float64))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (float64)(top)

	gofunc(slotval1)
}

func (this *QDoubleValidator) DecimalsChanged(decimals int) {
	QDoubleValidator_DecimalsChanged(this.h, (int)(decimals))
}

func (this *QDoubleValidator) OnDecimalsChanged(slot func(decimals int)) {
	QDoubleValidator_connect_DecimalsChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleValidator_DecimalsChanged
func miqt_exec_callback_QDoubleValidator_DecimalsChanged(cb intptr_t, decimals int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(decimals int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(decimals)

	gofunc(slotval1)
}

func (this *QDoubleValidator) NotationChanged(notation QDoubleValidator__Notation) {
	QDoubleValidator_NotationChanged(this.h, (int)(notation))
}

func (this *QDoubleValidator) OnNotationChanged(slot func(notation QDoubleValidator__Notation)) {
	QDoubleValidator_connect_NotationChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleValidator_NotationChanged
func miqt_exec_callback_QDoubleValidator_NotationChanged(cb intptr_t, notation int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(notation QDoubleValidator__Notation))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QDoubleValidator__Notation)(notation)

	gofunc(slotval1)
}

func QDoubleValidator_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDoubleValidator_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDoubleValidator_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDoubleValidator_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDoubleValidator) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDoubleValidator_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDoubleValidator) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleValidator_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleValidator_MetaObject
func miqt_exec_callback_QDoubleValidator_MetaObject(self QDoubleValidator, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDoubleValidator{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDoubleValidator) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDoubleValidator_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDoubleValidator) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDoubleValidator_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDoubleValidator_Metacast
func miqt_exec_callback_QDoubleValidator_Metacast(self QDoubleValidator, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QDoubleValidator{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QRegularExpressionValidator struct {
	h          uintptr
	isSubclass bool
}

// NewQRegularExpressionValidator constructs a new QRegularExpressionValidator object.
func NewQRegularExpressionValidator() *QRegularExpressionValidator {
	ret := newQRegularExpressionValidator(QRegularExpressionValidator_new())
	ret.isSubclass = true
	return ret
}

// NewQRegularExpressionValidator2 constructs a new QRegularExpressionValidator object.
func NewQRegularExpressionValidator2(re *QRegularExpression) *QRegularExpressionValidator {
	ret := newQRegularExpressionValidator(QRegularExpressionValidator_new2(re.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRegularExpressionValidator3 constructs a new QRegularExpressionValidator object.
func NewQRegularExpressionValidator3(parent *QObject) *QRegularExpressionValidator {
	ret := newQRegularExpressionValidator(QRegularExpressionValidator_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRegularExpressionValidator4 constructs a new QRegularExpressionValidator object.
func NewQRegularExpressionValidator4(re *QRegularExpression, parent *QObject) *QRegularExpressionValidator {
	ret := newQRegularExpressionValidator(QRegularExpressionValidator_new4(re.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRegularExpressionValidator) MetaObject() *QMetaObject {
	return newQMetaObject(QRegularExpressionValidator_MetaObject(this.h))
}

func (this *QRegularExpressionValidator) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QRegularExpressionValidator_Metacast(this.h, param1_Cstring))
}

func QRegularExpressionValidator_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QRegularExpressionValidator_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpressionValidator) Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))
	return (QValidator__State)(QRegularExpressionValidator_Validate(this.h, input_ms, (*int)(unsafe.Pointer(pos))))
}

func (this *QRegularExpressionValidator) RegularExpression() *QRegularExpression {
	_goptr := newQRegularExpression(QRegularExpressionValidator_RegularExpression(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpressionValidator) SetRegularExpression(re *QRegularExpression) {
	QRegularExpressionValidator_SetRegularExpression(this.h, re.cPointer())
}

func (this *QRegularExpressionValidator) RegularExpressionChanged(re *QRegularExpression) {
	QRegularExpressionValidator_RegularExpressionChanged(this.h, re.cPointer())
}

func (this *QRegularExpressionValidator) OnRegularExpressionChanged(slot func(re *QRegularExpression)) {
	QRegularExpressionValidator_connect_RegularExpressionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRegularExpressionValidator_RegularExpressionChanged
func miqt_exec_callback_QRegularExpressionValidator_RegularExpressionChanged(cb intptr_t, re *QRegularExpression) {
	gofunc, ok := cgo.Handle(cb).Value().(func(re *QRegularExpression))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQRegularExpression(re)

	gofunc(slotval1)
}

func QRegularExpressionValidator_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRegularExpressionValidator_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRegularExpressionValidator_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QRegularExpressionValidator_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpressionValidator) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QRegularExpressionValidator_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QRegularExpressionValidator) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRegularExpressionValidator_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRegularExpressionValidator_MetaObject
func miqt_exec_callback_QRegularExpressionValidator_MetaObject(self QRegularExpressionValidator, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QRegularExpressionValidator{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QRegularExpressionValidator) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QRegularExpressionValidator_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QRegularExpressionValidator) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QRegularExpressionValidator_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QRegularExpressionValidator_Metacast
func miqt_exec_callback_QRegularExpressionValidator_Metacast(self QRegularExpressionValidator, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QRegularExpressionValidator{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
