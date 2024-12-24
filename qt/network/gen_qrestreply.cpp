// +build ignore

#include <QByteArray>
#include <QNetworkReply>
#include <QRestReply>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qrestreply.h>
#include "gen_qrestreply.h"
QRestReply* QRestReply_new(QNetworkReply* reply) {
return new QRestReply(reply);
}
void QRestReply_Swap(QRestReply* self, QRestReply* other) {
	self->swap(*other);
}
QNetworkReply* QRestReply_NetworkReply(const QRestReply* self) {
	return self->networkReply();
}
struct miqt_string QRestReply_ReadBody(QRestReply* self) {
	QByteArray _qb = self->readBody();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
struct miqt_string QRestReply_ReadText(QRestReply* self) {
	QString _ret = self->readText();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
bool QRestReply_IsSuccess(const QRestReply* self) {
	return self->isSuccess();
}
int QRestReply_HttpStatus(const QRestReply* self) {
	return self->httpStatus();
}
bool QRestReply_IsHttpStatusSuccess(const QRestReply* self) {
	return self->isHttpStatusSuccess();
}
bool QRestReply_HasError(const QRestReply* self) {
	return self->hasError();
}
int QRestReply_Error(const QRestReply* self) {
	QNetworkReply::NetworkError _ret = self->error();
	return static_cast<int>(_ret);
}
struct miqt_string QRestReply_ErrorString(const QRestReply* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QRestReply_Delete(QRestReply* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QRestReply*>( self );
	} else {
		delete self;
	}
}
