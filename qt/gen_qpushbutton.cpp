// +build ignore

#include <QAbstractButton>
#include <QEvent>
#include <QFocusEvent>
#include <QIcon>
#include <QKeyEvent>
#include <QMenu>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QPushButton>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionButton>
#include <QWidget>
#include <qpushbutton.h>
#include "gen_qpushbutton.h"
class MiqtVirtualQPushButton : public virtual QPushButton {
public:
MiqtVirtualQPushButton(QWidget* parent): QPushButton(parent) {};
MiqtVirtualQPushButton(): QPushButton() {};
MiqtVirtualQPushButton(const QString& text): QPushButton(text) {};
MiqtVirtualQPushButton(const QIcon& icon, const QString& text): QPushButton(icon, text) {};
MiqtVirtualQPushButton(const QString& text, QWidget* parent): QPushButton(text, parent) {};
MiqtVirtualQPushButton(const QIcon& icon, const QString& text, QWidget* parent): QPushButton(icon, text, parent) {};

virtual ~MiqtVirtualQPushButton() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPushButton::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPushButton_MetaObject(const_cast<MiqtVirtualQPushButton*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPushButton::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPushButton::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPushButton_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPushButton::qt_metacast(param1);

	}
};
QPushButton* QPushButton_new(QWidget* parent) {
return new MiqtVirtualQPushButton(parent);
}
QPushButton* QPushButton_new2() {
return new MiqtVirtualQPushButton();
}
QPushButton* QPushButton_new3(struct miqt_string text) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQPushButton(text_QString);
}
QPushButton* QPushButton_new4(QIcon* icon, struct miqt_string text) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQPushButton(*icon, text_QString);
}
QPushButton* QPushButton_new5(struct miqt_string text, QWidget* parent) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQPushButton(text_QString, parent);
}
QPushButton* QPushButton_new6(QIcon* icon, struct miqt_string text, QWidget* parent) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQPushButton(*icon, text_QString, parent);
}
void QPushButton_virtbase(QPushButton* src
, QAbstractButton** outptr_QAbstractButton
) {
*outptr_QAbstractButton = static_cast<QAbstractButton*>(src);
}
QMetaObject* QPushButton_MetaObject(const QPushButton* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPushButton_Metacast(QPushButton* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPushButton_Tr(const char* s) {
	QString _ret = QPushButton::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QPushButton_SizeHint(const QPushButton* self) {
	return new QSize(self->sizeHint());
}
QSize* QPushButton_MinimumSizeHint(const QPushButton* self) {
	return new QSize(self->minimumSizeHint());
}
bool QPushButton_AutoDefault(const QPushButton* self) {
	return self->autoDefault();
}
void QPushButton_SetAutoDefault(QPushButton* self, bool autoDefault) {
	self->setAutoDefault(autoDefault);
}
bool QPushButton_IsDefault(const QPushButton* self) {
	return self->isDefault();
}
void QPushButton_SetDefault(QPushButton* self, bool defaultVal) {
	self->setDefault(defaultVal);
}
void QPushButton_SetMenu(QPushButton* self, QMenu* menu) {
	self->setMenu(menu);
}
QMenu* QPushButton_Menu(const QPushButton* self) {
	return self->menu();
}
void QPushButton_SetFlat(QPushButton* self, bool flat) {
	self->setFlat(flat);
}
bool QPushButton_IsFlat(const QPushButton* self) {
	return self->isFlat();
}
void QPushButton_ShowMenu(QPushButton* self) {
	self->showMenu();
}
struct miqt_string QPushButton_Tr2(const char* s, const char* c) {
	QString _ret = QPushButton::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPushButton_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPushButton::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPushButton_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPushButton*>( (QPushButton*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPushButton_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPushButton*)(self) )->virtualbase_MetaObject();
}
void QPushButton_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPushButton*>( (QPushButton*)(self) )->handle__Metacast = slot;
}
void* QPushButton_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPushButton*)(self) )->virtualbase_Metacast(param1);
}
void QPushButton_Delete(QPushButton* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPushButton*>( self );
	} else {
		delete self;
	}
}
