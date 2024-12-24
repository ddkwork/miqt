// +build ignore

#include <QAbstractPrintDialog>
#include <QDialog>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPrinter>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qabstractprintdialog.h>
#include "gen_qabstractprintdialog.h"

class MiqtVirtualQAbstractPrintDialog : public virtual QAbstractPrintDialog {
public:

	MiqtVirtualQAbstractPrintDialog(QPrinter* printer): QAbstractPrintDialog(printer) {};
	MiqtVirtualQAbstractPrintDialog(QPrinter* printer, QWidget* parent): QAbstractPrintDialog(printer, parent) {};

	virtual ~MiqtVirtualQAbstractPrintDialog() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAbstractPrintDialog::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAbstractPrintDialog_MetaObject(const_cast<MiqtVirtualQAbstractPrintDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAbstractPrintDialog::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAbstractPrintDialog::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAbstractPrintDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAbstractPrintDialog::qt_metacast(param1);

	}

};

QAbstractPrintDialog* QAbstractPrintDialog_new(QPrinter* printer) {
	return new MiqtVirtualQAbstractPrintDialog(printer);
}

QAbstractPrintDialog* QAbstractPrintDialog_new2(QPrinter* printer, QWidget* parent) {
	return new MiqtVirtualQAbstractPrintDialog(printer, parent);
}

void QAbstractPrintDialog_virtbase(QAbstractPrintDialog* src, QDialog** outptr_QDialog) {
	*outptr_QDialog = static_cast<QDialog*>(src);
}

QMetaObject* QAbstractPrintDialog_MetaObject(const QAbstractPrintDialog* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractPrintDialog_Metacast(QAbstractPrintDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractPrintDialog_Tr(const char* s) {
	QString _ret = QAbstractPrintDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractPrintDialog_SetOptionTabs(QAbstractPrintDialog* self, struct miqt_array /* of QWidget* */  tabs) {
	QList<QWidget *> tabs_QList;
	tabs_QList.reserve(tabs.len);
	QWidget** tabs_arr = static_cast<QWidget**>(tabs.data);
	for(size_t i = 0; i < tabs.len; ++i) {
		tabs_QList.push_back(tabs_arr[i]);
	}
	self->setOptionTabs(tabs_QList);
}

void QAbstractPrintDialog_SetPrintRange(QAbstractPrintDialog* self, PrintRange rangeVal) {
	self->setPrintRange(rangeVal);
}

PrintRange QAbstractPrintDialog_PrintRange(const QAbstractPrintDialog* self) {
	return self->printRange();
}

void QAbstractPrintDialog_SetMinMax(QAbstractPrintDialog* self, int min, int max) {
	self->setMinMax(static_cast<int>(min), static_cast<int>(max));
}

int QAbstractPrintDialog_MinPage(const QAbstractPrintDialog* self) {
	return self->minPage();
}

int QAbstractPrintDialog_MaxPage(const QAbstractPrintDialog* self) {
	return self->maxPage();
}

void QAbstractPrintDialog_SetFromTo(QAbstractPrintDialog* self, int fromPage, int toPage) {
	self->setFromTo(static_cast<int>(fromPage), static_cast<int>(toPage));
}

int QAbstractPrintDialog_FromPage(const QAbstractPrintDialog* self) {
	return self->fromPage();
}

int QAbstractPrintDialog_ToPage(const QAbstractPrintDialog* self) {
	return self->toPage();
}

QPrinter* QAbstractPrintDialog_Printer(const QAbstractPrintDialog* self) {
	return self->printer();
}

struct miqt_string QAbstractPrintDialog_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractPrintDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractPrintDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractPrintDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractPrintDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractPrintDialog*>( (QAbstractPrintDialog*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAbstractPrintDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAbstractPrintDialog*)(self) )->virtualbase_MetaObject();
}

void QAbstractPrintDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractPrintDialog*>( (QAbstractPrintDialog*)(self) )->handle__Metacast = slot;
}

void* QAbstractPrintDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAbstractPrintDialog*)(self) )->virtualbase_Metacast(param1);
}

void QAbstractPrintDialog_Delete(QAbstractPrintDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractPrintDialog*>( self );
	} else {
		delete self;
	}
}

