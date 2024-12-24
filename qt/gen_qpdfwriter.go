package qt

import (
	"unsafe"
)

type QPdfWriter__ColorModel int

const (
	QPdfWriter__RGB       QPdfWriter__ColorModel = 0
	QPdfWriter__Grayscale QPdfWriter__ColorModel = 1
	QPdfWriter__CMYK      QPdfWriter__ColorModel = 2
	QPdfWriter__Auto      QPdfWriter__ColorModel = 3
)

type QPdfWriter struct {
	h          uintptr
	isSubclass bool
}

// NewQPdfWriter constructs a new QPdfWriter object.
func NewQPdfWriter(filename string) *QPdfWriter {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))

	ret := newQPdfWriter(QPdfWriter_new(filename_ms))
	ret.isSubclass = true
	return ret
}

// NewQPdfWriter2 constructs a new QPdfWriter object.
func NewQPdfWriter2(device *QIODevice) *QPdfWriter {
	ret := newQPdfWriter(QPdfWriter_new2(device.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QPdfWriter) MetaObject() *QMetaObject {
	return newQMetaObject(QPdfWriter_MetaObject(this.h))
}

func (this *QPdfWriter) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QPdfWriter_Metacast(this.h, param1_Cstring))
}

func QPdfWriter_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QPdfWriter_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfWriter) SetPdfVersion(version PdfVersion) {
	QPdfWriter_SetPdfVersion(this.h, version)
}

func (this *QPdfWriter) PdfVersion() PdfVersion {
	xxxxxxxxx
}

func (this *QPdfWriter) Title() string {
	var _ms struct_miqt_string = QPdfWriter_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfWriter) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QPdfWriter_SetTitle(this.h, title_ms)
}

func (this *QPdfWriter) Creator() string {
	var _ms struct_miqt_string = QPdfWriter_Creator(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfWriter) SetCreator(creator string) {
	creator_ms := struct_miqt_string{}
	creator_ms.data = CString(creator)
	creator_ms.len = size_t(len(creator))
	defer free(unsafe.Pointer(creator_ms.data))
	QPdfWriter_SetCreator(this.h, creator_ms)
}

func (this *QPdfWriter) DocumentId() *QUuid {
	_goptr := newQUuid(QPdfWriter_DocumentId(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPdfWriter) SetDocumentId(documentId QUuid) {
	QPdfWriter_SetDocumentId(this.h, documentId.cPointer())
}

func (this *QPdfWriter) Author() string {
	var _ms struct_miqt_string = QPdfWriter_Author(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfWriter) SetAuthor(author string) {
	author_ms := struct_miqt_string{}
	author_ms.data = CString(author)
	author_ms.len = size_t(len(author))
	defer free(unsafe.Pointer(author_ms.data))
	QPdfWriter_SetAuthor(this.h, author_ms)
}

func (this *QPdfWriter) NewPage() bool {
	return (bool)(QPdfWriter_NewPage(this.h))
}

func (this *QPdfWriter) SetResolution(resolution int) {
	QPdfWriter_SetResolution(this.h, (int)(resolution))
}

func (this *QPdfWriter) Resolution() int {
	return (int)(QPdfWriter_Resolution(this.h))
}

func (this *QPdfWriter) SetDocumentXmpMetadata(xmpMetadata []byte) {
	xmpMetadata_alias := struct_miqt_string{}
	xmpMetadata_alias.data = (char)(unsafe.Pointer(&xmpMetadata[0]))
	xmpMetadata_alias.len = size_t(len(xmpMetadata))
	QPdfWriter_SetDocumentXmpMetadata(this.h, xmpMetadata_alias)
}

func (this *QPdfWriter) DocumentXmpMetadata() []byte {
	var _bytearray struct_miqt_string = QPdfWriter_DocumentXmpMetadata(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QPdfWriter) AddFileAttachment(fileName string, data []byte) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	QPdfWriter_AddFileAttachment(this.h, fileName_ms, data_alias)
}

func (this *QPdfWriter) ColorModel() ColorModel {
	xxxxxxxxx
}

func (this *QPdfWriter) SetColorModel(model ColorModel) {
	QPdfWriter_SetColorModel(this.h, model)
}

func (this *QPdfWriter) OutputIntent() *QPdfOutputIntent {
	_goptr := newQPdfOutputIntent(QPdfWriter_OutputIntent(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QPdfWriter) SetOutputIntent(intent *QPdfOutputIntent) {
	QPdfWriter_SetOutputIntent(this.h, intent.cPointer())
}

func QPdfWriter_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPdfWriter_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QPdfWriter_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QPdfWriter_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QPdfWriter) AddFileAttachment3(fileName string, data []byte, mimeType string) {
	fileName_ms := struct_miqt_string{}
	fileName_ms.data = CString(fileName)
	fileName_ms.len = size_t(len(fileName))
	defer free(unsafe.Pointer(fileName_ms.data))
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	mimeType_ms := struct_miqt_string{}
	mimeType_ms.data = CString(mimeType)
	mimeType_ms.len = size_t(len(mimeType))
	defer free(unsafe.Pointer(mimeType_ms.data))
	QPdfWriter_AddFileAttachment3(this.h, fileName_ms, data_alias, mimeType_ms)
}

func (this *QPdfWriter) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QPdfWriter_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QPdfWriter) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPdfWriter_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPdfWriter_MetaObject
func miqt_exec_callback_QPdfWriter_MetaObject(self QPdfWriter, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QPdfWriter{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QPdfWriter) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QPdfWriter_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QPdfWriter) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QPdfWriter_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QPdfWriter_Metacast
func miqt_exec_callback_QPdfWriter_Metacast(self QPdfWriter, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QPdfWriter{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
