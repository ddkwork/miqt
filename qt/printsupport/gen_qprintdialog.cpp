// +build ignore

#include <QAbstractPrintDialog>
#include <QDialog>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPrintDialog>
#include <QPrinter>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qprintdialog.h>
#include "gen_qprintdialog.h"

class MiqtVirtualQPrintDialog : public virtual QPrintDialog {
public:

	MiqtVirtualQPrintDialog(QWidget* parent): QPrintDialog(parent) {};
	MiqtVirtualQPrintDialog(QPrinter* printer): QPrintDialog(printer) {};
	MiqtVirtualQPrintDialog(): QPrintDialog() {};
	MiqtVirtualQPrintDialog(QPrinter* printer, QWidget* parent): QPrintDialog(printer, parent) {};

	virtual ~MiqtVirtualQPrintDialog() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QPrintDialog::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QPrintDialog_MetaObject(const_cast<MiqtVirtualQPrintDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPrintDialog::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QPrintDialog::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPrintDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPrintDialog::qt_metacast(param1);

	}

};

QPrintDialog* QPrintDialog_new(QWidget* parent) {
	return new MiqtVirtualQPrintDialog(parent);
}

QPrintDialog* QPrintDialog_new2(QPrinter* printer) {
	return new MiqtVirtualQPrintDialog(printer);
}

QPrintDialog* QPrintDialog_new3() {
	return new MiqtVirtualQPrintDialog();
}

QPrintDialog* QPrintDialog_new4(QPrinter* printer, QWidget* parent) {
	return new MiqtVirtualQPrintDialog(printer, parent);
}

void QPrintDialog_virtbase(QPrintDialog* src, QAbstractPrintDialog** outptr_QAbstractPrintDialog) {
	*outptr_QAbstractPrintDialog = static_cast<QAbstractPrintDialog*>(src);
}

QMetaObject* QPrintDialog_MetaObject(const QPrintDialog* self) {
	return (QMetaObject*) self->metaObject();
}

void* QPrintDialog_Metacast(QPrintDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QPrintDialog_Tr(const char* s) {
	QString _ret = QPrintDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

int QPrintDialog_Exec(QPrintDialog* self) {
	return self->exec();
}

void QPrintDialog_Done(QPrintDialog* self, int result) {
	self->done(static_cast<int>(result));
}

void QPrintDialog_SetOption(QPrintDialog* self, PrintDialogOption option) {
	self->setOption(option);
}

bool QPrintDialog_TestOption(const QPrintDialog* self, PrintDialogOption option) {
	return self->testOption(option);
}

void QPrintDialog_SetOptions(QPrintDialog* self, PrintDialogOptions options) {
	self->setOptions(options);
}

PrintDialogOptions QPrintDialog_Options(const QPrintDialog* self) {
	return self->options();
}

void QPrintDialog_SetVisible(QPrintDialog* self, bool visible) {
	self->setVisible(visible);
}

void QPrintDialog_Accepted(QPrintDialog* self, QPrinter* printer) {
	self->accepted(printer);
}

void QPrintDialog_connect_Accepted(QPrintDialog* self, intptr_t slot) {
	MiqtVirtualQPrintDialog::connect(self, static_cast<void (QPrintDialog::*)(QPrinter*)>(&QPrintDialog::accepted), self, [=](QPrinter* printer) {
		QPrinter* sigval1 = printer;
		miqt_exec_callback_QPrintDialog_Accepted(slot, sigval1);
	});
}

struct miqt_string QPrintDialog_Tr2(const char* s, const char* c) {
	QString _ret = QPrintDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPrintDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPrintDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrintDialog_SetOption2(QPrintDialog* self, PrintDialogOption option, bool on) {
	self->setOption(option, on);
}

void QPrintDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrintDialog*>( (QPrintDialog*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QPrintDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPrintDialog*)(self) )->virtualbase_MetaObject();
}

void QPrintDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrintDialog*>( (QPrintDialog*)(self) )->handle__Metacast = slot;
}

void* QPrintDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPrintDialog*)(self) )->virtualbase_Metacast(param1);
}

void QPrintDialog_Delete(QPrintDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPrintDialog*>( self );
	} else {
		delete self;
	}
}

