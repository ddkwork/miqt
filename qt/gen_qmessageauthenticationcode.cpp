// +build ignore

#include <QByteArray>
#include <QByteArrayView>
#include <QIODevice>
#include <QMessageAuthenticationCode>
#include <qmessageauthenticationcode.h>
#include "gen_qmessageauthenticationcode.h"

#ifndef _Bool
#define _Bool bool
#endif

QMessageAuthenticationCode* QMessageAuthenticationCode_new(int method) {
	return new QMessageAuthenticationCode(static_cast<QCryptographicHash::Algorithm>(method));
}

QMessageAuthenticationCode* QMessageAuthenticationCode_new2(int method, QByteArrayView* key) {
	return new QMessageAuthenticationCode(static_cast<QCryptographicHash::Algorithm>(method), *key);
}

void QMessageAuthenticationCode_Swap(QMessageAuthenticationCode* self, QMessageAuthenticationCode* other) {
	self->swap(*other);
}

void QMessageAuthenticationCode_Reset(QMessageAuthenticationCode* self) {
	self->reset();
}

void QMessageAuthenticationCode_SetKey(QMessageAuthenticationCode* self, QByteArrayView* key) {
	self->setKey(*key);
}

void QMessageAuthenticationCode_AddData(QMessageAuthenticationCode* self, const char* data, ptrdiff_t length) {
	self->addData(data, (qsizetype)(length));
}

void QMessageAuthenticationCode_AddDataWithData(QMessageAuthenticationCode* self, QByteArrayView* data) {
	self->addData(*data);
}

bool QMessageAuthenticationCode_AddDataWithDevice(QMessageAuthenticationCode* self, QIODevice* device) {
	return self->addData(device);
}

QByteArrayView* QMessageAuthenticationCode_ResultView(const QMessageAuthenticationCode* self) {
	return new QByteArrayView(self->resultView());
}

struct miqt_string QMessageAuthenticationCode_Result(const QMessageAuthenticationCode* self) {
	QByteArray _qb = self->result();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_string QMessageAuthenticationCode_Hash(QByteArrayView* message, QByteArrayView* key, int method) {
	QByteArray _qb = QMessageAuthenticationCode::hash(*message, *key, static_cast<QCryptographicHash::Algorithm>(method));
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QByteArrayView* QMessageAuthenticationCode_HashInto(QSpan<char> buffer, QByteArrayView* message, QByteArrayView* key, int method) {
	return new QByteArrayView(QMessageAuthenticationCode::hashInto(buffer, *message, *key, static_cast<QCryptographicHash::Algorithm>(method)));
}

QByteArrayView* QMessageAuthenticationCode_HashInto2(QSpan<uchar> buffer, QByteArrayView* message, QByteArrayView* key, int method) {
	return new QByteArrayView(QMessageAuthenticationCode::hashInto(buffer, *message, *key, static_cast<QCryptographicHash::Algorithm>(method)));
}

QByteArrayView* QMessageAuthenticationCode_HashInto3(QSpan<std::byte> buffer, QByteArrayView* message, QByteArrayView* key, int method) {
	return new QByteArrayView(QMessageAuthenticationCode::hashInto(buffer, *message, *key, static_cast<QCryptographicHash::Algorithm>(method)));
}

QByteArrayView* QMessageAuthenticationCode_HashInto4(QSpan<char> buffer, QSpan<const QByteArrayView> messageParts, QByteArrayView* key, int method) {
	return new QByteArrayView(QMessageAuthenticationCode::hashInto(buffer, messageParts, *key, static_cast<QCryptographicHash::Algorithm>(method)));
}

QByteArrayView* QMessageAuthenticationCode_HashInto5(QSpan<uchar> buffer, QSpan<const QByteArrayView> messageParts, QByteArrayView* key, int method) {
	return new QByteArrayView(QMessageAuthenticationCode::hashInto(buffer, messageParts, *key, static_cast<QCryptographicHash::Algorithm>(method)));
}

QByteArrayView* QMessageAuthenticationCode_HashInto6(QSpan<std::byte> buffer, QSpan<const QByteArrayView> message, QByteArrayView* key, int method) {
	return new QByteArrayView(QMessageAuthenticationCode::hashInto(buffer, message, *key, static_cast<QCryptographicHash::Algorithm>(method)));
}

void QMessageAuthenticationCode_Delete(QMessageAuthenticationCode* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMessageAuthenticationCode*>( self );
	} else {
		delete self;
	}
}

