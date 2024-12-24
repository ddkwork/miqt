// +build ignore

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QFrame>
#include <QList>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QSize>
#include <QSplitter>
#include <QSplitterHandle>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTextStream>
#include <QWidget>
#include <qsplitter.h>
#include "gen_qsplitter.h"
void QTextStream_Delete(QTextStream* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<QTextStream*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQSplitter : public virtual QSplitter {
public:
MiqtVirtualQSplitter(QWidget* parent): QSplitter(parent) {};
MiqtVirtualQSplitter(): QSplitter() {};
MiqtVirtualQSplitter(Qt::Orientation param1): QSplitter(param1) {};
MiqtVirtualQSplitter(Qt::Orientation param1, QWidget* parent): QSplitter(param1, parent) {};

virtual ~MiqtVirtualQSplitter() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSplitter::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSplitter_MetaObject(const_cast<MiqtVirtualQSplitter*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSplitter::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSplitter::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSplitter_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSplitter::qt_metacast(param1);

	}
};
QSplitter* QSplitter_new(QWidget* parent) {
return new MiqtVirtualQSplitter(parent);
}
QSplitter* QSplitter_new2() {
return new MiqtVirtualQSplitter();
}
QSplitter* QSplitter_new3(int param1) {
return new MiqtVirtualQSplitter(static_cast<Qt::Orientation>(param1));
}
QSplitter* QSplitter_new4(int param1, QWidget* parent) {
return new MiqtVirtualQSplitter(static_cast<Qt::Orientation>(param1), parent);
}
void QSplitter_virtbase(QSplitter* src
, QFrame** outptr_QFrame
) {
*outptr_QFrame = static_cast<QFrame*>(src);
}
QMetaObject* QSplitter_MetaObject(const QSplitter* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSplitter_Metacast(QSplitter* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSplitter_Tr(const char* s) {
	QString _ret = QSplitter::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
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
void QSplitter_SetOrientation(QSplitter* self, int orientation) {
	self->setOrientation(static_cast<Qt::Orientation>(orientation));
}
int QSplitter_Orientation(const QSplitter* self) {
	Qt::Orientation _ret = self->orientation();
	return static_cast<int>(_ret);
}
void QSplitter_SetChildrenCollapsible(QSplitter* self, bool childrenCollapsible) {
	self->setChildrenCollapsible(childrenCollapsible);
}
bool QSplitter_ChildrenCollapsible(const QSplitter* self) {
	return self->childrenCollapsible();
}
void QSplitter_SetCollapsible(QSplitter* self, int index, bool param2) {
	self->setCollapsible(static_cast<int>(index), param2);
}
bool QSplitter_IsCollapsible(const QSplitter* self, int index) {
	return self->isCollapsible(static_cast<int>(index));
}
void QSplitter_SetOpaqueResize(QSplitter* self) {
	self->setOpaqueResize();
}
bool QSplitter_OpaqueResize(const QSplitter* self) {
	return self->opaqueResize();
}
void QSplitter_Refresh(QSplitter* self) {
	self->refresh();
}
QSize* QSplitter_SizeHint(const QSplitter* self) {
	return new QSize(self->sizeHint());
}
QSize* QSplitter_MinimumSizeHint(const QSplitter* self) {
	return new QSize(self->minimumSizeHint());
}
struct miqt_array /* of int */  QSplitter_Sizes(const QSplitter* self) {
	QList<int> _ret = self->sizes();
	// Convert QList<> from C++ memory to manually-managed C memory
	int* _arr = static_cast<int*>(malloc(sizeof(int) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}
void QSplitter_SetSizes(QSplitter* self, struct miqt_array /* of int */  list) {
	QList<int> list_QList;
	list_QList.reserve(list.len);
	int* list_arr = static_cast<int*>(list.data);
	for(size_t i = 0; i < list.len; ++i) {
		list_QList.push_back(static_cast<int>(list_arr[i]));
	}
	self->setSizes(list_QList);
}
struct miqt_string QSplitter_SaveState(const QSplitter* self) {
	QByteArray _qb = self->saveState();
	struct miqt_string _ms;
	_ms.len = _qb.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _qb.data(), _ms.len);
	return _ms;
}
bool QSplitter_RestoreState(QSplitter* self, struct miqt_string state) {
	QByteArray state_QByteArray(state.data, state.len);
	return self->restoreState(state_QByteArray);
}
int QSplitter_HandleWidth(const QSplitter* self) {
	return self->handleWidth();
}
void QSplitter_SetHandleWidth(QSplitter* self, int handleWidth) {
	self->setHandleWidth(static_cast<int>(handleWidth));
}
int QSplitter_IndexOf(const QSplitter* self, QWidget* w) {
	return self->indexOf(w);
}
QWidget* QSplitter_Widget(const QSplitter* self, int index) {
	return self->widget(static_cast<int>(index));
}
int QSplitter_Count(const QSplitter* self) {
	return self->count();
}
void QSplitter_GetRange(const QSplitter* self, int index, int* param2, int* param3) {
	self->getRange(static_cast<int>(index), static_cast<int*>(param2), static_cast<int*>(param3));
}
QSplitterHandle* QSplitter_Handle(const QSplitter* self, int index) {
	return self->handle(static_cast<int>(index));
}
void QSplitter_SetStretchFactor(QSplitter* self, int index, int stretch) {
	self->setStretchFactor(static_cast<int>(index), static_cast<int>(stretch));
}
void QSplitter_SplitterMoved(QSplitter* self, int pos, int index) {
	self->splitterMoved(static_cast<int>(pos), static_cast<int>(index));
}
void QSplitter_connect_SplitterMoved(QSplitter* self, intptr_t slot) {
	MiqtVirtualQSplitter::connect(self, static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), self, [=](int pos, int index) {
		int sigval1 = pos;
		int sigval2 = index;
		miqt_exec_callback_QSplitter_SplitterMoved(slot, sigval1, sigval2);
	});
}
struct miqt_string QSplitter_Tr2(const char* s, const char* c) {
	QString _ret = QSplitter::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSplitter_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSplitter::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSplitter_SetOpaqueResize1(QSplitter* self, bool opaque) {
	self->setOpaqueResize(opaque);
}
void QSplitter_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSplitter*>( (QSplitter*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSplitter_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSplitter*)(self) )->virtualbase_MetaObject();
}
void QSplitter_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSplitter*>( (QSplitter*)(self) )->handle__Metacast = slot;
}
void* QSplitter_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSplitter*)(self) )->virtualbase_Metacast(param1);
}
void QSplitter_Delete(QSplitter* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSplitter*>( self );
	} else {
		delete self;
	}
}
class MiqtVirtualQSplitterHandle : public virtual QSplitterHandle {
public:
MiqtVirtualQSplitterHandle(Qt::Orientation o, QSplitter* parent): QSplitterHandle(o, parent) {};

virtual ~MiqtVirtualQSplitterHandle() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSplitterHandle::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSplitterHandle_MetaObject(const_cast<MiqtVirtualQSplitterHandle*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSplitterHandle::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSplitterHandle::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSplitterHandle_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSplitterHandle::qt_metacast(param1);

	}
};
QSplitterHandle* QSplitterHandle_new(int o, QSplitter* parent) {
return new MiqtVirtualQSplitterHandle(static_cast<Qt::Orientation>(o), parent);
}
void QSplitterHandle_virtbase(QSplitterHandle* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QSplitterHandle_MetaObject(const QSplitterHandle* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSplitterHandle_Metacast(QSplitterHandle* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSplitterHandle_Tr(const char* s) {
	QString _ret = QSplitterHandle::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSplitterHandle_SetOrientation(QSplitterHandle* self, int o) {
	self->setOrientation(static_cast<Qt::Orientation>(o));
}
int QSplitterHandle_Orientation(const QSplitterHandle* self) {
	Qt::Orientation _ret = self->orientation();
	return static_cast<int>(_ret);
}
bool QSplitterHandle_OpaqueResize(const QSplitterHandle* self) {
	return self->opaqueResize();
}
QSplitter* QSplitterHandle_Splitter(const QSplitterHandle* self) {
	return self->splitter();
}
QSize* QSplitterHandle_SizeHint(const QSplitterHandle* self) {
	return new QSize(self->sizeHint());
}
struct miqt_string QSplitterHandle_Tr2(const char* s, const char* c) {
	QString _ret = QSplitterHandle::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSplitterHandle_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSplitterHandle::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSplitterHandle_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSplitterHandle*>( (QSplitterHandle*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSplitterHandle_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSplitterHandle*)(self) )->virtualbase_MetaObject();
}
void QSplitterHandle_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSplitterHandle*>( (QSplitterHandle*)(self) )->handle__Metacast = slot;
}
void* QSplitterHandle_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSplitterHandle*)(self) )->virtualbase_Metacast(param1);
}
void QSplitterHandle_Delete(QSplitterHandle* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSplitterHandle*>( self );
	} else {
		delete self;
	}
}
