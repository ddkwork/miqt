// +build ignore

#include <QDialog>
#include <QEvent>
#include <QFont>
#include <QFontDialog>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qfontdialog.h>
#include "gen_qfontdialog.h"

class MiqtVirtualQFontDialog : public virtual QFontDialog {
public:

	MiqtVirtualQFontDialog(QWidget* parent): QFontDialog(parent) {};
	MiqtVirtualQFontDialog(): QFontDialog() {};
	MiqtVirtualQFontDialog(const QFont& initial): QFontDialog(initial) {};
	MiqtVirtualQFontDialog(const QFont& initial, QWidget* parent): QFontDialog(initial, parent) {};

	virtual ~MiqtVirtualQFontDialog() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QFontDialog::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QFontDialog_MetaObject(const_cast<MiqtVirtualQFontDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QFontDialog::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QFontDialog::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QFontDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QFontDialog::qt_metacast(param1);

	}

};

QFontDialog* QFontDialog_new(QWidget* parent) {
	return new MiqtVirtualQFontDialog(parent);
}

QFontDialog* QFontDialog_new2() {
	return new MiqtVirtualQFontDialog();
}

QFontDialog* QFontDialog_new3(QFont* initial) {
	return new MiqtVirtualQFontDialog(*initial);
}

QFontDialog* QFontDialog_new4(QFont* initial, QWidget* parent) {
	return new MiqtVirtualQFontDialog(*initial, parent);
}

void QFontDialog_virtbase(QFontDialog* src, QDialog** outptr_QDialog) {
	*outptr_QDialog = static_cast<QDialog*>(src);
}

QMetaObject* QFontDialog_MetaObject(const QFontDialog* self) {
	return (QMetaObject*) self->metaObject();
}

void* QFontDialog_Metacast(QFontDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QFontDialog_Tr(const char* s) {
	QString _ret = QFontDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFontDialog_SetCurrentFont(QFontDialog* self, QFont* font) {
	self->setCurrentFont(*font);
}

QFont* QFontDialog_CurrentFont(const QFontDialog* self) {
	return new QFont(self->currentFont());
}

QFont* QFontDialog_SelectedFont(const QFontDialog* self) {
	return new QFont(self->selectedFont());
}

void QFontDialog_SetOption(QFontDialog* self, FontDialogOption option) {
	self->setOption(option);
}

bool QFontDialog_TestOption(const QFontDialog* self, FontDialogOption option) {
	return self->testOption(option);
}

void QFontDialog_SetOptions(QFontDialog* self, FontDialogOptions options) {
	self->setOptions(options);
}

FontDialogOptions QFontDialog_Options(const QFontDialog* self) {
	return self->options();
}

void QFontDialog_SetVisible(QFontDialog* self, bool visible) {
	self->setVisible(visible);
}

QFont* QFontDialog_GetFont(bool* ok) {
	return new QFont(QFontDialog::getFont(ok));
}

QFont* QFontDialog_GetFont2(bool* ok, QFont* initial) {
	return new QFont(QFontDialog::getFont(ok, *initial));
}

void QFontDialog_CurrentFontChanged(QFontDialog* self, QFont* font) {
	self->currentFontChanged(*font);
}

void QFontDialog_connect_CurrentFontChanged(QFontDialog* self, intptr_t slot) {
	MiqtVirtualQFontDialog::connect(self, static_cast<void (QFontDialog::*)(const QFont&)>(&QFontDialog::currentFontChanged), self, [=](const QFont& font) {
		const QFont& font_ret = font;
		// Cast returned reference into pointer
		QFont* sigval1 = const_cast<QFont*>(&font_ret);
		miqt_exec_callback_QFontDialog_CurrentFontChanged(slot, sigval1);
	});
}

void QFontDialog_FontSelected(QFontDialog* self, QFont* font) {
	self->fontSelected(*font);
}

void QFontDialog_connect_FontSelected(QFontDialog* self, intptr_t slot) {
	MiqtVirtualQFontDialog::connect(self, static_cast<void (QFontDialog::*)(const QFont&)>(&QFontDialog::fontSelected), self, [=](const QFont& font) {
		const QFont& font_ret = font;
		// Cast returned reference into pointer
		QFont* sigval1 = const_cast<QFont*>(&font_ret);
		miqt_exec_callback_QFontDialog_FontSelected(slot, sigval1);
	});
}

struct miqt_string QFontDialog_Tr2(const char* s, const char* c) {
	QString _ret = QFontDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QFontDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QFontDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QFontDialog_SetOption2(QFontDialog* self, FontDialogOption option, bool on) {
	self->setOption(option, on);
}

QFont* QFontDialog_GetFont22(bool* ok, QWidget* parent) {
	return new QFont(QFontDialog::getFont(ok, parent));
}

QFont* QFontDialog_GetFont3(bool* ok, QFont* initial, QWidget* parent) {
	return new QFont(QFontDialog::getFont(ok, *initial, parent));
}

QFont* QFontDialog_GetFont4(bool* ok, QFont* initial, QWidget* parent, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	return new QFont(QFontDialog::getFont(ok, *initial, parent, title_QString));
}

QFont* QFontDialog_GetFont5(bool* ok, QFont* initial, QWidget* parent, struct miqt_string title, FontDialogOptions options) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	return new QFont(QFontDialog::getFont(ok, *initial, parent, title_QString, options));
}

void QFontDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFontDialog*>( (QFontDialog*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QFontDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQFontDialog*)(self) )->virtualbase_MetaObject();
}

void QFontDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQFontDialog*>( (QFontDialog*)(self) )->handle__Metacast = slot;
}

void* QFontDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQFontDialog*)(self) )->virtualbase_Metacast(param1);
}

void QFontDialog_Delete(QFontDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQFontDialog*>( self );
	} else {
		delete self;
	}
}

