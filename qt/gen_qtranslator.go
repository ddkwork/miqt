package qt

import (
	"unsafe"
)

type QTranslator struct {
	h          uintptr
	isSubclass bool
}

// NewQTranslator constructs a new QTranslator object.
func NewQTranslator() *QTranslator {

	ret := newQTranslator(QTranslator_new())
	ret.isSubclass = true
	return ret
}

// NewQTranslator2 constructs a new QTranslator object.
func NewQTranslator2(parent *QObject) *QTranslator {

	ret := newQTranslator(QTranslator_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTranslator) MetaObject() *QMetaObject {
	return newQMetaObject(QTranslator_MetaObject(this.h))
}

func (this *QTranslator) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTranslator_Metacast(this.h, param1_Cstring))
}

func QTranslator_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTranslator_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	context_Cstring := CString(context)
	defer free(unsafe.Pointer(context_Cstring))
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QTranslator_Translate(this.h, context_Cstring, sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) IsEmpty() bool {
	return (bool)(QTranslator_IsEmpty(this.h))
}

func (this *QTranslator) Language() string {
	var _ms struct_miqt_string = QTranslator_Language(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) FilePath() string {
	var _ms struct_miqt_string = QTranslator_FilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) Load(filename string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QTranslator_Load(this.h, filename_ms))
}

func (this *QTranslator) Load2(locale *QLocale, filename string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QTranslator_Load2(this.h, locale.cPointer(), filename_ms))
}

func (this *QTranslator) Load3(data *byte, lenVal int) bool {
	return (bool)(QTranslator_Load3(this.h, (*uchar)(unsafe.Pointer(data)), (int)(lenVal)))
}

func QTranslator_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTranslator_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTranslator_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTranslator_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) Load22(filename string, directory string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	return (bool)(QTranslator_Load22(this.h, filename_ms, directory_ms))
}

func (this *QTranslator) Load32(filename string, directory string, search_delimiters string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	search_delimiters_ms := struct_miqt_string{}
	search_delimiters_ms.data = CString(search_delimiters)
	search_delimiters_ms.len = size_t(len(search_delimiters))
	defer free(unsafe.Pointer(search_delimiters_ms.data))
	return (bool)(QTranslator_Load32(this.h, filename_ms, directory_ms, search_delimiters_ms))
}

func (this *QTranslator) Load4(filename string, directory string, search_delimiters string, suffix string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	search_delimiters_ms := struct_miqt_string{}
	search_delimiters_ms.data = CString(search_delimiters)
	search_delimiters_ms.len = size_t(len(search_delimiters))
	defer free(unsafe.Pointer(search_delimiters_ms.data))
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	return (bool)(QTranslator_Load4(this.h, filename_ms, directory_ms, search_delimiters_ms, suffix_ms))
}

func (this *QTranslator) Load33(locale *QLocale, filename string, prefix string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	return (bool)(QTranslator_Load33(this.h, locale.cPointer(), filename_ms, prefix_ms))
}

func (this *QTranslator) Load42(locale *QLocale, filename string, prefix string, directory string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	return (bool)(QTranslator_Load42(this.h, locale.cPointer(), filename_ms, prefix_ms, directory_ms))
}

func (this *QTranslator) Load5(locale *QLocale, filename string, prefix string, directory string, suffix string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	return (bool)(QTranslator_Load5(this.h, locale.cPointer(), filename_ms, prefix_ms, directory_ms, suffix_ms))
}

func (this *QTranslator) Load34(data *byte, lenVal int, directory string) bool {
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	return (bool)(QTranslator_Load34(this.h, (*uchar)(unsafe.Pointer(data)), (int)(lenVal), directory_ms))
}

func (this *QTranslator) callVirtualBase_Translate(context string, sourceText string, disambiguation string, n int) string {
	context_Cstring := CString(context)
	defer free(unsafe.Pointer(context_Cstring))
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))

	var _ms struct_miqt_string = QTranslator_virtualbase_Translate(unsafe.Pointer(this.h), context_Cstring, sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QTranslator) OnTranslate(slot func(super func(context string, sourceText string, disambiguation string, n int) string, context string, sourceText string, disambiguation string, n int) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_Translate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_Translate
func miqt_exec_callback_QTranslator_Translate(self QTranslator, cb intptr_t, context *const_char, sourceText *const_char, disambiguation *const_char, n int) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(context string, sourceText string, disambiguation string, n int) string, context string, sourceText string, disambiguation string, n int) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	context_ret := context
	slotval1 := GoString(context_ret)

	sourceText_ret := sourceText
	slotval2 := GoString(sourceText_ret)

	disambiguation_ret := disambiguation
	slotval3 := GoString(disambiguation_ret)

	slotval4 := (int)(n)

	virtualReturn := gofunc((&QTranslator{h: self}).callVirtualBase_Translate, slotval1, slotval2, slotval3, slotval4)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QTranslator) callVirtualBase_IsEmpty() bool {

	return (bool)(QTranslator_virtualbase_IsEmpty(unsafe.Pointer(this.h)))

}
func (this *QTranslator) OnIsEmpty(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_IsEmpty(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_IsEmpty
func miqt_exec_callback_QTranslator_IsEmpty(self QTranslator, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTranslator{h: self}).callVirtualBase_IsEmpty)

	return (bool)(virtualReturn)

}

func (this *QTranslator) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTranslator_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTranslator) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_Event
func miqt_exec_callback_QTranslator_Event(self QTranslator, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTranslator{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTranslator) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QTranslator_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QTranslator) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_EventFilter
func miqt_exec_callback_QTranslator_EventFilter(self QTranslator, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QTranslator{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QTranslator) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QTranslator_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTranslator) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_TimerEvent
func miqt_exec_callback_QTranslator_TimerEvent(self QTranslator, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QTranslator{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QTranslator) callVirtualBase_ChildEvent(event *QChildEvent) {

	QTranslator_virtualbase_ChildEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTranslator) OnChildEvent(slot func(super func(event *QChildEvent), event *QChildEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_ChildEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_ChildEvent
func miqt_exec_callback_QTranslator_ChildEvent(self QTranslator, cb intptr_t, event *QChildEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QChildEvent), event *QChildEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQChildEvent(event)

	gofunc((&QTranslator{h: self}).callVirtualBase_ChildEvent, slotval1)

}

func (this *QTranslator) callVirtualBase_CustomEvent(event *QEvent) {

	QTranslator_virtualbase_CustomEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTranslator) OnCustomEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_CustomEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_CustomEvent
func miqt_exec_callback_QTranslator_CustomEvent(self QTranslator, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QTranslator{h: self}).callVirtualBase_CustomEvent, slotval1)

}

func (this *QTranslator) callVirtualBase_ConnectNotify(signal *QMetaMethod) {

	QTranslator_virtualbase_ConnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTranslator) OnConnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_ConnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_ConnectNotify
func miqt_exec_callback_QTranslator_ConnectNotify(self QTranslator, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTranslator{h: self}).callVirtualBase_ConnectNotify, slotval1)

}

func (this *QTranslator) callVirtualBase_DisconnectNotify(signal *QMetaMethod) {

	QTranslator_virtualbase_DisconnectNotify(unsafe.Pointer(this.h), signal.cPointer())

}
func (this *QTranslator) OnDisconnectNotify(slot func(super func(signal *QMetaMethod), signal *QMetaMethod)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_DisconnectNotify(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_DisconnectNotify
func miqt_exec_callback_QTranslator_DisconnectNotify(self QTranslator, cb intptr_t, signal *QMetaMethod) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(signal *QMetaMethod), signal *QMetaMethod))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMetaMethod(signal)

	gofunc((&QTranslator{h: self}).callVirtualBase_DisconnectNotify, slotval1)

}
