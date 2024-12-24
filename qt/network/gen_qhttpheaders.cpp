// +build ignore

#include <QAnyStringView>
#include <QByteArray>
#include <QByteArrayView>
#include <QHttpHeaders>
#include <QList>
#include <qhttpheaders.h>
#include "gen_qhttpheaders.h"

#ifndef _Bool
#define _Bool bool
#endif

QHttpHeaders* QHttpHeaders_new() {
	return new QHttpHeaders();
}

QHttpHeaders* QHttpHeaders_new2(QHttpHeaders* other) {
	return new QHttpHeaders(*other);
}

void QHttpHeaders_OperatorAssign(QHttpHeaders* self, QHttpHeaders* other) {
	self->operator=(*other);
}

void QHttpHeaders_Swap(QHttpHeaders* self, QHttpHeaders* other) {
	self->swap(*other);
}

bool QHttpHeaders_Append(QHttpHeaders* self, QAnyStringView* name, QAnyStringView* value) {
	return self->append(*name, *value);
}

bool QHttpHeaders_Append2(QHttpHeaders* self, WellKnownHeader name, QAnyStringView* value) {
	return self->append(name, *value);
}

bool QHttpHeaders_Insert(QHttpHeaders* self, ptrdiff_t i, QAnyStringView* name, QAnyStringView* value) {
	return self->insert((qsizetype)(i), *name, *value);
}

bool QHttpHeaders_Insert2(QHttpHeaders* self, ptrdiff_t i, WellKnownHeader name, QAnyStringView* value) {
	return self->insert((qsizetype)(i), name, *value);
}

bool QHttpHeaders_Replace(QHttpHeaders* self, ptrdiff_t i, QAnyStringView* name, QAnyStringView* newValue) {
	return self->replace((qsizetype)(i), *name, *newValue);
}

bool QHttpHeaders_Replace2(QHttpHeaders* self, ptrdiff_t i, WellKnownHeader name, QAnyStringView* newValue) {
	return self->replace((qsizetype)(i), name, *newValue);
}

bool QHttpHeaders_ReplaceOrAppend(QHttpHeaders* self, QAnyStringView* name, QAnyStringView* newValue) {
	return self->replaceOrAppend(*name, *newValue);
}

bool QHttpHeaders_ReplaceOrAppend2(QHttpHeaders* self, WellKnownHeader name, QAnyStringView* newValue) {
	return self->replaceOrAppend(name, *newValue);
}

bool QHttpHeaders_Contains(const QHttpHeaders* self, QAnyStringView* name) {
	return self->contains(*name);
}

bool QHttpHeaders_ContainsWithName(const QHttpHeaders* self, WellKnownHeader name) {
	return self->contains(name);
}

void QHttpHeaders_Clear(QHttpHeaders* self) {
	self->clear();
}

void QHttpHeaders_RemoveAll(QHttpHeaders* self, QAnyStringView* name) {
	self->removeAll(*name);
}

void QHttpHeaders_RemoveAllWithName(QHttpHeaders* self, WellKnownHeader name) {
	self->removeAll(name);
}

void QHttpHeaders_RemoveAt(QHttpHeaders* self, ptrdiff_t i) {
	self->removeAt((qsizetype)(i));
}

QByteArrayView* QHttpHeaders_Value(const QHttpHeaders* self, QAnyStringView* name) {
	return new QByteArrayView(self->value(*name));
}

QByteArrayView* QHttpHeaders_ValueWithName(const QHttpHeaders* self, WellKnownHeader name) {
	return new QByteArrayView(self->value(name));
}

struct miqt_array /* of struct miqt_string */  QHttpHeaders_Values(const QHttpHeaders* self, QAnyStringView* name) {
	QList<QByteArray> _ret = self->values(*name);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QByteArray _lv_qb = _ret[i];
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_qb.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_qb.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_array /* of struct miqt_string */  QHttpHeaders_ValuesWithName(const QHttpHeaders* self, WellKnownHeader name) {
	QList<QByteArray> _ret = self->values(name);
	// Convert QList<> from C++ memory to manually-managed C memory
	struct miqt_string* _arr = static_cast<struct miqt_string*>(malloc(sizeof(struct miqt_string) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		QByteArray _lv_qb = _ret[i];
		struct miqt_string _lv_ms;
		_lv_ms.len = _lv_qb.length();
		_lv_ms.data = static_cast<char*>(malloc(_lv_ms.len));
		memcpy(_lv_ms.data, _lv_qb.data(), _lv_ms.len);
		_arr[i] = _lv_ms;
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QByteArrayView* QHttpHeaders_ValueAt(const QHttpHeaders* self, ptrdiff_t i) {
	return new QByteArrayView(self->valueAt((qsizetype)(i)));
}

struct miqt_string QHttpHeaders_CombinedValue(const QHttpHeaders* self, QAnyStringView* name) {
	QByteArray _qb = self->combinedValue(*name);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

struct miqt_string QHttpHeaders_CombinedValueWithName(const QHttpHeaders* self, WellKnownHeader name) {
	QByteArray _qb = self->combinedValue(name);
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}

ptrdiff_t QHttpHeaders_Size(const QHttpHeaders* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

void QHttpHeaders_Reserve(QHttpHeaders* self, ptrdiff_t size) {
	self->reserve((qsizetype)(size));
}

bool QHttpHeaders_IsEmpty(const QHttpHeaders* self) {
	return self->isEmpty();
}

QByteArrayView* QHttpHeaders_WellKnownHeaderName(WellKnownHeader name) {
	return new QByteArrayView(QHttpHeaders::wellKnownHeaderName(name));
}

QByteArrayView* QHttpHeaders_Value2(const QHttpHeaders* self, QAnyStringView* name, QByteArrayView* defaultValue) {
	return new QByteArrayView(self->value(*name, *defaultValue));
}

QByteArrayView* QHttpHeaders_Value22(const QHttpHeaders* self, WellKnownHeader name, QByteArrayView* defaultValue) {
	return new QByteArrayView(self->value(name, *defaultValue));
}

void QHttpHeaders_Delete(QHttpHeaders* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QHttpHeaders*>( self );
	} else {
		delete self;
	}
}

