package qt

import (
	"unsafe"
)

type QCalendarWidget__HorizontalHeaderFormat int

const (
	QCalendarWidget__NoHorizontalHeader   QCalendarWidget__HorizontalHeaderFormat = 0
	QCalendarWidget__SingleLetterDayNames QCalendarWidget__HorizontalHeaderFormat = 1
	QCalendarWidget__ShortDayNames        QCalendarWidget__HorizontalHeaderFormat = 2
	QCalendarWidget__LongDayNames         QCalendarWidget__HorizontalHeaderFormat = 3
)

type QCalendarWidget__VerticalHeaderFormat int

const (
	QCalendarWidget__NoVerticalHeader QCalendarWidget__VerticalHeaderFormat = 0
	QCalendarWidget__ISOWeekNumbers   QCalendarWidget__VerticalHeaderFormat = 1
)

type QCalendarWidget__SelectionMode int

const (
	QCalendarWidget__NoSelection     QCalendarWidget__SelectionMode = 0
	QCalendarWidget__SingleSelection QCalendarWidget__SelectionMode = 1
)

type QCalendarWidget struct {
	h          uintptr
	isSubclass bool
}

// NewQCalendarWidget constructs a new QCalendarWidget object.
func NewQCalendarWidget(parent *QWidget) *QCalendarWidget {
	g := newQCalendarWidget(QCalendarWidget_new(parent.cPointer()))
	g.isSubclass = true
	return g
}

// NewQCalendarWidget2 constructs a new QCalendarWidget object.
func NewQCalendarWidget2() *QCalendarWidget {
	g := newQCalendarWidget(QCalendarWidget_new2())
	g.isSubclass = true
	return g
}

func (this *QCalendarWidget) MetaObject() *QMetaObject {
	return newQMetaObject(QCalendarWidget_MetaObject(this.h))
}

func (this *QCalendarWidget) Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))
	return (unsafe.Pointer)(QCalendarWidget_Metacast(this.h, param1_Cstring))
}

func QCalendarWidget_Tr(s string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	var _ms struct_miqt_string = QCalendarWidget_Tr(s_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendarWidget) SizeHint() *QSize {
	_goptr := newQSize(QCalendarWidget_SizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) MinimumSizeHint() *QSize {
	_goptr := newQSize(QCalendarWidget_MinimumSizeHint(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SelectedDate() *QDate {
	_goptr := newQDate(QCalendarWidget_SelectedDate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) YearShown() int {
	return (int)(QCalendarWidget_YearShown(this.h))
}

func (this *QCalendarWidget) MonthShown() int {
	return (int)(QCalendarWidget_MonthShown(this.h))
}

func (this *QCalendarWidget) MinimumDate() *QDate {
	_goptr := newQDate(QCalendarWidget_MinimumDate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SetMinimumDate(date QDate) {
	QCalendarWidget_SetMinimumDate(this.h, date.cPointer())
}

func (this *QCalendarWidget) ClearMinimumDate() {
	QCalendarWidget_ClearMinimumDate(this.h)
}

func (this *QCalendarWidget) MaximumDate() *QDate {
	_goptr := newQDate(QCalendarWidget_MaximumDate(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SetMaximumDate(date QDate) {
	QCalendarWidget_SetMaximumDate(this.h, date.cPointer())
}

func (this *QCalendarWidget) ClearMaximumDate() {
	QCalendarWidget_ClearMaximumDate(this.h)
}

func (this *QCalendarWidget) FirstDayOfWeek() DayOfWeek {
	return (DayOfWeek)(QCalendarWidget_FirstDayOfWeek(this.h))
}

func (this *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek DayOfWeek) {
	QCalendarWidget_SetFirstDayOfWeek(this.h, (int)(dayOfWeek))
}

func (this *QCalendarWidget) IsNavigationBarVisible() bool {
	return (bool)(QCalendarWidget_IsNavigationBarVisible(this.h))
}

func (this *QCalendarWidget) IsGridVisible() bool {
	return (bool)(QCalendarWidget_IsGridVisible(this.h))
}

func (this *QCalendarWidget) Calendar() *QCalendar {
	_goptr := newQCalendar(QCalendarWidget_Calendar(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SetCalendar(calendar QCalendar) {
	QCalendarWidget_SetCalendar(this.h, calendar.cPointer())
}

func (this *QCalendarWidget) SelectionMode() SelectionMode {
	xxxxxxxxx
}

func (this *QCalendarWidget) SetSelectionMode(mode SelectionMode) {
	QCalendarWidget_SetSelectionMode(this.h, mode)
}

func (this *QCalendarWidget) HorizontalHeaderFormat() HorizontalHeaderFormat {
	xxxxxxxxx
}

func (this *QCalendarWidget) SetHorizontalHeaderFormat(format HorizontalHeaderFormat) {
	QCalendarWidget_SetHorizontalHeaderFormat(this.h, format)
}

func (this *QCalendarWidget) VerticalHeaderFormat() VerticalHeaderFormat {
	xxxxxxxxx
}

func (this *QCalendarWidget) SetVerticalHeaderFormat(format VerticalHeaderFormat) {
	QCalendarWidget_SetVerticalHeaderFormat(this.h, format)
}

func (this *QCalendarWidget) HeaderTextFormat() *QTextCharFormat {
	_goptr := newQTextCharFormat(QCalendarWidget_HeaderTextFormat(this.h))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SetHeaderTextFormat(format *QTextCharFormat) {
	QCalendarWidget_SetHeaderTextFormat(this.h, format.cPointer())
}

func (this *QCalendarWidget) WeekdayTextFormat(dayOfWeek DayOfWeek) *QTextCharFormat {
	_goptr := newQTextCharFormat(QCalendarWidget_WeekdayTextFormat(this.h, (int)(dayOfWeek)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek DayOfWeek, format *QTextCharFormat) {
	QCalendarWidget_SetWeekdayTextFormat(this.h, (int)(dayOfWeek), format.cPointer())
}

func (this *QCalendarWidget) DateTextFormat() map[QDate]QTextCharFormat {
	var _mm struct_miqt_map = QCalendarWidget_DateTextFormat(this.h)
	_ret := make(map[QDate]QTextCharFormat, int(_mm.len))
	_Keys := (*[0xffff]*QDate)(unsafe.Pointer(_mm.keys))
	_Values := (*[0xffff]*QTextCharFormat)(unsafe.Pointer(_mm.values))
	for i := 0; i < int(_mm.len); i++ {
		_mapkey_goptr := newQDate(_Keys[i])
		_mapkey_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Key := *_mapkey_goptr

		_mapval_goptr := newQTextCharFormat(_Values[i])
		_mapval_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
		_entry_Value := *_mapval_goptr

		_ret[_entry_Key] = _entry_Value
	}
	return _ret
}

func (this *QCalendarWidget) DateTextFormatWithDate(date QDate) *QTextCharFormat {
	_goptr := newQTextCharFormat(QCalendarWidget_DateTextFormatWithDate(this.h, date.cPointer()))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr
}

func (this *QCalendarWidget) SetDateTextFormat(date QDate, format *QTextCharFormat) {
	QCalendarWidget_SetDateTextFormat(this.h, date.cPointer(), format.cPointer())
}

func (this *QCalendarWidget) IsDateEditEnabled() bool {
	return (bool)(QCalendarWidget_IsDateEditEnabled(this.h))
}

func (this *QCalendarWidget) SetDateEditEnabled(enable bool) {
	QCalendarWidget_SetDateEditEnabled(this.h, (bool)(enable))
}

func (this *QCalendarWidget) DateEditAcceptDelay() int {
	return (int)(QCalendarWidget_DateEditAcceptDelay(this.h))
}

func (this *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	QCalendarWidget_SetDateEditAcceptDelay(this.h, (int)(delay))
}

func (this *QCalendarWidget) SetSelectedDate(date QDate) {
	QCalendarWidget_SetSelectedDate(this.h, date.cPointer())
}

func (this *QCalendarWidget) SetDateRange(min QDate, max QDate) {
	QCalendarWidget_SetDateRange(this.h, min.cPointer(), max.cPointer())
}

func (this *QCalendarWidget) SetCurrentPage(year int, month int) {
	QCalendarWidget_SetCurrentPage(this.h, (int)(year), (int)(month))
}

func (this *QCalendarWidget) SetGridVisible(show bool) {
	QCalendarWidget_SetGridVisible(this.h, (bool)(show))
}

func (this *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	QCalendarWidget_SetNavigationBarVisible(this.h, (bool)(visible))
}

func (this *QCalendarWidget) ShowNextMonth() {
	QCalendarWidget_ShowNextMonth(this.h)
}

func (this *QCalendarWidget) ShowPreviousMonth() {
	QCalendarWidget_ShowPreviousMonth(this.h)
}

func (this *QCalendarWidget) ShowNextYear() {
	QCalendarWidget_ShowNextYear(this.h)
}

func (this *QCalendarWidget) ShowPreviousYear() {
	QCalendarWidget_ShowPreviousYear(this.h)
}

func (this *QCalendarWidget) ShowSelectedDate() {
	QCalendarWidget_ShowSelectedDate(this.h)
}

func (this *QCalendarWidget) ShowToday() {
	QCalendarWidget_ShowToday(this.h)
}

func (this *QCalendarWidget) SelectionChanged() {
	QCalendarWidget_SelectionChanged(this.h)
}

func (this *QCalendarWidget) OnSelectionChanged(slot func()) {
	QCalendarWidget_connect_SelectionChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_SelectionChanged
func miqt_exec_callback_QCalendarWidget_SelectionChanged(cb intptr_t) {
	gofunc, ok := cgo.Handle(cb).Value().(func())
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	gofunc()
}

func (this *QCalendarWidget) Clicked(date QDate) {
	QCalendarWidget_Clicked(this.h, date.cPointer())
}

func (this *QCalendarWidget) OnClicked(slot func(date QDate)) {
	QCalendarWidget_connect_Clicked(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_Clicked
func miqt_exec_callback_QCalendarWidget_Clicked(cb intptr_t, date *QDate) {
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

func (this *QCalendarWidget) Activated(date QDate) {
	QCalendarWidget_Activated(this.h, date.cPointer())
}

func (this *QCalendarWidget) OnActivated(slot func(date QDate)) {
	QCalendarWidget_connect_Activated(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_Activated
func miqt_exec_callback_QCalendarWidget_Activated(cb intptr_t, date *QDate) {
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

func (this *QCalendarWidget) CurrentPageChanged(year int, month int) {
	QCalendarWidget_CurrentPageChanged(this.h, (int)(year), (int)(month))
}

func (this *QCalendarWidget) OnCurrentPageChanged(slot func(year int, month int)) {
	QCalendarWidget_connect_CurrentPageChanged(this.h, intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_CurrentPageChanged
func miqt_exec_callback_QCalendarWidget_CurrentPageChanged(cb intptr_t, year int, month int) {
	gofunc, ok := cgo.Handle(cb).Value().(func(year int, month int))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(year)

	slotval2 := (int)(month)

	gofunc(slotval1, slotval2)
}

func QCalendarWidget_Tr2(s string, c string) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCalendarWidget_Tr2(s_Cstring, c_Cstring)
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func QCalendarWidget_Tr3(s string, c string, n int) string {
	s_Cstring := CString(s)
	defer free(unsafe.Pointer(s_Cstring))
	c_Cstring := CString(c)
	defer free(unsafe.Pointer(c_Cstring))
	var _ms struct_miqt_string = QCalendarWidget_Tr3(s_Cstring, c_Cstring, (int)(n))
	_ret := GoStringN(_ms.data, int(int64(_ms.len)))
	free(unsafe.Pointer(_ms.data))
	return _ret
}

func (this *QCalendarWidget) callVirtualBase_MetaObject() *QMetaObject {
	return newQMetaObject(QCalendarWidget_virtualbase_MetaObject(unsafe.Pointer(this.h)))
}

func (this *QCalendarWidget) OnMetaObject(slot func(super func() *QMetaObject) *QMetaObject) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MetaObject(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MetaObject
func miqt_exec_callback_QCalendarWidget_MetaObject(self QCalendarWidget, cb intptr_t) *QMetaObject {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QMetaObject) *QMetaObject)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_MetaObject)

	return virtualReturn.cPointer()
}

func (this *QCalendarWidget) callVirtualBase_Metacast(param1 string) unsafe.Pointer {
	param1_Cstring := CString(param1)
	defer free(unsafe.Pointer(param1_Cstring))

	return (unsafe.Pointer)(QCalendarWidget_virtualbase_Metacast(unsafe.Pointer(this.h), param1_Cstring))
}

func (this *QCalendarWidget) OnMetacast(slot func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_Metacast(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_Metacast
func miqt_exec_callback_QCalendarWidget_Metacast(self QCalendarWidget, cb intptr_t, param1 *const_char) unsafe.Pointer {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 string) unsafe.Pointer, param1 string) unsafe.Pointer)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}
	// Convert all CABI parameters to Go parameters
	param1_ret := param1
	slotval1 := GoString(param1_ret)
	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_Metacast, slotval1)

	return virtualReturn
}
