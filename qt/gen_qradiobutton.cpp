// +build ignore

#include <QAbstractButton>
#include <QEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QRadioButton>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QStyleOptionButton>
#include <QWidget>
#include <qradiobutton.h>
#include "gen_qradiobutton.h"
class MiqtVirtualQRadioButton : public virtual QRadioButton {
public:
MiqtVirtualQRadioButton(QWidget* parent): QRadioButton(parent) {};
MiqtVirtualQRadioButton(): QRadioButton() {};
MiqtVirtualQRadioButton(const QString& text): QRadioButton(text) {};
MiqtVirtualQRadioButton(const QString& text, QWidget* parent): QRadioButton(text, parent) {};

virtual ~MiqtVirtualQRadioButton() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QRadioButton::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QRadioButton_MetaObject(const_cast<MiqtVirtualQRadioButton*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QRadioButton::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QRadioButton::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QRadioButton_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QRadioButton::qt_metacast(param1);

	}
};
QRadioButton* QRadioButton_new(QWidget* parent) {
return new MiqtVirtualQRadioButton(parent);
}
QRadioButton* QRadioButton_new2() {
return new MiqtVirtualQRadioButton();
}
QRadioButton* QRadioButton_new3(struct miqt_string text) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQRadioButton(text_QString);
}
QRadioButton* QRadioButton_new4(struct miqt_string text, QWidget* parent) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQRadioButton(text_QString, parent);
}
void QRadioButton_virtbase(QRadioButton* src
, QAbstractButton** outptr_QAbstractButton
) {
*outptr_QAbstractButton = static_cast<QAbstractButton*>(src);
}
QMetaObject* QRadioButton_MetaObject(const QRadioButton* self) {
	return (QMetaObject*) self->metaObject();
}
void* QRadioButton_Metacast(QRadioButton* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QRadioButton_Tr(const char* s) {
	QString _ret = QRadioButton::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QRadioButton_SizeHint(const QRadioButton* self) {
	return new QSize(self->sizeHint());
}
QSize* QRadioButton_MinimumSizeHint(const QRadioButton* self) {
	return new QSize(self->minimumSizeHint());
}
struct miqt_string QRadioButton_Tr2(const char* s, const char* c) {
	QString _ret = QRadioButton::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QRadioButton_Tr3(const char* s, const char* c, int n) {
	QString _ret = QRadioButton::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QRadioButton_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QRadioButton_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQRadioButton*)(self) )->virtualbase_MetaObject();
}
void QRadioButton_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQRadioButton*>( (QRadioButton*)(self) )->handle__Metacast = slot;
}
void* QRadioButton_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQRadioButton*)(self) )->virtualbase_Metacast(param1);
}
void QRadioButton_Delete(QRadioButton* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQRadioButton*>( self );
	} else {
		delete self;
	}
}
