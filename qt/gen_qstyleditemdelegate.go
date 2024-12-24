package qt

import (
	"unsafe"
)

type QStyledItemDelegate struct {
	h          uintptr
	isSubclass bool
}

// NewQStyledItemDelegate constructs a new QStyledItemDelegate object.
func NewQStyledItemDelegate() *QStyledItemDelegate {
	g := newQStyledItemDelegate(QStyledItemDelegate_new())
	g.isSubclass = true
	return g
}

// NewQStyledItemDelegate2 constructs a new QStyledItemDelegate object.
func NewQStyledItemDelegate2(parent *QObject) *QStyledItemDelegate {
	g := newQStyledItemDelegate(QStyledItemDelegate_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QStyledItemDelegate) MetaObject() *QMetaObject {
	return newQMetaObject(QStyledItemDelegate_MetaObject(this.h))
}

func (this *QStyledItemDelegate) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStyledItemDelegate_Metacast(this.h, param1_Cstring))
}

func QStyledItemDelegate_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStyledItemDelegate_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStyledItemDelegate) Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	QStyledItemDelegate_Paint(this.h, painter.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	_goptr := newQSize(QStyledItemDelegate_SizeHint(this.h, option.cPointer(), index.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QStyledItemDelegate) CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	return newQWidget(QStyledItemDelegate_CreateEditor(this.h, parent.cPointer(), option.cPointer(), index.cPointer()))
}

func (this *QStyledItemDelegate) SetEditorData(editor *QWidget, index *QModelIndex) {
	QStyledItemDelegate_SetEditorData(this.h, editor.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	QStyledItemDelegate_SetModelData(this.h, editor.cPointer(), model.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	QStyledItemDelegate_UpdateEditorGeometry(this.h, editor.cPointer(), option.cPointer(), index.cPointer())
}

func (this *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	return newQItemEditorFactory(QStyledItemDelegate_ItemEditorFactory(this.h))
}

func (this *QStyledItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory) {
	QStyledItemDelegate_SetItemEditorFactory(this.h, factory.cPointer())
}

func (this *QStyledItemDelegate) DisplayText(value *QVariant, locale *QLocale) string {
	var _ms struct_miqt_string = QStyledItemDelegate_DisplayText(this.h, value.cPointer(), locale.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStyledItemDelegate_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStyledItemDelegate_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStyledItemDelegate_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStyledItemDelegate_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStyledItemDelegate) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStyledItemDelegate_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStyledItemDelegate) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_MetaObject
func miqt_exec_callback_QStyledItemDelegate_MetaObject(self QStyledItemDelegate, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStyledItemDelegate) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStyledItemDelegate_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStyledItemDelegate) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStyledItemDelegate_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStyledItemDelegate_Metacast
func miqt_exec_callback_QStyledItemDelegate_Metacast(self QStyledItemDelegate, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QStyledItemDelegate{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
