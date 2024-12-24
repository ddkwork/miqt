// +build ignore

#include <QAbstractSpinBox>
#include <QCalendar>
#include <QCalendarWidget>
#include <QDate>
#include <QDateEdit>
#include <QDateTime>
#include <QDateTimeEdit>
#include <QEvent>
#include <QFocusEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionSpinBox>
#include <QTime>
#include <QTimeEdit>
#include <QTimeZone>
#include <QWheelEvent>
#include <QWidget>
#include <qdatetimeedit.h>
#include "gen_qdatetimeedit.h"
class MiqtVirtualQDateTimeEdit : public virtual QDateTimeEdit {
public:
MiqtVirtualQDateTimeEdit(QWidget* parent): QDateTimeEdit(parent) {};
MiqtVirtualQDateTimeEdit(): QDateTimeEdit() {};
MiqtVirtualQDateTimeEdit(const QDateTime& dt): QDateTimeEdit(dt) {};
MiqtVirtualQDateTimeEdit(QDate d): QDateTimeEdit(d) {};
MiqtVirtualQDateTimeEdit(QTime t): QDateTimeEdit(t) {};
MiqtVirtualQDateTimeEdit(const QDateTime& dt, QWidget* parent): QDateTimeEdit(dt, parent) {};
MiqtVirtualQDateTimeEdit(QDate d, QWidget* parent): QDateTimeEdit(d, parent) {};
MiqtVirtualQDateTimeEdit(QTime t, QWidget* parent): QDateTimeEdit(t, parent) {};

virtual ~MiqtVirtualQDateTimeEdit() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QDateTimeEdit::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QDateTimeEdit_MetaObject(const_cast<MiqtVirtualQDateTimeEdit*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QDateTimeEdit::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QDateTimeEdit::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QDateTimeEdit_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QDateTimeEdit::qt_metacast(param1);

	}
};
QDateTimeEdit* QDateTimeEdit_new(QWidget* parent) {
return new MiqtVirtualQDateTimeEdit(parent);
}
QDateTimeEdit* QDateTimeEdit_new2() {
return new MiqtVirtualQDateTimeEdit();
}
QDateTimeEdit* QDateTimeEdit_new3(QDateTime* dt) {
return new MiqtVirtualQDateTimeEdit(*dt);
}
QDateTimeEdit* QDateTimeEdit_new4(QDate* d) {
return new MiqtVirtualQDateTimeEdit(*d);
}
QDateTimeEdit* QDateTimeEdit_new5(QTime* t) {
return new MiqtVirtualQDateTimeEdit(*t);
}
QDateTimeEdit* QDateTimeEdit_new6(QDateTime* dt, QWidget* parent) {
return new MiqtVirtualQDateTimeEdit(*dt, parent);
}
QDateTimeEdit* QDateTimeEdit_new7(QDate* d, QWidget* parent) {
return new MiqtVirtualQDateTimeEdit(*d, parent);
}
QDateTimeEdit* QDateTimeEdit_new8(QTime* t, QWidget* parent) {
return new MiqtVirtualQDateTimeEdit(*t, parent);
}
void QDateTimeEdit_virtbase(QDateTimeEdit* src
, QAbstractSpinBox** outptr_QAbstractSpinBox
) {
*outptr_QAbstractSpinBox = static_cast<QAbstractSpinBox*>(src);
}
QMetaObject* QDateTimeEdit_MetaObject(const QDateTimeEdit* self) {
	return (QMetaObject*) self->metaObject();
}
void* QDateTimeEdit_Metacast(QDateTimeEdit* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QDateTimeEdit_Tr(const char* s) {
	QString _ret = QDateTimeEdit::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QDateTime* QDateTimeEdit_DateTime(const QDateTimeEdit* self) {
	return new QDateTime(self->dateTime());
}
QDate* QDateTimeEdit_Date(const QDateTimeEdit* self) {
	return new QDate(self->date());
}
QTime* QDateTimeEdit_Time(const QDateTimeEdit* self) {
	return new QTime(self->time());
}
QCalendar* QDateTimeEdit_Calendar(const QDateTimeEdit* self) {
	return new QCalendar(self->calendar());
}
void QDateTimeEdit_SetCalendar(QDateTimeEdit* self, QCalendar* calendar) {
	self->setCalendar(*calendar);
}
QDateTime* QDateTimeEdit_MinimumDateTime(const QDateTimeEdit* self) {
	return new QDateTime(self->minimumDateTime());
}
void QDateTimeEdit_ClearMinimumDateTime(QDateTimeEdit* self) {
	self->clearMinimumDateTime();
}
void QDateTimeEdit_SetMinimumDateTime(QDateTimeEdit* self, QDateTime* dt) {
	self->setMinimumDateTime(*dt);
}
QDateTime* QDateTimeEdit_MaximumDateTime(const QDateTimeEdit* self) {
	return new QDateTime(self->maximumDateTime());
}
void QDateTimeEdit_ClearMaximumDateTime(QDateTimeEdit* self) {
	self->clearMaximumDateTime();
}
void QDateTimeEdit_SetMaximumDateTime(QDateTimeEdit* self, QDateTime* dt) {
	self->setMaximumDateTime(*dt);
}
void QDateTimeEdit_SetDateTimeRange(QDateTimeEdit* self, QDateTime* min, QDateTime* max) {
	self->setDateTimeRange(*min, *max);
}
QDate* QDateTimeEdit_MinimumDate(const QDateTimeEdit* self) {
	return new QDate(self->minimumDate());
}
void QDateTimeEdit_SetMinimumDate(QDateTimeEdit* self, QDate* min) {
	self->setMinimumDate(*min);
}
void QDateTimeEdit_ClearMinimumDate(QDateTimeEdit* self) {
	self->clearMinimumDate();
}
QDate* QDateTimeEdit_MaximumDate(const QDateTimeEdit* self) {
	return new QDate(self->maximumDate());
}
void QDateTimeEdit_SetMaximumDate(QDateTimeEdit* self, QDate* max) {
	self->setMaximumDate(*max);
}
void QDateTimeEdit_ClearMaximumDate(QDateTimeEdit* self) {
	self->clearMaximumDate();
}
void QDateTimeEdit_SetDateRange(QDateTimeEdit* self, QDate* min, QDate* max) {
	self->setDateRange(*min, *max);
}
QTime* QDateTimeEdit_MinimumTime(const QDateTimeEdit* self) {
	return new QTime(self->minimumTime());
}
void QDateTimeEdit_SetMinimumTime(QDateTimeEdit* self, QTime* min) {
	self->setMinimumTime(*min);
}
void QDateTimeEdit_ClearMinimumTime(QDateTimeEdit* self) {
	self->clearMinimumTime();
}
QTime* QDateTimeEdit_MaximumTime(const QDateTimeEdit* self) {
	return new QTime(self->maximumTime());
}
void QDateTimeEdit_SetMaximumTime(QDateTimeEdit* self, QTime* max) {
	self->setMaximumTime(*max);
}
void QDateTimeEdit_ClearMaximumTime(QDateTimeEdit* self) {
	self->clearMaximumTime();
}
void QDateTimeEdit_SetTimeRange(QDateTimeEdit* self, QTime* min, QTime* max) {
	self->setTimeRange(*min, *max);
}
Sections QDateTimeEdit_DisplayedSections(const QDateTimeEdit* self) {
	return self->displayedSections();
}
Section QDateTimeEdit_CurrentSection(const QDateTimeEdit* self) {
	return self->currentSection();
}
Section QDateTimeEdit_SectionAt(const QDateTimeEdit* self, int index) {
	return self->sectionAt(static_cast<int>(index));
}
void QDateTimeEdit_SetCurrentSection(QDateTimeEdit* self, Section section) {
	self->setCurrentSection(section);
}
int QDateTimeEdit_CurrentSectionIndex(const QDateTimeEdit* self) {
	return self->currentSectionIndex();
}
void QDateTimeEdit_SetCurrentSectionIndex(QDateTimeEdit* self, int index) {
	self->setCurrentSectionIndex(static_cast<int>(index));
}
QCalendarWidget* QDateTimeEdit_CalendarWidget(const QDateTimeEdit* self) {
	return self->calendarWidget();
}
void QDateTimeEdit_SetCalendarWidget(QDateTimeEdit* self, QCalendarWidget* calendarWidget) {
	self->setCalendarWidget(calendarWidget);
}
int QDateTimeEdit_SectionCount(const QDateTimeEdit* self) {
	return self->sectionCount();
}
void QDateTimeEdit_SetSelectedSection(QDateTimeEdit* self, Section section) {
	self->setSelectedSection(section);
}
struct miqt_string QDateTimeEdit_SectionText(const QDateTimeEdit* self, Section section) {
	QString _ret = self->sectionText(section);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QDateTimeEdit_DisplayFormat(const QDateTimeEdit* self) {
	QString _ret = self->displayFormat();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QDateTimeEdit_SetDisplayFormat(QDateTimeEdit* self, struct miqt_string format) {
	QString format_QString = QString::fromUtf8(format.data, format.len);
	self->setDisplayFormat(format_QString);
}
bool QDateTimeEdit_CalendarPopup(const QDateTimeEdit* self) {
	return self->calendarPopup();
}
void QDateTimeEdit_SetCalendarPopup(QDateTimeEdit* self, bool enable) {
	self->setCalendarPopup(enable);
}
int QDateTimeEdit_TimeSpec(const QDateTimeEdit* self) {
	Qt::TimeSpec _ret = self->timeSpec();
	return static_cast<int>(_ret);
}
void QDateTimeEdit_SetTimeSpec(QDateTimeEdit* self, int spec) {
	self->setTimeSpec(static_cast<Qt::TimeSpec>(spec));
}
QTimeZone* QDateTimeEdit_TimeZone(const QDateTimeEdit* self) {
	return new QTimeZone(self->timeZone());
}
void QDateTimeEdit_SetTimeZone(QDateTimeEdit* self, QTimeZone* zone) {
	self->setTimeZone(*zone);
}
QSize* QDateTimeEdit_SizeHint(const QDateTimeEdit* self) {
	return new QSize(self->sizeHint());
}
void QDateTimeEdit_Clear(QDateTimeEdit* self) {
	self->clear();
}
void QDateTimeEdit_StepBy(QDateTimeEdit* self, int steps) {
	self->stepBy(static_cast<int>(steps));
}
bool QDateTimeEdit_Event(QDateTimeEdit* self, QEvent* event) {
	return self->event(event);
}
void QDateTimeEdit_DateTimeChanged(QDateTimeEdit* self, QDateTime* dateTime) {
	self->dateTimeChanged(*dateTime);
}
void QDateTimeEdit_connect_DateTimeChanged(QDateTimeEdit* self, intptr_t slot) {
	MiqtVirtualQDateTimeEdit::connect(self, static_cast<void (QDateTimeEdit::*)(const QDateTime&)>(&QDateTimeEdit::dateTimeChanged), self, [=](const QDateTime& dateTime) {
		const QDateTime& dateTime_ret = dateTime;
		// Cast returned reference into pointer
		QDateTime* sigval1 = const_cast<QDateTime*>(&dateTime_ret);
		miqt_exec_callback_QDateTimeEdit_DateTimeChanged(slot, sigval1);
	});
}
void QDateTimeEdit_TimeChanged(QDateTimeEdit* self, QTime* time) {
	self->timeChanged(*time);
}
void QDateTimeEdit_connect_TimeChanged(QDateTimeEdit* self, intptr_t slot) {
	MiqtVirtualQDateTimeEdit::connect(self, static_cast<void (QDateTimeEdit::*)(QTime)>(&QDateTimeEdit::timeChanged), self, [=](QTime time) {
		QTime* sigval1 = new QTime(time);
		miqt_exec_callback_QDateTimeEdit_TimeChanged(slot, sigval1);
	});
}
void QDateTimeEdit_DateChanged(QDateTimeEdit* self, QDate* date) {
	self->dateChanged(*date);
}
void QDateTimeEdit_connect_DateChanged(QDateTimeEdit* self, intptr_t slot) {
	MiqtVirtualQDateTimeEdit::connect(self, static_cast<void (QDateTimeEdit::*)(QDate)>(&QDateTimeEdit::dateChanged), self, [=](QDate date) {
		QDate* sigval1 = new QDate(date);
		miqt_exec_callback_QDateTimeEdit_DateChanged(slot, sigval1);
	});
}
void QDateTimeEdit_SetDateTime(QDateTimeEdit* self, QDateTime* dateTime) {
	self->setDateTime(*dateTime);
}
void QDateTimeEdit_SetDate(QDateTimeEdit* self, QDate* date) {
	self->setDate(*date);
}
void QDateTimeEdit_SetTime(QDateTimeEdit* self, QTime* time) {
	self->setTime(*time);
}
struct miqt_string QDateTimeEdit_Tr2(const char* s, const char* c) {
	QString _ret = QDateTimeEdit::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QDateTimeEdit_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDateTimeEdit::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QDateTimeEdit_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDateTimeEdit*>( (QDateTimeEdit*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QDateTimeEdit_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQDateTimeEdit*)(self) )->virtualbase_MetaObject();
}
void QDateTimeEdit_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDateTimeEdit*>( (QDateTimeEdit*)(self) )->handle__Metacast = slot;
}
void* QDateTimeEdit_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQDateTimeEdit*)(self) )->virtualbase_Metacast(param1);
}
void QDateTimeEdit_Delete(QDateTimeEdit* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDateTimeEdit*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQTimeEdit : public virtual QTimeEdit {
public:
MiqtVirtualQTimeEdit(QWidget* parent): QTimeEdit(parent) {};
MiqtVirtualQTimeEdit(): QTimeEdit() {};
MiqtVirtualQTimeEdit(QTime time): QTimeEdit(time) {};
MiqtVirtualQTimeEdit(QTime time, QWidget* parent): QTimeEdit(time, parent) {};

virtual ~MiqtVirtualQTimeEdit() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QTimeEdit::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QTimeEdit_MetaObject(const_cast<MiqtVirtualQTimeEdit*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QTimeEdit::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QTimeEdit::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QTimeEdit_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QTimeEdit::qt_metacast(param1);

	}
};
QTimeEdit* QTimeEdit_new(QWidget* parent) {
return new MiqtVirtualQTimeEdit(parent);
}
QTimeEdit* QTimeEdit_new2() {
return new MiqtVirtualQTimeEdit();
}
QTimeEdit* QTimeEdit_new3(QTime* time) {
return new MiqtVirtualQTimeEdit(*time);
}
QTimeEdit* QTimeEdit_new4(QTime* time, QWidget* parent) {
return new MiqtVirtualQTimeEdit(*time, parent);
}
void QTimeEdit_virtbase(QTimeEdit* src
, QDateTimeEdit** outptr_QDateTimeEdit
) {
*outptr_QDateTimeEdit = static_cast<QDateTimeEdit*>(src);
}
QMetaObject* QTimeEdit_MetaObject(const QTimeEdit* self) {
	return (QMetaObject*) self->metaObject();
}
void* QTimeEdit_Metacast(QTimeEdit* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QTimeEdit_Tr(const char* s) {
	QString _ret = QTimeEdit::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTimeEdit_UserTimeChanged(QTimeEdit* self, QTime* time) {
	self->userTimeChanged(*time);
}
void QTimeEdit_connect_UserTimeChanged(QTimeEdit* self, intptr_t slot) {
	MiqtVirtualQTimeEdit::connect(self, static_cast<void (QTimeEdit::*)(QTime)>(&QTimeEdit::userTimeChanged), self, [=](QTime time) {
		QTime* sigval1 = new QTime(time);
		miqt_exec_callback_QTimeEdit_UserTimeChanged(slot, sigval1);
	});
}
struct miqt_string QTimeEdit_Tr2(const char* s, const char* c) {
	QString _ret = QTimeEdit::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QTimeEdit_Tr3(const char* s, const char* c, int n) {
	QString _ret = QTimeEdit::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QTimeEdit_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTimeEdit*>( (QTimeEdit*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QTimeEdit_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQTimeEdit*)(self) )->virtualbase_MetaObject();
}
void QTimeEdit_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQTimeEdit*>( (QTimeEdit*)(self) )->handle__Metacast = slot;
}
void* QTimeEdit_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQTimeEdit*)(self) )->virtualbase_Metacast(param1);
}
void QTimeEdit_Delete(QTimeEdit* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQTimeEdit*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQDateEdit : public virtual QDateEdit {
public:
MiqtVirtualQDateEdit(QWidget* parent): QDateEdit(parent) {};
MiqtVirtualQDateEdit(): QDateEdit() {};
MiqtVirtualQDateEdit(QDate date): QDateEdit(date) {};
MiqtVirtualQDateEdit(QDate date, QWidget* parent): QDateEdit(date, parent) {};

virtual ~MiqtVirtualQDateEdit() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QDateEdit::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QDateEdit_MetaObject(const_cast<MiqtVirtualQDateEdit*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QDateEdit::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QDateEdit::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QDateEdit_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QDateEdit::qt_metacast(param1);

	}
};
QDateEdit* QDateEdit_new(QWidget* parent) {
return new MiqtVirtualQDateEdit(parent);
}
QDateEdit* QDateEdit_new2() {
return new MiqtVirtualQDateEdit();
}
QDateEdit* QDateEdit_new3(QDate* date) {
return new MiqtVirtualQDateEdit(*date);
}
QDateEdit* QDateEdit_new4(QDate* date, QWidget* parent) {
return new MiqtVirtualQDateEdit(*date, parent);
}
void QDateEdit_virtbase(QDateEdit* src
, QDateTimeEdit** outptr_QDateTimeEdit
) {
*outptr_QDateTimeEdit = static_cast<QDateTimeEdit*>(src);
}
QMetaObject* QDateEdit_MetaObject(const QDateEdit* self) {
	return (QMetaObject*) self->metaObject();
}
void* QDateEdit_Metacast(QDateEdit* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QDateEdit_Tr(const char* s) {
	QString _ret = QDateEdit::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QDateEdit_UserDateChanged(QDateEdit* self, QDate* date) {
	self->userDateChanged(*date);
}
void QDateEdit_connect_UserDateChanged(QDateEdit* self, intptr_t slot) {
	MiqtVirtualQDateEdit::connect(self, static_cast<void (QDateEdit::*)(QDate)>(&QDateEdit::userDateChanged), self, [=](QDate date) {
		QDate* sigval1 = new QDate(date);
		miqt_exec_callback_QDateEdit_UserDateChanged(slot, sigval1);
	});
}
struct miqt_string QDateEdit_Tr2(const char* s, const char* c) {
	QString _ret = QDateEdit::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QDateEdit_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDateEdit::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QDateEdit_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDateEdit*>( (QDateEdit*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QDateEdit_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQDateEdit*)(self) )->virtualbase_MetaObject();
}
void QDateEdit_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDateEdit*>( (QDateEdit*)(self) )->handle__Metacast = slot;
}
void* QDateEdit_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQDateEdit*)(self) )->virtualbase_Metacast(param1);
}
void QDateEdit_Delete(QDateEdit* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDateEdit*>( self );
	} else {
		delete self;
	}
}
