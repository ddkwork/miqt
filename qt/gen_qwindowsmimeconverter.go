package qt

import (
	"unsafe"
)

type QWindowsMimeConverter struct {
	h          uintptr
	isSubclass bool
}

// NewQWindowsMimeConverter constructs a new QWindowsMimeConverter object.
func NewQWindowsMimeConverter() *QWindowsMimeConverter {
	ret := newQWindowsMimeConverter(QWindowsMimeConverter_new())
	ret.isSubclass = true
	return ret
}

func QWindowsMimeConverter_RegisterMimeType(mimeType string) int {
	mimeType_ms := struct_miqt_string{}
	mimeType_ms.data = CString(mimeType)
	mimeType_ms.len = size_t(len(mimeType))
	defer free(unsafe.Pointer(mimeType_ms.data))
	return (int)(QWindowsMimeConverter_RegisterMimeType(mimeType_ms))
}

func (this *QWindowsMimeConverter) CanConvertFromMime(formatetc *tagFORMATETC, mimeData *QMimeData) bool {
	return (bool)(QWindowsMimeConverter_CanConvertFromMime(this.h, formatetc, mimeData.cPointer()))
}

func (this *QWindowsMimeConverter) ConvertFromMime(formatetc *tagFORMATETC, mimeData *QMimeData, pmedium *tagSTGMEDIUM) bool {
	return (bool)(QWindowsMimeConverter_ConvertFromMime(this.h, formatetc, mimeData.cPointer(), pmedium))
}

func (this *QWindowsMimeConverter) FormatsForMime(mimeType string, mimeData *QMimeData) []tagFORMATETC {
	mimeType_ms := struct_miqt_string{}
	mimeType_ms.data = CString(mimeType)
	mimeType_ms.len = size_t(len(mimeType))
	defer free(unsafe.Pointer(mimeType_ms.data))
	var _ma struct_miqt_array = QWindowsMimeConverter_FormatsForMime(this.h, mimeType_ms, mimeData.cPointer())
	_ret := make([]tagFORMATETC, int(_ma.len))
	_outCast := (*[0xffff]tagFORMATETC)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		xxxxxxxxx
	}
	return _ret
}

func (this *QWindowsMimeConverter) CanConvertToMime(mimeType string, pDataObj *IDataObject) bool {
	mimeType_ms := struct_miqt_string{}
	mimeType_ms.data = CString(mimeType)
	mimeType_ms.len = size_t(len(mimeType))
	defer free(unsafe.Pointer(mimeType_ms.data))
	return (bool)(QWindowsMimeConverter_CanConvertToMime(this.h, mimeType_ms, pDataObj))
}

func (this *QWindowsMimeConverter) ConvertToMime(mimeType string, pDataObj *IDataObject, preferredType QMetaType) *QVariant {
	mimeType_ms := struct_miqt_string{}
	mimeType_ms.data = CString(mimeType)
	mimeType_ms.len = size_t(len(mimeType))
	defer free(unsafe.Pointer(mimeType_ms.data))
	_goptr := newQVariant(QWindowsMimeConverter_ConvertToMime(this.h, mimeType_ms, pDataObj, preferredType.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWindowsMimeConverter) MimeForFormat(formatetc *tagFORMATETC) string {
	var _ms struct_miqt_string = QWindowsMimeConverter_MimeForFormat(this.h, formatetc)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
