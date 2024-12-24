package qt

import (
	"unsafe"
)

type QStylePlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQStylePlugin constructs a new QStylePlugin object.
func NewQStylePlugin() *QStylePlugin {
	g := newQStylePlugin(QStylePlugin_new())
	g.isSubclass = true
	return g
}

// NewQStylePlugin2 constructs a new QStylePlugin object.
func NewQStylePlugin2(parent *QObject) *QStylePlugin {
	g := newQStylePlugin(QStylePlugin_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QStylePlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QStylePlugin_MetaObject(this.h))
}

func (this *QStylePlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QStylePlugin_Metacast(this.h, param1_Cstring))
}

func QStylePlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QStylePlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStylePlugin) Create(key string) *QStyle {
	key_ms := struct_miqt_string{}
	key_ms.data = CString(key)
	key_ms.len = size_t(len(key))
	defer free(unsafe.Pointer(key_ms.data))
	return newQStyle(QStylePlugin_Create(this.h, key_ms))
}

func QStylePlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStylePlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QStylePlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QStylePlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStylePlugin) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QStylePlugin_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QStylePlugin) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_MetaObject
func miqt_exec_callback_QStylePlugin_MetaObject(self QStylePlugin, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QStylePlugin{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QStylePlugin) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QStylePlugin_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QStylePlugin) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QStylePlugin_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QStylePlugin_Metacast
func miqt_exec_callback_QStylePlugin_Metacast(self QStylePlugin, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QStylePlugin{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
