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

func (this *QDateTimeEdit) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QDateTimeEdit_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDateTimeEdit) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_SizeHint
func miqt_exec_callback_QDateTimeEdit_SizeHint(self QDateTimeEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QDateTimeEdit) callVirtualBase_Clear() {

	QDateTimeEdit_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QDateTimeEdit) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_Clear
func miqt_exec_callback_QDateTimeEdit_Clear(self QDateTimeEdit, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_Clear)

}

func (this *QDateTimeEdit) callVirtualBase_StepBy(steps int) {

	QDateTimeEdit_virtualbase_StepBy(unsafe.Pointer(this.h), (int)(steps))

}
func (this *QDateTimeEdit) OnStepBy(slot func(super func(steps int), steps int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_StepBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_StepBy
func miqt_exec_callback_QDateTimeEdit_StepBy(self QDateTimeEdit, cb intptr_t, steps int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(steps int), steps int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(steps)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_StepBy, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QDateTimeEdit_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QDateTimeEdit) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_Event
func miqt_exec_callback_QDateTimeEdit_Event(self QDateTimeEdit, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDateTimeEdit) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QDateTimeEdit_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_KeyPressEvent
func miqt_exec_callback_QDateTimeEdit_KeyPressEvent(self QDateTimeEdit, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QDateTimeEdit_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_WheelEvent
func miqt_exec_callback_QDateTimeEdit_WheelEvent(self QDateTimeEdit, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QDateTimeEdit_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_FocusInEvent
func miqt_exec_callback_QDateTimeEdit_FocusInEvent(self QDateTimeEdit, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QDateTimeEdit_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QDateTimeEdit) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_FocusNextPrevChild
func miqt_exec_callback_QDateTimeEdit_FocusNextPrevChild(self QDateTimeEdit, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDateTimeEdit) callVirtualBase_Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	return (QValidator__State)(QDateTimeEdit_virtualbase_Validate(unsafe.Pointer(this.h), input_ms, (*int)(unsafe.Pointer(pos))))

}
func (this *QDateTimeEdit) OnValidate(slot func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_Validate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_Validate
func miqt_exec_callback_QDateTimeEdit_Validate(self QDateTimeEdit, cb intptr_t, input struct_miqt_string, pos *int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret
	slotval2 := (*int)(unsafe.Pointer(pos))

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_Validate, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QDateTimeEdit) callVirtualBase_Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	QDateTimeEdit_virtualbase_Fixup(unsafe.Pointer(this.h), input_ms)

}
func (this *QDateTimeEdit) OnFixup(slot func(super func(input string), input string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_Fixup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_Fixup
func miqt_exec_callback_QDateTimeEdit_Fixup(self QDateTimeEdit, cb intptr_t, input struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string), input string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_Fixup, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_DateTimeFromText(text string) *QDateTime {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	_goptr := newQDateTime(QDateTimeEdit_virtualbase_DateTimeFromText(unsafe.Pointer(this.h), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDateTimeEdit) OnDateTimeFromText(slot func(super func(text string) *QDateTime, text string) *QDateTime) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_DateTimeFromText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_DateTimeFromText
func miqt_exec_callback_QDateTimeEdit_DateTimeFromText(self QDateTimeEdit, cb intptr_t, text struct_miqt_string) *QDateTime {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(text string) *QDateTime, text string) *QDateTime)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_DateTimeFromText, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDateTimeEdit) callVirtualBase_TextFromDateTime(dt *QDateTime) string {

	var _ms struct_miqt_string = QDateTimeEdit_virtualbase_TextFromDateTime(unsafe.Pointer(this.h), dt.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QDateTimeEdit) OnTextFromDateTime(slot func(super func(dt *QDateTime) string, dt *QDateTime) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_TextFromDateTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_TextFromDateTime
func miqt_exec_callback_QDateTimeEdit_TextFromDateTime(self QDateTimeEdit, cb intptr_t, dt *QDateTime) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dt *QDateTime) string, dt *QDateTime) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDateTime(dt)

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_TextFromDateTime, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QDateTimeEdit) callVirtualBase_StepEnabled() StepEnabled {

	xxxxxxxxx
}
func (this *QDateTimeEdit) OnStepEnabled(slot func(super func() StepEnabled) StepEnabled) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_StepEnabled(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_StepEnabled
func miqt_exec_callback_QDateTimeEdit_StepEnabled(self QDateTimeEdit, cb intptr_t) StepEnabled {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() StepEnabled) StepEnabled)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_StepEnabled)

	return virtualReturn

}

func (this *QDateTimeEdit) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QDateTimeEdit_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_MousePressEvent
func miqt_exec_callback_QDateTimeEdit_MousePressEvent(self QDateTimeEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QDateTimeEdit_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_PaintEvent
func miqt_exec_callback_QDateTimeEdit_PaintEvent(self QDateTimeEdit, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_InitStyleOption(option *QStyleOptionSpinBox) {

	QDateTimeEdit_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QDateTimeEdit) OnInitStyleOption(slot func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_InitStyleOption
func miqt_exec_callback_QDateTimeEdit_InitStyleOption(self QDateTimeEdit, cb intptr_t, option *QStyleOptionSpinBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSpinBox(option)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_InitStyleOption, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QDateTimeEdit_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDateTimeEdit) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_MinimumSizeHint
func miqt_exec_callback_QDateTimeEdit_MinimumSizeHint(self QDateTimeEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QDateTimeEdit) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QDateTimeEdit_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDateTimeEdit) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_InputMethodQuery
func miqt_exec_callback_QDateTimeEdit_InputMethodQuery(self QDateTimeEdit, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QDateTimeEdit{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDateTimeEdit) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QDateTimeEdit_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_ResizeEvent
func miqt_exec_callback_QDateTimeEdit_ResizeEvent(self QDateTimeEdit, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QDateTimeEdit_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_KeyReleaseEvent
func miqt_exec_callback_QDateTimeEdit_KeyReleaseEvent(self QDateTimeEdit, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QDateTimeEdit_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_FocusOutEvent
func miqt_exec_callback_QDateTimeEdit_FocusOutEvent(self QDateTimeEdit, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QDateTimeEdit_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_ContextMenuEvent
func miqt_exec_callback_QDateTimeEdit_ContextMenuEvent(self QDateTimeEdit, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_ChangeEvent(event *QEvent) {

	QDateTimeEdit_virtualbase_ChangeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnChangeEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_ChangeEvent
func miqt_exec_callback_QDateTimeEdit_ChangeEvent(self QDateTimeEdit, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QDateTimeEdit_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_CloseEvent
func miqt_exec_callback_QDateTimeEdit_CloseEvent(self QDateTimeEdit, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_HideEvent(event *QHideEvent) {

	QDateTimeEdit_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_HideEvent
func miqt_exec_callback_QDateTimeEdit_HideEvent(self QDateTimeEdit, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QDateTimeEdit_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_MouseReleaseEvent
func miqt_exec_callback_QDateTimeEdit_MouseReleaseEvent(self QDateTimeEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QDateTimeEdit_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_MouseMoveEvent
func miqt_exec_callback_QDateTimeEdit_MouseMoveEvent(self QDateTimeEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_TimerEvent(event *QTimerEvent) {

	QDateTimeEdit_virtualbase_TimerEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnTimerEvent(slot func(super func(event *QTimerEvent), event *QTimerEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_TimerEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_TimerEvent
func miqt_exec_callback_QDateTimeEdit_TimerEvent(self QDateTimeEdit, cb intptr_t, event *QTimerEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTimerEvent), event *QTimerEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTimerEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_TimerEvent, slotval1)

}

func (this *QDateTimeEdit) callVirtualBase_ShowEvent(event *QShowEvent) {

	QDateTimeEdit_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateTimeEdit) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateTimeEdit_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateTimeEdit_ShowEvent
func miqt_exec_callback_QDateTimeEdit_ShowEvent(self QDateTimeEdit, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QDateTimeEdit{h: self}).callVirtualBase_ShowEvent, slotval1)

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

func (this *QTimeEdit) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QTimeEdit_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTimeEdit) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_SizeHint
func miqt_exec_callback_QTimeEdit_SizeHint(self QTimeEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QTimeEdit) callVirtualBase_Clear() {

	QTimeEdit_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QTimeEdit) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_Clear
func miqt_exec_callback_QTimeEdit_Clear(self QTimeEdit, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QTimeEdit{h: self}).callVirtualBase_Clear)

}

func (this *QTimeEdit) callVirtualBase_StepBy(steps int) {

	QTimeEdit_virtualbase_StepBy(unsafe.Pointer(this.h), (int)(steps))

}
func (this *QTimeEdit) OnStepBy(slot func(super func(steps int), steps int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_StepBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_StepBy
func miqt_exec_callback_QTimeEdit_StepBy(self QTimeEdit, cb intptr_t, steps int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(steps int), steps int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(steps)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_StepBy, slotval1)

}

func (this *QTimeEdit) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QTimeEdit_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QTimeEdit) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_Event
func miqt_exec_callback_QTimeEdit_Event(self QTimeEdit, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTimeEdit) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QTimeEdit_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeEdit) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_KeyPressEvent
func miqt_exec_callback_QTimeEdit_KeyPressEvent(self QTimeEdit, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QTimeEdit) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QTimeEdit_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeEdit) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_WheelEvent
func miqt_exec_callback_QTimeEdit_WheelEvent(self QTimeEdit, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QTimeEdit) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QTimeEdit_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeEdit) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_FocusInEvent
func miqt_exec_callback_QTimeEdit_FocusInEvent(self QTimeEdit, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QTimeEdit) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QTimeEdit_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QTimeEdit) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_FocusNextPrevChild
func miqt_exec_callback_QTimeEdit_FocusNextPrevChild(self QTimeEdit, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QTimeEdit) callVirtualBase_Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	return (QValidator__State)(QTimeEdit_virtualbase_Validate(unsafe.Pointer(this.h), input_ms, (*int)(unsafe.Pointer(pos))))

}
func (this *QTimeEdit) OnValidate(slot func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_Validate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_Validate
func miqt_exec_callback_QTimeEdit_Validate(self QTimeEdit, cb intptr_t, input struct_miqt_string, pos *int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret
	slotval2 := (*int)(unsafe.Pointer(pos))

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_Validate, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QTimeEdit) callVirtualBase_Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	QTimeEdit_virtualbase_Fixup(unsafe.Pointer(this.h), input_ms)

}
func (this *QTimeEdit) OnFixup(slot func(super func(input string), input string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_Fixup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_Fixup
func miqt_exec_callback_QTimeEdit_Fixup(self QTimeEdit, cb intptr_t, input struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string), input string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret

	gofunc((&QTimeEdit{h: self}).callVirtualBase_Fixup, slotval1)

}

func (this *QTimeEdit) callVirtualBase_DateTimeFromText(text string) *QDateTime {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	_goptr := newQDateTime(QTimeEdit_virtualbase_DateTimeFromText(unsafe.Pointer(this.h), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QTimeEdit) OnDateTimeFromText(slot func(super func(text string) *QDateTime, text string) *QDateTime) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_DateTimeFromText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_DateTimeFromText
func miqt_exec_callback_QTimeEdit_DateTimeFromText(self QTimeEdit, cb intptr_t, text struct_miqt_string) *QDateTime {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(text string) *QDateTime, text string) *QDateTime)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_DateTimeFromText, slotval1)

	return virtualReturn.cPointer()

}

func (this *QTimeEdit) callVirtualBase_TextFromDateTime(dt *QDateTime) string {

	var _ms struct_miqt_string = QTimeEdit_virtualbase_TextFromDateTime(unsafe.Pointer(this.h), dt.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QTimeEdit) OnTextFromDateTime(slot func(super func(dt *QDateTime) string, dt *QDateTime) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_TextFromDateTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_TextFromDateTime
func miqt_exec_callback_QTimeEdit_TextFromDateTime(self QTimeEdit, cb intptr_t, dt *QDateTime) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dt *QDateTime) string, dt *QDateTime) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDateTime(dt)

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_TextFromDateTime, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QTimeEdit) callVirtualBase_StepEnabled() StepEnabled {

	xxxxxxxxx
}
func (this *QTimeEdit) OnStepEnabled(slot func(super func() StepEnabled) StepEnabled) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_StepEnabled(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_StepEnabled
func miqt_exec_callback_QTimeEdit_StepEnabled(self QTimeEdit, cb intptr_t) StepEnabled {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() StepEnabled) StepEnabled)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QTimeEdit{h: self}).callVirtualBase_StepEnabled)

	return virtualReturn

}

func (this *QTimeEdit) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QTimeEdit_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeEdit) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_MousePressEvent
func miqt_exec_callback_QTimeEdit_MousePressEvent(self QTimeEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QTimeEdit) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QTimeEdit_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QTimeEdit) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_PaintEvent
func miqt_exec_callback_QTimeEdit_PaintEvent(self QTimeEdit, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QTimeEdit) callVirtualBase_InitStyleOption(option *QStyleOptionSpinBox) {

	QTimeEdit_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QTimeEdit) OnInitStyleOption(slot func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QTimeEdit_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QTimeEdit_InitStyleOption
func miqt_exec_callback_QTimeEdit_InitStyleOption(self QTimeEdit, cb intptr_t, option *QStyleOptionSpinBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSpinBox(option)

	gofunc((&QTimeEdit{h: self}).callVirtualBase_InitStyleOption, slotval1)

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

func (this *QDateEdit) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QDateEdit_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDateEdit) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_SizeHint
func miqt_exec_callback_QDateEdit_SizeHint(self QDateEdit, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QDateEdit) callVirtualBase_Clear() {

	QDateEdit_virtualbase_Clear(unsafe.Pointer(this.h))

}
func (this *QDateEdit) OnClear(slot func(super func())) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_Clear(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_Clear
func miqt_exec_callback_QDateEdit_Clear(self QDateEdit, cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func()))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc((&QDateEdit{h: self}).callVirtualBase_Clear)

}

func (this *QDateEdit) callVirtualBase_StepBy(steps int) {

	QDateEdit_virtualbase_StepBy(unsafe.Pointer(this.h), (int)(steps))

}
func (this *QDateEdit) OnStepBy(slot func(super func(steps int), steps int)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_StepBy(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_StepBy
func miqt_exec_callback_QDateEdit_StepBy(self QDateEdit, cb intptr_t, steps int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(steps int), steps int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(steps)

	gofunc((&QDateEdit{h: self}).callVirtualBase_StepBy, slotval1)

}

func (this *QDateEdit) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QDateEdit_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QDateEdit) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_Event
func miqt_exec_callback_QDateEdit_Event(self QDateEdit, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDateEdit) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QDateEdit_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateEdit) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_KeyPressEvent
func miqt_exec_callback_QDateEdit_KeyPressEvent(self QDateEdit, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QDateEdit{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QDateEdit) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QDateEdit_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateEdit) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_WheelEvent
func miqt_exec_callback_QDateEdit_WheelEvent(self QDateEdit, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QDateEdit{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QDateEdit) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QDateEdit_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateEdit) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_FocusInEvent
func miqt_exec_callback_QDateEdit_FocusInEvent(self QDateEdit, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QDateEdit{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QDateEdit) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QDateEdit_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QDateEdit) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_FocusNextPrevChild
func miqt_exec_callback_QDateEdit_FocusNextPrevChild(self QDateEdit, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}

func (this *QDateEdit) callVirtualBase_Validate(input string, pos *int) QValidator__State {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	return (QValidator__State)(QDateEdit_virtualbase_Validate(unsafe.Pointer(this.h), input_ms, (*int)(unsafe.Pointer(pos))))

}
func (this *QDateEdit) OnValidate(slot func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_Validate(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_Validate
func miqt_exec_callback_QDateEdit_Validate(self QDateEdit, cb intptr_t, input struct_miqt_string, pos *int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string, pos *int) QValidator__State, input string, pos *int) QValidator__State)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret
	slotval2 := (*int)(unsafe.Pointer(pos))

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_Validate, slotval1, slotval2)

	return (int)(virtualReturn)

}

func (this *QDateEdit) callVirtualBase_Fixup(input string) {
	input_ms := struct_miqt_string{}
	input_ms.data = CString(input)
	input_ms.len = size_t(len(input))
	defer free(unsafe.Pointer(input_ms.data))

	QDateEdit_virtualbase_Fixup(unsafe.Pointer(this.h), input_ms)

}
func (this *QDateEdit) OnFixup(slot func(super func(input string), input string)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_Fixup(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_Fixup
func miqt_exec_callback_QDateEdit_Fixup(self QDateEdit, cb intptr_t, input struct_miqt_string) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(input string), input string))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var input_ms struct_miqt_string = input
	input_ret := GoStringN(input_ms.data, int(int64(input_ms.len)))
	free(unsafe.Pointer(input_ms.data))
	slotval1 := input_ret

	gofunc((&QDateEdit{h: self}).callVirtualBase_Fixup, slotval1)

}

func (this *QDateEdit) callVirtualBase_DateTimeFromText(text string) *QDateTime {
	text_ms := struct_miqt_string{}
	text_ms.data = CString(text)
	text_ms.len = size_t(len(text))
	defer free(unsafe.Pointer(text_ms.data))

	_goptr := newQDateTime(QDateEdit_virtualbase_DateTimeFromText(unsafe.Pointer(this.h), text_ms))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QDateEdit) OnDateTimeFromText(slot func(super func(text string) *QDateTime, text string) *QDateTime) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_DateTimeFromText(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_DateTimeFromText
func miqt_exec_callback_QDateEdit_DateTimeFromText(self QDateEdit, cb intptr_t, text struct_miqt_string) *QDateTime {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(text string) *QDateTime, text string) *QDateTime)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var text_ms struct_miqt_string = text
	text_ret := GoStringN(text_ms.data, int(int64(text_ms.len)))
	free(unsafe.Pointer(text_ms.data))
	slotval1 := text_ret

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_DateTimeFromText, slotval1)

	return virtualReturn.cPointer()

}

func (this *QDateEdit) callVirtualBase_TextFromDateTime(dt *QDateTime) string {

	var _ms struct_miqt_string = QDateEdit_virtualbase_TextFromDateTime(unsafe.Pointer(this.h), dt.cPointer())
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}
func (this *QDateEdit) OnTextFromDateTime(slot func(super func(dt *QDateTime) string, dt *QDateTime) string) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_TextFromDateTime(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_TextFromDateTime
func miqt_exec_callback_QDateEdit_TextFromDateTime(self QDateEdit, cb intptr_t, dt *QDateTime) struct_miqt_string {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(dt *QDateTime) string, dt *QDateTime) string)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDateTime(dt)

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_TextFromDateTime, slotval1)
	virtualReturn_ms := struct_miqt_string{}
	virtualReturn_ms.data = CString(virtualReturn)
	virtualReturn_ms.len = size_t(len(virtualReturn))
	defer free(unsafe.Pointer(virtualReturn_ms.data))

	return virtualReturn_ms

}

func (this *QDateEdit) callVirtualBase_StepEnabled() StepEnabled {

	xxxxxxxxx
}
func (this *QDateEdit) OnStepEnabled(slot func(super func() StepEnabled) StepEnabled) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_StepEnabled(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_StepEnabled
func miqt_exec_callback_QDateEdit_StepEnabled(self QDateEdit, cb intptr_t) StepEnabled {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() StepEnabled) StepEnabled)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QDateEdit{h: self}).callVirtualBase_StepEnabled)

	return virtualReturn

}

func (this *QDateEdit) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QDateEdit_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateEdit) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_MousePressEvent
func miqt_exec_callback_QDateEdit_MousePressEvent(self QDateEdit, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QDateEdit{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QDateEdit) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QDateEdit_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QDateEdit) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_PaintEvent
func miqt_exec_callback_QDateEdit_PaintEvent(self QDateEdit, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QDateEdit{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QDateEdit) callVirtualBase_InitStyleOption(option *QStyleOptionSpinBox) {

	QDateEdit_virtualbase_InitStyleOption(unsafe.Pointer(this.h), option.cPointer())

}
func (this *QDateEdit) OnInitStyleOption(slot func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QDateEdit_override_virtual_InitStyleOption(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QDateEdit_InitStyleOption
func miqt_exec_callback_QDateEdit_InitStyleOption(self QDateEdit, cb intptr_t, option *QStyleOptionSpinBox) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(option *QStyleOptionSpinBox), option *QStyleOptionSpinBox))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQStyleOptionSpinBox(option)

	gofunc((&QDateEdit{h: self}).callVirtualBase_InitStyleOption, slotval1)

}
