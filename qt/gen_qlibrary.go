package qt

import (
	"unsafe"
)

type QLibrary__LoadHint int

const (
	QLibrary__ResolveAllSymbolsHint     QLibrary__LoadHint = 1
	QLibrary__ExportExternalSymbolsHint QLibrary__LoadHint = 2
	QLibrary__LoadArchiveMemberHint     QLibrary__LoadHint = 4
	QLibrary__PreventUnloadHint         QLibrary__LoadHint = 8
	QLibrary__DeepBindHint              QLibrary__LoadHint = 16
)

type QLibrary struct {
	h          uintptr
	isSubclass bool
}

// NewQLibrary constructs a new QLibrary object.
func NewQLibrary() *QLibrary {
	g := newQLibrary(QLibrary_new())
	g.isSubclass = true
	return g
}

// NewQLibrary2 constructs a new QLibrary object.
func NewQLibrary2(fileName string) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQLibrary(QLibrary_new2(fileName_ms))
	g.isSubclass = true
	return g
}

// NewQLibrary3 constructs a new QLibrary object.
func NewQLibrary3(fileName string, verNum int) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQLibrary(QLibrary_new3(fileName_ms, (int)(verNum)))
	g.isSubclass = true
	return g
}

// NewQLibrary4 constructs a new QLibrary object.
func NewQLibrary4(fileName string, version string) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))
	g := newQLibrary(QLibrary_new4(fileName_ms, version_ms))
	g.isSubclass = true
	return g
}

// NewQLibrary5 constructs a new QLibrary object.
func NewQLibrary5(parent *QObject) *QLibrary {
	g := newQLibrary(QLibrary_new5(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQLibrary6 constructs a new QLibrary object.
func NewQLibrary6(fileName string, parent *QObject) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQLibrary(QLibrary_new6(fileName_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQLibrary7 constructs a new QLibrary object.
func NewQLibrary7(fileName string, verNum int, parent *QObject) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	g := newQLibrary(QLibrary_new7(fileName_ms, (int)(verNum), parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQLibrary8 constructs a new QLibrary object.
func NewQLibrary8(fileName string, version string, parent *QObject) *QLibrary {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))
	g := newQLibrary(QLibrary_new8(fileName_ms, version_ms, parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QLibrary) MetaObject() *QMetaObject {
	return newQMetaObject(QLibrary_MetaObject(this.h))
}

func (this *QLibrary) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLibrary_Metacast(this.h, param1_Cstring))
}

func QLibrary_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLibrary_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) Load() bool {
	return (bool)(QLibrary_Load(this.h))
}

func (this *QLibrary) Unload() bool {
	return (bool)(QLibrary_Unload(this.h))
}

func (this *QLibrary) IsLoaded() bool {
	return (bool)(QLibrary_IsLoaded(this.h))
}

func QLibrary_IsLibrary(fileName string) bool {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return (bool)(QLibrary_IsLibrary(fileName_ms))
}

func (this *QLibrary) SetFileName(fileName string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QLibrary_SetFileName(this.h, fileName_ms)
}

func (this *QLibrary) FileName() string {
	var _ms struct_miqt_string = QLibrary_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) SetFileNameAndVersion(fileName string, verNum int) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	QLibrary_SetFileNameAndVersion(this.h, fileName_ms, (int)(verNum))
}

func (this *QLibrary) SetFileNameAndVersion2(fileName string, version string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	version_ms := struct_miqt_string{}
	version_ms.data = CString(version)
	version_ms.len = size_t(len(version))
	defer free(unsafe.Pointer(version_ms.data))
	QLibrary_SetFileNameAndVersion2(this.h, fileName_ms, version_ms)
}

func (this *QLibrary) ErrorString() string {
	var _ms struct_miqt_string = QLibrary_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) SetLoadHints(hints LoadHints) {
	QLibrary_SetLoadHints(this.h, hints)
}

func (this *QLibrary) LoadHints() LoadHints {
	xxxxxxxxx
}

func QLibrary_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLibrary_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLibrary_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLibrary_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLibrary) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QLibrary_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QLibrary) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_MetaObject
func miqt_exec_callback_QLibrary_MetaObject(self QLibrary, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLibrary{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QLibrary) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLibrary_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLibrary) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLibrary_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLibrary_Metacast
func miqt_exec_callback_QLibrary_Metacast(self QLibrary, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QLibrary{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
