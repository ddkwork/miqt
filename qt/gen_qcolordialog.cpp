// +build ignore

#include <QColor>
#include <QColorDialog>
#include <QDialog>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qcolordialog.h>
#include "gen_qcolordialog.h"

class MiqtVirtualQColorDialog : public virtual QColorDialog {
public:

	MiqtVirtualQColorDialog(QWidget* parent): QColorDialog(parent) {};
	MiqtVirtualQColorDialog(): QColorDialog() {};
	MiqtVirtualQColorDialog(const QColor& initial): QColorDialog(initial) {};
	MiqtVirtualQColorDialog(const QColor& initial, QWidget* parent): QColorDialog(initial, parent) {};

	virtual ~MiqtVirtualQColorDialog() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QColorDialog::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QColorDialog_MetaObject(const_cast<MiqtVirtualQColorDialog*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QColorDialog::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QColorDialog::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QColorDialog_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QColorDialog::qt_metacast(param1);

	}

};

QColorDialog* QColorDialog_new(QWidget* parent) {
	return new MiqtVirtualQColorDialog(parent);
}

QColorDialog* QColorDialog_new2() {
	return new MiqtVirtualQColorDialog();
}

QColorDialog* QColorDialog_new3(QColor* initial) {
	return new MiqtVirtualQColorDialog(*initial);
}

QColorDialog* QColorDialog_new4(QColor* initial, QWidget* parent) {
	return new MiqtVirtualQColorDialog(*initial, parent);
}

void QColorDialog_virtbase(QColorDialog* src, QDialog** outptr_QDialog) {
	*outptr_QDialog = static_cast<QDialog*>(src);
}

QMetaObject* QColorDialog_MetaObject(const QColorDialog* self) {
	return (QMetaObject*) self->metaObject();
}

void* QColorDialog_Metacast(QColorDialog* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QColorDialog_Tr(const char* s) {
	QString _ret = QColorDialog::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QColorDialog_SetCurrentColor(QColorDialog* self, QColor* color) {
	self->setCurrentColor(*color);
}

QColor* QColorDialog_CurrentColor(const QColorDialog* self) {
	return new QColor(self->currentColor());
}

QColor* QColorDialog_SelectedColor(const QColorDialog* self) {
	return new QColor(self->selectedColor());
}

void QColorDialog_SetOption(QColorDialog* self, ColorDialogOption option) {
	self->setOption(option);
}

bool QColorDialog_TestOption(const QColorDialog* self, ColorDialogOption option) {
	return self->testOption(option);
}

void QColorDialog_SetOptions(QColorDialog* self, ColorDialogOptions options) {
	self->setOptions(options);
}

ColorDialogOptions QColorDialog_Options(const QColorDialog* self) {
	return self->options();
}

void QColorDialog_SetVisible(QColorDialog* self, bool visible) {
	self->setVisible(visible);
}

QColor* QColorDialog_GetColor() {
	return new QColor(QColorDialog::getColor());
}

int QColorDialog_CustomCount() {
	return QColorDialog::customCount();
}

QColor* QColorDialog_CustomColor(int index) {
	return new QColor(QColorDialog::customColor(static_cast<int>(index)));
}

void QColorDialog_SetCustomColor(int index, QColor* color) {
	QColorDialog::setCustomColor(static_cast<int>(index), *color);
}

QColor* QColorDialog_StandardColor(int index) {
	return new QColor(QColorDialog::standardColor(static_cast<int>(index)));
}

void QColorDialog_SetStandardColor(int index, QColor* color) {
	QColorDialog::setStandardColor(static_cast<int>(index), *color);
}

void QColorDialog_CurrentColorChanged(QColorDialog* self, QColor* color) {
	self->currentColorChanged(*color);
}

void QColorDialog_connect_CurrentColorChanged(QColorDialog* self, intptr_t slot) {
	MiqtVirtualQColorDialog::connect(self, static_cast<void (QColorDialog::*)(const QColor&)>(&QColorDialog::currentColorChanged), self, [=](const QColor& color) {
		const QColor& color_ret = color;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&color_ret);
		miqt_exec_callback_QColorDialog_CurrentColorChanged(slot, sigval1);
	});
}

void QColorDialog_ColorSelected(QColorDialog* self, QColor* color) {
	self->colorSelected(*color);
}

void QColorDialog_connect_ColorSelected(QColorDialog* self, intptr_t slot) {
	MiqtVirtualQColorDialog::connect(self, static_cast<void (QColorDialog::*)(const QColor&)>(&QColorDialog::colorSelected), self, [=](const QColor& color) {
		const QColor& color_ret = color;
		// Cast returned reference into pointer
		QColor* sigval1 = const_cast<QColor*>(&color_ret);
		miqt_exec_callback_QColorDialog_ColorSelected(slot, sigval1);
	});
}

struct miqt_string QColorDialog_Tr2(const char* s, const char* c) {
	QString _ret = QColorDialog::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QColorDialog_Tr3(const char* s, const char* c, int n) {
	QString _ret = QColorDialog::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QColorDialog_SetOption2(QColorDialog* self, ColorDialogOption option, bool on) {
	self->setOption(option, on);
}

QColor* QColorDialog_GetColor1(QColor* initial) {
	return new QColor(QColorDialog::getColor(*initial));
}

QColor* QColorDialog_GetColor2(QColor* initial, QWidget* parent) {
	return new QColor(QColorDialog::getColor(*initial, parent));
}

QColor* QColorDialog_GetColor3(QColor* initial, QWidget* parent, struct miqt_string title) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	return new QColor(QColorDialog::getColor(*initial, parent, title_QString));
}

QColor* QColorDialog_GetColor4(QColor* initial, QWidget* parent, struct miqt_string title, ColorDialogOptions options) {
	QString title_QString = QString::fromUtf8(title.data, title.len);
	return new QColor(QColorDialog::getColor(*initial, parent, title_QString, options));
}

void QColorDialog_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQColorDialog*>( (QColorDialog*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QColorDialog_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQColorDialog*)(self) )->virtualbase_MetaObject();
}

void QColorDialog_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQColorDialog*>( (QColorDialog*)(self) )->handle__Metacast = slot;
}

void* QColorDialog_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQColorDialog*)(self) )->virtualbase_Metacast(param1);
}

void QColorDialog_Delete(QColorDialog* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQColorDialog*>( self );
	} else {
		delete self;
	}
}

