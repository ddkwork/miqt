package qt

import (
	"unsafe"
)

type QWizard__WizardButton int

const (
	QWizard__BackButton       QWizard__WizardButton = 0
	QWizard__NextButton       QWizard__WizardButton = 1
	QWizard__CommitButton     QWizard__WizardButton = 2
	QWizard__FinishButton     QWizard__WizardButton = 3
	QWizard__CancelButton     QWizard__WizardButton = 4
	QWizard__HelpButton       QWizard__WizardButton = 5
	QWizard__CustomButton1    QWizard__WizardButton = 6
	QWizard__CustomButton2    QWizard__WizardButton = 7
	QWizard__CustomButton3    QWizard__WizardButton = 8
	QWizard__Stretch          QWizard__WizardButton = 9
	QWizard__NoButton         QWizard__WizardButton = -1
	QWizard__NStandardButtons QWizard__WizardButton = 6
	QWizard__NButtons         QWizard__WizardButton = 9
)

type QWizard__WizardPixmap int

const (
	QWizard__WatermarkPixmap  QWizard__WizardPixmap = 0
	QWizard__LogoPixmap       QWizard__WizardPixmap = 1
	QWizard__BannerPixmap     QWizard__WizardPixmap = 2
	QWizard__BackgroundPixmap QWizard__WizardPixmap = 3
	QWizard__NPixmaps         QWizard__WizardPixmap = 4
)

type QWizard__WizardStyle int

const (
	QWizard__ClassicStyle QWizard__WizardStyle = 0
	QWizard__ModernStyle  QWizard__WizardStyle = 1
	QWizard__MacStyle     QWizard__WizardStyle = 2
	QWizard__AeroStyle    QWizard__WizardStyle = 3
	QWizard__NStyles      QWizard__WizardStyle = 4
)

type QWizard__WizardOption int

const (
	QWizard__IndependentPages             QWizard__WizardOption = 1
	QWizard__IgnoreSubTitles              QWizard__WizardOption = 2
	QWizard__ExtendedWatermarkPixmap      QWizard__WizardOption = 4
	QWizard__NoDefaultButton              QWizard__WizardOption = 8
	QWizard__NoBackButtonOnStartPage      QWizard__WizardOption = 16
	QWizard__NoBackButtonOnLastPage       QWizard__WizardOption = 32
	QWizard__DisabledBackButtonOnLastPage QWizard__WizardOption = 64
	QWizard__HaveNextButtonOnLastPage     QWizard__WizardOption = 128
	QWizard__HaveFinishButtonOnEarlyPages QWizard__WizardOption = 256
	QWizard__NoCancelButton               QWizard__WizardOption = 512
	QWizard__CancelButtonOnLeft           QWizard__WizardOption = 1024
	QWizard__HaveHelpButton               QWizard__WizardOption = 2048
	QWizard__HelpButtonOnRight            QWizard__WizardOption = 4096
	QWizard__HaveCustomButton1            QWizard__WizardOption = 8192
	QWizard__HaveCustomButton2            QWizard__WizardOption = 16384
	QWizard__HaveCustomButton3            QWizard__WizardOption = 32768
	QWizard__NoCancelButtonOnLastPage     QWizard__WizardOption = 65536
)

type QWizard struct {
	h          uintptr
	isSubclass bool
}

// NewQWizard constructs a new QWizard object.
func NewQWizard(parent *QWidget) *QWizard {

	ret := newQWizard(QWizard_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQWizard2 constructs a new QWizard object.
func NewQWizard2() *QWizard {

	ret := newQWizard(QWizard_new2())
	ret.isSubclass = true
	return ret
}

// NewQWizard3 constructs a new QWizard object.
func NewQWizard3(parent *QWidget, flags WindowType) *QWizard {

	ret := newQWizard(QWizard_new3(parent.cPointer(), (int)(flags)))
	ret.isSubclass = true
	return ret
}

func (this *QWizard) MetaObject() *QMetaObject {
	return newQMetaObject(QWizard_MetaObject(this.h))
}

func (this *QWizard) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWizard_Metacast(this.h, param1_Cstring))
}

func QWizard_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWizard_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizard) AddPage(page *QWizardPage) int {
	return (int)(QWizard_AddPage(this.h, page.cPointer()))
}

func (this *QWizard) SetPage(id int, page *QWizardPage) {
	QWizard_SetPage(this.h, (int)(id), page.cPointer())
}

func (this *QWizard) RemovePage(id int) {
	QWizard_RemovePage(this.h, (int)(id))
}

func (this *QWizard) Page(id int) *QWizardPage {
	return newQWizardPage(QWizard_Page(this.h, (int)(id)))
}

func (this *QWizard) HasVisitedPage(id int) bool {
	return (bool)(QWizard_HasVisitedPage(this.h, (int)(id)))
}

func (this *QWizard) VisitedIds() []int {
	var _ma struct_miqt_array = QWizard_VisitedIds(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QWizard) PageIds() []int {
	var _ma struct_miqt_array = QWizard_PageIds(this.h)
	_ret := make([]int, int(_ma.len))
	_outCast := (*[0xffff]int)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = (int)(_outCast[i])
	}
	return _ret
}

func (this *QWizard) SetStartId(id int) {
	QWizard_SetStartId(this.h, (int)(id))
}

func (this *QWizard) StartId() int {
	return (int)(QWizard_StartId(this.h))
}

func (this *QWizard) CurrentPage() *QWizardPage {
	return newQWizardPage(QWizard_CurrentPage(this.h))
}

func (this *QWizard) CurrentId() int {
	return (int)(QWizard_CurrentId(this.h))
}

func (this *QWizard) ValidateCurrentPage() bool {
	return (bool)(QWizard_ValidateCurrentPage(this.h))
}

func (this *QWizard) NextId() int {
	return (int)(QWizard_NextId(this.h))
}

func (this *QWizard) SetField(name string, value *QVariant) {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	QWizard_SetField(this.h, name_ms, value.cPointer())
}

func (this *QWizard) Field(name string) *QVariant {
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	_goptr := newQVariant(QWizard_Field(this.h, name_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWizard) SetWizardStyle(style WizardStyle) {
	QWizard_SetWizardStyle(this.h, style)
}

func (this *QWizard) WizardStyle() WizardStyle {
	xxxxxxxxx
}

func (this *QWizard) SetOption(option WizardOption) {
	QWizard_SetOption(this.h, option)
}

func (this *QWizard) TestOption(option WizardOption) bool {
	return (bool)(QWizard_TestOption(this.h, option))
}

func (this *QWizard) SetOptions(options WizardOptions) {
	QWizard_SetOptions(this.h, options)
}

func (this *QWizard) Options() WizardOptions {
	xxxxxxxxx
}

func (this *QWizard) SetButtonText(which WizardButton, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QWizard_SetButtonText(this.h, which, text_ms)
}

func (this *QWizard) ButtonText(which WizardButton) string {
	var _ms struct_miqt_string = QWizard_ButtonText(this.h, which)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizard) SetButtonLayout(layout []WizardButton) {
	layout_CArray := (*[0xffff]WizardButton)(malloc(size_t(8 * len(layout))))
	defer free(unsafe.Pointer(layout_CArray))
	for i := range layout {
		layout_CArray[i] = layout[i]
	}
	layout_ma := struct_miqt_array{len: size_t(len(layout)), data: unsafe.Pointer(layout_CArray)}
	QWizard_SetButtonLayout(this.h, layout_ma)
}

func (this *QWizard) SetButton(which WizardButton, button *QAbstractButton) {
	QWizard_SetButton(this.h, which, button.cPointer())
}

func (this *QWizard) Button(which WizardButton) *QAbstractButton {
	return newQAbstractButton(QWizard_Button(this.h, which))
}

func (this *QWizard) SetTitleFormat(format TextFormat) {
	QWizard_SetTitleFormat(this.h, (int)(format))
}

func (this *QWizard) TitleFormat() TextFormat {
	return (TextFormat)(QWizard_TitleFormat(this.h))
}

func (this *QWizard) SetSubTitleFormat(format TextFormat) {
	QWizard_SetSubTitleFormat(this.h, (int)(format))
}

func (this *QWizard) SubTitleFormat() TextFormat {
	return (TextFormat)(QWizard_SubTitleFormat(this.h))
}

func (this *QWizard) SetPixmap(which WizardPixmap, pixmap *QPixmap) {
	QWizard_SetPixmap(this.h, which, pixmap.cPointer())
}

func (this *QWizard) Pixmap(which WizardPixmap) *QPixmap {
	_goptr := newQPixmap(QWizard_Pixmap(this.h, which))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWizard) SetSideWidget(widget *QWidget) {
	QWizard_SetSideWidget(this.h, widget.cPointer())
}

func (this *QWizard) SideWidget() *QWidget {
	return newQWidget(QWizard_SideWidget(this.h))
}

func (this *QWizard) SetDefaultProperty(className string, property string, changedSignal string) {
	className_Cstring := CString(className)
	defer free(unsafe.Pointer(className_Cstring))
	property_Cstring := CString(property)
	defer free(unsafe.Pointer(property_Cstring))
	changedSignal_Cstring := CString(changedSignal)
	defer free(unsafe.Pointer(changedSignal_Cstring))
	QWizard_SetDefaultProperty(this.h, className_Cstring, property_Cstring, changedSignal_Cstring)
}

func (this *QWizard) SetVisible(visible bool) {
	QWizard_SetVisible(this.h, (bool)(visible))
}

func (this *QWizard) SizeHint() *QSize {
	_goptr := newQSize(QWizard_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWizard) CurrentIdChanged(id int) {
	QWizard_CurrentIdChanged(this.h, (int)(id))
}
func (this *QWizard) OnCurrentIdChanged(slot func(id int)) {
	QWizard_connect_CurrentIdChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_CurrentIdChanged
func miqt_exec_callback_QWizard_CurrentIdChanged(cb intptr_t, id int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc(slotval1)
}

func (this *QWizard) HelpRequested() {
	QWizard_HelpRequested(this.h)
}
func (this *QWizard) OnHelpRequested(slot func()) {
	QWizard_connect_HelpRequested(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_HelpRequested
func miqt_exec_callback_QWizard_HelpRequested(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QWizard) CustomButtonClicked(which int) {
	QWizard_CustomButtonClicked(this.h, (int)(which))
}
func (this *QWizard) OnCustomButtonClicked(slot func(which int)) {
	QWizard_connect_CustomButtonClicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_CustomButtonClicked
func miqt_exec_callback_QWizard_CustomButtonClicked(cb intptr_t, which int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(which int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(which)

	gofunc(slotval1)
}

func (this *QWizard) PageAdded(id int) {
	QWizard_PageAdded(this.h, (int)(id))
}
func (this *QWizard) OnPageAdded(slot func(id int)) {
	QWizard_connect_PageAdded(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_PageAdded
func miqt_exec_callback_QWizard_PageAdded(cb intptr_t, id int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc(slotval1)
}

func (this *QWizard) PageRemoved(id int) {
	QWizard_PageRemoved(this.h, (int)(id))
}
func (this *QWizard) OnPageRemoved(slot func(id int)) {
	QWizard_connect_PageRemoved(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_PageRemoved
func miqt_exec_callback_QWizard_PageRemoved(cb intptr_t, id int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc(slotval1)
}

func (this *QWizard) Back() {
	QWizard_Back(this.h)
}

func (this *QWizard) Next() {
	QWizard_Next(this.h)
}

func (this *QWizard) SetCurrentId(id int) {
	QWizard_SetCurrentId(this.h, (int)(id))
}

func (this *QWizard) Restart() {
	QWizard_Restart(this.h)
}

func QWizard_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWizard_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWizard_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWizard_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizard) SetOption2(option WizardOption, on bool) {
	QWizard_SetOption2(this.h, option, (bool)(on))
}

func (this *QWizard) callVirtualBase_ValidateCurrentPage() bool {

	return (bool)(QWizard_virtualbase_ValidateCurrentPage(unsafe.Pointer(this.h)))

}
func (this *QWizard) OnValidateCurrentPage(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_ValidateCurrentPage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_ValidateCurrentPage
func miqt_exec_callback_QWizard_ValidateCurrentPage(self QWizard, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_ValidateCurrentPage)

	return (bool)(virtualReturn)

}

func (this *QWizard) callVirtualBase_NextId() int {

	return (int)(QWizard_virtualbase_NextId(unsafe.Pointer(this.h)))

}
func (this *QWizard) OnNextId(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_NextId(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_NextId
func miqt_exec_callback_QWizard_NextId(self QWizard, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_NextId)

	return (int)(virtualReturn)

}

func (this *QWizard) callVirtualBase_SetVisible(visible bool) {

	QWizard_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QWizard) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_SetVisible
func miqt_exec_callback_QWizard_SetVisible(self QWizard, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QWizard{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QWizard) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QWizard_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWizard) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_SizeHint
func miqt_exec_callback_QWizard_SizeHint(self QWizard, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QWizard) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QWizard_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QWizard) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Event
func miqt_exec_callback_QWizard_Event(self QWizard, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWizard) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QWizard_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizard) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_ResizeEvent
func miqt_exec_callback_QWizard_ResizeEvent(self QWizard, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QWizard{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QWizard) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QWizard_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizard) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_PaintEvent
func miqt_exec_callback_QWizard_PaintEvent(self QWizard, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QWizard{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QWizard) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QWizard_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QWizard) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_NativeEvent
func miqt_exec_callback_QWizard_NativeEvent(self QWizard, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QWizard) callVirtualBase_Done(result int) {

	QWizard_virtualbase_Done(unsafe.Pointer(this.h), (int)(result))

}
func (this *QWizard) OnDone(slot func(super func(result int), result int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Done(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Done
func miqt_exec_callback_QWizard_Done(self QWizard, cb intptr_t, result int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(result int), result int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(result)

	gofunc((&QWizard{h: self}).callVirtualBase_Done, slotval1)

}

func (this *QWizard) callVirtualBase_InitializePage(id int) {

	QWizard_virtualbase_InitializePage(unsafe.Pointer(this.h), (int)(id))

}
func (this *QWizard) OnInitializePage(slot func(super func(id int), id int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_InitializePage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_InitializePage
func miqt_exec_callback_QWizard_InitializePage(self QWizard, cb intptr_t, id int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(id int), id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc((&QWizard{h: self}).callVirtualBase_InitializePage, slotval1)

}

func (this *QWizard) callVirtualBase_CleanupPage(id int) {

	QWizard_virtualbase_CleanupPage(unsafe.Pointer(this.h), (int)(id))

}
func (this *QWizard) OnCleanupPage(slot func(super func(id int), id int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_CleanupPage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_CleanupPage
func miqt_exec_callback_QWizard_CleanupPage(self QWizard, cb intptr_t, id int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(id int), id int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(id)

	gofunc((&QWizard{h: self}).callVirtualBase_CleanupPage, slotval1)

}

func (this *QWizard) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QWizard_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWizard) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_MinimumSizeHint
func miqt_exec_callback_QWizard_MinimumSizeHint(self QWizard, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QWizard) callVirtualBase_Open() {

	QWizard_virtualbase_Open(unsafe.Pointer(this.h))

}
func (this *QWizard) OnOpen(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Open(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Open
func miqt_exec_callback_QWizard_Open(self QWizard, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QWizard{h: self}).callVirtualBase_Open)

}

func (this *QWizard) callVirtualBase_Exec() int {

	return (int)(QWizard_virtualbase_Exec(unsafe.Pointer(this.h)))

}
func (this *QWizard) OnExec(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Exec(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Exec
func miqt_exec_callback_QWizard_Exec(self QWizard, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_Exec)

	return (int)(virtualReturn)

}

func (this *QWizard) callVirtualBase_Accept() {

	QWizard_virtualbase_Accept(unsafe.Pointer(this.h))

}
func (this *QWizard) OnAccept(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Accept(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Accept
func miqt_exec_callback_QWizard_Accept(self QWizard, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QWizard{h: self}).callVirtualBase_Accept)

}

func (this *QWizard) callVirtualBase_Reject() {

	QWizard_virtualbase_Reject(unsafe.Pointer(this.h))

}
func (this *QWizard) OnReject(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Reject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Reject
func miqt_exec_callback_QWizard_Reject(self QWizard, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QWizard{h: self}).callVirtualBase_Reject)

}

func (this *QWizard) callVirtualBase_KeyPressEvent(param1 *QKeyEvent) {

	QWizard_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWizard) OnKeyPressEvent(slot func(super func(param1 *QKeyEvent), param1 *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_KeyPressEvent
func miqt_exec_callback_QWizard_KeyPressEvent(self QWizard, cb intptr_t, param1 *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QKeyEvent), param1 *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(param1)

	gofunc((&QWizard{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QWizard) callVirtualBase_CloseEvent(param1 *QCloseEvent) {

	QWizard_virtualbase_CloseEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWizard) OnCloseEvent(slot func(super func(param1 *QCloseEvent), param1 *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_CloseEvent
func miqt_exec_callback_QWizard_CloseEvent(self QWizard, cb intptr_t, param1 *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QCloseEvent), param1 *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(param1)

	gofunc((&QWizard{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QWizard) callVirtualBase_ShowEvent(param1 *QShowEvent) {

	QWizard_virtualbase_ShowEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWizard) OnShowEvent(slot func(super func(param1 *QShowEvent), param1 *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_ShowEvent
func miqt_exec_callback_QWizard_ShowEvent(self QWizard, cb intptr_t, param1 *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QShowEvent), param1 *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(param1)

	gofunc((&QWizard{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QWizard) callVirtualBase_ContextMenuEvent(param1 *QContextMenuEvent) {

	QWizard_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWizard) OnContextMenuEvent(slot func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_ContextMenuEvent
func miqt_exec_callback_QWizard_ContextMenuEvent(self QWizard, cb intptr_t, param1 *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QContextMenuEvent), param1 *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(param1)

	gofunc((&QWizard{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QWizard) callVirtualBase_EventFilter(param1 *QObject, param2 *QEvent) bool {

	return (bool)(QWizard_virtualbase_EventFilter(unsafe.Pointer(this.h), param1.cPointer(), param2.cPointer()))

}
func (this *QWizard) OnEventFilter(slot func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_EventFilter
func miqt_exec_callback_QWizard_EventFilter(self QWizard, cb intptr_t, param1 *QObject, param2 *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QObject, param2 *QEvent) bool, param1 *QObject, param2 *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(param1)

	slotval2 := newQEvent(param2)

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

type QWizardPage struct {
	h          uintptr
	isSubclass bool
}

// NewQWizardPage constructs a new QWizardPage object.
func NewQWizardPage(parent *QWidget) *QWizardPage {

	ret := newQWizardPage(QWizardPage_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQWizardPage2 constructs a new QWizardPage object.
func NewQWizardPage2() *QWizardPage {

	ret := newQWizardPage(QWizardPage_new2())
	ret.isSubclass = true
	return ret
}

func (this *QWizardPage) MetaObject() *QMetaObject {
	return newQMetaObject(QWizardPage_MetaObject(this.h))
}

func (this *QWizardPage) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QWizardPage_Metacast(this.h, param1_Cstring))
}

func QWizardPage_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QWizardPage_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizardPage) SetTitle(title string) {
	title_ms := struct_miqt_string{}
	title_ms.data = CString(title)
	title_ms.len = size_t(len(title))
	defer free(unsafe.Pointer(title_ms.data))
	QWizardPage_SetTitle(this.h, title_ms)
}

func (this *QWizardPage) Title() string {
	var _ms struct_miqt_string = QWizardPage_Title(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizardPage) SetSubTitle(subTitle string) {
	subTitle_ms := struct_miqt_string{}
	subTitle_ms.data = CString(subTitle)
	subTitle_ms.len = size_t(len(subTitle))
	defer free(unsafe.Pointer(subTitle_ms.data))
	QWizardPage_SetSubTitle(this.h, subTitle_ms)
}

func (this *QWizardPage) SubTitle() string {
	var _ms struct_miqt_string = QWizardPage_SubTitle(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizardPage) SetPixmap(which QWizard__WizardPixmap, pixmap *QPixmap) {
	QWizardPage_SetPixmap(this.h, (int)(which), pixmap.cPointer())
}

func (this *QWizardPage) Pixmap(which QWizard__WizardPixmap) *QPixmap {
	_goptr := newQPixmap(QWizardPage_Pixmap(this.h, (int)(which)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QWizardPage) SetFinalPage(finalPage bool) {
	QWizardPage_SetFinalPage(this.h, (bool)(finalPage))
}

func (this *QWizardPage) IsFinalPage() bool {
	return (bool)(QWizardPage_IsFinalPage(this.h))
}

func (this *QWizardPage) SetCommitPage(commitPage bool) {
	QWizardPage_SetCommitPage(this.h, (bool)(commitPage))
}

func (this *QWizardPage) IsCommitPage() bool {
	return (bool)(QWizardPage_IsCommitPage(this.h))
}

func (this *QWizardPage) SetButtonText(which QWizard__WizardButton, text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QWizardPage_SetButtonText(this.h, (int)(which), text_ms)
}

func (this *QWizardPage) ButtonText(which QWizard__WizardButton) string {
	var _ms struct_miqt_string = QWizardPage_ButtonText(this.h, (int)(which))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizardPage) InitializePage() {
	QWizardPage_InitializePage(this.h)
}

func (this *QWizardPage) CleanupPage() {
	QWizardPage_CleanupPage(this.h)
}

func (this *QWizardPage) ValidatePage() bool {
	return (bool)(QWizardPage_ValidatePage(this.h))
}

func (this *QWizardPage) IsComplete() bool {
	return (bool)(QWizardPage_IsComplete(this.h))
}

func (this *QWizardPage) NextId() int {
	return (int)(QWizardPage_NextId(this.h))
}

func (this *QWizardPage) CompleteChanged() {
	QWizardPage_CompleteChanged(this.h)
}
func (this *QWizardPage) OnCompleteChanged(slot func()) {
	QWizardPage_connect_CompleteChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_CompleteChanged
func miqt_exec_callback_QWizardPage_CompleteChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func QWizardPage_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWizardPage_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QWizardPage_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QWizardPage_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QWizardPage) callVirtualBase_InitializePage() {

	QWizardPage_virtualbase_InitializePage(unsafe.Pointer(this.h))

}
func (this *QWizardPage) OnInitializePage(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_InitializePage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_InitializePage
func miqt_exec_callback_QWizardPage_InitializePage(self QWizardPage, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QWizardPage{h: self}).callVirtualBase_InitializePage)

}

func (this *QWizardPage) callVirtualBase_CleanupPage() {

	QWizardPage_virtualbase_CleanupPage(unsafe.Pointer(this.h))

}
func (this *QWizardPage) OnCleanupPage(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_CleanupPage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_CleanupPage
func miqt_exec_callback_QWizardPage_CleanupPage(self QWizardPage, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QWizardPage{h: self}).callVirtualBase_CleanupPage)

}

func (this *QWizardPage) callVirtualBase_ValidatePage() bool {

	return (bool)(QWizardPage_virtualbase_ValidatePage(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnValidatePage(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_ValidatePage(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_ValidatePage
func miqt_exec_callback_QWizardPage_ValidatePage(self QWizardPage, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_ValidatePage)

	return (bool)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_IsComplete() bool {

	return (bool)(QWizardPage_virtualbase_IsComplete(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnIsComplete(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_IsComplete(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_IsComplete
func miqt_exec_callback_QWizardPage_IsComplete(self QWizardPage, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_IsComplete)

	return (bool)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_NextId() int {

	return (int)(QWizardPage_virtualbase_NextId(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnNextId(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_NextId(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_NextId
func miqt_exec_callback_QWizardPage_NextId(self QWizardPage, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_NextId)

	return (int)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_DevType() int {

	return (int)(QWizardPage_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_DevType
func miqt_exec_callback_QWizardPage_DevType(self QWizardPage, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_SetVisible(visible bool) {

	QWizardPage_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QWizardPage) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_SetVisible
func miqt_exec_callback_QWizardPage_SetVisible(self QWizardPage, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QWizardPage{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QWizardPage) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QWizardPage_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWizardPage) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_SizeHint
func miqt_exec_callback_QWizardPage_SizeHint(self QWizardPage, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QWizardPage) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QWizardPage_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWizardPage) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MinimumSizeHint
func miqt_exec_callback_QWizardPage_MinimumSizeHint(self QWizardPage, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QWizardPage) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QWizardPage_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QWizardPage) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_HeightForWidth
func miqt_exec_callback_QWizardPage_HeightForWidth(self QWizardPage, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QWizardPage_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_HasHeightForWidth
func miqt_exec_callback_QWizardPage_HasHeightForWidth(self QWizardPage, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QWizardPage_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_PaintEngine
func miqt_exec_callback_QWizardPage_PaintEngine(self QWizardPage, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QWizardPage) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QWizardPage_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QWizardPage) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_Event
func miqt_exec_callback_QWizardPage_Event(self QWizardPage, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QWizardPage_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MousePressEvent
func miqt_exec_callback_QWizardPage_MousePressEvent(self QWizardPage, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QWizardPage_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MouseReleaseEvent
func miqt_exec_callback_QWizardPage_MouseReleaseEvent(self QWizardPage, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QWizardPage_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MouseDoubleClickEvent
func miqt_exec_callback_QWizardPage_MouseDoubleClickEvent(self QWizardPage, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QWizardPage_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MouseMoveEvent
func miqt_exec_callback_QWizardPage_MouseMoveEvent(self QWizardPage, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QWizardPage_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_WheelEvent
func miqt_exec_callback_QWizardPage_WheelEvent(self QWizardPage, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QWizardPage_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_KeyPressEvent
func miqt_exec_callback_QWizardPage_KeyPressEvent(self QWizardPage, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QWizardPage_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_KeyReleaseEvent
func miqt_exec_callback_QWizardPage_KeyReleaseEvent(self QWizardPage, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QWizardPage_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_FocusInEvent
func miqt_exec_callback_QWizardPage_FocusInEvent(self QWizardPage, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QWizardPage_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_FocusOutEvent
func miqt_exec_callback_QWizardPage_FocusOutEvent(self QWizardPage, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QWizardPage_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_EnterEvent
func miqt_exec_callback_QWizardPage_EnterEvent(self QWizardPage, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_LeaveEvent(event *QEvent) {

	QWizardPage_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_LeaveEvent
func miqt_exec_callback_QWizardPage_LeaveEvent(self QWizardPage, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QWizardPage_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_PaintEvent
func miqt_exec_callback_QWizardPage_PaintEvent(self QWizardPage, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QWizardPage_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MoveEvent
func miqt_exec_callback_QWizardPage_MoveEvent(self QWizardPage, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QWizardPage_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_ResizeEvent
func miqt_exec_callback_QWizardPage_ResizeEvent(self QWizardPage, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QWizardPage_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_CloseEvent
func miqt_exec_callback_QWizardPage_CloseEvent(self QWizardPage, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QWizardPage_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_ContextMenuEvent
func miqt_exec_callback_QWizardPage_ContextMenuEvent(self QWizardPage, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QWizardPage_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_TabletEvent
func miqt_exec_callback_QWizardPage_TabletEvent(self QWizardPage, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_ActionEvent(event *QActionEvent) {

	QWizardPage_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_ActionEvent
func miqt_exec_callback_QWizardPage_ActionEvent(self QWizardPage, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QWizardPage_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_DragEnterEvent
func miqt_exec_callback_QWizardPage_DragEnterEvent(self QWizardPage, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QWizardPage_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_DragMoveEvent
func miqt_exec_callback_QWizardPage_DragMoveEvent(self QWizardPage, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QWizardPage_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_DragLeaveEvent
func miqt_exec_callback_QWizardPage_DragLeaveEvent(self QWizardPage, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_DropEvent(event *QDropEvent) {

	QWizardPage_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_DropEvent
func miqt_exec_callback_QWizardPage_DropEvent(self QWizardPage, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_ShowEvent(event *QShowEvent) {

	QWizardPage_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_ShowEvent
func miqt_exec_callback_QWizardPage_ShowEvent(self QWizardPage, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_HideEvent(event *QHideEvent) {

	QWizardPage_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QWizardPage) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_HideEvent
func miqt_exec_callback_QWizardPage_HideEvent(self QWizardPage, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QWizardPage{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QWizardPage_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QWizardPage) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_NativeEvent
func miqt_exec_callback_QWizardPage_NativeEvent(self QWizardPage, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QWizardPage_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWizardPage) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_ChangeEvent
func miqt_exec_callback_QWizardPage_ChangeEvent(self QWizardPage, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QWizardPage{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QWizardPage_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QWizardPage) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_Metric
func miqt_exec_callback_QWizardPage_Metric(self QWizardPage, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QWizardPage) callVirtualBase_InitPainter(painter *QPainter) {

	QWizardPage_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QWizardPage) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_InitPainter
func miqt_exec_callback_QWizardPage_InitPainter(self QWizardPage, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QWizardPage{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QWizardPage) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QWizardPage_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QWizardPage) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_Redirected
func miqt_exec_callback_QWizardPage_Redirected(self QWizardPage, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QWizardPage) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QWizardPage_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QWizardPage) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_SharedPainter
func miqt_exec_callback_QWizardPage_SharedPainter(self QWizardPage, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QWizardPage) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QWizardPage_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QWizardPage) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_InputMethodEvent
func miqt_exec_callback_QWizardPage_InputMethodEvent(self QWizardPage, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QWizardPage{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QWizardPage) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QWizardPage_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QWizardPage) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_InputMethodQuery
func miqt_exec_callback_QWizardPage_InputMethodQuery(self QWizardPage, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QWizardPage) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QWizardPage_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QWizardPage) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_FocusNextPrevChild
func miqt_exec_callback_QWizardPage_FocusNextPrevChild(self QWizardPage, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
