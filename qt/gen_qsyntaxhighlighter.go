package qt

import (
	"unsafe"
)

type QSyntaxHighlighter struct {
	h          uintptr
	isSubclass bool
}

// NewQSyntaxHighlighter constructs a new QSyntaxHighlighter object.
func NewQSyntaxHighlighter(parent *QObject) *QSyntaxHighlighter {
	ret := newQSyntaxHighlighter(QSyntaxHighlighter_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQSyntaxHighlighter2 constructs a new QSyntaxHighlighter object.
func NewQSyntaxHighlighter2(parent *QTextDocument) *QSyntaxHighlighter {
	ret := newQSyntaxHighlighter(QSyntaxHighlighter_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QSyntaxHighlighter) MetaObject() *QMetaObject {
	return newQMetaObject(QSyntaxHighlighter_MetaObject(this.h))
}

func (this *QSyntaxHighlighter) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSyntaxHighlighter_Metacast(this.h, param1_Cstring))
}

func QSyntaxHighlighter_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSyntaxHighlighter_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSyntaxHighlighter) SetDocument(doc *QTextDocument) {
	QSyntaxHighlighter_SetDocument(this.h, doc.cPointer())
}

func (this *QSyntaxHighlighter) Document() *QTextDocument {
	return newQTextDocument(QSyntaxHighlighter_Document(this.h))
}

func (this *QSyntaxHighlighter) Rehighlight() {
	QSyntaxHighlighter_Rehighlight(this.h)
}

func (this *QSyntaxHighlighter) RehighlightBlock(block *QTextBlock) {
	QSyntaxHighlighter_RehighlightBlock(this.h, block.cPointer())
}

func QSyntaxHighlighter_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSyntaxHighlighter_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSyntaxHighlighter_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSyntaxHighlighter_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSyntaxHighlighter) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSyntaxHighlighter_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSyntaxHighlighter) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_MetaObject
func miqt_exec_callback_QSyntaxHighlighter_MetaObject(self QSyntaxHighlighter, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSyntaxHighlighter) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSyntaxHighlighter_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSyntaxHighlighter) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_Metacast
func miqt_exec_callback_QSyntaxHighlighter_Metacast(self QSyntaxHighlighter, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
