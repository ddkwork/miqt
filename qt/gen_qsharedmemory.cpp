// +build ignore

#include <QMetaObject>
#include <QNativeIpcKey>
#include <QObject>
#include <QSharedMemory>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qsharedmemory.h>
#include "gen_qsharedmemory.h"
class MiqtVirtualQSharedMemory : public virtual QSharedMemory {
public:
MiqtVirtualQSharedMemory(): QSharedMemory() {};
MiqtVirtualQSharedMemory(const QNativeIpcKey& key): QSharedMemory(key) {};
MiqtVirtualQSharedMemory(const QString& key): QSharedMemory(key) {};
MiqtVirtualQSharedMemory(QObject* parent): QSharedMemory(parent) {};
MiqtVirtualQSharedMemory(const QNativeIpcKey& key, QObject* parent): QSharedMemory(key, parent) {};
MiqtVirtualQSharedMemory(const QString& key, QObject* parent): QSharedMemory(key, parent) {};

virtual ~MiqtVirtualQSharedMemory() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSharedMemory::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSharedMemory_MetaObject(const_cast<MiqtVirtualQSharedMemory*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSharedMemory::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSharedMemory::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSharedMemory_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSharedMemory::qt_metacast(param1);

	}
};
QSharedMemory* QSharedMemory_new() {
return new MiqtVirtualQSharedMemory();
}
QSharedMemory* QSharedMemory_new2(QNativeIpcKey* key) {
return new MiqtVirtualQSharedMemory(*key);
}
QSharedMemory* QSharedMemory_new3(struct miqt_string key) {
QString key_QString = QString::fromUtf8(key.data, key.len);
	return new MiqtVirtualQSharedMemory(key_QString);
}
QSharedMemory* QSharedMemory_new4(QObject* parent) {
return new MiqtVirtualQSharedMemory(parent);
}
QSharedMemory* QSharedMemory_new5(QNativeIpcKey* key, QObject* parent) {
return new MiqtVirtualQSharedMemory(*key, parent);
}
QSharedMemory* QSharedMemory_new6(struct miqt_string key, QObject* parent) {
QString key_QString = QString::fromUtf8(key.data, key.len);
	return new MiqtVirtualQSharedMemory(key_QString, parent);
}
void QSharedMemory_virtbase(QSharedMemory* src
, QObject** outptr_QObject
) {
*outptr_QObject = static_cast<QObject*>(src);
}
QMetaObject* QSharedMemory_MetaObject(const QSharedMemory* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSharedMemory_Metacast(QSharedMemory* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSharedMemory_Tr(const char* s) {
	QString _ret = QSharedMemory::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSharedMemory_SetKey(QSharedMemory* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setKey(key_QString);
}
struct miqt_string QSharedMemory_Key(const QSharedMemory* self) {
	QString _ret = self->key();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSharedMemory_SetNativeKey(QSharedMemory* self, QNativeIpcKey* key) {
	self->setNativeKey(*key);
}
void QSharedMemory_SetNativeKeyWithKey(QSharedMemory* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setNativeKey(key_QString);
}
struct miqt_string QSharedMemory_NativeKey(const QSharedMemory* self) {
	QString _ret = self->nativeKey();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QNativeIpcKey* QSharedMemory_NativeIpcKey(const QSharedMemory* self) {
	return new QNativeIpcKey(self->nativeIpcKey());
}
bool QSharedMemory_Create(QSharedMemory* self, ptrdiff_t size) {
	return self->create((qsizetype)(size));
}
ptrdiff_t QSharedMemory_Size(const QSharedMemory* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}
bool QSharedMemory_Attach(QSharedMemory* self) {
	return self->attach();
}
bool QSharedMemory_IsAttached(const QSharedMemory* self) {
	return self->isAttached();
}
bool QSharedMemory_Detach(QSharedMemory* self) {
	return self->detach();
}
void* QSharedMemory_Data(QSharedMemory* self) {
	return self->data();
}
const void* QSharedMemory_ConstData(const QSharedMemory* self) {
	return (const void*) self->constData();
}
const void* QSharedMemory_Data2(const QSharedMemory* self) {
	return (const void*) self->data();
}
bool QSharedMemory_Lock(QSharedMemory* self) {
	return self->lock();
}
bool QSharedMemory_Unlock(QSharedMemory* self) {
	return self->unlock();
}
SharedMemoryError QSharedMemory_Error(const QSharedMemory* self) {
	return self->error();
}
struct miqt_string QSharedMemory_ErrorString(const QSharedMemory* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QSharedMemory_IsKeyTypeSupported(uint16_t typeVal) {
	return QSharedMemory::isKeyTypeSupported(static_cast<QNativeIpcKey::Type>(typeVal));
}
QNativeIpcKey* QSharedMemory_PlatformSafeKey(struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSharedMemory::platformSafeKey(key_QString));
}
QNativeIpcKey* QSharedMemory_LegacyNativeKey(struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSharedMemory::legacyNativeKey(key_QString));
}
struct miqt_string QSharedMemory_Tr2(const char* s, const char* c) {
	QString _ret = QSharedMemory::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSharedMemory_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSharedMemory::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSharedMemory_SetNativeKey2(QSharedMemory* self, struct miqt_string key, uint16_t typeVal) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setNativeKey(key_QString, static_cast<QNativeIpcKey::Type>(typeVal));
}
bool QSharedMemory_Create2(QSharedMemory* self, ptrdiff_t size, AccessMode mode) {
	return self->create((qsizetype)(size), mode);
}
bool QSharedMemory_Attach1(QSharedMemory* self, AccessMode mode) {
	return self->attach(mode);
}
QNativeIpcKey* QSharedMemory_PlatformSafeKey2(struct miqt_string key, uint16_t typeVal) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSharedMemory::platformSafeKey(key_QString, static_cast<QNativeIpcKey::Type>(typeVal)));
}
QNativeIpcKey* QSharedMemory_LegacyNativeKey2(struct miqt_string key, uint16_t typeVal) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSharedMemory::legacyNativeKey(key_QString, static_cast<QNativeIpcKey::Type>(typeVal)));
}
void QSharedMemory_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSharedMemory*>( (QSharedMemory*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSharedMemory_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSharedMemory*)(self) )->virtualbase_MetaObject();
}
void QSharedMemory_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSharedMemory*>( (QSharedMemory*)(self) )->handle__Metacast = slot;
}
void* QSharedMemory_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSharedMemory*)(self) )->virtualbase_Metacast(param1);
}
void QSharedMemory_Delete(QSharedMemory* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSharedMemory*>( self );
	} else {
		delete self;
	}
}
