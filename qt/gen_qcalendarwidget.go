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

	ret := newQCalendarWidget(QCalendarWidget_new(parent.cPointer()))
	ret.isSubclass = true
	return ret
}

// NewQCalendarWidget2 constructs a new QCalendarWidget object.
func NewQCalendarWidget2() *QCalendarWidget {

	ret := newQCalendarWidget(QCalendarWidget_new2())
	ret.isSubclass = true
	return ret
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

func (this *QCalendarWidget) callVirtualBase_SizeHint() *QSize {

	_goptr := newQSize(QCalendarWidget_virtualbase_SizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCalendarWidget) OnSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_SizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_SizeHint
func miqt_exec_callback_QCalendarWidget_SizeHint(self QCalendarWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_SizeHint)

	return virtualReturn.cPointer()

}

func (this *QCalendarWidget) callVirtualBase_MinimumSizeHint() *QSize {

	_goptr := newQSize(QCalendarWidget_virtualbase_MinimumSizeHint(unsafe.Pointer(this.h)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCalendarWidget) OnMinimumSizeHint(slot func(super func() *QSize) *QSize) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MinimumSizeHint(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MinimumSizeHint
func miqt_exec_callback_QCalendarWidget_MinimumSizeHint(self QCalendarWidget, cb intptr_t) *QSize {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QSize) *QSize)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_MinimumSizeHint)

	return virtualReturn.cPointer()

}

func (this *QCalendarWidget) callVirtualBase_Event(event *QEvent) bool {

	return (bool)(QCalendarWidget_virtualbase_Event(unsafe.Pointer(this.h), event.cPointer()))

}
func (this *QCalendarWidget) OnEvent(slot func(super func(event *QEvent) bool, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_Event(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_Event
func miqt_exec_callback_QCalendarWidget_Event(self QCalendarWidget, cb intptr_t, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent) bool, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_Event, slotval1)

	return (bool)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_EventFilter(watched *QObject, event *QEvent) bool {

	return (bool)(QCalendarWidget_virtualbase_EventFilter(unsafe.Pointer(this.h), watched.cPointer(), event.cPointer()))

}
func (this *QCalendarWidget) OnEventFilter(slot func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_EventFilter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_EventFilter
func miqt_exec_callback_QCalendarWidget_EventFilter(self QCalendarWidget, cb intptr_t, watched *QObject, event *QEvent) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(watched *QObject, event *QEvent) bool, watched *QObject, event *QEvent) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQObject(watched)

	slotval2 := newQEvent(event)

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_EventFilter, slotval1, slotval2)

	return (bool)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_MousePressEvent(event *QMouseEvent) {

	QCalendarWidget_virtualbase_MousePressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnMousePressEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MousePressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MousePressEvent
func miqt_exec_callback_QCalendarWidget_MousePressEvent(self QCalendarWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_MousePressEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_ResizeEvent(event *QResizeEvent) {

	QCalendarWidget_virtualbase_ResizeEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnResizeEvent(slot func(super func(event *QResizeEvent), event *QResizeEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_ResizeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_ResizeEvent
func miqt_exec_callback_QCalendarWidget_ResizeEvent(self QCalendarWidget, cb intptr_t, event *QResizeEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QResizeEvent), event *QResizeEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQResizeEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_ResizeEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_KeyPressEvent(event *QKeyEvent) {

	QCalendarWidget_virtualbase_KeyPressEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnKeyPressEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_KeyPressEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_KeyPressEvent
func miqt_exec_callback_QCalendarWidget_KeyPressEvent(self QCalendarWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_KeyPressEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_PaintCell(painter *QPainter, rect *QRect, date QDate) {

	QCalendarWidget_virtualbase_PaintCell(unsafe.Pointer(this.h), painter.cPointer(), rect.cPointer(), date.cPointer())

}
func (this *QCalendarWidget) OnPaintCell(slot func(super func(painter *QPainter, rect *QRect, date QDate), painter *QPainter, rect *QRect, date QDate)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_PaintCell(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_PaintCell
func miqt_exec_callback_QCalendarWidget_PaintCell(self QCalendarWidget, cb intptr_t, painter *QPainter, rect *QRect, date *QDate) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter, rect *QRect, date QDate), painter *QPainter, rect *QRect, date QDate))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	slotval2 := newQRect(rect)

	date_goptr := newQDate(date)
	date_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	slotval3 := *date_goptr

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_PaintCell, slotval1, slotval2, slotval3)

}

func (this *QCalendarWidget) callVirtualBase_DevType() int {

	return (int)(QCalendarWidget_virtualbase_DevType(unsafe.Pointer(this.h)))

}
func (this *QCalendarWidget) OnDevType(slot func(super func() int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_DevType(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_DevType
func miqt_exec_callback_QCalendarWidget_DevType(self QCalendarWidget, cb intptr_t) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_DevType)

	return (int)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_SetVisible(visible bool) {

	QCalendarWidget_virtualbase_SetVisible(unsafe.Pointer(this.h), (bool)(visible))

}
func (this *QCalendarWidget) OnSetVisible(slot func(super func(visible bool), visible bool)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_SetVisible(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_SetVisible
func miqt_exec_callback_QCalendarWidget_SetVisible(self QCalendarWidget, cb intptr_t, visible bool) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(visible bool), visible bool))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(visible)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_SetVisible, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_HeightForWidth(param1 int) int {

	return (int)(QCalendarWidget_virtualbase_HeightForWidth(unsafe.Pointer(this.h), (int)(param1)))

}
func (this *QCalendarWidget) OnHeightForWidth(slot func(super func(param1 int) int, param1 int) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_HeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_HeightForWidth
func miqt_exec_callback_QCalendarWidget_HeightForWidth(self QCalendarWidget, cb intptr_t, param1 int) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 int) int, param1 int) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (int)(param1)

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_HeightForWidth, slotval1)

	return (int)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_HasHeightForWidth() bool {

	return (bool)(QCalendarWidget_virtualbase_HasHeightForWidth(unsafe.Pointer(this.h)))

}
func (this *QCalendarWidget) OnHasHeightForWidth(slot func(super func() bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_HasHeightForWidth(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_HasHeightForWidth
func miqt_exec_callback_QCalendarWidget_HasHeightForWidth(self QCalendarWidget, cb intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_HasHeightForWidth)

	return (bool)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_PaintEngine() *QPaintEngine {

	return newQPaintEngine(QCalendarWidget_virtualbase_PaintEngine(unsafe.Pointer(this.h)))

}
func (this *QCalendarWidget) OnPaintEngine(slot func(super func() *QPaintEngine) *QPaintEngine) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_PaintEngine(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_PaintEngine
func miqt_exec_callback_QCalendarWidget_PaintEngine(self QCalendarWidget, cb intptr_t) *QPaintEngine {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPaintEngine) *QPaintEngine)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_PaintEngine)

	return virtualReturn.cPointer()

}

func (this *QCalendarWidget) callVirtualBase_MouseReleaseEvent(event *QMouseEvent) {

	QCalendarWidget_virtualbase_MouseReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnMouseReleaseEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MouseReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MouseReleaseEvent
func miqt_exec_callback_QCalendarWidget_MouseReleaseEvent(self QCalendarWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_MouseReleaseEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_MouseDoubleClickEvent(event *QMouseEvent) {

	QCalendarWidget_virtualbase_MouseDoubleClickEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnMouseDoubleClickEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MouseDoubleClickEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MouseDoubleClickEvent
func miqt_exec_callback_QCalendarWidget_MouseDoubleClickEvent(self QCalendarWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_MouseDoubleClickEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_MouseMoveEvent(event *QMouseEvent) {

	QCalendarWidget_virtualbase_MouseMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnMouseMoveEvent(slot func(super func(event *QMouseEvent), event *QMouseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MouseMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MouseMoveEvent
func miqt_exec_callback_QCalendarWidget_MouseMoveEvent(self QCalendarWidget, cb intptr_t, event *QMouseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMouseEvent), event *QMouseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMouseEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_MouseMoveEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_WheelEvent(event *QWheelEvent) {

	QCalendarWidget_virtualbase_WheelEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnWheelEvent(slot func(super func(event *QWheelEvent), event *QWheelEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_WheelEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_WheelEvent
func miqt_exec_callback_QCalendarWidget_WheelEvent(self QCalendarWidget, cb intptr_t, event *QWheelEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QWheelEvent), event *QWheelEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQWheelEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_WheelEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_KeyReleaseEvent(event *QKeyEvent) {

	QCalendarWidget_virtualbase_KeyReleaseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnKeyReleaseEvent(slot func(super func(event *QKeyEvent), event *QKeyEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_KeyReleaseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_KeyReleaseEvent
func miqt_exec_callback_QCalendarWidget_KeyReleaseEvent(self QCalendarWidget, cb intptr_t, event *QKeyEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QKeyEvent), event *QKeyEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQKeyEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_KeyReleaseEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_FocusInEvent(event *QFocusEvent) {

	QCalendarWidget_virtualbase_FocusInEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnFocusInEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_FocusInEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_FocusInEvent
func miqt_exec_callback_QCalendarWidget_FocusInEvent(self QCalendarWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_FocusInEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_FocusOutEvent(event *QFocusEvent) {

	QCalendarWidget_virtualbase_FocusOutEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnFocusOutEvent(slot func(super func(event *QFocusEvent), event *QFocusEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_FocusOutEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_FocusOutEvent
func miqt_exec_callback_QCalendarWidget_FocusOutEvent(self QCalendarWidget, cb intptr_t, event *QFocusEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QFocusEvent), event *QFocusEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQFocusEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_FocusOutEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_EnterEvent(event *QEnterEvent) {

	QCalendarWidget_virtualbase_EnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnEnterEvent(slot func(super func(event *QEnterEvent), event *QEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_EnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_EnterEvent
func miqt_exec_callback_QCalendarWidget_EnterEvent(self QCalendarWidget, cb intptr_t, event *QEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEnterEvent), event *QEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEnterEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_EnterEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_LeaveEvent(event *QEvent) {

	QCalendarWidget_virtualbase_LeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnLeaveEvent(slot func(super func(event *QEvent), event *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_LeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_LeaveEvent
func miqt_exec_callback_QCalendarWidget_LeaveEvent(self QCalendarWidget, cb intptr_t, event *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QEvent), event *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_LeaveEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_PaintEvent(event *QPaintEvent) {

	QCalendarWidget_virtualbase_PaintEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnPaintEvent(slot func(super func(event *QPaintEvent), event *QPaintEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_PaintEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_PaintEvent
func miqt_exec_callback_QCalendarWidget_PaintEvent(self QCalendarWidget, cb intptr_t, event *QPaintEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QPaintEvent), event *QPaintEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPaintEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_PaintEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_MoveEvent(event *QMoveEvent) {

	QCalendarWidget_virtualbase_MoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnMoveEvent(slot func(super func(event *QMoveEvent), event *QMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_MoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_MoveEvent
func miqt_exec_callback_QCalendarWidget_MoveEvent(self QCalendarWidget, cb intptr_t, event *QMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QMoveEvent), event *QMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQMoveEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_MoveEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_CloseEvent(event *QCloseEvent) {

	QCalendarWidget_virtualbase_CloseEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnCloseEvent(slot func(super func(event *QCloseEvent), event *QCloseEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_CloseEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_CloseEvent
func miqt_exec_callback_QCalendarWidget_CloseEvent(self QCalendarWidget, cb intptr_t, event *QCloseEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QCloseEvent), event *QCloseEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQCloseEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_CloseEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_ContextMenuEvent(event *QContextMenuEvent) {

	QCalendarWidget_virtualbase_ContextMenuEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnContextMenuEvent(slot func(super func(event *QContextMenuEvent), event *QContextMenuEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_ContextMenuEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_ContextMenuEvent
func miqt_exec_callback_QCalendarWidget_ContextMenuEvent(self QCalendarWidget, cb intptr_t, event *QContextMenuEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QContextMenuEvent), event *QContextMenuEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQContextMenuEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_ContextMenuEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_TabletEvent(event *QTabletEvent) {

	QCalendarWidget_virtualbase_TabletEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnTabletEvent(slot func(super func(event *QTabletEvent), event *QTabletEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_TabletEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_TabletEvent
func miqt_exec_callback_QCalendarWidget_TabletEvent(self QCalendarWidget, cb intptr_t, event *QTabletEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QTabletEvent), event *QTabletEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQTabletEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_TabletEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_ActionEvent(event *QActionEvent) {

	QCalendarWidget_virtualbase_ActionEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnActionEvent(slot func(super func(event *QActionEvent), event *QActionEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_ActionEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_ActionEvent
func miqt_exec_callback_QCalendarWidget_ActionEvent(self QCalendarWidget, cb intptr_t, event *QActionEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QActionEvent), event *QActionEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQActionEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_ActionEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_DragEnterEvent(event *QDragEnterEvent) {

	QCalendarWidget_virtualbase_DragEnterEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnDragEnterEvent(slot func(super func(event *QDragEnterEvent), event *QDragEnterEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_DragEnterEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_DragEnterEvent
func miqt_exec_callback_QCalendarWidget_DragEnterEvent(self QCalendarWidget, cb intptr_t, event *QDragEnterEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragEnterEvent), event *QDragEnterEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragEnterEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_DragEnterEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_DragMoveEvent(event *QDragMoveEvent) {

	QCalendarWidget_virtualbase_DragMoveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnDragMoveEvent(slot func(super func(event *QDragMoveEvent), event *QDragMoveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_DragMoveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_DragMoveEvent
func miqt_exec_callback_QCalendarWidget_DragMoveEvent(self QCalendarWidget, cb intptr_t, event *QDragMoveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragMoveEvent), event *QDragMoveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragMoveEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_DragMoveEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_DragLeaveEvent(event *QDragLeaveEvent) {

	QCalendarWidget_virtualbase_DragLeaveEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnDragLeaveEvent(slot func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_DragLeaveEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_DragLeaveEvent
func miqt_exec_callback_QCalendarWidget_DragLeaveEvent(self QCalendarWidget, cb intptr_t, event *QDragLeaveEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDragLeaveEvent), event *QDragLeaveEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDragLeaveEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_DragLeaveEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_DropEvent(event *QDropEvent) {

	QCalendarWidget_virtualbase_DropEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnDropEvent(slot func(super func(event *QDropEvent), event *QDropEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_DropEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_DropEvent
func miqt_exec_callback_QCalendarWidget_DropEvent(self QCalendarWidget, cb intptr_t, event *QDropEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QDropEvent), event *QDropEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQDropEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_DropEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_ShowEvent(event *QShowEvent) {

	QCalendarWidget_virtualbase_ShowEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnShowEvent(slot func(super func(event *QShowEvent), event *QShowEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_ShowEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_ShowEvent
func miqt_exec_callback_QCalendarWidget_ShowEvent(self QCalendarWidget, cb intptr_t, event *QShowEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QShowEvent), event *QShowEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQShowEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_ShowEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_HideEvent(event *QHideEvent) {

	QCalendarWidget_virtualbase_HideEvent(unsafe.Pointer(this.h), event.cPointer())

}
func (this *QCalendarWidget) OnHideEvent(slot func(super func(event *QHideEvent), event *QHideEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_HideEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_HideEvent
func miqt_exec_callback_QCalendarWidget_HideEvent(self QCalendarWidget, cb intptr_t, event *QHideEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(event *QHideEvent), event *QHideEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQHideEvent(event)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_HideEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_NativeEvent(eventType []byte, message unsafe.Pointer, result *uintptr) bool {
	eventType_alias := struct_miqt_string{}
	eventType_alias.data = (char)(unsafe.Pointer(&eventType[0]))
	eventType_alias.len = size_t(len(eventType))

	return (bool)(QCalendarWidget_virtualbase_NativeEvent(unsafe.Pointer(this.h), eventType_alias, message, (*intptr_t)(unsafe.Pointer(result))))

}
func (this *QCalendarWidget) OnNativeEvent(slot func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_NativeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_NativeEvent
func miqt_exec_callback_QCalendarWidget_NativeEvent(self QCalendarWidget, cb intptr_t, eventType struct_miqt_string, message unsafe.Pointer, result *intptr_t) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(eventType []byte, message unsafe.Pointer, result *uintptr) bool, eventType []byte, message unsafe.Pointer, result *uintptr) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	var eventType_bytearray struct_miqt_string = eventType
	eventType_ret := GoBytes(unsafe.Pointer(eventType_bytearray.data), int(int64(eventType_bytearray.len)))
	free(unsafe.Pointer(eventType_bytearray.data))
	slotval1 := eventType_ret
	slotval2 := (unsafe.Pointer)(message)

	slotval3 := (*uintptr)(unsafe.Pointer(result))

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_NativeEvent, slotval1, slotval2, slotval3)

	return (bool)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_ChangeEvent(param1 *QEvent) {

	QCalendarWidget_virtualbase_ChangeEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCalendarWidget) OnChangeEvent(slot func(super func(param1 *QEvent), param1 *QEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_ChangeEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_ChangeEvent
func miqt_exec_callback_QCalendarWidget_ChangeEvent(self QCalendarWidget, cb intptr_t, param1 *QEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QEvent), param1 *QEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQEvent(param1)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_ChangeEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_Metric(param1 PaintDeviceMetric) int {

	return (int)(QCalendarWidget_virtualbase_Metric(unsafe.Pointer(this.h), param1))

}
func (this *QCalendarWidget) OnMetric(slot func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_Metric(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_Metric
func miqt_exec_callback_QCalendarWidget_Metric(self QCalendarWidget, cb intptr_t, param1 PaintDeviceMetric) int {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 PaintDeviceMetric) int, param1 PaintDeviceMetric) int)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	xxxxxxxxx

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_Metric, slotval1)

	return (int)(virtualReturn)

}

func (this *QCalendarWidget) callVirtualBase_InitPainter(painter *QPainter) {

	QCalendarWidget_virtualbase_InitPainter(unsafe.Pointer(this.h), painter.cPointer())

}
func (this *QCalendarWidget) OnInitPainter(slot func(super func(painter *QPainter), painter *QPainter)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_InitPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_InitPainter
func miqt_exec_callback_QCalendarWidget_InitPainter(self QCalendarWidget, cb intptr_t, painter *QPainter) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(painter *QPainter), painter *QPainter))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPainter(painter)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_InitPainter, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_Redirected(offset *QPoint) *QPaintDevice {

	return newQPaintDevice(QCalendarWidget_virtualbase_Redirected(unsafe.Pointer(this.h), offset.cPointer()))

}
func (this *QCalendarWidget) OnRedirected(slot func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_Redirected(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_Redirected
func miqt_exec_callback_QCalendarWidget_Redirected(self QCalendarWidget, cb intptr_t, offset *QPoint) *QPaintDevice {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(offset *QPoint) *QPaintDevice, offset *QPoint) *QPaintDevice)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQPoint(offset)

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_Redirected, slotval1)

	return virtualReturn.cPointer()

}

func (this *QCalendarWidget) callVirtualBase_SharedPainter() *QPainter {

	return newQPainter(QCalendarWidget_virtualbase_SharedPainter(unsafe.Pointer(this.h)))

}
func (this *QCalendarWidget) OnSharedPainter(slot func(super func() *QPainter) *QPainter) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_SharedPainter(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_SharedPainter
func miqt_exec_callback_QCalendarWidget_SharedPainter(self QCalendarWidget, cb intptr_t) *QPainter {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func() *QPainter) *QPainter)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_SharedPainter)

	return virtualReturn.cPointer()

}

func (this *QCalendarWidget) callVirtualBase_InputMethodEvent(param1 *QInputMethodEvent) {

	QCalendarWidget_virtualbase_InputMethodEvent(unsafe.Pointer(this.h), param1.cPointer())

}
func (this *QCalendarWidget) OnInputMethodEvent(slot func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent)) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_InputMethodEvent(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_InputMethodEvent
func miqt_exec_callback_QCalendarWidget_InputMethodEvent(self QCalendarWidget, cb intptr_t, param1 *QInputMethodEvent) {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 *QInputMethodEvent), param1 *QInputMethodEvent))
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := newQInputMethodEvent(param1)

	gofunc((&QCalendarWidget{h: self}).callVirtualBase_InputMethodEvent, slotval1)

}

func (this *QCalendarWidget) callVirtualBase_InputMethodQuery(param1 InputMethodQuery) *QVariant {

	_goptr := newQVariant(QCalendarWidget_virtualbase_InputMethodQuery(unsafe.Pointer(this.h), (int)(param1)))
	_goptr.GoGC() // Qt uses pass-by-value semantics for this type. Mimic with finalizer
	return _goptr

}
func (this *QCalendarWidget) OnInputMethodQuery(slot func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_InputMethodQuery(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_InputMethodQuery
func miqt_exec_callback_QCalendarWidget_InputMethodQuery(self QCalendarWidget, cb intptr_t, param1 int) *QVariant {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(param1 InputMethodQuery) *QVariant, param1 InputMethodQuery) *QVariant)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (InputMethodQuery)(param1)

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_InputMethodQuery, slotval1)

	return virtualReturn.cPointer()

}

func (this *QCalendarWidget) callVirtualBase_FocusNextPrevChild(next bool) bool {

	return (bool)(QCalendarWidget_virtualbase_FocusNextPrevChild(unsafe.Pointer(this.h), (bool)(next)))

}
func (this *QCalendarWidget) OnFocusNextPrevChild(slot func(super func(next bool) bool, next bool) bool) {
	if !this.isSubclass {
		panic("miqt: can only override virtual methods for directly constructed types")
	}
	QCalendarWidget_override_virtual_FocusNextPrevChild(unsafe.Pointer(this.h), intptr_t(cgo.NewHandle(slot)))
}

//export miqt_exec_callback_QCalendarWidget_FocusNextPrevChild
func miqt_exec_callback_QCalendarWidget_FocusNextPrevChild(self QCalendarWidget, cb intptr_t, next bool) bool {
	gofunc, ok := cgo.Handle(cb).Value().(func(super func(next bool) bool, next bool) bool)
	if !ok {
		panic("miqt: callback of non-callback type (heap corruption?)")
	}

	// Convert all CABI parameters to Go parameters
	slotval1 := (bool)(next)

	virtualReturn := gofunc((&QCalendarWidget{h: self}).callVirtualBase_FocusNextPrevChild, slotval1)

	return (bool)(virtualReturn)

}
