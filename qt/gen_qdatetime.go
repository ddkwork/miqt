package qt

import (
	"unsafe"
)

type QDateTime__TransitionResolution int

const (
	QDateTime__Reject               QDateTime__TransitionResolution = 0
	QDateTime__RelativeToBefore     QDateTime__TransitionResolution = 1
	QDateTime__RelativeToAfter      QDateTime__TransitionResolution = 2
	QDateTime__PreferBefore         QDateTime__TransitionResolution = 3
	QDateTime__PreferAfter          QDateTime__TransitionResolution = 4
	QDateTime__PreferStandard       QDateTime__TransitionResolution = 5
	QDateTime__PreferDaylightSaving QDateTime__TransitionResolution = 6
	QDateTime__LegacyBehavior       QDateTime__TransitionResolution = 1
)

type QDateTime__YearRange int

const (
	QDateTime__First QDateTime__YearRange = -292275056
	QDateTime__Last  QDateTime__YearRange = 292278994
)

type QDate struct {
	h          uintptr
	isSubclass bool
}

// NewQDate constructs a new QDate object.
func NewQDate() *QDate {

	ret := newQDate(QDate_new())
	ret.isSubclass = true
	return ret
}

// NewQDate2 constructs a new QDate object.
func NewQDate2(y int, m int, d int) *QDate {

	ret := newQDate(QDate_new2((int)(y), (int)(m), (int)(d)))
	ret.isSubclass = true
	return ret
}

// NewQDate3 constructs a new QDate object.
func NewQDate3(y int, m int, d int, cal QCalendar) *QDate {

	ret := newQDate(QDate_new3((int)(y), (int)(m), (int)(d), cal.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDate4 constructs a new QDate object.
func NewQDate4(param1 *QDate) *QDate {

	ret := newQDate(QDate_new4(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDate) IsNull() bool {
	return (bool)(QDate_IsNull(this.h))
}

func (this *QDate) IsValid() bool {
	return (bool)(QDate_IsValid(this.h))
}

func (this *QDate) Year() int {
	return (int)(QDate_Year(this.h))
}

func (this *QDate) Month() int {
	return (int)(QDate_Month(this.h))
}

func (this *QDate) Day() int {
	return (int)(QDate_Day(this.h))
}

func (this *QDate) DayOfWeek() int {
	return (int)(QDate_DayOfWeek(this.h))
}

func (this *QDate) DayOfYear() int {
	return (int)(QDate_DayOfYear(this.h))
}

func (this *QDate) DaysInMonth() int {
	return (int)(QDate_DaysInMonth(this.h))
}

func (this *QDate) DaysInYear() int {
	return (int)(QDate_DaysInYear(this.h))
}

func (this *QDate) WeekNumber() int {
	return (int)(QDate_WeekNumber(this.h))
}

func (this *QDate) YearWithCal(cal QCalendar) int {
	return (int)(QDate_YearWithCal(this.h, cal.cPointer()))
}

func (this *QDate) MonthWithCal(cal QCalendar) int {
	return (int)(QDate_MonthWithCal(this.h, cal.cPointer()))
}

func (this *QDate) DayWithCal(cal QCalendar) int {
	return (int)(QDate_DayWithCal(this.h, cal.cPointer()))
}

func (this *QDate) DayOfWeekWithCal(cal QCalendar) int {
	return (int)(QDate_DayOfWeekWithCal(this.h, cal.cPointer()))
}

func (this *QDate) DayOfYearWithCal(cal QCalendar) int {
	return (int)(QDate_DayOfYearWithCal(this.h, cal.cPointer()))
}

func (this *QDate) DaysInMonthWithCal(cal QCalendar) int {
	return (int)(QDate_DaysInMonthWithCal(this.h, cal.cPointer()))
}

func (this *QDate) DaysInYearWithCal(cal QCalendar) int {
	return (int)(QDate_DaysInYearWithCal(this.h, cal.cPointer()))
}

func (this *QDate) StartOfDay(spec TimeSpec) *QDateTime {
	_goptr := newQDateTime(QDate_StartOfDay(this.h, (int)(spec)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) EndOfDay(spec TimeSpec) *QDateTime {
	_goptr := newQDateTime(QDate_EndOfDay(this.h, (int)(spec)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) StartOfDayWithZone(zone *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDate_StartOfDayWithZone(this.h, zone.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) EndOfDayWithZone(zone *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDate_EndOfDayWithZone(this.h, zone.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) StartOfDay2() *QDateTime {
	_goptr := newQDateTime(QDate_StartOfDay2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) EndOfDay2() *QDateTime {
	_goptr := newQDateTime(QDate_EndOfDay2(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) ToString() string {
	var _ms struct_miqt_string = QDate_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDate) ToStringWithFormat(format string) string {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	var _ms struct_miqt_string = QDate_ToStringWithFormat(this.h, format_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDate) ToString2(format string, cal QCalendar) string {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	var _ms struct_miqt_string = QDate_ToString2(this.h, format_ms, cal.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDate) SetDate(year int, month int, day int) bool {
	return (bool)(QDate_SetDate(this.h, (int)(year), (int)(month), (int)(day)))
}

func (this *QDate) SetDate2(year int, month int, day int, cal QCalendar) bool {
	return (bool)(QDate_SetDate2(this.h, (int)(year), (int)(month), (int)(day), cal.cPointer()))
}

func (this *QDate) GetDate(year *int, month *int, day *int) {
	QDate_GetDate(this.h, (*int)(unsafe.Pointer(year)), (*int)(unsafe.Pointer(month)), (*int)(unsafe.Pointer(day)))
}

func (this *QDate) AddDays(days int64) *QDate {
	_goptr := newQDate(QDate_AddDays(this.h, (longlong)(days)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) AddMonths(months int) *QDate {
	_goptr := newQDate(QDate_AddMonths(this.h, (int)(months)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) AddYears(years int) *QDate {
	_goptr := newQDate(QDate_AddYears(this.h, (int)(years)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) AddMonths2(months int, cal QCalendar) *QDate {
	_goptr := newQDate(QDate_AddMonths2(this.h, (int)(months), cal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) AddYears2(years int, cal QCalendar) *QDate {
	_goptr := newQDate(QDate_AddYears2(this.h, (int)(years), cal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) DaysTo(d QDate) int64 {
	return (int64)(QDate_DaysTo(this.h, d.cPointer()))
}

func QDate_CurrentDate() *QDate {
	_goptr := newQDate(QDate_CurrentDate())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDate_FromStringWithStringVal(stringVal string) *QDate {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQDate(QDate_FromStringWithStringVal(stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDate_FromString4(stringVal string, format string, cal QCalendar) *QDate {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDate(QDate_FromString4(stringVal_ms, format_ms, cal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDate_FromString9(stringVal string, format string) *QDate {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDate(QDate_FromString9(stringVal_ms, format_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDate_FromString10(stringVal string, format string, baseYear int, cal QCalendar) *QDate {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDate(QDate_FromString10(stringVal_ms, format_ms, (int)(baseYear), cal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDate_IsValid2(y int, m int, d int) bool {
	return (bool)(QDate_IsValid2((int)(y), (int)(m), (int)(d)))
}

func QDate_IsLeapYear(year int) bool {
	return (bool)(QDate_IsLeapYear((int)(year)))
}

func QDate_FromJulianDay(jd_ int64) *QDate {
	_goptr := newQDate(QDate_FromJulianDay((longlong)(jd_)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) ToJulianDay() int64 {
	return (int64)(QDate_ToJulianDay(this.h))
}

func (this *QDate) WeekNumber1(yearNum *int) int {
	return (int)(QDate_WeekNumber1(this.h, (*int)(unsafe.Pointer(yearNum))))
}

func (this *QDate) StartOfDay22(spec TimeSpec, offsetSeconds int) *QDateTime {
	_goptr := newQDateTime(QDate_StartOfDay22(this.h, (int)(spec), (int)(offsetSeconds)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) EndOfDay22(spec TimeSpec, offsetSeconds int) *QDateTime {
	_goptr := newQDateTime(QDate_EndOfDay22(this.h, (int)(spec), (int)(offsetSeconds)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDate) ToString1(format DateFormat) string {
	var _ms struct_miqt_string = QDate_ToString1(this.h, (int)(format))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDate_FromString23(stringVal string, format DateFormat) *QDate {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQDate(QDate_FromString23(stringVal_ms, (int)(format)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDate_FromString34(stringVal string, format string, baseYear int) *QDate {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDate(QDate_FromString34(stringVal_ms, format_ms, (int)(baseYear)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

type QTime struct {
	h          uintptr
	isSubclass bool
}

// NewQTime constructs a new QTime object.
func NewQTime() *QTime {

	ret := newQTime(QTime_new())
	ret.isSubclass = true
	return ret
}

// NewQTime2 constructs a new QTime object.
func NewQTime2(h int, m int) *QTime {

	ret := newQTime(QTime_new2((int)(h), (int)(m)))
	ret.isSubclass = true
	return ret
}

// NewQTime3 constructs a new QTime object.
func NewQTime3(param1 *QTime) *QTime {

	ret := newQTime(QTime_new3(param1.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTime4 constructs a new QTime object.
func NewQTime4(h int, m int, s int) *QTime {

	ret := newQTime(QTime_new4((int)(h), (int)(m), (int)(s)))
	ret.isSubclass = true
	return ret
}

// NewQTime5 constructs a new QTime object.
func NewQTime5(h int, m int, s int, ms int) *QTime {

	ret := newQTime(QTime_new5((int)(h), (int)(m), (int)(s), (int)(ms)))
	ret.isSubclass = true
	return ret
}

func (this *QTime) IsNull() bool {
	return (bool)(QTime_IsNull(this.h))
}

func (this *QTime) IsValid() bool {
	return (bool)(QTime_IsValid(this.h))
}

func (this *QTime) Hour() int {
	return (int)(QTime_Hour(this.h))
}

func (this *QTime) Minute() int {
	return (int)(QTime_Minute(this.h))
}

func (this *QTime) Second() int {
	return (int)(QTime_Second(this.h))
}

func (this *QTime) Msec() int {
	return (int)(QTime_Msec(this.h))
}

func (this *QTime) ToString() string {
	var _ms struct_miqt_string = QTime_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTime) ToStringWithFormat(format string) string {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	var _ms struct_miqt_string = QTime_ToStringWithFormat(this.h, format_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTime) SetHMS(h int, m int, s int) bool {
	return (bool)(QTime_SetHMS(this.h, (int)(h), (int)(m), (int)(s)))
}

func (this *QTime) AddSecs(secs int) *QTime {
	_goptr := newQTime(QTime_AddSecs(this.h, (int)(secs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTime) SecsTo(t QTime) int {
	return (int)(QTime_SecsTo(this.h, t.cPointer()))
}

func (this *QTime) AddMSecs(ms int) *QTime {
	_goptr := newQTime(QTime_AddMSecs(this.h, (int)(ms)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTime) MsecsTo(t QTime) int {
	return (int)(QTime_MsecsTo(this.h, t.cPointer()))
}

func QTime_FromMSecsSinceStartOfDay(msecs int) *QTime {
	_goptr := newQTime(QTime_FromMSecsSinceStartOfDay((int)(msecs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QTime) MsecsSinceStartOfDay() int {
	return (int)(QTime_MsecsSinceStartOfDay(this.h))
}

func QTime_CurrentTime() *QTime {
	_goptr := newQTime(QTime_CurrentTime())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTime_FromStringWithStringVal(stringVal string) *QTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQTime(QTime_FromStringWithStringVal(stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTime_FromString4(stringVal string, format string) *QTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQTime(QTime_FromString4(stringVal_ms, format_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTime_IsValid2(h int, m int, s int) bool {
	return (bool)(QTime_IsValid2((int)(h), (int)(m), (int)(s)))
}

func (this *QTime) ToString1(f DateFormat) string {
	var _ms struct_miqt_string = QTime_ToString1(this.h, (int)(f))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTime) SetHMS4(h int, m int, s int, ms int) bool {
	return (bool)(QTime_SetHMS4(this.h, (int)(h), (int)(m), (int)(s), (int)(ms)))
}

func QTime_FromString23(stringVal string, format DateFormat) *QTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQTime(QTime_FromString23(stringVal_ms, (int)(format)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QTime_IsValid4(h int, m int, s int, ms int) bool {
	return (bool)(QTime_IsValid4((int)(h), (int)(m), (int)(s), (int)(ms)))
}

type QDateTime struct {
	h          uintptr
	isSubclass bool
}

// NewQDateTime constructs a new QDateTime object.
func NewQDateTime() *QDateTime {

	ret := newQDateTime(QDateTime_new())
	ret.isSubclass = true
	return ret
}

// NewQDateTime2 constructs a new QDateTime object.
func NewQDateTime2(date QDate, time QTime, spec TimeSpec) *QDateTime {

	ret := newQDateTime(QDateTime_new2(date.cPointer(), time.cPointer(), (int)(spec)))
	ret.isSubclass = true
	return ret
}

// NewQDateTime3 constructs a new QDateTime object.
func NewQDateTime3(date QDate, time QTime, timeZone *QTimeZone) *QDateTime {

	ret := newQDateTime(QDateTime_new3(date.cPointer(), time.cPointer(), timeZone.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTime4 constructs a new QDateTime object.
func NewQDateTime4(date QDate, time QTime) *QDateTime {

	ret := newQDateTime(QDateTime_new4(date.cPointer(), time.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTime5 constructs a new QDateTime object.
func NewQDateTime5(other *QDateTime) *QDateTime {

	ret := newQDateTime(QDateTime_new5(other.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTime6 constructs a new QDateTime object.
func NewQDateTime6(date QDate, time QTime, spec TimeSpec, offsetSeconds int) *QDateTime {

	ret := newQDateTime(QDateTime_new6(date.cPointer(), time.cPointer(), (int)(spec), (int)(offsetSeconds)))
	ret.isSubclass = true
	return ret
}

// NewQDateTime7 constructs a new QDateTime object.
func NewQDateTime7(date QDate, time QTime, timeZone *QTimeZone, resolve TransitionResolution) *QDateTime {

	ret := newQDateTime(QDateTime_new7(date.cPointer(), time.cPointer(), timeZone.cPointer(), resolve))
	ret.isSubclass = true
	return ret
}

// NewQDateTime8 constructs a new QDateTime object.
func NewQDateTime8(date QDate, time QTime, resolve TransitionResolution) *QDateTime {

	ret := newQDateTime(QDateTime_new8(date.cPointer(), time.cPointer(), resolve))
	ret.isSubclass = true
	return ret
}

func (this *QDateTime) OperatorAssign(other *QDateTime) {
	QDateTime_OperatorAssign(this.h, other.cPointer())
}

func (this *QDateTime) Swap(other *QDateTime) {
	QDateTime_Swap(this.h, other.cPointer())
}

func (this *QDateTime) IsNull() bool {
	return (bool)(QDateTime_IsNull(this.h))
}

func (this *QDateTime) IsValid() bool {
	return (bool)(QDateTime_IsValid(this.h))
}

func (this *QDateTime) Date() *QDate {
	_goptr := newQDate(QDateTime_Date(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) Time() *QTime {
	_goptr := newQTime(QDateTime_Time(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) TimeSpec() TimeSpec {
	return (TimeSpec)(QDateTime_TimeSpec(this.h))
}

func (this *QDateTime) OffsetFromUtc() int {
	return (int)(QDateTime_OffsetFromUtc(this.h))
}

func (this *QDateTime) TimeRepresentation() *QTimeZone {
	_goptr := newQTimeZone(QDateTime_TimeRepresentation(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) TimeZone() *QTimeZone {
	_goptr := newQTimeZone(QDateTime_TimeZone(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) TimeZoneAbbreviation() string {
	var _ms struct_miqt_string = QDateTime_TimeZoneAbbreviation(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTime) IsDaylightTime() bool {
	return (bool)(QDateTime_IsDaylightTime(this.h))
}

func (this *QDateTime) ToMSecsSinceEpoch() int64 {
	return (int64)(QDateTime_ToMSecsSinceEpoch(this.h))
}

func (this *QDateTime) ToSecsSinceEpoch() int64 {
	return (int64)(QDateTime_ToSecsSinceEpoch(this.h))
}

func (this *QDateTime) SetDate(date QDate) {
	QDateTime_SetDate(this.h, date.cPointer())
}

func (this *QDateTime) SetTime(time QTime) {
	QDateTime_SetTime(this.h, time.cPointer())
}

func (this *QDateTime) SetTimeSpec(spec TimeSpec) {
	QDateTime_SetTimeSpec(this.h, (int)(spec))
}

func (this *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	QDateTime_SetOffsetFromUtc(this.h, (int)(offsetSeconds))
}

func (this *QDateTime) SetTimeZone(toZone *QTimeZone) {
	QDateTime_SetTimeZone(this.h, toZone.cPointer())
}

func (this *QDateTime) SetMSecsSinceEpoch(msecs int64) {
	QDateTime_SetMSecsSinceEpoch(this.h, (longlong)(msecs))
}

func (this *QDateTime) SetSecsSinceEpoch(secs int64) {
	QDateTime_SetSecsSinceEpoch(this.h, (longlong)(secs))
}

func (this *QDateTime) ToString() string {
	var _ms struct_miqt_string = QDateTime_ToString(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTime) ToStringWithFormat(format string) string {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	var _ms struct_miqt_string = QDateTime_ToStringWithFormat(this.h, format_ms)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTime) ToString2(format string, cal QCalendar) string {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	var _ms struct_miqt_string = QDateTime_ToString2(this.h, format_ms, cal.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTime) AddDays(days int64) *QDateTime {
	_goptr := newQDateTime(QDateTime_AddDays(this.h, (longlong)(days)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) AddMonths(months int) *QDateTime {
	_goptr := newQDateTime(QDateTime_AddMonths(this.h, (int)(months)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) AddYears(years int) *QDateTime {
	_goptr := newQDateTime(QDateTime_AddYears(this.h, (int)(years)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) AddSecs(secs int64) *QDateTime {
	_goptr := newQDateTime(QDateTime_AddSecs(this.h, (longlong)(secs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) AddMSecs(msecs int64) *QDateTime {
	_goptr := newQDateTime(QDateTime_AddMSecs(this.h, (longlong)(msecs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) ToTimeSpec(spec TimeSpec) *QDateTime {
	_goptr := newQDateTime(QDateTime_ToTimeSpec(this.h, (int)(spec)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) ToLocalTime() *QDateTime {
	_goptr := newQDateTime(QDateTime_ToLocalTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) ToUTC() *QDateTime {
	_goptr := newQDateTime(QDateTime_ToUTC(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) ToOffsetFromUtc(offsetSeconds int) *QDateTime {
	_goptr := newQDateTime(QDateTime_ToOffsetFromUtc(this.h, (int)(offsetSeconds)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) ToTimeZone(toZone *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDateTime_ToTimeZone(this.h, toZone.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTime) DaysTo(param1 *QDateTime) int64 {
	return (int64)(QDateTime_DaysTo(this.h, param1.cPointer()))
}

func (this *QDateTime) SecsTo(param1 *QDateTime) int64 {
	return (int64)(QDateTime_SecsTo(this.h, param1.cPointer()))
}

func (this *QDateTime) MsecsTo(param1 *QDateTime) int64 {
	return (int64)(QDateTime_MsecsTo(this.h, param1.cPointer()))
}

func QDateTime_CurrentDateTime(zone *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDateTime_CurrentDateTime(zone.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_CurrentDateTime2() *QDateTime {
	_goptr := newQDateTime(QDateTime_CurrentDateTime2())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_CurrentDateTimeUtc() *QDateTime {
	_goptr := newQDateTime(QDateTime_CurrentDateTimeUtc())
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromStringWithStringVal(stringVal string) *QDateTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQDateTime(QDateTime_FromStringWithStringVal(stringVal_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromString4(stringVal string, format string, cal QCalendar) *QDateTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDateTime(QDateTime_FromString4(stringVal_ms, format_ms, cal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromString9(stringVal string, format string) *QDateTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDateTime(QDateTime_FromString9(stringVal_ms, format_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromString10(stringVal string, format string, baseYear int, cal QCalendar) *QDateTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDateTime(QDateTime_FromString10(stringVal_ms, format_ms, (int)(baseYear), cal.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromMSecsSinceEpoch(msecs int64, spec TimeSpec) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromMSecsSinceEpoch((longlong)(msecs), (int)(spec)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromSecsSinceEpoch(secs int64, spec TimeSpec) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromSecsSinceEpoch((longlong)(secs), (int)(spec)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromMSecsSinceEpoch2(msecs int64, timeZone *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromMSecsSinceEpoch2((longlong)(msecs), timeZone.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromSecsSinceEpoch2(secs int64, timeZone *QTimeZone) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromSecsSinceEpoch2((longlong)(secs), timeZone.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromMSecsSinceEpochWithMsecs(msecs int64) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromMSecsSinceEpochWithMsecs((longlong)(msecs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromSecsSinceEpochWithSecs(secs int64) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromSecsSinceEpochWithSecs((longlong)(secs)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_CurrentMSecsSinceEpoch() int64 {
	return (int64)(QDateTime_CurrentMSecsSinceEpoch())
}

func QDateTime_CurrentSecsSinceEpoch() int64 {
	return (int64)(QDateTime_CurrentSecsSinceEpoch())
}

func (this *QDateTime) SetDate2(date QDate, resolve TransitionResolution) {
	QDateTime_SetDate2(this.h, date.cPointer(), resolve)
}

func (this *QDateTime) SetTime2(time QTime, resolve TransitionResolution) {
	QDateTime_SetTime2(this.h, time.cPointer(), resolve)
}

func (this *QDateTime) SetTimeZone2(toZone *QTimeZone, resolve TransitionResolution) {
	QDateTime_SetTimeZone2(this.h, toZone.cPointer(), resolve)
}

func (this *QDateTime) ToString1(format DateFormat) string {
	var _ms struct_miqt_string = QDateTime_ToString1(this.h, (int)(format))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDateTime_FromString23(stringVal string, format DateFormat) *QDateTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	_goptr := newQDateTime(QDateTime_FromString23(stringVal_ms, (int)(format)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromString34(stringVal string, format string, baseYear int) *QDateTime {
	stringVal_ms := struct_miqt_string{}
	stringVal_ms.data = CString(stringVal)
	stringVal_ms.len = size_t(len(stringVal))
	defer free(unsafe.Pointer(stringVal_ms.data))
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	_goptr := newQDateTime(QDateTime_FromString34(stringVal_ms, format_ms, (int)(baseYear)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromMSecsSinceEpoch3(msecs int64, spec TimeSpec, offsetFromUtc int) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromMSecsSinceEpoch3((longlong)(msecs), (int)(spec), (int)(offsetFromUtc)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func QDateTime_FromSecsSinceEpoch3(secs int64, spec TimeSpec, offsetFromUtc int) *QDateTime {
	_goptr := newQDateTime(QDateTime_FromSecsSinceEpoch3((longlong)(secs), (int)(spec), (int)(offsetFromUtc)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}
