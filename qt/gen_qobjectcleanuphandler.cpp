// +build ignore

#include <QMetaObject>
#include <QObject>
#include <QObjectCleanupHandler>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qobjectcleanuphandler.h>
#include "gen_qobjectcleanuphandler.h"
class MiqtVirtualQObjectCleanupHandler : public virtual QObjectCleanupHandler {
public:
MiqtVirtualQObjectCleanupHandler(): QObjectCleanupHandler() {};

virtual ~MiqtVirtualQObjectCleanupHandler() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QObjectCleanupHandler::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QObjectCleanupHandler_MetaObject(const_cast<MiqtVirtualQObjectCleanupHandler*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QObjectCleanupHandler::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QObjectCleanupHandler::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QObjectCleanupHandler_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QObjectCleanupHandler::qt_metacast(param1);

	}
};
QObjectCleanupHandler* QObjectCleanupHandler_new() {
return new MiqtVirtualQObjectCleanupHandler();
}
void QObjectCleanupHandler_virtbase(QObjectCleanupHandler* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QObjectCleanupHandler_MetaObject(const QObjectCleanupHandler* self) {
	return (QMetaObject*) self->metaObject();
}
void* QObjectCleanupHandler_Metacast(QObjectCleanupHandler* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QObjectCleanupHandler_Tr(const char* s) {
	QString _ret = QObjectCleanupHandler::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QObject* QObjectCleanupHandler_Add(QObjectCleanupHandler* self, QObject* object) {
	return self->add(object);
}
void QObjectCleanupHandler_Remove(QObjectCleanupHandler* self, QObject* object) {
	self->remove(object);
}
bool QObjectCleanupHandler_IsEmpty(const QObjectCleanupHandler* self) {
	return self->isEmpty();
}
void QObjectCleanupHandler_Clear(QObjectCleanupHandler* self) {
	self->clear();
}
struct miqt_string QObjectCleanupHandler_Tr2(const char* s, const char* c) {
	QString _ret = QObjectCleanupHandler::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QObjectCleanupHandler_Tr3(const char* s, const char* c, int n) {
	QString _ret = QObjectCleanupHandler::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QObjectCleanupHandler_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQObjectCleanupHandler*>( (QObjectCleanupHandler*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QObjectCleanupHandler_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQObjectCleanupHandler*)(self) )->virtualbase_MetaObject();
}
void QObjectCleanupHandler_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQObjectCleanupHandler*>( (QObjectCleanupHandler*)(self) )->handle__Metacast = slot;
}
void* QObjectCleanupHandler_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQObjectCleanupHandler*)(self) )->virtualbase_Metacast(param1);
}
void QObjectCleanupHandler_Delete(QObjectCleanupHandler* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQObjectCleanupHandler*>( self );
	} else {
		delete self;
	}
}
