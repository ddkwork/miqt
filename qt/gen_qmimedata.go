package qt

import (
	"unsafe"
)

type QMimeData struct {
	h          uintptr
	isSubclass bool
}

// NewQMimeData constructs a new QMimeData object.
func NewQMimeData() *QMimeData {
	g := newQMimeData(QMimeData_new())
	g.isSubclass = true
	return g
}

func (this *QMimeData) MetaObject() *QMetaObject {
	return newQMetaObject(QMimeData_MetaObject(this.h))
}

func (this *QMimeData) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QMimeData_Metacast(this.h, param1_Cstring))
}

func QMimeData_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QMimeData_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeData) Urls() []QUrl {
	var _ma struct_miqt_array = QMimeData_Urls(this.h)
	_ret := make([]QUrl, int(_ma.len))
	_outCast := (*[0xffff]*QUrl)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQUrl(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QMimeData) SetUrls(urls []QUrl) {
	urls_CArray := (*[0xffff]*QUrl)(malloc(size_t(8 * len(urls))))
	defer free(unsafe.Pointer(urls_CArray))
	for i := range urls {
		urls_CArray[i] = urls[i].cPointer()
	}
	urls_ma := struct_miqt_array{len: size_t(len(urls)), data: unsafe.Pointer(urls_CArray)}
	QMimeData_SetUrls(this.h, urls_ma)
}

func (this *QMimeData) HasUrls() bool {
	return (bool)(QMimeData_HasUrls(this.h))
}

func (this *QMimeData) Text() string {
	var _ms struct_miqt_string = QMimeData_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeData) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QMimeData_SetText(this.h, text_ms)
}

func (this *QMimeData) HasText() bool {
	return (bool)(QMimeData_HasText(this.h))
}

func (this *QMimeData) Html() string {
	var _ms struct_miqt_string = QMimeData_Html(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeData) SetHtml(html string) {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	QMimeData_SetHtml(this.h, html_ms)
}

func (this *QMimeData) HasHtml() bool {
	return (bool)(QMimeData_HasHtml(this.h))
}

func (this *QMimeData) ImageData() *QVariant {
	_goptr := newQVariant(QMimeData_ImageData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeData) SetImageData(image *QVariant) {
	QMimeData_SetImageData(this.h, image.cPointer())
}

func (this *QMimeData) HasImage() bool {
	return (bool)(QMimeData_HasImage(this.h))
}

func (this *QMimeData) ColorData() *QVariant {
	_goptr := newQVariant(QMimeData_ColorData(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QMimeData) SetColorData(color *QVariant) {
	QMimeData_SetColorData(this.h, color.cPointer())
}

func (this *QMimeData) HasColor() bool {
	return (bool)(QMimeData_HasColor(this.h))
}

func (this *QMimeData) Data(mimetype string) []byte {
	mimetype_ms := struct_miqt_string{}
	mimetype_ms.data = CString(mimetype)
	mimetype_ms.len = size_t(len(mimetype))
	defer free(unsafe.Pointer(mimetype_ms.data))
	var _bytearray struct_miqt_string = QMimeData_Data(this.h, mimetype_ms)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QMimeData) SetData(mimetype string, data []byte) {
	mimetype_ms := struct_miqt_string{}
	mimetype_ms.data = CString(mimetype)
	mimetype_ms.len = size_t(len(mimetype))
	defer free(unsafe.Pointer(mimetype_ms.data))
	data_alias := struct_miqt_string{}
	data_alias.data = (char)(unsafe.Pointer(&data[0]))
	data_alias.len = size_t(len(data))
	QMimeData_SetData(this.h, mimetype_ms, data_alias)
}

func (this *QMimeData) RemoveFormat(mimetype string) {
	mimetype_ms := struct_miqt_string{}
	mimetype_ms.data = CString(mimetype)
	mimetype_ms.len = size_t(len(mimetype))
	defer free(unsafe.Pointer(mimetype_ms.data))
	QMimeData_RemoveFormat(this.h, mimetype_ms)
}

func (this *QMimeData) HasFormat(mimetype string) bool {
	mimetype_ms := struct_miqt_string{}
	mimetype_ms.data = CString(mimetype)
	mimetype_ms.len = size_t(len(mimetype))
	defer free(unsafe.Pointer(mimetype_ms.data))
	return (bool)(QMimeData_HasFormat(this.h, mimetype_ms))
}

func (this *QMimeData) Formats() []string {
	var _ma struct_miqt_array = QMimeData_Formats(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QMimeData) Clear() {
	QMimeData_Clear(this.h)
}

func QMimeData_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMimeData_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QMimeData_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QMimeData_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QMimeData) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QMimeData_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QMimeData) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMimeData_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMimeData_MetaObject
func miqt_exec_callback_QMimeData_MetaObject(self QMimeData, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QMimeData{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QMimeData) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QMimeData_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QMimeData) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QMimeData_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QMimeData_Metacast
func miqt_exec_callback_QMimeData_Metacast(self QMimeData, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QMimeData{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
