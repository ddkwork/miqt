#include <QKeyCombination>
#include <QKeySequence>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qkeysequence.h>
#include "gen_qkeysequence.h"
#include "_cgo_export.h"

void QKeySequence_new(QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence();
	*outptr_QKeySequence = ret;
}

void QKeySequence_new2(struct miqt_string key, QKeySequence** outptr_QKeySequence) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	QKeySequence* ret = new QKeySequence(key_QString);
	*outptr_QKeySequence = ret;
}

void QKeySequence_new3(int k1, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(static_cast<int>(k1));
	*outptr_QKeySequence = ret;
}

void QKeySequence_new4(QKeyCombination* k1, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(*k1);
	*outptr_QKeySequence = ret;
}

void QKeySequence_new5(QKeySequence* ks, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(*ks);
	*outptr_QKeySequence = ret;
}

void QKeySequence_new6(int key, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
	*outptr_QKeySequence = ret;
}

void QKeySequence_new7(struct miqt_string key, int format, QKeySequence** outptr_QKeySequence) {
	QString key_QString = QString::fromUtf8(key.data, key.len);
	QKeySequence* ret = new QKeySequence(key_QString, static_cast<QKeySequence::SequenceFormat>(format));
	*outptr_QKeySequence = ret;
}

void QKeySequence_new8(int k1, int k2, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(static_cast<int>(k1), static_cast<int>(k2));
	*outptr_QKeySequence = ret;
}

void QKeySequence_new9(int k1, int k2, int k3, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(static_cast<int>(k1), static_cast<int>(k2), static_cast<int>(k3));
	*outptr_QKeySequence = ret;
}

void QKeySequence_new10(int k1, int k2, int k3, int k4, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(static_cast<int>(k1), static_cast<int>(k2), static_cast<int>(k3), static_cast<int>(k4));
	*outptr_QKeySequence = ret;
}

void QKeySequence_new11(QKeyCombination* k1, QKeyCombination* k2, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(*k1, *k2);
	*outptr_QKeySequence = ret;
}

void QKeySequence_new12(QKeyCombination* k1, QKeyCombination* k2, QKeyCombination* k3, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(*k1, *k2, *k3);
	*outptr_QKeySequence = ret;
}

void QKeySequence_new13(QKeyCombination* k1, QKeyCombination* k2, QKeyCombination* k3, QKeyCombination* k4, QKeySequence** outptr_QKeySequence) {
	QKeySequence* ret = new QKeySequence(*k1, *k2, *k3, *k4);
	*outptr_QKeySequence = ret;
}

int QKeySequence_Count(const QKeySequence* self) {
	return self->count();
}

bool QKeySequence_IsEmpty(const QKeySequence* self) {
	return self->isEmpty();
}

struct miqt_string QKeySequence_ToString(const QKeySequence* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QKeySequence* QKeySequence_FromString(struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	return new QKeySequence(QKeySequence::fromString(str_QString));
}

struct miqt_array /* of QKeySequence* */  QKeySequence_ListFromString(struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QList<QKeySequence> _ret = QKeySequence::listFromString(str_QString);
	// Convert QList<> from C++ memory to manually-managed C memory
	QKeySequence** _arr = static_cast<QKeySequence**>(malloc(sizeof(QKeySequence*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QKeySequence(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QKeySequence_ListToString(struct miqt_array /* of QKeySequence* */  list) {
	QList<QKeySequence> list_QList;
	list_QList.reserve(list.len);
	QKeySequence** list_arr = static_cast<QKeySequence**>(list.data);
	for(size_t i = 0; i < list.len; ++i) {
		list_QList.push_back(*(list_arr[i]));
	}
	QString _ret = QKeySequence::listToString(list_QList);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QKeySequence_Matches(const QKeySequence* self, QKeySequence* seq) {
	QKeySequence::SequenceMatch _ret = self->matches(*seq);
	return static_cast<int>(_ret);
}

QKeySequence* QKeySequence_Mnemonic(struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return new QKeySequence(QKeySequence::mnemonic(text_QString));
}

struct miqt_array /* of QKeySequence* */  QKeySequence_KeyBindings(int key) {
	QList<QKeySequence> _ret = QKeySequence::keyBindings(static_cast<QKeySequence::StandardKey>(key));
	// Convert QList<> from C++ memory to manually-managed C memory
	QKeySequence** _arr = static_cast<QKeySequence**>(malloc(sizeof(QKeySequence*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QKeySequence(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

QKeyCombination* QKeySequence_OperatorSubscript(const QKeySequence* self, unsigned int i) {
	return new QKeyCombination(self->operator[](static_cast<uint>(i)));
}

void QKeySequence_OperatorAssign(QKeySequence* self, QKeySequence* other) {
	self->operator=(*other);
}

void QKeySequence_Swap(QKeySequence* self, QKeySequence* other) {
	self->swap(*other);
}

bool QKeySequence_OperatorEqual(const QKeySequence* self, QKeySequence* other) {
	return self->operator==(*other);
}

bool QKeySequence_OperatorNotEqual(const QKeySequence* self, QKeySequence* other) {
	return self->operator!=(*other);
}

bool QKeySequence_OperatorLesser(const QKeySequence* self, QKeySequence* ks) {
	return self->operator<(*ks);
}

bool QKeySequence_OperatorGreater(const QKeySequence* self, QKeySequence* other) {
	return self->operator>(*other);
}

bool QKeySequence_OperatorLesserOrEqual(const QKeySequence* self, QKeySequence* other) {
	return self->operator<=(*other);
}

bool QKeySequence_OperatorGreaterOrEqual(const QKeySequence* self, QKeySequence* other) {
	return self->operator>=(*other);
}

bool QKeySequence_IsDetached(const QKeySequence* self) {
	return self->isDetached();
}

struct miqt_string QKeySequence_ToString1(const QKeySequence* self, int format) {
	QString _ret = self->toString(static_cast<QKeySequence::SequenceFormat>(format));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QKeySequence* QKeySequence_FromString2(struct miqt_string str, int format) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	return new QKeySequence(QKeySequence::fromString(str_QString, static_cast<QKeySequence::SequenceFormat>(format)));
}

struct miqt_array /* of QKeySequence* */  QKeySequence_ListFromString2(struct miqt_string str, int format) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	QList<QKeySequence> _ret = QKeySequence::listFromString(str_QString, static_cast<QKeySequence::SequenceFormat>(format));
	// Convert QList<> from C++ memory to manually-managed C memory
	QKeySequence** _arr = static_cast<QKeySequence**>(malloc(sizeof(QKeySequence*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = new QKeySequence(_ret[i]);
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

struct miqt_string QKeySequence_ListToString2(struct miqt_array /* of QKeySequence* */  list, int format) {
	QList<QKeySequence> list_QList;
	list_QList.reserve(list.len);
	QKeySequence** list_arr = static_cast<QKeySequence**>(list.data);
	for(size_t i = 0; i < list.len; ++i) {
		list_QList.push_back(*(list_arr[i]));
	}
	QString _ret = QKeySequence::listToString(list_QList, static_cast<QKeySequence::SequenceFormat>(format));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QKeySequence_Delete(QKeySequence* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QKeySequence*>( self );
	} else {
		delete self;
	}
}

