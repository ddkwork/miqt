package qt

import (
	"unsafe"
)

type QFontDialog__FontDialogOption int

const (
	QFontDialog__NoButtons           QFontDialog__FontDialogOption = 1
	QFontDialog__DontUseNativeDialog QFontDialog__FontDialogOption = 2
	QFontDialog__ScalableFonts       QFontDialog__FontDialogOption = 4
	QFontDialog__NonScalableFonts    QFontDialog__FontDialogOption = 8
	QFontDialog__MonospacedFonts     QFontDialog__FontDialogOption = 16
	QFontDialog__ProportionalFonts   QFontDialog__FontDialogOption = 32
)

type QFontDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQFontDialog constructs a new QFontDialog object.
func NewQFontDialog(parent *QWidget) *QFontDialog {
	g := newQFontDialog(QFontDialog_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQFontDialog2 constructs a new QFontDialog object.
func NewQFontDialog2() *QFontDialog {
	g := newQFontDialog(QFontDialog_new2())
	g.isSubclass = true
	return g
}

// NewQFontDialog3 constructs a new QFontDialog object.
func NewQFontDialog3(initial *QFont) *QFontDialog {
	g := newQFontDialog(QFontDialog_new3(initial.cPointer()))
	g.isSubclass = true
	return g
}

// NewQFontDialog4 constructs a new QFontDialog object.
func NewQFontDialog4(initial *QFont, parent *QWidget) *QFontDialog {
	g := newQFontDialog(QFontDialog_new4(initial.cPointer(), parent.cPointer()))
	g.isSubclass = true
	return g
}

func (this *QFontDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QFontDialog_MetaObject(this.h))
}

func (this *QFontDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QFontDialog_Metacast(this.h, param1_Cstring))
}

func QFontDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QFontDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontDialog) SetCurrentFont(font *QFont) {
	QFontDialog_SetCurrentFont(this.h, font.cPointer())
}

func (this *QFontDialog) CurrentFont() *QFont {
	_goptr := newQFont(QFontDialog_CurrentFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) SelectedFont() *QFont {
	_goptr := newQFont(QFontDialog_SelectedFont(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) SetOption(option FontDialogOption) {
	QFontDialog_SetOption(this.h, option)
}

func (this *QFontDialog) TestOption(option FontDialogOption) bool {
	return (bool)(QFontDialog_TestOption(this.h, option))
}

func (this *QFontDialog) SetOptions(options FontDialogOptions) {
	QFontDialog_SetOptions(this.h, options)
}

func (this *QFontDialog) Options() FontDialogOptions {
	xxxxxxxxx
}

func (this *QFontDialog) SetVisible(visible bool) {
	QFontDialog_SetVisible(this.h, (bool)(visible))
}

func QFontDialog_GetFont(ok *bool) *QFont {
	_goptr := newQFont(QFontDialog_GetFont((*bool)(unsafe.Pointer(ok))))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont2(ok *bool, initial *QFont) *QFont {
	_goptr := newQFont(QFontDialog_GetFont2((*bool)(unsafe.Pointer(ok)), initial.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) CurrentFontChanged(font *QFont) {
	QFontDialog_CurrentFontChanged(this.h, font.cPointer())
}

func (this *QFontDialog) OnCurrentFontChanged(slot func(font *QFont)) {
	QFontDialog_connect_CurrentFontChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_CurrentFontChanged
func miqt_exec_callback_QFontDialog_CurrentFontChanged(cb intptr_t, font *QFont) {
	gofunc, ok := cgo.Handle(cb).Value().(func(font *QFont))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFont(font)

	gofunc(slotval1)
}

func (this *QFontDialog) FontSelected(font *QFont) {
	QFontDialog_FontSelected(this.h, font.cPointer())
}

func (this *QFontDialog) OnFontSelected(slot func(font *QFont)) {
	QFontDialog_connect_FontSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_FontSelected
func miqt_exec_callback_QFontDialog_FontSelected(cb intptr_t, font *QFont) {
	gofunc, ok := cgo.Handle(cb).Value().(func(font *QFont))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFont(font)

	gofunc(slotval1)
}

func QFontDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFontDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QFontDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QFontDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QFontDialog) SetOption2(option FontDialogOption, on bool) {
	QFontDialog_SetOption2(this.h, option, (bool)(on))
}

func QFontDialog_GetFont22(ok *bool, parent *QWidget) *QFont {
	_goptr := newQFont(QFontDialog_GetFont22((*bool)(unsafe.Pointer(ok)), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont3(ok *bool, initial *QFont, parent *QWidget) *QFont {
	_goptr := newQFont(QFontDialog_GetFont3((*bool)(unsafe.Pointer(ok)), initial.cPointer(), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont4(ok *bool, initial *QFont, parent *QWidget, title string) *QFont {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQFont(QFontDialog_GetFont4((*bool)(unsafe.Pointer(ok)), initial.cPointer(), parent.cPointer(), title_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QFontDialog_GetFont5(ok *bool, initial *QFont, parent *QWidget, title string, options FontDialogOptions) *QFont {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQFont(QFontDialog_GetFont5((*bool)(unsafe.Pointer(ok)), initial.cPointer(), parent.cPointer(), title_ms, options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QFontDialog) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QFontDialog_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QFontDialog) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_MetaObject
func miqt_exec_callback_QFontDialog_MetaObject(self QFontDialog, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontDialog{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QFontDialog) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QFontDialog_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QFontDialog) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontDialog_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontDialog_Metacast
func miqt_exec_callback_QFontDialog_Metacast(self QFontDialog, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QFontDialog{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
