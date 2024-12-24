// +build ignore

#include <QCborMap>
#define WORKAROUND_INNER_CLASS_DEFINITION_QCborMap__ConstIterator
#define WORKAROUND_INNER_CLASS_DEFINITION_QCborMap__Iterator
#include <QCborValue>
#include <QCborValueConstRef>
#include <QCborValueRef>
#include <QJsonObject>
#include <QList>
#include <QMap>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QVariant>
#include <qcbormap.h>
#include "gen_qcbormap.h"

#ifndef _Bool
#define _Bool bool
#endif

QCborMap* QCborMap_new() {
	return new QCborMap();
}

QCborMap* QCborMap_new2(QCborMap* other) {
	return new QCborMap(*other);
}

void QCborMap_OperatorAssign(QCborMap* self, QCborMap* other) {
	self->operator=(*other);
}

void QCborMap_Swap(QCborMap* self, QCborMap* other) {
	self->swap(*other);
}

QCborValue* QCborMap_ToCborValue(const QCborMap* self) {
	return new QCborValue(self->toCborValue());
}

ptrdiff_t QCborMap_Size(const QCborMap* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

bool QCborMap_IsEmpty(const QCborMap* self) {
	return self->isEmpty();
}

void QCborMap_Clear(QCborMap* self) {
	self->clear();
}

struct miqt_array /* of QCborValue* */  QCborMap_Keys(const QCborMap* self) {
	QList<QCborValue> _ret = self->keys();
	// Convert QList<> from C++ memory to manually-managed C memory
	QCborValue** _arr = static_cast<QCborValue**>(malloc(sizeof(QCborValue*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QCborValue(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QCborValue* QCborMap_Value(const QCborMap* self, long long key) {
	return new QCborValue(self->value(static_cast<qint64>(key)));
}

QCborValue* QCborMap_Value2(const QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QCborValue(self->value(key_QString));
}

QCborValue* QCborMap_Value3(const QCborMap* self, QCborValue* key) {
	return new QCborValue(self->value(*key));
}

QCborValue* QCborMap_OperatorSubscript(const QCborMap* self, long long key) {
	return new QCborValue(self->operator[](static_cast<qint64>(key)));
}

QCborValue* QCborMap_OperatorSubscript2(const QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QCborValue(self->operator[](key_QString));
}

QCborValue* QCborMap_OperatorSubscript3(const QCborMap* self, QCborValue* key) {
	return new QCborValue(self->operator[](*key));
}

QCborValueRef* QCborMap_OperatorSubscript4(QCborMap* self, long long key) {
	return new QCborValueRef(self->operator[](static_cast<qint64>(key)));
}

QCborValueRef* QCborMap_OperatorSubscript6(QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QCborValueRef(self->operator[](key_QString));
}

QCborValueRef* QCborMap_OperatorSubscript7(QCborMap* self, QCborValue* key) {
	return new QCborValueRef(self->operator[](*key));
}

QCborValue* QCborMap_Take(QCborMap* self, long long key) {
	return new QCborValue(self->take(static_cast<qint64>(key)));
}

QCborValue* QCborMap_Take2(QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return new QCborValue(self->take(key_QString));
}

QCborValue* QCborMap_Take3(QCborMap* self, QCborValue* key) {
	return new QCborValue(self->take(*key));
}

void QCborMap_Remove(QCborMap* self, long long key) {
	self->remove(static_cast<qint64>(key));
}

void QCborMap_Remove2(QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	self->remove(key_QString);
}

void QCborMap_Remove3(QCborMap* self, QCborValue* key) {
	self->remove(*key);
}

bool QCborMap_Contains(const QCborMap* self, long long key) {
	return self->contains(static_cast<qint64>(key));
}

bool QCborMap_Contains2(const QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->contains(key_QString);
}

bool QCborMap_Contains3(const QCborMap* self, QCborValue* key) {
	return self->contains(*key);
}

int QCborMap_Compare(const QCborMap* self, QCborMap* other) {
	return self->compare(*other);
}

iterator QCborMap_Begin(QCborMap* self) {
	return self->begin();
}

const_iterator QCborMap_ConstBegin(const QCborMap* self) {
	return self->constBegin();
}

const_iterator QCborMap_Begin2(const QCborMap* self) {
	return self->begin();
}

const_iterator QCborMap_Cbegin(const QCborMap* self) {
	return self->cbegin();
}

iterator QCborMap_End(QCborMap* self) {
	return self->end();
}

const_iterator QCborMap_ConstEnd(const QCborMap* self) {
	return self->constEnd();
}

const_iterator QCborMap_End2(const QCborMap* self) {
	return self->end();
}

const_iterator QCborMap_Cend(const QCborMap* self) {
	return self->cend();
}

iterator QCborMap_Erase(QCborMap* self, iterator it) {
	return self->erase(it);
}

iterator QCborMap_EraseWithIt(QCborMap* self, const_iterator it) {
	return self->erase(it);
}

QCborValue* QCborMap_Extract(QCborMap* self, iterator it) {
	return new QCborValue(self->extract(it));
}

QCborValue* QCborMap_ExtractWithIt(QCborMap* self, const_iterator it) {
	return new QCborValue(self->extract(it));
}

bool QCborMap_Empty(const QCborMap* self) {
	return self->empty();
}

iterator QCborMap_Find(QCborMap* self, long long key) {
	return self->find(static_cast<qint64>(key));
}

iterator QCborMap_Find2(QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->find(key_QString);
}

iterator QCborMap_Find3(QCborMap* self, QCborValue* key) {
	return self->find(*key);
}

const_iterator QCborMap_ConstFind(const QCborMap* self, long long key) {
	return self->constFind(static_cast<qint64>(key));
}

const_iterator QCborMap_ConstFind2(const QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->constFind(key_QString);
}

const_iterator QCborMap_ConstFind3(const QCborMap* self, QCborValue* key) {
	return self->constFind(*key);
}

const_iterator QCborMap_Find4(const QCborMap* self, long long key) {
	return self->find(static_cast<qint64>(key));
}

const_iterator QCborMap_Find6(const QCborMap* self, struct miqt_string key) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->find(key_QString);
}

const_iterator QCborMap_Find7(const QCborMap* self, QCborValue* key) {
	return self->find(*key);
}

iterator QCborMap_Insert(QCborMap* self, long long key, QCborValue* value_) {
	return self->insert(static_cast<qint64>(key), *value_);
}

iterator QCborMap_Insert3(QCborMap* self, struct miqt_string key, QCborValue* value_) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	return self->insert(key_QString, *value_);
}

iterator QCborMap_Insert4(QCborMap* self, QCborValue* key, QCborValue* value_) {
	return self->insert(*key, *value_);
}

iterator QCborMap_InsertWithValueType(QCborMap* self, value_type v) {
	return self->insert(v);
}

QCborMap* QCborMap_FromVariantMap(struct miqt_map /* of struct miqt_string to QVariant* */  mapVal) {
	QVariantMap mapVal_QMap;
	struct miqt_string* mapVal_karr = static_cast<struct miqt_string*>(mapVal.keys);
	QVariant** mapVal_varr = static_cast<QVariant**>(mapVal.values);
	for(size_t i = 0; i < mapVal.len; ++i) {
		QString mapVal_karr_i_QString = QString::fromUtf8(mapVal_karr[i].data, mapVal_karr[i].len);
		mapVal_QMap[mapVal_karr_i_QString] = *(mapVal_varr[i]);
	}
	return new QCborMap(QCborMap::fromVariantMap(mapVal_QMap));
}

QCborMap* QCborMap_FromVariantHash(struct miqt_map /* of struct miqt_string to QVariant* */  hash) {
	QVariantHash hash_QMap;
	hash_QMap.reserve(hash.len);
	struct miqt_string* hash_karr = static_cast<struct miqt_string*>(hash.keys);
	QVariant** hash_varr = static_cast<QVariant**>(hash.values);
	for(size_t i = 0; i < hash.len; ++i) {
		QString hash_karr_i_QString = QString::fromUtf8(hash_karr[i].data, hash_karr[i].len);
		hash_QMap[hash_karr_i_QString] = *(hash_varr[i]);
	}
	return new QCborMap(QCborMap::fromVariantHash(hash_QMap));
}

QCborMap* QCborMap_FromJsonObject(QJsonObject* o) {
	return new QCborMap(QCborMap::fromJsonObject(*o));
}

struct miqt_map /* of struct miqt_string to QVariant* */  QCborMap_ToVariantMap(const QCborMap* self) {
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

struct miqt_map /* of struct miqt_string to QVariant* */  QCborMap_ToVariantHash(const QCborMap* self) {
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

QJsonObject* QCborMap_ToJsonObject(const QCborMap* self) {
	return new QJsonObject(self->toJsonObject());
}

void QCborMap_Delete(QCborMap* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCborMap*>( self );
	} else {
		delete self;
	}
}

QCborMap__Iterator* QCborMap__Iterator_new() {
	return new QCborMap::Iterator();
}

QCborMap__Iterator* QCborMap__Iterator_new2(const Iterator* param1) {
	return new QCborMap::Iterator(*param1);
}

void QCborMap__Iterator_OperatorAssign(QCborMap__Iterator* self, const Iterator* other) {
	self->operator=(*other);
}

value_type QCborMap__Iterator_OperatorMultiply(const QCborMap__Iterator* self) {
	return self->operator*();
}

value_type QCborMap__Iterator_OperatorSubscript(const QCborMap__Iterator* self, ptrdiff_t j) {
	return self->operator[]((qsizetype)(j));
}

QCborValueRef* QCborMap__Iterator_OperatorMinusGreater(QCborMap__Iterator* self) {
	return self->operator->();
}

QCborValueConstRef* QCborMap__Iterator_OperatorMinusGreater2(const QCborMap__Iterator* self) {
	return (QCborValueConstRef*) self->operator->();
}

QCborValue* QCborMap__Iterator_Key(const QCborMap__Iterator* self) {
	return new QCborValue(self->key());
}

QCborValueRef* QCborMap__Iterator_Value(const QCborMap__Iterator* self) {
	return new QCborValueRef(self->value());
}

Iterator* QCborMap__Iterator_OperatorPlusPlus(QCborMap__Iterator* self) {
	return &self->operator++();
}

Iterator QCborMap__Iterator_OperatorPlusPlusWithInt(QCborMap__Iterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

Iterator* QCborMap__Iterator_OperatorMinusMinus(QCborMap__Iterator* self) {
	return &self->operator--();
}

Iterator QCborMap__Iterator_OperatorMinusMinusWithInt(QCborMap__Iterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

Iterator* QCborMap__Iterator_OperatorPlusAssign(QCborMap__Iterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

Iterator* QCborMap__Iterator_OperatorMinusAssign(QCborMap__Iterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

Iterator QCborMap__Iterator_OperatorPlus(const QCborMap__Iterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

Iterator QCborMap__Iterator_OperatorMinus(const QCborMap__Iterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

ptrdiff_t QCborMap__Iterator_OperatorMinusWithIterator(const QCborMap__Iterator* self, Iterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QCborMap__Iterator_Delete(QCborMap__Iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCborMap::Iterator*>( self );
	} else {
		delete self;
	}
}

QCborMap__ConstIterator* QCborMap__ConstIterator_new() {
	return new QCborMap::ConstIterator();
}

QCborMap__ConstIterator* QCborMap__ConstIterator_new2(const ConstIterator* param1) {
	return new QCborMap::ConstIterator(*param1);
}

void QCborMap__ConstIterator_OperatorAssign(QCborMap__ConstIterator* self, const ConstIterator* other) {
	self->operator=(*other);
}

value_type QCborMap__ConstIterator_OperatorMultiply(const QCborMap__ConstIterator* self) {
	return self->operator*();
}

value_type QCborMap__ConstIterator_OperatorSubscript(const QCborMap__ConstIterator* self, ptrdiff_t j) {
	return self->operator[]((qsizetype)(j));
}

QCborValueConstRef* QCborMap__ConstIterator_OperatorMinusGreater(const QCborMap__ConstIterator* self) {
	return (QCborValueConstRef*) self->operator->();
}

QCborValue* QCborMap__ConstIterator_Key(const QCborMap__ConstIterator* self) {
	return new QCborValue(self->key());
}

QCborValueConstRef* QCborMap__ConstIterator_Value(const QCborMap__ConstIterator* self) {
	return new QCborValueConstRef(self->value());
}

ConstIterator* QCborMap__ConstIterator_OperatorPlusPlus(QCborMap__ConstIterator* self) {
	return &self->operator++();
}

ConstIterator QCborMap__ConstIterator_OperatorPlusPlusWithInt(QCborMap__ConstIterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

ConstIterator* QCborMap__ConstIterator_OperatorMinusMinus(QCborMap__ConstIterator* self) {
	return &self->operator--();
}

ConstIterator QCborMap__ConstIterator_OperatorMinusMinusWithInt(QCborMap__ConstIterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

ConstIterator* QCborMap__ConstIterator_OperatorPlusAssign(QCborMap__ConstIterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

ConstIterator* QCborMap__ConstIterator_OperatorMinusAssign(QCborMap__ConstIterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

ConstIterator QCborMap__ConstIterator_OperatorPlus(const QCborMap__ConstIterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

ConstIterator QCborMap__ConstIterator_OperatorMinus(const QCborMap__ConstIterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

ptrdiff_t QCborMap__ConstIterator_OperatorMinusWithConstIterator(const QCborMap__ConstIterator* self, ConstIterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QCborMap__ConstIterator_Delete(QCborMap__ConstIterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCborMap::ConstIterator*>( self );
	} else {
		delete self;
	}
}

