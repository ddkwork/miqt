package qt

import (
	"unsafe"
)

type QPluginLoader struct {
	h          uintptr
	isSubclass bool
}

// NewQPluginLoader constructs a new QPluginLoader object.
func NewQPluginLoader() *QPluginLoader {
	ret := newQPluginLoader(QPluginLoader_new())
	ret.isSubclass = true
	return ret
}

// NewQPluginLoader2 constructs a new QPluginLoader object.
func NewQPluginLoader2(fileName string) *QPluginLoader {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQPluginLoader(QPluginLoader_new2(fileName_ms))
	ret.isSubclass = true
	return ret
}

// NewQPluginLoader3 constructs a new QPluginLoader object.
func NewQPluginLoader3(parent *QObject) *QPluginLoader {
	ret := newQPluginLoader(QPluginLoader_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQPluginLoader4 constructs a new QPluginLoader object.
func NewQPluginLoader4(fileName string, parent *QObject) *QPluginLoader {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))

	ret := newQPluginLoader(QPluginLoader_new4(fileName_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPluginLoader) MetaObject() *QMetaObject {
	return newQMetaObject(QPluginLoader_MetaObject(this.h))
}

func (this *QPluginLoader) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPluginLoader_Metacast(this.h, param1_Cstring))
}

func QPluginLoader_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPluginLoader_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) Instance() *QObject {
	return newQObject(QPluginLoader_Instance(this.h))
}

func (this *QPluginLoader) MetaData() *QJsonObject {
	_goptr := newQJsonObject(QPluginLoader_MetaData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QPluginLoader_StaticInstances() []*QObject {
	var _ma struct_miqt_array = QPluginLoader_StaticInstances()
	_ret := make([]*QObject, int(_ma.len))
	_outCast := (*[0xffff]*QObject)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQObject(_outCast[i])
	}
	return _ret
}

func QPluginLoader_StaticPlugins() []QStaticPlugin {
	var _ma struct_miqt_array = QPluginLoader_StaticPlugins()
	_ret := make([]QStaticPlugin, int(_ma.len))
	_outCast := (*[0xffff]*QStaticPlugin)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQStaticPlugin(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QPluginLoader) Load() bool {
	return (bool)(QPluginLoader_Load(this.h))
}

func (this *QPluginLoader) Unload() bool {
	return (bool)(QPluginLoader_Unload(this.h))
}

func (this *QPluginLoader) IsLoaded() bool {
	return (bool)(QPluginLoader_IsLoaded(this.h))
}

func (this *QPluginLoader) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QPluginLoader_SetFileName(this.h, fileName_ms)
}

func (this *QPluginLoader) FileName() string {
	var _ms struct_miqt_string = QPluginLoader_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) ErrorString() string {
	var _ms struct_miqt_string = QPluginLoader_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) SetLoadHints(loadHints LoadHint) {
	QPluginLoader_SetLoadHints(this.h, (int)(loadHints))
}

func (this *QPluginLoader) LoadHints() LoadHint {
	return (LoadHint)(QPluginLoader_LoadHints(this.h))
}

func QPluginLoader_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPluginLoader_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPluginLoader_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPluginLoader_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPluginLoader) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPluginLoader_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPluginLoader) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_MetaObject
func miqt_exec_callback_QPluginLoader_MetaObject(self QPluginLoader, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPluginLoader{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPluginLoader) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPluginLoader_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPluginLoader) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPluginLoader_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPluginLoader_Metacast
func miqt_exec_callback_QPluginLoader_Metacast(self QPluginLoader, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QPluginLoader{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
