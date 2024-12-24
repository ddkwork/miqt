package qt

import (
	"unsafe"
)

type QTextStream struct {
	h          uintptr
	isSubclass bool
}

type QTextDocumentFragment struct {
	h          uintptr
	isSubclass bool
}

// NewQTextDocumentFragment constructs a new QTextDocumentFragment object.
func NewQTextDocumentFragment() *QTextDocumentFragment {
	ret := newQTextDocumentFragment(QTextDocumentFragment_new())
	ret.isSubclass = true
	return ret
}

// NewQTextDocumentFragment2 constructs a new QTextDocumentFragment object.
func NewQTextDocumentFragment2(document *QTextDocument) *QTextDocumentFragment {
	ret := newQTextDocumentFragment(QTextDocumentFragment_new2(document.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextDocumentFragment3 constructs a new QTextDocumentFragment object.
func NewQTextDocumentFragment3(rangeVal *QTextCursor) *QTextDocumentFragment {
	ret := newQTextDocumentFragment(QTextDocumentFragment_new3(rangeVal.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextDocumentFragment4 constructs a new QTextDocumentFragment object.
func NewQTextDocumentFragment4(rhs *QTextDocumentFragment) *QTextDocumentFragment {
	ret := newQTextDocumentFragment(QTextDocumentFragment_new4(rhs.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTextDocumentFragment) OperatorAssign(rhs *QTextDocumentFragment) {
	QTextDocumentFragment_OperatorAssign(this.h, rhs.cPointer())
}

func (this *QTextDocumentFragment) IsEmpty() bool {
	return (bool)(QTextDocumentFragment_IsEmpty(this.h))
}

func (this *QTextDocumentFragment) ToPlainText() string {
	var _ms struct_miqt_string = QTextDocumentFragment_ToPlainText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocumentFragment) ToRawText() string {
	var _ms struct_miqt_string = QTextDocumentFragment_ToRawText(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocumentFragment) ToHtml() string {
	var _ms struct_miqt_string = QTextDocumentFragment_ToHtml(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextDocumentFragment) ToMarkdown() string {
	var _ms struct_miqt_string = QTextDocumentFragment_ToMarkdown(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextDocumentFragment_FromPlainText(plainText string) *QTextDocumentFragment {
	plainText_ms := struct_miqt_string{}
	plainText_ms.data = CString(plainText)
	plainText_ms.len = size_t(len(plainText))
	defer free(unsafe.Pointer(plainText_ms.data))
	_goptr := newQTextDocumentFragment(QTextDocumentFragment_FromPlainText(plainText_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTextDocumentFragment_FromHtml(html string) *QTextDocumentFragment {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	_goptr := newQTextDocumentFragment(QTextDocumentFragment_FromHtml(html_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTextDocumentFragment_FromMarkdown(markdown string) *QTextDocumentFragment {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	_goptr := newQTextDocumentFragment(QTextDocumentFragment_FromMarkdown(markdown_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTextDocumentFragment) ToMarkdown1(features MarkdownFeature) string {
	var _ms struct_miqt_string = QTextDocumentFragment_ToMarkdown1(this.h, (int)(features))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTextDocumentFragment_FromHtml2(html string, resourceProvider *QTextDocument) *QTextDocumentFragment {
	html_ms := struct_miqt_string{}
	html_ms.data = CString(html)
	html_ms.len = size_t(len(html))
	defer free(unsafe.Pointer(html_ms.data))
	_goptr := newQTextDocumentFragment(QTextDocumentFragment_FromHtml2(html_ms, resourceProvider.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTextDocumentFragment_FromMarkdown2(markdown string, features MarkdownFeature) *QTextDocumentFragment {
	markdown_ms := struct_miqt_string{}
	markdown_ms.data = CString(markdown)
	markdown_ms.len = size_t(len(markdown))
	defer free(unsafe.Pointer(markdown_ms.data))
	_goptr := newQTextDocumentFragment(QTextDocumentFragment_FromMarkdown2(markdown_ms, (int)(features)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
