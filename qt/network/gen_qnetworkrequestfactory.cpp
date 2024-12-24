// +build ignore

#include <QByteArray>
#include <QHttpHeaders>
#include <QNetworkRequest>
#include <QNetworkRequestFactory>
#include <QSslConfiguration>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUrl>
#include <QUrlQuery>
#include <QVariant>
#include <qnetworkrequestfactory.h>
#include "gen_qnetworkrequestfactory.h"
QNetworkRequestFactory* QNetworkRequestFactory_new() {
return new QNetworkRequestFactory();
}
QNetworkRequestFactory* QNetworkRequestFactory_new2(QUrl* baseUrl) {
return new QNetworkRequestFactory(*baseUrl);
}
QNetworkRequestFactory* QNetworkRequestFactory_new3(QNetworkRequestFactory* other) {
return new QNetworkRequestFactory(*other);
}
void QNetworkRequestFactory_OperatorAssign(QNetworkRequestFactory* self, QNetworkRequestFactory* other) {
	self->operator=(*other);
}
void QNetworkRequestFactory_Swap(QNetworkRequestFactory* self, QNetworkRequestFactory* other) {
	self->swap(*other);
}
QUrl* QNetworkRequestFactory_BaseUrl(const QNetworkRequestFactory* self) {
	return new QUrl(self->baseUrl());
}
void QNetworkRequestFactory_SetBaseUrl(QNetworkRequestFactory* self, QUrl* url) {
	self->setBaseUrl(*url);
}
QSslConfiguration* QNetworkRequestFactory_SslConfiguration(const QNetworkRequestFactory* self) {
	return new QSslConfiguration(self->sslConfiguration());
}
void QNetworkRequestFactory_SetSslConfiguration(QNetworkRequestFactory* self, QSslConfiguration* configuration) {
	self->setSslConfiguration(*configuration);
}
QNetworkRequest* QNetworkRequestFactory_CreateRequest(const QNetworkRequestFactory* self) {
	return new QNetworkRequest(self->createRequest());
}
QNetworkRequest* QNetworkRequestFactory_CreateRequestWithQuery(const QNetworkRequestFactory* self, QUrlQuery* query) {
	return new QNetworkRequest(self->createRequest(*query));
}
QNetworkRequest* QNetworkRequestFactory_CreateRequestWithPath(const QNetworkRequestFactory* self, struct miqt_string path) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	return new QNetworkRequest(self->createRequest(path_QString));
}
QNetworkRequest* QNetworkRequestFactory_CreateRequest2(const QNetworkRequestFactory* self, struct miqt_string path, QUrlQuery* query) {
	QString path_QString = QString::fromUtf8(path.data, path.len);
	return new QNetworkRequest(self->createRequest(path_QString, *query));
}
void QNetworkRequestFactory_SetCommonHeaders(QNetworkRequestFactory* self, QHttpHeaders* headers) {
	self->setCommonHeaders(*headers);
}
QHttpHeaders* QNetworkRequestFactory_CommonHeaders(const QNetworkRequestFactory* self) {
	return new QHttpHeaders(self->commonHeaders());
}
void QNetworkRequestFactory_ClearCommonHeaders(QNetworkRequestFactory* self) {
	self->clearCommonHeaders();
}
struct miqt_string QNetworkRequestFactory_BearerToken(const QNetworkRequestFactory* self) {
	QByteArray _qb = self->bearerToken();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
void QNetworkRequestFactory_SetBearerToken(QNetworkRequestFactory* self, struct miqt_string token) {
	QByteArray token_QByteArray(token.data, token.len);
	self->setBearerToken(token_QByteArray);
}
void QNetworkRequestFactory_ClearBearerToken(QNetworkRequestFactory* self) {
	self->clearBearerToken();
}
struct miqt_string QNetworkRequestFactory_UserName(const QNetworkRequestFactory* self) {
	QString _ret = self->userName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QNetworkRequestFactory_SetUserName(QNetworkRequestFactory* self, struct miqt_string userName) {
	QString userName_QString = QString::fromUtf8(userName.data, userName.len);
	self->setUserName(userName_QString);
}
void QNetworkRequestFactory_ClearUserName(QNetworkRequestFactory* self) {
	self->clearUserName();
}
struct miqt_string QNetworkRequestFactory_Password(const QNetworkRequestFactory* self) {
	QString _ret = self->password();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QNetworkRequestFactory_SetPassword(QNetworkRequestFactory* self, struct miqt_string password) {
	QString password_QString = QString::fromUtf8(password.data, password.len);
	self->setPassword(password_QString);
}
void QNetworkRequestFactory_ClearPassword(QNetworkRequestFactory* self) {
	self->clearPassword();
}
QUrlQuery* QNetworkRequestFactory_QueryParameters(const QNetworkRequestFactory* self) {
	return new QUrlQuery(self->queryParameters());
}
void QNetworkRequestFactory_SetQueryParameters(QNetworkRequestFactory* self, QUrlQuery* query) {
	self->setQueryParameters(*query);
}
void QNetworkRequestFactory_ClearQueryParameters(QNetworkRequestFactory* self) {
	self->clearQueryParameters();
}
void QNetworkRequestFactory_SetPriority(QNetworkRequestFactory* self, int priority) {
	self->setPriority(static_cast<QNetworkRequest::Priority>(priority));
}
int QNetworkRequestFactory_Priority(const QNetworkRequestFactory* self) {
	QNetworkRequest::Priority _ret = self->priority();
	return static_cast<int>(_ret);
}
QVariant* QNetworkRequestFactory_Attribute(const QNetworkRequestFactory* self, int attribute) {
	return new QVariant(self->attribute(static_cast<QNetworkRequest::Attribute>(attribute)));
}
QVariant* QNetworkRequestFactory_Attribute2(const QNetworkRequestFactory* self, int attribute, QVariant* defaultValue) {
	return new QVariant(self->attribute(static_cast<QNetworkRequest::Attribute>(attribute), *defaultValue));
}
void QNetworkRequestFactory_SetAttribute(QNetworkRequestFactory* self, int attribute, QVariant* value) {
	self->setAttribute(static_cast<QNetworkRequest::Attribute>(attribute), *value);
}
void QNetworkRequestFactory_ClearAttribute(QNetworkRequestFactory* self, int attribute) {
	self->clearAttribute(static_cast<QNetworkRequest::Attribute>(attribute));
}
void QNetworkRequestFactory_ClearAttributes(QNetworkRequestFactory* self) {
	self->clearAttributes();
}
void QNetworkRequestFactory_Delete(QNetworkRequestFactory* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QNetworkRequestFactory*>( self );
	} else {
		delete self;
	}
}
