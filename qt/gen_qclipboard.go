package qt

import (
	"unsafe"
)

type QClipboard__Mode int

const (
	QClipboard__Clipboard  QClipboard__Mode = 0
	QClipboard__Selection  QClipboard__Mode = 1
	QClipboard__FindBuffer QClipboard__Mode = 2
	QClipboard__LastMode   QClipboard__Mode = 2
)

type QClipboard struct {
	h          uintptr
	isSubclass bool
}

func (this *QClipboard) MetaObject() *QMetaObject {
	return newQMetaObject(QClipboard_MetaObject(this.h))
}

func (this *QClipboard) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QClipboard_Metacast(this.h, param1_Cstring))
}

func QClipboard_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QClipboard_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QClipboard) Clear() {
	QClipboard_Clear(this.h)
}

func (this *QClipboard) SupportsSelection() bool {
	return (bool)(QClipboard_SupportsSelection(this.h))
}

func (this *QClipboard) SupportsFindBuffer() bool {
	return (bool)(QClipboard_SupportsFindBuffer(this.h))
}

func (this *QClipboard) OwnsSelection() bool {
	return (bool)(QClipboard_OwnsSelection(this.h))
}

func (this *QClipboard) OwnsClipboard() bool {
	return (bool)(QClipboard_OwnsClipboard(this.h))
}

func (this *QClipboard) OwnsFindBuffer() bool {
	return (bool)(QClipboard_OwnsFindBuffer(this.h))
}

func (this *QClipboard) Text() string {
	var _ms struct_miqt_string = QClipboard_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QClipboard) TextWithSubtype(subtype string) string {
	subtype_ms := struct_miqt_string{}
	subtype_ms.data = CString(subtype)
	subtype_ms.len = size_t(len(subtype))
	defer free(unsafe.Pointer(subtype_ms.data))
	var _ms struct_miqt_string = QClipboard_TextWithSubtype(this.h, subtype_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QClipboard) SetText(param1 string) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QClipboard_SetText(this.h, param1_ms)
}

func (this *QClipboard) MimeData() *QMimeData {
	return newQMimeData(QClipboard_MimeData(this.h))
}

func (this *QClipboard) SetMimeData(data *QMimeData) {
	QClipboard_SetMimeData(this.h, data.cPointer())
}

func (this *QClipboard) Image() *QImage {
	_goptr := newQImage(QClipboard_Image(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QClipboard) Pixmap() *QPixmap {
	_goptr := newQPixmap(QClipboard_Pixmap(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QClipboard) SetImage(param1 *QImage) {
	QClipboard_SetImage(this.h, param1.cPointer())
}

func (this *QClipboard) SetPixmap(param1 *QPixmap) {
	QClipboard_SetPixmap(this.h, param1.cPointer())
}

func (this *QClipboard) Changed(mode QClipboard__Mode) {
	QClipboard_Changed(this.h, (int)(mode))
}

func (this *QClipboard) OnChanged(slot func(mode QClipboard__Mode)) {
	QClipboard_connect_Changed(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QClipboard_Changed
func miqt_exec_callback_QClipboard_Changed(cb intptr_t, mode int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(mode QClipboard__Mode))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (QClipboard__Mode)(mode)

	gofunc(slotval1)
}

func (this *QClipboard) SelectionChanged() {
	QClipboard_SelectionChanged(this.h)
}

func (this *QClipboard) OnSelectionChanged(slot func()) {
	QClipboard_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QClipboard_SelectionChanged
func miqt_exec_callback_QClipboard_SelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QClipboard) FindBufferChanged() {
	QClipboard_FindBufferChanged(this.h)
}

func (this *QClipboard) OnFindBufferChanged(slot func()) {
	QClipboard_connect_FindBufferChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QClipboard_FindBufferChanged
func miqt_exec_callback_QClipboard_FindBufferChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QClipboard) DataChanged() {
	QClipboard_DataChanged(this.h)
}

func (this *QClipboard) OnDataChanged(slot func()) {
	QClipboard_connect_DataChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QClipboard_DataChanged
func miqt_exec_callback_QClipboard_DataChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QClipboard_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QClipboard_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QClipboard_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QClipboard_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QClipboard) Clear1(mode Mode) {
	QClipboard_Clear1(this.h, mode)
}

func (this *QClipboard) Text1(mode Mode) string {
	var _ms struct_miqt_string = QClipboard_Text1(this.h, mode)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QClipboard) Text2(subtype string, mode Mode) string {
	subtype_ms := struct_miqt_string{}
	subtype_ms.data = CString(subtype)
	subtype_ms.len = size_t(len(subtype))
	defer free(unsafe.Pointer(subtype_ms.data))
	var _ms struct_miqt_string = QClipboard_Text2(this.h, subtype_ms, mode)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QClipboard) SetText2(param1 string, mode Mode) {
	param1_ms := struct_miqt_string{}
	param1_ms.data = CString(param1)
	param1_ms.len = size_t(len(param1))
	defer free(unsafe.Pointer(param1_ms.data))
	QClipboard_SetText2(this.h, param1_ms, mode)
}

func (this *QClipboard) MimeData1(mode Mode) *QMimeData {
	return newQMimeData(QClipboard_MimeData1(this.h, mode))
}

func (this *QClipboard) SetMimeData2(data *QMimeData, mode Mode) {
	QClipboard_SetMimeData2(this.h, data.cPointer(), mode)
}

func (this *QClipboard) Image1(mode Mode) *QImage {
	_goptr := newQImage(QClipboard_Image1(this.h, mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QClipboard) Pixmap1(mode Mode) *QPixmap {
	_goptr := newQPixmap(QClipboard_Pixmap1(this.h, mode))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QClipboard) SetImage2(param1 *QImage, mode Mode) {
	QClipboard_SetImage2(this.h, param1.cPointer(), mode)
}

func (this *QClipboard) SetPixmap2(param1 *QPixmap, mode Mode) {
	QClipboard_SetPixmap2(this.h, param1.cPointer(), mode)
}
