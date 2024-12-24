package qt

import (
	"unsafe"
)

type QTimeZone__Initialization int

const (
	QTimeZone__LocalTime QTimeZone__Initialization = 0
	QTimeZone__UTC       QTimeZone__Initialization = 1
)

type QTimeZone__TimeType int

const (
	QTimeZone__StandardTime QTimeZone__TimeType = 0
	QTimeZone__DaylightTime QTimeZone__TimeType = 1
	QTimeZone__GenericTime  QTimeZone__TimeType = 2
)

type QTimeZone__NameType int

const (
	QTimeZone__DefaultName QTimeZone__NameType = 0
	QTimeZone__LongName    QTimeZone__NameType = 1
	QTimeZone__ShortName   QTimeZone__NameType = 2
	QTimeZone__OffsetName  QTimeZone__NameType = 3
)

type QTimeZone struct {
	h          uintptr
	isSubclass bool
}

// NewQTimeZone constructs a new QTimeZone object.
func NewQTimeZone() *QTimeZone {
	ret := newQTimeZone(QTimeZone_new())
	ret.isSubclass = true
	return ret
}

// NewQTimeZone2 constructs a new QTimeZone object.
func NewQTimeZone2(spec Initialization) *QTimeZone {
	ret := newQTimeZone(QTimeZone_new2(spec))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone3 constructs a new QTimeZone object.
func NewQTimeZone3(offsetSeconds int) *QTimeZone {
	ret := newQTimeZone(QTimeZone_new3((int)(offsetSeconds)))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone4 constructs a new QTimeZone object.
func NewQTimeZone4(ianaId []byte) *QTimeZone {
	ianaId_alias := struct_miqt_string{}
	ianaId_alias.data = (char)(unsafe.Pointer(&ianaId[0]))
	ianaId_alias.len = size_t(len(ianaId))

	ret := newQTimeZone(QTimeZone_new4(ianaId_alias))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone5 constructs a new QTimeZone object.
func NewQTimeZone5(zoneId []byte, offsetSeconds int, name string, abbreviation string) *QTimeZone {
	zoneId_alias := struct_miqt_string{}
	zoneId_alias.data = (char)(unsafe.Pointer(&zoneId[0]))
	zoneId_alias.len = size_t(len(zoneId))
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	abbreviation_ms := struct_miqt_string{}
	abbreviation_ms.data = CString(abbreviation)
	abbreviation_ms.len = size_t(len(abbreviation))
	defer free(unsafe.Pointer(abbreviation_ms.data))

	ret := newQTimeZone(QTimeZone_new5(zoneId_alias, (int)(offsetSeconds), name_ms, abbreviation_ms))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone6 constructs a new QTimeZone object.
func NewQTimeZone6(other *QTimeZone) *QTimeZone {
	ret := newQTimeZone(QTimeZone_new6(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone7 constructs a new QTimeZone object.
func NewQTimeZone7(zoneId []byte, offsetSeconds int, name string, abbreviation string, territory Country) *QTimeZone {
	zoneId_alias := struct_miqt_string{}
	zoneId_alias.data = (char)(unsafe.Pointer(&zoneId[0]))
	zoneId_alias.len = size_t(len(zoneId))
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	abbreviation_ms := struct_miqt_string{}
	abbreviation_ms.data = CString(abbreviation)
	abbreviation_ms.len = size_t(len(abbreviation))
	defer free(unsafe.Pointer(abbreviation_ms.data))

	ret := newQTimeZone(QTimeZone_new7(zoneId_alias, (int)(offsetSeconds), name_ms, abbreviation_ms, territory))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone8 constructs a new QTimeZone object.
func NewQTimeZone8(zoneId []byte, offsetSeconds int, name string, abbreviation string, territory Country, comment string) *QTimeZone {
	zoneId_alias := struct_miqt_string{}
	zoneId_alias.data = (char)(unsafe.Pointer(&zoneId[0]))
	zoneId_alias.len = size_t(len(zoneId))
	name_ms := struct_miqt_string{}
	name_ms.data = CString(name)
	name_ms.len = size_t(len(name))
	defer free(unsafe.Pointer(name_ms.data))
	abbreviation_ms := struct_miqt_string{}
	abbreviation_ms.data = CString(abbreviation)
	abbreviation_ms.len = size_t(len(abbreviation))
	defer free(unsafe.Pointer(abbreviation_ms.data))
	comment_ms := struct_miqt_string{}
	comment_ms.data = CString(comment)
	comment_ms.len = size_t(len(comment))
	defer free(unsafe.Pointer(comment_ms.data))

	ret := newQTimeZone(QTimeZone_new8(zoneId_alias, (int)(offsetSeconds), name_ms, abbreviation_ms, territory, comment_ms))
	ret.isSubclass = true
	return ret
}

func (this *QTimeZone) OperatorAssign(other *QTimeZone) {
	QTimeZone_OperatorAssign(this.h, other.cPointer())
}

func (this *QTimeZone) Swap(other *QTimeZone) {
	QTimeZone_Swap(this.h, other.cPointer())
}

func (this *QTimeZone) IsValid() bool {
	return (bool)(QTimeZone_IsValid(this.h))
}

func QTimeZone_FromSecondsAheadOfUtc(offset int) *QTimeZone {
	_goptr := newQTimeZone(QTimeZone_FromSecondsAheadOfUtc((int)(offset)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTimeZone) TimeSpec() TimeSpec {
	return (TimeSpec)(QTimeZone_TimeSpec(this.h))
}

func (this *QTimeZone) FixedSecondsAheadOfUtc() int {
	return (int)(QTimeZone_FixedSecondsAheadOfUtc(this.h))
}

func QTimeZone_IsUtcOrFixedOffset(spec TimeSpec) bool {
	return (bool)(QTimeZone_IsUtcOrFixedOffset((int)(spec)))
}

func (this *QTimeZone) IsUtcOrFixedOffset2() bool {
	return (bool)(QTimeZone_IsUtcOrFixedOffset2(this.h))
}

func (this *QTimeZone) AsBackendZone() *QTimeZone {
	_goptr := newQTimeZone(QTimeZone_AsBackendZone(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTimeZone) HasAlternativeName(alias QByteArrayView) bool {
	return (bool)(QTimeZone_HasAlternativeName(this.h, alias.cPointer()))
}

func (this *QTimeZone) Id() []byte {
	var _bytearray struct_miqt_string = QTimeZone_Id(this.h)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func (this *QTimeZone) Territory() Country {
	return (Country)(QTimeZone_Territory(this.h))
}

func (this *QTimeZone) Country() QLocale__Country {
	return (QLocale__Country)(QTimeZone_Country(this.h))
}

func (this *QTimeZone) Comment() string {
	var _ms struct_miqt_string = QTimeZone_Comment(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) DisplayName(atDateTime *QDateTime) string {
	var _ms struct_miqt_string = QTimeZone_DisplayName(this.h, atDateTime.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) DisplayNameWithTimeType(timeType TimeType) string {
	var _ms struct_miqt_string = QTimeZone_DisplayNameWithTimeType(this.h, timeType)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) Abbreviation(atDateTime *QDateTime) string {
	var _ms struct_miqt_string = QTimeZone_Abbreviation(this.h, atDateTime.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) OffsetFromUtc(atDateTime *QDateTime) int {
	return (int)(QTimeZone_OffsetFromUtc(this.h, atDateTime.cPointer()))
}

func (this *QTimeZone) StandardTimeOffset(atDateTime *QDateTime) int {
	return (int)(QTimeZone_StandardTimeOffset(this.h, atDateTime.cPointer()))
}

func (this *QTimeZone) DaylightTimeOffset(atDateTime *QDateTime) int {
	return (int)(QTimeZone_DaylightTimeOffset(this.h, atDateTime.cPointer()))
}

func (this *QTimeZone) HasDaylightTime() bool {
	return (bool)(QTimeZone_HasDaylightTime(this.h))
}

func (this *QTimeZone) IsDaylightTime(atDateTime *QDateTime) bool {
	return (bool)(QTimeZone_IsDaylightTime(this.h, atDateTime.cPointer()))
}

func (this *QTimeZone) OffsetData(forDateTime *QDateTime) OffsetData {
	xxxxxxxxx
}

func (this *QTimeZone) HasTransitions() bool {
	return (bool)(QTimeZone_HasTransitions(this.h))
}

func (this *QTimeZone) NextTransition(afterDateTime *QDateTime) OffsetData {
	xxxxxxxxx
}

func (this *QTimeZone) PreviousTransition(beforeDateTime *QDateTime) OffsetData {
	xxxxxxxxx
}

func (this *QTimeZone) Transitions(fromDateTime *QDateTime, toDateTime *QDateTime) OffsetDataList {
	xxxxxxxxx
}

func QTimeZone_SystemTimeZoneId() []byte {
	var _bytearray struct_miqt_string = QTimeZone_SystemTimeZoneId()
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QTimeZone_SystemTimeZone() *QTimeZone {
	_goptr := newQTimeZone(QTimeZone_SystemTimeZone())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTimeZone_Utc() *QTimeZone {
	_goptr := newQTimeZone(QTimeZone_Utc())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTimeZone_IsTimeZoneIdAvailable(ianaId []byte) bool {
	ianaId_alias := struct_miqt_string{}
	ianaId_alias.data = (char)(unsafe.Pointer(&ianaId[0]))
	ianaId_alias.len = size_t(len(ianaId))
	return (bool)(QTimeZone_IsTimeZoneIdAvailable(ianaId_alias))
}

func QTimeZone_AvailableTimeZoneIds() [][]byte {
	var _ma struct_miqt_array = QTimeZone_AvailableTimeZoneIds()
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QTimeZone_AvailableTimeZoneIdsWithTerritory(territory Country) [][]byte {
	var _ma struct_miqt_array = QTimeZone_AvailableTimeZoneIdsWithTerritory(territory)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QTimeZone_AvailableTimeZoneIdsWithOffsetSeconds(offsetSeconds int) [][]byte {
	var _ma struct_miqt_array = QTimeZone_AvailableTimeZoneIdsWithOffsetSeconds((int)(offsetSeconds))
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QTimeZone_IanaIdToWindowsId(ianaId []byte) []byte {
	ianaId_alias := struct_miqt_string{}
	ianaId_alias.data = (char)(unsafe.Pointer(&ianaId[0]))
	ianaId_alias.len = size_t(len(ianaId))
	var _bytearray struct_miqt_string = QTimeZone_IanaIdToWindowsId(ianaId_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QTimeZone_WindowsIdToDefaultIanaId(windowsId []byte) []byte {
	windowsId_alias := struct_miqt_string{}
	windowsId_alias.data = (char)(unsafe.Pointer(&windowsId[0]))
	windowsId_alias.len = size_t(len(windowsId))
	var _bytearray struct_miqt_string = QTimeZone_WindowsIdToDefaultIanaId(windowsId_alias)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QTimeZone_WindowsIdToDefaultIanaId2(windowsId []byte, territory Country) []byte {
	windowsId_alias := struct_miqt_string{}
	windowsId_alias.data = (char)(unsafe.Pointer(&windowsId[0]))
	windowsId_alias.len = size_t(len(windowsId))
	var _bytearray struct_miqt_string = QTimeZone_WindowsIdToDefaultIanaId2(windowsId_alias, territory)
	_ret := GoBytes(unsafe.Pointer(_bytearray.data), int(int64(_bytearray.len)))
	free(unsafe.Pointer(_bytearray.data))
	return _ret
}

func QTimeZone_WindowsIdToIanaIds(windowsId []byte) [][]byte {
	windowsId_alias := struct_miqt_string{}
	windowsId_alias.data = (char)(unsafe.Pointer(&windowsId[0]))
	windowsId_alias.len = size_t(len(windowsId))
	var _ma struct_miqt_array = QTimeZone_WindowsIdToIanaIds(windowsId_alias)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func QTimeZone_WindowsIdToIanaIds2(windowsId []byte, territory Country) [][]byte {
	windowsId_alias := struct_miqt_string{}
	windowsId_alias.data = (char)(unsafe.Pointer(&windowsId[0]))
	windowsId_alias.len = size_t(len(windowsId))
	var _ma struct_miqt_array = QTimeZone_WindowsIdToIanaIds2(windowsId_alias, territory)
	_ret := make([][]byte, int(_ma.len))
	_outCast := (*[0xffff]struct_miqt_string)(unsafe.Pointer(_ma.data)) // hey ya
	for i := 0; i < int(_ma.len); i++ {
		var _lv_bytearray struct_miqt_string = _outCast[i]
		_lv_ret := GoBytes(unsafe.Pointer(_lv_bytearray.data), int(int64(_lv_bytearray.len)))
		free(unsafe.Pointer(_lv_bytearray.data))
		_ret[i] = _lv_ret
	}
	return _ret
}

func (this *QTimeZone) DisplayName2(atDateTime *QDateTime, nameType NameType) string {
	var _ms struct_miqt_string = QTimeZone_DisplayName2(this.h, atDateTime.cPointer(), nameType)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) DisplayName3(atDateTime *QDateTime, nameType NameType, locale *QLocale) string {
	var _ms struct_miqt_string = QTimeZone_DisplayName3(this.h, atDateTime.cPointer(), nameType, locale.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) DisplayName22(timeType TimeType, nameType NameType) string {
	var _ms struct_miqt_string = QTimeZone_DisplayName22(this.h, timeType, nameType)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeZone) DisplayName32(timeType TimeType, nameType NameType, locale *QLocale) string {
	var _ms struct_miqt_string = QTimeZone_DisplayName32(this.h, timeType, nameType, locale.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QTimeZone__OffsetData struct {
	h          uintptr
	isSubclass bool
}

// NewQTimeZone__OffsetData constructs a new QTimeZone::OffsetData object.
func NewQTimeZone__OffsetData(param1 *OffsetData) *QTimeZone__OffsetData {
	ret := newQTimeZone__OffsetData(QTimeZone__OffsetData_new(param1))
	ret.isSubclass = true
	return ret
}

// NewQTimeZone__OffsetData2 constructs a new QTimeZone::OffsetData object.
func NewQTimeZone__OffsetData2() *QTimeZone__OffsetData {
	ret := newQTimeZone__OffsetData(QTimeZone__OffsetData_new2())
	ret.isSubclass = true
	return ret
}

func (this *QTimeZone__OffsetData) OperatorAssign(param1 *OffsetData) {
	QTimeZone__OffsetData_OperatorAssign(this.h, param1)
}
