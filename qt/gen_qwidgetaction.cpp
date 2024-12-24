// +build ignore

#include <QAction>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <QWidgetAction>
#include <qwidgetaction.h>
#include "gen_qwidgetaction.h"
class MiqtVirtualQWidgetAction : public virtual QWidgetAction {
public:
MiqtVirtualQWidgetAction(QObject* parent): QWidgetAction(parent) {};

virtual ~MiqtVirtualQWidgetAction() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QWidgetAction::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QWidgetAction_MetaObject(const_cast<MiqtVirtualQWidgetAction*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QWidgetAction::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QWidgetAction::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QWidgetAction_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QWidgetAction::qt_metacast(param1);

	}
};
QWidgetAction* QWidgetAction_new(QObject* parent) {
return new MiqtVirtualQWidgetAction(parent);
}
void QWidgetAction_virtbase(QWidgetAction* src
, QAction** outptr_QAction
) {
*outptr_QAction = static_cast<QAction*>(src);
}
QMetaObject* QWidgetAction_MetaObject(const QWidgetAction* self) {
	return (QMetaObject*) self->metaObject();
}
void* QWidgetAction_Metacast(QWidgetAction* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QWidgetAction_Tr(const char* s) {
	QString _ret = QWidgetAction::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWidgetAction_SetDefaultWidget(QWidgetAction* self, QWidget* w) {
	self->setDefaultWidget(w);
}
QWidget* QWidgetAction_DefaultWidget(const QWidgetAction* self) {
	return self->defaultWidget();
}
QWidget* QWidgetAction_RequestWidget(QWidgetAction* self, QWidget* parent) {
	return self->requestWidget(parent);
}
void QWidgetAction_ReleaseWidget(QWidgetAction* self, QWidget* widget) {
	self->releaseWidget(widget);
}
struct miqt_string QWidgetAction_Tr2(const char* s, const char* c) {
	QString _ret = QWidgetAction::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QWidgetAction_Tr3(const char* s, const char* c, int n) {
	QString _ret = QWidgetAction::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QWidgetAction_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetAction*>( (QWidgetAction*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QWidgetAction_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQWidgetAction*)(self) )->virtualbase_MetaObject();
}
void QWidgetAction_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQWidgetAction*>( (QWidgetAction*)(self) )->handle__Metacast = slot;
}
void* QWidgetAction_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQWidgetAction*)(self) )->virtualbase_Metacast(param1);
}
void QWidgetAction_Delete(QWidgetAction* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQWidgetAction*>( self );
	} else {
		delete self;
	}
}
