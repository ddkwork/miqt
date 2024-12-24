// +build ignore

#include <QAbstractEventDispatcher>
#define WORKAROUND_INNER_CLASS_DEFINITION_QAbstractEventDispatcher__TimerInfo
#define WORKAROUND_INNER_CLASS_DEFINITION_QAbstractEventDispatcher__TimerInfoV2
#include <QAbstractEventDispatcherV2>
#include <QAbstractNativeEventFilter>
#include <QByteArray>
#include <QDeadlineTimer>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QSocketNotifier>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QThread>
#include <qabstracteventdispatcher.h>
#include "gen_qabstracteventdispatcher.h"

#ifndef _Bool
#define _Bool bool
#endif

void _GUID_Delete(_GUID* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<_GUID*>( self );
	} else {
		delete self;
	}
}

void type_info_Delete(type_info* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<type_info*>( self );
	} else {
		delete self;
	}
}

void QAbstractEventDispatcher_virtbase(QAbstractEventDispatcher* src, QObject** outptr_QObject) {
	*outptr_QObject = static_cast<QObject*>(src);
}

QMetaObject* QAbstractEventDispatcher_MetaObject(const QAbstractEventDispatcher* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractEventDispatcher_Metacast(QAbstractEventDispatcher* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractEventDispatcher_Tr(const char* s) {
	QString _ret = QAbstractEventDispatcher::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QAbstractEventDispatcher* QAbstractEventDispatcher_Instance() {
	return QAbstractEventDispatcher::instance();
}

bool QAbstractEventDispatcher_ProcessEvents(QAbstractEventDispatcher* self, int flags) {
	return self->processEvents(static_cast<QEventLoop::ProcessEventsFlags>(flags));
}

void QAbstractEventDispatcher_RegisterSocketNotifier(QAbstractEventDispatcher* self, QSocketNotifier* notifier) {
	self->registerSocketNotifier(notifier);
}

void QAbstractEventDispatcher_UnregisterSocketNotifier(QAbstractEventDispatcher* self, QSocketNotifier* notifier) {
	self->unregisterSocketNotifier(notifier);
}

int QAbstractEventDispatcher_RegisterTimer(QAbstractEventDispatcher* self, Duration interval, int timerType, QObject* object) {
	Qt::TimerId _ret = self->registerTimer(interval, static_cast<Qt::TimerType>(timerType), object);
	return static_cast<int>(_ret);
}

int QAbstractEventDispatcher_RegisterTimer2(QAbstractEventDispatcher* self, long long interval, int timerType, QObject* object) {
	return self->registerTimer(static_cast<qint64>(interval), static_cast<Qt::TimerType>(timerType), object);
}

void QAbstractEventDispatcher_RegisterTimer3(QAbstractEventDispatcher* self, int timerId, long long interval, int timerType, QObject* object) {
	self->registerTimer(static_cast<int>(timerId), static_cast<qint64>(interval), static_cast<Qt::TimerType>(timerType), object);
}

bool QAbstractEventDispatcher_UnregisterTimer(QAbstractEventDispatcher* self, int timerId) {
	return self->unregisterTimer(static_cast<int>(timerId));
}

bool QAbstractEventDispatcher_UnregisterTimers(QAbstractEventDispatcher* self, QObject* object) {
	return self->unregisterTimers(object);
}

struct miqt_array /* of TimerInfo */  QAbstractEventDispatcher_RegisteredTimers(const QAbstractEventDispatcher* self, QObject* object) {
	QList<TimerInfo> _ret = self->registeredTimers(object);
	// Convert QList<> from C++ memory to manually-managed C memory
	TimerInfo* _arr = static_cast<TimerInfo*>(malloc(sizeof(TimerInfo) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

int QAbstractEventDispatcher_RemainingTime(QAbstractEventDispatcher* self, int timerId) {
	return self->remainingTime(static_cast<int>(timerId));
}

void QAbstractEventDispatcher_RegisterTimer4(QAbstractEventDispatcher* self, int timerId, Duration interval, int timerType, QObject* object) {
	self->registerTimer(static_cast<Qt::TimerId>(timerId), interval, static_cast<Qt::TimerType>(timerType), object);
}

bool QAbstractEventDispatcher_UnregisterTimerWithTimerId(QAbstractEventDispatcher* self, int timerId) {
	return self->unregisterTimer(static_cast<Qt::TimerId>(timerId));
}

struct miqt_array /* of TimerInfoV2 */  QAbstractEventDispatcher_TimersForObject(const QAbstractEventDispatcher* self, QObject* object) {
	QList<TimerInfoV2> _ret = self->timersForObject(object);
	// Convert QList<> from C++ memory to manually-managed C memory
	TimerInfoV2* _arr = static_cast<TimerInfoV2*>(malloc(sizeof(TimerInfoV2) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

Duration QAbstractEventDispatcher_RemainingTimeWithTimerId(const QAbstractEventDispatcher* self, int timerId) {
	return self->remainingTime(static_cast<Qt::TimerId>(timerId));
}

void QAbstractEventDispatcher_WakeUp(QAbstractEventDispatcher* self) {
	self->wakeUp();
}

void QAbstractEventDispatcher_Interrupt(QAbstractEventDispatcher* self) {
	self->interrupt();
}

void QAbstractEventDispatcher_StartingUp(QAbstractEventDispatcher* self) {
	self->startingUp();
}

void QAbstractEventDispatcher_ClosingDown(QAbstractEventDispatcher* self) {
	self->closingDown();
}

void QAbstractEventDispatcher_InstallNativeEventFilter(QAbstractEventDispatcher* self, QAbstractNativeEventFilter* filterObj) {
	self->installNativeEventFilter(filterObj);
}

void QAbstractEventDispatcher_RemoveNativeEventFilter(QAbstractEventDispatcher* self, QAbstractNativeEventFilter* filterObj) {
	self->removeNativeEventFilter(filterObj);
}

bool QAbstractEventDispatcher_FilterNativeEvent(QAbstractEventDispatcher* self, struct miqt_string eventType, void* message, intptr_t* result) {
	QByteArray eventType_QByteArray(eventType.data, eventType.len);
	return self->filterNativeEvent(eventType_QByteArray, message, (qintptr*)(result));
}

void QAbstractEventDispatcher_AboutToBlock(QAbstractEventDispatcher* self) {
	self->aboutToBlock();
}

void QAbstractEventDispatcher_connect_AboutToBlock(QAbstractEventDispatcher* self, intptr_t slot) {
	QAbstractEventDispatcher::connect(self, static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::aboutToBlock), self, [=]() {
		miqt_exec_callback_QAbstractEventDispatcher_AboutToBlock(slot);
	});
}

void QAbstractEventDispatcher_Awake(QAbstractEventDispatcher* self) {
	self->awake();
}

void QAbstractEventDispatcher_connect_Awake(QAbstractEventDispatcher* self, intptr_t slot) {
	QAbstractEventDispatcher::connect(self, static_cast<void (QAbstractEventDispatcher::*)()>(&QAbstractEventDispatcher::awake), self, [=]() {
		miqt_exec_callback_QAbstractEventDispatcher_Awake(slot);
	});
}

struct miqt_string QAbstractEventDispatcher_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractEventDispatcher::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractEventDispatcher_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractEventDispatcher::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QAbstractEventDispatcher* QAbstractEventDispatcher_Instance1(QThread* thread) {
	return QAbstractEventDispatcher::instance(thread);
}

void QAbstractEventDispatcher_Delete(QAbstractEventDispatcher* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractEventDispatcher*>( self );
	} else {
		delete self;
	}
}

QAbstractEventDispatcherV2* QAbstractEventDispatcherV2_new() {
	return new QAbstractEventDispatcherV2();
}

QAbstractEventDispatcherV2* QAbstractEventDispatcherV2_new2(QObject* parent) {
	return new QAbstractEventDispatcherV2(parent);
}

void QAbstractEventDispatcherV2_virtbase(QAbstractEventDispatcherV2* src, QAbstractEventDispatcher** outptr_QAbstractEventDispatcher) {
	*outptr_QAbstractEventDispatcher = static_cast<QAbstractEventDispatcher*>(src);
}

QMetaObject* QAbstractEventDispatcherV2_MetaObject(const QAbstractEventDispatcherV2* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractEventDispatcherV2_Metacast(QAbstractEventDispatcherV2* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractEventDispatcherV2_Tr(const char* s) {
	QString _ret = QAbstractEventDispatcherV2::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractEventDispatcherV2_RegisterTimer(QAbstractEventDispatcherV2* self, int timerId, Duration interval, int timerType, QObject* object) {
	self->registerTimer(static_cast<Qt::TimerId>(timerId), interval, static_cast<Qt::TimerType>(timerType), object);
}

bool QAbstractEventDispatcherV2_UnregisterTimer(QAbstractEventDispatcherV2* self, int timerId) {
	return self->unregisterTimer(static_cast<Qt::TimerId>(timerId));
}

struct miqt_array /* of TimerInfoV2 */  QAbstractEventDispatcherV2_TimersForObject(const QAbstractEventDispatcherV2* self, QObject* object) {
	QList<TimerInfoV2> _ret = self->timersForObject(object);
	// Convert QList<> from C++ memory to manually-managed C memory
	TimerInfoV2* _arr = static_cast<TimerInfoV2*>(malloc(sizeof(TimerInfoV2) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

Duration QAbstractEventDispatcherV2_RemainingTime(const QAbstractEventDispatcherV2* self, int timerId) {
	return self->remainingTime(static_cast<Qt::TimerId>(timerId));
}

bool QAbstractEventDispatcherV2_ProcessEventsWithDeadline(QAbstractEventDispatcherV2* self, int flags, QDeadlineTimer* deadline) {
	return self->processEventsWithDeadline(static_cast<QEventLoop::ProcessEventsFlags>(flags), *deadline);
}

struct miqt_string QAbstractEventDispatcherV2_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractEventDispatcherV2::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractEventDispatcherV2_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractEventDispatcherV2::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractEventDispatcherV2_Delete(QAbstractEventDispatcherV2* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractEventDispatcherV2*>( self );
	} else {
		delete self;
	}
}

QAbstractEventDispatcher__TimerInfo* QAbstractEventDispatcher__TimerInfo_new(int id, int i, int t) {
	return new QAbstractEventDispatcher::TimerInfo(static_cast<int>(id), static_cast<int>(i), static_cast<Qt::TimerType>(t));
}

QAbstractEventDispatcher__TimerInfo* QAbstractEventDispatcher__TimerInfo_new2(const TimerInfo* param1) {
	return new QAbstractEventDispatcher::TimerInfo(*param1);
}

void QAbstractEventDispatcher__TimerInfo_Delete(QAbstractEventDispatcher__TimerInfo* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractEventDispatcher::TimerInfo*>( self );
	} else {
		delete self;
	}
}

QAbstractEventDispatcher__TimerInfoV2* QAbstractEventDispatcher__TimerInfoV2_new() {
	return new QAbstractEventDispatcher::TimerInfoV2();
}

QAbstractEventDispatcher__TimerInfoV2* QAbstractEventDispatcher__TimerInfoV2_new2(const TimerInfoV2* param1) {
	return new QAbstractEventDispatcher::TimerInfoV2(*param1);
}

void QAbstractEventDispatcher__TimerInfoV2_Delete(QAbstractEventDispatcher__TimerInfoV2* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAbstractEventDispatcher::TimerInfoV2*>( self );
	} else {
		delete self;
	}
}

