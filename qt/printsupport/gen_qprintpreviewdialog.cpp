// +build ignore

#include <QDialog>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPrintPreviewDialog>
#include <QPrinter>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qprintpreviewdialog.h>
#include "gen_qprintpreviewdialog.h"

class MiqtVirtualQPrintPreviewDialog : public virtual QPrintPreviewDialog {
public:

	MiqtVirtualQPrintPreviewDialog(QWidget* parent): QPrintPreviewDialog(parent) {};
	MiqtVirtualQPrintPreviewDialog(): QPrintPreviewDialog() {};
	MiqtVirtualQPrintPreviewDialog(QPrinter* printer): QPrintPreviewDialog(printer) {};
	MiqtVirtualQPrintPreviewDialog(QWidget* parent, Qt::WindowFlags flags): QPrintPreviewDialog(parent, flags) {};
	MiqtVirtualQPrintPreviewDialog(QPrinter* printer, QWidget* parent): QPrintPreviewDialog(printer, parent) {};
	MiqtVirtualQPrintPreviewDialog(QPrinter* printer, QWidget* parent, Qt::WindowFlags flags): QPrintPreviewDialog(printer, parent, flags) {};

	virtual ~MiqtVirtualQPrintPreviewDialog() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QPrintPreviewDialog::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QPrintPreviewDialog_MetaObject(const_cast<MiqtVirtualQPrintPreviewDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QPrintPreviewDialog::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QPrintPreviewDialog::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QPrintPreviewDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QPrintPreviewDialog::qt_metacast(param1);

	}

};

QPrintPreviewDialog* QPrintPreviewDialog_new(QWidget* parent) {
	return new MiqtVirtualQPrintPreviewDialog(parent);
}

QPrintPreviewDialog* QPrintPreviewDialog_new2() {
	return new MiqtVirtualQPrintPreviewDialog();
}

QPrintPreviewDialog* QPrintPreviewDialog_new3(QPrinter* printer) {
	return new MiqtVirtualQPrintPreviewDialog(printer);
}

QPrintPreviewDialog* QPrintPreviewDialog_new4(QWidget* parent, int flags) {
	return new MiqtVirtualQPrintPreviewDialog(parent, static_cast<Qt::WindowFlags>(flags));
}

QPrintPreviewDialog* QPrintPreviewDialog_new5(QPrinter* printer, QWidget* parent) {
	return new MiqtVirtualQPrintPreviewDialog(printer, parent);
}

QPrintPreviewDialog* QPrintPreviewDialog_new6(QPrinter* printer, QWidget* parent, int flags) {
	return new MiqtVirtualQPrintPreviewDialog(printer, parent, static_cast<Qt::WindowFlags>(flags));
}

void QPrintPreviewDialog_virtbase(QPrintPreviewDialog* src, QDialog** outptr_QDialog) {
	*outptr_QDialog = static_cast<QDialog*>(src);
}

QMetaObject* QPrintPreviewDialog_MetaObject(const QPrintPreviewDialog* self) {
	return (QMetaObject*) self->metaObject();
}

void* QPrintPreviewDialog_Metacast(QPrintPreviewDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QPrintPreviewDialog_Tr(const char* s) {
	QString _ret = QPrintPreviewDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

QPrinter* QPrintPreviewDialog_Printer(QPrintPreviewDialog* self) {
	return self->printer();
}

void QPrintPreviewDialog_SetVisible(QPrintPreviewDialog* self, bool visible) {
	self->setVisible(visible);
}

void QPrintPreviewDialog_Done(QPrintPreviewDialog* self, int result) {
	self->done(static_cast<int>(result));
}

void QPrintPreviewDialog_PaintRequested(QPrintPreviewDialog* self, QPrinter* printer) {
	self->paintRequested(printer);
}

void QPrintPreviewDialog_connect_PaintRequested(QPrintPreviewDialog* self, intptr_t slot) {
	MiqtVirtualQPrintPreviewDialog::connect(self, static_cast<void (QPrintPreviewDialog::*)(QPrinter*)>(&QPrintPreviewDialog::paintRequested), self, [=](QPrinter* printer) {
		QPrinter* sigval1 = printer;
		miqt_exec_callback_QPrintPreviewDialog_PaintRequested(slot, sigval1);
	});
}

struct miqt_string QPrintPreviewDialog_Tr2(const char* s, const char* c) {
	QString _ret = QPrintPreviewDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QPrintPreviewDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QPrintPreviewDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QPrintPreviewDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrintPreviewDialog*>( (QPrintPreviewDialog*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QPrintPreviewDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQPrintPreviewDialog*)(self) )->virtualbase_MetaObject();
}

void QPrintPreviewDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQPrintPreviewDialog*>( (QPrintPreviewDialog*)(self) )->handle__Metacast = slot;
}

void* QPrintPreviewDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQPrintPreviewDialog*)(self) )->virtualbase_Metacast(param1);
}

void QPrintPreviewDialog_Delete(QPrintPreviewDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQPrintPreviewDialog*>( self );
	} else {
		delete self;
	}
}

