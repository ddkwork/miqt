// +build ignore

#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWinEventNotifier>
#include <qwineventnotifier.h>
#include "gen_qwineventnotifier.h"
class MiqtVirtualQWinEventNotifier : public virtual QWinEventNotifier {
public:
MiqtVirtualQWinEventNotifier(): QWinEventNotifier() {};
MiqtVirtualQWinEventNotifier(HANDLE hEvent): QWinEventNotifier(hEvent) {};
MiqtVirtualQWinEventNotifier(QObject* parent): QWinEventNotifier(parent) {};
MiqtVirtualQWinEventNotifier(HANDLE hEvent, QObject* parent): QWinEventNotifier(hEvent, parent) {};

virtual ~MiqtVirtualQWinEventNotifier() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWinEventNotifier::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWinEventNotifier_MetaObject(const_cast<MiqtVirtualQWinEventNotifier*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWinEventNotifier::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWinEventNotifier::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWinEventNotifier_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWinEventNotifier::qt_metacast(param1);

	}
};
QWinEventNotifier* QWinEventNotifier_new() {
return new MiqtVirtualQWinEventNotifier();
}
QWinEventNotifier* QWinEventNotifier_new2(HANDLE hEvent) {
return new MiqtVirtualQWinEventNotifier(hEvent);
}
QWinEventNotifier* QWinEventNotifier_new3(QObject* parent) {
return new MiqtVirtualQWinEventNotifier(parent);
}
QWinEventNotifier* QWinEventNotifier_new4(HANDLE hEvent, QObject* parent) {
return new MiqtVirtualQWinEventNotifier(hEvent, parent);
}
void QWinEventNotifier_virtbase(QWinEventNotifier* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QWinEventNotifier_MetaObject(const QWinEventNotifier* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWinEventNotifier_Metacast(QWinEventNotifier* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWinEventNotifier_Tr(const char* s) {
	QString _ret = QWinEventNotifier::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWinEventNotifier_SetHandle(QWinEventNotifier* self, HANDLE hEvent) {
	self->setHandle(hEvent);
}
HANDLE QWinEventNotifier_Handle(const QWinEventNotifier* self) {
	return self->handle();
}
bool QWinEventNotifier_IsEnabled(const QWinEventNotifier* self) {
	return self->isEnabled();
}
void QWinEventNotifier_SetEnabled(QWinEventNotifier* self, bool enable) {
	self->setEnabled(enable);
}
struct miqt_string QWinEventNotifier_Tr2(const char* s, const char* c) {
	QString _ret = QWinEventNotifier::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWinEventNotifier_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWinEventNotifier::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWinEventNotifier_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWinEventNotifier_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_MetaObject();
}
void QWinEventNotifier_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWinEventNotifier*>( (QWinEventNotifier*)(self) )->handle__Metacast = slot;
}
void* QWinEventNotifier_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWinEventNotifier*)(self) )->virtualbase_Metacast(param1);
}
void QWinEventNotifier_Delete(QWinEventNotifier* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWinEventNotifier*>( self );
	} else {
		delete self;
	}
}
