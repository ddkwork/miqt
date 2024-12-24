package qt

import (
	"unsafe"
)

type QDateTimeEdit__Section int

const (
	QDateTimeEdit__NoSection         QDateTimeEdit__Section = 0
	QDateTimeEdit__AmPmSection       QDateTimeEdit__Section = 1
	QDateTimeEdit__MSecSection       QDateTimeEdit__Section = 2
	QDateTimeEdit__SecondSection     QDateTimeEdit__Section = 4
	QDateTimeEdit__MinuteSection     QDateTimeEdit__Section = 8
	QDateTimeEdit__HourSection       QDateTimeEdit__Section = 16
	QDateTimeEdit__DaySection        QDateTimeEdit__Section = 256
	QDateTimeEdit__MonthSection      QDateTimeEdit__Section = 512
	QDateTimeEdit__YearSection       QDateTimeEdit__Section = 1024
	QDateTimeEdit__TimeSections_Mask QDateTimeEdit__Section = 31
	QDateTimeEdit__DateSections_Mask QDateTimeEdit__Section = 1792
)

type QDateTimeEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQDateTimeEdit constructs a new QDateTimeEdit object.
func NewQDateTimeEdit(parent *QWidget) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit2 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit2() *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit3 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit3(dt *QDateTime) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new3(dt.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit4 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit4(d QDate) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new4(d.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit5 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit5(t QTime) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new5(t.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit6 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit6(dt *QDateTime, parent *QWidget) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new6(dt.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit7 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit7(d QDate, parent *QWidget) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new7(d.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateTimeEdit8 constructs a new QDateTimeEdit object.
func NewQDateTimeEdit8(t QTime, parent *QWidget) *QDateTimeEdit {
	ret := newQDateTimeEdit(QDateTimeEdit_new8(t.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDateTimeEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QDateTimeEdit_MetaObject(this.h))
}

func (this *QDateTimeEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDateTimeEdit_Metacast(this.h, param1_Cstring))
}

func QDateTimeEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDateTimeEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTimeEdit) DateTime() *QDateTime {
	_goptr := newQDateTime(QDateTimeEdit_DateTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) Date() *QDate {
	_goptr := newQDate(QDateTimeEdit_Date(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) Time() *QTime {
	_goptr := newQTime(QDateTimeEdit_Time(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) Calendar() *QCalendar {
	_goptr := newQCalendar(QDateTimeEdit_Calendar(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) SetCalendar(calendar QCalendar) {
	QDateTimeEdit_SetCalendar(this.h, calendar.cPointer())
}

func (this *QDateTimeEdit) MinimumDateTime() *QDateTime {
	_goptr := newQDateTime(QDateTimeEdit_MinimumDateTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) ClearMinimumDateTime() {
	QDateTimeEdit_ClearMinimumDateTime(this.h)
}

func (this *QDateTimeEdit) SetMinimumDateTime(dt *QDateTime) {
	QDateTimeEdit_SetMinimumDateTime(this.h, dt.cPointer())
}

func (this *QDateTimeEdit) MaximumDateTime() *QDateTime {
	_goptr := newQDateTime(QDateTimeEdit_MaximumDateTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) ClearMaximumDateTime() {
	QDateTimeEdit_ClearMaximumDateTime(this.h)
}

func (this *QDateTimeEdit) SetMaximumDateTime(dt *QDateTime) {
	QDateTimeEdit_SetMaximumDateTime(this.h, dt.cPointer())
}

func (this *QDateTimeEdit) SetDateTimeRange(min *QDateTime, max *QDateTime) {
	QDateTimeEdit_SetDateTimeRange(this.h, min.cPointer(), max.cPointer())
}

func (this *QDateTimeEdit) MinimumDate() *QDate {
	_goptr := newQDate(QDateTimeEdit_MinimumDate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) SetMinimumDate(min QDate) {
	QDateTimeEdit_SetMinimumDate(this.h, min.cPointer())
}

func (this *QDateTimeEdit) ClearMinimumDate() {
	QDateTimeEdit_ClearMinimumDate(this.h)
}

func (this *QDateTimeEdit) MaximumDate() *QDate {
	_goptr := newQDate(QDateTimeEdit_MaximumDate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) SetMaximumDate(max QDate) {
	QDateTimeEdit_SetMaximumDate(this.h, max.cPointer())
}

func (this *QDateTimeEdit) ClearMaximumDate() {
	QDateTimeEdit_ClearMaximumDate(this.h)
}

func (this *QDateTimeEdit) SetDateRange(min QDate, max QDate) {
	QDateTimeEdit_SetDateRange(this.h, min.cPointer(), max.cPointer())
}

func (this *QDateTimeEdit) MinimumTime() *QTime {
	_goptr := newQTime(QDateTimeEdit_MinimumTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) SetMinimumTime(min QTime) {
	QDateTimeEdit_SetMinimumTime(this.h, min.cPointer())
}

func (this *QDateTimeEdit) ClearMinimumTime() {
	QDateTimeEdit_ClearMinimumTime(this.h)
}

func (this *QDateTimeEdit) MaximumTime() *QTime {
	_goptr := newQTime(QDateTimeEdit_MaximumTime(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) SetMaximumTime(max QTime) {
	QDateTimeEdit_SetMaximumTime(this.h, max.cPointer())
}

func (this *QDateTimeEdit) ClearMaximumTime() {
	QDateTimeEdit_ClearMaximumTime(this.h)
}

func (this *QDateTimeEdit) SetTimeRange(min QTime, max QTime) {
	QDateTimeEdit_SetTimeRange(this.h, min.cPointer(), max.cPointer())
}

func (this *QDateTimeEdit) DisplayedSections() Sections {
	xxxxxxxxx
}

func (this *QDateTimeEdit) CurrentSection() Section {
	xxxxxxxxx
}

func (this *QDateTimeEdit) SectionAt(index int) Section {
	xxxxxxxxx
}

func (this *QDateTimeEdit) SetCurrentSection(section Section) {
	QDateTimeEdit_SetCurrentSection(this.h, section)
}

func (this *QDateTimeEdit) CurrentSectionIndex() int {
	return (int)(QDateTimeEdit_CurrentSectionIndex(this.h))
}

func (this *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	QDateTimeEdit_SetCurrentSectionIndex(this.h, (int)(index))
}

func (this *QDateTimeEdit) CalendarWidget() *QCalendarWidget {
	return newQCalendarWidget(QDateTimeEdit_CalendarWidget(this.h))
}

func (this *QDateTimeEdit) SetCalendarWidget(calendarWidget *QCalendarWidget) {
	QDateTimeEdit_SetCalendarWidget(this.h, calendarWidget.cPointer())
}

func (this *QDateTimeEdit) SectionCount() int {
	return (int)(QDateTimeEdit_SectionCount(this.h))
}

func (this *QDateTimeEdit) SetSelectedSection(section Section) {
	QDateTimeEdit_SetSelectedSection(this.h, section)
}

func (this *QDateTimeEdit) SectionText(section Section) string {
	var _ms struct_miqt_string = QDateTimeEdit_SectionText(this.h, section)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTimeEdit) DisplayFormat() string {
	var _ms struct_miqt_string = QDateTimeEdit_DisplayFormat(this.h)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTimeEdit) SetDisplayFormat(format string) {
	format_ms := struct_miqt_string{}
	format_ms.data = CString(format)
	format_ms.len = size_t(len(format))
	defer free(unsafe.Pointer(format_ms.data))
	QDateTimeEdit_SetDisplayFormat(this.h, format_ms)
}

func (this *QDateTimeEdit) CalendarPopup() bool {
	return (bool)(QDateTimeEdit_CalendarPopup(this.h))
}

func (this *QDateTimeEdit) SetCalendarPopup(enable bool) {
	QDateTimeEdit_SetCalendarPopup(this.h, (bool)(enable))
}

func (this *QDateTimeEdit) TimeSpec() TimeSpec {
	return (TimeSpec)(QDateTimeEdit_TimeSpec(this.h))
}

func (this *QDateTimeEdit) SetTimeSpec(spec TimeSpec) {
	QDateTimeEdit_SetTimeSpec(this.h, (int)(spec))
}

func (this *QDateTimeEdit) TimeZone() *QTimeZone {
	_goptr := newQTimeZone(QDateTimeEdit_TimeZone(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) SetTimeZone(zone *QTimeZone) {
	QDateTimeEdit_SetTimeZone(this.h, zone.cPointer())
}

func (this *QDateTimeEdit) SizeHint() *QSize {
	_goptr := newQSize(QDateTimeEdit_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QDateTimeEdit) Clear() {
	QDateTimeEdit_Clear(this.h)
}

func (this *QDateTimeEdit) StepBy(steps int) {
	QDateTimeEdit_StepBy(this.h, (int)(steps))
}

func (this *QDateTimeEdit) Event(event *QEvent) bool {
	return (bool)(QDateTimeEdit_Event(this.h, event.cPointer()))
}

func (this *QDateTimeEdit) DateTimeChanged(dateTime *QDateTime) {
	QDateTimeEdit_DateTimeChanged(this.h, dateTime.cPointer())
}

func (this *QDateTimeEdit) OnDateTimeChanged(slot func(dateTime *QDateTime)) {
	QDateTimeEdit_connect_DateTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_DateTimeChanged
func miqt_exec_callback_QDateTimeEdit_DateTimeChanged(cb intptr_t, dateTime *QDateTime) {
	gofunc, ok := cgo.Handle(cb).Value().(func(dateTime *QDateTime))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDateTime(dateTime)

	gofunc(slotval1)
}

func (this *QDateTimeEdit) TimeChanged(time QTime) {
	QDateTimeEdit_TimeChanged(this.h, time.cPointer())
}

func (this *QDateTimeEdit) OnTimeChanged(slot func(time QTime)) {
	QDateTimeEdit_connect_TimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_TimeChanged
func miqt_exec_callback_QDateTimeEdit_TimeChanged(cb intptr_t, time *QTime) {
	gofunc, ok := cgo.Handle(cb).Value().(func(time QTime))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	time_goptr := newQTime(time)
	time_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *time_goptr

	gofunc(slotval1)
}

func (this *QDateTimeEdit) DateChanged(date QDate) {
	QDateTimeEdit_DateChanged(this.h, date.cPointer())
}

func (this *QDateTimeEdit) OnDateChanged(slot func(date QDate)) {
	QDateTimeEdit_connect_DateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_DateChanged
func miqt_exec_callback_QDateTimeEdit_DateChanged(cb intptr_t, date *QDate) {
	gofunc, ok := cgo.Handle(cb).Value().(func(date QDate))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	date_goptr := newQDate(date)
	date_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *date_goptr

	gofunc(slotval1)
}

func (this *QDateTimeEdit) SetDateTime(dateTime *QDateTime) {
	QDateTimeEdit_SetDateTime(this.h, dateTime.cPointer())
}

func (this *QDateTimeEdit) SetDate(date QDate) {
	QDateTimeEdit_SetDate(this.h, date.cPointer())
}

func (this *QDateTimeEdit) SetTime(time QTime) {
	QDateTimeEdit_SetTime(this.h, time.cPointer())
}

func QDateTimeEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDateTimeEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDateTimeEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDateTimeEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateTimeEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDateTimeEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDateTimeEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_MetaObject
func miqt_exec_callback_QDateTimeEdit_MetaObject(self QDateTimeEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDateTimeEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDateTimeEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDateTimeEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_Metacast
func miqt_exec_callback_QDateTimeEdit_Metacast(self QDateTimeEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QTimeEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQTimeEdit constructs a new QTimeEdit object.
func NewQTimeEdit(parent *QWidget) *QTimeEdit {
	ret := newQTimeEdit(QTimeEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTimeEdit2 constructs a new QTimeEdit object.
func NewQTimeEdit2() *QTimeEdit {
	ret := newQTimeEdit(QTimeEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQTimeEdit3 constructs a new QTimeEdit object.
func NewQTimeEdit3(time QTime) *QTimeEdit {
	ret := newQTimeEdit(QTimeEdit_new3(time.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQTimeEdit4 constructs a new QTimeEdit object.
func NewQTimeEdit4(time QTime, parent *QWidget) *QTimeEdit {
	ret := newQTimeEdit(QTimeEdit_new4(time.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QTimeEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QTimeEdit_MetaObject(this.h))
}

func (this *QTimeEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QTimeEdit_Metacast(this.h, param1_Cstring))
}

func QTimeEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QTimeEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeEdit) UserTimeChanged(time QTime) {
	QTimeEdit_UserTimeChanged(this.h, time.cPointer())
}

func (this *QTimeEdit) OnUserTimeChanged(slot func(time QTime)) {
	QTimeEdit_connect_UserTimeChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_UserTimeChanged
func miqt_exec_callback_QTimeEdit_UserTimeChanged(cb intptr_t, time *QTime) {
	gofunc, ok := cgo.Handle(cb).Value().(func(time QTime))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	time_goptr := newQTime(time)
	time_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *time_goptr

	gofunc(slotval1)
}

func QTimeEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimeEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QTimeEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QTimeEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QTimeEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QTimeEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QTimeEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_MetaObject
func miqt_exec_callback_QTimeEdit_MetaObject(self QTimeEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QTimeEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QTimeEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QTimeEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_Metacast
func miqt_exec_callback_QTimeEdit_Metacast(self QTimeEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}

type QDateEdit struct {
	h          uintptr
	isSubclass bool
}

// NewQDateEdit constructs a new QDateEdit object.
func NewQDateEdit(parent *QWidget) *QDateEdit {
	ret := newQDateEdit(QDateEdit_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateEdit2 constructs a new QDateEdit object.
func NewQDateEdit2() *QDateEdit {
	ret := newQDateEdit(QDateEdit_new2())
	ret.isSubclass = true
	return ret
}

// NewQDateEdit3 constructs a new QDateEdit object.
func NewQDateEdit3(date QDate) *QDateEdit {
	ret := newQDateEdit(QDateEdit_new3(date.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQDateEdit4 constructs a new QDateEdit object.
func NewQDateEdit4(date QDate, parent *QWidget) *QDateEdit {
	ret := newQDateEdit(QDateEdit_new4(date.cPointer(), parent.cPointer()))
	ret.isSubclass = true
	return ret
}

func (this *QDateEdit) MetaObject() *QMetaObject {
	return newQMetaObject(QDateEdit_MetaObject(this.h))
}

func (this *QDateEdit) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QDateEdit_Metacast(this.h, param1_Cstring))
}

func QDateEdit_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QDateEdit_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateEdit) UserDateChanged(date QDate) {
	QDateEdit_UserDateChanged(this.h, date.cPointer())
}

func (this *QDateEdit) OnUserDateChanged(slot func(date QDate)) {
	QDateEdit_connect_UserDateChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_UserDateChanged
func miqt_exec_callback_QDateEdit_UserDateChanged(cb intptr_t, date *QDate) {
	gofunc, ok := cgo.Handle(cb).Value().(func(date QDate))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	date_goptr := newQDate(date)
	date_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval1 := *date_goptr

	gofunc(slotval1)
}

func QDateEdit_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDateEdit_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QDateEdit_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QDateEdit_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QDateEdit) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QDateEdit_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QDateEdit) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_MetaObject
func miqt_exec_callback_QDateEdit_MetaObject(self QDateEdit, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QDateEdit) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QDateEdit_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QDateEdit) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_Metacast
func miqt_exec_callback_QDateEdit_Metacast(self QDateEdit, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
