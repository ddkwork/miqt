// +build ignore

#include <QNativeIpcKey>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QSystemSemaphore>
#include <qsystemsemaphore.h>
#include "gen_qsystemsemaphore.h"

QSystemSemaphore* QSystemSemaphore_new(QNativeIpcKey* key) {
	return new QSystemSemaphore(*key);
}

QSystemSemaphore* QSystemSemaphore_new2(struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QSystemSemaphore(key_QString);
}

QSystemSemaphore* QSystemSemaphore_new3(QNativeIpcKey* key, int initialValue) {
	return new QSystemSemaphore(*key, static_cast<int>(initialValue));
}

QSystemSemaphore* QSystemSemaphore_new4(QNativeIpcKey* key, int initialValue, AccessMode param3) {
	return new QSystemSemaphore(*key, static_cast<int>(initialValue), param3);
}

QSystemSemaphore* QSystemSemaphore_new5(struct miqt_string key, int initialValue) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QSystemSemaphore(key_QString, static_cast<int>(initialValue));
}

QSystemSemaphore* QSystemSemaphore_new6(struct miqt_string key, int initialValue, AccessMode mode) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QSystemSemaphore(key_QString, static_cast<int>(initialValue), mode);
}

struct miqt_string QSystemSemaphore_Tr(const char* sourceText) {
	QString _ret = QSystemSemaphore::tr(sourceText);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSystemSemaphore_SetNativeKey(QSystemSemaphore* self, QNativeIpcKey* key) {
	self->setNativeKey(*key);
}

void QSystemSemaphore_SetNativeKeyWithKey(QSystemSemaphore* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setNativeKey(key_QString);
}

QNativeIpcKey* QSystemSemaphore_NativeIpcKey(const QSystemSemaphore* self) {
	return new QNativeIpcKey(self->nativeIpcKey());
}

void QSystemSemaphore_SetKey(QSystemSemaphore* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setKey(key_QString);
}

struct miqt_string QSystemSemaphore_Key(const QSystemSemaphore* self) {
	QString _ret = self->key();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QSystemSemaphore_Acquire(QSystemSemaphore* self) {
	return self->acquire();
}

bool QSystemSemaphore_Release(QSystemSemaphore* self) {
	return self->release();
}

SystemSemaphoreError QSystemSemaphore_Error(const QSystemSemaphore* self) {
	return self->error();
}

struct miqt_string QSystemSemaphore_ErrorString(const QSystemSemaphore* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QSystemSemaphore_IsKeyTypeSupported(uint16_t typeVal) {
	return QSystemSemaphore::isKeyTypeSupported(static_cast<QNativeIpcKey::Type>(typeVal));
}

QNativeIpcKey* QSystemSemaphore_PlatformSafeKey(struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSystemSemaphore::platformSafeKey(key_QString));
}

QNativeIpcKey* QSystemSemaphore_LegacyNativeKey(struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSystemSemaphore::legacyNativeKey(key_QString));
}

struct miqt_string QSystemSemaphore_Tr2(const char* sourceText, const char* disambiguation) {
	QString _ret = QSystemSemaphore::tr(sourceText, disambiguation);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QSystemSemaphore_Tr3(const char* sourceText, const char* disambiguation, int n) {
	QString _ret = QSystemSemaphore::tr(sourceText, disambiguation, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QSystemSemaphore_SetNativeKey2(QSystemSemaphore* self, QNativeIpcKey* key, int initialValue) {
	self->setNativeKey(*key, static_cast<int>(initialValue));
}

void QSystemSemaphore_SetNativeKey3(QSystemSemaphore* self, QNativeIpcKey* key, int initialValue, AccessMode param3) {
	self->setNativeKey(*key, static_cast<int>(initialValue), param3);
}

void QSystemSemaphore_SetNativeKey22(QSystemSemaphore* self, struct miqt_string key, int initialValue) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setNativeKey(key_QString, static_cast<int>(initialValue));
}

void QSystemSemaphore_SetNativeKey32(QSystemSemaphore* self, struct miqt_string key, int initialValue, AccessMode mode) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setNativeKey(key_QString, static_cast<int>(initialValue), mode);
}

void QSystemSemaphore_SetNativeKey4(QSystemSemaphore* self, struct miqt_string key, int initialValue, AccessMode mode, uint16_t typeVal) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setNativeKey(key_QString, static_cast<int>(initialValue), mode, static_cast<QNativeIpcKey::Type>(typeVal));
}

void QSystemSemaphore_SetKey2(QSystemSemaphore* self, struct miqt_string key, int initialValue) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setKey(key_QString, static_cast<int>(initialValue));
}

void QSystemSemaphore_SetKey3(QSystemSemaphore* self, struct miqt_string key, int initialValue, AccessMode mode) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->setKey(key_QString, static_cast<int>(initialValue), mode);
}

bool QSystemSemaphore_Release1(QSystemSemaphore* self, int n) {
	return self->release(static_cast<int>(n));
}

QNativeIpcKey* QSystemSemaphore_PlatformSafeKey2(struct miqt_string key, uint16_t typeVal) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSystemSemaphore::platformSafeKey(key_QString, static_cast<QNativeIpcKey::Type>(typeVal)));
}

QNativeIpcKey* QSystemSemaphore_LegacyNativeKey2(struct miqt_string key, uint16_t typeVal) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QNativeIpcKey(QSystemSemaphore::legacyNativeKey(key_QString, static_cast<QNativeIpcKey::Type>(typeVal)));
}

void QSystemSemaphore_Delete(QSystemSemaphore* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QSystemSemaphore*>( self );
	} else {
		delete self;
	}
}

