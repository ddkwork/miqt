// +build ignore

#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEvent>
#include <QResizeEvent>
#include <QShowEvent>
#include <QStatusBar>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qstatusbar.h>
#include "gen_qstatusbar.h"
class MiqtVirtualQStatusBar : public virtual QStatusBar {
public:
MiqtVirtualQStatusBar(QWidget* parent): QStatusBar(parent) {};
MiqtVirtualQStatusBar(): QStatusBar() {};

virtual ~MiqtVirtualQStatusBar() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QStatusBar::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QStatusBar_MetaObject(const_cast<MiqtVirtualQStatusBar*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QStatusBar::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QStatusBar::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QStatusBar_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QStatusBar::qt_metacast(param1);

	}
};
QStatusBar* QStatusBar_new(QWidget* parent) {
return new MiqtVirtualQStatusBar(parent);
}
QStatusBar* QStatusBar_new2() {
return new MiqtVirtualQStatusBar();
}
void QStatusBar_virtbase(QStatusBar* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QStatusBar_MetaObject(const QStatusBar* self) {
	return (QMetaObject*) self->metaObject();
}
void* QStatusBar_Metacast(QStatusBar* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QStatusBar_Tr(const char* s) {
	QString _ret = QStatusBar::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStatusBar_AddWidget(QStatusBar* self, QWidget* widget) {
	self->addWidget(widget);
}
int QStatusBar_InsertWidget(QStatusBar* self, int index, QWidget* widget) {
	return self->insertWidget(static_cast<int>(index), widget);
}
void QStatusBar_AddPermanentWidget(QStatusBar* self, QWidget* widget) {
	self->addPermanentWidget(widget);
}
int QStatusBar_InsertPermanentWidget(QStatusBar* self, int index, QWidget* widget) {
	return self->insertPermanentWidget(static_cast<int>(index), widget);
}
void QStatusBar_RemoveWidget(QStatusBar* self, QWidget* widget) {
	self->removeWidget(widget);
}
void QStatusBar_SetSizeGripEnabled(QStatusBar* self, bool sizeGripEnabled) {
	self->setSizeGripEnabled(sizeGripEnabled);
}
bool QStatusBar_IsSizeGripEnabled(const QStatusBar* self) {
	return self->isSizeGripEnabled();
}
struct miqt_string QStatusBar_CurrentMessage(const QStatusBar* self) {
	QString _ret = self->currentMessage();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStatusBar_ShowMessage(QStatusBar* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->showMessage(text_QString);
}
void QStatusBar_ClearMessage(QStatusBar* self) {
	self->clearMessage();
}
void QStatusBar_MessageChanged(QStatusBar* self, struct miqt_string text) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->messageChanged(text_QString);
}
void QStatusBar_connect_MessageChanged(QStatusBar* self, intptr_t slot) {
	MiqtVirtualQStatusBar::connect(self, static_cast<void (QStatusBar::*)(const QString&)>(&QStatusBar::messageChanged), self, [=](const QString& text) {
		const QString text_ret = text;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray text_b = text_ret.toUtf8();
		struct miqt_string text_ms;
		text_ms.len = text_b.length();
		text_ms.data = static_cast<char*>(malloc(text_ms.len));
		memcpy(text_ms.data, text_b.data(), text_ms.len);
		struct miqt_string sigval1 = text_ms;
		miqt_exec_callback_QStatusBar_MessageChanged(slot, sigval1);
	});
}
struct miqt_string QStatusBar_Tr2(const char* s, const char* c) {
	QString _ret = QStatusBar::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QStatusBar_Tr3(const char* s, const char* c, int n) {
	QString _ret = QStatusBar::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QStatusBar_AddWidget2(QStatusBar* self, QWidget* widget, int stretch) {
	self->addWidget(widget, static_cast<int>(stretch));
}
int QStatusBar_InsertWidget3(QStatusBar* self, int index, QWidget* widget, int stretch) {
	return self->insertWidget(static_cast<int>(index), widget, static_cast<int>(stretch));
}
void QStatusBar_AddPermanentWidget2(QStatusBar* self, QWidget* widget, int stretch) {
	self->addPermanentWidget(widget, static_cast<int>(stretch));
}
int QStatusBar_InsertPermanentWidget3(QStatusBar* self, int index, QWidget* widget, int stretch) {
	return self->insertPermanentWidget(static_cast<int>(index), widget, static_cast<int>(stretch));
}
void QStatusBar_ShowMessage2(QStatusBar* self, struct miqt_string text, int timeout) {
	QString text_QString = QString::fromUtf8(text.data, text.len);
	self->showMessage(text_QString, static_cast<int>(timeout));
}
void QStatusBar_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStatusBar*>( (QStatusBar*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QStatusBar_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQStatusBar*)(self) )->virtualbase_MetaObject();
}
void QStatusBar_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQStatusBar*>( (QStatusBar*)(self) )->handle__Metacast = slot;
}
void* QStatusBar_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQStatusBar*)(self) )->virtualbase_Metacast(param1);
}
void QStatusBar_Delete(QStatusBar* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQStatusBar*>( self );
	} else {
		delete self;
	}
}
