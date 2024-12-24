package qt

import (
	"unsafe"
)

type QMimeDatabase__MatchMode int

const (
	QMimeDatabase__MatchDefault   QMimeDatabase__MatchMode = 0
	QMimeDatabase__MatchExtension QMimeDatabase__MatchMode = 1
	QMimeDatabase__MatchContent   QMimeDatabase__MatchMode = 2
)

type QMimeDatabase struct {
	h          uintptr
	isSubclass bool
}

// NewQMimeDatabase constructs a new QMimeDatabase object.
func NewQMimeDatabase() *QMimeDatabase {
	ret := newQMimeDatabase(QMimeDatabase_new())
	ret.isSubclass = true
	return ret
}

func (this *QMimeDatabase) MimeTypeForName(nameOrAlias string) *QMimeType {
	nameOrAlias_ms := struct_miqt_string{}
	nameOrAlias_ms.data = CString(nameOrAlias)
	nameOrAlias_ms.len = size_t(len(nameOrAlias))
	defer free(unsafe.Pointer(nameOrAlias_ms.data))
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForName(this.h, nameOrAlias_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForFile(fileName string) *QMimeType {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForFile(this.h, fileName_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForFileWithFileInfo(fileInfo *QFileInfo) *QMimeType {
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForFileWithFileInfo(this.h, fileInfo.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypesForFileName(fileName string) []QMimeType {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ma struct_miqt_array = QMimeDatabase_MimeTypesForFileName(this.h, fileName_ms)
	_ret := make([]QMimeType, int(_ma.len))
	_outCast := (*[0xffff]*QMimeType)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQMimeType(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMimeDatabase) MimeTypeForData(data []byte) *QMimeType {
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForData(this.h, data_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForDataWithDevice(device *QIODevice) *QMimeType {
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForDataWithDevice(this.h, device.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForUrl(url *QUrl) *QMimeType {
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForUrl(this.h, url.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForFileNameAndData(fileName string, device *QIODevice) *QMimeType {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForFileNameAndData(this.h, fileName_ms, device.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForFileNameAndData2(fileName string, data []byte) *QMimeType {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForFileNameAndData2(this.h, fileName_ms, data_alias))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) SuffixForFileName(fileName string) string {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	var _ms struct_miqt_string = QMimeDatabase_SuffixForFileName(this.h, fileName_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeDatabase) AllMimeTypes() []QMimeType {
	var _ma struct_miqt_array = QMimeDatabase_AllMimeTypes(this.h)
	_ret := make([]QMimeType, int(_ma.len))
	_outCast := (*[0xffff]*QMimeType)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQMimeType(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMimeDatabase) MimeTypeForFile2(fileName string, mode MatchMode) *QMimeType {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForFile2(this.h, fileName_ms, mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeDatabase) MimeTypeForFile22(fileInfo *QFileInfo, mode MatchMode) *QMimeType {
	_goptr := newQMimeType(QMimeDatabase_MimeTypeForFile22(this.h, fileInfo.cPointer(), mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
