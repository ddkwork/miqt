// +build ignore

#include <QAbstractItemView>
#include <QAbstractScrollArea>
#include <QFrame>
#include <QIcon>
#include <QListView>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QUndoGroup>
#include <QUndoStack>
#include <QUndoView>
#include <QWidget>
#include <qundoview.h>
#include "gen_qundoview.h"

class MiqtVirtualQUndoView : public virtual QUndoView {
public:

	MiqtVirtualQUndoView(QWidget* parent): QUndoView(parent) {};
	MiqtVirtualQUndoView(): QUndoView() {};
	MiqtVirtualQUndoView(QUndoStack* stack): QUndoView(stack) {};
	MiqtVirtualQUndoView(QUndoGroup* group): QUndoView(group) {};
	MiqtVirtualQUndoView(QUndoStack* stack, QWidget* parent): QUndoView(stack, parent) {};
	MiqtVirtualQUndoView(QUndoGroup* group, QWidget* parent): QUndoView(group, parent) {};

	virtual ~MiqtVirtualQUndoView() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QUndoView::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QUndoView_MetaObject(const_cast<MiqtVirtualQUndoView*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QUndoView::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QUndoView::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QUndoView_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QUndoView::qt_metacast(param1);

	}

};

QUndoView* QUndoView_new(QWidget* parent) {
	return new MiqtVirtualQUndoView(parent);
}

QUndoView* QUndoView_new2() {
	return new MiqtVirtualQUndoView();
}

QUndoView* QUndoView_new3(QUndoStack* stack) {
	return new MiqtVirtualQUndoView(stack);
}

QUndoView* QUndoView_new4(QUndoGroup* group) {
	return new MiqtVirtualQUndoView(group);
}

QUndoView* QUndoView_new5(QUndoStack* stack, QWidget* parent) {
	return new MiqtVirtualQUndoView(stack, parent);
}

QUndoView* QUndoView_new6(QUndoGroup* group, QWidget* parent) {
	return new MiqtVirtualQUndoView(group, parent);
}

void QUndoView_virtbase(QUndoView* src, QListView** outptr_QListView) {
	*outptr_QListView = static_cast<QListView*>(src);
}

QMetaObject* QUndoView_MetaObject(const QUndoView* self) {
	return (QMetaObject*) self->metaObject();
}

void* QUndoView_Metacast(QUndoView* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QUndoView_Tr(const char* s) {
	QString _ret = QUndoView::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QUndoStack* QUndoView_Stack(const QUndoView* self) {
	return self->stack();
}

QUndoGroup* QUndoView_Group(const QUndoView* self) {
	return self->group();
}

void QUndoView_SetEmptyLabel(QUndoView* self, struct miqt_string label) {
	QString label_QString = QString::fromUtf8(label.data, label.len);
	self->setEmptyLabel(label_QString);
}

struct miqt_string QUndoView_EmptyLabel(const QUndoView* self) {
	QString _ret = self->emptyLabel();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUndoView_SetCleanIcon(QUndoView* self, QIcon* icon) {
	self->setCleanIcon(*icon);
}

QIcon* QUndoView_CleanIcon(const QUndoView* self) {
	return new QIcon(self->cleanIcon());
}

void QUndoView_SetStack(QUndoView* self, QUndoStack* stack) {
	self->setStack(stack);
}

void QUndoView_SetGroup(QUndoView* self, QUndoGroup* group) {
	self->setGroup(group);
}

struct miqt_string QUndoView_Tr2(const char* s, const char* c) {
	QString _ret = QUndoView::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QUndoView_Tr3(const char* s, const char* c, int n) {
	QString _ret = QUndoView::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QUndoView_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQUndoView*>( (QUndoView*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QUndoView_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQUndoView*)(self) )->virtualbase_MetaObject();
}

void QUndoView_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQUndoView*>( (QUndoView*)(self) )->handle__Metacast = slot;
}

void* QUndoView_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQUndoView*)(self) )->virtualbase_Metacast(param1);
}

void QUndoView_Delete(QUndoView* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQUndoView*>( self );
	} else {
		delete self;
	}
}

