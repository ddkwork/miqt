package qt

import (
	"unsafe"
)

type QStringConverterBase__Flag int

const (
	QStringConverterBase__Default              QStringConverterBase__Flag = 0
	QStringConverterBase__Stateless            QStringConverterBase__Flag = 1
	QStringConverterBase__ConvertInvalidToNull QStringConverterBase__Flag = 2
	QStringConverterBase__WriteBom             QStringConverterBase__Flag = 4
	QStringConverterBase__ConvertInitialBom    QStringConverterBase__Flag = 8
	QStringConverterBase__UsesIcu              QStringConverterBase__Flag = 16
)

type QStringConverter__Encoding int

const (
	QStringConverter__Utf8         QStringConverter__Encoding = 0
	QStringConverter__Utf16        QStringConverter__Encoding = 1
	QStringConverter__Utf16LE      QStringConverter__Encoding = 2
	QStringConverter__Utf16BE      QStringConverter__Encoding = 3
	QStringConverter__Utf32        QStringConverter__Encoding = 4
	QStringConverter__Utf32LE      QStringConverter__Encoding = 5
	QStringConverter__Utf32BE      QStringConverter__Encoding = 6
	QStringConverter__Latin1       QStringConverter__Encoding = 7
	QStringConverter__System       QStringConverter__Encoding = 8
	QStringConverter__LastEncoding QStringConverter__Encoding = 8
)

type QStringConverterBase struct {
	h          uintptr
	isSubclass bool
}

// NewQStringConverterBase constructs a new QStringConverterBase object.
func NewQStringConverterBase(param1 *QStringConverterBase) *QStringConverterBase {

	ret := newQStringConverterBase(QStringConverterBase_new(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQStringConverterBase2 constructs a new QStringConverterBase object.
func NewQStringConverterBase2() *QStringConverterBase {

	ret := newQStringConverterBase(QStringConverterBase_new2())
	ret.isSubclass = true
	return ret
}

type QStringConverter struct {
	h          uintptr
	isSubclass bool
}

func (this *QStringConverter) IsValid() bool {
	return (bool)(QStringConverter_IsValid(this.h))
}

func (this *QStringConverter) ResetState() {
	QStringConverter_ResetState(this.h)
}

func (this *QStringConverter) HasError() bool {
	return (bool)(QStringConverter_HasError(this.h))
}

func (this *QStringConverter) Name() string {
	_ret := QStringConverter_Name(this.h)
	return GoString(_ret)
}

func QStringConverter_NameForEncoding(e Encoding) string {
	_ret := QStringConverter_NameForEncoding(e)
	return GoString(_ret)
}

func QStringConverter_AvailableCodecs() []string {
	var _ma struct_miqt_array = QStringConverter_AvailableCodecs()
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

type QStringConverterBase__State struct {
	h          uintptr
	isSubclass bool
}

// NewQStringConverterBase__State constructs a new QStringConverterBase::State object.
func NewQStringConverterBase__State() *QStringConverterBase__State {

	ret := newQStringConverterBase__State(QStringConverterBase__State_new())
	ret.isSubclass = true
	return ret
}

// NewQStringConverterBase__State2 constructs a new QStringConverterBase::State object.
func NewQStringConverterBase__State2(f Flags) *QStringConverterBase__State {

	ret := newQStringConverterBase__State(QStringConverterBase__State_new2(f))
	ret.isSubclass = true
	return ret
}

func (this *QStringConverterBase__State) Clear() {
	QStringConverterBase__State_Clear(this.h)
}

func (this *QStringConverterBase__State) Reset() {
	QStringConverterBase__State_Reset(this.h)
}
