package qt

import (
	"unsafe"
)

type QTextBoundaryFinder__BoundaryType int

const (
	QTextBoundaryFinder__Grapheme QTextBoundaryFinder__BoundaryType = 0
	QTextBoundaryFinder__Word     QTextBoundaryFinder__BoundaryType = 1
	QTextBoundaryFinder__Sentence QTextBoundaryFinder__BoundaryType = 2
	QTextBoundaryFinder__Line     QTextBoundaryFinder__BoundaryType = 3
)

type QTextBoundaryFinder__BoundaryReason int

const (
	QTextBoundaryFinder__NotAtBoundary    QTextBoundaryFinder__BoundaryReason = 0
	QTextBoundaryFinder__BreakOpportunity QTextBoundaryFinder__BoundaryReason = 31
	QTextBoundaryFinder__StartOfItem      QTextBoundaryFinder__BoundaryReason = 32
	QTextBoundaryFinder__EndOfItem        QTextBoundaryFinder__BoundaryReason = 64
	QTextBoundaryFinder__MandatoryBreak   QTextBoundaryFinder__BoundaryReason = 128
	QTextBoundaryFinder__SoftHyphen       QTextBoundaryFinder__BoundaryReason = 256
)

type QTextBoundaryFinder struct {
	h          uintptr
	isSubclass bool
}

// NewQTextBoundaryFinder constructs a new QTextBoundaryFinder object.
func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	ret := newQTextBoundaryFinder(QTextBoundaryFinder_new())
	ret.isSubclass = true
	return ret
}

// NewQTextBoundaryFinder2 constructs a new QTextBoundaryFinder object.
func NewQTextBoundaryFinder2(other *QTextBoundaryFinder) *QTextBoundaryFinder {
	ret := newQTextBoundaryFinder(QTextBoundaryFinder_new2(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTextBoundaryFinder3 constructs a new QTextBoundaryFinder object.
func NewQTextBoundaryFinder3(typeVal BoundaryType, stringVal string) *QTextBoundaryFinder {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))

	ret := newQTextBoundaryFinder(QTextBoundaryFinder_new3(typeVal, stringVal_ms))
	ret.isSubclass = true
	return ret
}

// NewQTextBoundaryFinder4 constructs a new QTextBoundaryFinder object.
func NewQTextBoundaryFinder4(typeVal BoundaryType, chars *QChar, length int64) *QTextBoundaryFinder {
	ret := newQTextBoundaryFinder(QTextBoundaryFinder_new4(typeVal, chars.cPointer(), (ptrdiff_t)(length)))
	ret.isSubclass = true
	return ret
}

// NewQTextBoundaryFinder5 constructs a new QTextBoundaryFinder object.
func NewQTextBoundaryFinder5(typeVal BoundaryType, chars *QChar, length int64, buffer *byte) *QTextBoundaryFinder {
	ret := newQTextBoundaryFinder(QTextBoundaryFinder_new5(typeVal, chars.cPointer(), (ptrdiff_t)(length), (*uchar)(unsafe.Pointer(buffer))))
	ret.isSubclass = true
	return ret
}

// NewQTextBoundaryFinder6 constructs a new QTextBoundaryFinder object.
func NewQTextBoundaryFinder6(typeVal BoundaryType, chars *QChar, length int64, buffer *byte, bufferSize int64) *QTextBoundaryFinder {
	ret := newQTextBoundaryFinder(QTextBoundaryFinder_new6(typeVal, chars.cPointer(), (ptrdiff_t)(length), (*uchar)(unsafe.Pointer(buffer)), (ptrdiff_t)(bufferSize)))
	ret.isSubclass = true
	return ret
}

func (this *QTextBoundaryFinder) OperatorAssign(other *QTextBoundaryFinder) {
	QTextBoundaryFinder_OperatorAssign(this.h, other.cPointer())
}

func (this *QTextBoundaryFinder) IsValid() bool {
	return (bool)(QTextBoundaryFinder_IsValid(this.h))
}

func (this *QTextBoundaryFinder) Type() BoundaryType {
	xxxxxxxxx
}

func (this *QTextBoundaryFinder) String() string {
	var _ms struct_miqt_string = QTextBoundaryFinder_String(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTextBoundaryFinder) ToStart() {
	QTextBoundaryFinder_ToStart(this.h)
}

func (this *QTextBoundaryFinder) ToEnd() {
	QTextBoundaryFinder_ToEnd(this.h)
}

func (this *QTextBoundaryFinder) Position() int64 {
	return (int64)(QTextBoundaryFinder_Position(this.h))
}

func (this *QTextBoundaryFinder) SetPosition(position int64) {
	QTextBoundaryFinder_SetPosition(this.h, (ptrdiff_t)(position))
}

func (this *QTextBoundaryFinder) ToNextBoundary() int64 {
	return (int64)(QTextBoundaryFinder_ToNextBoundary(this.h))
}

func (this *QTextBoundaryFinder) ToPreviousBoundary() int64 {
	return (int64)(QTextBoundaryFinder_ToPreviousBoundary(this.h))
}

func (this *QTextBoundaryFinder) IsAtBoundary() bool {
	return (bool)(QTextBoundaryFinder_IsAtBoundary(this.h))
}

func (this *QTextBoundaryFinder) BoundaryReasons() BoundaryReasons {
	xxxxxxxxx
}
