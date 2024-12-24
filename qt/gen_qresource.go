package qt

import (
	"unsafe"
)

type QResource__Compression int

const (
	QResource__NoCompression   QResource__Compression = 0
	QResource__ZlibCompression QResource__Compression = 1
	QResource__ZstdCompression QResource__Compression = 2
)

type QResource struct {
	h          uintptr
	isSubclass bool
}

// NewQResource constructs a new QResource object.
func NewQResource() *QResource {
	g := newQResource(QResource_new())
	g.isSubclass = true
	return g
}

// NewQResource2 constructs a new QResource object.
func NewQResource2(file string) *QResource {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	g := newQResource(QResource_new2(file_ms))
	g.isSubclass = true
	return g
}

// NewQResource3 constructs a new QResource object.
func NewQResource3(file string, locale *QLocale) *QResource {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	g := newQResource(QResource_new3(file_ms, locale.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QResource) SetFileName(file string) {
	file_ms := struct_miqt_string{}
	file_ms.data = CString(file)
	file_ms.len = size_t(len(file))
	defer free(unsafe.Pointer(file_ms.data))
	QResource_SetFileName(this.h, file_ms)
}

func (this *QResource) FileName() string {
	var _ms struct_miqt_string = QResource_FileName(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QResource) AbsoluteFilePath() string {
	var _ms struct_miqt_string = QResource_AbsoluteFilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QResource) SetLocale(locale *QLocale) {
	QResource_SetLocale(this.h, locale.cPointer())
}

func (this *QResource) Locale() *QLocale {
	_goptr := newQLocale(QResource_Locale(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QResource) IsValid() bool {
	return (bool)(QResource_IsValid(this.h))
}

func (this *QResource) CompressionAlgorithm() Compression {
	xxxxxxxxx
}

func (this *QResource) Size() int64 {
	return (int64)(QResource_Size(this.h))
}

func (this *QResource) Data() *byte {
	return (*byte)(unsafe.Pointer(QResource_Data(this.h)))
}

func (this *QResource) UncompressedSize() int64 {
	return (int64)(QResource_UncompressedSize(this.h))
}

func (this *QResource) UncompressedData() []byte {
	var _bytearray struct_miqt_string = QResource_UncompressedData(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QResource) LastModified() *QDateTime {
	_goptr := newQDateTime(QResource_LastModified(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QResource_RegisterResource(rccFilename string) bool {
	rccFilename_ms := struct_miqt_string{}
	rccFilename_ms.data = CString(rccFilename)
	rccFilename_ms.len = size_t(len(rccFilename))
	defer free(unsafe.Pointer(rccFilename_ms.data))
	return (bool)(QResource_RegisterResource(rccFilename_ms))
}

func QResource_UnregisterResource(rccFilename string) bool {
	rccFilename_ms := struct_miqt_string{}
	rccFilename_ms.data = CString(rccFilename)
	rccFilename_ms.len = size_t(len(rccFilename))
	defer free(unsafe.Pointer(rccFilename_ms.data))
	return (bool)(QResource_UnregisterResource(rccFilename_ms))
}

func QResource_RegisterResourceWithRccData(rccData *byte) bool {
	return (bool)(QResource_RegisterResourceWithRccData((*uchar)(unsafe.Pointer(rccData))))
}

func QResource_UnregisterResourceWithRccData(rccData *byte) bool {
	return (bool)(QResource_UnregisterResourceWithRccData((*uchar)(unsafe.Pointer(rccData))))
}

func QResource_RegisterResource2(rccFilename string, resourceRoot string) bool {
	rccFilename_ms := struct_miqt_string{}
	rccFilename_ms.data = CString(rccFilename)
	rccFilename_ms.len = size_t(len(rccFilename))
	defer free(unsafe.Pointer(rccFilename_ms.data))
	resourceRoot_ms := struct_miqt_string{}
	resourceRoot_ms.data = CString(resourceRoot)
	resourceRoot_ms.len = size_t(len(resourceRoot))
	defer free(unsafe.Pointer(resourceRoot_ms.data))
	return (bool)(QResource_RegisterResource2(rccFilename_ms, resourceRoot_ms))
}

func QResource_UnregisterResource2(rccFilename string, resourceRoot string) bool {
	rccFilename_ms := struct_miqt_string{}
	rccFilename_ms.data = CString(rccFilename)
	rccFilename_ms.len = size_t(len(rccFilename))
	defer free(unsafe.Pointer(rccFilename_ms.data))
	resourceRoot_ms := struct_miqt_string{}
	resourceRoot_ms.data = CString(resourceRoot)
	resourceRoot_ms.len = size_t(len(resourceRoot))
	defer free(unsafe.Pointer(resourceRoot_ms.data))
	return (bool)(QResource_UnregisterResource2(rccFilename_ms, resourceRoot_ms))
}

func QResource_RegisterResource22(rccData *byte, resourceRoot string) bool {
	resourceRoot_ms := struct_miqt_string{}
	resourceRoot_ms.data = CString(resourceRoot)
	resourceRoot_ms.len = size_t(len(resourceRoot))
	defer free(unsafe.Pointer(resourceRoot_ms.data))
	return (bool)(QResource_RegisterResource22((*uchar)(unsafe.Pointer(rccData)), resourceRoot_ms))
}

func QResource_UnregisterResource22(rccData *byte, resourceRoot string) bool {
	resourceRoot_ms := struct_miqt_string{}
	resourceRoot_ms.data = CString(resourceRoot)
	resourceRoot_ms.len = size_t(len(resourceRoot))
	defer free(unsafe.Pointer(resourceRoot_ms.data))
	return (bool)(QResource_UnregisterResource22((*uchar)(unsafe.Pointer(rccData)), resourceRoot_ms))
}
