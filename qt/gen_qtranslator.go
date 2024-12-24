package qt

import (
	"unsafe"
)

type QTranslator struct {
	h          uintptr
	isSubclass bool
}

// NewQTranslator constructs a new QTranslator object.
func NewQTranslator() *QTranslator {
	ret := newQTranslator(QTranslator_new())
	ret.isSubclass = true
	return ret
}

// NewQTranslator2 constructs a new QTranslator object.
func NewQTranslator2(parent *QObject) *QTranslator {
	ret := newQTranslator(QTranslator_new2(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTranslator) MetaObject() *QMetaObject {
	return newQMetaObject(QTranslator_MetaObject(this.h))
}

func (this *QTranslator) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTranslator_Metacast(this.h, param1_Cstring))
}

func QTranslator_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTranslator_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	context_Cstring := CString(context)
	defer free(unsafe.Pointer(context_Cstring))
	sourceText_Cstring := CString(sourceText)
	defer free(unsafe.Pointer(sourceText_Cstring))
	disambiguation_Cstring := CString(disambiguation)
	defer free(unsafe.Pointer(disambiguation_Cstring))
	var _ms struct_miqt_string = QTranslator_Translate(this.h, context_Cstring, sourceText_Cstring, disambiguation_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) IsEmpty() bool {
	return (bool)(QTranslator_IsEmpty(this.h))
}

func (this *QTranslator) Language() string {
	var _ms struct_miqt_string = QTranslator_Language(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) FilePath() string {
	var _ms struct_miqt_string = QTranslator_FilePath(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) Load(filename string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QTranslator_Load(this.h, filename_ms))
}

func (this *QTranslator) Load2(locale *QLocale, filename string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	return (bool)(QTranslator_Load2(this.h, locale.cPointer(), filename_ms))
}

func (this *QTranslator) Load3(data *byte, lenVal int) bool {
	return (bool)(QTranslator_Load3(this.h, (*uchar)(unsafe.Pointer(data)), (int)(lenVal)))
}

func QTranslator_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTranslator_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTranslator_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTranslator_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTranslator) Load22(filename string, directory string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	return (bool)(QTranslator_Load22(this.h, filename_ms, directory_ms))
}

func (this *QTranslator) Load32(filename string, directory string, search_delimiters string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	search_delimiters_ms := struct_miqt_string{}
	search_delimiters_ms.data = CString(search_delimiters)
	search_delimiters_ms.len = size_t(len(search_delimiters))
	defer free(unsafe.Pointer(search_delimiters_ms.data))
	return (bool)(QTranslator_Load32(this.h, filename_ms, directory_ms, search_delimiters_ms))
}

func (this *QTranslator) Load4(filename string, directory string, search_delimiters string, suffix string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	search_delimiters_ms := struct_miqt_string{}
	search_delimiters_ms.data = CString(search_delimiters)
	search_delimiters_ms.len = size_t(len(search_delimiters))
	defer free(unsafe.Pointer(search_delimiters_ms.data))
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	return (bool)(QTranslator_Load4(this.h, filename_ms, directory_ms, search_delimiters_ms, suffix_ms))
}

func (this *QTranslator) Load33(locale *QLocale, filename string, prefix string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	return (bool)(QTranslator_Load33(this.h, locale.cPointer(), filename_ms, prefix_ms))
}

func (this *QTranslator) Load42(locale *QLocale, filename string, prefix string, directory string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	return (bool)(QTranslator_Load42(this.h, locale.cPointer(), filename_ms, prefix_ms, directory_ms))
}

func (this *QTranslator) Load5(locale *QLocale, filename string, prefix string, directory string, suffix string) bool {
	filename_ms := struct_miqt_string{}
	filename_ms.data = CString(filename)
	filename_ms.len = size_t(len(filename))
	defer free(unsafe.Pointer(filename_ms.data))
	prefix_ms := struct_miqt_string{}
	prefix_ms.data = CString(prefix)
	prefix_ms.len = size_t(len(prefix))
	defer free(unsafe.Pointer(prefix_ms.data))
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	suffix_ms := struct_miqt_string{}
	suffix_ms.data = CString(suffix)
	suffix_ms.len = size_t(len(suffix))
	defer free(unsafe.Pointer(suffix_ms.data))
	return (bool)(QTranslator_Load5(this.h, locale.cPointer(), filename_ms, prefix_ms, directory_ms, suffix_ms))
}

func (this *QTranslator) Load34(data *byte, lenVal int, directory string) bool {
	directory_ms := struct_miqt_string{}
	directory_ms.data = CString(directory)
	directory_ms.len = size_t(len(directory))
	defer free(unsafe.Pointer(directory_ms.data))
	return (bool)(QTranslator_Load34(this.h, (*uchar)(unsafe.Pointer(data)), (int)(lenVal), directory_ms))
}

func (this *QTranslator) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTranslator_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTranslator) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_MetaObject
func miqt_exec_callback_QTranslator_MetaObject(self QTranslator, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTranslator{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTranslator) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTranslator_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTranslator) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTranslator_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTranslator_Metacast
func miqt_exec_callback_QTranslator_Metacast(self QTranslator, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTranslator{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
