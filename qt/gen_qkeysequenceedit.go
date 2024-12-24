package qt

import (
	"unsafe"
)

type QKeySequenceEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQKeySequenceEdit constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit(parent *QWidget) *QKeySequenceEdit {
	ret := newQKeySequenceEdit(QKeySequenceEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQKeySequenceEdit2 constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit2() *QKeySequenceEdit {
	ret := newQKeySequenceEdit(QKeySequenceEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQKeySequenceEdit3 constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit3(keySequence *QKeySequence) *QKeySequenceEdit {
	ret := newQKeySequenceEdit(QKeySequenceEdit_new3(keySequence.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQKeySequenceEdit4 constructs a new QKeySequenceEdit object.
func NewQKeySequenceEdit4(keySequence *QKeySequence, parent *QWidget) *QKeySequenceEdit {
	ret := newQKeySequenceEdit(QKeySequenceEdit_new4(keySequence.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QKeySequenceEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QKeySequenceEdit_MetaObject(this.h))
}

func (this *QKeySequenceEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QKeySequenceEdit_Metacast(this.h, param1_Cstring))
}

func QKeySequenceEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QKeySequenceEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QKeySequenceEdit) KeySequence() *QKeySequence {
	_goptr := newQKeySequence(QKeySequenceEdit_KeySequence(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QKeySequenceEdit) MaximumSequenceLength() int64 {
	return (int64)(QKeySequenceEdit_MaximumSequenceLength(this.h))
}

func (this *QKeySequenceEdit) SetClearButtonEnabled(enable bool) {
	QKeySequenceEdit_SetClearButtonEnabled(this.h, (bool)(enable))
}

func (this *QKeySequenceEdit) IsClearButtonEnabled() bool {
	return (bool)(QKeySequenceEdit_IsClearButtonEnabled(this.h))
}

func (this *QKeySequenceEdit) SetFinishingKeyCombinations(finishingKeyCombinations []QKeyCombination) {
	finishingKeyCombinations_CArray := (*[0xffff]*QKeyCombination)(malloc(size_t(8 * len(finishingKeyCombinations))))
	defer free(unsafe.Pointer(finishingKeyCombinations_CArray))
	for i := range finishingKeyCombinations {
		finishingKeyCombinations_CArray[i] = finishingKeyCombinations[i].cPointer()
	}
	finishingKeyCombinations_ma := struct_miqt_array{len: size_t(len(finishingKeyCombinations)), data: unsafe.Pointer(finishingKeyCombinations_CArray)}
	QKeySequenceEdit_SetFinishingKeyCombinations(this.h, finishingKeyCombinations_ma)
}

func (this *QKeySequenceEdit) FinishingKeyCombinations() []QKeyCombination {
	var _ma struct_miqt_array = QKeySequenceEdit_FinishingKeyCombinations(this.h)
	_ret := make([]QKeyCombination, int(_ma.len))
	_outCast := (*[0xffff]*QKeyCombination)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQKeyCombination(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QKeySequenceEdit) SetKeySequence(keySequence *QKeySequence) {
	QKeySequenceEdit_SetKeySequence(this.h, keySequence.cPointer())
}

func (this *QKeySequenceEdit) Clear() {
	QKeySequenceEdit_Clear(this.h)
}

func (this *QKeySequenceEdit) SetMaximumSequenceLength(count int64) {
	QKeySequenceEdit_SetMaximumSequenceLength(this.h, (ptrdiff_t)(count))
}

func (this *QKeySequenceEdit) EditingFinished() {
	QKeySequenceEdit_EditingFinished(this.h)
}

func (this *QKeySequenceEdit) OnEditingFinished(slot func()) {
	QKeySequenceEdit_connect_EditingFinished(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_EditingFinished
func miqt_exec_callback_QKeySequenceEdit_EditingFinished(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QKeySequenceEdit) KeySequenceChanged(keySequence *QKeySequence) {
	QKeySequenceEdit_KeySequenceChanged(this.h, keySequence.cPointer())
}

func (this *QKeySequenceEdit) OnKeySequenceChanged(slot func(keySequence *QKeySequence)) {
	QKeySequenceEdit_connect_KeySequenceChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_KeySequenceChanged
func miqt_exec_callback_QKeySequenceEdit_KeySequenceChanged(cb intptr_t, keySequence *QKeySequence) {
	gofunc, ok := cgo.Handle(cb).Value().(func(keySequence *QKeySequence))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeySequence(keySequence)

	gofunc(slotval1)
}

func QKeySequenceEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QKeySequenceEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QKeySequenceEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QKeySequenceEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QKeySequenceEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QKeySequenceEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QKeySequenceEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_MetaObject
func miqt_exec_callback_QKeySequenceEdit_MetaObject(self QKeySequenceEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QKeySequenceEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QKeySequenceEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QKeySequenceEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QKeySequenceEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QKeySequenceEdit_Metacast
func miqt_exec_callback_QKeySequenceEdit_Metacast(self QKeySequenceEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QKeySequenceEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
