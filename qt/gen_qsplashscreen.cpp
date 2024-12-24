// +build ignore

#include <QColor>
#include <QEvent>
#include <QMetaObject>
#include <QMouseEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPainter>
#include <QPixmap>
#include <QScreen>
#include <QSplashScreen>
#include <QString>
#include <QByteArray>
#include <cstring>
#include <QWidget>
#include <qsplashscreen.h>
#include "gen_qsplashscreen.h"
class MiqtVirtualQSplashScreen : public virtual QSplashScreen {
public:
MiqtVirtualQSplashScreen(): QSplashScreen() {};
MiqtVirtualQSplashScreen(QScreen* screen): QSplashScreen(screen) {};
MiqtVirtualQSplashScreen(const QPixmap& pixmap): QSplashScreen(pixmap) {};
MiqtVirtualQSplashScreen(const QPixmap& pixmap, Qt::WindowFlags f): QSplashScreen(pixmap, f) {};
MiqtVirtualQSplashScreen(QScreen* screen, const QPixmap& pixmap): QSplashScreen(screen, pixmap) {};
MiqtVirtualQSplashScreen(QScreen* screen, const QPixmap& pixmap, Qt::WindowFlags f): QSplashScreen(screen, pixmap, f) {};

virtual ~MiqtVirtualQSplashScreen() = default;
// cgo.Handle value for overwritten implementation
	intptr_t handle__MetaObject = 0;
// Subclass to allow providing a Go implementation
	virtual const QMetaObject* metaObject() const override {
if (handle__MetaObject == 0) {
return QSplashScreen::metaObject();
}
QMetaObject* callback_return_value = miqt_exec_callback_QSplashScreen_MetaObject(const_cast<MiqtVirtualQSplashScreen*>(this), handle__MetaObject);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	QMetaObject* virtualbase_MetaObject() const {

		return (QMetaObject*) QSplashScreen::metaObject();

	}
// cgo.Handle value for overwritten implementation
	intptr_t handle__Metacast = 0;
// Subclass to allow providing a Go implementation
	virtual void* qt_metacast(const char* param1) override {
if (handle__Metacast == 0) {
return QSplashScreen::qt_metacast(param1);
}
const char* sigval1 = (const char*) param1;

		void* callback_return_value = miqt_exec_callback_QSplashScreen_Metacast(this, handle__Metacast, sigval1);

		return callback_return_value;
	}
// Wrapper to allow calling protected method
	void* virtualbase_Metacast(const char* param1) {

		return QSplashScreen::qt_metacast(param1);

	}
};
QSplashScreen* QSplashScreen_new() {
return new MiqtVirtualQSplashScreen();
}
QSplashScreen* QSplashScreen_new2(QScreen* screen) {
return new MiqtVirtualQSplashScreen(screen);
}
QSplashScreen* QSplashScreen_new3(QPixmap* pixmap) {
return new MiqtVirtualQSplashScreen(*pixmap);
}
QSplashScreen* QSplashScreen_new4(QPixmap* pixmap, int f) {
return new MiqtVirtualQSplashScreen(*pixmap, static_cast<Qt::WindowFlags>(f));
}
QSplashScreen* QSplashScreen_new5(QScreen* screen, QPixmap* pixmap) {
return new MiqtVirtualQSplashScreen(screen, *pixmap);
}
QSplashScreen* QSplashScreen_new6(QScreen* screen, QPixmap* pixmap, int f) {
return new MiqtVirtualQSplashScreen(screen, *pixmap, static_cast<Qt::WindowFlags>(f));
}
void QSplashScreen_virtbase(QSplashScreen* src
, QWidget** outptr_QWidget
) {
*outptr_QWidget = static_cast<QWidget*>(src);
}
QMetaObject* QSplashScreen_MetaObject(const QSplashScreen* self) {
	return (QMetaObject*) self->metaObject();
}
void* QSplashScreen_Metacast(QSplashScreen* self, const char* param1) {
	return self->qt_metacast(param1);
}
struct miqt_string QSplashScreen_Tr(const char* s) {
	QString _ret = QSplashScreen::tr(s);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSplashScreen_SetPixmap(QSplashScreen* self, QPixmap* pixmap) {
	self->setPixmap(*pixmap);
}
QPixmap* QSplashScreen_Pixmap(const QSplashScreen* self) {
	return new QPixmap(self->pixmap());
}
void QSplashScreen_Finish(QSplashScreen* self, QWidget* w) {
	self->finish(w);
}
void QSplashScreen_Repaint(QSplashScreen* self) {
	self->repaint();
}
struct miqt_string QSplashScreen_Message(const QSplashScreen* self) {
	QString _ret = self->message();
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSplashScreen_ShowMessage(QSplashScreen* self, struct miqt_string message) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->showMessage(message_QString);
}
void QSplashScreen_ClearMessage(QSplashScreen* self) {
	self->clearMessage();
}
void QSplashScreen_MessageChanged(QSplashScreen* self, struct miqt_string message) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->messageChanged(message_QString);
}
void QSplashScreen_connect_MessageChanged(QSplashScreen* self, intptr_t slot) {
	MiqtVirtualQSplashScreen::connect(self, static_cast<void (QSplashScreen::*)(const QString&)>(&QSplashScreen::messageChanged), self, [=](const QString& message) {
		const QString message_ret = message;
		// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
		QByteArray message_b = message_ret.toUtf8();
		struct miqt_string message_ms;
		message_ms.len = message_b.length();
		message_ms.data = static_cast<char*>(malloc(message_ms.len));
		memcpy(message_ms.data, message_b.data(), message_ms.len);
		struct miqt_string sigval1 = message_ms;
		miqt_exec_callback_QSplashScreen_MessageChanged(slot, sigval1);
	});
}
struct miqt_string QSplashScreen_Tr2(const char* s, const char* c) {
	QString _ret = QSplashScreen::tr(s, c);
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
struct miqt_string QSplashScreen_Tr3(const char* s, const char* c, int n) {
	QString _ret = QSplashScreen::tr(s, c, static_cast<int>(n));
	// Convert QString from UTF-16 in C++ RAII memory to UTF-8 in manually-managed C memory
	QByteArray _b = _ret.toUtf8();
	struct miqt_string _ms;
	_ms.len = _b.length();
	_ms.data = static_cast<char*>(malloc(_ms.len));
	memcpy(_ms.data, _b.data(), _ms.len);
	return _ms;
}
void QSplashScreen_ShowMessage2(QSplashScreen* self, struct miqt_string message, int alignment) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->showMessage(message_QString, static_cast<int>(alignment));
}
void QSplashScreen_ShowMessage3(QSplashScreen* self, struct miqt_string message, int alignment, QColor* color) {
	QString message_QString = QString::fromUtf8(message.data, message.len);
	self->showMessage(message_QString, static_cast<int>(alignment), *color);
}
void QSplashScreen_override_virtual_MetaObject(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSplashScreen*>( (QSplashScreen*)(self) )->handle__MetaObject = slot;
}
QMetaObject* QSplashScreen_virtualbase_MetaObject(const void* self) {
	return ( (const MiqtVirtualQSplashScreen*)(self) )->virtualbase_MetaObject();
}
void QSplashScreen_override_virtual_Metacast(void* self, intptr_t slot) {
	dynamic_cast<MiqtVirtualQSplashScreen*>( (QSplashScreen*)(self) )->handle__Metacast = slot;
}
void* QSplashScreen_virtualbase_Metacast(void* self, const char* param1) {
	return ( (MiqtVirtualQSplashScreen*)(self) )->virtualbase_Metacast(param1);
}
void QSplashScreen_Delete(QSplashScreen* self, bool isSubclass) {
	if (isSubclass) {
		delete dynamic_cast<MiqtVirtualQSplashScreen*>( self );
	} else {
		delete self;
	}
}
