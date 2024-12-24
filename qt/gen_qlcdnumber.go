package qt

import (
	"unsafe"
)

type QLCDNumber__Mode int

const (
	QLCDNumber__Hex QLCDNumber__Mode = 0
	QLCDNumber__Dec QLCDNumber__Mode = 1
	QLCDNumber__Oct QLCDNumber__Mode = 2
	QLCDNumber__Bin QLCDNumber__Mode = 3
)

type QLCDNumber__SegmentStyle int

const (
	QLCDNumber__Outline QLCDNumber__SegmentStyle = 0
	QLCDNumber__Filled  QLCDNumber__SegmentStyle = 1
	QLCDNumber__Flat    QLCDNumber__SegmentStyle = 2
)

type QLCDNumber struct {
	h          uintptr
	isSubclass bool
}

// NewQLCDNumber constructs a new QLCDNumber object.
func NewQLCDNumber(parent *QWidget) *QLCDNumber {
	g := newQLCDNumber(QLCDNumber_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQLCDNumber2 constructs a new QLCDNumber object.
func NewQLCDNumber2() *QLCDNumber {
	g := newQLCDNumber(QLCDNumber_new2())
	g.isSubclass = true
	return g
}

// NewQLCDNumber3 constructs a new QLCDNumber object.
func NewQLCDNumber3(numDigits uint) *QLCDNumber {
	g := newQLCDNumber(QLCDNumber_new3((uint)(numDigits)))
	g.isSubclass = true
	return g
}

// NewQLCDNumber4 constructs a new QLCDNumber object.
func NewQLCDNumber4(numDigits uint, parent *QWidget) *QLCDNumber {
	g := newQLCDNumber(QLCDNumber_new4((uint)(numDigits), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QLCDNumber) MetaObject() *QMetaObject {
	return newQMetaObject(QLCDNumber_MetaObject(this.h))
}

func (this *QLCDNumber) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLCDNumber_Metacast(this.h, param1_Cstring))
}

func QLCDNumber_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLCDNumber_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLCDNumber) SmallDecimalPoint() bool {
	return (bool)(QLCDNumber_SmallDecimalPoint(this.h))
}

func (this *QLCDNumber) DigitCount() int {
	return (int)(QLCDNumber_DigitCount(this.h))
}

func (this *QLCDNumber) SetDigitCount(nDigits int) {
	QLCDNumber_SetDigitCount(this.h, (int)(nDigits))
}

func (this *QLCDNumber) CheckOverflow(num float64) bool {
	return (bool)(QLCDNumber_CheckOverflow(this.h, (double)(num)))
}

func (this *QLCDNumber) CheckOverflowWithNum(num int) bool {
	return (bool)(QLCDNumber_CheckOverflowWithNum(this.h, (int)(num)))
}

func (this *QLCDNumber) Mode() Mode {
	xxxxxxxxx
}

func (this *QLCDNumber) SetMode(mode Mode) {
	QLCDNumber_SetMode(this.h, mode)
}

func (this *QLCDNumber) SegmentStyle() SegmentStyle {
	xxxxxxxxx
}

func (this *QLCDNumber) SetSegmentStyle(segmentStyle SegmentStyle) {
	QLCDNumber_SetSegmentStyle(this.h, segmentStyle)
}

func (this *QLCDNumber) Value() float64 {
	return (float64)(QLCDNumber_Value(this.h))
}

func (this *QLCDNumber) IntValue() int {
	return (int)(QLCDNumber_IntValue(this.h))
}

func (this *QLCDNumber) SizeHint() *QSize {
	_goptr := newQSize(QLCDNumber_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLCDNumber) Display(str string) {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	QLCDNumber_Display(this.h, str_ms)
}

func (this *QLCDNumber) DisplayWithNum(num int) {
	QLCDNumber_DisplayWithNum(this.h, (int)(num))
}

func (this *QLCDNumber) Display2(num float64) {
	QLCDNumber_Display2(this.h, (double)(num))
}

func (this *QLCDNumber) SetHexMode() {
	QLCDNumber_SetHexMode(this.h)
}

func (this *QLCDNumber) SetDecMode() {
	QLCDNumber_SetDecMode(this.h)
}

func (this *QLCDNumber) SetOctMode() {
	QLCDNumber_SetOctMode(this.h)
}

func (this *QLCDNumber) SetBinMode() {
	QLCDNumber_SetBinMode(this.h)
}

func (this *QLCDNumber) SetSmallDecimalPoint(smallDecimalPoint bool) {
	QLCDNumber_SetSmallDecimalPoint(this.h, (bool)(smallDecimalPoint))
}

func (this *QLCDNumber) Overflow() {
	QLCDNumber_Overflow(this.h)
}

func (this *QLCDNumber) OnOverflow(slot func()) {
	QLCDNumber_connect_Overflow(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLCDNumber_Overflow
func miqt_exec_callback_QLCDNumber_Overflow(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QLCDNumber_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLCDNumber_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLCDNumber_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLCDNumber_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLCDNumber) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QLCDNumber_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QLCDNumber) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLCDNumber_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLCDNumber_MetaObject
func miqt_exec_callback_QLCDNumber_MetaObject(self QLCDNumber, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLCDNumber{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QLCDNumber) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLCDNumber_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLCDNumber) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLCDNumber_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLCDNumber_Metacast
func miqt_exec_callback_QLCDNumber_Metacast(self QLCDNumber, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QLCDNumber{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
