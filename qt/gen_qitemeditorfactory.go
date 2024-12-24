package qt

import (
	"unsafe"
)

type QItemEditorCreatorBase struct {
	h          uintptr
	isSubclass bool
}

func (this *QItemEditorCreatorBase) CreateWidget(parent *QWidget) *QWidget {
	return newQWidget(QItemEditorCreatorBase_CreateWidget(this.h, parent.cPointer()))
}

func (this *QItemEditorCreatorBase) ValuePropertyName() []byte {
	var _bytearray struct_miqt_string = QItemEditorCreatorBase_ValuePropertyName(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QItemEditorCreatorBase) OperatorAssign(param1 *QItemEditorCreatorBase) {
	QItemEditorCreatorBase_OperatorAssign(this.h, param1.cPointer())
}

type QItemEditorFactory struct {
	h          uintptr
	isSubclass bool
}

// NewQItemEditorFactory constructs a new QItemEditorFactory object.
func NewQItemEditorFactory() *QItemEditorFactory {

	ret := newQItemEditorFactory(QItemEditorFactory_new())
	ret.isSubclass = true
	return ret
}

// NewQItemEditorFactory2 constructs a new QItemEditorFactory object.
func NewQItemEditorFactory2(param1 *QItemEditorFactory) *QItemEditorFactory {

	ret := newQItemEditorFactory(QItemEditorFactory_new2(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QItemEditorFactory) CreateEditor(userType int, parent *QWidget) *QWidget {
	return newQWidget(QItemEditorFactory_CreateEditor(this.h, (int)(userType), parent.cPointer()))
}

func (this *QItemEditorFactory) ValuePropertyName(userType int) []byte {
	var _bytearray struct_miqt_string = QItemEditorFactory_ValuePropertyName(this.h, (int)(userType))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QItemEditorFactory) RegisterEditor(userType int, creator *QItemEditorCreatorBase) {
	QItemEditorFactory_RegisterEditor(this.h, (int)(userType), creator.cPointer())
}

func QItemEditorFactory_DefaultFactory() *QItemEditorFactory {
	return newQItemEditorFactory(QItemEditorFactory_DefaultFactory())
}

func QItemEditorFactory_SetDefaultFactory(factory *QItemEditorFactory) {
	QItemEditorFactory_SetDefaultFactory(factory.cPointer())
}

func (this *QItemEditorFactory) callVirtualBase_CreateEditor(userType int, parent *QWidget) *QWidget {

	return newQWidget(QItemEditorFactory_virtualbase_CreateEditor(unsafe.Pointer(this.h), (int)(userType), parent.cPointer()))

}
func (this *QItemEditorFactory) OnCreateEditor(slot func(super func(userType int, parent *QWidget) *QWidget, userType int, parent *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemEditorFactory_override_virtual_CreateEditor(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemEditorFactory_CreateEditor
func miqt_exec_callback_QItemEditorFactory_CreateEditor(self QItemEditorFactory, cb intptr_t, userType int, parent *QWidget) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(userType int, parent *QWidget) *QWidget, userType int, parent *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(userType)

	slotval2 := newQWidget(parent)

	virtualReturn := gofunc((&QItemEditorFactory{h: self}).callVirtualBase_CreateEditor, slotval1, slotval2)

	return virtualReturn.cPointer()

}

func (this *QItemEditorFactory) callVirtualBase_ValuePropertyName(userType int) []byte {

	var _bytearray struct_miqt_string = QItemEditorFactory_virtualbase_ValuePropertyName(unsafe.Pointer(this.h), (int)(userType))
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}
func (this *QItemEditorFactory) OnValuePropertyName(slot func(super func(userType int) []byte, userType int) []byte) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QItemEditorFactory_override_virtual_ValuePropertyName(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QItemEditorFactory_ValuePropertyName
func miqt_exec_callback_QItemEditorFactory_ValuePropertyName(self QItemEditorFactory, cb intptr_t, userType int) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(userType int) []byte, userType int) []byte)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(userType)

	virtualReturn := gofunc((&QItemEditorFactory{h: self}).callVirtualBase_ValuePropertyName, slotval1)
	virtualReturn_alias := struct_miqt_string{}
	virtualReturn_alias.data = (char)(unsafe.Pointer(&virtualReturn[0]))
	virtualReturn_alias.len = size_t(len(virtualReturn))

	return virtualReturn_alias

}
