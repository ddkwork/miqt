// +build ignore

#include <QEvent>
#include <QFrame>
#include <QIcon>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QShowEvent>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QToolBox>
#include <QWidget>
#include <qtoolbox.h>
#include "gen_qtoolbox.h"
class MiqtVirtualQToolBox : public virtual QToolBox {
public:
MiqtVirtualQToolBox(QWidget* parent): QToolBox(parent) {};
MiqtVirtualQToolBox(): QToolBox() {};
MiqtVirtualQToolBox(QWidget* parent, Qt::WindowFlags f): QToolBox(parent, f) {};

virtual ~MiqtVirtualQToolBox() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QToolBox::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QToolBox_MetaObject(const_cast<MiqtVirtualQToolBox*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QToolBox::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QToolBox::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QToolBox_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QToolBox::qt_metacast(param1);

	}
};
QToolBox* QToolBox_new(QWidget* parent) {
return new MiqtVirtualQToolBox(parent);
}
QToolBox* QToolBox_new2() {
return new MiqtVirtualQToolBox();
}
QToolBox* QToolBox_new3(QWidget* parent, int f) {
return new MiqtVirtualQToolBox(parent, static_cast<Qt::WindowFlags>(f));
}
void QToolBox_virtbase(QToolBox* src
, QFrame** outptr_QFrame
) {
*outptr_QFrame = static_cast<QFrame*>(src);
}
QMetaObject* QToolBox_MetaObject(const QToolBox* self) {
	return (QMetaObject*) self->metaObject();
}
void* QToolBox_Metacast(QToolBox* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QToolBox_Tr(const char* s) {
	QString _ret = QToolBox::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QToolBox_AddItem(QToolBox* self, QWidget* widget, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addItem(widget, text_QString);
}
int QToolBox_AddItem2(QToolBox* self, QWidget* widget, QIcon* icon, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addItem(widget, *icon, text_QString);
}
int QToolBox_InsertItem(QToolBox* self, int index, QWidget* widget, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->insertItem(static_cast<int>(index), widget, text_QString);
}
int QToolBox_InsertItem2(QToolBox* self, int index, QWidget* widget, QIcon* icon, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->insertItem(static_cast<int>(index), widget, *icon, text_QString);
}
void QToolBox_RemoveItem(QToolBox* self, int index) {
	self->removeItem(static_cast<int>(index));
}
void QToolBox_SetItemEnabled(QToolBox* self, int index, bool enabled) {
	self->setItemEnabled(static_cast<int>(index), enabled);
}
bool QToolBox_IsItemEnabled(const QToolBox* self, int index) {
	return self->isItemEnabled(static_cast<int>(index));
}
void QToolBox_SetItemText(QToolBox* self, int index, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setItemText(static_cast<int>(index), text_QString);
}
struct miqt_string QToolBox_ItemText(const QToolBox* self, int index) {
	QString _ret = self->itemText(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QToolBox_SetItemIcon(QToolBox* self, int index, QIcon* icon) {
	self->setItemIcon(static_cast<int>(index), *icon);
}
QIcon* QToolBox_ItemIcon(const QToolBox* self, int index) {
	return new QIcon(self->itemIcon(static_cast<int>(index)));
}
void QToolBox_SetItemToolTip(QToolBox* self, int index, struct miqt_string toolTip) {
	QString toolTip_QString = QString::fromUtf8(toolTip.data, toolTip.len);
	self->setItemToolTip(static_cast<int>(index), toolTip_QString);
}
struct miqt_string QToolBox_ItemToolTip(const QToolBox* self, int index) {
	QString _ret = self->itemToolTip(static_cast<int>(index));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QToolBox_CurrentIndex(const QToolBox* self) {
	return self->currentIndex();
}
QWidget* QToolBox_CurrentWidget(const QToolBox* self) {
	return self->currentWidget();
}
QWidget* QToolBox_Widget(const QToolBox* self, int index) {
	return self->widget(static_cast<int>(index));
}
int QToolBox_IndexOf(const QToolBox* self, QWidget* widget) {
	return self->indexOf(widget);
}
int QToolBox_Count(const QToolBox* self) {
	return self->count();
}
void QToolBox_SetCurrentIndex(QToolBox* self, int index) {
	self->setCurrentIndex(static_cast<int>(index));
}
void QToolBox_SetCurrentWidget(QToolBox* self, QWidget* widget) {
	self->setCurrentWidget(widget);
}
void QToolBox_CurrentChanged(QToolBox* self, int index) {
	self->currentChanged(static_cast<int>(index));
}
void QToolBox_connect_CurrentChanged(QToolBox* self, intptr_t slot) {
	MiqtVirtualQToolBox::connect(self, static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), self, [=](int index) {
		int sigval1 = index;
		miqt_exec_callback_QToolBox_CurrentChanged(slot, sigval1);
	});
}
struct miqt_string QToolBox_Tr2(const char* s, const char* c) {
	QString _ret = QToolBox::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QToolBox_Tr3(const char* s, const char* c, int n) {
	QString _ret = QToolBox::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QToolBox_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQToolBox*>( (QToolBox*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QToolBox_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQToolBox*)(self) )->virtualbase_MetaObject();
}
void QToolBox_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQToolBox*>( (QToolBox*)(self) )->handle__Metacast = slot;
}
void* QToolBox_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQToolBox*)(self) )->virtualbase_Metacast(param1);
}
void QToolBox_Delete(QToolBox* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQToolBox*>( self );
	} else {
		delete self;
	}
}
