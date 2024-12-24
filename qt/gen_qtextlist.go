package qt

import (
	"unsafe"
)

type QTextList struct {
	h          uintptr
	isSubclass bool
}

// NewQTextList constructs a new QTextList object.
func NewQTextList(doc *QTextDocument) *QTextList {
	ret := newQTextList(QTextList_new(doc.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextList) MetaObject() *QMetaObject {
	return newQMetaObject(QTextList_MetaObject(this.h))
}

func (this *QTextList) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextList_Metacast(this.h, param1_Cstring))
}

func QTextList_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextList_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextList) Count() int {
	return (int)(QTextList_Count(this.h))
}

func (this *QTextList) Item(i int) *QTextBlock {
	_goptr := newQTextBlock(QTextList_Item(this.h, (int)(i)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextList) ItemNumber(param1 *QTextBlock) int {
	return (int)(QTextList_ItemNumber(this.h, param1.cPointer()))
}

func (this *QTextList) ItemText(param1 *QTextBlock) string {
	var _ms struct_miqt_string = QTextList_ItemText(this.h, param1.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextList) RemoveItem(i int) {
	QTextList_RemoveItem(this.h, (int)(i))
}

func (this *QTextList) Remove(param1 *QTextBlock) {
	QTextList_Remove(this.h, param1.cPointer())
}

func (this *QTextList) Add(block *QTextBlock) {
	QTextList_Add(this.h, block.cPointer())
}

func (this *QTextList) SetFormat(format *QTextListFormat) {
	QTextList_SetFormat(this.h, format.cPointer())
}

func (this *QTextList) Format() *QTextListFormat {
	_goptr := newQTextListFormat(QTextList_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTextList_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextList_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextList_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextList_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextList) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTextList_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTextList) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextList_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextList_MetaObject
func miqt_exec_callback_QTextList_MetaObject(self QTextList, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTextList{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTextList) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTextList_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTextList) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTextList_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTextList_Metacast
func miqt_exec_callback_QTextList_Metacast(self QTextList, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTextList{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
