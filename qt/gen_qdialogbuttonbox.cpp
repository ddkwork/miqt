// +build ignore

#include <QAbstractButton>
#include <QDialogButtonBox>
#include <QEvent>
#include <QList>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPushButton>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qdialogbuttonbox.h>
#include "gen_qdialogbuttonbox.h"

class MiqtVirtualQDialogButtonBox : public virtual QDialogButtonBox {
public:

	MiqtVirtualQDialogButtonBox(QWidget* parent): QDialogButtonBox(parent) {};
	MiqtVirtualQDialogButtonBox(): QDialogButtonBox() {};
	MiqtVirtualQDialogButtonBox(Qt::Orientation orientation): QDialogButtonBox(orientation) {};
	MiqtVirtualQDialogButtonBox(StandardButtons buttons): QDialogButtonBox(buttons) {};
	MiqtVirtualQDialogButtonBox(StandardButtons buttons, Qt::Orientation orientation): QDialogButtonBox(buttons, orientation) {};
	MiqtVirtualQDialogButtonBox(Qt::Orientation orientation, QWidget* parent): QDialogButtonBox(orientation, parent) {};
	MiqtVirtualQDialogButtonBox(StandardButtons buttons, QWidget* parent): QDialogButtonBox(buttons, parent) {};
	MiqtVirtualQDialogButtonBox(StandardButtons buttons, Qt::Orientation orientation, QWidget* parent): QDialogButtonBox(buttons, orientation, parent) {};

	virtual ~MiqtVirtualQDialogButtonBox() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QDialogButtonBox::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QDialogButtonBox_MetaObject(const_cast<MiqtVirtualQDialogButtonBox*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QDialogButtonBox::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QDialogButtonBox::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QDialogButtonBox_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QDialogButtonBox::qt_metacast(param1);

	}

};

QDialogButtonBox* QDialogButtonBox_new(QWidget* parent) {
	return new MiqtVirtualQDialogButtonBox(parent);
}

QDialogButtonBox* QDialogButtonBox_new2() {
	return new MiqtVirtualQDialogButtonBox();
}

QDialogButtonBox* QDialogButtonBox_new3(int orientation) {
	return new MiqtVirtualQDialogButtonBox(static_cast<Qt::Orientation>(orientation));
}

QDialogButtonBox* QDialogButtonBox_new4(StandardButtons buttons) {
	return new MiqtVirtualQDialogButtonBox(buttons);
}

QDialogButtonBox* QDialogButtonBox_new5(StandardButtons buttons, int orientation) {
	return new MiqtVirtualQDialogButtonBox(buttons, static_cast<Qt::Orientation>(orientation));
}

QDialogButtonBox* QDialogButtonBox_new6(int orientation, QWidget* parent) {
	return new MiqtVirtualQDialogButtonBox(static_cast<Qt::Orientation>(orientation), parent);
}

QDialogButtonBox* QDialogButtonBox_new7(StandardButtons buttons, QWidget* parent) {
	return new MiqtVirtualQDialogButtonBox(buttons, parent);
}

QDialogButtonBox* QDialogButtonBox_new8(StandardButtons buttons, int orientation, QWidget* parent) {
	return new MiqtVirtualQDialogButtonBox(buttons, static_cast<Qt::Orientation>(orientation), parent);
}

void QDialogButtonBox_virtbase(QDialogButtonBox* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QDialogButtonBox_MetaObject(const QDialogButtonBox* self) {
	return (QMetaObject*) self->metaObject();
}

void* QDialogButtonBox_Metacast(QDialogButtonBox* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QDialogButtonBox_Tr(const char* s) {
	QString _ret = QDialogButtonBox::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDialogButtonBox_SetOrientation(QDialogButtonBox* self, int orientation) {
	self->setOrientation(static_cast<Qt::Orientation>(orientation));
}

int QDialogButtonBox_Orientation(const QDialogButtonBox* self) {
	Qt::Orientation _ret = self->orientation();
	return static_cast<int>(_ret);
}

void QDialogButtonBox_AddButton(QDialogButtonBox* self, QAbstractButton* button, ButtonRole role) {
	self->addButton(button, role);
}

QPushButton* QDialogButtonBox_AddButton2(QDialogButtonBox* self, struct miqt_string text, ButtonRole role) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	return self->addButton(text_QString, role);
}

QPushButton* QDialogButtonBox_AddButtonWithButton(QDialogButtonBox* self, StandardButton button) {
	return self->addButton(button);
}

void QDialogButtonBox_RemoveButton(QDialogButtonBox* self, QAbstractButton* button) {
	self->removeButton(button);
}

void QDialogButtonBox_Clear(QDialogButtonBox* self) {
	self->clear();
}

struct miqt_array /* of QAbstractButton* */  QDialogButtonBox_Buttons(const QDialogButtonBox* self) {
	QList<QAbstractButton *> _ret = self->buttons();
	// Convert QList<> from C++ memory to manually-managed C memory
	QAbstractButton** _arr = static_cast<QAbstractButton**>(malloc(sizeof(QAbstractButton*) * _ret.length()));
	for (size_t i = 0, e = _ret.length(); i < e; ++i) {
		_arr[i] = _ret[i];
	}
	struct miqt_array _out;
	_out.len = _ret.length();
	_out.data = static_cast<void*>(_arr);
	return _out;
}

ButtonRole QDialogButtonBox_ButtonRole(const QDialogButtonBox* self, QAbstractButton* button) {
	return self->buttonRole(button);
}

void QDialogButtonBox_SetStandardButtons(QDialogButtonBox* self, StandardButtons buttons) {
	self->setStandardButtons(buttons);
}

StandardButtons QDialogButtonBox_StandardButtons(const QDialogButtonBox* self) {
	return self->standardButtons();
}

StandardButton QDialogButtonBox_StandardButton(const QDialogButtonBox* self, QAbstractButton* button) {
	return self->standardButton(button);
}

QPushButton* QDialogButtonBox_Button(const QDialogButtonBox* self, StandardButton which) {
	return self->button(which);
}

void QDialogButtonBox_SetCenterButtons(QDialogButtonBox* self, bool center) {
	self->setCenterButtons(center);
}

bool QDialogButtonBox_CenterButtons(const QDialogButtonBox* self) {
	return self->centerButtons();
}

void QDialogButtonBox_Clicked(QDialogButtonBox* self, QAbstractButton* button) {
	self->clicked(button);
}

void QDialogButtonBox_connect_Clicked(QDialogButtonBox* self, intptr_t slot) {
	MiqtVirtualQDialogButtonBox::connect(self, static_cast<void (QDialogButtonBox::*)(QAbstractButton*)>(&QDialogButtonBox::clicked), self, [=](QAbstractButton* button) {
		QAbstractButton* sigval1 = button;
		miqt_exec_callback_QDialogButtonBox_Clicked(slot, sigval1);
	});
}

void QDialogButtonBox_Accepted(QDialogButtonBox* self) {
	self->accepted();
}

void QDialogButtonBox_connect_Accepted(QDialogButtonBox* self, intptr_t slot) {
	MiqtVirtualQDialogButtonBox::connect(self, static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::accepted), self, [=]() {
		miqt_exec_callback_QDialogButtonBox_Accepted(slot);
	});
}

void QDialogButtonBox_HelpRequested(QDialogButtonBox* self) {
	self->helpRequested();
}

void QDialogButtonBox_connect_HelpRequested(QDialogButtonBox* self, intptr_t slot) {
	MiqtVirtualQDialogButtonBox::connect(self, static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::helpRequested), self, [=]() {
		miqt_exec_callback_QDialogButtonBox_HelpRequested(slot);
	});
}

void QDialogButtonBox_Rejected(QDialogButtonBox* self) {
	self->rejected();
}

void QDialogButtonBox_connect_Rejected(QDialogButtonBox* self, intptr_t slot) {
	MiqtVirtualQDialogButtonBox::connect(self, static_cast<void (QDialogButtonBox::*)()>(&QDialogButtonBox::rejected), self, [=]() {
		miqt_exec_callback_QDialogButtonBox_Rejected(slot);
	});
}

struct miqt_string QDialogButtonBox_Tr2(const char* s, const char* c) {
	QString _ret = QDialogButtonBox::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QDialogButtonBox_Tr3(const char* s, const char* c, int n) {
	QString _ret = QDialogButtonBox::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QDialogButtonBox_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDialogButtonBox*>( (QDialogButtonBox*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QDialogButtonBox_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQDialogButtonBox*)(self) )->virtualbase_MetaObject();
}

void QDialogButtonBox_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQDialogButtonBox*>( (QDialogButtonBox*)(self) )->handle__Metacast = slot;
}

void* QDialogButtonBox_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQDialogButtonBox*)(self) )->virtualbase_Metacast(param1);
}

void QDialogButtonBox_Delete(QDialogButtonBox* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQDialogButtonBox*>( self );
	} else {
		delete self;
	}
}

