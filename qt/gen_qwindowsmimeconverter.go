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
func (this *QWindowsMimeConverter) OnCanConvertFromMime(slot func(formatetc *tagFORMATETC, mimeData *QMimeData) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowsMimeConverter_override_virtual_CanConvertFromMime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowsMimeConverter_CanConvertFromMime
func miqt_exec_callback_QWindowsMimeConverter_CanConvertFromMime(self QWindowsMimeConverter, cb intptr_t, formatetc *tagFORMATETC, mimeData *QMimeData) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(formatetc *tagFORMATETC, mimeData *QMimeData) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*tagFORMATETC)(unsafe.Pointer(formatetc))

	slotval2 := newQMimeData(mimeData)

	virtualReturn := gofunc(slotval1, slotval2)

	return (bool)(virtualReturn)

}
func (this *QWindowsMimeConverter) OnConvertFromMime(slot func(formatetc *tagFORMATETC, mimeData *QMimeData, pmedium *tagSTGMEDIUM) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowsMimeConverter_override_virtual_ConvertFromMime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowsMimeConverter_ConvertFromMime
func miqt_exec_callback_QWindowsMimeConverter_ConvertFromMime(self QWindowsMimeConverter, cb intptr_t, formatetc *tagFORMATETC, mimeData *QMimeData, pmedium *tagSTGMEDIUM) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(formatetc *tagFORMATETC, mimeData *QMimeData, pmedium *tagSTGMEDIUM) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*tagFORMATETC)(unsafe.Pointer(formatetc))

	slotval2 := newQMimeData(mimeData)

	slotval3 := (*tagSTGMEDIUM)(unsafe.Pointer(pmedium))

	virtualReturn := gofunc(slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}
func (this *QWindowsMimeConverter) OnFormatsForMime(slot func(mimeType string, mimeData *QMimeData) []tagFORMATETC) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowsMimeConverter_override_virtual_FormatsForMime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowsMimeConverter_FormatsForMime
func miqt_exec_callback_QWindowsMimeConverter_FormatsForMime(self QWindowsMimeConverter, cb intptr_t, mimeType struct_miqt_string, mimeData *QMimeData) struct_miqt_array {
	gofunc, ok := cgo.Handle(cb).Value().(func(mimeType string, mimeData *QMimeData) []tagFORMATETC)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var mimeType_ms struct_miqt_string = mimeType
	mimeType_ret := GoStringN(mimeType_ms.data, int(int64(mimeType_ms.len)))
	free(unsafe.Pointer(mimeType_ms.data))
	slotval1 := mimeType_ret
	slotval2 := newQMimeData(mimeData)

	virtualReturn := gofunc(slotval1, slotval2)
	virtualReturn_CArray := (*[0xffff]tagFORMATETC)(malloc(size_t(8 * len(virtualReturn))))
	defer free(unsafe.Pointer(virtualReturn_CArray))
	for i := range virtualReturn {
		virtualReturn_CArray[i] = virtualReturn[i]
	}
	virtualReturn_ma := struct_miqt_array{len: size_t(len(virtualReturn)), data: unsafe.Pointer(virtualReturn_CArray)}

	return virtualReturn_ma

}
func (this *QWindowsMimeConverter) OnCanConvertToMime(slot func(mimeType string, pDataObj *IDataObject) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowsMimeConverter_override_virtual_CanConvertToMime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowsMimeConverter_CanConvertToMime
func miqt_exec_callback_QWindowsMimeConverter_CanConvertToMime(self QWindowsMimeConverter, cb intptr_t, mimeType struct_miqt_string, pDataObj *IDataObject) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(mimeType string, pDataObj *IDataObject) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var mimeType_ms struct_miqt_string = mimeType
	mimeType_ret := GoStringN(mimeType_ms.data, int(int64(mimeType_ms.len)))
	free(unsafe.Pointer(mimeType_ms.data))
	slotval1 := mimeType_ret
	xxxxxxxxx

	virtualReturn := gofunc(slotval1, slotval2)

	return (bool)(virtualReturn)

}
func (this *QWindowsMimeConverter) OnConvertToMime(slot func(mimeType string, pDataObj *IDataObject, preferredType QMetaType) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowsMimeConverter_override_virtual_ConvertToMime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowsMimeConverter_ConvertToMime
func miqt_exec_callback_QWindowsMimeConverter_ConvertToMime(self QWindowsMimeConverter, cb intptr_t, mimeType struct_miqt_string, pDataObj *IDataObject, preferredType *QMetaType) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(mimeType string, pDataObj *IDataObject, preferredType QMetaType) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var mimeType_ms struct_miqt_string = mimeType
	mimeType_ret := GoStringN(mimeType_ms.data, int(int64(mimeType_ms.len)))
	free(unsafe.Pointer(mimeType_ms.data))
	slotval1 := mimeType_ret
	xxxxxxxxx
	preferredType_goptr := newQMetaType(preferredType)
	preferredType_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval3 := *preferredType_goptr

	virtualReturn := gofunc(slotval1, slotval2, slotval3)

	return virtualReturn.cPointer()

}
func (this *QWindowsMimeConverter) OnMimeForFormat(slot func(formatetc *tagFORMATETC) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWindowsMimeConverter_override_virtual_MimeForFormat(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWindowsMimeConverter_MimeForFormat
func miqt_exec_callback_QWindowsMimeConverter_MimeForFormat(self QWindowsMimeConverter, cb intptr_t, formatetc *tagFORMATETC) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(formatetc *tagFORMATETC) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (*tagFORMATETC)(unsafe.Pointer(formatetc))

	virtualReturn := gofunc(slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}
