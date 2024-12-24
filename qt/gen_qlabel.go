package qt

import (
	"unsafe"
)

type QLabel struct {
	h          uintptr
	isSubclass bool
}

// NewQLabel constructs a new QLabel object.
func NewQLabel(parent *QWidget) *QLabel {
	ret := newQLabel(QLabel_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLabel2 constructs a new QLabel object.
func NewQLabel2() *QLabel {
	ret := newQLabel(QLabel_new2())
	ret.isSubclass = true
	return ret
}

// NewQLabel3 constructs a new QLabel object.
func NewQLabel3(text string) *QLabel {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQLabel(QLabel_new3(text_ms))
	ret.isSubclass = true
	return ret
}

// NewQLabel4 constructs a new QLabel object.
func NewQLabel4(parent *QWidget, f WindowType) *QLabel {
	ret := newQLabel(QLabel_new4(parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

// NewQLabel5 constructs a new QLabel object.
func NewQLabel5(text string, parent *QWidget) *QLabel {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQLabel(QLabel_new5(text_ms, parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQLabel6 constructs a new QLabel object.
func NewQLabel6(text string, parent *QWidget, f WindowType) *QLabel {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	ret := newQLabel(QLabel_new6(text_ms, parent.cPointer(), (int)(f)))
	ret.isSubclass = true
	return ret
}

func (this *QLabel) MetaObject() *QMetaObject {
	return newQMetaObject(QLabel_MetaObject(this.h))
}

func (this *QLabel) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QLabel_Metacast(this.h, param1_Cstring))
}

func QLabel_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QLabel_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLabel) Text() string {
	var _ms struct_miqt_string = QLabel_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLabel) Pixmap(param1 ReturnByValueConstant) *QPixmap {
	_goptr := newQPixmap(QLabel_Pixmap(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLabel) Pixmap2() *QPixmap {
	_goptr := newQPixmap(QLabel_Pixmap2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLabel) Picture(param1 ReturnByValueConstant) *QPicture {
	_goptr := newQPicture(QLabel_Picture(this.h, (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLabel) Picture2() *QPicture {
	_goptr := newQPicture(QLabel_Picture2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLabel) Movie() *QMovie {
	return newQMovie(QLabel_Movie(this.h))
}

func (this *QLabel) TextFormat() TextFormat {
	return (TextFormat)(QLabel_TextFormat(this.h))
}

func (this *QLabel) SetTextFormat(textFormat TextFormat) {
	QLabel_SetTextFormat(this.h, (int)(textFormat))
}

func (this *QLabel) Alignment() AlignmentFlag {
	return (AlignmentFlag)(QLabel_Alignment(this.h))
}

func (this *QLabel) SetAlignment(alignment AlignmentFlag) {
	QLabel_SetAlignment(this.h, (int)(alignment))
}

func (this *QLabel) SetWordWrap(on bool) {
	QLabel_SetWordWrap(this.h, (bool)(on))
}

func (this *QLabel) WordWrap() bool {
	return (bool)(QLabel_WordWrap(this.h))
}

func (this *QLabel) Indent() int {
	return (int)(QLabel_Indent(this.h))
}

func (this *QLabel) SetIndent(indent int) {
	QLabel_SetIndent(this.h, (int)(indent))
}

func (this *QLabel) Margin() int {
	return (int)(QLabel_Margin(this.h))
}

func (this *QLabel) SetMargin(margin int) {
	QLabel_SetMargin(this.h, (int)(margin))
}

func (this *QLabel) HasScaledContents() bool {
	return (bool)(QLabel_HasScaledContents(this.h))
}

func (this *QLabel) SetScaledContents(scaledContents bool) {
	QLabel_SetScaledContents(this.h, (bool)(scaledContents))
}

func (this *QLabel) SizeHint() *QSize {
	_goptr := newQSize(QLabel_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLabel) MinimumSizeHint() *QSize {
	_goptr := newQSize(QLabel_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QLabel) SetBuddy(buddy *QWidget) {
	QLabel_SetBuddy(this.h, buddy.cPointer())
}

func (this *QLabel) Buddy() *QWidget {
	return newQWidget(QLabel_Buddy(this.h))
}

func (this *QLabel) HeightForWidth(param1 int) int {
	return (int)(QLabel_HeightForWidth(this.h, (int)(param1)))
}

func (this *QLabel) OpenExternalLinks() bool {
	return (bool)(QLabel_OpenExternalLinks(this.h))
}

func (this *QLabel) SetOpenExternalLinks(open bool) {
	QLabel_SetOpenExternalLinks(this.h, (bool)(open))
}

func (this *QLabel) SetTextInteractionFlags(flags TextInteractionFlag) {
	QLabel_SetTextInteractionFlags(this.h, (int)(flags))
}

func (this *QLabel) TextInteractionFlags() TextInteractionFlag {
	return (TextInteractionFlag)(QLabel_TextInteractionFlags(this.h))
}

func (this *QLabel) SetSelection(param1 int, param2 int) {
	QLabel_SetSelection(this.h, (int)(param1), (int)(param2))
}

func (this *QLabel) HasSelectedText() bool {
	return (bool)(QLabel_HasSelectedText(this.h))
}

func (this *QLabel) SelectedText() string {
	var _ms struct_miqt_string = QLabel_SelectedText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLabel) SelectionStart() int {
	return (int)(QLabel_SelectionStart(this.h))
}

func (this *QLabel) SetText(text string) {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))
	QLabel_SetText(this.h, text_ms)
}

func (this *QLabel) SetPixmap(pixmap *QPixmap) {
	QLabel_SetPixmap(this.h, pixmap.cPointer())
}

func (this *QLabel) SetPicture(picture *QPicture) {
	QLabel_SetPicture(this.h, picture.cPointer())
}

func (this *QLabel) SetMovie(movie *QMovie) {
	QLabel_SetMovie(this.h, movie.cPointer())
}

func (this *QLabel) SetNum(num int) {
	QLabel_SetNum(this.h, (int)(num))
}

func (this *QLabel) SetNumWithNum(num float64) {
	QLabel_SetNumWithNum(this.h, (double)(num))
}

func (this *QLabel) Clear() {
	QLabel_Clear(this.h)
}

func (this *QLabel) LinkActivated(link string) {
	link_ms := struct_miqt_string{}
	link_ms.data = CString(link)
	link_ms.len = size_t(len(link))
	defer free(unsafe.Pointer(link_ms.data))
	QLabel_LinkActivated(this.h, link_ms)
}

func (this *QLabel) OnLinkActivated(slot func(link string)) {
	QLabel_connect_LinkActivated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLabel_LinkActivated
func miqt_exec_callback_QLabel_LinkActivated(cb intptr_t, link struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(link string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var link_ms struct_miqt_string = link
	link_ret := GoStringN(link_ms.data, int(int64(link_ms.len)))
	free(unsafe.Pointer(link_ms.data))
	slotval1 := link_ret

	gofunc(slotval1)
}

func (this *QLabel) LinkHovered(link string) {
	link_ms := struct_miqt_string{}
	link_ms.data = CString(link)
	link_ms.len = size_t(len(link))
	defer free(unsafe.Pointer(link_ms.data))
	QLabel_LinkHovered(this.h, link_ms)
}

func (this *QLabel) OnLinkHovered(slot func(link string)) {
	QLabel_connect_LinkHovered(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLabel_LinkHovered
func miqt_exec_callback_QLabel_LinkHovered(cb intptr_t, link struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(link string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var link_ms struct_miqt_string = link
	link_ret := GoStringN(link_ms.data, int(int64(link_ms.len)))
	free(unsafe.Pointer(link_ms.data))
	slotval1 := link_ret

	gofunc(slotval1)
}

func QLabel_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLabel_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QLabel_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QLabel_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QLabel) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QLabel_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QLabel) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLabel_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLabel_MetaObject
func miqt_exec_callback_QLabel_MetaObject(self QLabel, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QLabel{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QLabel) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QLabel_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QLabel) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QLabel_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QLabel_Metacast
func miqt_exec_callback_QLabel_Metacast(self QLabel, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QLabel{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
