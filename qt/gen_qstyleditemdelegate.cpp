// +build ignore

#include <QAbstractItemDelegate>
#include <QAbstractItemModel>
#include <QEvent>
#include <QItemEditorFactory>
#include <QLocale>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPainter>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionViewItem>
#include <QStyledItemDelegate>
#include <QVariant>
#include <QWidget>
#include <qstyleditemdelegate.h>
#include "gen_qstyleditemdelegate.h"
class MiqtVirtualQStyledItemDelegate : public virtual QStyledItemDelegate {
public:
MiqtVirtualQStyledItemDelegate(): QStyledItemDelegate() {};
MiqtVirtualQStyledItemDelegate(QObject* parent): QStyledItemDelegate(parent) {};

virtual ~MiqtVirtualQStyledItemDelegate() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QStyledItemDelegate::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QStyledItemDelegate_MetaObject(const_cast<MiqtVirtualQStyledItemDelegate*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStyledItemDelegate::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QStyledItemDelegate::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStyledItemDelegate_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStyledItemDelegate::qt_metacast(param1);

	}
};
QStyledItemDelegate* QStyledItemDelegate_new() {
return new MiqtVirtualQStyledItemDelegate();
}
QStyledItemDelegate* QStyledItemDelegate_new2(QObject* parent) {
return new MiqtVirtualQStyledItemDelegate(parent);
}
void QStyledItemDelegate_virtbase(QStyledItemDelegate* src
, QAbstractItemDelegate** outptr_QAbstractItemDelegate
) {
*outptr_QAbstractItemDelegate = static_cast<QAbstractItemDelegate*>(src);
}
QMetaObject* QStyledItemDelegate_MetaObject(const QStyledItemDelegate* self) {
	return (QMetaObject*) self->metaObject();
}
void* QStyledItemDelegate_Metacast(QStyledItemDelegate* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QStyledItemDelegate_Tr(const char* s) {
	QString _ret = QStyledItemDelegate::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStyledItemDelegate_Paint(const QStyledItemDelegate* self, QPainter* painter, QStyleOptionViewItem* option, QModelIndex* index) {
	self->paint(painter, *option, *index);
}
QSize* QStyledItemDelegate_SizeHint(const QStyledItemDelegate* self, QStyleOptionViewItem* option, QModelIndex* index) {
	return new QSize(self->sizeHint(*option, *index));
}
QWidget* QStyledItemDelegate_CreateEditor(const QStyledItemDelegate* self, QWidget* parent, QStyleOptionViewItem* option, QModelIndex* index) {
	return self->createEditor(parent, *option, *index);
}
void QStyledItemDelegate_SetEditorData(const QStyledItemDelegate* self, QWidget* editor, QModelIndex* index) {
	self->setEditorData(editor, *index);
}
void QStyledItemDelegate_SetModelData(const QStyledItemDelegate* self, QWidget* editor, QAbstractItemModel* model, QModelIndex* index) {
	self->setModelData(editor, model, *index);
}
void QStyledItemDelegate_UpdateEditorGeometry(const QStyledItemDelegate* self, QWidget* editor, QStyleOptionViewItem* option, QModelIndex* index) {
	self->updateEditorGeometry(editor, *option, *index);
}
QItemEditorFactory* QStyledItemDelegate_ItemEditorFactory(const QStyledItemDelegate* self) {
	return self->itemEditorFactory();
}
void QStyledItemDelegate_SetItemEditorFactory(QStyledItemDelegate* self, QItemEditorFactory* factory) {
	self->setItemEditorFactory(factory);
}
struct miqt_string QStyledItemDelegate_DisplayText(const QStyledItemDelegate* self, QVariant* value, QLocale* locale) {
	QString _ret = self->displayText(*value, *locale);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStyledItemDelegate_Tr2(const char* s, const char* c) {
	QString _ret = QStyledItemDelegate::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStyledItemDelegate_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStyledItemDelegate::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStyledItemDelegate_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStyledItemDelegate*>( (QStyledItemDelegate*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QStyledItemDelegate_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStyledItemDelegate*)(self) )->virtualbase_MetaObject();
}
void QStyledItemDelegate_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStyledItemDelegate*>( (QStyledItemDelegate*)(self) )->handle__Metacast = slot;
}
void* QStyledItemDelegate_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStyledItemDelegate*)(self) )->virtualbase_Metacast(param1);
}
void QStyledItemDelegate_Delete(QStyledItemDelegate* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStyledItemDelegate*>( self );
	} else {
		delete self;
	}
}
