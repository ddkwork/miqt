package qt

import (
	"unsafe"
)

type QFontComboBox__FontFilter int

const (
	QFontComboBox__AllFonts          QFontComboBox__FontFilter = 0
	QFontComboBox__ScalableFonts     QFontComboBox__FontFilter = 1
	QFontComboBox__NonScalableFonts  QFontComboBox__FontFilter = 2
	QFontComboBox__MonospacedFonts   QFontComboBox__FontFilter = 4
	QFontComboBox__ProportionalFonts QFontComboBox__FontFilter = 8
)

type QFontComboBox struct {
	h          uintptr
	isSubclass bool
}

// NewQFontComboBox constructs a new QFontComboBox object.
func NewQFontComboBox(parent *QWidget) *QFontComboBox {
	ret := newQFontComboBox(QFontComboBox_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQFontComboBox2 constructs a new QFontComboBox object.
func NewQFontComboBox2() *QFontComboBox {
	ret := newQFontComboBox(QFontComboBox_new2())
	ret.isSubclass = true
	return ret
}

func (this *QFontComboBox) MetaObject() *QMetaObject {
	return newQMetaObject(QFontComboBox_MetaObject(this.h))
}

func (this *QFontComboBox) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFontComboBox_Metacast(this.h, param1_Cstring))
}

func QFontComboBox_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFontComboBox_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontComboBox) SetWritingSystem(writingSystem QFontDatabase__WritingSystem) {
	QFontComboBox_SetWritingSystem(this.h, (int)(writingSystem))
}

func (this *QFontComboBox) WritingSystem() QFontDatabase__WritingSystem {
	return (QFontDatabase__WritingSystem)(QFontComboBox_WritingSystem(this.h))
}

func (this *QFontComboBox) SetFontFilters(filters FontFilters) {
	QFontComboBox_SetFontFilters(this.h, filters)
}

func (this *QFontComboBox) FontFilters() FontFilters {
	xxxxxxxxx
}

func (this *QFontComboBox) CurrentFont() *QFont {
	_goptr := newQFont(QFontComboBox_CurrentFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontComboBox) SizeHint() *QSize {
	_goptr := newQSize(QFontComboBox_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontComboBox) SetSampleTextForSystem(writingSystem QFontDatabase__WritingSystem, sampleText string) {
	sampleText_ms := struct_miqt_string{}
	sampleText_ms.data = CString(sampleText)
	sampleText_ms.len = size_t(len(sampleText))
	defer free(unsafe.Pointer(sampleText_ms.data))
	QFontComboBox_SetSampleTextForSystem(this.h, (int)(writingSystem), sampleText_ms)
}

func (this *QFontComboBox) SampleTextForSystem(writingSystem QFontDatabase__WritingSystem) string {
	var _ms struct_miqt_string = QFontComboBox_SampleTextForSystem(this.h, (int)(writingSystem))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontComboBox) SetSampleTextForFont(fontFamily string, sampleText string) {
	fontFamily_ms := struct_miqt_string{}
	fontFamily_ms.data = CString(fontFamily)
	fontFamily_ms.len = size_t(len(fontFamily))
	defer free(unsafe.Pointer(fontFamily_ms.data))
	sampleText_ms := struct_miqt_string{}
	sampleText_ms.data = CString(sampleText)
	sampleText_ms.len = size_t(len(sampleText))
	defer free(unsafe.Pointer(sampleText_ms.data))
	QFontComboBox_SetSampleTextForFont(this.h, fontFamily_ms, sampleText_ms)
}

func (this *QFontComboBox) SampleTextForFont(fontFamily string) string {
	fontFamily_ms := struct_miqt_string{}
	fontFamily_ms.data = CString(fontFamily)
	fontFamily_ms.len = size_t(len(fontFamily))
	defer free(unsafe.Pointer(fontFamily_ms.data))
	var _ms struct_miqt_string = QFontComboBox_SampleTextForFont(this.h, fontFamily_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontComboBox) SetDisplayFont(fontFamily string, font *QFont) {
	fontFamily_ms := struct_miqt_string{}
	fontFamily_ms.data = CString(fontFamily)
	fontFamily_ms.len = size_t(len(fontFamily))
	defer free(unsafe.Pointer(fontFamily_ms.data))
	QFontComboBox_SetDisplayFont(this.h, fontFamily_ms, font.cPointer())
}

func (this *QFontComboBox) SetCurrentFont(f *QFont) {
	QFontComboBox_SetCurrentFont(this.h, f.cPointer())
}

func (this *QFontComboBox) CurrentFontChanged(f *QFont) {
	QFontComboBox_CurrentFontChanged(this.h, f.cPointer())
}

func (this *QFontComboBox) OnCurrentFontChanged(slot func(f *QFont)) {
	QFontComboBox_connect_CurrentFontChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_CurrentFontChanged
func miqt_exec_callback_QFontComboBox_CurrentFontChanged(cb intptr_t, f *QFont) {
	gofunc, ok := cgo.Handle(cb).Value().(func(f *QFont))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFont(f)

	gofunc(slotval1)
}

func QFontComboBox_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFontComboBox_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontComboBox_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFontComboBox_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontComboBox) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFontComboBox_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFontComboBox) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_MetaObject
func miqt_exec_callback_QFontComboBox_MetaObject(self QFontComboBox, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontComboBox{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFontComboBox) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFontComboBox_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFontComboBox) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_Metacast
func miqt_exec_callback_QFontComboBox_Metacast(self QFontComboBox, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QFontComboBox{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
