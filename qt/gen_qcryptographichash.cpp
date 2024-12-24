// +build ignore

#include <QByteArray>
#include <QByteArrayView>
#include <QCryptographicHash>
#include <QIODevice>
#include <qcryptographichash.h>
#include "gen_qcryptographichash.h"

QCryptographicHash* QCryptographicHash_new(Algorithm method) {
	return new QCryptographicHash(method);
}

void QCryptographicHash_Swap(QCryptographicHash* self, QCryptographicHash* other) {
	self->swap(*other);
}

void QCryptographicHash_Reset(QCryptographicHash* self) {
	self->reset();
}

Algorithm QCryptographicHash_Algorithm(const QCryptographicHash* self) {
	return self->algorithm();
}

void QCryptographicHash_AddData(QCryptographicHash* self, const char* data, ptrdiff_t length) {
	self->addData(data, (qsizetype)(length));
}

void QCryptographicHash_AddDataWithData(QCryptographicHash* self, QByteArrayView* data) {
	self->addData(*data);
}

bool QCryptographicHash_AddDataWithDevice(QCryptographicHash* self, QIODevice* device) {
	return self->addData(device);
}

struct miqt_string QCryptographicHash_Result(const QCryptographicHash* self) {
	QByteArray _qb = self->result();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QByteArrayView* QCryptographicHash_ResultView(const QCryptographicHash* self) {
	return new QByteArrayView(self->resultView());
}

struct miqt_string QCryptographicHash_Hash(QByteArrayView* data, Algorithm method) {
	QByteArray _qb = QCryptographicHash::hash(*data, method);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

QByteArrayView* QCryptographicHash_HashInto(QSpan<char> buffer, QByteArrayView* data, Algorithm method) {
	return new QByteArrayView(QCryptographicHash::hashInto(buffer, *data, method));
}

QByteArrayView* QCryptographicHash_HashInto2(QSpan<uchar> buffer, QByteArrayView* data, Algorithm method) {
	return new QByteArrayView(QCryptographicHash::hashInto(buffer, *data, method));
}

QByteArrayView* QCryptographicHash_HashInto3(QSpan<std::byte> buffer, QByteArrayView* data, Algorithm method) {
	return new QByteArrayView(QCryptographicHash::hashInto(buffer, *data, method));
}

QByteArrayView* QCryptographicHash_HashInto4(QSpan<char> buffer, QSpan<const QByteArrayView> data, Algorithm method) {
	return new QByteArrayView(QCryptographicHash::hashInto(buffer, data, method));
}

QByteArrayView* QCryptographicHash_HashInto5(QSpan<uchar> buffer, QSpan<const QByteArrayView> data, Algorithm method) {
	return new QByteArrayView(QCryptographicHash::hashInto(buffer, data, method));
}

QByteArrayView* QCryptographicHash_HashInto6(QSpan<std::byte> buffer, QSpan<const QByteArrayView> data, Algorithm method) {
	return new QByteArrayView(QCryptographicHash::hashInto(buffer, data, method));
}

int QCryptographicHash_HashLength(Algorithm method) {
	return QCryptographicHash::hashLength(method);
}

bool QCryptographicHash_SupportsAlgorithm(Algorithm method) {
	return QCryptographicHash::supportsAlgorithm(method);
}

void QCryptographicHash_Delete(QCryptographicHash* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCryptographicHash*>( self );
	} else {
		delete self;
	}
}

