package qt

import (
	"unsafe"
)

type QGenericPlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQGenericPlugin constructs a new QGenericPlugin object.
func NewQGenericPlugin() *QGenericPlugin {
	ret := newQGenericPlugin(QGenericPlugin_new())
	ret.isSubclass = true
	return ret
}

// NewQGenericPlugin2 constructs a new QGenericPlugin object.
func NewQGenericPlugin2(parent *QObject) *QGenericPlugin {
	ret := newQGenericPlugin(QGenericPlugin_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QGenericPlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QGenericPlugin_MetaObject(this.h))
}

func (this *QGenericPlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QGenericPlugin_Metacast(this.h, param1_Cstring))
}

func QGenericPlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QGenericPlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGenericPlugin) Create(name string, spec string) *QObject {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	spec_ms := struct_miqt_string{}
	spec_ms.data = CString(spec)
	spec_ms.len = size_t(len(spec))
	defer free(unsafe.Pointer(spec_ms.data))
	return newQObject(QGenericPlugin_Create(this.h, name_ms, spec_ms))
}

func QGenericPlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGenericPlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QGenericPlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QGenericPlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QGenericPlugin) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QGenericPlugin_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QGenericPlugin) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_MetaObject
func miqt_exec_callback_QGenericPlugin_MetaObject(self QGenericPlugin, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QGenericPlugin{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QGenericPlugin) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QGenericPlugin_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QGenericPlugin) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QGenericPlugin_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QGenericPlugin_Metacast
func miqt_exec_callback_QGenericPlugin_Metacast(self QGenericPlugin, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QGenericPlugin{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
