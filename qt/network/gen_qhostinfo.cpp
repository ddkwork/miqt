#include <QHostAddress>
#include <QHostInfo>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qhostinfo.h>
#include "gen_qhostinfo.h"
#include "_cgo_export.h"

void QHostInfo_new(QHostInfo** outptr_QHostInfo) {
	QHostInfo* ret = new QHostInfo();
	*outptr_QHostInfo = ret;
}

void QHostInfo_new2(QHostInfo* d, QHostInfo** outptr_QHostInfo) {
	QHostInfo* ret = new QHostInfo(*d);
	*outptr_QHostInfo = ret;
}

void QHostInfo_new3(int lookupId, QHostInfo** outptr_QHostInfo) {
	QHostInfo* ret = new QHostInfo(static_cast<int>(lookupId));
	*outptr_QHostInfo = ret;
}

void QHostInfo_OperatorAssign(QHostInfo* self, QHostInfo* d) {
	self->operator=(*d);
}

void QHostInfo_Swap(QHostInfo* self, QHostInfo* other) {
	self->swap(*other);
}

struct miqt_string QHostInfo_HostName(const QHostInfo* self) {
	QString _ret = self->hostName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QHostInfo_SetHostName(QHostInfo* self, struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	self->setHostName(name_QString);
}

struct miqt_array /* of QHostAddress* */  QHostInfo_Addresses(const QHostInfo* self) {
	QList<QHostAddress> _ret = self->addresses();
	// Convert QList<> from C++ memory to manually-managed C memory
	QHostAddress** _arr = static_cast<QHostAddress**>(malloc(sizeof(QHostAddress*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QHostAddress(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

void QHostInfo_SetAddresses(QHostInfo* self, struct miqt_array /* of QHostAddress* */  addresses) {
	QList<QHostAddress> addresses_QList;
	addresses_QList.reserve(addresses.len);
	QHostAddress** addresses_arr = static_cast<QHostAddress**>(addresses.data);
	for(size_t i = 0; i < addresses.len; ++i) {
		addresses_QList.push_back(*(addresses_arr[i]));
	}
	self->setAddresses(addresses_QList);
}

int QHostInfo_Error(const QHostInfo* self) {
	QHostInfo::HostInfoError _ret = self->error();
	return static_cast<int>(_ret);
}

void QHostInfo_SetError(QHostInfo* self, int error) {
	self->setError(static_cast<QHostInfo::HostInfoError>(error));
}

struct miqt_string QHostInfo_ErrorString(const QHostInfo* self) {
	QString _ret = self->errorString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QHostInfo_SetErrorString(QHostInfo* self, struct miqt_string errorString) {
	QString errorString_QString = QString::fromUtf8(errorString.data, errorString.len);
	self->setErrorString(errorString_QString);
}

void QHostInfo_SetLookupId(QHostInfo* self, int id) {
	self->setLookupId(static_cast<int>(id));
}

int QHostInfo_LookupId(const QHostInfo* self) {
	return self->lookupId();
}

void QHostInfo_AbortHostLookup(int lookupId) {
	QHostInfo::abortHostLookup(static_cast<int>(lookupId));
}

QHostInfo* QHostInfo_FromName(struct miqt_string name) {
	QString name_QString = QString::fromUtf8(name.data, name.len);
	return new QHostInfo(QHostInfo::fromName(name_QString));
}

struct miqt_string QHostInfo_LocalHostName() {
	QString _ret = QHostInfo::localHostName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QHostInfo_LocalDomainName() {
	QString _ret = QHostInfo::localDomainName();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QHostInfo_Delete(QHostInfo* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHostInfo*>( self );
	} else {
		delete self;
	}
}

