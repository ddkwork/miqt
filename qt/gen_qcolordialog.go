package qt

import (
	"unsafe"
)

type QColorDialog__ColorDialogOption int

const (
	QColorDialog__ShowAlphaChannel    QColorDialog__ColorDialogOption = 1
	QColorDialog__NoButtons           QColorDialog__ColorDialogOption = 2
	QColorDialog__DontUseNativeDialog QColorDialog__ColorDialogOption = 4
	QColorDialog__NoEyeDropperButton  QColorDialog__ColorDialogOption = 8
)

type QColorDialog struct {
	h          uintptr
	isSubclass bool
}

// NewQColorDialog constructs a new QColorDialog object.
func NewQColorDialog(parent *QWidget) *QColorDialog {
	ret := newQColorDialog(QColorDialog_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQColorDialog2 constructs a new QColorDialog object.
func NewQColorDialog2() *QColorDialog {
	ret := newQColorDialog(QColorDialog_new2())
	ret.isSubclass = true
	return ret
}

// NewQColorDialog3 constructs a new QColorDialog object.
func NewQColorDialog3(initial *QColor) *QColorDialog {
	ret := newQColorDialog(QColorDialog_new3(initial.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQColorDialog4 constructs a new QColorDialog object.
func NewQColorDialog4(initial *QColor, parent *QWidget) *QColorDialog {
	ret := newQColorDialog(QColorDialog_new4(initial.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QColorDialog) MetaObject() *QMetaObject {
	return newQMetaObject(QColorDialog_MetaObject(this.h))
}

func (this *QColorDialog) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QColorDialog_Metacast(this.h, param1_Cstring))
}

func QColorDialog_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QColorDialog_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColorDialog) SetCurrentColor(color *QColor) {
	QColorDialog_SetCurrentColor(this.h, color.cPointer())
}

func (this *QColorDialog) CurrentColor() *QColor {
	_goptr := newQColor(QColorDialog_CurrentColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorDialog) SelectedColor() *QColor {
	_goptr := newQColor(QColorDialog_SelectedColor(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorDialog) SetOption(option ColorDialogOption) {
	QColorDialog_SetOption(this.h, option)
}

func (this *QColorDialog) TestOption(option ColorDialogOption) bool {
	return (bool)(QColorDialog_TestOption(this.h, option))
}

func (this *QColorDialog) SetOptions(options ColorDialogOptions) {
	QColorDialog_SetOptions(this.h, options)
}

func (this *QColorDialog) Options() ColorDialogOptions {
	xxxxxxxxx
}

func (this *QColorDialog) SetVisible(visible bool) {
	QColorDialog_SetVisible(this.h, (bool)(visible))
}

func QColorDialog_GetColor() *QColor {
	_goptr := newQColor(QColorDialog_GetColor())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_CustomCount() int {
	return (int)(QColorDialog_CustomCount())
}

func QColorDialog_CustomColor(index int) *QColor {
	_goptr := newQColor(QColorDialog_CustomColor((int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_SetCustomColor(index int, color QColor) {
	QColorDialog_SetCustomColor((int)(index), color.cPointer())
}

func QColorDialog_StandardColor(index int) *QColor {
	_goptr := newQColor(QColorDialog_StandardColor((int)(index)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_SetStandardColor(index int, color QColor) {
	QColorDialog_SetStandardColor((int)(index), color.cPointer())
}

func (this *QColorDialog) CurrentColorChanged(color *QColor) {
	QColorDialog_CurrentColorChanged(this.h, color.cPointer())
}

func (this *QColorDialog) OnCurrentColorChanged(slot func(color *QColor)) {
	QColorDialog_connect_CurrentColorChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_CurrentColorChanged
func miqt_exec_callback_QColorDialog_CurrentColorChanged(cb intptr_t, color *QColor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(color *QColor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQColor(color)

	gofunc(slotval1)
}

func (this *QColorDialog) ColorSelected(color *QColor) {
	QColorDialog_ColorSelected(this.h, color.cPointer())
}

func (this *QColorDialog) OnColorSelected(slot func(color *QColor)) {
	QColorDialog_connect_ColorSelected(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_ColorSelected
func miqt_exec_callback_QColorDialog_ColorSelected(cb intptr_t, color *QColor) {
	gofunc, ok := cgo.Handle(cb).Value().(func(color *QColor))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQColor(color)

	gofunc(slotval1)
}

func QColorDialog_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QColorDialog_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QColorDialog_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QColorDialog_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QColorDialog) SetOption2(option ColorDialogOption, on bool) {
	QColorDialog_SetOption2(this.h, option, (bool)(on))
}

func QColorDialog_GetColor1(initial *QColor) *QColor {
	_goptr := newQColor(QColorDialog_GetColor1(initial.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_GetColor2(initial *QColor, parent *QWidget) *QColor {
	_goptr := newQColor(QColorDialog_GetColor2(initial.cPointer(), parent.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_GetColor3(initial *QColor, parent *QWidget, title string) *QColor {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQColor(QColorDialog_GetColor3(initial.cPointer(), parent.cPointer(), title_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QColorDialog_GetColor4(initial *QColor, parent *QWidget, title string, options ColorDialogOptions) *QColor {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	_goptr := newQColor(QColorDialog_GetColor4(initial.cPointer(), parent.cPointer(), title_ms, options))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QColorDialog) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QColorDialog_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QColorDialog) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_MetaObject
func miqt_exec_callback_QColorDialog_MetaObject(self QColorDialog, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QColorDialog{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QColorDialog) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QColorDialog_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QColorDialog) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QColorDialog_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QColorDialog_Metacast
func miqt_exec_callback_QColorDialog_Metacast(self QColorDialog, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QColorDialog{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
