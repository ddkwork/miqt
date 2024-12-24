// +build ignore

#include <QAbstractButton>
#include <QButtonGroup>
#include <QEvent>
#include <QFocusEvent>
#include <QIcon>
#include <QKeyEvent>
#include <QKeySequence>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QPoint>
#include <QSize>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QTimerEvent>
#include <QWidget>
#include <qabstractbutton.h>
#include "gen_qabstractbutton.h"

class MiqtVirtualQAbstractButton : public virtual QAbstractButton {
public:

	MiqtVirtualQAbstractButton(QWidget* parent): QAbstractButton(parent) {};
	MiqtVirtualQAbstractButton(): QAbstractButton() {};

	virtual ~MiqtVirtualQAbstractButton() = default;

	// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;

	// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
		if (handle__MetaObject == 0) {
			return QAbstractButton::metaObject();
		}
		

		QMetaObject* callback_return_value = miqt_exec_callback_QAbstractButton_MetaObject(const_cast<MiqtVirtualQAbstractButton*>(this), handle__MetaObject);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QAbstractButton::metaObject();

	}

	// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;

	// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
		if (handle__Metacast == 0) {
			return QAbstractButton::qt_metacast(param1);
		}
		
		const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QAbstractButton_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}

	// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QAbstractButton::qt_metacast(param1);

	}

};

QAbstractButton* QAbstractButton_new(QWidget* parent) {
	return new MiqtVirtualQAbstractButton(parent);
}

QAbstractButton* QAbstractButton_new2() {
	return new MiqtVirtualQAbstractButton();
}

void QAbstractButton_virtbase(QAbstractButton* src, QWidget** outptr_QWidget) {
	*outptr_QWidget = static_cast<QWidget*>(src);
}

QMetaObject* QAbstractButton_MetaObject(const QAbstractButton* self) {
	return (QMetaObject*) self->metaObject();
}

void* QAbstractButton_Metacast(QAbstractButton* self, const char* param1) {
	return self->qt_metacast(param1);
}

struct miqt_string QAbstractButton_Tr(const char* s) {
	QString _ret = QAbstractButton::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractButton_SetText(QAbstractButton* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->setText(text_QString);
}

struct miqt_string QAbstractButton_Text(const QAbstractButton* self) {
	QString _ret = self->text();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractButton_SetIcon(QAbstractButton* self, QIcon* icon) {
	self->setIcon(*icon);
}

QIcon* QAbstractButton_Icon(const QAbstractButton* self) {
	return new QIcon(self->icon());
}

QSize* QAbstractButton_IconSize(const QAbstractButton* self) {
	return new QSize(self->iconSize());
}

void QAbstractButton_SetShortcut(QAbstractButton* self, QKeySequence* key) {
	self->setShortcut(*key);
}

QKeySequence* QAbstractButton_Shortcut(const QAbstractButton* self) {
	return new QKeySequence(self->shortcut());
}

void QAbstractButton_SetCheckable(QAbstractButton* self, bool checkable) {
	self->setCheckable(checkable);
}

bool QAbstractButton_IsCheckable(const QAbstractButton* self) {
	return self->isCheckable();
}

bool QAbstractButton_IsChecked(const QAbstractButton* self) {
	return self->isChecked();
}

void QAbstractButton_SetDown(QAbstractButton* self, bool down) {
	self->setDown(down);
}

bool QAbstractButton_IsDown(const QAbstractButton* self) {
	return self->isDown();
}

void QAbstractButton_SetAutoRepeat(QAbstractButton* self, bool autoRepeat) {
	self->setAutoRepeat(autoRepeat);
}

bool QAbstractButton_AutoRepeat(const QAbstractButton* self) {
	return self->autoRepeat();
}

void QAbstractButton_SetAutoRepeatDelay(QAbstractButton* self, int autoRepeatDelay) {
	self->setAutoRepeatDelay(static_cast<int>(autoRepeatDelay));
}

int QAbstractButton_AutoRepeatDelay(const QAbstractButton* self) {
	return self->autoRepeatDelay();
}

void QAbstractButton_SetAutoRepeatInterval(QAbstractButton* self, int autoRepeatInterval) {
	self->setAutoRepeatInterval(static_cast<int>(autoRepeatInterval));
}

int QAbstractButton_AutoRepeatInterval(const QAbstractButton* self) {
	return self->autoRepeatInterval();
}

void QAbstractButton_SetAutoExclusive(QAbstractButton* self, bool autoExclusive) {
	self->setAutoExclusive(autoExclusive);
}

bool QAbstractButton_AutoExclusive(const QAbstractButton* self) {
	return self->autoExclusive();
}

QButtonGroup* QAbstractButton_Group(const QAbstractButton* self) {
	return self->group();
}

void QAbstractButton_SetIconSize(QAbstractButton* self, QSize* size) {
	self->setIconSize(*size);
}

void QAbstractButton_AnimateClick(QAbstractButton* self) {
	self->animateClick();
}

void QAbstractButton_Click(QAbstractButton* self) {
	self->click();
}

void QAbstractButton_Toggle(QAbstractButton* self) {
	self->toggle();
}

void QAbstractButton_SetChecked(QAbstractButton* self, bool checked) {
	self->setChecked(checked);
}

void QAbstractButton_Pressed(QAbstractButton* self) {
	self->pressed();
}

void QAbstractButton_connect_Pressed(QAbstractButton* self, intptr_t slot) {
	MiqtVirtualQAbstractButton::connect(self, static_cast<void (QAbstractButton::*)()>(&QAbstractButton::pressed), self, [=]() {
		miqt_exec_callback_QAbstractButton_Pressed(slot);
	});
}

void QAbstractButton_Released(QAbstractButton* self) {
	self->released();
}

void QAbstractButton_connect_Released(QAbstractButton* self, intptr_t slot) {
	MiqtVirtualQAbstractButton::connect(self, static_cast<void (QAbstractButton::*)()>(&QAbstractButton::released), self, [=]() {
		miqt_exec_callback_QAbstractButton_Released(slot);
	});
}

void QAbstractButton_Clicked(QAbstractButton* self) {
	self->clicked();
}

void QAbstractButton_connect_Clicked(QAbstractButton* self, intptr_t slot) {
	MiqtVirtualQAbstractButton::connect(self, static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), self, [=]() {
		miqt_exec_callback_QAbstractButton_Clicked(slot);
	});
}

void QAbstractButton_Toggled(QAbstractButton* self, bool checked) {
	self->toggled(checked);
}

void QAbstractButton_connect_Toggled(QAbstractButton* self, intptr_t slot) {
	MiqtVirtualQAbstractButton::connect(self, static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::toggled), self, [=](bool checked) {
		bool sigval1 = checked;
		miqt_exec_callback_QAbstractButton_Toggled(slot, sigval1);
	});
}

struct miqt_string QAbstractButton_Tr2(const char* s, const char* c) {
	QString _ret = QAbstractButton::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

struct miqt_string QAbstractButton_Tr3(const char* s, const char* c, int n) {
	QString _ret = QAbstractButton::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}

void QAbstractButton_Clicked1(QAbstractButton* self, bool checked) {
	self->clicked(checked);
}

void QAbstractButton_connect_Clicked1(QAbstractButton* self, intptr_t slot) {
	MiqtVirtualQAbstractButton::connect(self, static_cast<void (QAbstractButton::*)(bool)>(&QAbstractButton::clicked), self, [=](bool checked) {
		bool sigval1 = checked;
		miqt_exec_callback_QAbstractButton_Clicked1(slot, sigval1);
	});
}

void QAbstractButton_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractButton*>( (QAbstractButton*)(self) )->handle__MetaObject = slot;
}

QMetaObject* QAbstractButton_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQAbstractButton*)(self) )->virtualbase_MetaObject();
}

void QAbstractButton_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQAbstractButton*>( (QAbstractButton*)(self) )->handle__Metacast = slot;
}

void* QAbstractButton_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQAbstractButton*)(self) )->virtualbase_Metacast(param1);
}

void QAbstractButton_Delete(QAbstractButton* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQAbstractButton*>( self );
	} else {
		delete self;
	}
}

