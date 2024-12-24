package qt

import (
	"unsafe"
)

type QTextObject struct {
	h          uintptr
	isSubclass bool
}

func (this *QTextObject) MetaObject() *QMetaObject {
	return newQMetaObject(QTextObject_MetaObject(this.h))
}

func (this *QTextObject) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextObject_Metacast(this.h, param1_Cstring))
}

func QTextObject_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextObject_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextObject) Format() *QTextFormat {
	_goptr := newQTextFormat(QTextObject_Format(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextObject) FormatIndex() int {
	return (int)(QTextObject_FormatIndex(this.h))
}

func (this *QTextObject) Document() *QTextDocument {
	return newQTextDocument(QTextObject_Document(this.h))
}

func (this *QTextObject) ObjectIndex() int {
	return (int)(QTextObject_ObjectIndex(this.h))
}

func QTextObject_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextObject_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextObject_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextObject_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QTextBlockGroup struct {
	h          uintptr
	isSubclass bool
}

func (this *QTextBlockGroup) MetaObject() *QMetaObject {
	return newQMetaObject(QTextBlockGroup_MetaObject(this.h))
}

func (this *QTextBlockGroup) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextBlockGroup_Metacast(this.h, param1_Cstring))
}

func QTextBlockGroup_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextBlockGroup_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextBlockGroup_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextBlockGroup_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextBlockGroup_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextBlockGroup_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QTextFrameLayoutData struct {
	h          uintptr
	isSubclass bool
}

func (this *QTextFrameLayoutData) OperatorAssign(param1 *QTextFrameLayoutData) {
	QTextFrameLayoutData_OperatorAssign(this.h, param1.cPointer())
}

type QTextFrame struct {
	h          uintptr
	isSubclass bool
}

// NewQTextFrame constructs a new QTextFrame object.
func NewQTextFrame(doc *QTextDocument) *QTextFrame {

	ret := newQTextFrame(QTextFrame_new(doc.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextFrame) MetaObject() *QMetaObject {
	return newQMetaObject(QTextFrame_MetaObject(this.h))
}

func (this *QTextFrame) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTextFrame_Metacast(this.h, param1_Cstring))
}

func QTextFrame_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTextFrame_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextFrame) SetFrameFormat(format *QTextFrameFormat) {
	QTextFrame_SetFrameFormat(this.h, format.cPointer())
}

func (this *QTextFrame) FrameFormat() *QTextFrameFormat {
	_goptr := newQTextFrameFormat(QTextFrame_FrameFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrame) FirstCursorPosition() *QTextCursor {
	_goptr := newQTextCursor(QTextFrame_FirstCursorPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrame) LastCursorPosition() *QTextCursor {
	_goptr := newQTextCursor(QTextFrame_LastCursorPosition(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrame) FirstPosition() int {
	return (int)(QTextFrame_FirstPosition(this.h))
}

func (this *QTextFrame) LastPosition() int {
	return (int)(QTextFrame_LastPosition(this.h))
}

func (this *QTextFrame) LayoutData() *QTextFrameLayoutData {
	return newQTextFrameLayoutData(QTextFrame_LayoutData(this.h))
}

func (this *QTextFrame) SetLayoutData(data *QTextFrameLayoutData) {
	QTextFrame_SetLayoutData(this.h, data.cPointer())
}

func (this *QTextFrame) ChildFrames() []*QTextFrame {
	var _ma struct_miqt_array = QTextFrame_ChildFrames(this.h)
	_ret := make([]*QTextFrame, int(_ma.len))
	_outCast := (*[0xffff]*QTextFrame)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_ret[i] = newQTextFrame(_outCast[i])
	}
	return _ret
}

func (this *QTextFrame) ParentFrame() *QTextFrame {
	return newQTextFrame(QTextFrame_ParentFrame(this.h))
}

func (this *QTextFrame) Begin() iterator {
	xxxxxxxxx
}

func (this *QTextFrame) End() iterator {
	xxxxxxxxx
}

func QTextFrame_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextFrame_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextFrame_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTextFrame_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QTextBlockUserData struct {
	h          uintptr
	isSubclass bool
}

func (this *QTextBlockUserData) OperatorAssign(param1 *QTextBlockUserData) {
	QTextBlockUserData_OperatorAssign(this.h, param1.cPointer())
}

type QTextBlock struct {
	h          uintptr
	isSubclass bool
}

// NewQTextBlock constructs a new QTextBlock object.
func NewQTextBlock() *QTextBlock {

	ret := newQTextBlock(QTextBlock_new())
	ret.isSubclass = true
	return ret
}

// NewQTextBlock2 constructs a new QTextBlock object.
func NewQTextBlock2(o *QTextBlock) *QTextBlock {

	ret := newQTextBlock(QTextBlock_new2(o.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextBlock) OperatorAssign(o *QTextBlock) {
	QTextBlock_OperatorAssign(this.h, o.cPointer())
}

func (this *QTextBlock) IsValid() bool {
	return (bool)(QTextBlock_IsValid(this.h))
}

func (this *QTextBlock) OperatorEqual(o *QTextBlock) bool {
	return (bool)(QTextBlock_OperatorEqual(this.h, o.cPointer()))
}

func (this *QTextBlock) OperatorNotEqual(o *QTextBlock) bool {
	return (bool)(QTextBlock_OperatorNotEqual(this.h, o.cPointer()))
}

func (this *QTextBlock) OperatorLesser(o *QTextBlock) bool {
	return (bool)(QTextBlock_OperatorLesser(this.h, o.cPointer()))
}

func (this *QTextBlock) Position() int {
	return (int)(QTextBlock_Position(this.h))
}

func (this *QTextBlock) Length() int {
	return (int)(QTextBlock_Length(this.h))
}

func (this *QTextBlock) Contains(position int) bool {
	return (bool)(QTextBlock_Contains(this.h, (int)(position)))
}

func (this *QTextBlock) Layout() *QTextLayout {
	return newQTextLayout(QTextBlock_Layout(this.h))
}

func (this *QTextBlock) ClearLayout() {
	QTextBlock_ClearLayout(this.h)
}

func (this *QTextBlock) BlockFormat() *QTextBlockFormat {
	_goptr := newQTextBlockFormat(QTextBlock_BlockFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBlock) BlockFormatIndex() int {
	return (int)(QTextBlock_BlockFormatIndex(this.h))
}

func (this *QTextBlock) CharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextBlock_CharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBlock) CharFormatIndex() int {
	return (int)(QTextBlock_CharFormatIndex(this.h))
}

func (this *QTextBlock) TextDirection() LayoutDirection {
	return (LayoutDirection)(QTextBlock_TextDirection(this.h))
}

func (this *QTextBlock) Text() string {
	var _ms struct_miqt_string = QTextBlock_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextBlock) TextFormats() []QTextLayout__FormatRange {
	var _ma struct_miqt_array = QTextBlock_TextFormats(this.h)
	_ret := make([]QTextLayout__FormatRange, int(_ma.len))
	_outCast := (*[0xffff]*QTextLayout__FormatRange)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQTextLayout__FormatRange(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextBlock) Document() *QTextDocument {
	return newQTextDocument(QTextBlock_Document(this.h))
}

func (this *QTextBlock) TextList() *QTextList {
	return newQTextList(QTextBlock_TextList(this.h))
}

func (this *QTextBlock) UserData() *QTextBlockUserData {
	return newQTextBlockUserData(QTextBlock_UserData(this.h))
}

func (this *QTextBlock) SetUserData(data *QTextBlockUserData) {
	QTextBlock_SetUserData(this.h, data.cPointer())
}

func (this *QTextBlock) UserState() int {
	return (int)(QTextBlock_UserState(this.h))
}

func (this *QTextBlock) SetUserState(state int) {
	QTextBlock_SetUserState(this.h, (int)(state))
}

func (this *QTextBlock) Revision() int {
	return (int)(QTextBlock_Revision(this.h))
}

func (this *QTextBlock) SetRevision(rev int) {
	QTextBlock_SetRevision(this.h, (int)(rev))
}

func (this *QTextBlock) IsVisible() bool {
	return (bool)(QTextBlock_IsVisible(this.h))
}

func (this *QTextBlock) SetVisible(visible bool) {
	QTextBlock_SetVisible(this.h, (bool)(visible))
}

func (this *QTextBlock) BlockNumber() int {
	return (int)(QTextBlock_BlockNumber(this.h))
}

func (this *QTextBlock) FirstLineNumber() int {
	return (int)(QTextBlock_FirstLineNumber(this.h))
}

func (this *QTextBlock) SetLineCount(count int) {
	QTextBlock_SetLineCount(this.h, (int)(count))
}

func (this *QTextBlock) LineCount() int {
	return (int)(QTextBlock_LineCount(this.h))
}

func (this *QTextBlock) Begin() iterator {
	xxxxxxxxx
}

func (this *QTextBlock) End() iterator {
	xxxxxxxxx
}

func (this *QTextBlock) Next() *QTextBlock {
	_goptr := newQTextBlock(QTextBlock_Next(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBlock) Previous() *QTextBlock {
	_goptr := newQTextBlock(QTextBlock_Previous(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBlock) FragmentIndex() int {
	return (int)(QTextBlock_FragmentIndex(this.h))
}

type QTextFragment struct {
	h          uintptr
	isSubclass bool
}

// NewQTextFragment constructs a new QTextFragment object.
func NewQTextFragment() *QTextFragment {

	ret := newQTextFragment(QTextFragment_new())
	ret.isSubclass = true
	return ret
}

// NewQTextFragment2 constructs a new QTextFragment object.
func NewQTextFragment2(o *QTextFragment) *QTextFragment {

	ret := newQTextFragment(QTextFragment_new2(o.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextFragment) OperatorAssign(o *QTextFragment) {
	QTextFragment_OperatorAssign(this.h, o.cPointer())
}

func (this *QTextFragment) IsValid() bool {
	return (bool)(QTextFragment_IsValid(this.h))
}

func (this *QTextFragment) OperatorEqual(o *QTextFragment) bool {
	return (bool)(QTextFragment_OperatorEqual(this.h, o.cPointer()))
}

func (this *QTextFragment) OperatorNotEqual(o *QTextFragment) bool {
	return (bool)(QTextFragment_OperatorNotEqual(this.h, o.cPointer()))
}

func (this *QTextFragment) OperatorLesser(o *QTextFragment) bool {
	return (bool)(QTextFragment_OperatorLesser(this.h, o.cPointer()))
}

func (this *QTextFragment) Position() int {
	return (int)(QTextFragment_Position(this.h))
}

func (this *QTextFragment) Length() int {
	return (int)(QTextFragment_Length(this.h))
}

func (this *QTextFragment) Contains(position int) bool {
	return (bool)(QTextFragment_Contains(this.h, (int)(position)))
}

func (this *QTextFragment) CharFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QTextFragment_CharFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFragment) CharFormatIndex() int {
	return (int)(QTextFragment_CharFormatIndex(this.h))
}

func (this *QTextFragment) Text() string {
	var _ms struct_miqt_string = QTextFragment_Text(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextFragment) GlyphRuns() []QGlyphRun {
	var _ma struct_miqt_array = QTextFragment_GlyphRuns(this.h)
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextFragment) GlyphRuns1(from int) []QGlyphRun {
	var _ma struct_miqt_array = QTextFragment_GlyphRuns1(this.h, (int)(from))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

func (this *QTextFragment) GlyphRuns2(from int, length int) []QGlyphRun {
	var _ma struct_miqt_array = QTextFragment_GlyphRuns2(this.h, (int)(from), (int)(length))
	_ret := make([]QGlyphRun, int(_ma.len))
	_outCast := (*[0xffff]*QGlyphRun)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		_lv_goptr := newQGlyphRun(_outCast[i])
		_lv_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_ret[i] = *_lv_goptr
	}
	return _ret
}

type QTextFrame__iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQTextFrame__iterator constructs a new QTextFrame::iterator object.
func NewQTextFrame__iterator() *QTextFrame__iterator {

	ret := newQTextFrame__iterator(QTextFrame__iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQTextFrame__iterator2 constructs a new QTextFrame::iterator object.
func NewQTextFrame__iterator2(param1 *iterator) *QTextFrame__iterator {

	ret := newQTextFrame__iterator(QTextFrame__iterator_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QTextFrame__iterator) ParentFrame() *QTextFrame {
	return newQTextFrame(QTextFrame__iterator_ParentFrame(this.h))
}

func (this *QTextFrame__iterator) CurrentFrame() *QTextFrame {
	return newQTextFrame(QTextFrame__iterator_CurrentFrame(this.h))
}

func (this *QTextFrame__iterator) CurrentBlock() *QTextBlock {
	_goptr := newQTextBlock(QTextFrame__iterator_CurrentBlock(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextFrame__iterator) AtEnd() bool {
	return (bool)(QTextFrame__iterator_AtEnd(this.h))
}

func (this *QTextFrame__iterator) OperatorEqual(o *iterator) bool {
	return (bool)(QTextFrame__iterator_OperatorEqual(this.h, o))
}

func (this *QTextFrame__iterator) OperatorNotEqual(o *iterator) bool {
	return (bool)(QTextFrame__iterator_OperatorNotEqual(this.h, o))
}

func (this *QTextFrame__iterator) OperatorPlusPlus() *iterator {
	xxxxxxxxx
}

func (this *QTextFrame__iterator) OperatorPlusPlusWithInt(param1 int) iterator {
	xxxxxxxxx
}

func (this *QTextFrame__iterator) OperatorMinusMinus() *iterator {
	xxxxxxxxx
}

func (this *QTextFrame__iterator) OperatorMinusMinusWithInt(param1 int) iterator {
	xxxxxxxxx
}

type QTextBlock__iterator struct {
	h          uintptr
	isSubclass bool
}

// NewQTextBlock__iterator constructs a new QTextBlock::iterator object.
func NewQTextBlock__iterator() *QTextBlock__iterator {

	ret := newQTextBlock__iterator(QTextBlock__iterator_new())
	ret.isSubclass = true
	return ret
}

// NewQTextBlock__iterator2 constructs a new QTextBlock::iterator object.
func NewQTextBlock__iterator2(param1 *iterator) *QTextBlock__iterator {

	ret := newQTextBlock__iterator(QTextBlock__iterator_new2(param1))
	ret.isSubclass = true
	return ret
}

func (this *QTextBlock__iterator) Fragment() *QTextFragment {
	_goptr := newQTextFragment(QTextBlock__iterator_Fragment(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextBlock__iterator) AtEnd() bool {
	return (bool)(QTextBlock__iterator_AtEnd(this.h))
}

func (this *QTextBlock__iterator) OperatorEqual(o *iterator) bool {
	return (bool)(QTextBlock__iterator_OperatorEqual(this.h, o))
}

func (this *QTextBlock__iterator) OperatorNotEqual(o *iterator) bool {
	return (bool)(QTextBlock__iterator_OperatorNotEqual(this.h, o))
}

func (this *QTextBlock__iterator) OperatorPlusPlus() *iterator {
	xxxxxxxxx
}

func (this *QTextBlock__iterator) OperatorPlusPlusWithInt(param1 int) iterator {
	xxxxxxxxx
}

func (this *QTextBlock__iterator) OperatorMinusMinus() *iterator {
	xxxxxxxxx
}

func (this *QTextBlock__iterator) OperatorMinusMinusWithInt(param1 int) iterator {
	xxxxxxxxx
}
