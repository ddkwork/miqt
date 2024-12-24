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

func (this *QFontComboBox) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QFontComboBox_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFontComboBox) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_SizeHint
func miqt_exec_callback_QFontComboBox_SizeHint(self QFontComboBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontComboBox{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QFontComboBox) callVirtualBase_Event(e *QEvent) bool {

	return (bool)(QFontComboBox_virtualbase_Event(unsafe.Pointer(this.h), e.cPointer()))

}
func (this *QFontComboBox) OnEvent(slot func(super func(e *QEvent) bool, e *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_Event
func miqt_exec_callback_QFontComboBox_Event(self QFontComboBox, cb intptr_t, e *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent) bool, e *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	virtualReturn := gofunc((&QFontComboBox{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QFontComboBox) callVirtualBase_SetModel(model *QAbstractItemModel) {

	QFontComboBox_virtualbase_SetModel(unsafe.Pointer(this.h), model.cPointer())

}
func (this *QFontComboBox) OnSetModel(slot func(super func(model *QAbstractItemModel), model *QAbstractItemModel)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_SetModel(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_SetModel
func miqt_exec_callback_QFontComboBox_SetModel(self QFontComboBox, cb intptr_t, model *QAbstractItemModel) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(model *QAbstractItemModel), model *QAbstractItemModel))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQAbstractItemModel(model)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_SetModel, slotval1)

}

func (this *QFontComboBox) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QFontComboBox_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFontComboBox) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_MinimumSizeHint
func miqt_exec_callback_QFontComboBox_MinimumSizeHint(self QFontComboBox, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QFontComboBox{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QFontComboBox) callVirtualBase_ShowPopup() {

	QFontComboBox_virtualbase_ShowPopup(unsafe.Pointer(this.h))

}
func (this *QFontComboBox) OnShowPopup(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_ShowPopup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_ShowPopup
func miqt_exec_callback_QFontComboBox_ShowPopup(self QFontComboBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFontComboBox{h: self}).callVirtualBase_ShowPopup)

}

func (this *QFontComboBox) callVirtualBase_HidePopup() {

	QFontComboBox_virtualbase_HidePopup(unsafe.Pointer(this.h))

}
func (this *QFontComboBox) OnHidePopup(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_HidePopup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_HidePopup
func miqt_exec_callback_QFontComboBox_HidePopup(self QFontComboBox, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QFontComboBox{h: self}).callVirtualBase_HidePopup)

}

func (this *QFontComboBox) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QFontComboBox_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QFontComboBox) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_InputMethodQuery
func miqt_exec_callback_QFontComboBox_InputMethodQuery(self QFontComboBox, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QFontComboBox{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QFontComboBox) callVirtualBase_FocusInEvent(e *QFocusEvent) {

	QFontComboBox_virtualbase_FocusInEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnFocusInEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_FocusInEvent
func miqt_exec_callback_QFontComboBox_FocusInEvent(self QFontComboBox, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_FocusOutEvent(e *QFocusEvent) {

	QFontComboBox_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnFocusOutEvent(slot func(super func(e *QFocusEvent), e *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_FocusOutEvent
func miqt_exec_callback_QFontComboBox_FocusOutEvent(self QFontComboBox, cb intptr_t, e *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QFocusEvent), e *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_ChangeEvent(e *QEvent) {

	QFontComboBox_virtualbase_ChangeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnChangeEvent(slot func(super func(e *QEvent), e *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_ChangeEvent
func miqt_exec_callback_QFontComboBox_ChangeEvent(self QFontComboBox, cb intptr_t, e *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QEvent), e *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_ResizeEvent(e *QResizeEvent) {

	QFontComboBox_virtualbase_ResizeEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnResizeEvent(slot func(super func(e *QResizeEvent), e *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_ResizeEvent
func miqt_exec_callback_QFontComboBox_ResizeEvent(self QFontComboBox, cb intptr_t, e *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QResizeEvent), e *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_PaintEvent(e *QPaintEvent) {

	QFontComboBox_virtualbase_PaintEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnPaintEvent(slot func(super func(e *QPaintEvent), e *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_PaintEvent
func miqt_exec_callback_QFontComboBox_PaintEvent(self QFontComboBox, cb intptr_t, e *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QPaintEvent), e *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_ShowEvent(e *QShowEvent) {

	QFontComboBox_virtualbase_ShowEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnShowEvent(slot func(super func(e *QShowEvent), e *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_ShowEvent
func miqt_exec_callback_QFontComboBox_ShowEvent(self QFontComboBox, cb intptr_t, e *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QShowEvent), e *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_HideEvent(e *QHideEvent) {

	QFontComboBox_virtualbase_HideEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnHideEvent(slot func(super func(e *QHideEvent), e *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_HideEvent
func miqt_exec_callback_QFontComboBox_HideEvent(self QFontComboBox, cb intptr_t, e *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QHideEvent), e *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_MousePressEvent(e *QMouseEvent) {

	QFontComboBox_virtualbase_MousePressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnMousePressEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_MousePressEvent
func miqt_exec_callback_QFontComboBox_MousePressEvent(self QFontComboBox, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_MouseReleaseEvent(e *QMouseEvent) {

	QFontComboBox_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnMouseReleaseEvent(slot func(super func(e *QMouseEvent), e *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_MouseReleaseEvent
func miqt_exec_callback_QFontComboBox_MouseReleaseEvent(self QFontComboBox, cb intptr_t, e *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QMouseEvent), e *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_KeyPressEvent(e *QKeyEvent) {

	QFontComboBox_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnKeyPressEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_KeyPressEvent
func miqt_exec_callback_QFontComboBox_KeyPressEvent(self QFontComboBox, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_KeyReleaseEvent(e *QKeyEvent) {

	QFontComboBox_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnKeyReleaseEvent(slot func(super func(e *QKeyEvent), e *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_KeyReleaseEvent
func miqt_exec_callback_QFontComboBox_KeyReleaseEvent(self QFontComboBox, cb intptr_t, e *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QKeyEvent), e *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_WheelEvent(e *QWheelEvent) {

	QFontComboBox_virtualbase_WheelEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnWheelEvent(slot func(super func(e *QWheelEvent), e *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_WheelEvent
func miqt_exec_callback_QFontComboBox_WheelEvent(self QFontComboBox, cb intptr_t, e *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QWheelEvent), e *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_ContextMenuEvent(e *QContextMenuEvent) {

	QFontComboBox_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), e.cPointer())

}
func (this *QFontComboBox) OnContextMenuEvent(slot func(super func(e *QContextMenuEvent), e *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_ContextMenuEvent
func miqt_exec_callback_QFontComboBox_ContextMenuEvent(self QFontComboBox, cb intptr_t, e *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(e *QContextMenuEvent), e *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(e)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QFontComboBox_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QFontComboBox) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_InputMethodEvent
func miqt_exec_callback_QFontComboBox_InputMethodEvent(self QFontComboBox, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QFontComboBox) callVirtualBase_InitStyleOption(option *QStyleOptionComboBox) {

	QFontComboBox_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QFontComboBox) OnInitStyleOption(slot func(super func(option *QStyleOptionComboBox), option *QStyleOptionComboBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QFontComboBox_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QFontComboBox_InitStyleOption
func miqt_exec_callback_QFontComboBox_InitStyleOption(self QFontComboBox, cb intptr_t, option *QStyleOptionComboBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionComboBox), option *QStyleOptionComboBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionComboBox(option)

	gofunc((&QFontComboBox{h: self}).callVirtualBase_InitStyleOption, slotval1)

}
