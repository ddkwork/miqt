// +build ignore

#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDialog>
#include <QEvent>
#include <QKeyEvent>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QResizeEvent>
#include <QShowEvent>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qdialog.h>
#include "gen_qdialog.h"
class MiqtVirtualQDialog : public virtual QDialog {
public:
MiqtVirtualQDialog(QWidget* parent): QDialog(parent) {};
MiqtVirtualQDialog(): QDialog() {};
MiqtVirtualQDialog(QWidget* parent, Qt::WindowFlags f): QDialog(parent, f) {};

virtual ~MiqtVirtualQDialog() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QDialog::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QDialog_MetaObject(const_cast<MiqtVirtualQDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QDialog::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QDialog::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QDialog::qt_metacast(param1);

	}
};
QDialog* QDialog_new(QWidget* parent) {
return new MiqtVirtualQDialog(parent);
}
QDialog* QDialog_new2() {
return new MiqtVirtualQDialog();
}
QDialog* QDialog_new3(QWidget* parent, int f) {
return new MiqtVirtualQDialog(parent, static_cast<Qt::WindowFlags>(f));
}
void QDialog_virtbase(QDialog* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QDialog_MetaObject(const QDialog* self) {
	return (QMetaObject*) self->metaObject();
}
void* QDialog_Metacast(QDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QDialog_Tr(const char* s) {
	QString _ret = QDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
int QDialog_Result(const QDialog* self) {
	return self->result();
}
void QDialog_SetVisible(QDialog* self, bool visible) {
	self->setVisible(visible);
}
QSize* QDialog_SizeHint(const QDialog* self) {
	return new QSize(self->sizeHint());
}
QSize* QDialog_MinimumSizeHint(const QDialog* self) {
	return new QSize(self->minimumSizeHint());
}
void QDialog_SetSizeGripEnabled(QDialog* self, bool sizeGripEnabled) {
	self->setSizeGripEnabled(sizeGripEnabled);
}
bool QDialog_IsSizeGripEnabled(const QDialog* self) {
	return self->isSizeGripEnabled();
}
void QDialog_SetModal(QDialog* self, bool modal) {
	self->setModal(modal);
}
void QDialog_SetResult(QDialog* self, int r) {
	self->setResult(static_cast<int>(r));
}
void QDialog_Finished(QDialog* self, int result) {
	self->finished(static_cast<int>(result));
}
void QDialog_connect_Finished(QDialog* self, intptr_t slot) {
	MiqtVirtualQDialog::connect(self, static_cast<void (QDialog::*)(int)>(&QDialog::finished), self, [=](int result) {
		int sigval1 = result;
		miqt_exec_callback_QDialog_Finished(slot, sigval1);
	});
}
void QDialog_Accepted(QDialog* self) {
	self->accepted();
}
void QDialog_connect_Accepted(QDialog* self, intptr_t slot) {
	MiqtVirtualQDialog::connect(self, static_cast<void (QDialog::*)()>(&QDialog::accepted), self, [=]() {
		miqt_exec_callback_QDialog_Accepted(slot);
	});
}
void QDialog_Rejected(QDialog* self) {
	self->rejected();
}
void QDialog_connect_Rejected(QDialog* self, intptr_t slot) {
	MiqtVirtualQDialog::connect(self, static_cast<void (QDialog::*)()>(&QDialog::rejected), self, [=]() {
		miqt_exec_callback_QDialog_Rejected(slot);
	});
}
void QDialog_Open(QDialog* self) {
	self->open();
}
int QDialog_Exec(QDialog* self) {
	return self->exec();
}
void QDialog_Done(QDialog* self, int param1) {
	self->done(static_cast<int>(param1));
}
void QDialog_Accept(QDialog* self) {
	self->accept();
}
void QDialog_Reject(QDialog* self) {
	self->reject();
}
struct miqt_string QDialog_Tr2(const char* s, const char* c) {
	QString _ret = QDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDialog*>( (QDialog*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQDialog*)(self) )->virtualbase_MetaObject();
}
void QDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDialog*>( (QDialog*)(self) )->handle__Metacast = slot;
}
void* QDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQDialog*)(self) )->virtualbase_Metacast(param1);
}
void QDialog_Delete(QDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDialog*>( self );
	} else {
		delete self;
	}
}
