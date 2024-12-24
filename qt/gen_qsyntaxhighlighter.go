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
func (this *QSyntaxHighlighter) OnHighlightBlock(slot func(text string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_HighlightBlock(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_HighlightBlock
func miqt_exec_callback_QSyntaxHighlighter_HighlightBlock(self QSyntaxHighlighter, cb intptr_t, text struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(text string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	gofunc(slotval1)

}

func (this *QSyntaxHighlighter) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QSyntaxHighlighter_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QSyntaxHighlighter) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_Event
func miqt_exec_callback_QSyntaxHighlighter_Event(self QSyntaxHighlighter, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QSyntaxHighlighter) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QSyntaxHighlighter_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QSyntaxHighlighter) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_EventFilter
func miqt_exec_callback_QSyntaxHighlighter_EventFilter(self QSyntaxHighlighter, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QSyntaxHighlighter) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QSyntaxHighlighter_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSyntaxHighlighter) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_TimerEvent
func miqt_exec_callback_QSyntaxHighlighter_TimerEvent(self QSyntaxHighlighter, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QSyntaxHighlighter) callVirtualBase_ChildEvent(event *QChildEvent) {

	QSyntaxHighlighter_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSyntaxHighlighter) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_ChildEvent
func miqt_exec_callback_QSyntaxHighlighter_ChildEvent(self QSyntaxHighlighter, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QSyntaxHighlighter) callVirtualBase_CustomEvent(event *QEvent) {

	QSyntaxHighlighter_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QSyntaxHighlighter) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_CustomEvent
func miqt_exec_callback_QSyntaxHighlighter_CustomEvent(self QSyntaxHighlighter, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QSyntaxHighlighter) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QSyntaxHighlighter_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSyntaxHighlighter) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_ConnectNotify
func miqt_exec_callback_QSyntaxHighlighter_ConnectNotify(self QSyntaxHighlighter, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QSyntaxHighlighter) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QSyntaxHighlighter_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QSyntaxHighlighter) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSyntaxHighlighter_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSyntaxHighlighter_DisconnectNotify
func miqt_exec_callback_QSyntaxHighlighter_DisconnectNotify(self QSyntaxHighlighter, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QSyntaxHighlighter{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
