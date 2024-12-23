#include <QUrl>
#include <QWebEngineQuotaRequest>
#include <qwebenginequotarequest.h>
#include "gen_qwebenginequotarequest.h"

#ifndef _Bool
#define _Bool bool
#endif
#include "_cgo_export.h"

QWebEngineQuotaRequest* QWebEngineQuotaRequest_new() {
	return new QWebEngineQuotaRequest();
}

QWebEngineQuotaRequest* QWebEngineQuotaRequest_new2(QWebEngineQuotaRequest* param1) {
	return new QWebEngineQuotaRequest(*param1);
}

void QWebEngineQuotaRequest_Accept(QWebEngineQuotaRequest* self) {
	self->accept();
}

void QWebEngineQuotaRequest_Reject(QWebEngineQuotaRequest* self) {
	self->reject();
}

QUrl* QWebEngineQuotaRequest_Origin(const QWebEngineQuotaRequest* self) {
	return new QUrl(self->origin());
}

long long QWebEngineQuotaRequest_RequestedSize(const QWebEngineQuotaRequest* self) {
	qint64 _ret = self->requestedSize();
	return static_cast<long long>(_ret);
}

bool QWebEngineQuotaRequest_OperatorEqual(const QWebEngineQuotaRequest* self, QWebEngineQuotaRequest* that) {
	return (*self == *that);
}

bool QWebEngineQuotaRequest_OperatorNotEqual(const QWebEngineQuotaRequest* self, QWebEngineQuotaRequest* that) {
	return (*self != *that);
}

void QWebEngineQuotaRequest_Delete(QWebEngineQuotaRequest* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QWebEngineQuotaRequest*>( self );
	} else {
		delete self;
	}
}

