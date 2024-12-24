// +build ignore

#include <QDialog>
#include <QMetaObject>
#include <QObject>
#include <QPageSetupDialog>
#include <QPaintDevice>
#include <QPrinter>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qpagesetupdialog.h>
#include "gen_qpagesetupdialog.h"
class MiqtVirtualQPageSetupDialog : public virtual QPageSetupDialog {
public:
MiqtVirtualQPageSetupDialog(QWidget* parent): QPageSetupDialog(parent) {};
MiqtVirtualQPageSetupDialog(QPrinter* printer): QPageSetupDialog(printer) {};
MiqtVirtualQPageSetupDialog(): QPageSetupDialog() {};
MiqtVirtualQPageSetupDialog(QPrinter* printer, QWidget* parent): QPageSetupDialog(printer, parent) {};

virtual ~MiqtVirtualQPageSetupDialog() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QPageSetupDialog::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QPageSetupDialog_MetaObject(const_cast<MiqtVirtualQPageSetupDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPageSetupDialog::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QPageSetupDialog::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPageSetupDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPageSetupDialog::qt_metacast(param1);

	}
};
QPageSetupDialog* QPageSetupDialog_new(QWidget* parent) {
return new MiqtVirtualQPageSetupDialog(parent);
}
QPageSetupDialog* QPageSetupDialog_new2(QPrinter* printer) {
return new MiqtVirtualQPageSetupDialog(printer);
}
QPageSetupDialog* QPageSetupDialog_new3() {
return new MiqtVirtualQPageSetupDialog();
}
QPageSetupDialog* QPageSetupDialog_new4(QPrinter* printer, QWidget* parent) {
return new MiqtVirtualQPageSetupDialog(printer, parent);
}
void QPageSetupDialog_virtbase(QPageSetupDialog* src
, QDialog** outptr_QDialog
) {
*outptr_QDialog = static_cast<QDialog*>(src);
}
QMetaObject* QPageSetupDialog_MetaObject(const QPageSetupDialog* self) {
	return (QMetaObject*) self->metaObject();
}
void* QPageSetupDialog_Metacast(QPageSetupDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QPageSetupDialog_Tr(const char* s) {
	QString _ret = QPageSetupDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPageSetupDialog_SetVisible(QPageSetupDialog* self, bool visible) {
	self->setVisible(visible);
}
int QPageSetupDialog_Exec(QPageSetupDialog* self) {
	return self->exec();
}
void QPageSetupDialog_Done(QPageSetupDialog* self, int result) {
	self->done(static_cast<int>(result));
}
QPrinter* QPageSetupDialog_Printer(QPageSetupDialog* self) {
	return self->printer();
}
struct miqt_string QPageSetupDialog_Tr2(const char* s, const char* c) {
	QString _ret = QPageSetupDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QPageSetupDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPageSetupDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QPageSetupDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPageSetupDialog*>( (QPageSetupDialog*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QPageSetupDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPageSetupDialog*)(self) )->virtualbase_MetaObject();
}
void QPageSetupDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPageSetupDialog*>( (QPageSetupDialog*)(self) )->handle__Metacast = slot;
}
void* QPageSetupDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPageSetupDialog*)(self) )->virtualbase_Metacast(param1);
}
void QPageSetupDialog_Delete(QPageSetupDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPageSetupDialog*>( self );
	} else {
		delete self;
	}
}
