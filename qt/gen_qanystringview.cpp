// +build ignore

#include <QAnyStringView>
#include <QByteArray>
#include <QChar>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <qanystringview.h>
#include "gen_qanystringview.h"

QAnyStringView* QAnyStringView_new() {
	return new QAnyStringView();
}

QAnyStringView* QAnyStringView_new2(struct miqt_string str) {
	QByteArray str_QByteArray(str.data, str.len);
	return new QAnyStringView(str_QByteArray);
}

QAnyStringView* QAnyStringView_new3(struct miqt_string str) {
	QString str_QString = QString::fromUtf8(str.data, str.len);
	return new QAnyStringView(str_QString);
}

QAnyStringView* QAnyStringView_new4(QAnyStringView* param1) {
	return new QAnyStringView(*param1);
}

QAnyStringView* QAnyStringView_Mid(const QAnyStringView* self, ptrdiff_t pos) {
	return new QAnyStringView(self->mid((qsizetype)(pos)));
}

QAnyStringView* QAnyStringView_Left(const QAnyStringView* self, ptrdiff_t n) {
	return new QAnyStringView(self->left((qsizetype)(n)));
}

QAnyStringView* QAnyStringView_Right(const QAnyStringView* self, ptrdiff_t n) {
	return new QAnyStringView(self->right((qsizetype)(n)));
}

QAnyStringView* QAnyStringView_Sliced(const QAnyStringView* self, ptrdiff_t pos) {
	return new QAnyStringView(self->sliced((qsizetype)(pos)));
}

QAnyStringView* QAnyStringView_Sliced2(const QAnyStringView* self, ptrdiff_t pos, ptrdiff_t n) {
	return new QAnyStringView(self->sliced((qsizetype)(pos), (qsizetype)(n)));
}

QAnyStringView* QAnyStringView_First(const QAnyStringView* self, ptrdiff_t n) {
	return new QAnyStringView(self->first((qsizetype)(n)));
}

QAnyStringView* QAnyStringView_Last(const QAnyStringView* self, ptrdiff_t n) {
	return new QAnyStringView(self->last((qsizetype)(n)));
}

QAnyStringView* QAnyStringView_Chopped(const QAnyStringView* self, ptrdiff_t n) {
	return new QAnyStringView(self->chopped((qsizetype)(n)));
}

QAnyStringView* QAnyStringView_Slice(QAnyStringView* self, ptrdiff_t pos) {
	QAnyStringView& _ret = self->slice((qsizetype)(pos));
	// Cast returned reference into pointer
	return &_ret;
}

QAnyStringView* QAnyStringView_Slice2(QAnyStringView* self, ptrdiff_t pos, ptrdiff_t n) {
	QAnyStringView& _ret = self->slice((qsizetype)(pos), (qsizetype)(n));
	// Cast returned reference into pointer
	return &_ret;
}

void QAnyStringView_Truncate(QAnyStringView* self, ptrdiff_t n) {
	self->truncate((qsizetype)(n));
}

void QAnyStringView_Chop(QAnyStringView* self, ptrdiff_t n) {
	self->chop((qsizetype)(n));
}

struct miqt_string QAnyStringView_ToString(const QAnyStringView* self) {
	QString _ret = self->toString();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

ptrdiff_t QAnyStringView_Size(const QAnyStringView* self) {
	qsizetype _ret = self->size();
	return static_cast<ptrdiff_t>(_ret);
}

const void* QAnyStringView_Data(const QAnyStringView* self) {
	return (const void*) self->data();
}

int QAnyStringView_Compare(QAnyStringView* lhs, QAnyStringView* rhs) {
	return QAnyStringView::compare(*lhs, *rhs);
}

bool QAnyStringView_Equal(QAnyStringView* lhs, QAnyStringView* rhs) {
	return QAnyStringView::equal(*lhs, *rhs);
}

QChar* QAnyStringView_Front(const QAnyStringView* self) {
	return new QChar(self->front());
}

QChar* QAnyStringView_Back(const QAnyStringView* self) {
	return new QChar(self->back());
}

bool QAnyStringView_Empty(const QAnyStringView* self) {
	return self->empty();
}

ptrdiff_t QAnyStringView_SizeBytes(const QAnyStringView* self) {
	qsizetype _ret = self->size_bytes();
	return static_cast<ptrdiff_t>(_ret);
}

ptrdiff_t QAnyStringView_MaxSize(const QAnyStringView* self) {
	qsizetype _ret = self->max_size();
	return static_cast<ptrdiff_t>(_ret);
}

bool QAnyStringView_IsNull(const QAnyStringView* self) {
	return self->isNull();
}

bool QAnyStringView_IsEmpty(const QAnyStringView* self) {
	return self->isEmpty();
}

ptrdiff_t QAnyStringView_Length(const QAnyStringView* self) {
	qsizetype _ret = self->length();
	return static_cast<ptrdiff_t>(_ret);
}

QAnyStringView* QAnyStringView_Mid2(const QAnyStringView* self, ptrdiff_t pos, ptrdiff_t n) {
	return new QAnyStringView(self->mid((qsizetype)(pos), (qsizetype)(n)));
}

int QAnyStringView_Compare3(QAnyStringView* lhs, QAnyStringView* rhs, int cs) {
	return QAnyStringView::compare(*lhs, *rhs, static_cast<Qt::CaseSensitivity>(cs));
}

void QAnyStringView_Delete(QAnyStringView* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QAnyStringView*>( self );
	} else {
		delete self;
	}
}

