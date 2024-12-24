package qt

import (
	"unsafe"
)

type QSignalMapper struct {
	h          uintptr
	isSubclass bool
}

// NewQSignalMapper constructs a new QSignalMapper object.
func NewQSignalMapper() *QSignalMapper {
	g := newQSignalMapper(QSignalMapper_new())
	g.isSubclass = true
	return g
}

// NewQSignalMapper2 constructs a new QSignalMapper object.
func NewQSignalMapper2(parent *QObject) *QSignalMapper {
	g := newQSignalMapper(QSignalMapper_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QSignalMapper) MetaObject() *QMetaObject {
	return newQMetaObject(QSignalMapper_MetaObject(this.h))
}

func (this *QSignalMapper) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QSignalMapper_Metacast(this.h, param1_Cstring))
}

func QSignalMapper_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QSignalMapper_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSignalMapper) SetMapping(sender *QObject, id int) {
	QSignalMapper_SetMapping(this.h, sender.cPointer(), (int)(id))
}

func (this *QSignalMapper) SetMapping2(sender *QObject, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QSignalMapper_SetMapping2(this.h, sender.cPointer(), text_ms)
}

func (this *QSignalMapper) SetMapping3(sender *QObject, object *QObject) {
	QSignalMapper_SetMapping3(this.h, sender.cPointer(), object.cPointer())
}

func (this *QSignalMapper) RemoveMappings(sender *QObject) {
	QSignalMapper_RemoveMappings(this.h, sender.cPointer())
}

func (this *QSignalMapper) Mapping(id int) *QObject {
	return newQObject(QSignalMapper_Mapping(this.h, (int)(id)))
}

func (this *QSignalMapper) MappingWithText(text string) *QObject {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	return newQObject(QSignalMapper_MappingWithText(this.h, text_ms))
}

func (this *QSignalMapper) MappingWithObject(object *QObject) *QObject {
	return newQObject(QSignalMapper_MappingWithObject(this.h, object.cPointer()))
}

func (this *QSignalMapper) MappedInt(param1 int) {
	QSignalMapper_MappedInt(this.h, (int)(param1))
}

func (this *QSignalMapper) OnMappedInt(slot func(param1 int)) {
	QSignalMapper_connect_MappedInt(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MappedInt
func miqt_exec_callback_QSignalMapper_MappedInt(cb intptr_t, param1 int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	gofunc(slotval1)
}

func (this *QSignalMapper) MappedString(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QSignalMapper_MappedString(this.h, param1_ms)
}

func (this *QSignalMapper) OnMappedString(slot func(param1 string)) {
	QSignalMapper_connect_MappedString(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MappedString
func miqt_exec_callback_QSignalMapper_MappedString(cb intptr_t, param1 struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var param1_ms struct_miqt_string = param1
	param1_ret := GoStringN(param1_ms.data, int(int64(param1_ms.len)))
	free(unsafe.Pointer(param1_ms.data))
	slotval1 := param1_ret

	gofunc(slotval1)
}

func (this *QSignalMapper) MappedObject(param1 *QObject) {
	QSignalMapper_MappedObject(this.h, param1.cPointer())
}

func (this *QSignalMapper) OnMappedObject(slot func(param1 *QObject)) {
	QSignalMapper_connect_MappedObject(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MappedObject
func miqt_exec_callback_QSignalMapper_MappedObject(cb intptr_t, param1 *QObject) {
	gofunc, ok := cgo.Handle(cb).Value().(func(param1 *QObject))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	gofunc(slotval1)
}

func (this *QSignalMapper) Map() {
	QSignalMapper_Map(this.h)
}

func (this *QSignalMapper) MapWithSender(sender *QObject) {
	QSignalMapper_MapWithSender(this.h, sender.cPointer())
}

func QSignalMapper_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSignalMapper_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QSignalMapper_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QSignalMapper_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QSignalMapper) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QSignalMapper_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QSignalMapper) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_MetaObject
func miqt_exec_callback_QSignalMapper_MetaObject(self QSignalMapper, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QSignalMapper{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QSignalMapper) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QSignalMapper_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QSignalMapper) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QSignalMapper_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QSignalMapper_Metacast
func miqt_exec_callback_QSignalMapper_Metacast(self QSignalMapper, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QSignalMapper{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
