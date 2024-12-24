// +build ignore

#include <QJsonArray>
#define WORKAROUND_INNER_CLASS_DEFINITION_QJsonArray__const_iterator
#define WORKAROUND_INNER_CLASS_DEFINITION_QJsonArray__iterator
#include <QJsonValue>
#include <QJsonValueConstRef>
#include <QJsonValueRef>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qjsonarray.h>
#include "gen_qjsonarray.h"

#ifndef _Bool
#define _Bool bool
#endif

QJsonArray* QJsonArray_new() {
	return new QJsonArray();
}

QJsonArray* QJsonArray_new2(QJsonArray* other) {
	return new QJsonArray(*other);
}

void QJsonArray_OperatorAssign(QJsonArray* self, QJsonArray* other) {
	self->operator=(*other);
}

QJsonArray* QJsonArray_FromStringList(struct miqt_array /* of struct miqt_string */  list) {
	QStringList list_QList;
	list_QList.reserve(list.len);
	struct miqt_string* list_arr = static_cast<struct miqt_string*>(list.data);
	for(size_t i = 0; i < list.len; ++i) {
		QString list_arr_i_QString = QString::fromUtf8(list_arr[i].data, list_arr[i].len);
		list_QList.push_back(list_arr_i_QString);
	}
	return new QJsonArray(QJsonArray::fromStringList(list_QList));
}

ptrdiff_t QJsonArray_Size(const QJsonArray* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QJsonArray_Count(const QJsonArray* self) {
	qsizetype _ret = self->count();
	return static_cast<ptrdiff_t>(_ret);
}

bool QJsonArray_IsEmpty(const QJsonArray* self) {
	return self->isEmpty();
}

QJsonValue* QJsonArray_At(const QJsonArray* self, ptrdiff_t i) {
	return new QJsonValue(self->at((qsizetype)(i)));
}

QJsonValue* QJsonArray_First(const QJsonArray* self) {
	return new QJsonValue(self->first());
}

QJsonValue* QJsonArray_Last(const QJsonArray* self) {
	return new QJsonValue(self->last());
}

void QJsonArray_Prepend(QJsonArray* self, QJsonValue* value) {
	self->prepend(*value);
}

void QJsonArray_Append(QJsonArray* self, QJsonValue* value) {
	self->append(*value);
}

void QJsonArray_RemoveAt(QJsonArray* self, ptrdiff_t i) {
	self->removeAt((qsizetype)(i));
}

QJsonValue* QJsonArray_TakeAt(QJsonArray* self, ptrdiff_t i) {
	return new QJsonValue(self->takeAt((qsizetype)(i)));
}

void QJsonArray_RemoveFirst(QJsonArray* self) {
	self->removeFirst();
}

void QJsonArray_RemoveLast(QJsonArray* self) {
	self->removeLast();
}

void QJsonArray_Insert(QJsonArray* self, ptrdiff_t i, QJsonValue* value) {
	self->insert((qsizetype)(i), *value);
}

void QJsonArray_Replace(QJsonArray* self, ptrdiff_t i, QJsonValue* value) {
	self->replace((qsizetype)(i), *value);
}

bool QJsonArray_Contains(const QJsonArray* self, QJsonValue* element) {
	return self->contains(*element);
}

QJsonValueRef* QJsonArray_OperatorSubscript(QJsonArray* self, ptrdiff_t i) {
	return new QJsonValueRef(self->operator[]((qsizetype)(i)));
}

QJsonValue* QJsonArray_OperatorSubscriptWithQsizetype(const QJsonArray* self, ptrdiff_t i) {
	return new QJsonValue(self->operator[]((qsizetype)(i)));
}

void QJsonArray_Swap(QJsonArray* self, QJsonArray* other) {
	self->swap(*other);
}

iterator QJsonArray_Begin(QJsonArray* self) {
	return self->begin();
}

const_iterator QJsonArray_Begin2(const QJsonArray* self) {
	return self->begin();
}

const_iterator QJsonArray_ConstBegin(const QJsonArray* self) {
	return self->constBegin();
}

const_iterator QJsonArray_Cbegin(const QJsonArray* self) {
	return self->cbegin();
}

iterator QJsonArray_End(QJsonArray* self) {
	return self->end();
}

const_iterator QJsonArray_End2(const QJsonArray* self) {
	return self->end();
}

const_iterator QJsonArray_ConstEnd(const QJsonArray* self) {
	return self->constEnd();
}

const_iterator QJsonArray_Cend(const QJsonArray* self) {
	return self->cend();
}

iterator QJsonArray_Insert2(QJsonArray* self, iterator before, QJsonValue* value) {
	return self->insert(before, *value);
}

iterator QJsonArray_Erase(QJsonArray* self, iterator it) {
	return self->erase(it);
}

QJsonArray* QJsonArray_OperatorPlus(const QJsonArray* self, QJsonValue* v) {
	return new QJsonArray(self->operator+(*v));
}

QJsonArray* QJsonArray_OperatorPlusAssign(QJsonArray* self, QJsonValue* v) {
	QJsonArray& _ret = self->operator+=(*v);
	// Cast returned reference into pointer
	return &_ret;
}

QJsonArray* QJsonArray_OperatorShiftLeft(QJsonArray* self, QJsonValue* v) {
	QJsonArray& _ret = self->operator<<(*v);
	// Cast returned reference into pointer
	return &_ret;
}

void QJsonArray_PushBack(QJsonArray* self, QJsonValue* t) {
	self->push_back(*t);
}

void QJsonArray_PushFront(QJsonArray* self, QJsonValue* t) {
	self->push_front(*t);
}

void QJsonArray_PopFront(QJsonArray* self) {
	self->pop_front();
}

void QJsonArray_PopBack(QJsonArray* self) {
	self->pop_back();
}

bool QJsonArray_Empty(const QJsonArray* self) {
	return self->empty();
}

void QJsonArray_Delete(QJsonArray* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonArray*>( self );
	} else {
		delete self;
	}
}

QJsonArray__iterator* QJsonArray__iterator_new() {
	return new QJsonArray::iterator();
}

QJsonArray__iterator* QJsonArray__iterator_new2(QJsonArray* array, ptrdiff_t index) {
	return new QJsonArray::iterator(array, (qsizetype)(index));
}

QJsonArray__iterator* QJsonArray__iterator_new3(const iterator* other) {
	return new QJsonArray::iterator(*other);
}

void QJsonArray__iterator_OperatorAssign(QJsonArray__iterator* self, const iterator* other) {
	self->operator=(*other);
}

QJsonValueRef* QJsonArray__iterator_OperatorMultiply(const QJsonArray__iterator* self) {
	return new QJsonValueRef(self->operator*());
}

QJsonValueConstRef* QJsonArray__iterator_OperatorMinusGreater(const QJsonArray__iterator* self) {
	return (QJsonValueConstRef*) self->operator->();
}

QJsonValueRef* QJsonArray__iterator_OperatorMinusGreater2(QJsonArray__iterator* self) {
	return self->operator->();
}

QJsonValueRef* QJsonArray__iterator_OperatorSubscript(const QJsonArray__iterator* self, ptrdiff_t j) {
	return new QJsonValueRef(self->operator[]((qsizetype)(j)));
}

iterator* QJsonArray__iterator_OperatorPlusPlus(QJsonArray__iterator* self) {
	return &self->operator++();
}

iterator QJsonArray__iterator_OperatorPlusPlusWithInt(QJsonArray__iterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

iterator* QJsonArray__iterator_OperatorMinusMinus(QJsonArray__iterator* self) {
	return &self->operator--();
}

iterator QJsonArray__iterator_OperatorMinusMinusWithInt(QJsonArray__iterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

iterator* QJsonArray__iterator_OperatorPlusAssign(QJsonArray__iterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

iterator* QJsonArray__iterator_OperatorMinusAssign(QJsonArray__iterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

iterator QJsonArray__iterator_OperatorPlus(const QJsonArray__iterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

iterator QJsonArray__iterator_OperatorMinus(const QJsonArray__iterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

ptrdiff_t QJsonArray__iterator_OperatorMinusWithIterator(const QJsonArray__iterator* self, iterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QJsonArray__iterator_Delete(QJsonArray__iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonArray::iterator*>( self );
	} else {
		delete self;
	}
}

QJsonArray__const_iterator* QJsonArray__const_iterator_new() {
	return new QJsonArray::const_iterator();
}

QJsonArray__const_iterator* QJsonArray__const_iterator_new2(QJsonArray* array, ptrdiff_t index) {
	return new QJsonArray::const_iterator(array, (qsizetype)(index));
}

QJsonArray__const_iterator* QJsonArray__const_iterator_new3(const iterator* o) {
	return new QJsonArray::const_iterator(*o);
}

QJsonArray__const_iterator* QJsonArray__const_iterator_new4(const const_iterator* other) {
	return new QJsonArray::const_iterator(*other);
}

void QJsonArray__const_iterator_OperatorAssign(QJsonArray__const_iterator* self, const const_iterator* other) {
	self->operator=(*other);
}

QJsonValueConstRef* QJsonArray__const_iterator_OperatorMultiply(const QJsonArray__const_iterator* self) {
	return new QJsonValueConstRef(self->operator*());
}

QJsonValueConstRef* QJsonArray__const_iterator_OperatorMinusGreater(const QJsonArray__const_iterator* self) {
	return (QJsonValueConstRef*) self->operator->();
}

QJsonValueConstRef* QJsonArray__const_iterator_OperatorSubscript(const QJsonArray__const_iterator* self, ptrdiff_t j) {
	return new QJsonValueConstRef(self->operator[]((qsizetype)(j)));
}

const_iterator* QJsonArray__const_iterator_OperatorPlusPlus(QJsonArray__const_iterator* self) {
	return &self->operator++();
}

const_iterator QJsonArray__const_iterator_OperatorPlusPlusWithInt(QJsonArray__const_iterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

const_iterator* QJsonArray__const_iterator_OperatorMinusMinus(QJsonArray__const_iterator* self) {
	return &self->operator--();
}

const_iterator QJsonArray__const_iterator_OperatorMinusMinusWithInt(QJsonArray__const_iterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

const_iterator* QJsonArray__const_iterator_OperatorPlusAssign(QJsonArray__const_iterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

const_iterator* QJsonArray__const_iterator_OperatorMinusAssign(QJsonArray__const_iterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

const_iterator QJsonArray__const_iterator_OperatorPlus(const QJsonArray__const_iterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

const_iterator QJsonArray__const_iterator_OperatorMinus(const QJsonArray__const_iterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

ptrdiff_t QJsonArray__const_iterator_OperatorMinusWithConstIterator(const QJsonArray__const_iterator* self, const_iterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QJsonArray__const_iterator_Delete(QJsonArray__const_iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QJsonArray::const_iterator*>( self );
	} else {
		delete self;
	}
}

