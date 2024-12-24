// +build ignore

#include <QCalendar>
#include <QCalendarWidget>
#include <QDate>
#include <QEvent>
#include <QKeyEvent>
#include <QMap>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPainter>
#include <QRect>
#include <QResizeEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTextCharFormat>
#include <QWidget>
#include <qcalendarwidget.h>
#include "gen_qcalendarwidget.h"
class MiqtVirtualQCalendarWidget : public virtual QCalendarWidget {
public:
MiqtVirtualQCalendarWidget(QWidget* parent): QCalendarWidget(parent) {};
MiqtVirtualQCalendarWidget(): QCalendarWidget() {};

virtual ~MiqtVirtualQCalendarWidget() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QCalendarWidget::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QCalendarWidget_MetaObject(const_cast<MiqtVirtualQCalendarWidget*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QCalendarWidget::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QCalendarWidget::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QCalendarWidget_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QCalendarWidget::qt_metacast(param1);

	}
};
QCalendarWidget* QCalendarWidget_new(QWidget* parent) {
return new MiqtVirtualQCalendarWidget(parent);
}
QCalendarWidget* QCalendarWidget_new2() {
return new MiqtVirtualQCalendarWidget();
}
void QCalendarWidget_virtbase(QCalendarWidget* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QCalendarWidget_MetaObject(const QCalendarWidget* self) {
	return (QMetaObject*) self->metaObject();
}
void* QCalendarWidget_Metacast(QCalendarWidget* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QCalendarWidget_Tr(const char* s) {
	QString _ret = QCalendarWidget::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QCalendarWidget_SizeHint(const QCalendarWidget* self) {
	return new QSize(self->sizeHint());
}
QSize* QCalendarWidget_MinimumSizeHint(const QCalendarWidget* self) {
	return new QSize(self->minimumSizeHint());
}
QDate* QCalendarWidget_SelectedDate(const QCalendarWidget* self) {
	return new QDate(self->selectedDate());
}
int QCalendarWidget_YearShown(const QCalendarWidget* self) {
	return self->yearShown();
}
int QCalendarWidget_MonthShown(const QCalendarWidget* self) {
	return self->monthShown();
}
QDate* QCalendarWidget_MinimumDate(const QCalendarWidget* self) {
	return new QDate(self->minimumDate());
}
void QCalendarWidget_SetMinimumDate(QCalendarWidget* self, QDate* date) {
	self->setMinimumDate(*date);
}
void QCalendarWidget_ClearMinimumDate(QCalendarWidget* self) {
	self->clearMinimumDate();
}
QDate* QCalendarWidget_MaximumDate(const QCalendarWidget* self) {
	return new QDate(self->maximumDate());
}
void QCalendarWidget_SetMaximumDate(QCalendarWidget* self, QDate* date) {
	self->setMaximumDate(*date);
}
void QCalendarWidget_ClearMaximumDate(QCalendarWidget* self) {
	self->clearMaximumDate();
}
int QCalendarWidget_FirstDayOfWeek(const QCalendarWidget* self) {
	Qt::DayOfWeek _ret = self->firstDayOfWeek();
	return static_cast<int>(_ret);
}
void QCalendarWidget_SetFirstDayOfWeek(QCalendarWidget* self, int dayOfWeek) {
	self->setFirstDayOfWeek(static_cast<Qt::DayOfWeek>(dayOfWeek));
}
bool QCalendarWidget_IsNavigationBarVisible(const QCalendarWidget* self) {
	return self->isNavigationBarVisible();
}
bool QCalendarWidget_IsGridVisible(const QCalendarWidget* self) {
	return self->isGridVisible();
}
QCalendar* QCalendarWidget_Calendar(const QCalendarWidget* self) {
	return new QCalendar(self->calendar());
}
void QCalendarWidget_SetCalendar(QCalendarWidget* self, QCalendar* calendar) {
	self->setCalendar(*calendar);
}
SelectionMode QCalendarWidget_SelectionMode(const QCalendarWidget* self) {
	return self->selectionMode();
}
void QCalendarWidget_SetSelectionMode(QCalendarWidget* self, SelectionMode mode) {
	self->setSelectionMode(mode);
}
HorizontalHeaderFormat QCalendarWidget_HorizontalHeaderFormat(const QCalendarWidget* self) {
	return self->horizontalHeaderFormat();
}
void QCalendarWidget_SetHorizontalHeaderFormat(QCalendarWidget* self, HorizontalHeaderFormat format) {
	self->setHorizontalHeaderFormat(format);
}
VerticalHeaderFormat QCalendarWidget_VerticalHeaderFormat(const QCalendarWidget* self) {
	return self->verticalHeaderFormat();
}
void QCalendarWidget_SetVerticalHeaderFormat(QCalendarWidget* self, VerticalHeaderFormat format) {
	self->setVerticalHeaderFormat(format);
}
QTextCharFormat* QCalendarWidget_HeaderTextFormat(const QCalendarWidget* self) {
	return new QTextCharFormat(self->headerTextFormat());
}
void QCalendarWidget_SetHeaderTextFormat(QCalendarWidget* self, QTextCharFormat* format) {
	self->setHeaderTextFormat(*format);
}
QTextCharFormat* QCalendarWidget_WeekdayTextFormat(const QCalendarWidget* self, int dayOfWeek) {
	return new QTextCharFormat(self->weekdayTextFormat(static_cast<Qt::DayOfWeek>(dayOfWeek)));
}
void QCalendarWidget_SetWeekdayTextFormat(QCalendarWidget* self, int dayOfWeek, QTextCharFormat* format) {
	self->setWeekdayTextFormat(static_cast<Qt::DayOfWeek>(dayOfWeek), *format);
}
struct miqt_map /* of QDate* to QTextCharFormat* */  QCalendarWidget_DateTextFormat(const QCalendarWidget* self) {
	QMap<QDate, QTextCharFormat> _ret = self->dateTextFormat();
	// Convert QMap<> from C++ memory to manually-managed C memory
	QDate** _karr = static_cast<QDate**>(malloc(sizeof(QDate*) * _ret.size()));
	QTextCharFormat** _varr = static_cast<QTextCharFormat**>(malloc(sizeof(QTextCharFormat*) * _ret.size()));
	int _ctr = 0;
	for (auto _itr = _ret.keyValueBegin(); _itr != _ret.keyValueEnd(); ++_itr) {
		_karr[_ctr] = new QDate(_itr->first);
		_varr[_ctr] = new QTextCharFormat(_itr->second);
		_ctr++;
	}
	struct miqt_map _out;
	_out.len = _ret.size();
	_out.keys = static_cast<void*>(_karr);
	_out.values = static_cast<void*>(_varr);
	return _out;
}
QTextCharFormat* QCalendarWidget_DateTextFormatWithDate(const QCalendarWidget* self, QDate* date) {
	return new QTextCharFormat(self->dateTextFormat(*date));
}
void QCalendarWidget_SetDateTextFormat(QCalendarWidget* self, QDate* date, QTextCharFormat* format) {
	self->setDateTextFormat(*date, *format);
}
bool QCalendarWidget_IsDateEditEnabled(const QCalendarWidget* self) {
	return self->isDateEditEnabled();
}
void QCalendarWidget_SetDateEditEnabled(QCalendarWidget* self, bool enable) {
	self->setDateEditEnabled(enable);
}
int QCalendarWidget_DateEditAcceptDelay(const QCalendarWidget* self) {
	return self->dateEditAcceptDelay();
}
void QCalendarWidget_SetDateEditAcceptDelay(QCalendarWidget* self, int delay) {
	self->setDateEditAcceptDelay(static_cast<int>(delay));
}
void QCalendarWidget_SetSelectedDate(QCalendarWidget* self, QDate* date) {
	self->setSelectedDate(*date);
}
void QCalendarWidget_SetDateRange(QCalendarWidget* self, QDate* min, QDate* max) {
	self->setDateRange(*min, *max);
}
void QCalendarWidget_SetCurrentPage(QCalendarWidget* self, int year, int month) {
	self->setCurrentPage(static_cast<int>(year), static_cast<int>(month));
}
void QCalendarWidget_SetGridVisible(QCalendarWidget* self, bool show) {
	self->setGridVisible(show);
}
void QCalendarWidget_SetNavigationBarVisible(QCalendarWidget* self, bool visible) {
	self->setNavigationBarVisible(visible);
}
void QCalendarWidget_ShowNextMonth(QCalendarWidget* self) {
	self->showNextMonth();
}
void QCalendarWidget_ShowPreviousMonth(QCalendarWidget* self) {
	self->showPreviousMonth();
}
void QCalendarWidget_ShowNextYear(QCalendarWidget* self) {
	self->showNextYear();
}
void QCalendarWidget_ShowPreviousYear(QCalendarWidget* self) {
	self->showPreviousYear();
}
void QCalendarWidget_ShowSelectedDate(QCalendarWidget* self) {
	self->showSelectedDate();
}
void QCalendarWidget_ShowToday(QCalendarWidget* self) {
	self->showToday();
}
void QCalendarWidget_SelectionChanged(QCalendarWidget* self) {
	self->selectionChanged();
}
void QCalendarWidget_connect_SelectionChanged(QCalendarWidget* self, intptr_t slot) {
	MiqtVirtualQCalendarWidget::connect(self, static_cast<void (QCalendarWidget::*)()>(&QCalendarWidget::selectionChanged), self, [=]() {
		miqt_exec_callback_QCalendarWidget_SelectionChanged(slot);
	});
}
void QCalendarWidget_Clicked(QCalendarWidget* self, QDate* date) {
	self->clicked(*date);
}
void QCalendarWidget_connect_Clicked(QCalendarWidget* self, intptr_t slot) {
	MiqtVirtualQCalendarWidget::connect(self, static_cast<void (QCalendarWidget::*)(QDate)>(&QCalendarWidget::clicked), self, [=](QDate date) {
		QDate* sigval1 = new QDate(date);
		miqt_exec_callback_QCalendarWidget_Clicked(slot, sigval1);
	});
}
void QCalendarWidget_Activated(QCalendarWidget* self, QDate* date) {
	self->activated(*date);
}
void QCalendarWidget_connect_Activated(QCalendarWidget* self, intptr_t slot) {
	MiqtVirtualQCalendarWidget::connect(self, static_cast<void (QCalendarWidget::*)(QDate)>(&QCalendarWidget::activated), self, [=](QDate date) {
		QDate* sigval1 = new QDate(date);
		miqt_exec_callback_QCalendarWidget_Activated(slot, sigval1);
	});
}
void QCalendarWidget_CurrentPageChanged(QCalendarWidget* self, int year, int month) {
	self->currentPageChanged(static_cast<int>(year), static_cast<int>(month));
}
void QCalendarWidget_connect_CurrentPageChanged(QCalendarWidget* self, intptr_t slot) {
	MiqtVirtualQCalendarWidget::connect(self, static_cast<void (QCalendarWidget::*)(int, int)>(&QCalendarWidget::currentPageChanged), self, [=](int year, int month) {
		int sigval1 = year;
		int sigval2 = month;
		miqt_exec_callback_QCalendarWidget_CurrentPageChanged(slot, sigval1, sigval2);
	});
}
struct miqt_string QCalendarWidget_Tr2(const char* s, const char* c) {
	QString _ret = QCalendarWidget::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QCalendarWidget_Tr3(const char* s, const char* c, int n) {
	QString _ret = QCalendarWidget::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QCalendarWidget_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCalendarWidget*>( (QCalendarWidget*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QCalendarWidget_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQCalendarWidget*)(self) )->virtualbase_MetaObject();
}
void QCalendarWidget_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCalendarWidget*>( (QCalendarWidget*)(self) )->handle__Metacast = slot;
}
void* QCalendarWidget_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQCalendarWidget*)(self) )->virtualbase_Metacast(param1);
}
void QCalendarWidget_Delete(QCalendarWidget* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQCalendarWidget*>( self );
	} else {
		delete self;
	}
}
