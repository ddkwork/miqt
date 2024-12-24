// +build ignore

#include <QJsonObject>
#define WORKAROUND_INNER_CLASS_DEFINITION_QJsonObject__const_iterator
#define WORKAROUND_INNER_CLASS_DEFINITION_QJsonObject__iterator
#include <QJsonValue>
#include <QJsonValueConstRef>
#include <QJsonValueRef>
#include <QList>
#include <QMap>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qjsonobject.h>
#include "gen_qjsonobject.h"

QJsonObject* QJsonObject_new() {
	return new QJsonObject();
}

QJsonObject* QJsonObject_new2(QJsonObject* other) {
	return new QJsonObject(*other);
}

void QJsonObject_OperatorAssign(QJsonObject* self, QJsonObject* other) {
	self->operator=(*other);
}

void QJsonObject_Swap(QJsonObject* self, QJsonObject* other) {
	self->swap(*other);
}

QJsonObject* QJsonObject_FromVariantMap(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal) {
	QVariantMap mapVal_QMap;
	struct miqt_string* mapVal_karr = static_cast<struct miqt_string*>(mapVal.keys);
	QVariant** mapVal_varr = static_cast<QVariant**>(mapVal.values);
	for(size_t i = 0; i < mapVal.len; ++i) {
		QString mapVal_karr_i_QString = QString::fromUtf8(mapVal_karr[i].data, mapVal_karr[i].len);
		mapVal_QMap[mapVal_karr_i_QString] = *(mapVal_varr[i]);
	}
	return new QJsonObject(QJsonObject::fromVariantMap(mapVal_QMap));
}

struct miqt_map /* of struct miqt_string to QVariant* */  QJsonObject_ToVariantMap(const QJsonObject* self) {
	QVariantMap _ret = self->toVariantMap();
	// Convert QMap<> from C++ memory to manually-managed C memory
	struct miqt_string* _karr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.size()));
	QVariant** _varr = static_cast<QVariant**>(malloc(sizeof(QVariant*) * _ret.size()));
	int _ctr = 0;
	for (auto _itr = _ret.keyValueBegin(); _itr != _ret.keyValueEnd(); ++_itr) {
		QString _mapkey_ret = _itr->first;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _mapkey_b = _mapkey_ret.toUtf8();
		struct miqt_string _mapkey_ms;
		_mapkey_ms.len = _mapkey_b.length();
		_mapkey_ms.data = static_cast<char*>(malloc(_mapkey_ms.len));
		memcpy(_mapkey_ms.data, _mapkey_b.data(), _mapkey_ms.len);
		_karr[_ctr] = _mapkey_ms;
		_varr[_ctr] = new QVariant(_itr->second);
		_ctr++;
	}
	struct miqt_map _out;
	_out.len = _ret.size();
	_out.keys = static_cast<void*>(_karr);
	_out.values = static_cast<void*>(_varr);
	return _out;
}

QJsonObject* QJsonObject_FromVariantHash(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal) {
	QVariantHash mapVal_QMap;
	mapVal_QMap.reserve(mapVal.len);
	struct miqt_string* mapVal_karr = static_cast<struct miqt_string*>(mapVal.keys);
	QVariant** mapVal_varr = static_cast<QVariant**>(mapVal.values);
	for(size_t i = 0; i < mapVal.len; ++i) {
		QString mapVal_karr_i_QString = QString::fromUtf8(mapVal_karr[i].data, mapVal_karr[i].len);
		mapVal_QMap[mapVal_karr_i_QString] = *(mapVal_varr[i]);
	}
	return new QJsonObject(QJsonObject::fromVariantHash(mapVal_QMap));
}

struct miqt_map /* of struct miqt_string to QVariant* */  QJsonObject_ToVariantHash(const QJsonObject* self) {
	QVariantHash _ret = self->toVariantHash();
	// Convert QMap<> from C++ memory to manually-managed C memory
	struct miqt_string* _karr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.size()));
	QVariant** _varr = static_cast<QVariant**>(malloc(sizeof(QVariant*) * _ret.size()));
	int _ctr = 0;
	for (auto _itr = _ret.keyValueBegin(); _itr != _ret.keyValueEnd(); ++_itr) {
		QString _hashkey_ret = _itr->first;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _hashkey_b = _hashkey_ret.toUtf8();
		struct miqt_string _hashkey_ms;
		_hashkey_ms.len = _hashkey_b.length();
		_hashkey_ms.data = static_cast<char*>(malloc(_hashkey_ms.len));
		memcpy(_hashkey_ms.data, _hashkey_b.data(), _hashkey_ms.len);
		_karr[_ctr] = _hashkey_ms;
		_varr[_ctr] = new QVariant(_itr->second);
		_ctr++;
	}
	struct miqt_map _out;
	_out.len = _ret.size();
	_out.keys = static_cast<void*>(_karr);
	_out.values = static_cast<void*>(_varr);
	return _out;
}

struct miqt_array /* of struct miqt_string */  QJsonObject_Keys(const QJsonObject* self) {
	QStringList _ret = self->keys();
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QString _lv_ret = _ret[i];
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray _lv_b = _lv_ret.toUtf8();
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_b.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_b.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

ptrdiff_t QJsonObject_Size(const QJsonObject* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QJsonObject_Count(const QJsonObject* self) {
	qsizetype _ret = self->count();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QJsonObject_Length(const QJsonObject* self) {
	qsizetype _ret = self->length();
	return static_cast<ptrdiff_t>(_ret);
}

bool QJsonObject_IsEmpty(const QJsonObject* self) {
	return self->isEmpty();
}

QJsonValue* QJsonObject_Value(const QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QJsonValue(self->value(key_QString));
}

QJsonValue* QJsonObject_OperatorSubscript(const QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QJsonValue(self->operator[](key_QString));
}

QJsonValueRef* QJsonObject_OperatorSubscriptWithKey(QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QJsonValueRef(self->operator[](key_QString));
}

void QJsonObject_Remove(QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->remove(key_QString);
}

QJsonValue* QJsonObject_Take(QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QJsonValue(self->take(key_QString));
}

bool QJsonObject_Contains(const QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->contains(key_QString);
}

iterator QJsonObject_Begin(QJsonObject* self) {
	return self->begin();
}

const_iterator QJsonObject_Begin2(const QJsonObject* self) {
	return self->begin();
}

const_iterator QJsonObject_ConstBegin(const QJsonObject* self) {
	return self->constBegin();
}

iterator QJsonObject_End(QJsonObject* self) {
	return self->end();
}

const_iterator QJsonObject_End2(const QJsonObject* self) {
	return self->end();
}

const_iterator QJsonObject_ConstEnd(const QJsonObject* self) {
	return self->constEnd();
}

iterator QJsonObject_Erase(QJsonObject* self, iterator it) {
	return self->erase(it);
}

iterator QJsonObject_Find(QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->find(key_QString);
}

const_iterator QJsonObject_FindWithKey(const QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->find(key_QString);
}

const_iterator QJsonObject_ConstFind(const QJsonObject* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->constFind(key_QString);
}

iterator QJsonObject_Insert(QJsonObject* self, struct miqt_string key, QJsonValue* value) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->insert(key_QString, *value);
}

bool QJsonObject_Empty(const QJsonObject* self) {
	return self->empty();
}

void QJsonObject_Delete(QJsonObject* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonObject*>( self );
	} else {
		delete self;
	}
}

QJsonObject__iterator* QJsonObject__iterator_new() {
	return new QJsonObject::iterator();
}

QJsonObject__iterator* QJsonObject__iterator_new2(QJsonObject* obj, ptrdiff_t index) {
	return new QJsonObject::iterator(obj, (qsizetype)(index));
}

QJsonObject__iterator* QJsonObject__iterator_new3(const iterator* other) {
	return new QJsonObject::iterator(*other);
}

void QJsonObject__iterator_OperatorAssign(QJsonObject__iterator* self, const iterator* other) {
	self->operator=(*other);
}

struct miqt_string QJsonObject__iterator_Key(const QJsonObject__iterator* self) {
	QString _ret = self->key();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QJsonValueRef* QJsonObject__iterator_Value(const QJsonObject__iterator* self) {
	return new QJsonValueRef(self->value());
}

QJsonValueRef* QJsonObject__iterator_OperatorMultiply(const QJsonObject__iterator* self) {
	return new QJsonValueRef(self->operator*());
}

QJsonValueConstRef* QJsonObject__iterator_OperatorMinusGreater(const QJsonObject__iterator* self) {
	return (QJsonValueConstRef*) self->operator->();
}

QJsonValueRef* QJsonObject__iterator_OperatorMinusGreater2(QJsonObject__iterator* self) {
	return self->operator->();
}

QJsonValueRef* QJsonObject__iterator_OperatorSubscript(const QJsonObject__iterator* self, ptrdiff_t j) {
	return new QJsonValueRef(self->operator[]((qsizetype)(j)));
}

iterator* QJsonObject__iterator_OperatorPlusPlus(QJsonObject__iterator* self) {
	return &self->operator++();
}

iterator QJsonObject__iterator_OperatorPlusPlusWithInt(QJsonObject__iterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

iterator* QJsonObject__iterator_OperatorMinusMinus(QJsonObject__iterator* self) {
	return &self->operator--();
}

iterator QJsonObject__iterator_OperatorMinusMinusWithInt(QJsonObject__iterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

iterator QJsonObject__iterator_OperatorPlus(const QJsonObject__iterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

iterator QJsonObject__iterator_OperatorMinus(const QJsonObject__iterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

iterator* QJsonObject__iterator_OperatorPlusAssign(QJsonObject__iterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

iterator* QJsonObject__iterator_OperatorMinusAssign(QJsonObject__iterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

ptrdiff_t QJsonObject__iterator_OperatorMinusWithIterator(const QJsonObject__iterator* self, iterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QJsonObject__iterator_Delete(QJsonObject__iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonObject::iterator*>( self );
	} else {
		delete self;
	}
}

QJsonObject__const_iterator* QJsonObject__const_iterator_new() {
	return new QJsonObject::const_iterator();
}

QJsonObject__const_iterator* QJsonObject__const_iterator_new2(QJsonObject* obj, ptrdiff_t index) {
	return new QJsonObject::const_iterator(obj, (qsizetype)(index));
}

QJsonObject__const_iterator* QJsonObject__const_iterator_new3(const iterator* other) {
	return new QJsonObject::const_iterator(*other);
}

QJsonObject__const_iterator* QJsonObject__const_iterator_new4(const const_iterator* other) {
	return new QJsonObject::const_iterator(*other);
}

void QJsonObject__const_iterator_OperatorAssign(QJsonObject__const_iterator* self, const const_iterator* other) {
	self->operator=(*other);
}

struct miqt_string QJsonObject__const_iterator_Key(const QJsonObject__const_iterator* self) {
	QString _ret = self->key();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QJsonValueConstRef* QJsonObject__const_iterator_Value(const QJsonObject__const_iterator* self) {
	return new QJsonValueConstRef(self->value());
}

QJsonValueConstRef* QJsonObject__const_iterator_OperatorMultiply(const QJsonObject__const_iterator* self) {
	return new QJsonValueConstRef(self->operator*());
}

QJsonValueConstRef* QJsonObject__const_iterator_OperatorMinusGreater(const QJsonObject__const_iterator* self) {
	return (QJsonValueConstRef*) self->operator->();
}

QJsonValueConstRef* QJsonObject__const_iterator_OperatorSubscript(const QJsonObject__const_iterator* self, ptrdiff_t j) {
	return new QJsonValueConstRef(self->operator[]((qsizetype)(j)));
}

const_iterator* QJsonObject__const_iterator_OperatorPlusPlus(QJsonObject__const_iterator* self) {
	return &self->operator++();
}

const_iterator QJsonObject__const_iterator_OperatorPlusPlusWithInt(QJsonObject__const_iterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

const_iterator* QJsonObject__const_iterator_OperatorMinusMinus(QJsonObject__const_iterator* self) {
	return &self->operator--();
}

const_iterator QJsonObject__const_iterator_OperatorMinusMinusWithInt(QJsonObject__const_iterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

const_iterator QJsonObject__const_iterator_OperatorPlus(const QJsonObject__const_iterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

const_iterator QJsonObject__const_iterator_OperatorMinus(const QJsonObject__const_iterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

const_iterator* QJsonObject__const_iterator_OperatorPlusAssign(QJsonObject__const_iterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

const_iterator* QJsonObject__const_iterator_OperatorMinusAssign(QJsonObject__const_iterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

ptrdiff_t QJsonObject__const_iterator_OperatorMinusWithConstIterator(const QJsonObject__const_iterator* self, const_iterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QJsonObject__const_iterator_Delete(QJsonObject__const_iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonObject::const_iterator*>( self );
	} else {
		delete self;
	}
}

