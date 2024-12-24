// +build ignore

#include <QList>
#include <QMediaMetaData>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qmediametadata.h>
#include "gen_qmediametadata.h"
QMediaMetaData* QMediaMetaData_new(QMediaMetaData* param1) {
return new QMediaMetaData(*param1);
}
QMediaMetaData* QMediaMetaData_new2() {
return new QMediaMetaData();
}
QVariant* QMediaMetaData_Value(const QMediaMetaData* self, Key k) {
	return new QVariant(self->value(k));
}
void QMediaMetaData_Insert(QMediaMetaData* self, Key k, QVariant* value) {
	self->insert(k, *value);
}
void QMediaMetaData_Remove(QMediaMetaData* self, Key k) {
	self->remove(k);
}
struct miqt_array /* of Key */  QMediaMetaData_Keys(const QMediaMetaData* self) {
	QList<Key> _ret = self->keys();
	// Convert QList<> from C++ memory to manually-managed C memory
	Key* _arr = static_cast<Key*>(malloc(sizeof(Key) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
QVariant* QMediaMetaData_OperatorSubscript(QMediaMetaData* self, Key k) {
	QVariant& _ret = self->operator[](k);
	// Cast returned reference into pointer
	return &_ret;
}
void QMediaMetaData_Clear(QMediaMetaData* self) {
	self->clear();
}
bool QMediaMetaData_IsEmpty(const QMediaMetaData* self) {
	return self->isEmpty();
}
struct miqt_string QMediaMetaData_StringValue(const QMediaMetaData* self, Key k) {
	QString _ret = self->stringValue(k);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QMediaMetaData_MetaDataKeyToString(Key k) {
	QString _ret = QMediaMetaData::metaDataKeyToString(k);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QMediaMetaData_Delete(QMediaMetaData* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QMediaMetaData*>( self );
	} else {
		delete self;
	}
}
