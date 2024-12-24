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

	ret := newQWidgetAction(QWidgetAction_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
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

func (this *QWidgetAction) callVirtualBase_Event(param1 *QEvent) bool {

	return (bool)(QWidgetAction_virtualbase_Event(unsafe.Pointer(this.h), param1.cPointer()))

}
func (this *QWidgetAction) OnEvent(slot func(super func(param1 *QEvent) bool, param1 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetAction_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetAction_Event
func miqt_exec_callback_QWidgetAction_Event(self QWidgetAction, cb intptr_t, param1 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent) bool, param1 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	virtualReturn := gofunc((&QWidgetAction{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWidgetAction) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QWidgetAction_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QWidgetAction) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetAction_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetAction_EventFilter
func miqt_exec_callback_QWidgetAction_EventFilter(self QWidgetAction, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QWidgetAction{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QWidgetAction) callVirtualBase_CreateWidget(parent *QWidget) *QWidget {

	return newQWidget(QWidgetAction_virtualbase_CreateWidget(unsafe.Pointer(this.h), parent.cPointer()))

}
func (this *QWidgetAction) OnCreateWidget(slot func(super func(parent *QWidget) *QWidget, parent *QWidget) *QWidget) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetAction_override_virtual_CreateWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetAction_CreateWidget
func miqt_exec_callback_QWidgetAction_CreateWidget(self QWidgetAction, cb intptr_t, parent *QWidget) *QWidget {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(parent *QWidget) *QWidget, parent *QWidget) *QWidget)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(parent)

	virtualReturn := gofunc((&QWidgetAction{h: self}).callVirtualBase_CreateWidget, slotval1)

	return virtualReturn.cPointer()

}

func (this *QWidgetAction) callVirtualBase_DeleteWidget(widget *QWidget) {

	QWidgetAction_virtualbase_DeleteWidget(unsafe.Pointer(this.h), widget.cPointer())

}
func (this *QWidgetAction) OnDeleteWidget(slot func(super func(widget *QWidget), widget *QWidget)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWidgetAction_override_virtual_DeleteWidget(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWidgetAction_DeleteWidget
func miqt_exec_callback_QWidgetAction_DeleteWidget(self QWidgetAction, cb intptr_t, widget *QWidget) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(widget *QWidget), widget *QWidget))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWidget(widget)

	gofunc((&QWidgetAction{h: self}).callVirtualBase_DeleteWidget, slotval1)

}
