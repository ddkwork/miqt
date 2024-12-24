package qt

import (
	"unsafe"
)

type QTemporaryFile struct {
	h          uintptr
	isSubclass bool
}

// NewQTemporaryFile constructs a new QTemporaryFile object.
func NewQTemporaryFile() *QTemporaryFile {
	ret := newQTemporaryFile(QTemporaryFile_new())
	ret.isSubclass = true
	return ret
}

// NewQTemporaryFile2 constructs a new QTemporaryFile object.
func NewQTemporaryFile2(templateName string) *QTemporaryFile {
	templateName_ms := struct_miqt_string{}
	templateName_ms.data = CString(templateName)
	templateName_ms.len = size_t(len(templateName))
	defer free(unsafe.Pointer(templateName_ms.data))

	ret := newQTemporaryFile(QTemporaryFile_new2(templateName_ms))
	ret.isSubclass = true
	return ret
}

// NewQTemporaryFile3 constructs a new QTemporaryFile object.
func NewQTemporaryFile3(parent *QObject) *QTemporaryFile {
	ret := newQTemporaryFile(QTemporaryFile_new3(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTemporaryFile4 constructs a new QTemporaryFile object.
func NewQTemporaryFile4(templateName string, parent *QObject) *QTemporaryFile {
	templateName_ms := struct_miqt_string{}
	templateName_ms.data = CString(templateName)
	templateName_ms.len = size_t(len(templateName))
	defer free(unsafe.Pointer(templateName_ms.data))

	ret := newQTemporaryFile(QTemporaryFile_new4(templateName_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTemporaryFile) MetaObject() *QMetaObject {
	return newQMetaObject(QTemporaryFile_MetaObject(this.h))
}

func (this *QTemporaryFile) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTemporaryFile_Metacast(this.h, param1_Cstring))
}

func QTemporaryFile_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTemporaryFile_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) AutoRemove() bool {
	return (bool)(QTemporaryFile_AutoRemove(this.h))
}

func (this *QTemporaryFile) SetAutoRemove(b bool) {
	QTemporaryFile_SetAutoRemove(this.h, (bool)(b))
}

func (this *QTemporaryFile) Open() bool {
	return (bool)(QTemporaryFile_Open(this.h))
}

func (this *QTemporaryFile) FileName() string {
	var _ms struct_miqt_string = QTemporaryFile_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) FileTemplate() string {
	var _ms struct_miqt_string = QTemporaryFile_FileTemplate(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) SetFileTemplate(name string) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QTemporaryFile_SetFileTemplate(this.h, name_ms)
}

func (this *QTemporaryFile) Rename(newName string) bool {
	newName_ms := struct_miqt_string{}
	newName_ms.data = CString(newName)
	newName_ms.len = size_t(len(newName))
	defer free(unsafe.Pointer(newName_ms.data))
	return (bool)(QTemporaryFile_Rename(this.h, newName_ms))
}

func QTemporaryFile_CreateNativeFile(fileName string) *QTemporaryFile {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	return newQTemporaryFile(QTemporaryFile_CreateNativeFile(fileName_ms))
}

func QTemporaryFile_CreateNativeFileWithFile(file *QFile) *QTemporaryFile {
	return newQTemporaryFile(QTemporaryFile_CreateNativeFileWithFile(file.cPointer()))
}

func QTemporaryFile_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTemporaryFile_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTemporaryFile_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTemporaryFile_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryFile) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTemporaryFile_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTemporaryFile) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_MetaObject
func miqt_exec_callback_QTemporaryFile_MetaObject(self QTemporaryFile, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTemporaryFile) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTemporaryFile_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTemporaryFile) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTemporaryFile_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTemporaryFile_Metacast
func miqt_exec_callback_QTemporaryFile_Metacast(self QTemporaryFile, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTemporaryFile{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
