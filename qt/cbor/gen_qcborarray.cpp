// +build ignore

#include <QCborArray>
#define WORKAROUND_INNER_CLASS_DEFINITION_QCborArray__ConstIterator
#define WORKAROUND_INNER_CLASS_DEFINITION_QCborArray__Iterator
#include <QCborValue>
#include <QCborValueConstRef>
#include <QCborValueRef>
#include <QJsonArray>
#include <QList>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qcborarray.h>
#include "gen_qcborarray.h"

QCborArray* QCborArray_new() {
	return new QCborArray();
}

QCborArray* QCborArray_new2(QCborArray* other) {
	return new QCborArray(*other);
}

void QCborArray_OperatorAssign(QCborArray* self, QCborArray* other) {
	self->operator=(*other);
}

void QCborArray_Swap(QCborArray* self, QCborArray* other) {
	self->swap(*other);
}

QCborValue* QCborArray_ToCborValue(const QCborArray* self) {
	return new QCborValue(self->toCborValue());
}

ptrdiff_t QCborArray_Size(const QCborArray* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

bool QCborArray_IsEmpty(const QCborArray* self) {
	return self->isEmpty();
}

void QCborArray_Clear(QCborArray* self) {
	self->clear();
}

QCborValue* QCborArray_At(const QCborArray* self, ptrdiff_t i) {
	return new QCborValue(self->at((qsizetype)(i)));
}

QCborValue* QCborArray_First(const QCborArray* self) {
	return new QCborValue(self->first());
}

QCborValue* QCborArray_Last(const QCborArray* self) {
	return new QCborValue(self->last());
}

QCborValue* QCborArray_OperatorSubscript(const QCborArray* self, ptrdiff_t i) {
	return new QCborValue(self->operator[]((qsizetype)(i)));
}

QCborValueRef* QCborArray_First2(QCborArray* self) {
	return new QCborValueRef(self->first());
}

QCborValueRef* QCborArray_Last2(QCborArray* self) {
	return new QCborValueRef(self->last());
}

QCborValueRef* QCborArray_OperatorSubscriptWithQsizetype(QCborArray* self, ptrdiff_t i) {
	return new QCborValueRef(self->operator[]((qsizetype)(i)));
}

void QCborArray_Insert(QCborArray* self, ptrdiff_t i, QCborValue* value) {
	self->insert((qsizetype)(i), *value);
}

void QCborArray_Prepend(QCborArray* self, QCborValue* value) {
	self->prepend(*value);
}

void QCborArray_Append(QCborArray* self, QCborValue* value) {
	self->append(*value);
}

QCborValue* QCborArray_Extract(QCborArray* self, ConstIterator it) {
	return new QCborValue(self->extract(it));
}

QCborValue* QCborArray_ExtractWithIt(QCborArray* self, Iterator it) {
	return new QCborValue(self->extract(it));
}

void QCborArray_RemoveAt(QCborArray* self, ptrdiff_t i) {
	self->removeAt((qsizetype)(i));
}

QCborValue* QCborArray_TakeAt(QCborArray* self, ptrdiff_t i) {
	return new QCborValue(self->takeAt((qsizetype)(i)));
}

void QCborArray_RemoveFirst(QCborArray* self) {
	self->removeFirst();
}

void QCborArray_RemoveLast(QCborArray* self) {
	self->removeLast();
}

QCborValue* QCborArray_TakeFirst(QCborArray* self) {
	return new QCborValue(self->takeFirst());
}

QCborValue* QCborArray_TakeLast(QCborArray* self) {
	return new QCborValue(self->takeLast());
}

bool QCborArray_Contains(const QCborArray* self, QCborValue* value) {
	return self->contains(*value);
}

int QCborArray_Compare(const QCborArray* self, QCborArray* other) {
	return self->compare(*other);
}

iterator QCborArray_Begin(QCborArray* self) {
	return self->begin();
}

const_iterator QCborArray_ConstBegin(const QCborArray* self) {
	return self->constBegin();
}

const_iterator QCborArray_Begin2(const QCborArray* self) {
	return self->begin();
}

const_iterator QCborArray_Cbegin(const QCborArray* self) {
	return self->cbegin();
}

iterator QCborArray_End(QCborArray* self) {
	return self->end();
}

const_iterator QCborArray_ConstEnd(const QCborArray* self) {
	return self->constEnd();
}

const_iterator QCborArray_End2(const QCborArray* self) {
	return self->end();
}

const_iterator QCborArray_Cend(const QCborArray* self) {
	return self->cend();
}

iterator QCborArray_Insert2(QCborArray* self, iterator before, QCborValue* value) {
	return self->insert(before, *value);
}

iterator QCborArray_Insert3(QCborArray* self, const_iterator before, QCborValue* value) {
	return self->insert(before, *value);
}

iterator QCborArray_Erase(QCborArray* self, iterator it) {
	return self->erase(it);
}

iterator QCborArray_EraseWithIt(QCborArray* self, const_iterator it) {
	return self->erase(it);
}

void QCborArray_PushBack(QCborArray* self, QCborValue* t) {
	self->push_back(*t);
}

void QCborArray_PushFront(QCborArray* self, QCborValue* t) {
	self->push_front(*t);
}

void QCborArray_PopFront(QCborArray* self) {
	self->pop_front();
}

void QCborArray_PopBack(QCborArray* self) {
	self->pop_back();
}

bool QCborArray_Empty(const QCborArray* self) {
	return self->empty();
}

QCborArray* QCborArray_OperatorPlus(const QCborArray* self, QCborValue* v) {
	return new QCborArray(self->operator+(*v));
}

QCborArray* QCborArray_OperatorPlusAssign(QCborArray* self, QCborValue* v) {
	QCborArray& _ret = self->operator+=(*v);
	// Cast returned reference into pointer
	return &_ret;
}

QCborArray* QCborArray_OperatorShiftLeft(QCborArray* self, QCborValue* v) {
	QCborArray& _ret = self->operator<<(*v);
	// Cast returned reference into pointer
	return &_ret;
}

QCborArray* QCborArray_FromStringList(struct miqt_array /* of struct miqt_string */  list) {
	QStringList list_QList;
	list_QList.reserve(list.len);
	struct miqt_string* list_arr = static_cast<struct miqt_string*>(list.data);
	for(size_t i = 0; i < list.len; ++i) {
		QString list_arr_i_QString = QString::fromUtf8(list_arr[i].data, list_arr[i].len);
		list_QList.push_back(list_arr_i_QString);
	}
	return new QCborArray(QCborArray::fromStringList(list_QList));
}

QCborArray* QCborArray_FromJsonArray(QJsonArray* array) {
	return new QCborArray(QCborArray::fromJsonArray(*array));
}

QJsonArray* QCborArray_ToJsonArray(const QCborArray* self) {
	return new QJsonArray(self->toJsonArray());
}

void QCborArray_Delete(QCborArray* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCborArray*>( self );
	} else {
		delete self;
	}
}

QCborArray__Iterator* QCborArray__Iterator_new() {
	return new QCborArray::Iterator();
}

QCborArray__Iterator* QCborArray__Iterator_new2(const Iterator* param1) {
	return new QCborArray::Iterator(*param1);
}

void QCborArray__Iterator_OperatorAssign(QCborArray__Iterator* self, const Iterator* other) {
	self->operator=(*other);
}

QCborValueRef* QCborArray__Iterator_OperatorMultiply(const QCborArray__Iterator* self) {
	return new QCborValueRef(self->operator*());
}

QCborValueRef* QCborArray__Iterator_OperatorMinusGreater(QCborArray__Iterator* self) {
	return self->operator->();
}

QCborValueConstRef* QCborArray__Iterator_OperatorMinusGreater2(const QCborArray__Iterator* self) {
	return (QCborValueConstRef*) self->operator->();
}

QCborValueRef* QCborArray__Iterator_OperatorSubscript(const QCborArray__Iterator* self, ptrdiff_t j) {
	return new QCborValueRef(self->operator[]((qsizetype)(j)));
}

Iterator* QCborArray__Iterator_OperatorPlusPlus(QCborArray__Iterator* self) {
	return &self->operator++();
}

Iterator QCborArray__Iterator_OperatorPlusPlusWithInt(QCborArray__Iterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

Iterator* QCborArray__Iterator_OperatorMinusMinus(QCborArray__Iterator* self) {
	return &self->operator--();
}

Iterator QCborArray__Iterator_OperatorMinusMinusWithInt(QCborArray__Iterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

Iterator* QCborArray__Iterator_OperatorPlusAssign(QCborArray__Iterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

Iterator* QCborArray__Iterator_OperatorMinusAssign(QCborArray__Iterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

Iterator QCborArray__Iterator_OperatorPlus(const QCborArray__Iterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

Iterator QCborArray__Iterator_OperatorMinus(const QCborArray__Iterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

ptrdiff_t QCborArray__Iterator_OperatorMinusWithIterator(const QCborArray__Iterator* self, Iterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QCborArray__Iterator_Delete(QCborArray__Iterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCborArray::Iterator*>( self );
	} else {
		delete self;
	}
}

QCborArray__ConstIterator* QCborArray__ConstIterator_new() {
	return new QCborArray::ConstIterator();
}

QCborArray__ConstIterator* QCborArray__ConstIterator_new2(const ConstIterator* param1) {
	return new QCborArray::ConstIterator(*param1);
}

void QCborArray__ConstIterator_OperatorAssign(QCborArray__ConstIterator* self, const ConstIterator* other) {
	self->operator=(*other);
}

QCborValueConstRef* QCborArray__ConstIterator_OperatorMultiply(const QCborArray__ConstIterator* self) {
	return new QCborValueConstRef(self->operator*());
}

QCborValueConstRef* QCborArray__ConstIterator_OperatorMinusGreater(const QCborArray__ConstIterator* self) {
	return (QCborValueConstRef*) self->operator->();
}

QCborValueConstRef* QCborArray__ConstIterator_OperatorSubscript(const QCborArray__ConstIterator* self, ptrdiff_t j) {
	return new QCborValueConstRef(self->operator[]((qsizetype)(j)));
}

ConstIterator* QCborArray__ConstIterator_OperatorPlusPlus(QCborArray__ConstIterator* self) {
	return &self->operator++();
}

ConstIterator QCborArray__ConstIterator_OperatorPlusPlusWithInt(QCborArray__ConstIterator* self, int param1) {
	return self->operator++(static_cast<int>(param1));
}

ConstIterator* QCborArray__ConstIterator_OperatorMinusMinus(QCborArray__ConstIterator* self) {
	return &self->operator--();
}

ConstIterator QCborArray__ConstIterator_OperatorMinusMinusWithInt(QCborArray__ConstIterator* self, int param1) {
	return self->operator--(static_cast<int>(param1));
}

ConstIterator* QCborArray__ConstIterator_OperatorPlusAssign(QCborArray__ConstIterator* self, ptrdiff_t j) {
	return &self->operator+=((qsizetype)(j));
}

ConstIterator* QCborArray__ConstIterator_OperatorMinusAssign(QCborArray__ConstIterator* self, ptrdiff_t j) {
	return &self->operator-=((qsizetype)(j));
}

ConstIterator QCborArray__ConstIterator_OperatorPlus(const QCborArray__ConstIterator* self, ptrdiff_t j) {
	return self->operator+((qsizetype)(j));
}

ConstIterator QCborArray__ConstIterator_OperatorMinus(const QCborArray__ConstIterator* self, ptrdiff_t j) {
	return self->operator-((qsizetype)(j));
}

ptrdiff_t QCborArray__ConstIterator_OperatorMinusWithConstIterator(const QCborArray__ConstIterator* self, ConstIterator j) {
	qsizetype _ret = self->operator-(j);
	return static_cast<ptrdiff_t>(_ret);
}

void QCborArray__ConstIterator_Delete(QCborArray__ConstIterator* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QCborArray::ConstIterator*>( self );
	} else {
		delete self;
	}
}

