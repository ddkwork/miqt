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
