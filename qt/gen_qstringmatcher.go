package qt

import (
	"unsafe"
)

type QStringMatcher struct {
	h          uintptr
	isSubclass bool
}

// NewQStringMatcher constructs a new QStringMatcher object.
func NewQStringMatcher() *QStringMatcher {
	g := newQStringMatcher(QStringMatcher_new())
	g.isSubclass = true
	return g
}

// NewQStringMatcher2 constructs a new QStringMatcher object.
func NewQStringMatcher2(pattern string) *QStringMatcher {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	g := newQStringMatcher(QStringMatcher_new2(pattern_ms))
	g.isSubclass = true
	return g
}

// NewQStringMatcher3 constructs a new QStringMatcher object.
func NewQStringMatcher3(uc *QChar, lenVal int64) *QStringMatcher {
	g := newQStringMatcher(QStringMatcher_new3(uc.cPointer(), (ptrdiff_t)(lenVal)))
	g.isSubclass = true
	return g
}

// NewQStringMatcher4 constructs a new QStringMatcher object.
func NewQStringMatcher4(other *QStringMatcher) *QStringMatcher {
	g := newQStringMatcher(QStringMatcher_new4(other.cPointer()))
	g.isSubclass = true
	return g
}

// NewQStringMatcher5 constructs a new QStringMatcher object.
func NewQStringMatcher5(pattern string, cs CaseSensitivity) *QStringMatcher {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	g := newQStringMatcher(QStringMatcher_new5(pattern_ms, (int)(cs)))
	g.isSubclass = true
	return g
}

// NewQStringMatcher6 constructs a new QStringMatcher object.
func NewQStringMatcher6(uc *QChar, lenVal int64, cs CaseSensitivity) *QStringMatcher {
	g := newQStringMatcher(QStringMatcher_new6(uc.cPointer(), (ptrdiff_t)(lenVal), (int)(cs)))
	g.isSubclass = true
	return g
}

func (this *QStringMatcher) OperatorAssign(other *QStringMatcher) {
	QStringMatcher_OperatorAssign(this.h, other.cPointer())
}

func (this *QStringMatcher) SetPattern(pattern string) {
	pattern_ms := struct_miqt_string{}
	pattern_ms.data = CString(pattern)
	pattern_ms.len = size_t(len(pattern))
	defer free(unsafe.Pointer(pattern_ms.data))
	QStringMatcher_SetPattern(this.h, pattern_ms)
}

func (this *QStringMatcher) SetCaseSensitivity(cs CaseSensitivity) {
	QStringMatcher_SetCaseSensitivity(this.h, (int)(cs))
}

func (this *QStringMatcher) IndexIn(str string) int64 {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	return (int64)(QStringMatcher_IndexIn(this.h, str_ms))
}

func (this *QStringMatcher) IndexIn2(str *QChar, length int64) int64 {
	return (int64)(QStringMatcher_IndexIn2(this.h, str.cPointer(), (ptrdiff_t)(length)))
}

func (this *QStringMatcher) Pattern() string {
	var _ms struct_miqt_string = QStringMatcher_Pattern(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QStringMatcher) CaseSensitivity() CaseSensitivity {
	return (CaseSensitivity)(QStringMatcher_CaseSensitivity(this.h))
}

func (this *QStringMatcher) IndexIn22(str string, from int64) int64 {
	str_ms := struct_miqt_string{}
	str_ms.data = CString(str)
	str_ms.len = size_t(len(str))
	defer free(unsafe.Pointer(str_ms.data))
	return (int64)(QStringMatcher_IndexIn22(this.h, str_ms, (ptrdiff_t)(from)))
}

func (this *QStringMatcher) IndexIn3(str *QChar, length int64, from int64) int64 {
	return (int64)(QStringMatcher_IndexIn3(this.h, str.cPointer(), (ptrdiff_t)(length), (ptrdiff_t)(from)))
}
