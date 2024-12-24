package qt

import (
	"unsafe"
)

type QIconEnginePlugin struct {
	h          uintptr
	isSubclass bool
}

// NewQIconEnginePlugin constructs a new QIconEnginePlugin object.
func NewQIconEnginePlugin() *QIconEnginePlugin {
	g := newQIconEnginePlugin(QIconEnginePlugin_new())
	g.isSubclass = true
	return g
}

// NewQIconEnginePlugin2 constructs a new QIconEnginePlugin object.
func NewQIconEnginePlugin2(parent *QObject) *QIconEnginePlugin {
	g := newQIconEnginePlugin(QIconEnginePlugin_new2(parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QIconEnginePlugin) MetaObject() *QMetaObject {
	return newQMetaObject(QIconEnginePlugin_MetaObject(this.h))
}

func (this *QIconEnginePlugin) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QIconEnginePlugin_Metacast(this.h, param1_Cstring))
}

func QIconEnginePlugin_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QIconEnginePlugin_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEnginePlugin) Create(filename string) *QIconEngine {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return newQIconEngine(QIconEnginePlugin_Create(this.h, filename_ms))
}

func QIconEnginePlugin_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIconEnginePlugin_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QIconEnginePlugin_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QIconEnginePlugin_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QIconEnginePlugin) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QIconEnginePlugin_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QIconEnginePlugin) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_MetaObject
func miqt_exec_callback_QIconEnginePlugin_MetaObject(self QIconEnginePlugin, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QIconEnginePlugin) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QIconEnginePlugin_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QIconEnginePlugin) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QIconEnginePlugin_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QIconEnginePlugin_Metacast
func miqt_exec_callback_QIconEnginePlugin_Metacast(self QIconEnginePlugin, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QIconEnginePlugin{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
