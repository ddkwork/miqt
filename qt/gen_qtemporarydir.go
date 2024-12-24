package qt

import (
	"unsafe"
)

type QTemporaryDir struct {
	h          uintptr
	isSubclass bool
}

// NewQTemporaryDir constructs a new QTemporaryDir object.
func NewQTemporaryDir() *QTemporaryDir {
	g := newQTemporaryDir(QTemporaryDir_new())
	g.isSubclass = true
	return g
}

// NewQTemporaryDir2 constructs a new QTemporaryDir object.
func NewQTemporaryDir2(templateName string) *QTemporaryDir {
	templateName_ms := struct_miqt_string{}
	templateName_ms.data = CString(templateName)
	templateName_ms.len = size_t(len(templateName))
	defer free(unsafe.Pointer(templateName_ms.data))
	g := newQTemporaryDir(QTemporaryDir_new2(templateName_ms))
	g.isSubclass = true
	return g
}

func (this *QTemporaryDir) Swap(other *QTemporaryDir) {
	QTemporaryDir_Swap(this.h, other.cPointer())
}

func (this *QTemporaryDir) IsValid() bool {
	return (bool)(QTemporaryDir_IsValid(this.h))
}

func (this *QTemporaryDir) ErrorString() string {
	var _ms struct_miqt_string = QTemporaryDir_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryDir) AutoRemove() bool {
	return (bool)(QTemporaryDir_AutoRemove(this.h))
}

func (this *QTemporaryDir) SetAutoRemove(b bool) {
	QTemporaryDir_SetAutoRemove(this.h, (bool)(b))
}

func (this *QTemporaryDir) Remove() bool {
	return (bool)(QTemporaryDir_Remove(this.h))
}

func (this *QTemporaryDir) Path() string {
	var _ms struct_miqt_string = QTemporaryDir_Path(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTemporaryDir) FilePath(fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QTemporaryDir_FilePath(this.h, fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
