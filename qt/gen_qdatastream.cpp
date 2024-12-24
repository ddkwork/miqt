// +build ignore

#include <QByteArray>
#include <QDataStream>
#include <QIODevice>
#include <QIODeviceBase>
#include <qdatastream.h>
#include "gen_qdatastream.h"

#ifndef _Bool
#define _Bool bool
#endif

QDataStream* QDataStream_new() {
	return new QDataStream();
}

QDataStream* QDataStream_new2(QIODevice* param1) {
	return new QDataStream(param1);
}

QDataStream* QDataStream_new3(struct miqt_string param1) {
	QByteArray param1_QByteArray(param1.data, param1.len);
	return new QDataStream(param1_QByteArray);
}

void QDataStream_virtbase(QDataStream* src, QIODeviceBase** outptr_QIODeviceBase) {
	*outptr_QIODeviceBase = static_cast<QIODeviceBase*>(src);
}

QIODevice* QDataStream_Device(const QDataStream* self) {
	return self->device();
}

void QDataStream_SetDevice(QDataStream* self, QIODevice* device) {
	self->setDevice(device);
}

bool QDataStream_AtEnd(const QDataStream* self) {
	return self->atEnd();
}

Status QDataStream_Status(const QDataStream* self) {
	return self->status();
}

void QDataStream_SetStatus(QDataStream* self, Status status) {
	self->setStatus(status);
}

void QDataStream_ResetStatus(QDataStream* self) {
	self->resetStatus();
}

FloatingPointPrecision QDataStream_FloatingPointPrecision(const QDataStream* self) {
	return self->floatingPointPrecision();
}

void QDataStream_SetFloatingPointPrecision(QDataStream* self, FloatingPointPrecision precision) {
	self->setFloatingPointPrecision(precision);
}

ByteOrder QDataStream_ByteOrder(const QDataStream* self) {
	return self->byteOrder();
}

void QDataStream_SetByteOrder(QDataStream* self, ByteOrder byteOrder) {
	self->setByteOrder(byteOrder);
}

int QDataStream_Version(const QDataStream* self) {
	return self->version();
}

void QDataStream_SetVersion(QDataStream* self, int version) {
	self->setVersion(static_cast<int>(version));
}

void QDataStream_OperatorShiftRight(QDataStream* self, char* i) {
	self->operator>>(static_cast<char&>(*i));
}

void QDataStream_OperatorShiftRightWithQint8(QDataStream* self, signed char* i) {
	self->operator>>(static_cast<qint8&>(*i));
}

void QDataStream_OperatorShiftRightWithQuint8(QDataStream* self, unsigned char* i) {
	self->operator>>(static_cast<quint8&>(*i));
}

void QDataStream_OperatorShiftRightWithQint16(QDataStream* self, int16_t* i) {
	self->operator>>(static_cast<qint16&>(*i));
}

void QDataStream_OperatorShiftRightWithQuint16(QDataStream* self, uint16_t* i) {
	self->operator>>(static_cast<quint16&>(*i));
}

void QDataStream_OperatorShiftRightWithQint32(QDataStream* self, int* i) {
	self->operator>>(static_cast<qint32&>(*i));
}

void QDataStream_OperatorShiftRightWithQuint32(QDataStream* self, unsigned int* i) {
	self->operator>>(static_cast<quint32&>(*i));
}

void QDataStream_OperatorShiftRightWithQint64(QDataStream* self, long long* i) {
	self->operator>>(static_cast<qint64&>(*i));
}

void QDataStream_OperatorShiftRightWithQuint64(QDataStream* self, unsigned long long* i) {
	self->operator>>(static_cast<quint64&>(*i));
}

void QDataStream_OperatorShiftRightWithBool(QDataStream* self, bool* i) {
	self->operator>>(*i);
}

void QDataStream_OperatorShiftRightWithFloat(QDataStream* self, float* f) {
	self->operator>>(static_cast<float&>(*f));
}

void QDataStream_OperatorShiftRightWithDouble(QDataStream* self, double* f) {
	self->operator>>(static_cast<double&>(*f));
}

void QDataStream_OperatorShiftRightWithStr(QDataStream* self, char* str) {
	self->operator>>(str);
}

void QDataStream_OperatorShiftLeft(QDataStream* self, char i) {
	self->operator<<(static_cast<char>(i));
}

void QDataStream_OperatorShiftLeftWithQint8(QDataStream* self, signed char i) {
	self->operator<<(static_cast<qint8>(i));
}

void QDataStream_OperatorShiftLeftWithQuint8(QDataStream* self, unsigned char i) {
	self->operator<<(static_cast<quint8>(i));
}

void QDataStream_OperatorShiftLeftWithQint16(QDataStream* self, int16_t i) {
	self->operator<<(static_cast<qint16>(i));
}

void QDataStream_OperatorShiftLeftWithQuint16(QDataStream* self, uint16_t i) {
	self->operator<<(static_cast<quint16>(i));
}

void QDataStream_OperatorShiftLeftWithQint32(QDataStream* self, int i) {
	self->operator<<(static_cast<qint32>(i));
}

void QDataStream_OperatorShiftLeftWithQuint32(QDataStream* self, unsigned int i) {
	self->operator<<(static_cast<quint32>(i));
}

void QDataStream_OperatorShiftLeftWithQint64(QDataStream* self, long long i) {
	self->operator<<(static_cast<qint64>(i));
}

void QDataStream_OperatorShiftLeftWithQuint64(QDataStream* self, unsigned long long i) {
	self->operator<<(static_cast<quint64>(i));
}

void QDataStream_OperatorShiftLeftWithFloat(QDataStream* self, float f) {
	self->operator<<(static_cast<float>(f));
}

void QDataStream_OperatorShiftLeftWithDouble(QDataStream* self, double f) {
	self->operator<<(static_cast<double>(f));
}

void QDataStream_OperatorShiftLeftWithStr(QDataStream* self, const char* str) {
	self->operator<<(str);
}

QDataStream* QDataStream_ReadBytes(QDataStream* self, char* param1, unsigned int* lenVal) {
	QDataStream& _ret = self->readBytes(param1, static_cast<uint&>(*lenVal));
	// Cast returned reference into pointer
	return &_ret;
}

QDataStream* QDataStream_ReadBytes2(QDataStream* self, char* param1, long long* lenVal) {
	QDataStream& _ret = self->readBytes(param1, static_cast<qint64&>(*lenVal));
	// Cast returned reference into pointer
	return &_ret;
}

long long QDataStream_ReadRawData(QDataStream* self, char* param1, long long lenVal) {
	qint64 _ret = self->readRawData(param1, static_cast<qint64>(lenVal));
	return static_cast<long long>(_ret);
}

void QDataStream_WriteBytes(QDataStream* self, const char* param1, long long lenVal) {
	self->writeBytes(param1, static_cast<qint64>(lenVal));
}

long long QDataStream_WriteRawData(QDataStream* self, const char* param1, long long lenVal) {
	qint64 _ret = self->writeRawData(param1, static_cast<qint64>(lenVal));
	return static_cast<long long>(_ret);
}

long long QDataStream_SkipRawData(QDataStream* self, long long lenVal) {
	qint64 _ret = self->skipRawData(static_cast<qint64>(lenVal));
	return static_cast<long long>(_ret);
}

void QDataStream_StartTransaction(QDataStream* self) {
	self->startTransaction();
}

bool QDataStream_CommitTransaction(QDataStream* self) {
	return self->commitTransaction();
}

void QDataStream_RollbackTransaction(QDataStream* self) {
	self->rollbackTransaction();
}

void QDataStream_AbortTransaction(QDataStream* self) {
	self->abortTransaction();
}

bool QDataStream_IsDeviceTransactionStarted(const QDataStream* self) {
	return self->isDeviceTransactionStarted();
}

void QDataStream_Delete(QDataStream* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QDataStream*>( self );
	} else {
		delete self;
	}
}

