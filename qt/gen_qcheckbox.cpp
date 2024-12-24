// +build ignore

#include <QAbstractButton>
#include <QCheckBox>
#include <QEvent>
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
#include <QStyleOptionButton>
#include <QWidget>
#include <qcheckbox.h>
#include "gen_qcheckbox.h"
class MiqtVirtualQCheckBox : public virtual QCheckBox {
public:
MiqtVirtualQCheckBox(QWidget* parent): QCheckBox(parent) {};
MiqtVirtualQCheckBox(): QCheckBox() {};
MiqtVirtualQCheckBox(const QString& text): QCheckBox(text) {};
MiqtVirtualQCheckBox(const QString& text, QWidget* parent): QCheckBox(text, parent) {};

virtual ~MiqtVirtualQCheckBox() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QCheckBox::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QCheckBox_MetaObject(const_cast<MiqtVirtualQCheckBox*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QCheckBox::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QCheckBox::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QCheckBox_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QCheckBox::qt_metacast(param1);

	}
};
QCheckBox* QCheckBox_new(QWidget* parent) {
return new MiqtVirtualQCheckBox(parent);
}
QCheckBox* QCheckBox_new2() {
return new MiqtVirtualQCheckBox();
}
QCheckBox* QCheckBox_new3(struct miqt_string text) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQCheckBox(text_QString);
}
QCheckBox* QCheckBox_new4(struct miqt_string text, QWidget* parent) {
QString text_QString = QString::fromUtf8(text.data, text.len);
	return new MiqtVirtualQCheckBox(text_QString, parent);
}
void QCheckBox_virtbase(QCheckBox* src
, QAbstractButton** outptr_QAbstractButton
) {
*outptr_QAbstractButton = static_cast<QAbstractButton*>(src);
}
QMetaObject* QCheckBox_MetaObject(const QCheckBox* self) {
	return (QMetaObject*) self->metaObject();
}
void* QCheckBox_Metacast(QCheckBox* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QCheckBox_Tr(const char* s) {
	QString _ret = QCheckBox::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
QSize* QCheckBox_SizeHint(const QCheckBox* self) {
	return new QSize(self->sizeHint());
}
QSize* QCheckBox_MinimumSizeHint(const QCheckBox* self) {
	return new QSize(self->minimumSizeHint());
}
void QCheckBox_SetTristate(QCheckBox* self) {
	self->setTristate();
}
bool QCheckBox_IsTristate(const QCheckBox* self) {
	return self->isTristate();
}
int QCheckBox_CheckState(const QCheckBox* self) {
	Qt::CheckState _ret = self->checkState();
	return static_cast<int>(_ret);
}
void QCheckBox_SetCheckState(QCheckBox* self, int state) {
	self->setCheckState(static_cast<Qt::CheckState>(state));
}
void QCheckBox_StateChanged(QCheckBox* self, int param1) {
	self->stateChanged(static_cast<int>(param1));
}
void QCheckBox_connect_StateChanged(QCheckBox* self, intptr_t slot) {
	MiqtVirtualQCheckBox::connect(self, static_cast<void (QCheckBox::*)(int)>(&QCheckBox::stateChanged), self, [=](int param1) {
		int sigval1 = param1;
		miqt_exec_callback_QCheckBox_StateChanged(slot, sigval1);
	});
}
void QCheckBox_CheckStateChanged(QCheckBox* self, int param1) {
	self->checkStateChanged(static_cast<Qt::CheckState>(param1));
}
void QCheckBox_connect_CheckStateChanged(QCheckBox* self, intptr_t slot) {
	MiqtVirtualQCheckBox::connect(self, static_cast<void (QCheckBox::*)(Qt::CheckState)>(&QCheckBox::checkStateChanged), self, [=](Qt::CheckState param1) {
		Qt::CheckState param1_ret = param1;
		int sigval1 = static_cast<int>(param1_ret);
		miqt_exec_callback_QCheckBox_CheckStateChanged(slot, sigval1);
	});
}
struct miqt_string QCheckBox_Tr2(const char* s, const char* c) {
	QString _ret = QCheckBox::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QCheckBox_Tr3(const char* s, const char* c, int n) {
	QString _ret = QCheckBox::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QCheckBox_SetTristate1(QCheckBox* self, bool y) {
	self->setTristate(y);
}
void QCheckBox_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCheckBox*>( (QCheckBox*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QCheckBox_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQCheckBox*)(self) )->virtualbase_MetaObject();
}
void QCheckBox_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQCheckBox*>( (QCheckBox*)(self) )->handle__Metacast = slot;
}
void* QCheckBox_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQCheckBox*)(self) )->virtualbase_Metacast(param1);
}
void QCheckBox_Delete(QCheckBox* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQCheckBox*>( self );
	} else {
		delete self;
	}
}
