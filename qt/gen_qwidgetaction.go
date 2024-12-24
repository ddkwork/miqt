package qt

import (
	"unsafe"
)

type QWidgetAction struct {
	h          uintptr
	isSubclass bool
}

// NewQWidgetAction constructs a new QWidgetAction object.
func NewQWidgetAction(parent *QObject) *QWidgetAction {
	g := newQWidgetAction(QWidgetAction_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QWidgetAction) MetaObject() *QMetaObject {
	return newQMetaObject(QWidgetAction_MetaObject(this.h))
}

func (this *QWidgetAction) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWidgetAction_Metacast(this.h, param1_Cstring))
}

func QWidgetAction_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWidgetAction_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidgetAction) SetDefaultWidget(w *QWidget) {
	QWidgetAction_SetDefaultWidget(this.h, w.cPointer())
}

func (this *QWidgetAction) DefaultWidget() *QWidget {
	return newQWidget(QWidgetAction_DefaultWidget(this.h))
}

func (this *QWidgetAction) RequestWidget(parent *QWidget) *QWidget {
	return newQWidget(QWidgetAction_RequestWidget(this.h, parent.cPointer()))
}

func (this *QWidgetAction) ReleaseWidget(widget *QWidget) {
	QWidgetAction_ReleaseWidget(this.h, widget.cPointer())
}

func QWidgetAction_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWidgetAction_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWidgetAction_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWidgetAction_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWidgetAction) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QWidgetAction_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QWidgetAction) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetAction_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetAction_MetaObject
func miqt_exec_callback_QWidgetAction_MetaObject(self QWidgetAction, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWidgetAction{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QWidgetAction) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWidgetAction_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWidgetAction) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetAction_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetAction_Metacast
func miqt_exec_callback_QWidgetAction_Metacast(self QWidgetAction, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QWidgetAction{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
