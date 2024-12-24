package qt

import (
	"unsafe"
)

type QDebug__VerbosityLevel int

const (
	QDebug__MinimumVerbosity QDebug__VerbosityLevel = 0
	QDebug__DefaultVerbosity QDebug__VerbosityLevel = 2
	QDebug__MaximumVerbosity QDebug__VerbosityLevel = 7
)

type QDebug struct {
	h          uintptr
	isSubclass bool
}

// NewQDebug constructs a new QDebug object.
func NewQDebug(device *QIODevice) *QDebug {

	ret := newQDebug(QDebug_new(device.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDebug2 constructs a new QDebug object.
func NewQDebug2(o *QDebug) *QDebug {

	ret := newQDebug(QDebug_new2(o.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDebug) OperatorAssign(other *QDebug) {
	QDebug_OperatorAssign(this.h, other.cPointer())
}

func (this *QDebug) Swap(other *QDebug) {
	QDebug_Swap(this.h, other.cPointer())
}

func (this *QDebug) ResetFormat() *QDebug {
	return newQDebug(QDebug_ResetFormat(this.h))
}

func (this *QDebug) Space() *QDebug {
	return newQDebug(QDebug_Space(this.h))
}

func (this *QDebug) Nospace() *QDebug {
	return newQDebug(QDebug_Nospace(this.h))
}

func (this *QDebug) MaybeSpace() *QDebug {
	return newQDebug(QDebug_MaybeSpace(this.h))
}

func (this *QDebug) Verbosity(verbosityLevel int) *QDebug {
	return newQDebug(QDebug_Verbosity(this.h, (int)(verbosityLevel)))
}

func (this *QDebug) Verbosity2() int {
	return (int)(QDebug_Verbosity2(this.h))
}

func (this *QDebug) SetVerbosity(verbosityLevel int) {
	QDebug_SetVerbosity(this.h, (int)(verbosityLevel))
}

func (this *QDebug) AutoInsertSpaces() bool {
	return (bool)(QDebug_AutoInsertSpaces(this.h))
}

func (this *QDebug) SetAutoInsertSpaces(b bool) {
	QDebug_SetAutoInsertSpaces(this.h, (bool)(b))
}

func (this *QDebug) QuoteStrings() bool {
	return (bool)(QDebug_QuoteStrings(this.h))
}

func (this *QDebug) SetQuoteStrings(b bool) {
	QDebug_SetQuoteStrings(this.h, (bool)(b))
}

func (this *QDebug) Quote() *QDebug {
	return newQDebug(QDebug_Quote(this.h))
}

func (this *QDebug) Noquote() *QDebug {
	return newQDebug(QDebug_Noquote(this.h))
}

func (this *QDebug) MaybeQuote() *QDebug {
	return newQDebug(QDebug_MaybeQuote(this.h))
}

func (this *QDebug) OperatorShiftLeft(t QChar) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeft(this.h, t.cPointer()))
}

func (this *QDebug) OperatorShiftLeftWithBool(t bool) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithBool(this.h, (bool)(t)))
}

func (this *QDebug) OperatorShiftLeftWithChar(t int8) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithChar(this.h, (char)(t)))
}

func (this *QDebug) OperatorShiftLeftWithShort(t int16) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithShort(this.h, (int16_t)(t)))
}

func (this *QDebug) OperatorShiftLeftWithUnsignedshort(t uint16) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithUnsignedshort(this.h, (uint16_t)(t)))
}

func (this *QDebug) OperatorShiftLeftWithInt(t int) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithInt(this.h, (int)(t)))
}

func (this *QDebug) OperatorShiftLeftWithUnsignedint(t uint) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithUnsignedint(this.h, (uint)(t)))
}

func (this *QDebug) OperatorShiftLeftWithLong(t int32) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithLong(this.h, (long)(t)))
}

func (this *QDebug) OperatorShiftLeftWithUnsignedlong(t uint32) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithUnsignedlong(this.h, (ulong)(t)))
}

func (this *QDebug) OperatorShiftLeftWithQint64(t int64) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithQint64(this.h, (longlong)(t)))
}

func (this *QDebug) OperatorShiftLeftWithQuint64(t uint64) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithQuint64(this.h, (ulonglong)(t)))
}

func (this *QDebug) OperatorShiftLeftWithFloat(t float32) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithFloat(this.h, (float)(t)))
}

func (this *QDebug) OperatorShiftLeftWithDouble(t float64) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithDouble(this.h, (double)(t)))
}

func (this *QDebug) OperatorShiftLeft2(t string) *QDebug {
	t_Cstring := CString(t)
	defer free(unsafe.Pointer(t_Cstring))
	return newQDebug(QDebug_OperatorShiftLeft2(this.h, t_Cstring))
}

func (this *QDebug) OperatorShiftLeftWithQString(t string) *QDebug {
	t_ms := struct_miqt_string{}
	t_ms.data = CString(t)
	t_ms.len = size_t(len(t))
	defer free(unsafe.Pointer(t_ms.data))
	return newQDebug(QDebug_OperatorShiftLeftWithQString(this.h, t_ms))
}

func (this *QDebug) OperatorShiftLeftWithQByteArray(t []byte) *QDebug {
	t_alias := struct_miqt_string{}
	t_alias.data = (char)(unsafe.Pointer(&t[0]))
	t_alias.len = size_t(len(t))
	return newQDebug(QDebug_OperatorShiftLeftWithQByteArray(this.h, t_alias))
}

func (this *QDebug) OperatorShiftLeftWithQByteArrayView(t QByteArrayView) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithQByteArrayView(this.h, t.cPointer()))
}

func (this *QDebug) OperatorShiftLeftWithVoid(t unsafe.Pointer) *QDebug {
	return newQDebug(QDebug_OperatorShiftLeftWithVoid(this.h, t))
}

func (this *QDebug) MaybeQuote1(c int8) *QDebug {
	return newQDebug(QDebug_MaybeQuote1(this.h, (char)(c)))
}

type QDebugStateSaver struct {
	h          uintptr
	isSubclass bool
}

// NewQDebugStateSaver constructs a new QDebugStateSaver object.
func NewQDebugStateSaver(dbg *QDebug) *QDebugStateSaver {

	ret := newQDebugStateSaver(QDebugStateSaver_new(dbg.cPointer()))
	ret.isSubclass = true
	return ret
}

type QNoDebug struct {
	h          uintptr
	isSubclass bool
}

func (this *QNoDebug) Space() *QNoDebug {
	return newQNoDebug(QNoDebug_Space(this.h))
}

func (this *QNoDebug) Nospace() *QNoDebug {
	return newQNoDebug(QNoDebug_Nospace(this.h))
}

func (this *QNoDebug) MaybeSpace() *QNoDebug {
	return newQNoDebug(QNoDebug_MaybeSpace(this.h))
}

func (this *QNoDebug) Quote() *QNoDebug {
	return newQNoDebug(QNoDebug_Quote(this.h))
}

func (this *QNoDebug) Noquote() *QNoDebug {
	return newQNoDebug(QNoDebug_Noquote(this.h))
}

func (this *QNoDebug) MaybeQuote() *QNoDebug {
	return newQNoDebug(QNoDebug_MaybeQuote(this.h))
}

func (this *QNoDebug) Verbosity(param1 int) *QNoDebug {
	return newQNoDebug(QNoDebug_Verbosity(this.h, (int)(param1)))
}

func (this *QNoDebug) MaybeQuote1(param1 int8) *QNoDebug {
	return newQNoDebug(QNoDebug_MaybeQuote1(this.h, (const_char)(param1)))
}
