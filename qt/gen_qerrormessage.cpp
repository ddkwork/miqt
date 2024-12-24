// +build ignore

#include <QDialog>
#include <QErrorMessage>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qerrormessage.h>
#include "gen_qerrormessage.h"
class MiqtVirtualQErrorMessage : public virtual QErrorMessage {
public:
MiqtVirtualQErrorMessage(QWidget* parent): QErrorMessage(parent) {};
MiqtVirtualQErrorMessage(): QErrorMessage() {};

virtual ~MiqtVirtualQErrorMessage() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QErrorMessage::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QErrorMessage_MetaObject(const_cast<MiqtVirtualQErrorMessage*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QErrorMessage::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QErrorMessage::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QErrorMessage_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QErrorMessage::qt_metacast(param1);

	}
};
QErrorMessage* QErrorMessage_new(QWidget* parent) {
return new MiqtVirtualQErrorMessage(parent);
}
QErrorMessage* QErrorMessage_new2() {
return new MiqtVirtualQErrorMessage();
}
void QErrorMessage_virtbase(QErrorMessage* src
, QDialog** outptr_QDialog
) {
*outptr_QDialog = static_cast<QDialog*>(src);
}
QMetaObject* QErrorMessage_MetaObject(const QErrorMessage* self) {
	return (QMetaObject*) self->metaObject();
}
void* QErrorMessage_Metacast(QErrorMessage* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QErrorMessage_Tr(const char* s) {
	QString _ret = QErrorMessage::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QErrorMessage* QErrorMessage_QtHandler() {
	return QErrorMessage::qtHandler();
}
void QErrorMessage_ShowMessage(QErrorMessage* self, struct miqt_string message) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->showMessage(message_QString);
}
void QErrorMessage_ShowMessage2(QErrorMessage* self, struct miqt_string message, struct miqt_string typeVal) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	QString typeVal_QString = QString::fromUtf8(typeVal.data, typeVal.len);
	self->showMessage(message_QString, typeVal_QString);
}
struct miqt_string QErrorMessage_Tr2(const char* s, const char* c) {
	QString _ret = QErrorMessage::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QErrorMessage_Tr3(const char* s, const char* c, int n) {
	QString _ret = QErrorMessage::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QErrorMessage_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQErrorMessage*>( (QErrorMessage*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QErrorMessage_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQErrorMessage*)(self) )->virtualbase_MetaObject();
}
void QErrorMessage_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQErrorMessage*>( (QErrorMessage*)(self) )->handle__Metacast = slot;
}
void* QErrorMessage_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQErrorMessage*)(self) )->virtualbase_Metacast(param1);
}
void QErrorMessage_Delete(QErrorMessage* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQErrorMessage*>( self );
	} else {
		delete self;
	}
}
