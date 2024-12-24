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

func (this *QWizard) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QWizard_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QWizard) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_MetaObject
func miqt_exec_callback_QWizard_MetaObject(self QWizard, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QWizard) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWizard_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWizard) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizard_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizard_Metacast
func miqt_exec_callback_QWizard_Metacast(self QWizard, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QWizard{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
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

func (this *QWizardPage) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QWizardPage_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QWizardPage) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_MetaObject
func miqt_exec_callback_QWizardPage_MetaObject(self QWizardPage, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QWizardPage) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QWizardPage_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QWizardPage) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QWizardPage_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QWizardPage_Metacast
func miqt_exec_callback_QWizardPage_Metacast(self QWizardPage, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QWizardPage{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
