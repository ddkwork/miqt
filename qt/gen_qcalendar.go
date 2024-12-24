package qt

import (
	"unsafe"
)

type QCalendar__ int

const (
	QCalendar__Unspecified QCalendar__ = -2147483648
)

type QCalendar__System int

const (
	QCalendar__Gregorian    QCalendar__System = 0
	QCalendar__Julian       QCalendar__System = 8
	QCalendar__Milankovic   QCalendar__System = 9
	QCalendar__Jalali       QCalendar__System = 10
	QCalendar__IslamicCivil QCalendar__System = 11
	QCalendar__Last         QCalendar__System = 11
	QCalendar__User         QCalendar__System = -1
)

type QCalendar struct {
	h          uintptr
	isSubclass bool
}

// NewQCalendar constructs a new QCalendar object.
func NewQCalendar() *QCalendar {

	ret := newQCalendar(QCalendar_new())
	ret.isSubclass = true
	return ret
}

// NewQCalendar2 constructs a new QCalendar object.
func NewQCalendar2(system System) *QCalendar {

	ret := newQCalendar(QCalendar_new2(system))
	ret.isSubclass = true
	return ret
}

// NewQCalendar3 constructs a new QCalendar object.
func NewQCalendar3(name QAnyStringView) *QCalendar {

	ret := newQCalendar(QCalendar_new3(name.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCalendar4 constructs a new QCalendar object.
func NewQCalendar4(id SystemId) *QCalendar {

	ret := newQCalendar(QCalendar_new4(id))
	ret.isSubclass = true
	return ret
}

func (this *QCalendar) IsValid() bool {
	return (bool)(QCalendar_IsValid(this.h))
}

func (this *QCalendar) DaysInMonth(month int) int {
	return (int)(QCalendar_DaysInMonth(this.h, (int)(month)))
}

func (this *QCalendar) DaysInYear(year int) int {
	return (int)(QCalendar_DaysInYear(this.h, (int)(year)))
}

func (this *QCalendar) MonthsInYear(year int) int {
	return (int)(QCalendar_MonthsInYear(this.h, (int)(year)))
}

func (this *QCalendar) IsDateValid(year int, month int, day int) bool {
	return (bool)(QCalendar_IsDateValid(this.h, (int)(year), (int)(month), (int)(day)))
}

func (this *QCalendar) IsLeapYear(year int) bool {
	return (bool)(QCalendar_IsLeapYear(this.h, (int)(year)))
}

func (this *QCalendar) IsGregorian() bool {
	return (bool)(QCalendar_IsGregorian(this.h))
}

func (this *QCalendar) IsLunar() bool {
	return (bool)(QCalendar_IsLunar(this.h))
}

func (this *QCalendar) IsLuniSolar() bool {
	return (bool)(QCalendar_IsLuniSolar(this.h))
}

func (this *QCalendar) IsSolar() bool {
	return (bool)(QCalendar_IsSolar(this.h))
}

func (this *QCalendar) IsProleptic() bool {
	return (bool)(QCalendar_IsProleptic(this.h))
}

func (this *QCalendar) HasYearZero() bool {
	return (bool)(QCalendar_HasYearZero(this.h))
}

func (this *QCalendar) MaximumDaysInMonth() int {
	return (int)(QCalendar_MaximumDaysInMonth(this.h))
}

func (this *QCalendar) MinimumDaysInMonth() int {
	return (int)(QCalendar_MinimumDaysInMonth(this.h))
}

func (this *QCalendar) MaximumMonthsInYear() int {
	return (int)(QCalendar_MaximumMonthsInYear(this.h))
}

func (this *QCalendar) Name() string {
	var _ms struct_miqt_string = QCalendar_Name(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) DateFromParts(year int, month int, day int) *QDate {
	_goptr := newQDate(QCalendar_DateFromParts(this.h, (int)(year), (int)(month), (int)(day)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendar) DateFromPartsWithParts(parts *YearMonthDay) *QDate {
	_goptr := newQDate(QCalendar_DateFromPartsWithParts(this.h, parts))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendar) MatchCenturyToWeekday(parts *YearMonthDay, dow int) *QDate {
	_goptr := newQDate(QCalendar_MatchCenturyToWeekday(this.h, parts, (int)(dow)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendar) PartsFromDate(date QDate) YearMonthDay {
	xxxxxxxxx
}

func (this *QCalendar) DayOfWeek(date QDate) int {
	return (int)(QCalendar_DayOfWeek(this.h, date.cPointer()))
}

func (this *QCalendar) MonthName(locale *QLocale, month int) string {
	var _ms struct_miqt_string = QCalendar_MonthName(this.h, locale.cPointer(), (int)(month))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) StandaloneMonthName(locale *QLocale, month int) string {
	var _ms struct_miqt_string = QCalendar_StandaloneMonthName(this.h, locale.cPointer(), (int)(month))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) WeekDayName(locale *QLocale, day int) string {
	var _ms struct_miqt_string = QCalendar_WeekDayName(this.h, locale.cPointer(), (int)(day))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) StandaloneWeekDayName(locale *QLocale, day int) string {
	var _ms struct_miqt_string = QCalendar_StandaloneWeekDayName(this.h, locale.cPointer(), (int)(day))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCalendar_AvailableCalendars() []string {
	var _ma struct_miqt_array = QCalendar_AvailableCalendars()
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

func (this *QCalendar) DaysInMonth2(month int, year int) int {
	return (int)(QCalendar_DaysInMonth2(this.h, (int)(month), (int)(year)))
}

func (this *QCalendar) MonthName3(locale *QLocale, month int, year int) string {
	var _ms struct_miqt_string = QCalendar_MonthName3(this.h, locale.cPointer(), (int)(month), (int)(year))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) MonthName4(locale *QLocale, month int, year int, format QLocale__FormatType) string {
	var _ms struct_miqt_string = QCalendar_MonthName4(this.h, locale.cPointer(), (int)(month), (int)(year), (int)(format))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) StandaloneMonthName3(locale *QLocale, month int, year int) string {
	var _ms struct_miqt_string = QCalendar_StandaloneMonthName3(this.h, locale.cPointer(), (int)(month), (int)(year))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) StandaloneMonthName4(locale *QLocale, month int, year int, format QLocale__FormatType) string {
	var _ms struct_miqt_string = QCalendar_StandaloneMonthName4(this.h, locale.cPointer(), (int)(month), (int)(year), (int)(format))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) WeekDayName3(locale *QLocale, day int, format QLocale__FormatType) string {
	var _ms struct_miqt_string = QCalendar_WeekDayName3(this.h, locale.cPointer(), (int)(day), (int)(format))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendar) StandaloneWeekDayName3(locale *QLocale, day int, format QLocale__FormatType) string {
	var _ms struct_miqt_string = QCalendar_StandaloneWeekDayName3(this.h, locale.cPointer(), (int)(day), (int)(format))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

type QCalendar__YearMonthDay struct {
	h          uintptr
	isSubclass bool
}

// NewQCalendar__YearMonthDay constructs a new QCalendar::YearMonthDay object.
func NewQCalendar__YearMonthDay() *QCalendar__YearMonthDay {

	ret := newQCalendar__YearMonthDay(QCalendar__YearMonthDay_new())
	ret.isSubclass = true
	return ret
}

// NewQCalendar__YearMonthDay2 constructs a new QCalendar::YearMonthDay object.
func NewQCalendar__YearMonthDay2(y int) *QCalendar__YearMonthDay {

	ret := newQCalendar__YearMonthDay(QCalendar__YearMonthDay_new2((int)(y)))
	ret.isSubclass = true
	return ret
}

// NewQCalendar__YearMonthDay3 constructs a new QCalendar::YearMonthDay object.
func NewQCalendar__YearMonthDay3(y int, m int) *QCalendar__YearMonthDay {

	ret := newQCalendar__YearMonthDay(QCalendar__YearMonthDay_new3((int)(y), (int)(m)))
	ret.isSubclass = true
	return ret
}

// NewQCalendar__YearMonthDay4 constructs a new QCalendar::YearMonthDay object.
func NewQCalendar__YearMonthDay4(y int, m int, d int) *QCalendar__YearMonthDay {

	ret := newQCalendar__YearMonthDay(QCalendar__YearMonthDay_new4((int)(y), (int)(m), (int)(d)))
	ret.isSubclass = true
	return ret
}

func (this *QCalendar__YearMonthDay) IsValid() bool {
	return (bool)(QCalendar__YearMonthDay_IsValid(this.h))
}

type QCalendar__SystemId struct {
	h          uintptr
	isSubclass bool
}

// NewQCalendar__SystemId constructs a new QCalendar::SystemId object.
func NewQCalendar__SystemId() *QCalendar__SystemId {

	ret := newQCalendar__SystemId(QCalendar__SystemId_new())
	ret.isSubclass = true
	return ret
}

func (this *QCalendar__SystemId) Index() uint64 {
	return (uint64)(QCalendar__SystemId_Index(this.h))
}

func (this *QCalendar__SystemId) IsValid() bool {
	return (bool)(QCalendar__SystemId_IsValid(this.h))
}
