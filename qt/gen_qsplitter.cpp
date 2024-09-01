#include <QByteArray>
#include <QList>
#include <QMetaObject>
#include <QSize>
#include <QSplitter>
#include <QSplitterHandle>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include "qsplitter.h"

#include "gen_qsplitter.h"

extern "C" {
    extern void miqt_exec_callback(void* cb, int argc, void* argv);
}

QSplitter* QSplitter_new() {
	return new QSplitter();
}

QSplitter* QSplitter_new2(uintptr_t param1) {
	return new QSplitter(static_cast<Qt::Orientation>(param1));
}

QSplitter* QSplitter_new3(QWidget* parent) {
	return new QSplitter(parent);
}

QSplitter* QSplitter_new4(uintptr_t param1, QWidget* parent) {
	return new QSplitter(static_cast<Qt::Orientation>(param1), parent);
}

QMetaObject* QSplitter_MetaObject(QSplitter* self) {
	return (QMetaObject*) const_cast<const QSplitter*>(self)->metaObject();
}

void QSplitter_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QSplitter::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitter_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QSplitter::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitter_AddWidget(QSplitter* self, QWidget* widget) {
	self->addWidget(widget);
}

void QSplitter_InsertWidget(QSplitter* self, int index, QWidget* widget) {
	self->insertWidget(static_cast<int>(index), widget);
}

QWidget* QSplitter_ReplaceWidget(QSplitter* self, int index, QWidget* widget) {
	return self->replaceWidget(static_cast<int>(index), widget);
}

void QSplitter_SetOrientation(QSplitter* self, uintptr_t orientation) {
	self->setOrientation(static_cast<Qt::Orientation>(orientation));
}

uintptr_t QSplitter_Orientation(QSplitter* self) {
	Qt::Orientation ret = const_cast<const QSplitter*>(self)->orientation();
	return static_cast<uintptr_t>(ret);
}

void QSplitter_SetChildrenCollapsible(QSplitter* self, bool childrenCollapsible) {
	self->setChildrenCollapsible(childrenCollapsible);
}

bool QSplitter_ChildrenCollapsible(QSplitter* self) {
	return const_cast<const QSplitter*>(self)->childrenCollapsible();
}

void QSplitter_SetCollapsible(QSplitter* self, int index, bool param2) {
	self->setCollapsible(static_cast<int>(index), param2);
}

bool QSplitter_IsCollapsible(QSplitter* self, int index) {
	return const_cast<const QSplitter*>(self)->isCollapsible(static_cast<int>(index));
}

void QSplitter_SetOpaqueResize(QSplitter* self) {
	self->setOpaqueResize();
}

bool QSplitter_OpaqueResize(QSplitter* self) {
	return const_cast<const QSplitter*>(self)->opaqueResize();
}

void QSplitter_Refresh(QSplitter* self) {
	self->refresh();
}

QSize* QSplitter_SizeHint(QSplitter* self) {
	QSize ret = const_cast<const QSplitter*>(self)->sizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

QSize* QSplitter_MinimumSizeHint(QSplitter* self) {
	QSize ret = const_cast<const QSplitter*>(self)->minimumSizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

void QSplitter_Sizes(QSplitter* self, int** _out, size_t* _out_len) {
	QList<int> ret = const_cast<const QSplitter*>(self)->sizes();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* __out = static_cast<int*>(malloc(sizeof(int) * ret.length()));
	for (size_t i = 0, e = ret.length(); i < e; ++i) {
		__out[i] = ret[i];
	}
	*_out = __out;
	*_out_len = ret.length();
}

void QSplitter_SetSizes(QSplitter* self, int* list, size_t list_len) {
	QList<int> list_QList;
	list_QList.reserve(list_len);
	for(size_t i = 0; i < list_len; ++i) {
		list_QList.push_back(list[i]);
	}
	self->setSizes(list_QList);
}

QByteArray* QSplitter_SaveState(QSplitter* self) {
	QByteArray ret = const_cast<const QSplitter*>(self)->saveState();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QByteArray*>(new QByteArray(ret));
}

bool QSplitter_RestoreState(QSplitter* self, QByteArray* state) {
	return self->restoreState(*state);
}

int QSplitter_HandleWidth(QSplitter* self) {
	return const_cast<const QSplitter*>(self)->handleWidth();
}

void QSplitter_SetHandleWidth(QSplitter* self, int handleWidth) {
	self->setHandleWidth(static_cast<int>(handleWidth));
}

int QSplitter_IndexOf(QSplitter* self, QWidget* w) {
	return const_cast<const QSplitter*>(self)->indexOf(w);
}

QWidget* QSplitter_Widget(QSplitter* self, int index) {
	return const_cast<const QSplitter*>(self)->widget(static_cast<int>(index));
}

int QSplitter_Count(QSplitter* self) {
	return const_cast<const QSplitter*>(self)->count();
}

void QSplitter_GetRange(QSplitter* self, int index, int* param2, int* param3) {
	const_cast<const QSplitter*>(self)->getRange(static_cast<int>(index), static_cast<int*>(param2), static_cast<int*>(param3));
}

QSplitterHandle* QSplitter_Handle(QSplitter* self, int index) {
	return const_cast<const QSplitter*>(self)->handle(static_cast<int>(index));
}

void QSplitter_SetStretchFactor(QSplitter* self, int index, int stretch) {
	self->setStretchFactor(static_cast<int>(index), static_cast<int>(stretch));
}

void QSplitter_SplitterMoved(QSplitter* self, int pos, int index) {
	self->splitterMoved(static_cast<int>(pos), static_cast<int>(index));
}

void QSplitter_connect_SplitterMoved(QSplitter* self, void* slot) {
	QSplitter::connect(self, static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), self, [=](int pos, int index) {
		miqt_exec_callback(slot, 0, nullptr);
	});
}

void QSplitter_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QSplitter::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitter_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QSplitter::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitter_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QSplitter::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitter_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QSplitter::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitter_SetOpaqueResize1(QSplitter* self, bool opaque) {
	self->setOpaqueResize(opaque);
}

void QSplitter_Delete(QSplitter* self) {
	delete self;
}

QSplitterHandle* QSplitterHandle_new(uintptr_t o, QSplitter* parent) {
	return new QSplitterHandle(static_cast<Qt::Orientation>(o), parent);
}

QMetaObject* QSplitterHandle_MetaObject(QSplitterHandle* self) {
	return (QMetaObject*) const_cast<const QSplitterHandle*>(self)->metaObject();
}

void QSplitterHandle_Tr(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QSplitterHandle::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitterHandle_TrUtf8(const char* s, char** _out, int* _out_Strlen) {
	QString ret = QSplitterHandle::trUtf8(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitterHandle_SetOrientation(QSplitterHandle* self, uintptr_t o) {
	self->setOrientation(static_cast<Qt::Orientation>(o));
}

uintptr_t QSplitterHandle_Orientation(QSplitterHandle* self) {
	Qt::Orientation ret = const_cast<const QSplitterHandle*>(self)->orientation();
	return static_cast<uintptr_t>(ret);
}

bool QSplitterHandle_OpaqueResize(QSplitterHandle* self) {
	return const_cast<const QSplitterHandle*>(self)->opaqueResize();
}

QSplitter* QSplitterHandle_Splitter(QSplitterHandle* self) {
	return const_cast<const QSplitterHandle*>(self)->splitter();
}

QSize* QSplitterHandle_SizeHint(QSplitterHandle* self) {
	QSize ret = const_cast<const QSplitterHandle*>(self)->sizeHint();
	// Copy-construct value returned type into heap-allocated copy
	return static_cast<QSize*>(new QSize(ret));
}

void QSplitterHandle_Tr2(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QSplitterHandle::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitterHandle_Tr3(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QSplitterHandle::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitterHandle_TrUtf82(const char* s, const char* c, char** _out, int* _out_Strlen) {
	QString ret = QSplitterHandle::trUtf8(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitterHandle_TrUtf83(const char* s, const char* c, int n, char** _out, int* _out_Strlen) {
	QString ret = QSplitterHandle::trUtf8(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray b = ret.toUtf8();
	*_out = static_cast<char*>(malloc(b.length()));
	memcpy(*_out, b.data(), b.length());
	*_out_Strlen = b.length();
}

void QSplitterHandle_Delete(QSplitterHandle* self) {
	delete self;
}
