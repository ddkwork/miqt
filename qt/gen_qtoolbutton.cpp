// +build ignore

#include <QAbstractButton>
#include <QAction>
#include <QActionEvent>
#include <QEnterEvent>
#include <QEvent>
#include <QMenu>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionToolButton>
#include <QTimerEvent>
#include <QToolButton>
#include <QWidget>
#include <qtoolbutton.h>
#include "gen_qtoolbutton.h"
class MiqtVirtualQToolButton : public virtual QToolButton {
public:
MiqtVirtualQToolButton(QWidget* parent): QToolButton(parent) {};
MiqtVirtualQToolButton(): QToolButton() {};

virtual ~MiqtVirtualQToolButton() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QToolButton::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QToolButton_MetaObject(const_cast<MiqtVirtualQToolButton*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QToolButton::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QToolButton::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QToolButton_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QToolButton::qt_metacast(param1);

	}
};
QToolButton* QToolButton_new(QWidget* parent) {
return new MiqtVirtualQToolButton(parent);
}
QToolButton* QToolButton_new2() {
return new MiqtVirtualQToolButton();
}
void QToolButton_virtbase(QToolButton* src
, QAbstractButton** outptr_QAbstractButton
) {
*outptr_QAbstractButton = static_cast<QAbstractButton*>(src);
}
QMetaObject* QToolButton_MetaObject(const QToolButton* self) {
	return (QMetaObject*) self->metaObject();
}
void* QToolButton_Metacast(QToolButton* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QToolButton_Tr(const char* s) {
	QString _ret = QToolButton::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QToolButton_SizeHint(const QToolButton* self) {
	return new QSize(self->sizeHint());
}
QSize* QToolButton_MinimumSizeHint(const QToolButton* self) {
	return new QSize(self->minimumSizeHint());
}
int QToolButton_ToolButtonStyle(const QToolButton* self) {
	Qt::ToolButtonStyle _ret = self->toolButtonStyle();
	return static_cast<int>(_ret);
}
int QToolButton_ArrowType(const QToolButton* self) {
	Qt::ArrowType _ret = self->arrowType();
	return static_cast<int>(_ret);
}
void QToolButton_SetArrowType(QToolButton* self, int typeVal) {
	self->setArrowType(static_cast<Qt::ArrowType>(typeVal));
}
void QToolButton_SetMenu(QToolButton* self, QMenu* menu) {
	self->setMenu(menu);
}
QMenu* QToolButton_Menu(const QToolButton* self) {
	return self->menu();
}
void QToolButton_SetPopupMode(QToolButton* self, ToolButtonPopupMode mode) {
	self->setPopupMode(mode);
}
ToolButtonPopupMode QToolButton_PopupMode(const QToolButton* self) {
	return self->popupMode();
}
QAction* QToolButton_DefaultAction(const QToolButton* self) {
	return self->defaultAction();
}
void QToolButton_SetAutoRaise(QToolButton* self, bool enable) {
	self->setAutoRaise(enable);
}
bool QToolButton_AutoRaise(const QToolButton* self) {
	return self->autoRaise();
}
void QToolButton_ShowMenu(QToolButton* self) {
	self->showMenu();
}
void QToolButton_SetToolButtonStyle(QToolButton* self, int style) {
	self->setToolButtonStyle(static_cast<Qt::ToolButtonStyle>(style));
}
void QToolButton_SetDefaultAction(QToolButton* self, QAction* defaultAction) {
	self->setDefaultAction(defaultAction);
}
void QToolButton_Triggered(QToolButton* self, QAction* param1) {
	self->triggered(param1);
}
void QToolButton_connect_Triggered(QToolButton* self, intptr_t slot) {
	MiqtVirtualQToolButton::connect(self, static_cast<void (QToolButton::*)(QAction*)>(&QToolButton::triggered), self, [=](QAction* param1) {
		QAction* sigval1 = param1;
		miqt_exec_callback_QToolButton_Triggered(slot, sigval1);
	});
}
struct miqt_string QToolButton_Tr2(const char* s, const char* c) {
	QString _ret = QToolButton::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QToolButton_Tr3(const char* s, const char* c, int n) {
	QString _ret = QToolButton::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QToolButton_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQToolButton*>( (QToolButton*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QToolButton_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQToolButton*)(self) )->virtualbase_MetaObject();
}
void QToolButton_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQToolButton*>( (QToolButton*)(self) )->handle__Metacast = slot;
}
void* QToolButton_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQToolButton*)(self) )->virtualbase_Metacast(param1);
}
void QToolButton_Delete(QToolButton* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQToolButton*>( self );
	} else {
		delete self;
	}
}
