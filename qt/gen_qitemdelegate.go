package qt

import (
	"unsafe"
)

type QItemDelegate struct {
	h          uintptr
	isSubclass bool
}

// NewQItemDelegate constructs a new QItemDelegate object.
func NewQItemDelegate() *QItemDelegate {
	g := newQItemDelegate(QItemDelegate_new())
	g.isSubclass = true
	return g
}

// NewQItemDelegate2 constructs a new QItemDelegate object.
func NewQItemDelegate2(parent *QObject) *QItemDelegate {
	g := newQItemDelegate(QItemDelegate_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QItemDelegate) MetaObject() *QMetaObject {
	return newQMetaObject(QItemDelegate_MetaObject(this.h))
}

func (this *QItemDelegate) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QItemDelegate_Metacast(this.h, param1_Cstring))
}

func QItemDelegate_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QItemDelegate_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QItemDelegate) HasClipping() bool {
	return (bool)(QItemDelegate_HasClipping(this.h))
}

func (this *QItemDelegate) SetClipping(clip bool) {
	QItemDelegate_SetClipping(this.h, (bool)(clip))
}

func (this *QItemDelegate) Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	QItemDelegate_Paint(this.h, painter.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QItemDelegate) SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	_goptr := newQSize(QItemDelegate_SizeHint(this.h, option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QItemDelegate) CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	return newQWidget(QItemDelegate_CreateEditor(this.h, parent.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QItemDelegate) SetEditorData(editor *QWidget, index *QModelIndex) {
	QItemDelegate_SetEditorData(this.h, editor.cPointer(), index.cPointer())
}

func (this *QItemDelegate) SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	QItemDelegate_SetModelData(this.h, editor.cPointer(), model.cPointer(), index.cPointer())
}

func (this *QItemDelegate) UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	QItemDelegate_UpdateEditorGeometry(this.h, editor.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	return newQItemEditorFactory(QItemDelegate_ItemEditorFactory(this.h))
}

func (this *QItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory) {
	QItemDelegate_SetItemEditorFactory(this.h, factory.cPointer())
}

func QItemDelegate_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QItemDelegate_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QItemDelegate_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QItemDelegate_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QItemDelegate) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QItemDelegate_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QItemDelegate) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_MetaObject
func miqt_exec_callback_QItemDelegate_MetaObject(self QItemDelegate, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QItemDelegate) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QItemDelegate_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QItemDelegate) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemDelegate_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemDelegate_Metacast
func miqt_exec_callback_QItemDelegate_Metacast(self QItemDelegate, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QItemDelegate{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
