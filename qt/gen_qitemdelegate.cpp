// +build ignore

#include <QAbstractItemDelegate>
#include <QAbstractItemModel>
#include <QEvent>
#include <QItemDelegate>
#include <QItemEditorFactory>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPainter>
#include <QPixmap>
#include <QRect>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionViewItem>
#include <QWidget>
#include <qitemdelegate.h>
#include "gen_qitemdelegate.h"

class MiqtVirtualQItemDelegate : public virtual QItemDelegate {
public:

	MiqtVirtualQItemDelegate(): QItemDelegate() {};
	MiqtVirtualQItemDelegate(QObject* parent): QItemDelegate(parent) {};

	virtual ~MiqtVirtualQItemDelegate() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QItemDelegate::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QItemDelegate_MetaObject(const_cast<MiqtVirtualQItemDelegate*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QItemDelegate::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QItemDelegate::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QItemDelegate_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QItemDelegate::qt_metacast(param1);

	}

};

QItemDelegate* QItemDelegate_new() {
	return new MiqtVirtualQItemDelegate();
}

QItemDelegate* QItemDelegate_new2(QObject* parent) {
	return new MiqtVirtualQItemDelegate(parent);
}

void QItemDelegate_virtbase(QItemDelegate* src, QAbstractItemDelegate** outptr_QAbstractItemDelegate) {
	*outptr_QAbstractItemDelegate = static_cast<QAbstractItemDelegate*>(src);
}

QMetaObject* QItemDelegate_MetaObject(const QItemDelegate* self) {
	return (QMetaObject*) self->metaObject();
}

void* QItemDelegate_Metacast(QItemDelegate* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QItemDelegate_Tr(const char* s) {
	QString _ret = QItemDelegate::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

bool QItemDelegate_HasClipping(const QItemDelegate* self) {
	return self->hasClipping();
}

void QItemDelegate_SetClipping(QItemDelegate* self, bool clip) {
	self->setClipping(clip);
}

void QItemDelegate_Paint(const QItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index) {
	self->paint(painter, *option, *index);
}

QSize* QItemDelegate_SizeHint(const QItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index) {
	return new QSize(self->sizeHint(*option, *index));
}

QWidget* QItemDelegate_CreateEditor(const QItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index) {
	return self->createEditor(parent, *option, *index);
}

void QItemDelegate_SetEditorData(const QItemDelegate* self, QWidget* editor, QModelIndex* index) {
	self->setEditorData(editor, *index);
}

void QItemDelegate_SetModelData(const QItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index) {
	self->setModelData(editor, model, *index);
}

void QItemDelegate_UpdateEditorGeometry(const QItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index) {
	self->updateEditorGeometry(editor, *option, *index);
}

QItemEditorFactory* QItemDelegate_ItemEditorFactory(const QItemDelegate* self) {
	return self->itemEditorFactory();
}

void QItemDelegate_SetItemEditorFactory(QItemDelegate* self, QItemEditorFactory* factory) {
	self->setItemEditorFactory(factory);
}

struct miqt_string QItemDelegate_Tr2(const char* s, const char* c) {
	QString _ret = QItemDelegate::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QItemDelegate_Tr3(const char* s, const char* c, int n) {
	QString _ret = QItemDelegate::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QItemDelegate_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQItemDelegate*>( (QItemDelegate*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QItemDelegate_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQItemDelegate*)(self) )->virtualbase_MetaObject();
}

void QItemDelegate_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQItemDelegate*>( (QItemDelegate*)(self) )->handle__Metacast = slot;
}

void* QItemDelegate_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQItemDelegate*)(self) )->virtualbase_Metacast(param1);
}

void QItemDelegate_Delete(QItemDelegate* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQItemDelegate*>( self );
	} else {
		delete self;
	}
}

