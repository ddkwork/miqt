#include <QFrame>
#include <QMetaObject>
#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include "qframe.h"

#include "gen_qframe.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QFrame* QFrame_new() {
	return new QFrame();
}

QFrame* QFrame_new2(QWidget* parent) {
	return new QFrame(parent);
}

QFrame* QFrame_new3(QWidget* parent, int f) {
	return new QFrame(parent, static_cast<Qt::WindowFlags>(f));
}

QMetaObject* QFrame_MetaObject(QFrame* self) {
	return (QMetaObject*) const_cast<const QFrame*>(self)->metaObject();
}

void QFrame_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QFrame::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFrame_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QFrame::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

int QFrame_FrameStyle(QFrame* self) {
	return const_cast<const QFrame*>(self)->frameStyle();
}

void QFrame_SetFrameStyle(QFrame* self, int frameStyle) {
	self->setFrameStyle(static_cast<int>(frameStyle));
}

int QFrame_FrameWidth(QFrame* self) {
	return const_cast<const QFrame*>(self)->frameWidth();
}

QSize* QFrame_SizeHint(QFrame* self) {
	QSize ret = const_cast<const QFrame*>(self)->sizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

uintptr_t QFrame_FrameShape(QFrame* self) {
	QFrame::Shape ret = const_cast<const QFrame*>(self)->frameShape();
	return static_cast<uintptr_t>(ret);
}

void QFrame_SetFrameShape(QFrame* self, uintptr_t frameShape) {
	self->setFrameShape(static_cast<QFrame::Shape>(frameShape));
}

uintptr_t QFrame_FrameShadow(QFrame* self) {
	QFrame::Shadow ret = const_cast<const QFrame*>(self)->frameShadow();
	return static_cast<uintptr_t>(ret);
}

void QFrame_SetFrameShadow(QFrame* self, uintptr_t frameShadow) {
	self->setFrameShadow(static_cast<QFrame::Shadow>(frameShadow));
}

int QFrame_LineWidth(QFrame* self) {
	return const_cast<const QFrame*>(self)->lineWidth();
}

void QFrame_SetLineWidth(QFrame* self, int lineWidth) {
	self->setLineWidth(static_cast<int>(lineWidth));
}

int QFrame_MidLineWidth(QFrame* self) {
	return const_cast<const QFrame*>(self)->midLineWidth();
}

void QFrame_SetMidLineWidth(QFrame* self, int midLineWidth) {
	self->setMidLineWidth(static_cast<int>(midLineWidth));
}

QRect* QFrame_FrameRect(QFrame* self) {
	QRect ret = const_cast<const QFrame*>(self)->frameRect();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QRect*>(new QRect(ret));
}

void QFrame_SetFrameRect(QFrame* self, QRect* frameRect) {
	self->setFrameRect(*frameRect);
}

void QFrame_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QFrame::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFrame_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QFrame::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFrame_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QFrame::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFrame_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QFrame::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QFrame_Delete(QFrame* self) {
	delete self;
}
