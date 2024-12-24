package qt

import (
	"unsafe"
)

type QRegularExpression__PatternOption int

const (
	QRegularExpression__NoPatternOption             QRegularExpression__PatternOption = 0
	QRegularExpression__CaseInsensitiveOption       QRegularExpression__PatternOption = 1
	QRegularExpression__DotMatchesEverythingOption  QRegularExpression__PatternOption = 2
	QRegularExpression__MultilineOption             QRegularExpression__PatternOption = 4
	QRegularExpression__ExtendedPatternSyntaxOption QRegularExpression__PatternOption = 8
	QRegularExpression__InvertedGreedinessOption    QRegularExpression__PatternOption = 16
	QRegularExpression__DontCaptureOption           QRegularExpression__PatternOption = 32
	QRegularExpression__UseUnicodePropertiesOption  QRegularExpression__PatternOption = 64
)

type QRegularExpression__MatchType int

const (
	QRegularExpression__NormalMatch                QRegularExpression__MatchType = 0
	QRegularExpression__PartialPreferCompleteMatch QRegularExpression__MatchType = 1
	QRegularExpression__PartialPreferFirstMatch    QRegularExpression__MatchType = 2
	QRegularExpression__NoMatch                    QRegularExpression__MatchType = 3
)

type QRegularExpression__MatchOption int

const (
	QRegularExpression__NoMatchOption                     QRegularExpression__MatchOption = 0
	QRegularExpression__AnchorAtOffsetMatchOption         QRegularExpression__MatchOption = 1
	QRegularExpression__AnchoredMatchOption               QRegularExpression__MatchOption = 1
	QRegularExpression__DontCheckSubjectStringMatchOption QRegularExpression__MatchOption = 2
)

type QRegularExpression__WildcardConversionOption int

const (
	QRegularExpression__DefaultWildcardConversion    QRegularExpression__WildcardConversionOption = 0
	QRegularExpression__UnanchoredWildcardConversion QRegularExpression__WildcardConversionOption = 1
	QRegularExpression__NonPathWildcardConversion    QRegularExpression__WildcardConversionOption = 2
)

type QRegularExpression struct {
	h          uintptr
	isSubclass bool
}

// NewQRegularExpression constructs a new QRegularExpression object.
func NewQRegularExpression() *QRegularExpression {
	ret := newQRegularExpression(QRegularExpression_new())
	ret.isSubclass = true
	return ret
}

// NewQRegularExpression2 constructs a new QRegularExpression object.
func NewQRegularExpression2(pattern string) *QRegularExpression {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))

	ret := newQRegularExpression(QRegularExpression_new2(pattern_ms))
	ret.isSubclass = true
	return ret
}

// NewQRegularExpression3 constructs a new QRegularExpression object.
func NewQRegularExpression3(re *QRegularExpression) *QRegularExpression {
	ret := newQRegularExpression(QRegularExpression_new3(re.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQRegularExpression4 constructs a new QRegularExpression object.
func NewQRegularExpression4(pattern string, options PatternOptions) *QRegularExpression {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))

	ret := newQRegularExpression(QRegularExpression_new4(pattern_ms, options))
	ret.isSubclass = true
	return ret
}

func (this *QRegularExpression) PatternOptions() PatternOptions {
	xxxxxxxxx
}

func (this *QRegularExpression) SetPatternOptions(options PatternOptions) {
	QRegularExpression_SetPatternOptions(this.h, options)
}

func (this *QRegularExpression) OperatorAssign(re *QRegularExpression) {
	QRegularExpression_OperatorAssign(this.h, re.cPointer())
}

func (this *QRegularExpression) Swap(other *QRegularExpression) {
	QRegularExpression_Swap(this.h, other.cPointer())
}

func (this *QRegularExpression) Pattern() string {
	var _ms struct_miqt_string = QRegularExpression_Pattern(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpression) SetPattern(pattern string) {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	QRegularExpression_SetPattern(this.h, pattern_ms)
}

func (this *QRegularExpression) IsValid() bool {
	return (bool)(QRegularExpression_IsValid(this.h))
}

func (this *QRegularExpression) PatternErrorOffset() int64 {
	return (int64)(QRegularExpression_PatternErrorOffset(this.h))
}

func (this *QRegularExpression) ErrorString() string {
	var _ms struct_miqt_string = QRegularExpression_ErrorString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpression) CaptureCount() int {
	return (int)(QRegularExpression_CaptureCount(this.h))
}

func (this *QRegularExpression) NamedCaptureGroups() []string {
	var _ma struct_miqt_array = QRegularExpression_NamedCaptureGroups(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QRegularExpression) Match(subject string) *QRegularExpressionMatch {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatch(QRegularExpression_Match(this.h, subject_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) GlobalMatch(subject string) *QRegularExpressionMatchIterator {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatchIterator(QRegularExpression_GlobalMatch(this.h, subject_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) Optimize() {
	QRegularExpression_Optimize(this.h)
}

func QRegularExpression_Escape(str string) string {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	var _ms struct_miqt_string = QRegularExpression_Escape(str_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRegularExpression_WildcardToRegularExpression(str string) string {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	var _ms struct_miqt_string = QRegularExpression_WildcardToRegularExpression(str_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QRegularExpression_AnchoredPattern(expression string) string {
	expression_ms := struct_miqt_string{}
	expression_ms.data = CString(expression)
	expression_ms.len = size_t(len(expression))
	defer free(unsafe.Pointer(expression_ms.data))
	var _ms struct_miqt_string = QRegularExpression_AnchoredPattern(expression_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpression) Match2(subject string, offset int64) *QRegularExpressionMatch {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatch(QRegularExpression_Match2(this.h, subject_ms, (ptrdiff_t)(offset)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) Match3(subject string, offset int64, matchType MatchType) *QRegularExpressionMatch {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatch(QRegularExpression_Match3(this.h, subject_ms, (ptrdiff_t)(offset), matchType))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) Match4(subject string, offset int64, matchType MatchType, matchOptions MatchOptions) *QRegularExpressionMatch {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatch(QRegularExpression_Match4(this.h, subject_ms, (ptrdiff_t)(offset), matchType, matchOptions))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) GlobalMatch2(subject string, offset int64) *QRegularExpressionMatchIterator {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatchIterator(QRegularExpression_GlobalMatch2(this.h, subject_ms, (ptrdiff_t)(offset)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) GlobalMatch3(subject string, offset int64, matchType MatchType) *QRegularExpressionMatchIterator {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatchIterator(QRegularExpression_GlobalMatch3(this.h, subject_ms, (ptrdiff_t)(offset), matchType))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpression) GlobalMatch4(subject string, offset int64, matchType MatchType, matchOptions MatchOptions) *QRegularExpressionMatchIterator {
	subject_ms := struct_miqt_string{}
	subject_ms.data = CString(subject)
	subject_ms.len = size_t(len(subject))
	defer free(unsafe.Pointer(subject_ms.data))
	_goptr := newQRegularExpressionMatchIterator(QRegularExpression_GlobalMatch4(this.h, subject_ms, (ptrdiff_t)(offset), matchType, matchOptions))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QRegularExpression_WildcardToRegularExpression2(str string, options WildcardConversionOptions) string {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	var _ms struct_miqt_string = QRegularExpression_WildcardToRegularExpression2(str_ms, options)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QRegularExpressionMatch struct {
	h          uintptr
	isSubclass bool
}

// NewQRegularExpressionMatch constructs a new QRegularExpressionMatch object.
func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	ret := newQRegularExpressionMatch(QRegularExpressionMatch_new())
	ret.isSubclass = true
	return ret
}

// NewQRegularExpressionMatch2 constructs a new QRegularExpressionMatch object.
func NewQRegularExpressionMatch2(match *QRegularExpressionMatch) *QRegularExpressionMatch {
	ret := newQRegularExpressionMatch(QRegularExpressionMatch_new2(match.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRegularExpressionMatch) OperatorAssign(match *QRegularExpressionMatch) {
	QRegularExpressionMatch_OperatorAssign(this.h, match.cPointer())
}

func (this *QRegularExpressionMatch) Swap(other *QRegularExpressionMatch) {
	QRegularExpressionMatch_Swap(this.h, other.cPointer())
}

func (this *QRegularExpressionMatch) RegularExpression() *QRegularExpression {
	_goptr := newQRegularExpression(QRegularExpressionMatch_RegularExpression(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpressionMatch) MatchType() QRegularExpression__MatchType {
	return (QRegularExpression__MatchType)(QRegularExpressionMatch_MatchType(this.h))
}

func (this *QRegularExpressionMatch) MatchOptions() MatchOption {
	return (MatchOption)(QRegularExpressionMatch_MatchOptions(this.h))
}

func (this *QRegularExpressionMatch) HasMatch() bool {
	return (bool)(QRegularExpressionMatch_HasMatch(this.h))
}

func (this *QRegularExpressionMatch) HasPartialMatch() bool {
	return (bool)(QRegularExpressionMatch_HasPartialMatch(this.h))
}

func (this *QRegularExpressionMatch) IsValid() bool {
	return (bool)(QRegularExpressionMatch_IsValid(this.h))
}

func (this *QRegularExpressionMatch) LastCapturedIndex() int {
	return (int)(QRegularExpressionMatch_LastCapturedIndex(this.h))
}

func (this *QRegularExpressionMatch) HasCaptured(name QAnyStringView) bool {
	return (bool)(QRegularExpressionMatch_HasCaptured(this.h, name.cPointer()))
}

func (this *QRegularExpressionMatch) HasCapturedWithNth(nth int) bool {
	return (bool)(QRegularExpressionMatch_HasCapturedWithNth(this.h, (int)(nth)))
}

func (this *QRegularExpressionMatch) Captured() string {
	var _ms struct_miqt_string = QRegularExpressionMatch_Captured(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpressionMatch) CapturedWithName(name QAnyStringView) string {
	var _ms struct_miqt_string = QRegularExpressionMatch_CapturedWithName(this.h, name.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpressionMatch) CapturedTexts() []string {
	var _ma struct_miqt_array = QRegularExpressionMatch_CapturedTexts(this.h)
	_ret := make([]string, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_ms struct_miqt_string = _outCast[i]
		_lv_ret := GoStringN(_lv_ms.data, int(int64(_lv_ms.len)))
		free(unsafe.Pointer(_lv_ms.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QRegularExpressionMatch) CapturedStart() int64 {
	return (int64)(QRegularExpressionMatch_CapturedStart(this.h))
}

func (this *QRegularExpressionMatch) CapturedLength() int64 {
	return (int64)(QRegularExpressionMatch_CapturedLength(this.h))
}

func (this *QRegularExpressionMatch) CapturedEnd() int64 {
	return (int64)(QRegularExpressionMatch_CapturedEnd(this.h))
}

func (this *QRegularExpressionMatch) CapturedStartWithName(name QAnyStringView) int64 {
	return (int64)(QRegularExpressionMatch_CapturedStartWithName(this.h, name.cPointer()))
}

func (this *QRegularExpressionMatch) CapturedLengthWithName(name QAnyStringView) int64 {
	return (int64)(QRegularExpressionMatch_CapturedLengthWithName(this.h, name.cPointer()))
}

func (this *QRegularExpressionMatch) CapturedEndWithName(name QAnyStringView) int64 {
	return (int64)(QRegularExpressionMatch_CapturedEndWithName(this.h, name.cPointer()))
}

func (this *QRegularExpressionMatch) Captured1(nth int) string {
	var _ms struct_miqt_string = QRegularExpressionMatch_Captured1(this.h, (int)(nth))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QRegularExpressionMatch) CapturedStart1(nth int) int64 {
	return (int64)(QRegularExpressionMatch_CapturedStart1(this.h, (int)(nth)))
}

func (this *QRegularExpressionMatch) CapturedLength1(nth int) int64 {
	return (int64)(QRegularExpressionMatch_CapturedLength1(this.h, (int)(nth)))
}

func (this *QRegularExpressionMatch) CapturedEnd1(nth int) int64 {
	return (int64)(QRegularExpressionMatch_CapturedEnd1(this.h, (int)(nth)))
}

type QRegularExpressionMatchIterator struct {
	h          uintptr
	isSubclass bool
}

// NewQRegularExpressionMatchIterator constructs a new QRegularExpressionMatchIterator object.
func NewQRegularExpressionMatchIterator() *QRegularExpressionMatchIterator {
	ret := newQRegularExpressionMatchIterator(QRegularExpressionMatchIterator_new())
	ret.isSubclass = true
	return ret
}

// NewQRegularExpressionMatchIterator2 constructs a new QRegularExpressionMatchIterator object.
func NewQRegularExpressionMatchIterator2(iterator *QRegularExpressionMatchIterator) *QRegularExpressionMatchIterator {
	ret := newQRegularExpressionMatchIterator(QRegularExpressionMatchIterator_new2(iterator.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QRegularExpressionMatchIterator) OperatorAssign(iterator *QRegularExpressionMatchIterator) {
	QRegularExpressionMatchIterator_OperatorAssign(this.h, iterator.cPointer())
}

func (this *QRegularExpressionMatchIterator) Swap(other *QRegularExpressionMatchIterator) {
	QRegularExpressionMatchIterator_Swap(this.h, other.cPointer())
}

func (this *QRegularExpressionMatchIterator) IsValid() bool {
	return (bool)(QRegularExpressionMatchIterator_IsValid(this.h))
}

func (this *QRegularExpressionMatchIterator) HasNext() bool {
	return (bool)(QRegularExpressionMatchIterator_HasNext(this.h))
}

func (this *QRegularExpressionMatchIterator) Next() *QRegularExpressionMatch {
	_goptr := newQRegularExpressionMatch(QRegularExpressionMatchIterator_Next(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpressionMatchIterator) PeekNext() *QRegularExpressionMatch {
	_goptr := newQRegularExpressionMatch(QRegularExpressionMatchIterator_PeekNext(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpressionMatchIterator) RegularExpression() *QRegularExpression {
	_goptr := newQRegularExpression(QRegularExpressionMatchIterator_RegularExpression(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QRegularExpressionMatchIterator) MatchType() QRegularExpression__MatchType {
	return (QRegularExpression__MatchType)(QRegularExpressionMatchIterator_MatchType(this.h))
}

func (this *QRegularExpressionMatchIterator) MatchOptions() MatchOption {
	return (MatchOption)(QRegularExpressionMatchIterator_MatchOptions(this.h))
}
